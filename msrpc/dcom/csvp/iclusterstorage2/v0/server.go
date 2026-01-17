package iclusterstorage2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
	_ = iunknown.GoPackage
)

// IClusterStorage2 server interface.
type ClusterStorage2Server interface {

	// IUnknown base class.
	iunknown.UnknownServer

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
	RawRead(context.Context, *RawReadRequest) (*RawReadResponse, error)

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
	RawWrite(context.Context, *RawWriteRequest) (*RawWriteResponse, error)

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
	PrepareNode(context.Context, *PrepareNodeRequest) (*PrepareNodeResponse, error)

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
	PrepareNodePhase2(context.Context, *PrepareNodePhase2Request) (*PrepareNodePhase2Response, error)

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
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)

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
	StopDefense(context.Context, *StopDefenseRequest) (*StopDefenseResponse, error)

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
	Online(context.Context, *OnlineRequest) (*OnlineResponse, error)

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
	VerifyUnique(context.Context, *VerifyUniqueRequest) (*VerifyUniqueResponse, error)

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
	WriteFileData(context.Context, *WriteFileDataRequest) (*WriteFileDataResponse, error)

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
	VerifyFileData(context.Context, *VerifyFileDataRequest) (*VerifyFileDataResponse, error)

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
	DeleteFile(context.Context, *DeleteFileRequest) (*DeleteFileResponse, error)

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
	Offline(context.Context, *OfflineRequest) (*OfflineResponse, error)

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
	GetUniqueIDs(context.Context, *GetUniqueIDsRequest) (*GetUniqueIDsResponse, error)

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
	Attach(context.Context, *AttachRequest) (*AttachResponse, error)

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
	PRArbitrate(context.Context, *PRArbitrateRequest) (*PRArbitrateResponse, error)

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
	PRRegister(context.Context, *PRRegisterRequest) (*PRRegisterResponse, error)

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
	PRUnregister(context.Context, *PRUnregisterRequest) (*PRUnregisterResponse, error)

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
	PRReserve(context.Context, *PRReserveRequest) (*PRReserveResponse, error)

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
	PRRelease(context.Context, *PRReleaseRequest) (*PRReleaseResponse, error)

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
	DiskPartitionIsNTFS(context.Context, *DiskPartitionIsNTFSRequest) (*DiskPartitionIsNTFSResponse, error)

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
	GetArbSectors(context.Context, *GetArbSectorsRequest) (*GetArbSectorsResponse, error)

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
	IsPRPresent(context.Context, *IsPRPresentRequest) (*IsPRPresentResponse, error)

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
	PRPreempt(context.Context, *PRPreemptRequest) (*PRPreemptResponse, error)

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
	PRClear(context.Context, *PRClearRequest) (*PRClearResponse, error)

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
	IsOnline(context.Context, *IsOnlineRequest) (*IsOnlineResponse, error)

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
	SetOnline(context.Context, *SetOnlineRequest) (*SetOnlineResponse, error)

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
	GetFSName(context.Context, *GetFSNameRequest) (*GetFSNameResponse, error)

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
	IsReadable(context.Context, *IsReadableRequest) (*IsReadableResponse, error)

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
	GetDSMs(context.Context, *GetDSMsRequest) (*GetDSMsResponse, error)
}

func RegisterClusterStorage2Server(conn dcerpc.Conn, o ClusterStorage2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewClusterStorage2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ClusterStorage2SyntaxV0_0))...)
}

func NewClusterStorage2ServerHandle(o ClusterStorage2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ClusterStorage2ServerHandle(ctx, o, opNum, r)
	}
}

func ClusterStorage2ServerHandle(ctx context.Context, o ClusterStorage2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // CprepDiskRawRead
		op := &xxx_RawReadOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RawReadRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RawRead(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // CprepDiskRawWrite
		op := &xxx_RawWriteOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RawWriteRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RawWrite(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // CprepPrepareNode
		op := &xxx_PrepareNodeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PrepareNodeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PrepareNode(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // CprepPrepareNodePhase2
		op := &xxx_PrepareNodePhase2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PrepareNodePhase2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PrepareNodePhase2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // CprepDiskGetProps
		op := &xxx_GetPropertiesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPropertiesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetProperties(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // Opnum8NotUsedOnWire
		// Opnum8NotUsedOnWire
		return nil, nil
	case 9: // Opnum9NotUsedOnWire
		// Opnum9NotUsedOnWire
		return nil, nil
	case 10: // Opnum10NotUsedOnWire
		// Opnum10NotUsedOnWire
		return nil, nil
	case 11: // Opnum11NotUsedOnWire
		// Opnum11NotUsedOnWire
		return nil, nil
	case 12: // CprepDiskStopDefense
		op := &xxx_StopDefenseOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StopDefenseRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.StopDefense(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // CprepDiskOnline
		op := &xxx_OnlineOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OnlineRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Online(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // CprepDiskVerifyUnique
		op := &xxx_VerifyUniqueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &VerifyUniqueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.VerifyUnique(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // Opnum15NotUsedOnWire
		// Opnum15NotUsedOnWire
		return nil, nil
	case 16: // Opnum16NotUsedOnWire
		// Opnum16NotUsedOnWire
		return nil, nil
	case 17: // CprepDiskWriteFileData
		op := &xxx_WriteFileDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &WriteFileDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.WriteFileData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // CprepDiskVerifyFileData
		op := &xxx_VerifyFileDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &VerifyFileDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.VerifyFileData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // CprepDiskDeleteFile
		op := &xxx_DeleteFileOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteFileRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteFile(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // CprepDiskOffline
		op := &xxx_OfflineOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OfflineRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Offline(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // Opnum21NotUsedOnWire
		// Opnum21NotUsedOnWire
		return nil, nil
	case 22: // CprepDiskGetUniqueIds
		op := &xxx_GetUniqueIDsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUniqueIDsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUniqueIDs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // CprepDiskAttach
		op := &xxx_AttachOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AttachRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Attach(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // CprepDiskPRArbitrate
		op := &xxx_PRArbitrateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PRArbitrateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PRArbitrate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // CprepDiskPRRegister
		op := &xxx_PRRegisterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PRRegisterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PRRegister(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // CprepDiskPRUnRegister
		op := &xxx_PRUnregisterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PRUnregisterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PRUnregister(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // CprepDiskPRReserve
		op := &xxx_PRReserveOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PRReserveRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PRReserve(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // CprepDiskPRRelease
		op := &xxx_PRReleaseOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PRReleaseRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PRRelease(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // CprepDiskDiskPartitionIsNtfs
		op := &xxx_DiskPartitionIsNTFSOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DiskPartitionIsNTFSRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DiskPartitionIsNTFS(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // CprepDiskGetArbSectors
		op := &xxx_GetArbSectorsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetArbSectorsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetArbSectors(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 31: // CprepDiskIsPRPresent
		op := &xxx_IsPRPresentOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsPRPresentRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsPRPresent(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 32: // CprepDiskPRPreempt
		op := &xxx_PRPreemptOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PRPreemptRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PRPreempt(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // CprepDiskPRClear
		op := &xxx_PRClearOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PRClearRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PRClear(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // CprepDiskIsOnline
		op := &xxx_IsOnlineOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsOnlineRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsOnline(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 35: // CprepDiskSetOnline
		op := &xxx_SetOnlineOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetOnlineRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetOnline(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 36: // CprepDiskGetFSName
		op := &xxx_GetFSNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFSNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFSName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 37: // CprepDiskIsReadable
		op := &xxx_IsReadableOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsReadableRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsReadable(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 38: // CprepDiskGetDsms
		op := &xxx_GetDSMsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDSMsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDSMs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IClusterStorage2
type UnimplementedClusterStorage2Server struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedClusterStorage2Server) RawRead(context.Context, *RawReadRequest) (*RawReadResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage2Server) RawWrite(context.Context, *RawWriteRequest) (*RawWriteResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage2Server) PrepareNode(context.Context, *PrepareNodeRequest) (*PrepareNodeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage2Server) PrepareNodePhase2(context.Context, *PrepareNodePhase2Request) (*PrepareNodePhase2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage2Server) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage2Server) StopDefense(context.Context, *StopDefenseRequest) (*StopDefenseResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage2Server) Online(context.Context, *OnlineRequest) (*OnlineResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage2Server) VerifyUnique(context.Context, *VerifyUniqueRequest) (*VerifyUniqueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage2Server) WriteFileData(context.Context, *WriteFileDataRequest) (*WriteFileDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage2Server) VerifyFileData(context.Context, *VerifyFileDataRequest) (*VerifyFileDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage2Server) DeleteFile(context.Context, *DeleteFileRequest) (*DeleteFileResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage2Server) Offline(context.Context, *OfflineRequest) (*OfflineResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage2Server) GetUniqueIDs(context.Context, *GetUniqueIDsRequest) (*GetUniqueIDsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage2Server) Attach(context.Context, *AttachRequest) (*AttachResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage2Server) PRArbitrate(context.Context, *PRArbitrateRequest) (*PRArbitrateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage2Server) PRRegister(context.Context, *PRRegisterRequest) (*PRRegisterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage2Server) PRUnregister(context.Context, *PRUnregisterRequest) (*PRUnregisterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage2Server) PRReserve(context.Context, *PRReserveRequest) (*PRReserveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage2Server) PRRelease(context.Context, *PRReleaseRequest) (*PRReleaseResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage2Server) DiskPartitionIsNTFS(context.Context, *DiskPartitionIsNTFSRequest) (*DiskPartitionIsNTFSResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage2Server) GetArbSectors(context.Context, *GetArbSectorsRequest) (*GetArbSectorsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage2Server) IsPRPresent(context.Context, *IsPRPresentRequest) (*IsPRPresentResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage2Server) PRPreempt(context.Context, *PRPreemptRequest) (*PRPreemptResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage2Server) PRClear(context.Context, *PRClearRequest) (*PRClearResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage2Server) IsOnline(context.Context, *IsOnlineRequest) (*IsOnlineResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage2Server) SetOnline(context.Context, *SetOnlineRequest) (*SetOnlineResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage2Server) GetFSName(context.Context, *GetFSNameRequest) (*GetFSNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage2Server) IsReadable(context.Context, *IsReadableRequest) (*IsReadableResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage2Server) GetDSMs(context.Context, *GetDSMsRequest) (*GetDSMsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ClusterStorage2Server = (*UnimplementedClusterStorage2Server)(nil)
