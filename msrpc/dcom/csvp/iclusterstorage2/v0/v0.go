package iclusterstorage2

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
)

var (
	// import guard
	GoPackage = "dcom/csvp"
)

var (
	// IClusterStorage2 interface identifier 12108a88-6858-4467-b92f-e6cf4568dfb6
	ClusterStorage2IID = &dcom.IID{Data1: 0x12108a88, Data2: 0x6858, Data3: 0x4467, Data4: []byte{0xb9, 0x2f, 0xe6, 0xcf, 0x45, 0x68, 0xdf, 0xb6}}
	// Syntax UUID
	ClusterStorage2SyntaxUUID = &uuid.UUID{TimeLow: 0x12108a88, TimeMid: 0x6858, TimeHiAndVersion: 0x4467, ClockSeqHiAndReserved: 0xb9, ClockSeqLow: 0x2f, Node: [6]uint8{0xe6, 0xcf, 0x45, 0x68, 0xdf, 0xb6}}
	// Syntax ID
	ClusterStorage2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ClusterStorage2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IClusterStorage2 interface.
type ClusterStorage2Client interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The CprepDiskRawRead method reads information directly from a single 512 byte sector
	// on a given disk.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+---------------------------------------+-----------------------------------------------------------------+
	//	|                RETURN                 |                                                                 |
	//	|              VALUE/CODE               |                           DESCRIPTION                           |
	//	|                                       |                                                                 |
	//	+---------------------------------------+-----------------------------------------------------------------+
	//	+---------------------------------------+-----------------------------------------------------------------+
	//	| 0x00000000 S_OK                       | The call was successful.                                        |
	//	+---------------------------------------+-----------------------------------------------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND       | The disk was not found.                                         |
	//	+---------------------------------------+-----------------------------------------------------------------+
	//	| 0x8007001E ERROR_READ_FAULT           | An attempt to read a buffer size larger than 512 was performed. |
	//	+---------------------------------------+-----------------------------------------------------------------+
	//	| 0x80070548 ERROR_INVALID_SERVER_STATE | The server's Prepare State is not Online.                       |
	//	+---------------------------------------+-----------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 3.
	RawRead(context.Context, *RawReadRequest, ...dcerpc.CallOption) (*RawReadResponse, error)

	// The CprepDiskRawWrite method writes information directly to a single 512 byte sector
	// on a given disk.
	//
	// Return Values: A signed, 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
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
	//	| 0x8007001D ERROR_WRITE_FAULT          | The size of the passed buffer was larger than 512 bytes. |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x80070548 ERROR_INVALID_SERVER_STATE | The server's Prepare State is not Online.                |
	//	+---------------------------------------+----------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 4.
	RawWrite(context.Context, *RawWriteRequest, ...dcerpc.CallOption) (*RawWriteResponse, error)

	// The CprepPrepareNode method prepares the server in an implementation-specific way
	// to execute the other methods in the interface. It also informs the client about version
	// information.
	//
	// This method is called before any other.
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
	PrepareNode(context.Context, *PrepareNodeRequest, ...dcerpc.CallOption) (*PrepareNodeResponse, error)

	// The CprepPrepareNodePhase2 method determines the number of disks accessible to the
	// system.
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
	//	| 0x80070548 ERROR_INVALID_SERVER_STATE | The server's Prepare State is not Online. |
	//	+---------------------------------------+-------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 6.
	PrepareNodePhase2(context.Context, *PrepareNodePhase2Request, ...dcerpc.CallOption) (*PrepareNodePhase2Response, error)

	// The CprepDiskGetProps method retrieves information about the configuration and status
	// of a given disk.
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
	// The opnum field value for this method is 7.
	GetProperties(context.Context, *GetPropertiesRequest, ...dcerpc.CallOption) (*GetPropertiesResponse, error)

	// Opnum8NotUsedOnWire operation.
	// Opnum8NotUsedOnWire

	// Opnum9NotUsedOnWire operation.
	// Opnum9NotUsedOnWire

	// Opnum10NotUsedOnWire operation.
	// Opnum10NotUsedOnWire

	// Opnum11NotUsedOnWire operation.
	// Opnum11NotUsedOnWire

	// The CprepDiskStopDefense method stops any implementation-specific method of maintaining
	// ownership of a disk.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+---------------------------------------+-------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                         |
	//	|              VALUE/CODE               |                               DESCRIPTION                               |
	//	|                                       |                                                                         |
	//	+---------------------------------------+-------------------------------------------------------------------------+
	//	+---------------------------------------+-------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                       | The call was successful.                                                |
	//	+---------------------------------------+-------------------------------------------------------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND       | The disk was not found.                                                 |
	//	+---------------------------------------+-------------------------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE        | The value of ClusPrepDisk.OwnedState is not equal to OwnedByThisServer. |
	//	+---------------------------------------+-------------------------------------------------------------------------+
	//	| 0x80070548 ERROR_INVALID_SERVER_STATE | The server's Prepare State is not Online.                               |
	//	+---------------------------------------+-------------------------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE        | The value of ClusPrepDisk.AttachedState is not equal to Attached.       |
	//	+---------------------------------------+-------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 12.
	StopDefense(context.Context, *StopDefenseRequest, ...dcerpc.CallOption) (*StopDefenseResponse, error)

	// The CprepDiskOnline method begins the transition of a ClusPrepDisk.OnlineState to
	// Online and then waits for the transition to complete.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+---------------------------------------+-------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                         |
	//	|              VALUE/CODE               |                               DESCRIPTION                               |
	//	|                                       |                                                                         |
	//	+---------------------------------------+-------------------------------------------------------------------------+
	//	+---------------------------------------+-------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                       | The call was successful.                                                |
	//	+---------------------------------------+-------------------------------------------------------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND       | The disk was not found.                                                 |
	//	+---------------------------------------+-------------------------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE        | The value of ClusPrepDisk.OwnedState is not equal to OwnedByThisServer. |
	//	+---------------------------------------+-------------------------------------------------------------------------+
	//	| 0x80070548 ERROR_INVALID_SERVER_STATE | The server's Prepare State is not Online.                               |
	//	+---------------------------------------+-------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 13.
	Online(context.Context, *OnlineRequest, ...dcerpc.CallOption) (*OnlineResponse, error)

	// The CprepDiskVerifyUnique method determines whether the same disk identifier is assigned
	// to more than one ClusPrepDisk in the attached state.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+---------------------------------------+---------------------------------------------------------------+
	//	|                RETURN                 |                                                               |
	//	|              VALUE/CODE               |                          DESCRIPTION                          |
	//	|                                       |                                                               |
	//	+---------------------------------------+---------------------------------------------------------------+
	//	+---------------------------------------+---------------------------------------------------------------+
	//	| 0x00000000 S_OK                       | The call was successful and only one ClusPrepDisk has the ID. |
	//	+---------------------------------------+---------------------------------------------------------------+
	//	| 0x800707DE ERROR_DUPLICATE_TAG        | There is more than one ClusPrepDisk with the given ID.        |
	//	+---------------------------------------+---------------------------------------------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND       | The disk was not found.                                       |
	//	+---------------------------------------+---------------------------------------------------------------+
	//	| 0x80070548 ERROR_INVALID_SERVER_STATE | The server's Prepare State is not Online.                     |
	//	+---------------------------------------+---------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 14.
	VerifyUnique(context.Context, *VerifyUniqueRequest, ...dcerpc.CallOption) (*VerifyUniqueResponse, error)

	// Opnum15NotUsedOnWire operation.
	// Opnum15NotUsedOnWire

	// Opnum16NotUsedOnWire operation.
	// Opnum16NotUsedOnWire

	// The CprepDiskWriteFileData method writes information to a file on a given partition
	// on a given disk.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	|                RETURN                 |                                                                     |
	//	|              VALUE/CODE               |                             DESCRIPTION                             |
	//	|                                       |                                                                     |
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 S_OK                       | The call was successful.                                            |
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE        | The ClusPrepDisk.OnlineState is not equal to Online.                |
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND       | The disk was not found or ulPartition cannot be mapped to a volume. |
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	| 0x800703ED ERROR_UNRECOGNIZED_VOLUME  | The volume does not contain a file system.                          |
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	| 0x80070548 ERROR_INVALID_SERVER_STATE | The server's Prepare State is not Online.                           |
	//	+---------------------------------------+---------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 17.
	WriteFileData(context.Context, *WriteFileDataRequest, ...dcerpc.CallOption) (*WriteFileDataResponse, error)

	// The CprepDiskVerifyFileData method verifies that the data in the file matches the
	// data passed to the method.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                       | The call was successful.                                                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE        | The ClusPrepDisk.OnlineState is not equal to Online.                             |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND       | The disk was not found or ulPartition cannot be mapped to a volume. The file     |
	//	|                                       | does not exist.                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800703ED ERROR_UNRECOGNIZED_VOLUME  | The volume does not contain a file system.                                       |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070548 ERROR_INVALID_SERVER_STATE | The server's Prepare State is not Online.                                        |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 18.
	VerifyFileData(context.Context, *VerifyFileDataRequest, ...dcerpc.CallOption) (*VerifyFileDataResponse, error)

	// The CprepDiskDeleteFile method deletes a file on a given partition on a given disk.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	|                RETURN                 |                                                                     |
	//	|              VALUE/CODE               |                             DESCRIPTION                             |
	//	|                                       |                                                                     |
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 S_OK                       | The call was successful.                                            |
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE        | The ClusPrepDisk.OnlineState is not equal to Online.                |
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND       | The disk was not found or ulPartition cannot be mapped to a volume. |
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	| 0x800703ED ERROR_UNRECOGNIZED_VOLUME  | The volume does not contain a file system.                          |
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	| 0x80070548 ERROR_INVALID_SERVER_STATE | The server's Prepare State is not Online.                           |
	//	+---------------------------------------+---------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 19.
	DeleteFile(context.Context, *DeleteFileRequest, ...dcerpc.CallOption) (*DeleteFileResponse, error)

	// The CprepDiskOffline method begins the transition of a ClusPrepDisk.OnlineState to
	// Not Online and then waits for the transition to complete.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+---------------------------------------+---------------------------------------------------------------+
	//	|                RETURN                 |                                                               |
	//	|              VALUE/CODE               |                          DESCRIPTION                          |
	//	|                                       |                                                               |
	//	+---------------------------------------+---------------------------------------------------------------+
	//	+---------------------------------------+---------------------------------------------------------------+
	//	| 0x00000000 S_OK                       | The call was successful.                                      |
	//	+---------------------------------------+---------------------------------------------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND       | The disk was not found.                                       |
	//	+---------------------------------------+---------------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE        | The value of ClusPrepDisk.OnlineState is not equal to Online. |
	//	+---------------------------------------+---------------------------------------------------------------+
	//	| 0x80070548 ERROR_INVALID_SERVER_STATE | The server's Prepare State is not Online.                     |
	//	+---------------------------------------+---------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 20.
	Offline(context.Context, *OfflineRequest, ...dcerpc.CallOption) (*OfflineResponse, error)

	// Opnum21NotUsedOnWire operation.
	// Opnum21NotUsedOnWire

	// The CprepDiskGetUniqueIds method returns device ID data about the ClusPrepDisk.
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
	//	| 0x8007007A ERROR_INSUFFICIENT_BUFFER  | pbData is not large enough.               |
	//	+---------------------------------------+-------------------------------------------+
	//	| 0x80070032 ERROR_NOT_SUPPORTED        | The disk does not support device ID data. |
	//	+---------------------------------------+-------------------------------------------+
	//	| 0x80070548 ERROR_INVALID_SERVER_STATE | The server's Prepare State is not Online. |
	//	+---------------------------------------+-------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 22.
	GetUniqueIDs(context.Context, *GetUniqueIDsRequest, ...dcerpc.CallOption) (*GetUniqueIDsResponse, error)

	// The CprepDiskAttach method offers implementations an opportunity to do disk-specific
	// setup before processing is done on a disk.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
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
	//	| 0x80070490 ERROR_NOT_FOUND            | The disk was not found.                   |
	//	+---------------------------------------+-------------------------------------------+
	//	| 0x80070548 ERROR_INVALID_SERVER_STATE | The server's Prepare State is not Online. |
	//	+---------------------------------------+-------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 23.
	Attach(context.Context, *AttachRequest, ...dcerpc.CallOption) (*AttachResponse, error)

	// The CprepDiskPRArbitrate method establishes ownership of a ClusPrepDisk.
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
	// The opnum field value for this method is 24.
	PRArbitrate(context.Context, *PRArbitrateRequest, ...dcerpc.CallOption) (*PRArbitrateResponse, error)

	// The CprepDiskPRRegister method performs a SCSI PERSISTENT RESERVE OUT command (see
	// [SPC-3] section 6.12) with a REGISTER AND IGNORE EXISTING KEY action.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+---------------------------------------+-------------------------------------------------------------------+
	//	|                RETURN                 |                                                                   |
	//	|              VALUE/CODE               |                            DESCRIPTION                            |
	//	|                                       |                                                                   |
	//	+---------------------------------------+-------------------------------------------------------------------+
	//	+---------------------------------------+-------------------------------------------------------------------+
	//	| 0x00000000 S_OK                       | The call was successful.                                          |
	//	+---------------------------------------+-------------------------------------------------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND       | The disk was not found.                                           |
	//	+---------------------------------------+-------------------------------------------------------------------+
	//	| 0x80070548 ERROR_INVALID_SERVER_STATE | The server's Prepare State is not Online.                         |
	//	+---------------------------------------+-------------------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE        | The value of ClusPrepDisk.AttachedState is not equal to Attached. |
	//	+---------------------------------------+-------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 25.
	PRRegister(context.Context, *PRRegisterRequest, ...dcerpc.CallOption) (*PRRegisterResponse, error)

	// The CprepDiskPRUnRegister method performs a SCSI PERSISTENT RESERVE OUT command (see
	// [SPC-3] section 6.12) with a REGISTER AND IGNORE EXISTING KEY action with a key of
	// 0.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+---------------------------------------+-------------------------------------------------------------------+
	//	|                RETURN                 |                                                                   |
	//	|              VALUE/CODE               |                            DESCRIPTION                            |
	//	|                                       |                                                                   |
	//	+---------------------------------------+-------------------------------------------------------------------+
	//	+---------------------------------------+-------------------------------------------------------------------+
	//	| 0x00000000 S_OK                       | The call was successful.                                          |
	//	+---------------------------------------+-------------------------------------------------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND       | The disk was not found.                                           |
	//	+---------------------------------------+-------------------------------------------------------------------+
	//	| 0x80070548 ERROR_INVALID_SERVER_STATE | The server's Prepare State is not Online.                         |
	//	+---------------------------------------+-------------------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE        | The value of ClusPrepDisk.AttachedState is not equal to Attached. |
	//	+---------------------------------------+-------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 26.
	PRUnregister(context.Context, *PRUnregisterRequest, ...dcerpc.CallOption) (*PRUnregisterResponse, error)

	// The CprepDiskPRReserve method performs a SCSI PERSISTENT RESERVE OUT command (see
	// [SPC-3] section 6.12) with a RESERVE action.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+---------------------------------------+-------------------------------------------------------------------+
	//	|                RETURN                 |                                                                   |
	//	|              VALUE/CODE               |                            DESCRIPTION                            |
	//	|                                       |                                                                   |
	//	+---------------------------------------+-------------------------------------------------------------------+
	//	+---------------------------------------+-------------------------------------------------------------------+
	//	| 0x00000000 S_OK                       | The call was successful.                                          |
	//	+---------------------------------------+-------------------------------------------------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND       | The disk was not found.                                           |
	//	+---------------------------------------+-------------------------------------------------------------------+
	//	| 0x80070548 ERROR_INVALID_SERVER_STATE | The server's Prepare State is not Online.                         |
	//	+---------------------------------------+-------------------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE        | The value of ClusPrepDisk.AttachedState is not equal to Attached. |
	//	+---------------------------------------+-------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 27.
	PRReserve(context.Context, *PRReserveRequest, ...dcerpc.CallOption) (*PRReserveResponse, error)

	// The CprepDiskPRRelease method performs a SCSI PERSISTENT RESERVE OUT command (see
	// [SPC-3] section 6.12) with a RELEASE action.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+---------------------------------------+-------------------------------------------------------------------+
	//	|                RETURN                 |                                                                   |
	//	|              VALUE/CODE               |                            DESCRIPTION                            |
	//	|                                       |                                                                   |
	//	+---------------------------------------+-------------------------------------------------------------------+
	//	+---------------------------------------+-------------------------------------------------------------------+
	//	| 0x00000000 S_OK                       | The call was successful.                                          |
	//	+---------------------------------------+-------------------------------------------------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND       | The disk was not found.                                           |
	//	+---------------------------------------+-------------------------------------------------------------------+
	//	| 0x80070548 ERROR_INVALID_SERVER_STATE | The server's Prepare State is not Online.                         |
	//	+---------------------------------------+-------------------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE        | The value of ClusPrepDisk.AttachedState is not equal to Attached. |
	//	+---------------------------------------+-------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 28.
	PRRelease(context.Context, *PRReleaseRequest, ...dcerpc.CallOption) (*PRReleaseResponse, error)

	// The CprepDiskDiskPartitionIsNtfs method determines whether the file system on a given
	// partition on a given disk is NTFS.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	|                RETURN                 |                                                                     |
	//	|              VALUE/CODE               |                             DESCRIPTION                             |
	//	|                                       |                                                                     |
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 S_OK                       | The call was successful.                                            |
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	| 0x80070022 ERROR_WRONG_DISK           | The partition on the disk has a file system other than NTFS.        |
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE        | The ClusPrepDisk.OnlineState value is not equal to Online.          |
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND       | The disk was not found or ulPartition cannot be mapped to a volume. |
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	| 0x80070548 ERROR_INVALID_SERVER_STATE | The server's Prepare State is not Online.                           |
	//	+---------------------------------------+---------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 29.
	DiskPartitionIsNTFS(context.Context, *DiskPartitionIsNTFSRequest, ...dcerpc.CallOption) (*DiskPartitionIsNTFSResponse, error)

	// The CprepDiskGetArbSectors method returns two sectors on the disk that can be used
	// as a "scratch pad" for raw reads/writes.
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
	// The opnum field value for this method is 30.
	GetArbSectors(context.Context, *GetArbSectorsRequest, ...dcerpc.CallOption) (*GetArbSectorsResponse, error)

	// The CprepDiskIsPRPresent method determines whether there are any PERSISTENT RESERVE
	// reservations on the disk.
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
	// The opnum field value for this method is 31.
	IsPRPresent(context.Context, *IsPRPresentRequest, ...dcerpc.CallOption) (*IsPRPresentResponse, error)

	// The CprepDiskPRPreempt method performs a SCSI PERSISTENT RESERVE OUT command (see
	// [SPC-3] section 6.12) with a PREEMPT action.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
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
	// The opnum field value for this method is 32.
	PRPreempt(context.Context, *PRPreemptRequest, ...dcerpc.CallOption) (*PRPreemptResponse, error)

	// The CprepDiskPRClear method performs a SCSI PERSISTENT RESERVE OUT command (see [SPC-3]
	// section 6.12) with a CLEAR action.
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
	// The opnum field value for this method is 33.
	PRClear(context.Context, *PRClearRequest, ...dcerpc.CallOption) (*PRClearResponse, error)

	// The CprepDiskIsOnline method reports whether the ClusPrepDisk, identified by the
	// DiskId parameter, has ClusPrepDisk.OnlineState equal to Online.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+---------------------------------------+--------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                          |
	//	|              VALUE/CODE               |                               DESCRIPTION                                |
	//	|                                       |                                                                          |
	//	+---------------------------------------+--------------------------------------------------------------------------+
	//	+---------------------------------------+--------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                       | The call was successful and ClusPrepDisk.OnlineState is equal to Online. |
	//	+---------------------------------------+--------------------------------------------------------------------------+
	//	| 0x80070015 ERROR_NOT_READY            | ClusPrepDisk.OnlineState is not equal to Online.                         |
	//	+---------------------------------------+--------------------------------------------------------------------------+
	//	| 0x80070548 ERROR_INVALID_SERVER_STATE | The server's Prepare State is not Online.                                |
	//	+---------------------------------------+--------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 34.
	IsOnline(context.Context, *IsOnlineRequest, ...dcerpc.CallOption) (*IsOnlineResponse, error)

	// The CprepDiskSetOnline method starts the process of transitioning ClusPrepDisk.OnlineState
	// to the Online state.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+---------------------------------------+-------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                         |
	//	|              VALUE/CODE               |                               DESCRIPTION                               |
	//	|                                       |                                                                         |
	//	+---------------------------------------+-------------------------------------------------------------------------+
	//	+---------------------------------------+-------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                       | The call was successful.                                                |
	//	+---------------------------------------+-------------------------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE        | The value of ClusPrepDisk.OwnedState is not equal to OwnedByThisServer. |
	//	+---------------------------------------+-------------------------------------------------------------------------+
	//	| 0x80070548 ERROR_INVALID_SERVER_STATE | The server's Prepare State is not Online.                               |
	//	+---------------------------------------+-------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 35.
	SetOnline(context.Context, *SetOnlineRequest, ...dcerpc.CallOption) (*SetOnlineResponse, error)

	// The CprepDiskGetFSName method returns the name of the file system on a given partition
	// on a given disk.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	|                RETURN                 |                                                                     |
	//	|              VALUE/CODE               |                             DESCRIPTION                             |
	//	|                                       |                                                                     |
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 S_OK                       | The call was successful.                                            |
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND       | The disk was not found or ulPartition cannot be mapped to a volume. |
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE        | The ClusPrepDisk.OnlineState is not equal to Online.                |
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	| 0x80070548 ERROR_INVALID_SERVER_STATE | The server's Prepare State is not Online.                           |
	//	+---------------------------------------+---------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 36.
	GetFSName(context.Context, *GetFSNameRequest, ...dcerpc.CallOption) (*GetFSNameResponse, error)

	// The CprepDiskIsReadable method determines whether the disk data on the disk can be
	// successfully read.
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
	// The opnum field value for this method is 37.
	IsReadable(context.Context, *IsReadableRequest, ...dcerpc.CallOption) (*IsReadableResponse, error)

	// The CprepDiskGetDsms method returns the DSMs installed on the system.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+----------------------------+-----------------------------------------------------------+
	//	|           RETURN           |                                                           |
	//	|         VALUE/CODE         |                        DESCRIPTION                        |
	//	|                            |                                                           |
	//	+----------------------------+-----------------------------------------------------------+
	//	+----------------------------+-----------------------------------------------------------+
	//	| 0x00000000 S_OK            | The call was successful.                                  |
	//	+----------------------------+-----------------------------------------------------------+
	//	| 0x800700EA ERROR_MORE_DATA | RegisteredDsms was not large enough to hold all the data. |
	//	+----------------------------+-----------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 38.
	GetDSMs(context.Context, *GetDSMsRequest, ...dcerpc.CallOption) (*GetDSMsResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ClusterStorage2Client
}

type xxx_DefaultClusterStorage2Client struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultClusterStorage2Client) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultClusterStorage2Client) RawRead(ctx context.Context, in *RawReadRequest, opts ...dcerpc.CallOption) (*RawReadResponse, error) {
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
	out := &RawReadResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage2Client) RawWrite(ctx context.Context, in *RawWriteRequest, opts ...dcerpc.CallOption) (*RawWriteResponse, error) {
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
	out := &RawWriteResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage2Client) PrepareNode(ctx context.Context, in *PrepareNodeRequest, opts ...dcerpc.CallOption) (*PrepareNodeResponse, error) {
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
	out := &PrepareNodeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage2Client) PrepareNodePhase2(ctx context.Context, in *PrepareNodePhase2Request, opts ...dcerpc.CallOption) (*PrepareNodePhase2Response, error) {
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
	out := &PrepareNodePhase2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage2Client) GetProperties(ctx context.Context, in *GetPropertiesRequest, opts ...dcerpc.CallOption) (*GetPropertiesResponse, error) {
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
	out := &GetPropertiesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage2Client) StopDefense(ctx context.Context, in *StopDefenseRequest, opts ...dcerpc.CallOption) (*StopDefenseResponse, error) {
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
	out := &StopDefenseResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage2Client) Online(ctx context.Context, in *OnlineRequest, opts ...dcerpc.CallOption) (*OnlineResponse, error) {
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
	out := &OnlineResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage2Client) VerifyUnique(ctx context.Context, in *VerifyUniqueRequest, opts ...dcerpc.CallOption) (*VerifyUniqueResponse, error) {
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
	out := &VerifyUniqueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage2Client) WriteFileData(ctx context.Context, in *WriteFileDataRequest, opts ...dcerpc.CallOption) (*WriteFileDataResponse, error) {
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
	out := &WriteFileDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage2Client) VerifyFileData(ctx context.Context, in *VerifyFileDataRequest, opts ...dcerpc.CallOption) (*VerifyFileDataResponse, error) {
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
	out := &VerifyFileDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage2Client) DeleteFile(ctx context.Context, in *DeleteFileRequest, opts ...dcerpc.CallOption) (*DeleteFileResponse, error) {
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
	out := &DeleteFileResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage2Client) Offline(ctx context.Context, in *OfflineRequest, opts ...dcerpc.CallOption) (*OfflineResponse, error) {
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
	out := &OfflineResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage2Client) GetUniqueIDs(ctx context.Context, in *GetUniqueIDsRequest, opts ...dcerpc.CallOption) (*GetUniqueIDsResponse, error) {
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
	out := &GetUniqueIDsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage2Client) Attach(ctx context.Context, in *AttachRequest, opts ...dcerpc.CallOption) (*AttachResponse, error) {
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
	out := &AttachResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage2Client) PRArbitrate(ctx context.Context, in *PRArbitrateRequest, opts ...dcerpc.CallOption) (*PRArbitrateResponse, error) {
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
	out := &PRArbitrateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage2Client) PRRegister(ctx context.Context, in *PRRegisterRequest, opts ...dcerpc.CallOption) (*PRRegisterResponse, error) {
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
	out := &PRRegisterResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage2Client) PRUnregister(ctx context.Context, in *PRUnregisterRequest, opts ...dcerpc.CallOption) (*PRUnregisterResponse, error) {
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
	out := &PRUnregisterResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage2Client) PRReserve(ctx context.Context, in *PRReserveRequest, opts ...dcerpc.CallOption) (*PRReserveResponse, error) {
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
	out := &PRReserveResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage2Client) PRRelease(ctx context.Context, in *PRReleaseRequest, opts ...dcerpc.CallOption) (*PRReleaseResponse, error) {
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
	out := &PRReleaseResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage2Client) DiskPartitionIsNTFS(ctx context.Context, in *DiskPartitionIsNTFSRequest, opts ...dcerpc.CallOption) (*DiskPartitionIsNTFSResponse, error) {
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
	out := &DiskPartitionIsNTFSResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage2Client) GetArbSectors(ctx context.Context, in *GetArbSectorsRequest, opts ...dcerpc.CallOption) (*GetArbSectorsResponse, error) {
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
	out := &GetArbSectorsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage2Client) IsPRPresent(ctx context.Context, in *IsPRPresentRequest, opts ...dcerpc.CallOption) (*IsPRPresentResponse, error) {
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
	out := &IsPRPresentResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage2Client) PRPreempt(ctx context.Context, in *PRPreemptRequest, opts ...dcerpc.CallOption) (*PRPreemptResponse, error) {
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
	out := &PRPreemptResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage2Client) PRClear(ctx context.Context, in *PRClearRequest, opts ...dcerpc.CallOption) (*PRClearResponse, error) {
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
	out := &PRClearResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage2Client) IsOnline(ctx context.Context, in *IsOnlineRequest, opts ...dcerpc.CallOption) (*IsOnlineResponse, error) {
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
	out := &IsOnlineResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage2Client) SetOnline(ctx context.Context, in *SetOnlineRequest, opts ...dcerpc.CallOption) (*SetOnlineResponse, error) {
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
	out := &SetOnlineResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage2Client) GetFSName(ctx context.Context, in *GetFSNameRequest, opts ...dcerpc.CallOption) (*GetFSNameResponse, error) {
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
	out := &GetFSNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage2Client) IsReadable(ctx context.Context, in *IsReadableRequest, opts ...dcerpc.CallOption) (*IsReadableResponse, error) {
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
	out := &IsReadableResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage2Client) GetDSMs(ctx context.Context, in *GetDSMsRequest, opts ...dcerpc.CallOption) (*GetDSMsResponse, error) {
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
	out := &GetDSMsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultClusterStorage2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultClusterStorage2Client) IPID(ctx context.Context, ipid *dcom.IPID) ClusterStorage2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultClusterStorage2Client{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewClusterStorage2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ClusterStorage2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ClusterStorage2SyntaxV0_0))...)
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
	return &xxx_DefaultClusterStorage2Client{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_RawReadOperation structure represents the CprepDiskRawRead operation
type xxx_RawReadOperation struct {
	This           *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID         *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	Sector         uint32              `idl:"name:ulSector" json:"sector"`
	DataLength     uint32              `idl:"name:cbData" json:"data_length"`
	Data           []byte              `idl:"name:pbData;size_is:(cbData);length_is:(pcbDataRead)" json:"data"`
	DataReadLength uint32              `idl:"name:pcbDataRead" json:"data_read_length"`
	Latency        uint32              `idl:"name:ulLatency" json:"latency"`
	Return         int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_RawReadOperation) OpNum() int { return 3 }

func (o *xxx_RawReadOperation) OpName() string { return "/IClusterStorage2/v0/CprepDiskRawRead" }

func (o *xxx_RawReadOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RawReadOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ulSector {in} (1:(uint32))
	{
		if err := w.WriteData(o.Sector); err != nil {
			return err
		}
	}
	// cbData {in} (1:(uint32))
	{
		if err := w.WriteData(o.DataLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RawReadOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ulSector {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Sector); err != nil {
			return err
		}
	}
	// cbData {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DataLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RawReadOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Data != nil && o.DataReadLength == 0 {
		o.DataReadLength = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RawReadOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbData {out} (1:{pointer=ref}*(1)[dim:0,size_is=cbData,length_is=pcbDataRead](uint8))
	{
		dimSize1 := uint64(o.DataLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.DataReadLength)
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
	}
	// pcbDataRead {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.DataReadLength); err != nil {
			return err
		}
	}
	// ulLatency {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Latency); err != nil {
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

func (o *xxx_RawReadOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbData {out} (1:{pointer=ref}*(1)[dim:0,size_is=cbData,length_is=pcbDataRead](uint8))
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
	// pcbDataRead {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.DataReadLength); err != nil {
			return err
		}
	}
	// ulLatency {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Latency); err != nil {
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

// RawReadRequest structure represents the CprepDiskRawRead operation request
type RawReadRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk representing the disk that holds the sector
	// from which to read.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	// ulSector: The sector number to read from.
	Sector uint32 `idl:"name:ulSector" json:"sector"`
	// cbData: The size, in bytes, of the buffer pbData.
	DataLength uint32 `idl:"name:cbData" json:"data_length"`
}

func (o *RawReadRequest) xxx_ToOp(ctx context.Context) *xxx_RawReadOperation {
	if o == nil {
		return &xxx_RawReadOperation{}
	}
	return &xxx_RawReadOperation{
		This:       o.This,
		DiskID:     o.DiskID,
		Sector:     o.Sector,
		DataLength: o.DataLength,
	}
}

func (o *RawReadRequest) xxx_FromOp(ctx context.Context, op *xxx_RawReadOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
	o.Sector = op.Sector
	o.DataLength = op.DataLength
}
func (o *RawReadRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RawReadRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RawReadOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RawReadResponse structure represents the CprepDiskRawRead operation response
type RawReadResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbData: The data to read from the disk.
	Data []byte `idl:"name:pbData;size_is:(cbData);length_is:(pcbDataRead)" json:"data"`
	// pcbDataRead: On successful completion, the server MUST set this to cbData. Otherwise
	// the client MUST ignore this value.
	DataReadLength uint32 `idl:"name:pcbDataRead" json:"data_read_length"`
	// ulLatency: The time, in milliseconds, that the read took to be performed.
	Latency uint32 `idl:"name:ulLatency" json:"latency"`
	// Return: The CprepDiskRawRead return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RawReadResponse) xxx_ToOp(ctx context.Context) *xxx_RawReadOperation {
	if o == nil {
		return &xxx_RawReadOperation{}
	}
	return &xxx_RawReadOperation{
		That:           o.That,
		Data:           o.Data,
		DataReadLength: o.DataReadLength,
		Latency:        o.Latency,
		Return:         o.Return,
	}
}

func (o *RawReadResponse) xxx_FromOp(ctx context.Context, op *xxx_RawReadOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Data = op.Data
	o.DataReadLength = op.DataReadLength
	o.Latency = op.Latency
	o.Return = op.Return
}
func (o *RawReadResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RawReadResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RawReadOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RawWriteOperation structure represents the CprepDiskRawWrite operation
type xxx_RawWriteOperation struct {
	This              *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID            *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	Sector            uint32              `idl:"name:ulSector" json:"sector"`
	DataLength        uint32              `idl:"name:cbData" json:"data_length"`
	Data              []byte              `idl:"name:pbData;size_is:(cbData)" json:"data"`
	DataWrittenLength uint32              `idl:"name:pcbDataWritten" json:"data_written_length"`
	Latency           uint32              `idl:"name:ulLatency" json:"latency"`
	Return            int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_RawWriteOperation) OpNum() int { return 4 }

func (o *xxx_RawWriteOperation) OpName() string { return "/IClusterStorage2/v0/CprepDiskRawWrite" }

func (o *xxx_RawWriteOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
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

func (o *xxx_RawWriteOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ulSector {in} (1:(uint32))
	{
		if err := w.WriteData(o.Sector); err != nil {
			return err
		}
	}
	// cbData {in} (1:(uint32))
	{
		if err := w.WriteData(o.DataLength); err != nil {
			return err
		}
	}
	// pbData {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbData](uint8))
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
	return nil
}

func (o *xxx_RawWriteOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ulSector {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Sector); err != nil {
			return err
		}
	}
	// cbData {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DataLength); err != nil {
			return err
		}
	}
	// pbData {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbData](uint8))
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
	return nil
}

func (o *xxx_RawWriteOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RawWriteOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pcbDataWritten {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.DataWrittenLength); err != nil {
			return err
		}
	}
	// ulLatency {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Latency); err != nil {
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

func (o *xxx_RawWriteOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pcbDataWritten {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.DataWrittenLength); err != nil {
			return err
		}
	}
	// ulLatency {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Latency); err != nil {
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

// RawWriteRequest structure represents the CprepDiskRawWrite operation request
type RawWriteRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk representing the disk that holds the sector
	// to which to write.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	// ulSector: The sector number to write to.
	Sector uint32 `idl:"name:ulSector" json:"sector"`
	// cbData: The size, in bytes of the buffer pbData.
	DataLength uint32 `idl:"name:cbData" json:"data_length"`
	// pbData: The data to write to the disk.
	Data []byte `idl:"name:pbData;size_is:(cbData)" json:"data"`
}

func (o *RawWriteRequest) xxx_ToOp(ctx context.Context) *xxx_RawWriteOperation {
	if o == nil {
		return &xxx_RawWriteOperation{}
	}
	return &xxx_RawWriteOperation{
		This:       o.This,
		DiskID:     o.DiskID,
		Sector:     o.Sector,
		DataLength: o.DataLength,
		Data:       o.Data,
	}
}

func (o *RawWriteRequest) xxx_FromOp(ctx context.Context, op *xxx_RawWriteOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
	o.Sector = op.Sector
	o.DataLength = op.DataLength
	o.Data = op.Data
}
func (o *RawWriteRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RawWriteRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RawWriteOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RawWriteResponse structure represents the CprepDiskRawWrite operation response
type RawWriteResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pcbDataWritten: If the CprepDiskRawWrite method is successful, the server MUST set
	// this value to 512. If an error occurs, the server MUST set pcbDataWritten to zero.
	DataWrittenLength uint32 `idl:"name:pcbDataWritten" json:"data_written_length"`
	// ulLatency: The time, in milliseconds, that the write took to be performed.
	Latency uint32 `idl:"name:ulLatency" json:"latency"`
	// Return: The CprepDiskRawWrite return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RawWriteResponse) xxx_ToOp(ctx context.Context) *xxx_RawWriteOperation {
	if o == nil {
		return &xxx_RawWriteOperation{}
	}
	return &xxx_RawWriteOperation{
		That:              o.That,
		DataWrittenLength: o.DataWrittenLength,
		Latency:           o.Latency,
		Return:            o.Return,
	}
}

func (o *RawWriteResponse) xxx_FromOp(ctx context.Context, op *xxx_RawWriteOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DataWrittenLength = op.DataWrittenLength
	o.Latency = op.Latency
	o.Return = op.Return
}
func (o *RawWriteResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RawWriteResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RawWriteOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PrepareNodeOperation structure represents the CprepPrepareNode operation
type xxx_PrepareNodeOperation struct {
	This                  *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                  *dcom.ORPCThat `idl:"name:That" json:"that"`
	MajorVersion          uint32         `idl:"name:pulMajorVersion" json:"major_version"`
	MinorVersion          uint32         `idl:"name:pulMinorVersion" json:"minor_version"`
	ClusterPrepareVersion uint32         `idl:"name:pdwCPrepVersion" json:"cluster_prepare_version"`
	Return                int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_PrepareNodeOperation) OpNum() int { return 5 }

func (o *xxx_PrepareNodeOperation) OpName() string { return "/IClusterStorage2/v0/CprepPrepareNode" }

func (o *xxx_PrepareNodeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PrepareNodeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PrepareNodeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_PrepareNodeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PrepareNodeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pulMajorVersion {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.MajorVersion); err != nil {
			return err
		}
	}
	// pulMinorVersion {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.MinorVersion); err != nil {
			return err
		}
	}
	// pdwCPrepVersion {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.ClusterPrepareVersion); err != nil {
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

func (o *xxx_PrepareNodeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pulMajorVersion {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.MajorVersion); err != nil {
			return err
		}
	}
	// pulMinorVersion {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.MinorVersion); err != nil {
			return err
		}
	}
	// pdwCPrepVersion {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.ClusterPrepareVersion); err != nil {
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

// PrepareNodeRequest structure represents the CprepPrepareNode operation request
type PrepareNodeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *PrepareNodeRequest) xxx_ToOp(ctx context.Context) *xxx_PrepareNodeOperation {
	if o == nil {
		return &xxx_PrepareNodeOperation{}
	}
	return &xxx_PrepareNodeOperation{
		This: o.This,
	}
}

func (o *PrepareNodeRequest) xxx_FromOp(ctx context.Context, op *xxx_PrepareNodeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *PrepareNodeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PrepareNodeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PrepareNodeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PrepareNodeResponse structure represents the CprepPrepareNode operation response
type PrepareNodeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pulMajorVersion: The server MUST set this to the operating system major version.
	MajorVersion uint32 `idl:"name:pulMajorVersion" json:"major_version"`
	// pulMinorVersion: The server MUST set this to the operating system minor version.
	MinorVersion uint32 `idl:"name:pulMinorVersion" json:"minor_version"`
	// pdwCPrepVersion: The client MUST ignore this value.
	ClusterPrepareVersion uint32 `idl:"name:pdwCPrepVersion" json:"cluster_prepare_version"`
	// Return: The CprepPrepareNode return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PrepareNodeResponse) xxx_ToOp(ctx context.Context) *xxx_PrepareNodeOperation {
	if o == nil {
		return &xxx_PrepareNodeOperation{}
	}
	return &xxx_PrepareNodeOperation{
		That:                  o.That,
		MajorVersion:          o.MajorVersion,
		MinorVersion:          o.MinorVersion,
		ClusterPrepareVersion: o.ClusterPrepareVersion,
		Return:                o.Return,
	}
}

func (o *PrepareNodeResponse) xxx_FromOp(ctx context.Context, op *xxx_PrepareNodeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MajorVersion = op.MajorVersion
	o.MinorVersion = op.MinorVersion
	o.ClusterPrepareVersion = op.ClusterPrepareVersion
	o.Return = op.Return
}
func (o *PrepareNodeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PrepareNodeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PrepareNodeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PrepareNodePhase2Operation structure represents the CprepPrepareNodePhase2 operation
type xxx_PrepareNodePhase2Operation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	Flags       uint32         `idl:"name:Flags" json:"flags"`
	DisksLength uint32         `idl:"name:pulNumDisks" json:"disks_length"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_PrepareNodePhase2Operation) OpNum() int { return 6 }

func (o *xxx_PrepareNodePhase2Operation) OpName() string {
	return "/IClusterStorage2/v0/CprepPrepareNodePhase2"
}

func (o *xxx_PrepareNodePhase2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PrepareNodePhase2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Flags {in} (1:(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PrepareNodePhase2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Flags {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PrepareNodePhase2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PrepareNodePhase2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pulNumDisks {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.DisksLength); err != nil {
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

func (o *xxx_PrepareNodePhase2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pulNumDisks {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.DisksLength); err != nil {
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

// PrepareNodePhase2Request structure represents the CprepPrepareNodePhase2 operation request
type PrepareNodePhase2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// Flags:   The client SHOULD<14> pass in one of the following values:
	//
	//	+------------------------------------------+-------------------------------------------------------------------+
	//	|                                          |                                                                   |
	//	|                  VALUE                   |                              MEANING                              |
	//	|                                          |                                                                   |
	//	+------------------------------------------+-------------------------------------------------------------------+
	//	+------------------------------------------+-------------------------------------------------------------------+
	//	| ForceOfflineNonClusteredDisks 0x00000001 | When set, the spaces on the nonclustered pool are force detached. |
	//	+------------------------------------------+-------------------------------------------------------------------+
	//	| SkipNonClusteredPools 0x00000002         | When set, the nonclustered pools are skipped.                     |
	//	+------------------------------------------+-------------------------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
}

func (o *PrepareNodePhase2Request) xxx_ToOp(ctx context.Context) *xxx_PrepareNodePhase2Operation {
	if o == nil {
		return &xxx_PrepareNodePhase2Operation{}
	}
	return &xxx_PrepareNodePhase2Operation{
		This:  o.This,
		Flags: o.Flags,
	}
}

func (o *PrepareNodePhase2Request) xxx_FromOp(ctx context.Context, op *xxx_PrepareNodePhase2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Flags = op.Flags
}
func (o *PrepareNodePhase2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PrepareNodePhase2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PrepareNodePhase2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PrepareNodePhase2Response structure represents the CprepPrepareNodePhase2 operation response
type PrepareNodePhase2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pulNumDisks: The number of disks accessible to the system.
	DisksLength uint32 `idl:"name:pulNumDisks" json:"disks_length"`
	// Return: The CprepPrepareNodePhase2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PrepareNodePhase2Response) xxx_ToOp(ctx context.Context) *xxx_PrepareNodePhase2Operation {
	if o == nil {
		return &xxx_PrepareNodePhase2Operation{}
	}
	return &xxx_PrepareNodePhase2Operation{
		That:        o.That,
		DisksLength: o.DisksLength,
		Return:      o.Return,
	}
}

func (o *PrepareNodePhase2Response) xxx_FromOp(ctx context.Context, op *xxx_PrepareNodePhase2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DisksLength = op.DisksLength
	o.Return = op.Return
}
func (o *PrepareNodePhase2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PrepareNodePhase2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PrepareNodePhase2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPropertiesOperation structure represents the CprepDiskGetProps operation
type xxx_GetPropertiesOperation struct {
	This           *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat       `idl:"name:That" json:"that"`
	DiskID         *csvp.ClusterDiskID  `idl:"name:DiskId" json:"disk_id"`
	DiskProperties *csvp.DiskProperties `idl:"name:DiskProps" json:"disk_properties"`
	Return         int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPropertiesOperation) OpNum() int { return 7 }

func (o *xxx_GetPropertiesOperation) OpName() string { return "/IClusterStorage2/v0/CprepDiskGetProps" }

func (o *xxx_GetPropertiesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPropertiesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPropertiesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// DiskProps {out} (1:{pointer=ref}*(1))(2:{alias=DISK_PROPS}(struct))
	{
		if o.DiskProperties != nil {
			if err := o.DiskProperties.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csvp.DiskProperties{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
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

func (o *xxx_GetPropertiesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// DiskProps {out} (1:{pointer=ref}*(1))(2:{alias=DISK_PROPS}(struct))
	{
		if o.DiskProperties == nil {
			o.DiskProperties = &csvp.DiskProperties{}
		}
		if err := o.DiskProperties.UnmarshalNDR(ctx, w); err != nil {
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

// GetPropertiesRequest structure represents the CprepDiskGetProps operation request
type GetPropertiesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk for which to get the disk properties.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
}

func (o *GetPropertiesRequest) xxx_ToOp(ctx context.Context) *xxx_GetPropertiesOperation {
	if o == nil {
		return &xxx_GetPropertiesOperation{}
	}
	return &xxx_GetPropertiesOperation{
		This:   o.This,
		DiskID: o.DiskID,
	}
}

func (o *GetPropertiesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
}
func (o *GetPropertiesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetPropertiesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPropertiesResponse structure represents the CprepDiskGetProps operation response
type GetPropertiesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// DiskProps: The properties of the selected ClusPrepDisk.
	DiskProperties *csvp.DiskProperties `idl:"name:DiskProps" json:"disk_properties"`
	// Return: The CprepDiskGetProps return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPropertiesResponse) xxx_ToOp(ctx context.Context) *xxx_GetPropertiesOperation {
	if o == nil {
		return &xxx_GetPropertiesOperation{}
	}
	return &xxx_GetPropertiesOperation{
		That:           o.That,
		DiskProperties: o.DiskProperties,
		Return:         o.Return,
	}
}

func (o *GetPropertiesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DiskProperties = op.DiskProperties
	o.Return = op.Return
}
func (o *GetPropertiesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetPropertiesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_StopDefenseOperation structure represents the CprepDiskStopDefense operation
type xxx_StopDefenseOperation struct {
	This   *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	Return int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_StopDefenseOperation) OpNum() int { return 12 }

func (o *xxx_StopDefenseOperation) OpName() string {
	return "/IClusterStorage2/v0/CprepDiskStopDefense"
}

func (o *xxx_StopDefenseOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopDefenseOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_StopDefenseOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_StopDefenseOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopDefenseOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_StopDefenseOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// StopDefenseRequest structure represents the CprepDiskStopDefense operation request
type StopDefenseRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk representing the disk.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
}

func (o *StopDefenseRequest) xxx_ToOp(ctx context.Context) *xxx_StopDefenseOperation {
	if o == nil {
		return &xxx_StopDefenseOperation{}
	}
	return &xxx_StopDefenseOperation{
		This:   o.This,
		DiskID: o.DiskID,
	}
}

func (o *StopDefenseRequest) xxx_FromOp(ctx context.Context, op *xxx_StopDefenseOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
}
func (o *StopDefenseRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *StopDefenseRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StopDefenseOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// StopDefenseResponse structure represents the CprepDiskStopDefense operation response
type StopDefenseResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CprepDiskStopDefense return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *StopDefenseResponse) xxx_ToOp(ctx context.Context) *xxx_StopDefenseOperation {
	if o == nil {
		return &xxx_StopDefenseOperation{}
	}
	return &xxx_StopDefenseOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *StopDefenseResponse) xxx_FromOp(ctx context.Context, op *xxx_StopDefenseOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *StopDefenseResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *StopDefenseResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StopDefenseOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OnlineOperation structure represents the CprepDiskOnline operation
type xxx_OnlineOperation struct {
	This               *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID             *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	MaxPartitionNumber uint32              `idl:"name:MaxPartitionNumber" json:"max_partition_number"`
	Return             int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_OnlineOperation) OpNum() int { return 13 }

func (o *xxx_OnlineOperation) OpName() string { return "/IClusterStorage2/v0/CprepDiskOnline" }

func (o *xxx_OnlineOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OnlineOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_OnlineOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_OnlineOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OnlineOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// MaxPartitionNumber {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.MaxPartitionNumber); err != nil {
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

func (o *xxx_OnlineOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// MaxPartitionNumber {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.MaxPartitionNumber); err != nil {
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

// OnlineRequest structure represents the CprepDiskOnline operation request
type OnlineRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk representing the disk whose associated
	// volumes will become online.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
}

func (o *OnlineRequest) xxx_ToOp(ctx context.Context) *xxx_OnlineOperation {
	if o == nil {
		return &xxx_OnlineOperation{}
	}
	return &xxx_OnlineOperation{
		This:   o.This,
		DiskID: o.DiskID,
	}
}

func (o *OnlineRequest) xxx_FromOp(ctx context.Context, op *xxx_OnlineOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
}
func (o *OnlineRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *OnlineRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OnlineOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OnlineResponse structure represents the CprepDiskOnline operation response
type OnlineResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// MaxPartitionNumber:  The number of partitions on the disk.
	MaxPartitionNumber uint32 `idl:"name:MaxPartitionNumber" json:"max_partition_number"`
	// Return: The CprepDiskOnline return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OnlineResponse) xxx_ToOp(ctx context.Context) *xxx_OnlineOperation {
	if o == nil {
		return &xxx_OnlineOperation{}
	}
	return &xxx_OnlineOperation{
		That:               o.That,
		MaxPartitionNumber: o.MaxPartitionNumber,
		Return:             o.Return,
	}
}

func (o *OnlineResponse) xxx_FromOp(ctx context.Context, op *xxx_OnlineOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MaxPartitionNumber = op.MaxPartitionNumber
	o.Return = op.Return
}
func (o *OnlineResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *OnlineResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OnlineOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_VerifyUniqueOperation structure represents the CprepDiskVerifyUnique operation
type xxx_VerifyUniqueOperation struct {
	This   *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	Return int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_VerifyUniqueOperation) OpNum() int { return 14 }

func (o *xxx_VerifyUniqueOperation) OpName() string {
	return "/IClusterStorage2/v0/CprepDiskVerifyUnique"
}

func (o *xxx_VerifyUniqueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_VerifyUniqueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_VerifyUniqueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_VerifyUniqueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_VerifyUniqueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_VerifyUniqueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// VerifyUniqueRequest structure represents the CprepDiskVerifyUnique operation request
type VerifyUniqueRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk representing the disk.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
}

func (o *VerifyUniqueRequest) xxx_ToOp(ctx context.Context) *xxx_VerifyUniqueOperation {
	if o == nil {
		return &xxx_VerifyUniqueOperation{}
	}
	return &xxx_VerifyUniqueOperation{
		This:   o.This,
		DiskID: o.DiskID,
	}
}

func (o *VerifyUniqueRequest) xxx_FromOp(ctx context.Context, op *xxx_VerifyUniqueOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
}
func (o *VerifyUniqueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *VerifyUniqueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_VerifyUniqueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// VerifyUniqueResponse structure represents the CprepDiskVerifyUnique operation response
type VerifyUniqueResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CprepDiskVerifyUnique return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *VerifyUniqueResponse) xxx_ToOp(ctx context.Context) *xxx_VerifyUniqueOperation {
	if o == nil {
		return &xxx_VerifyUniqueOperation{}
	}
	return &xxx_VerifyUniqueOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *VerifyUniqueResponse) xxx_FromOp(ctx context.Context, op *xxx_VerifyUniqueOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *VerifyUniqueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *VerifyUniqueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_VerifyUniqueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WriteFileDataOperation structure represents the CprepDiskWriteFileData operation
type xxx_WriteFileDataOperation struct {
	This         *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID       *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	Partition    uint32              `idl:"name:ulPartition" json:"partition"`
	FileName     string              `idl:"name:FileName;string" json:"file_name"`
	DataInLength uint32              `idl:"name:cbDataIn" json:"data_in_length"`
	DataIn       []byte              `idl:"name:DataIn;size_is:(cbDataIn)" json:"data_in"`
	Return       int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_WriteFileDataOperation) OpNum() int { return 17 }

func (o *xxx_WriteFileDataOperation) OpName() string {
	return "/IClusterStorage2/v0/CprepDiskWriteFileData"
}

func (o *xxx_WriteFileDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.DataIn != nil && o.DataInLength == 0 {
		o.DataInLength = uint32(len(o.DataIn))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WriteFileDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ulPartition {in} (1:(uint32))
	{
		if err := w.WriteData(o.Partition); err != nil {
			return err
		}
	}
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.FileName); err != nil {
			return err
		}
	}
	// cbDataIn {in} (1:(uint32))
	{
		if err := w.WriteData(o.DataInLength); err != nil {
			return err
		}
	}
	// DataIn {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbDataIn](uint8))
	{
		dimSize1 := uint64(o.DataInLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.DataIn {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.DataIn[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.DataIn); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_WriteFileDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ulPartition {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Partition); err != nil {
			return err
		}
	}
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.FileName); err != nil {
			return err
		}
	}
	// cbDataIn {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DataInLength); err != nil {
			return err
		}
	}
	// DataIn {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbDataIn](uint8))
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
			return fmt.Errorf("buffer overflow for size %d of array o.DataIn", sizeInfo[0])
		}
		o.DataIn = make([]byte, sizeInfo[0])
		for i1 := range o.DataIn {
			i1 := i1
			if err := w.ReadData(&o.DataIn[i1]); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_WriteFileDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WriteFileDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_WriteFileDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// WriteFileDataRequest structure represents the CprepDiskWriteFileData operation request
type WriteFileDataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk representing the disk that holds the file
	// to write to.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	// ulPartition: The partition number of the partition associated with the volume on
	// the disk that holds the file to write to.
	Partition uint32 `idl:"name:ulPartition" json:"partition"`
	// FileName: The path and name of the file to write to.
	FileName string `idl:"name:FileName;string" json:"file_name"`
	// cbDataIn: The size, in bytes, of the buffer DataIn.
	DataInLength uint32 `idl:"name:cbDataIn" json:"data_in_length"`
	// DataIn: The data to write to the file.
	DataIn []byte `idl:"name:DataIn;size_is:(cbDataIn)" json:"data_in"`
}

func (o *WriteFileDataRequest) xxx_ToOp(ctx context.Context) *xxx_WriteFileDataOperation {
	if o == nil {
		return &xxx_WriteFileDataOperation{}
	}
	return &xxx_WriteFileDataOperation{
		This:         o.This,
		DiskID:       o.DiskID,
		Partition:    o.Partition,
		FileName:     o.FileName,
		DataInLength: o.DataInLength,
		DataIn:       o.DataIn,
	}
}

func (o *WriteFileDataRequest) xxx_FromOp(ctx context.Context, op *xxx_WriteFileDataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
	o.Partition = op.Partition
	o.FileName = op.FileName
	o.DataInLength = op.DataInLength
	o.DataIn = op.DataIn
}
func (o *WriteFileDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *WriteFileDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WriteFileDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WriteFileDataResponse structure represents the CprepDiskWriteFileData operation response
type WriteFileDataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CprepDiskWriteFileData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *WriteFileDataResponse) xxx_ToOp(ctx context.Context) *xxx_WriteFileDataOperation {
	if o == nil {
		return &xxx_WriteFileDataOperation{}
	}
	return &xxx_WriteFileDataOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *WriteFileDataResponse) xxx_FromOp(ctx context.Context, op *xxx_WriteFileDataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *WriteFileDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *WriteFileDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WriteFileDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_VerifyFileDataOperation structure represents the CprepDiskVerifyFileData operation
type xxx_VerifyFileDataOperation struct {
	This         *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID       *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	Partition    uint32              `idl:"name:ulPartition" json:"partition"`
	FileName     string              `idl:"name:FileName;string" json:"file_name"`
	DataInLength uint32              `idl:"name:cbDataIn" json:"data_in_length"`
	DataIn       []byte              `idl:"name:DataIn;size_is:(cbDataIn)" json:"data_in"`
	Return       int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_VerifyFileDataOperation) OpNum() int { return 18 }

func (o *xxx_VerifyFileDataOperation) OpName() string {
	return "/IClusterStorage2/v0/CprepDiskVerifyFileData"
}

func (o *xxx_VerifyFileDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.DataIn != nil && o.DataInLength == 0 {
		o.DataInLength = uint32(len(o.DataIn))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_VerifyFileDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ulPartition {in} (1:(uint32))
	{
		if err := w.WriteData(o.Partition); err != nil {
			return err
		}
	}
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.FileName); err != nil {
			return err
		}
	}
	// cbDataIn {in} (1:(uint32))
	{
		if err := w.WriteData(o.DataInLength); err != nil {
			return err
		}
	}
	// DataIn {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbDataIn](uint8))
	{
		dimSize1 := uint64(o.DataInLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.DataIn {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.DataIn[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.DataIn); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_VerifyFileDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ulPartition {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Partition); err != nil {
			return err
		}
	}
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.FileName); err != nil {
			return err
		}
	}
	// cbDataIn {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DataInLength); err != nil {
			return err
		}
	}
	// DataIn {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbDataIn](uint8))
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
			return fmt.Errorf("buffer overflow for size %d of array o.DataIn", sizeInfo[0])
		}
		o.DataIn = make([]byte, sizeInfo[0])
		for i1 := range o.DataIn {
			i1 := i1
			if err := w.ReadData(&o.DataIn[i1]); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_VerifyFileDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_VerifyFileDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_VerifyFileDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// VerifyFileDataRequest structure represents the CprepDiskVerifyFileData operation request
type VerifyFileDataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId:  The identifier of the ClusPrepDisk representing the disk that holds the
	// file to verify.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	// ulPartition: The partition number of the partition associated with the volume on
	// the disk that holds the file to verify from.
	Partition uint32 `idl:"name:ulPartition" json:"partition"`
	// FileName:  The path and name of the file to verify from.
	FileName string `idl:"name:FileName;string" json:"file_name"`
	// cbDataIn:  The size, in bytes, of the buffer DataIn.
	DataInLength uint32 `idl:"name:cbDataIn" json:"data_in_length"`
	// DataIn:  The data to verify against the file.
	DataIn []byte `idl:"name:DataIn;size_is:(cbDataIn)" json:"data_in"`
}

func (o *VerifyFileDataRequest) xxx_ToOp(ctx context.Context) *xxx_VerifyFileDataOperation {
	if o == nil {
		return &xxx_VerifyFileDataOperation{}
	}
	return &xxx_VerifyFileDataOperation{
		This:         o.This,
		DiskID:       o.DiskID,
		Partition:    o.Partition,
		FileName:     o.FileName,
		DataInLength: o.DataInLength,
		DataIn:       o.DataIn,
	}
}

func (o *VerifyFileDataRequest) xxx_FromOp(ctx context.Context, op *xxx_VerifyFileDataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
	o.Partition = op.Partition
	o.FileName = op.FileName
	o.DataInLength = op.DataInLength
	o.DataIn = op.DataIn
}
func (o *VerifyFileDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *VerifyFileDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_VerifyFileDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// VerifyFileDataResponse structure represents the CprepDiskVerifyFileData operation response
type VerifyFileDataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CprepDiskVerifyFileData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *VerifyFileDataResponse) xxx_ToOp(ctx context.Context) *xxx_VerifyFileDataOperation {
	if o == nil {
		return &xxx_VerifyFileDataOperation{}
	}
	return &xxx_VerifyFileDataOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *VerifyFileDataResponse) xxx_FromOp(ctx context.Context, op *xxx_VerifyFileDataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *VerifyFileDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *VerifyFileDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_VerifyFileDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteFileOperation structure represents the CprepDiskDeleteFile operation
type xxx_DeleteFileOperation struct {
	This      *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID    *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	Partition uint32              `idl:"name:ulPartition" json:"partition"`
	FileName  string              `idl:"name:FileName;string" json:"file_name"`
	Return    int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteFileOperation) OpNum() int { return 19 }

func (o *xxx_DeleteFileOperation) OpName() string { return "/IClusterStorage2/v0/CprepDiskDeleteFile" }

func (o *xxx_DeleteFileOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteFileOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ulPartition {in} (1:(uint32))
	{
		if err := w.WriteData(o.Partition); err != nil {
			return err
		}
	}
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.FileName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteFileOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ulPartition {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Partition); err != nil {
			return err
		}
	}
	// FileName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.FileName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteFileOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteFileOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteFileOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteFileRequest structure represents the CprepDiskDeleteFile operation request
type DeleteFileRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId:  The identifier of the ClusPrepDisk representing the disk that holds the
	// file to be deleted.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	// ulPartition: The partition number of the partition associated with the volume on
	// the disk that holds the file to be deleted.
	Partition uint32 `idl:"name:ulPartition" json:"partition"`
	// FileName:  The path and name of the file to delete.
	FileName string `idl:"name:FileName;string" json:"file_name"`
}

func (o *DeleteFileRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteFileOperation {
	if o == nil {
		return &xxx_DeleteFileOperation{}
	}
	return &xxx_DeleteFileOperation{
		This:      o.This,
		DiskID:    o.DiskID,
		Partition: o.Partition,
		FileName:  o.FileName,
	}
}

func (o *DeleteFileRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteFileOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
	o.Partition = op.Partition
	o.FileName = op.FileName
}
func (o *DeleteFileRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteFileRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteFileOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteFileResponse structure represents the CprepDiskDeleteFile operation response
type DeleteFileResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CprepDiskDeleteFile return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteFileResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteFileOperation {
	if o == nil {
		return &xxx_DeleteFileOperation{}
	}
	return &xxx_DeleteFileOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *DeleteFileResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteFileOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteFileResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteFileResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteFileOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OfflineOperation structure represents the CprepDiskOffline operation
type xxx_OfflineOperation struct {
	This   *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	Return int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_OfflineOperation) OpNum() int { return 20 }

func (o *xxx_OfflineOperation) OpName() string { return "/IClusterStorage2/v0/CprepDiskOffline" }

func (o *xxx_OfflineOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OfflineOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_OfflineOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_OfflineOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OfflineOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_OfflineOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// OfflineRequest structure represents the CprepDiskOffline operation request
type OfflineRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk representing the disk whose associated
	// volumes will become Offline.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
}

func (o *OfflineRequest) xxx_ToOp(ctx context.Context) *xxx_OfflineOperation {
	if o == nil {
		return &xxx_OfflineOperation{}
	}
	return &xxx_OfflineOperation{
		This:   o.This,
		DiskID: o.DiskID,
	}
}

func (o *OfflineRequest) xxx_FromOp(ctx context.Context, op *xxx_OfflineOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
}
func (o *OfflineRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *OfflineRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OfflineOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OfflineResponse structure represents the CprepDiskOffline operation response
type OfflineResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CprepDiskOffline return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OfflineResponse) xxx_ToOp(ctx context.Context) *xxx_OfflineOperation {
	if o == nil {
		return &xxx_OfflineOperation{}
	}
	return &xxx_OfflineOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *OfflineResponse) xxx_FromOp(ctx context.Context, op *xxx_OfflineOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *OfflineResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *OfflineResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OfflineOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetUniqueIDsOperation structure represents the CprepDiskGetUniqueIds operation
type xxx_GetUniqueIDsOperation struct {
	This          *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID        *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	DataLength    uint32              `idl:"name:cbData" json:"data_length"`
	Data          []byte              `idl:"name:pbData;size_is:(cbData);length_is:(pcbDataOut)" json:"data"`
	DataOutLength uint32              `idl:"name:pcbDataOut" json:"data_out_length"`
	NeededLength  uint32              `idl:"name:pcbNeeded" json:"needed_length"`
	Return        int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_GetUniqueIDsOperation) OpNum() int { return 22 }

func (o *xxx_GetUniqueIDsOperation) OpName() string {
	return "/IClusterStorage2/v0/CprepDiskGetUniqueIds"
}

func (o *xxx_GetUniqueIDsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUniqueIDsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// cbData {in} (1:(uint32))
	{
		if err := w.WriteData(o.DataLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUniqueIDsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// cbData {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DataLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUniqueIDsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Data != nil && o.DataOutLength == 0 {
		o.DataOutLength = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUniqueIDsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbData {out} (1:{pointer=ref}*(1)[dim:0,size_is=cbData,length_is=pcbDataOut](uint8))
	{
		dimSize1 := uint64(o.DataLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.DataOutLength)
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
	}
	// pcbDataOut {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.DataOutLength); err != nil {
			return err
		}
	}
	// pcbNeeded {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.NeededLength); err != nil {
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

func (o *xxx_GetUniqueIDsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbData {out} (1:{pointer=ref}*(1)[dim:0,size_is=cbData,length_is=pcbDataOut](uint8))
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
	// pcbDataOut {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.DataOutLength); err != nil {
			return err
		}
	}
	// pcbNeeded {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.NeededLength); err != nil {
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

// GetUniqueIDsRequest structure represents the CprepDiskGetUniqueIds operation request
type GetUniqueIDsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier representing the ClusPrepDisk for which to retrieve the device
	// ID data.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	// cbData: The size, in bytes, of the pbData buffer passed to the server.
	DataLength uint32 `idl:"name:cbData" json:"data_length"`
}

func (o *GetUniqueIDsRequest) xxx_ToOp(ctx context.Context) *xxx_GetUniqueIDsOperation {
	if o == nil {
		return &xxx_GetUniqueIDsOperation{}
	}
	return &xxx_GetUniqueIDsOperation{
		This:       o.This,
		DiskID:     o.DiskID,
		DataLength: o.DataLength,
	}
}

func (o *GetUniqueIDsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetUniqueIDsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
	o.DataLength = op.DataLength
}
func (o *GetUniqueIDsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetUniqueIDsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetUniqueIDsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetUniqueIDsResponse structure represents the CprepDiskGetUniqueIds operation response
type GetUniqueIDsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbData: The output buffer for the device ID data.
	Data []byte `idl:"name:pbData;size_is:(cbData);length_is:(pcbDataOut)" json:"data"`
	// pcbDataOut: The size, in bytes, of the amount of data written to pbData on a successful
	// return.
	DataOutLength uint32 `idl:"name:pcbDataOut" json:"data_out_length"`
	// pcbNeeded: If ERROR_INSUFFICIENT_BUFFER is returned, then this parameter contains
	// the size, in bytes, of the buffer required for a successful call.
	NeededLength uint32 `idl:"name:pcbNeeded" json:"needed_length"`
	// Return: The CprepDiskGetUniqueIds return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetUniqueIDsResponse) xxx_ToOp(ctx context.Context) *xxx_GetUniqueIDsOperation {
	if o == nil {
		return &xxx_GetUniqueIDsOperation{}
	}
	return &xxx_GetUniqueIDsOperation{
		That:          o.That,
		Data:          o.Data,
		DataOutLength: o.DataOutLength,
		NeededLength:  o.NeededLength,
		Return:        o.Return,
	}
}

func (o *GetUniqueIDsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetUniqueIDsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Data = op.Data
	o.DataOutLength = op.DataOutLength
	o.NeededLength = op.NeededLength
	o.Return = op.Return
}
func (o *GetUniqueIDsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetUniqueIDsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetUniqueIDsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AttachOperation structure represents the CprepDiskAttach operation
type xxx_AttachOperation struct {
	This   *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	Return int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_AttachOperation) OpNum() int { return 23 }

func (o *xxx_AttachOperation) OpName() string { return "/IClusterStorage2/v0/CprepDiskAttach" }

func (o *xxx_AttachOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AttachOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AttachOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_AttachOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AttachOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AttachOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// AttachRequest structure represents the CprepDiskAttach operation request
type AttachRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk representing the disk.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
}

func (o *AttachRequest) xxx_ToOp(ctx context.Context) *xxx_AttachOperation {
	if o == nil {
		return &xxx_AttachOperation{}
	}
	return &xxx_AttachOperation{
		This:   o.This,
		DiskID: o.DiskID,
	}
}

func (o *AttachRequest) xxx_FromOp(ctx context.Context, op *xxx_AttachOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
}
func (o *AttachRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AttachRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AttachOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AttachResponse structure represents the CprepDiskAttach operation response
type AttachResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CprepDiskAttach return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AttachResponse) xxx_ToOp(ctx context.Context) *xxx_AttachOperation {
	if o == nil {
		return &xxx_AttachOperation{}
	}
	return &xxx_AttachOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *AttachResponse) xxx_FromOp(ctx context.Context, op *xxx_AttachOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *AttachResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AttachResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AttachOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PRArbitrateOperation structure represents the CprepDiskPRArbitrate operation
type xxx_PRArbitrateOperation struct {
	This   *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	Return int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_PRArbitrateOperation) OpNum() int { return 24 }

func (o *xxx_PRArbitrateOperation) OpName() string {
	return "/IClusterStorage2/v0/CprepDiskPRArbitrate"
}

func (o *xxx_PRArbitrateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRArbitrateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PRArbitrateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_PRArbitrateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRArbitrateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PRArbitrateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// PRArbitrateRequest structure represents the CprepDiskPRArbitrate operation request
type PRArbitrateRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk representing the disk.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
}

func (o *PRArbitrateRequest) xxx_ToOp(ctx context.Context) *xxx_PRArbitrateOperation {
	if o == nil {
		return &xxx_PRArbitrateOperation{}
	}
	return &xxx_PRArbitrateOperation{
		This:   o.This,
		DiskID: o.DiskID,
	}
}

func (o *PRArbitrateRequest) xxx_FromOp(ctx context.Context, op *xxx_PRArbitrateOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
}
func (o *PRArbitrateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PRArbitrateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PRArbitrateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PRArbitrateResponse structure represents the CprepDiskPRArbitrate operation response
type PRArbitrateResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CprepDiskPRArbitrate return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PRArbitrateResponse) xxx_ToOp(ctx context.Context) *xxx_PRArbitrateOperation {
	if o == nil {
		return &xxx_PRArbitrateOperation{}
	}
	return &xxx_PRArbitrateOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *PRArbitrateResponse) xxx_FromOp(ctx context.Context, op *xxx_PRArbitrateOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *PRArbitrateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PRArbitrateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PRArbitrateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PRRegisterOperation structure represents the CprepDiskPRRegister operation
type xxx_PRRegisterOperation struct {
	This   *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	Return int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_PRRegisterOperation) OpNum() int { return 25 }

func (o *xxx_PRRegisterOperation) OpName() string { return "/IClusterStorage2/v0/CprepDiskPRRegister" }

func (o *xxx_PRRegisterOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRRegisterOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PRRegisterOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_PRRegisterOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRRegisterOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PRRegisterOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// PRRegisterRequest structure represents the CprepDiskPRRegister operation request
type PRRegisterRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk representing the disk.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
}

func (o *PRRegisterRequest) xxx_ToOp(ctx context.Context) *xxx_PRRegisterOperation {
	if o == nil {
		return &xxx_PRRegisterOperation{}
	}
	return &xxx_PRRegisterOperation{
		This:   o.This,
		DiskID: o.DiskID,
	}
}

func (o *PRRegisterRequest) xxx_FromOp(ctx context.Context, op *xxx_PRRegisterOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
}
func (o *PRRegisterRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PRRegisterRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PRRegisterOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PRRegisterResponse structure represents the CprepDiskPRRegister operation response
type PRRegisterResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CprepDiskPRRegister return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PRRegisterResponse) xxx_ToOp(ctx context.Context) *xxx_PRRegisterOperation {
	if o == nil {
		return &xxx_PRRegisterOperation{}
	}
	return &xxx_PRRegisterOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *PRRegisterResponse) xxx_FromOp(ctx context.Context, op *xxx_PRRegisterOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *PRRegisterResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PRRegisterResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PRRegisterOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PRUnregisterOperation structure represents the CprepDiskPRUnRegister operation
type xxx_PRUnregisterOperation struct {
	This   *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	Return int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_PRUnregisterOperation) OpNum() int { return 26 }

func (o *xxx_PRUnregisterOperation) OpName() string {
	return "/IClusterStorage2/v0/CprepDiskPRUnRegister"
}

func (o *xxx_PRUnregisterOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRUnregisterOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PRUnregisterOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_PRUnregisterOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRUnregisterOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PRUnregisterOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// PRUnregisterRequest structure represents the CprepDiskPRUnRegister operation request
type PRUnregisterRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk representing the disk.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
}

func (o *PRUnregisterRequest) xxx_ToOp(ctx context.Context) *xxx_PRUnregisterOperation {
	if o == nil {
		return &xxx_PRUnregisterOperation{}
	}
	return &xxx_PRUnregisterOperation{
		This:   o.This,
		DiskID: o.DiskID,
	}
}

func (o *PRUnregisterRequest) xxx_FromOp(ctx context.Context, op *xxx_PRUnregisterOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
}
func (o *PRUnregisterRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PRUnregisterRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PRUnregisterOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PRUnregisterResponse structure represents the CprepDiskPRUnRegister operation response
type PRUnregisterResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CprepDiskPRUnRegister return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PRUnregisterResponse) xxx_ToOp(ctx context.Context) *xxx_PRUnregisterOperation {
	if o == nil {
		return &xxx_PRUnregisterOperation{}
	}
	return &xxx_PRUnregisterOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *PRUnregisterResponse) xxx_FromOp(ctx context.Context, op *xxx_PRUnregisterOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *PRUnregisterResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PRUnregisterResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PRUnregisterOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PRReserveOperation structure represents the CprepDiskPRReserve operation
type xxx_PRReserveOperation struct {
	This   *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	Return int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_PRReserveOperation) OpNum() int { return 27 }

func (o *xxx_PRReserveOperation) OpName() string { return "/IClusterStorage2/v0/CprepDiskPRReserve" }

func (o *xxx_PRReserveOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRReserveOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PRReserveOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_PRReserveOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRReserveOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PRReserveOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// PRReserveRequest structure represents the CprepDiskPRReserve operation request
type PRReserveRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk representing the disk.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
}

func (o *PRReserveRequest) xxx_ToOp(ctx context.Context) *xxx_PRReserveOperation {
	if o == nil {
		return &xxx_PRReserveOperation{}
	}
	return &xxx_PRReserveOperation{
		This:   o.This,
		DiskID: o.DiskID,
	}
}

func (o *PRReserveRequest) xxx_FromOp(ctx context.Context, op *xxx_PRReserveOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
}
func (o *PRReserveRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PRReserveRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PRReserveOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PRReserveResponse structure represents the CprepDiskPRReserve operation response
type PRReserveResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CprepDiskPRReserve return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PRReserveResponse) xxx_ToOp(ctx context.Context) *xxx_PRReserveOperation {
	if o == nil {
		return &xxx_PRReserveOperation{}
	}
	return &xxx_PRReserveOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *PRReserveResponse) xxx_FromOp(ctx context.Context, op *xxx_PRReserveOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *PRReserveResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PRReserveResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PRReserveOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PRReleaseOperation structure represents the CprepDiskPRRelease operation
type xxx_PRReleaseOperation struct {
	This   *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	Return int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_PRReleaseOperation) OpNum() int { return 28 }

func (o *xxx_PRReleaseOperation) OpName() string { return "/IClusterStorage2/v0/CprepDiskPRRelease" }

func (o *xxx_PRReleaseOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRReleaseOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PRReleaseOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_PRReleaseOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRReleaseOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PRReleaseOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// PRReleaseRequest structure represents the CprepDiskPRRelease operation request
type PRReleaseRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk representing the disk.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
}

func (o *PRReleaseRequest) xxx_ToOp(ctx context.Context) *xxx_PRReleaseOperation {
	if o == nil {
		return &xxx_PRReleaseOperation{}
	}
	return &xxx_PRReleaseOperation{
		This:   o.This,
		DiskID: o.DiskID,
	}
}

func (o *PRReleaseRequest) xxx_FromOp(ctx context.Context, op *xxx_PRReleaseOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
}
func (o *PRReleaseRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PRReleaseRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PRReleaseOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PRReleaseResponse structure represents the CprepDiskPRRelease operation response
type PRReleaseResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CprepDiskPRRelease return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PRReleaseResponse) xxx_ToOp(ctx context.Context) *xxx_PRReleaseOperation {
	if o == nil {
		return &xxx_PRReleaseOperation{}
	}
	return &xxx_PRReleaseOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *PRReleaseResponse) xxx_FromOp(ctx context.Context, op *xxx_PRReleaseOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *PRReleaseResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PRReleaseResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PRReleaseOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DiskPartitionIsNTFSOperation structure represents the CprepDiskDiskPartitionIsNtfs operation
type xxx_DiskPartitionIsNTFSOperation struct {
	This      *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID    *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	Partition uint32              `idl:"name:ulPartition" json:"partition"`
	Return    int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_DiskPartitionIsNTFSOperation) OpNum() int { return 29 }

func (o *xxx_DiskPartitionIsNTFSOperation) OpName() string {
	return "/IClusterStorage2/v0/CprepDiskDiskPartitionIsNtfs"
}

func (o *xxx_DiskPartitionIsNTFSOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DiskPartitionIsNTFSOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ulPartition {in} (1:(uint32))
	{
		if err := w.WriteData(o.Partition); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DiskPartitionIsNTFSOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ulPartition {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Partition); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DiskPartitionIsNTFSOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DiskPartitionIsNTFSOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DiskPartitionIsNTFSOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DiskPartitionIsNTFSRequest structure represents the CprepDiskDiskPartitionIsNtfs operation request
type DiskPartitionIsNTFSRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk representing the disk.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	// ulPartition: The partition number of the partition associated with the volume to
	// query for file system information.
	Partition uint32 `idl:"name:ulPartition" json:"partition"`
}

func (o *DiskPartitionIsNTFSRequest) xxx_ToOp(ctx context.Context) *xxx_DiskPartitionIsNTFSOperation {
	if o == nil {
		return &xxx_DiskPartitionIsNTFSOperation{}
	}
	return &xxx_DiskPartitionIsNTFSOperation{
		This:      o.This,
		DiskID:    o.DiskID,
		Partition: o.Partition,
	}
}

func (o *DiskPartitionIsNTFSRequest) xxx_FromOp(ctx context.Context, op *xxx_DiskPartitionIsNTFSOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
	o.Partition = op.Partition
}
func (o *DiskPartitionIsNTFSRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DiskPartitionIsNTFSRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DiskPartitionIsNTFSOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DiskPartitionIsNTFSResponse structure represents the CprepDiskDiskPartitionIsNtfs operation response
type DiskPartitionIsNTFSResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CprepDiskDiskPartitionIsNtfs return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DiskPartitionIsNTFSResponse) xxx_ToOp(ctx context.Context) *xxx_DiskPartitionIsNTFSOperation {
	if o == nil {
		return &xxx_DiskPartitionIsNTFSOperation{}
	}
	return &xxx_DiskPartitionIsNTFSOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *DiskPartitionIsNTFSResponse) xxx_FromOp(ctx context.Context, op *xxx_DiskPartitionIsNTFSOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DiskPartitionIsNTFSResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DiskPartitionIsNTFSResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DiskPartitionIsNTFSOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetArbSectorsOperation structure represents the CprepDiskGetArbSectors operation
type xxx_GetArbSectorsOperation struct {
	This    *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID  *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	SectorX uint32              `idl:"name:SectorX" json:"sector_x"`
	SectorY uint32              `idl:"name:SectorY" json:"sector_y"`
	Return  int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_GetArbSectorsOperation) OpNum() int { return 30 }

func (o *xxx_GetArbSectorsOperation) OpName() string {
	return "/IClusterStorage2/v0/CprepDiskGetArbSectors"
}

func (o *xxx_GetArbSectorsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetArbSectorsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetArbSectorsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetArbSectorsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetArbSectorsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// SectorX {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.SectorX); err != nil {
			return err
		}
	}
	// SectorY {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.SectorY); err != nil {
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

func (o *xxx_GetArbSectorsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// SectorX {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.SectorX); err != nil {
			return err
		}
	}
	// SectorY {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.SectorY); err != nil {
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

// GetArbSectorsRequest structure represents the CprepDiskGetArbSectors operation request
type GetArbSectorsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk representing the disk.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
}

func (o *GetArbSectorsRequest) xxx_ToOp(ctx context.Context) *xxx_GetArbSectorsOperation {
	if o == nil {
		return &xxx_GetArbSectorsOperation{}
	}
	return &xxx_GetArbSectorsOperation{
		This:   o.This,
		DiskID: o.DiskID,
	}
}

func (o *GetArbSectorsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetArbSectorsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
}
func (o *GetArbSectorsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetArbSectorsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetArbSectorsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetArbSectorsResponse structure represents the CprepDiskGetArbSectors operation response
type GetArbSectorsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// SectorX: The first sector number that is available.
	SectorX uint32 `idl:"name:SectorX" json:"sector_x"`
	// SectorY: The second sector number that is available.
	SectorY uint32 `idl:"name:SectorY" json:"sector_y"`
	// Return: The CprepDiskGetArbSectors return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetArbSectorsResponse) xxx_ToOp(ctx context.Context) *xxx_GetArbSectorsOperation {
	if o == nil {
		return &xxx_GetArbSectorsOperation{}
	}
	return &xxx_GetArbSectorsOperation{
		That:    o.That,
		SectorX: o.SectorX,
		SectorY: o.SectorY,
		Return:  o.Return,
	}
}

func (o *GetArbSectorsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetArbSectorsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SectorX = op.SectorX
	o.SectorY = op.SectorY
	o.Return = op.Return
}
func (o *GetArbSectorsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetArbSectorsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetArbSectorsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IsPRPresentOperation structure represents the CprepDiskIsPRPresent operation
type xxx_IsPRPresentOperation struct {
	This    *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID  *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	Present uint32              `idl:"name:Present" json:"present"`
	Return  int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_IsPRPresentOperation) OpNum() int { return 31 }

func (o *xxx_IsPRPresentOperation) OpName() string {
	return "/IClusterStorage2/v0/CprepDiskIsPRPresent"
}

func (o *xxx_IsPRPresentOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsPRPresentOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_IsPRPresentOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_IsPRPresentOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsPRPresentOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Present {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Present); err != nil {
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

func (o *xxx_IsPRPresentOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Present {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Present); err != nil {
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

// IsPRPresentRequest structure represents the CprepDiskIsPRPresent operation request
type IsPRPresentRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk representing the disk.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
}

func (o *IsPRPresentRequest) xxx_ToOp(ctx context.Context) *xxx_IsPRPresentOperation {
	if o == nil {
		return &xxx_IsPRPresentOperation{}
	}
	return &xxx_IsPRPresentOperation{
		This:   o.This,
		DiskID: o.DiskID,
	}
}

func (o *IsPRPresentRequest) xxx_FromOp(ctx context.Context, op *xxx_IsPRPresentOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
}
func (o *IsPRPresentRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *IsPRPresentRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsPRPresentOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsPRPresentResponse structure represents the CprepDiskIsPRPresent operation response
type IsPRPresentResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Present: MUST be 0x00000000 if no reserves are present. MUST be 0x00000001 or 0x00000002
	// if reserves are present.
	//
	//	+------------+--------------------------------------------------------------------------+
	//	|            |                                                                          |
	//	|   VALUE    |                                 MEANING                                  |
	//	|            |                                                                          |
	//	+------------+--------------------------------------------------------------------------+
	//	+------------+--------------------------------------------------------------------------+
	//	| 0x00000000 | No reserves are present.                                                 |
	//	+------------+--------------------------------------------------------------------------+
	//	| 0x00000001 | A persistent reservation is present and is not held by the local server. |
	//	+------------+--------------------------------------------------------------------------+
	//	| 0x00000002 | A persistent reservation is present and is held by the local server.<15> |
	//	+------------+--------------------------------------------------------------------------+
	Present uint32 `idl:"name:Present" json:"present"`
	// Return: The CprepDiskIsPRPresent return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *IsPRPresentResponse) xxx_ToOp(ctx context.Context) *xxx_IsPRPresentOperation {
	if o == nil {
		return &xxx_IsPRPresentOperation{}
	}
	return &xxx_IsPRPresentOperation{
		That:    o.That,
		Present: o.Present,
		Return:  o.Return,
	}
}

func (o *IsPRPresentResponse) xxx_FromOp(ctx context.Context, op *xxx_IsPRPresentOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Present = op.Present
	o.Return = op.Return
}
func (o *IsPRPresentResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *IsPRPresentResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsPRPresentOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PRPreemptOperation structure represents the CprepDiskPRPreempt operation
type xxx_PRPreemptOperation struct {
	This   *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	Return int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_PRPreemptOperation) OpNum() int { return 32 }

func (o *xxx_PRPreemptOperation) OpName() string { return "/IClusterStorage2/v0/CprepDiskPRPreempt" }

func (o *xxx_PRPreemptOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRPreemptOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PRPreemptOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_PRPreemptOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRPreemptOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PRPreemptOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// PRPreemptRequest structure represents the CprepDiskPRPreempt operation request
type PRPreemptRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk representing the disk.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
}

func (o *PRPreemptRequest) xxx_ToOp(ctx context.Context) *xxx_PRPreemptOperation {
	if o == nil {
		return &xxx_PRPreemptOperation{}
	}
	return &xxx_PRPreemptOperation{
		This:   o.This,
		DiskID: o.DiskID,
	}
}

func (o *PRPreemptRequest) xxx_FromOp(ctx context.Context, op *xxx_PRPreemptOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
}
func (o *PRPreemptRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PRPreemptRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PRPreemptOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PRPreemptResponse structure represents the CprepDiskPRPreempt operation response
type PRPreemptResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CprepDiskPRPreempt return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PRPreemptResponse) xxx_ToOp(ctx context.Context) *xxx_PRPreemptOperation {
	if o == nil {
		return &xxx_PRPreemptOperation{}
	}
	return &xxx_PRPreemptOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *PRPreemptResponse) xxx_FromOp(ctx context.Context, op *xxx_PRPreemptOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *PRPreemptResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PRPreemptResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PRPreemptOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PRClearOperation structure represents the CprepDiskPRClear operation
type xxx_PRClearOperation struct {
	This   *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	Return int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_PRClearOperation) OpNum() int { return 33 }

func (o *xxx_PRClearOperation) OpName() string { return "/IClusterStorage2/v0/CprepDiskPRClear" }

func (o *xxx_PRClearOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRClearOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PRClearOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_PRClearOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRClearOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PRClearOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// PRClearRequest structure represents the CprepDiskPRClear operation request
type PRClearRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk representing the disk.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
}

func (o *PRClearRequest) xxx_ToOp(ctx context.Context) *xxx_PRClearOperation {
	if o == nil {
		return &xxx_PRClearOperation{}
	}
	return &xxx_PRClearOperation{
		This:   o.This,
		DiskID: o.DiskID,
	}
}

func (o *PRClearRequest) xxx_FromOp(ctx context.Context, op *xxx_PRClearOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
}
func (o *PRClearRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PRClearRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PRClearOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PRClearResponse structure represents the CprepDiskPRClear operation response
type PRClearResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CprepDiskPRClear return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PRClearResponse) xxx_ToOp(ctx context.Context) *xxx_PRClearOperation {
	if o == nil {
		return &xxx_PRClearOperation{}
	}
	return &xxx_PRClearOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *PRClearResponse) xxx_FromOp(ctx context.Context, op *xxx_PRClearOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *PRClearResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PRClearResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PRClearOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IsOnlineOperation structure represents the CprepDiskIsOnline operation
type xxx_IsOnlineOperation struct {
	This   *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	Return int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_IsOnlineOperation) OpNum() int { return 34 }

func (o *xxx_IsOnlineOperation) OpName() string { return "/IClusterStorage2/v0/CprepDiskIsOnline" }

func (o *xxx_IsOnlineOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsOnlineOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_IsOnlineOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_IsOnlineOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsOnlineOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_IsOnlineOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// IsOnlineRequest structure represents the CprepDiskIsOnline operation request
type IsOnlineRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier representing the ClusPrepDisk.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
}

func (o *IsOnlineRequest) xxx_ToOp(ctx context.Context) *xxx_IsOnlineOperation {
	if o == nil {
		return &xxx_IsOnlineOperation{}
	}
	return &xxx_IsOnlineOperation{
		This:   o.This,
		DiskID: o.DiskID,
	}
}

func (o *IsOnlineRequest) xxx_FromOp(ctx context.Context, op *xxx_IsOnlineOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
}
func (o *IsOnlineRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *IsOnlineRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsOnlineOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsOnlineResponse structure represents the CprepDiskIsOnline operation response
type IsOnlineResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CprepDiskIsOnline return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *IsOnlineResponse) xxx_ToOp(ctx context.Context) *xxx_IsOnlineOperation {
	if o == nil {
		return &xxx_IsOnlineOperation{}
	}
	return &xxx_IsOnlineOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *IsOnlineResponse) xxx_FromOp(ctx context.Context, op *xxx_IsOnlineOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *IsOnlineResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *IsOnlineResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsOnlineOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetOnlineOperation structure represents the CprepDiskSetOnline operation
type xxx_SetOnlineOperation struct {
	This   *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	Return int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_SetOnlineOperation) OpNum() int { return 35 }

func (o *xxx_SetOnlineOperation) OpName() string { return "/IClusterStorage2/v0/CprepDiskSetOnline" }

func (o *xxx_SetOnlineOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOnlineOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetOnlineOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_SetOnlineOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOnlineOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetOnlineOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetOnlineRequest structure represents the CprepDiskSetOnline operation request
type SetOnlineRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier representing the ClusPrepDisk.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
}

func (o *SetOnlineRequest) xxx_ToOp(ctx context.Context) *xxx_SetOnlineOperation {
	if o == nil {
		return &xxx_SetOnlineOperation{}
	}
	return &xxx_SetOnlineOperation{
		This:   o.This,
		DiskID: o.DiskID,
	}
}

func (o *SetOnlineRequest) xxx_FromOp(ctx context.Context, op *xxx_SetOnlineOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
}
func (o *SetOnlineRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetOnlineRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetOnlineOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetOnlineResponse structure represents the CprepDiskSetOnline operation response
type SetOnlineResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CprepDiskSetOnline return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetOnlineResponse) xxx_ToOp(ctx context.Context) *xxx_SetOnlineOperation {
	if o == nil {
		return &xxx_SetOnlineOperation{}
	}
	return &xxx_SetOnlineOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetOnlineResponse) xxx_FromOp(ctx context.Context, op *xxx_SetOnlineOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetOnlineResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetOnlineResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetOnlineOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetFSNameOperation structure represents the CprepDiskGetFSName operation
type xxx_GetFSNameOperation struct {
	This      *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID    *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	Partition uint32              `idl:"name:Partition" json:"partition"`
	FSName    []uint16            `idl:"name:FsName" json:"fs_name"`
	Return    int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFSNameOperation) OpNum() int { return 36 }

func (o *xxx_GetFSNameOperation) OpName() string { return "/IClusterStorage2/v0/CprepDiskGetFSName" }

func (o *xxx_GetFSNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFSNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Partition {in} (1:(uint32))
	{
		if err := w.WriteData(o.Partition); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFSNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Partition {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Partition); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFSNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFSNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// FsName {out} (1:[100](wchar))
	{
		for i1 := range o.FSName {
			i1 := i1
			if uint64(i1) >= 100 {
				break
			}
			if err := w.WriteData(o.FSName[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.FSName); uint64(i1) < 100; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
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

func (o *xxx_GetFSNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// FsName {out} (1:[100](wchar))
	{
		o.FSName = make([]uint16, 100)
		for i1 := range o.FSName {
			i1 := i1
			if err := w.ReadData(&o.FSName[i1]); err != nil {
				return err
			}
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

// GetFSNameRequest structure represents the CprepDiskGetFSName operation request
type GetFSNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk representing the disk.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	// Partition: The partition number of the partition associated with the volume to query
	// for file system information.
	Partition uint32 `idl:"name:Partition" json:"partition"`
}

func (o *GetFSNameRequest) xxx_ToOp(ctx context.Context) *xxx_GetFSNameOperation {
	if o == nil {
		return &xxx_GetFSNameOperation{}
	}
	return &xxx_GetFSNameOperation{
		This:      o.This,
		DiskID:    o.DiskID,
		Partition: o.Partition,
	}
}

func (o *GetFSNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFSNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
	o.Partition = op.Partition
}
func (o *GetFSNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetFSNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFSNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFSNameResponse structure represents the CprepDiskGetFSName operation response
type GetFSNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// FsName: A null-terminated output string that contains the name of the file system.
	// The value MUST be "NTFS" if the partition has the NTFS file system. The value MUST
	// be "FAT" for the file allocation table (FAT) file system. No file system and unrecognized
	// file systems MUST be "RAW". Other values can be used for file systems not specified
	// here.
	//
	//	+--------+-----------------------------------------------------------+
	//	|        |                                                           |
	//	| VALUE  |                          MEANING                          |
	//	|        |                                                           |
	//	+--------+-----------------------------------------------------------+
	//	+--------+-----------------------------------------------------------+
	//	| "NTFS" | The partition file system is NTFS.                        |
	//	+--------+-----------------------------------------------------------+
	//	| "FAT"  | The partition file system is FAT.                         |
	//	+--------+-----------------------------------------------------------+
	//	| "RAW"  | There is no partition file system, or it is unrecognized. |
	//	+--------+-----------------------------------------------------------+
	FSName []uint16 `idl:"name:FsName" json:"fs_name"`
	// Return: The CprepDiskGetFSName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetFSNameResponse) xxx_ToOp(ctx context.Context) *xxx_GetFSNameOperation {
	if o == nil {
		return &xxx_GetFSNameOperation{}
	}
	return &xxx_GetFSNameOperation{
		That:   o.That,
		FSName: o.FSName,
		Return: o.Return,
	}
}

func (o *GetFSNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFSNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FSName = op.FSName
	o.Return = op.Return
}
func (o *GetFSNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetFSNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFSNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IsReadableOperation structure represents the CprepDiskIsReadable operation
type xxx_IsReadableOperation struct {
	This   *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	Return int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_IsReadableOperation) OpNum() int { return 37 }

func (o *xxx_IsReadableOperation) OpName() string { return "/IClusterStorage2/v0/CprepDiskIsReadable" }

func (o *xxx_IsReadableOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsReadableOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_IsReadableOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_IsReadableOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsReadableOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_IsReadableOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// IsReadableRequest structure represents the CprepDiskIsReadable operation request
type IsReadableRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk representing the disk.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
}

func (o *IsReadableRequest) xxx_ToOp(ctx context.Context) *xxx_IsReadableOperation {
	if o == nil {
		return &xxx_IsReadableOperation{}
	}
	return &xxx_IsReadableOperation{
		This:   o.This,
		DiskID: o.DiskID,
	}
}

func (o *IsReadableRequest) xxx_FromOp(ctx context.Context, op *xxx_IsReadableOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
}
func (o *IsReadableRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *IsReadableRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsReadableOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsReadableResponse structure represents the CprepDiskIsReadable operation response
type IsReadableResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CprepDiskIsReadable return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *IsReadableResponse) xxx_ToOp(ctx context.Context) *xxx_IsReadableOperation {
	if o == nil {
		return &xxx_IsReadableOperation{}
	}
	return &xxx_IsReadableOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *IsReadableResponse) xxx_FromOp(ctx context.Context, op *xxx_IsReadableOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *IsReadableResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *IsReadableResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsReadableOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDSMsOperation structure represents the CprepDiskGetDsms operation
type xxx_GetDSMsOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	Size            uint32         `idl:"name:Size" json:"size"`
	ResgisteredDSMs uint32         `idl:"name:pResgisteredDsms" json:"resgistered_dsms"`
	RegisteredDSMs  []byte         `idl:"name:RegisteredDsms;size_is:(Size);length_is:(pResgisteredDsms)" json:"registered_dsms"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDSMsOperation) OpNum() int { return 38 }

func (o *xxx_GetDSMsOperation) OpName() string { return "/IClusterStorage2/v0/CprepDiskGetDsms" }

func (o *xxx_GetDSMsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDSMsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Size {in} (1:(uint32))
	{
		if err := w.WriteData(o.Size); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDSMsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Size {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Size); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDSMsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.RegisteredDSMs != nil && o.ResgisteredDSMs == 0 {
		o.ResgisteredDSMs = uint32(len(o.RegisteredDSMs))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDSMsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pResgisteredDsms {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.ResgisteredDSMs); err != nil {
			return err
		}
	}
	// RegisteredDsms {out} (1:{pointer=ref}*(1)[dim:0,size_is=Size,length_is=pResgisteredDsms](uint8))
	{
		dimSize1 := uint64(o.Size)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.ResgisteredDSMs)
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
		for i1 := range o.RegisteredDSMs {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.RegisteredDSMs[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.RegisteredDSMs); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
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

func (o *xxx_GetDSMsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pResgisteredDsms {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.ResgisteredDSMs); err != nil {
			return err
		}
	}
	// RegisteredDsms {out} (1:{pointer=ref}*(1)[dim:0,size_is=Size,length_is=pResgisteredDsms](uint8))
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
			return fmt.Errorf("buffer overflow for size %d of array o.RegisteredDSMs", sizeInfo[0])
		}
		o.RegisteredDSMs = make([]byte, sizeInfo[0])
		for i1 := range o.RegisteredDSMs {
			i1 := i1
			if err := w.ReadData(&o.RegisteredDSMs[i1]); err != nil {
				return err
			}
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

// GetDSMsRequest structure represents the CprepDiskGetDsms operation request
type GetDSMsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// Size: The size, in bytes, of the RegisteredDsms parameter.
	Size uint32 `idl:"name:Size" json:"size"`
}

func (o *GetDSMsRequest) xxx_ToOp(ctx context.Context) *xxx_GetDSMsOperation {
	if o == nil {
		return &xxx_GetDSMsOperation{}
	}
	return &xxx_GetDSMsOperation{
		This: o.This,
		Size: o.Size,
	}
}

func (o *GetDSMsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDSMsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Size = op.Size
}
func (o *GetDSMsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetDSMsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDSMsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDSMsResponse structure represents the CprepDiskGetDsms operation response
type GetDSMsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	ResgisteredDSMs uint32         `idl:"name:pResgisteredDsms" json:"resgistered_dsms"`
	// RegisteredDsms: The buffer that holds the DSM data. The format of the buffer is a
	// REGISTERED_DSMS structure.
	RegisteredDSMs []byte `idl:"name:RegisteredDsms;size_is:(Size);length_is:(pResgisteredDsms)" json:"registered_dsms"`
	// Return: The CprepDiskGetDsms return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDSMsResponse) xxx_ToOp(ctx context.Context) *xxx_GetDSMsOperation {
	if o == nil {
		return &xxx_GetDSMsOperation{}
	}
	return &xxx_GetDSMsOperation{
		That:            o.That,
		ResgisteredDSMs: o.ResgisteredDSMs,
		RegisteredDSMs:  o.RegisteredDSMs,
		Return:          o.Return,
	}
}

func (o *GetDSMsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDSMsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ResgisteredDSMs = op.ResgisteredDSMs
	o.RegisteredDSMs = op.RegisteredDSMs
	o.Return = op.Return
}
func (o *GetDSMsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetDSMsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDSMsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
