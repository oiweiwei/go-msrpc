package iclusterstorage2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
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
		in := &RawReadRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RawRead(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // CprepDiskRawWrite
		in := &RawWriteRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RawWrite(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // CprepPrepareNode
		in := &PrepareNodeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PrepareNode(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // CprepPrepareNodePhase2
		in := &PrepareNodePhase2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PrepareNodePhase2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // CprepDiskGetProps
		in := &GetPropertiesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetProperties(ctx, in)
		return resp.xxx_ToOp(ctx), err
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
		in := &StopDefenseRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.StopDefense(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // CprepDiskOnline
		in := &OnlineRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Online(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // CprepDiskVerifyUnique
		in := &VerifyUniqueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.VerifyUnique(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // Opnum15NotUsedOnWire
		// Opnum15NotUsedOnWire
		return nil, nil
	case 16: // Opnum16NotUsedOnWire
		// Opnum16NotUsedOnWire
		return nil, nil
	case 17: // CprepDiskWriteFileData
		in := &WriteFileDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.WriteFileData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // CprepDiskVerifyFileData
		in := &VerifyFileDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.VerifyFileData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // CprepDiskDeleteFile
		in := &DeleteFileRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteFile(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // CprepDiskOffline
		in := &OfflineRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Offline(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // Opnum21NotUsedOnWire
		// Opnum21NotUsedOnWire
		return nil, nil
	case 22: // CprepDiskGetUniqueIds
		in := &GetUniqueIDsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetUniqueIDs(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // CprepDiskAttach
		in := &AttachRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Attach(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // CprepDiskPRArbitrate
		in := &PRArbitrateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PRArbitrate(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // CprepDiskPRRegister
		in := &PRRegisterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PRRegister(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 26: // CprepDiskPRUnRegister
		in := &PRUnregisterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PRUnregister(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 27: // CprepDiskPRReserve
		in := &PRReserveRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PRReserve(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 28: // CprepDiskPRRelease
		in := &PRReleaseRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PRRelease(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 29: // CprepDiskDiskPartitionIsNtfs
		in := &DiskPartitionIsNTFSRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DiskPartitionIsNTFS(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 30: // CprepDiskGetArbSectors
		in := &GetArbSectorsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetArbSectors(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 31: // CprepDiskIsPRPresent
		in := &IsPRPresentRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.IsPRPresent(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 32: // CprepDiskPRPreempt
		in := &PRPreemptRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PRPreempt(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 33: // CprepDiskPRClear
		in := &PRClearRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PRClear(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 34: // CprepDiskIsOnline
		in := &IsOnlineRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.IsOnline(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 35: // CprepDiskSetOnline
		in := &SetOnlineRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetOnline(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 36: // CprepDiskGetFSName
		in := &GetFSNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFSName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 37: // CprepDiskIsReadable
		in := &IsReadableRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.IsReadable(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 38: // CprepDiskGetDsms
		in := &GetDSMsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDSMs(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
