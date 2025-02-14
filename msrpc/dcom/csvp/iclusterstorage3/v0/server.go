package iclusterstorage3

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

// IClusterStorage3 server interface.
type ClusterStorage3Server interface {

	// IUnknown base class.
	iunknown.UnknownServer

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
	GetUniqueIDs3(context.Context, *GetUniqueIDs3Request) (*GetUniqueIDs3Response, error)

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
	CheckNetFTBindings3(context.Context, *CheckNetFTBindings3Request) (*CheckNetFTBindings3Response, error)

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
	CSVTestSetup3(context.Context, *CSVTestSetup3Request) (*CSVTestSetup3Response, error)

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
	IsNodeClustered3(context.Context, *IsNodeClustered3Request) (*IsNodeClustered3Response, error)

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
	CreateNewSMBShares3(context.Context, *CreateNewSMBShares3Request) (*CreateNewSMBShares3Response, error)

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
	ConnectToNewSMBShares3(context.Context, *ConnectToNewSMBShares3Request) (*ConnectToNewSMBShares3Response, error)

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
	GetProperties3(context.Context, *GetProperties3Request) (*GetProperties3Response, error)

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
	IsReadOnly3(context.Context, *IsReadOnly3Request) (*IsReadOnly3Response, error)

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
	PRRegister3(context.Context, *PRRegister3Request) (*PRRegister3Response, error)

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
	FindKey3(context.Context, *FindKey3Request) (*FindKey3Response, error)

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
	PRPreempt3(context.Context, *PRPreempt3Request) (*PRPreempt3Response, error)

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
	PRReserve3(context.Context, *PRReserve3Request) (*PRReserve3Response, error)

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
	IsPRPresent3(context.Context, *IsPRPresent3Request) (*IsPRPresent3Response, error)

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
	PRRelease3(context.Context, *PRRelease3Request) (*PRRelease3Response, error)

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
	PRClear3(context.Context, *PRClear3Request) (*PRClear3Response, error)
}

func RegisterClusterStorage3Server(conn dcerpc.Conn, o ClusterStorage3Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewClusterStorage3ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ClusterStorage3SyntaxV0_0))...)
}

func NewClusterStorage3ServerHandle(o ClusterStorage3Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ClusterStorage3ServerHandle(ctx, o, opNum, r)
	}
}

func ClusterStorage3ServerHandle(ctx context.Context, o ClusterStorage3Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // CprepDiskGetUniqueIds3
		op := &xxx_GetUniqueIDs3Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUniqueIDs3Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUniqueIDs3(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // CprepCheckNetFtBindings3
		op := &xxx_CheckNetFTBindings3Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CheckNetFTBindings3Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CheckNetFTBindings3(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // CprepCsvTestSetup3
		op := &xxx_CSVTestSetup3Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CSVTestSetup3Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CSVTestSetup3(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // CprepIsNodeClustered3
		op := &xxx_IsNodeClustered3Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsNodeClustered3Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsNodeClustered3(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // CprepCreateNewSmbShares3
		op := &xxx_CreateNewSMBShares3Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateNewSMBShares3Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateNewSMBShares3(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // CprepConnectToNewSmbShares3
		op := &xxx_ConnectToNewSMBShares3Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ConnectToNewSMBShares3Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ConnectToNewSMBShares3(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // CprepDiskGetProps3
		op := &xxx_GetProperties3Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetProperties3Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetProperties3(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // CprepDiskIsReadOnly3
		op := &xxx_IsReadOnly3Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsReadOnly3Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsReadOnly3(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // CprepDiskPRRegister3
		op := &xxx_PRRegister3Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PRRegister3Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PRRegister3(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // CprepDiskFindKey3
		op := &xxx_FindKey3Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FindKey3Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FindKey3(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // CprepDiskPRPreempt3
		op := &xxx_PRPreempt3Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PRPreempt3Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PRPreempt3(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // CprepDiskPRReserve3
		op := &xxx_PRReserve3Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PRReserve3Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PRReserve3(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // CprepDiskIsPRPresent3
		op := &xxx_IsPRPresent3Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsPRPresent3Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsPRPresent3(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // CprepDiskPRRelease3
		op := &xxx_PRRelease3Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PRRelease3Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PRRelease3(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // CprepDiskPRClear3
		op := &xxx_PRClear3Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PRClear3Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PRClear3(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IClusterStorage3
type UnimplementedClusterStorage3Server struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedClusterStorage3Server) GetUniqueIDs3(context.Context, *GetUniqueIDs3Request) (*GetUniqueIDs3Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage3Server) CheckNetFTBindings3(context.Context, *CheckNetFTBindings3Request) (*CheckNetFTBindings3Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage3Server) CSVTestSetup3(context.Context, *CSVTestSetup3Request) (*CSVTestSetup3Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage3Server) IsNodeClustered3(context.Context, *IsNodeClustered3Request) (*IsNodeClustered3Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage3Server) CreateNewSMBShares3(context.Context, *CreateNewSMBShares3Request) (*CreateNewSMBShares3Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage3Server) ConnectToNewSMBShares3(context.Context, *ConnectToNewSMBShares3Request) (*ConnectToNewSMBShares3Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage3Server) GetProperties3(context.Context, *GetProperties3Request) (*GetProperties3Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage3Server) IsReadOnly3(context.Context, *IsReadOnly3Request) (*IsReadOnly3Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage3Server) PRRegister3(context.Context, *PRRegister3Request) (*PRRegister3Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage3Server) FindKey3(context.Context, *FindKey3Request) (*FindKey3Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage3Server) PRPreempt3(context.Context, *PRPreempt3Request) (*PRPreempt3Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage3Server) PRReserve3(context.Context, *PRReserve3Request) (*PRReserve3Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage3Server) IsPRPresent3(context.Context, *IsPRPresent3Request) (*IsPRPresent3Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage3Server) PRRelease3(context.Context, *PRRelease3Request) (*PRRelease3Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterStorage3Server) PRClear3(context.Context, *PRClear3Request) (*PRClear3Response, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ClusterStorage3Server = (*UnimplementedClusterStorage3Server)(nil)
