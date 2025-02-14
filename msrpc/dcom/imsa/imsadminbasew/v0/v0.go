package imsadminbasew

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	imsa "github.com/oiweiwei/go-msrpc/msrpc/dcom/imsa"
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
	_ = imsa.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/imsa"
)

var (
	// IMSAdminBaseW interface identifier 70b51430-b6ca-11d0-b9b9-00a0c922e750
	IMSAdminBaseWIID = &dcom.IID{Data1: 0x70b51430, Data2: 0xb6ca, Data3: 0x11d0, Data4: []byte{0xb9, 0xb9, 0x00, 0xa0, 0xc9, 0x22, 0xe7, 0x50}}
	// Syntax UUID
	IMSAdminBaseWSyntaxUUID = &uuid.UUID{TimeLow: 0x70b51430, TimeMid: 0xb6ca, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xb9, ClockSeqLow: 0xb9, Node: [6]uint8{0x0, 0xa0, 0xc9, 0x22, 0xe7, 0x50}}
	// Syntax ID
	IMSAdminBaseWSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: IMSAdminBaseWSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IMSAdminBaseW interface.
type IMSAdminBaseWClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The AddKey method creates a node and adds it to the metabase as a subnode of an existing
	// node at the specified path.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+---------------------------------+--------------------------------------------------------+
	//	|             RETURN              |                                                        |
	//	|           VALUE/CODE            |                      DESCRIPTION                       |
	//	|                                 |                                                        |
	//	+---------------------------------+--------------------------------------------------------+
	//	+---------------------------------+--------------------------------------------------------+
	//	| 0x00000000 S_OK                 | The call was successful.                               |
	//	+---------------------------------+--------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED       | General access denied error.                           |
	//	+---------------------------------+--------------------------------------------------------+
	//	| 0x80070006 ERROR_INVALID_HANDLE | The handle is invalid.                                 |
	//	+---------------------------------+--------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG         | One or more arguments are invalid.                     |
	//	+---------------------------------+--------------------------------------------------------+
	//	| 0x800700B7 ERROR_ALREADY_EXISTS | Cannot create a file because that file already exists. |
	//	+---------------------------------+--------------------------------------------------------+
	//
	// The opnum field value for this method is 3.
	AddKey(context.Context, *AddKeyRequest, ...dcerpc.CallOption) (*AddKeyResponse, error)

	// The DeleteKey method deletes a node and all its data from the metabase. All of the
	// subnodes are recursively deleted.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+---------------------------------+--------------------------------------------+
	//	|             RETURN              |                                            |
	//	|           VALUE/CODE            |                DESCRIPTION                 |
	//	|                                 |                                            |
	//	+---------------------------------+--------------------------------------------+
	//	+---------------------------------+--------------------------------------------+
	//	| 0x00000000 S_OK                 | The call was successful.                   |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070003 ERROR_PATH_NOT_FOUND | The system cannot find the path specified. |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED       | General access denied error.               |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070006 ERROR_INVALID_HANDLE | The handle is invalid.                     |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070057 E_INVALIDARG         | One or more arguments are invalid.         |
	//	+---------------------------------+--------------------------------------------+
	//
	// The opnum field value for this method is 4.
	DeleteKey(context.Context, *DeleteKeyRequest, ...dcerpc.CallOption) (*DeleteKeyResponse, error)

	// The DeleteChildKeys method deletes all subnodes of the specified node and any data
	// they contain. It also recursively deletes all nodes below the subnodes.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+---------------------------------+--------------------------------------------+
	//	|             RETURN              |                                            |
	//	|           VALUE/CODE            |                DESCRIPTION                 |
	//	|                                 |                                            |
	//	+---------------------------------+--------------------------------------------+
	//	+---------------------------------+--------------------------------------------+
	//	| 0x00000000 S_OK                 | The call was successful.                   |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070003 ERROR_PATH_NOT_FOUND | The system cannot find the path specified. |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED       | General access denied error.               |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070006 ERROR_INVALID_HANDLE | The handle is invalid.                     |
	//	+---------------------------------+--------------------------------------------+
	//
	// The opnum field value for this method is 5.
	DeleteChildKeys(context.Context, *DeleteChildKeysRequest, ...dcerpc.CallOption) (*DeleteChildKeysResponse, error)

	// The EnumKeys method enumerates the subnodes of the specified node.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+---------------------------------+--------------------------------------------+
	//	|             RETURN              |                                            |
	//	|           VALUE/CODE            |                DESCRIPTION                 |
	//	|                                 |                                            |
	//	+---------------------------------+--------------------------------------------+
	//	+---------------------------------+--------------------------------------------+
	//	| 0x00000000 S_OK                 | The call was successful.                   |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070003 ERROR_PATH_NOT_FOUND | The system cannot find the path specified. |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED       | General access denied error.               |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070006 ERROR_INVALID_HANDLE | The handle is invalid.                     |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070057 E_INVALIDARG         | One or more arguments are invalid.         |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070103 ERROR_NO_MORE_ITEMS  | No more data is available.                 |
	//	+---------------------------------+--------------------------------------------+
	//
	// The opnum field value for this method is 6.
	//
	// A subnode can be enumerated once per call. Subnodes are numbered from zero to (NumKeys
	// - 1), with NumKeys equal to the number of subnodes below the node.
	EnumKeys(context.Context, *EnumKeysRequest, ...dcerpc.CallOption) (*EnumKeysResponse, error)

	// The CopyKey method copies or moves a node, including its subnodes and data, to a
	// specified destination. The copied or moved node becomes a subnode of the destination
	// node.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+---------------------------------+--------------------------------------------+
	//	|             RETURN              |                                            |
	//	|           VALUE/CODE            |                DESCRIPTION                 |
	//	|                                 |                                            |
	//	+---------------------------------+--------------------------------------------+
	//	+---------------------------------+--------------------------------------------+
	//	| 0x00000000 S_OK                 | The call was successful.                   |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070003 ERROR_PATH_NOT_FOUND | The system cannot find the path specified. |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED       | General access denied error.               |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070006 ERROR_INVALID_HANDLE | The handle is invalid.                     |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070057 E_INVALIDARG         | One or more arguments are invalid.         |
	//	+---------------------------------+--------------------------------------------+
	//
	// The opnum field value for this method is 7.
	CopyKey(context.Context, *CopyKeyRequest, ...dcerpc.CallOption) (*CopyKeyResponse, error)

	// The RenameKey method renames a node in the metabase.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+---------------------------------+----------------------------------------------------+
	//	|             RETURN              |                                                    |
	//	|           VALUE/CODE            |                    DESCRIPTION                     |
	//	|                                 |                                                    |
	//	+---------------------------------+----------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------+
	//	| 0x00000000 S_OK                 | The call was successful.                           |
	//	+---------------------------------+----------------------------------------------------+
	//	| 0x80070003 ERROR_PATH_NOT_FOUND | The system cannot find the path specified.         |
	//	+---------------------------------+----------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED       | General access denied error.                       |
	//	+---------------------------------+----------------------------------------------------+
	//	| 0x80070006 ERROR_INVALID_HANDLE | The handle is invalid.                             |
	//	+---------------------------------+----------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG         | An invalid parameter value was specified.          |
	//	+---------------------------------+----------------------------------------------------+
	//	| 0x800700B7 ERROR_ALREADY_EXISTS | A key of that name already exists in the database. |
	//	+---------------------------------+----------------------------------------------------+
	//
	// The opnum field value for this method is 8.
	RenameKey(context.Context, *RenameKeyRequest, ...dcerpc.CallOption) (*RenameKeyResponse, error)

	// The R_SetData method sets a data item for a particular node in the metabase.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16 27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                       RETURN                       |                                                                                  |
	//	|                     VALUE/CODE                     |                                   DESCRIPTION                                    |
	//	|                                                    |                                                                                  |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                                    | The call was successful.                                                         |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070003 ERROR_PATH_NOT_FOUND                    | The system cannot find the path specified.                                       |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED                          | General access denied error.                                                     |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                            | An invalid parameter value was specified.                                        |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800CC808 MD_ERROR_CANNOT_REMOVE_SECURE_ATTRIBUTE | The METADATA_SECURE attribute cannot be removed from a data item via the         |
	//	|                                                    | R_GetData method. Use the DeleteData method to remove the secure data.           |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY                 | There is not enough memory to complete the operation.                            |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 9.
	SetData(context.Context, *SetDataRequest, ...dcerpc.CallOption) (*SetDataResponse, error)

	// The R_GetData method returns a data entry from a particular node in the metabase.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+--------------------------------------+-----------------------------------------------------+
	//	|                RETURN                |                                                     |
	//	|              VALUE/CODE              |                     DESCRIPTION                     |
	//	|                                      |                                                     |
	//	+--------------------------------------+-----------------------------------------------------+
	//	+--------------------------------------+-----------------------------------------------------+
	//	| 0x00000000 S_OK                      | The call was successful.                            |
	//	+--------------------------------------+-----------------------------------------------------+
	//	| 0x80070003 ERROR_PATH_NOT_FOUND      | The system cannot find the path specified.          |
	//	+--------------------------------------+-----------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED            | General access denied error.                        |
	//	+--------------------------------------+-----------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG              | An invalid parameter value was specified.           |
	//	+--------------------------------------+-----------------------------------------------------+
	//	| 0x8007007A ERROR_INSUFFICIENT_BUFFER | The data area passed to a system call is too small. |
	//	+--------------------------------------+-----------------------------------------------------+
	//	| 0x800CC801 MD_ERROR_DATA_NOT_FOUND   | The specified metadata was not found.               |
	//	+--------------------------------------+-----------------------------------------------------+
	//
	// The opnum field value for this method is 10.
	//
	// The client describes the data it is requesting by initializing the METADATA_RECORD
	// passed in the pmdrMDData parameter.
	//
	// The client indicates how much decrypted data it is ready to receive by passing the
	// number of bytes in the dwMDDataLen field of the pmdrMDData parameter.
	//
	// The pbMDData field of the pmdrMDData parameter is not used to transfer the actual
	// data value. The client MUST set the pbMDData field of pmdrMDData to NULL. The IIS_CRYPTO_BLOB
	// structure is used to transfer the actual data value returned by the server and can
	// be encrypted when the server sends data marked as secure.
	GetData(context.Context, *GetDataRequest, ...dcerpc.CallOption) (*GetDataResponse, error)

	// The DeleteData method deletes specific data entries from a node in the metabase.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+------------------------------------+--------------------------------------------+
	//	|               RETURN               |                                            |
	//	|             VALUE/CODE             |                DESCRIPTION                 |
	//	|                                    |                                            |
	//	+------------------------------------+--------------------------------------------+
	//	+------------------------------------+--------------------------------------------+
	//	| 0x00000000 S_OK                    | The call was successful.                   |
	//	+------------------------------------+--------------------------------------------+
	//	| 0x80070003 ERROR_PATH_NOT_FOUND    | The system cannot find the path specified. |
	//	+------------------------------------+--------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED          | General access denied error.               |
	//	+------------------------------------+--------------------------------------------+
	//	| 0x80070006 ERROR_INVALID_HANDLE    | The handle is invalid.                     |
	//	+------------------------------------+--------------------------------------------+
	//	| 0x800CC801 MD_ERROR_DATA_NOT_FOUND | The specified metadata was not found.      |
	//	+------------------------------------+--------------------------------------------+
	//
	// The opnum field value for this method is 11.
	DeleteData(context.Context, *DeleteDataRequest, ...dcerpc.CallOption) (*DeleteDataResponse, error)

	// The R_EnumData method enumerates the data entries of a node in the metabase.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+--------------------------------------+-----------------------------------------------------+
	//	|                RETURN                |                                                     |
	//	|              VALUE/CODE              |                     DESCRIPTION                     |
	//	|                                      |                                                     |
	//	+--------------------------------------+-----------------------------------------------------+
	//	+--------------------------------------+-----------------------------------------------------+
	//	| 0x00000000 S_OK                      | The call was successful.                            |
	//	+--------------------------------------+-----------------------------------------------------+
	//	| 0x80070003 ERROR_PATH_NOT_FOUND      | The system cannot find the path specified.          |
	//	+--------------------------------------+-----------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED            | General access denied error.                        |
	//	+--------------------------------------+-----------------------------------------------------+
	//	| 0x80070006 ERROR_INVALID_HANDLE      | The handle is invalid.                              |
	//	+--------------------------------------+-----------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG              | One or more arguments are invalid.                  |
	//	+--------------------------------------+-----------------------------------------------------+
	//	| 0x8007007A ERROR_INSUFFICIENT_BUFFER | The data area passed to a system call is too small. |
	//	+--------------------------------------+-----------------------------------------------------+
	//	| 0x80070103 ERROR_NO_MORE_ITEMS       | No more data is available.                          |
	//	+--------------------------------------+-----------------------------------------------------+
	//
	// The opnum field value for this method is 12.
	//
	// The client indicates how much decrypted data it is ready to receive by passing the
	// number of bytes in the dwMDDataLen field of the pmdrMDData parameter. If this value
	// is too small to contain the decrypted data value, the server MUST return ERROR_INSUFFICIENT_BUFFER
	// and return the number of bytes required to hold the data in the pdwMDRequiredDataLen
	// parameter.
	EnumData(context.Context, *EnumDataRequest, ...dcerpc.CallOption) (*EnumDataResponse, error)

	// The R_GetAllData method returns all data associated with a node in the metabase,
	// including all values that the node inherits.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	// Note  Invalid dwMDUserType or dwMDDataType parameters result in a E_INVALIDARG return
	// status.
	//
	//	+--------------------------------------+-----------------------------------------------------+
	//	|                RETURN                |                                                     |
	//	|              VALUE/CODE              |                     DESCRIPTION                     |
	//	|                                      |                                                     |
	//	+--------------------------------------+-----------------------------------------------------+
	//	+--------------------------------------+-----------------------------------------------------+
	//	| 0x00000000 S_OK                      | The call was successful.                            |
	//	+--------------------------------------+-----------------------------------------------------+
	//	| 0x80070003 ERROR_PATH_NOT_FOUND      | The system cannot find the path specified.          |
	//	+--------------------------------------+-----------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED            | General access denied error.                        |
	//	+--------------------------------------+-----------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG              | An invalid parameter value was specified.           |
	//	+--------------------------------------+-----------------------------------------------------+
	//	| 0x8007007A ERROR_INSUFFICIENT_BUFFER | The data area passed to a system call is too small. |
	//	+--------------------------------------+-----------------------------------------------------+
	//
	// The opnum field value for this method is 13.
	//
	// The client indicates how much decrypted data it is ready to receive by passing the
	// number of bytes in the dwMDBufferSize parameter.
	GetAllData(context.Context, *GetAllDataRequest, ...dcerpc.CallOption) (*GetAllDataResponse, error)

	// The DeleteAllData method deletes all or a subset of local data associated with a
	// particular node.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+---------------------------------+------------------------------+
	//	|             RETURN              |                              |
	//	|           VALUE/CODE            |         DESCRIPTION          |
	//	|                                 |                              |
	//	+---------------------------------+------------------------------+
	//	+---------------------------------+------------------------------+
	//	| 0x00000000 S_OK                 | The call was successful.     |
	//	+---------------------------------+------------------------------+
	//	| 0x80070005 E_ACCESSDENIED       | General access denied error. |
	//	+---------------------------------+------------------------------+
	//	| 0x80070006 ERROR_INVALID_HANDLE | The handle is invalid.       |
	//	+---------------------------------+------------------------------+
	//
	// The opnum field value for this method is 14.
	DeleteAllData(context.Context, *DeleteAllDataRequest, ...dcerpc.CallOption) (*DeleteAllDataResponse, error)

	// The CopyData method copies or moves data between nodes.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+------------------------------------+-------------------------------------------------------+
	//	|               RETURN               |                                                       |
	//	|             VALUE/CODE             |                      DESCRIPTION                      |
	//	|                                    |                                                       |
	//	+------------------------------------+-------------------------------------------------------+
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x00000000 S_OK                    | The call was successful.                              |
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x80070003 ERROR_PATH_NOT_FOUND    | The system cannot find the path specified.            |
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED          | General access denied error.                          |
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x80070006 ERROR_INVALID_HANDLE    | The handle is invalid.                                |
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG            | One or more arguments are invalid.                    |
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY | There is not enough memory to complete the operation. |
	//	+------------------------------------+-------------------------------------------------------+
	//
	// The opnum field value for this method is 15.
	CopyData(context.Context, *CopyDataRequest, ...dcerpc.CallOption) (*CopyDataResponse, error)

	// The GetDataPaths method returns the paths of all nodes in the subtree relative to
	// a specified starting node that contains the supplied identifier.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+--------------------------------------+-----------------------------------------------------+
	//	|                RETURN                |                                                     |
	//	|              VALUE/CODE              |                     DESCRIPTION                     |
	//	|                                      |                                                     |
	//	+--------------------------------------+-----------------------------------------------------+
	//	+--------------------------------------+-----------------------------------------------------+
	//	| 0x00000000 S_OK                      | The call was successful.                            |
	//	+--------------------------------------+-----------------------------------------------------+
	//	| 0x80070003 ERROR_PATH_NOT_FOUND      | The system cannot find the path specified.          |
	//	+--------------------------------------+-----------------------------------------------------+
	//	| 0x80070006 ERROR_INVALID_HANDLE      | The handle is invalid.                              |
	//	+--------------------------------------+-----------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG              | One or more arguments are invalid.                  |
	//	+--------------------------------------+-----------------------------------------------------+
	//	| 0x8007007A ERROR_INSUFFICIENT_BUFFER | The data area passed to a system call is too small. |
	//	+--------------------------------------+-----------------------------------------------------+
	//
	// The opnum field value for this method is 16.
	GetDataPaths(context.Context, *GetDataPathsRequest, ...dcerpc.CallOption) (*GetDataPathsResponse, error)

	// The OpenKey method opens a node for read access, write access, or both. The returned
	// handle can be used by several of the other methods in the IMSAdminBaseW interface.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+---------------------------------+-------------------------------------------------+
	//	|             RETURN              |                                                 |
	//	|           VALUE/CODE            |                   DESCRIPTION                   |
	//	|                                 |                                                 |
	//	+---------------------------------+-------------------------------------------------+
	//	+---------------------------------+-------------------------------------------------+
	//	| 0x00000000 S_OK                 | The call was successful.                        |
	//	+---------------------------------+-------------------------------------------------+
	//	| 0x80070003 ERROR_PATH_NOT_FOUND | The system cannot find the path specified.      |
	//	+---------------------------------+-------------------------------------------------+
	//	| 0x80070006 ERROR_INVALID_HANDLE | The handle is invalid.                          |
	//	+---------------------------------+-------------------------------------------------+
	//	| 0x80070094 ERROR_PATH_BUSY      | The path specified cannot be used at this time. |
	//	+---------------------------------+-------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG         | One or more arguments are invalid.              |
	//	+---------------------------------+-------------------------------------------------+
	//
	// The opnum field value for this method is 17.
	OpenKey(context.Context, *OpenKeyRequest, ...dcerpc.CallOption) (*OpenKeyResponse, error)

	// The CloseKey method closes a handle to a node.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+---------------------------------+--------------------------+
	//	|             RETURN              |                          |
	//	|           VALUE/CODE            |       DESCRIPTION        |
	//	|                                 |                          |
	//	+---------------------------------+--------------------------+
	//	+---------------------------------+--------------------------+
	//	| 0x00000000 S_OK                 | The call was successful. |
	//	+---------------------------------+--------------------------+
	//	| 0x80070006 ERROR_INVALID_HANDLE | The handle is invalid.   |
	//	+---------------------------------+--------------------------+
	//
	// The opnum field value for this method is 18.
	CloseKey(context.Context, *CloseKeyRequest, ...dcerpc.CallOption) (*CloseKeyResponse, error)

	// The ChangePermissions method changes permissions on an open handle.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+---------------------------------+-------------------------------------------------+
	//	|             RETURN              |                                                 |
	//	|           VALUE/CODE            |                   DESCRIPTION                   |
	//	|                                 |                                                 |
	//	+---------------------------------+-------------------------------------------------+
	//	+---------------------------------+-------------------------------------------------+
	//	| 0x00000000 S_OK                 | The call was successful.                        |
	//	+---------------------------------+-------------------------------------------------+
	//	| 0x80070006 ERROR_INVALID_HANDLE | The handle is invalid.                          |
	//	+---------------------------------+-------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG         | One or more arguments are invalid.              |
	//	+---------------------------------+-------------------------------------------------+
	//	| 0x80070094 ERROR_PATH_BUSY      | The path specified cannot be used at this time. |
	//	+---------------------------------+-------------------------------------------------+
	//
	// The opnum field value for this method is 19.
	ChangePermissions(context.Context, *ChangePermissionsRequest, ...dcerpc.CallOption) (*ChangePermissionsResponse, error)

	// The SaveData method explicitly flushes the metabase data resident in memory to configuration
	// storage.
	//
	// This method has no parameters.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+----------------------------+-------------------------------------------------+
	//	|           RETURN           |                                                 |
	//	|         VALUE/CODE         |                   DESCRIPTION                   |
	//	|                            |                                                 |
	//	+----------------------------+-------------------------------------------------+
	//	+----------------------------+-------------------------------------------------+
	//	| 0x00000000 S_OK            | The call was successful.                        |
	//	+----------------------------+-------------------------------------------------+
	//	| 0x80070094 ERROR_PATH_BUSY | The path specified cannot be used at this time. |
	//	+----------------------------+-------------------------------------------------+
	//
	// The opnum field value for this method is 20.
	SaveData(context.Context, *SaveDataRequest, ...dcerpc.CallOption) (*SaveDataResponse, error)

	// The GetHandleInfo method returns information associated with the specified metabase
	// handle.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+---------------------------------+--------------------------+
	//	|             RETURN              |                          |
	//	|           VALUE/CODE            |       DESCRIPTION        |
	//	|                                 |                          |
	//	+---------------------------------+--------------------------+
	//	+---------------------------------+--------------------------+
	//	| 0x00000000 S_OK                 | The call was successful. |
	//	+---------------------------------+--------------------------+
	//	| 0x80070006 ERROR_INVALID_HANDLE | The handle is invalid.   |
	//	+---------------------------------+--------------------------+
	//
	// The opnum field value for this method is 21.
	GetHandleInfo(context.Context, *GetHandleInfoRequest, ...dcerpc.CallOption) (*GetHandleInfoResponse, error)

	// The GetSystemChangeNumber method returns the number of changes made to data since
	// the metabase was created.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
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
	// The opnum field value for this method is 22.
	GetSystemChangeNumber(context.Context, *GetSystemChangeNumberRequest, ...dcerpc.CallOption) (*GetSystemChangeNumberResponse, error)

	// The GetDataSetNumber method returns the dataset number associated with a node in
	// the metabase. A dataset number is a unique number identifying the data items at that
	// node, including inherited data items. Nodes with the same dataset number have identical
	// data.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
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
	// The opnum field value for this method is 23.
	GetDataSetNumber(context.Context, *GetDataSetNumberRequest, ...dcerpc.CallOption) (*GetDataSetNumberResponse, error)

	// The SetLastChangeTime method sets the last change time associated with a node in
	// the metabase.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+---------------------------------+--------------------------------------------+
	//	|             RETURN              |                                            |
	//	|           VALUE/CODE            |                DESCRIPTION                 |
	//	|                                 |                                            |
	//	+---------------------------------+--------------------------------------------+
	//	+---------------------------------+--------------------------------------------+
	//	| 0x00000000 S_OK                 | The call was successful.                   |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070003 ERROR_PATH_NOT_FOUND | The system cannot find the path specified. |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED       | General access denied error.               |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070057 E_INVALIDARG         | One or more arguments are invalid.         |
	//	+---------------------------------+--------------------------------------------+
	//
	// The opnum field value for this method is 24.
	SetLastChangeTime(context.Context, *SetLastChangeTimeRequest, ...dcerpc.CallOption) (*SetLastChangeTimeResponse, error)

	// The GetLastChangeTime method returns the last change time associated with a node
	// in the metabase.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+---------------------------------+--------------------------------------------+
	//	|             RETURN              |                                            |
	//	|           VALUE/CODE            |                DESCRIPTION                 |
	//	|                                 |                                            |
	//	+---------------------------------+--------------------------------------------+
	//	+---------------------------------+--------------------------------------------+
	//	| 0x00000000 S_OK                 | The call was successful.                   |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070003 ERROR_PATH_NOT_FOUND | The system cannot find the path specified. |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070057 E_INVALIDARG         | One or more arguments are invalid.         |
	//	+---------------------------------+--------------------------------------------+
	//
	// The opnum field value for this method is 25.
	GetLastChangeTime(context.Context, *GetLastChangeTimeRequest, ...dcerpc.CallOption) (*GetLastChangeTimeResponse, error)

	// The R_KeyExchangePhase1 method receives a pair of encrypted client keys and returns
	// server encryption and session keys.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
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
	// The opnum field value for this method is 26.
	KeyExchangePhase1(context.Context, *KeyExchangePhase1Request, ...dcerpc.CallOption) (*KeyExchangePhase1Response, error)

	// The R_KeyExchangePhase2 method receives the encrypted client session and hash keys
	// in response to the R_KeyExchangePhase1 method and returns the encrypted server hash
	// keys.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
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
	// The opnum field value for this method is 27.
	KeyExchangePhase2(context.Context, *KeyExchangePhase2Request, ...dcerpc.CallOption) (*KeyExchangePhase2Response, error)

	// The Backup method backs up the metabase.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                    | The call was successful.                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG            | One or more arguments are invalid.                                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000CC809 MD_WARNING_SAVE FAILED  | The metadata save prior to backup failed. The previous version of the data was   |
	//	|                                    | backed up.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY | There is not enough memory to complete the operation.                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 28.
	//
	// The location string can be up to 100 Unicode characters in length. Multiple metabase
	// backups can be stored with the same name.
	Backup(context.Context, *BackupRequest, ...dcerpc.CallOption) (*BackupResponse, error)

	// The Restore method restores the metabase from a backup.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+-------------------------------------+-------------------------------------------------------+
	//	|               RETURN                |                                                       |
	//	|             VALUE/CODE              |                      DESCRIPTION                      |
	//	|                                     |                                                       |
	//	+-------------------------------------+-------------------------------------------------------+
	//	+-------------------------------------+-------------------------------------------------------+
	//	| 0x00000000 S_OK                     | The call was successful.                              |
	//	+-------------------------------------+-------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG             | One or more arguments are invalid.                    |
	//	+-------------------------------------+-------------------------------------------------------+
	//	| 0x80070013 ERROR_INVALID_DATA       | The data is invalid.                                  |
	//	+-------------------------------------+-------------------------------------------------------+
	//	| 0x800CC802 MD_ERROR_INVALID_VERSION | The version specified by dwMDVersion is invalid.      |
	//	+-------------------------------------+-------------------------------------------------------+
	//	| 0x000CC805L MD_WARNING_INVALID_DATA | Invalid metabase data.                                |
	//	+-------------------------------------+-------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY  | There is not enough memory to complete the operation. |
	//	+-------------------------------------+-------------------------------------------------------+
	//
	// The opnum field value for this method is 29.
	Restore(context.Context, *RestoreRequest, ...dcerpc.CallOption) (*RestoreResponse, error)

	// The EnumBackups method enumerates metabase backups with a specified backup name or
	// all backups.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+--------------------------------+----------------------------------------+
	//	|             RETURN             |                                        |
	//	|           VALUE/CODE           |              DESCRIPTION               |
	//	|                                |                                        |
	//	+--------------------------------+----------------------------------------+
	//	+--------------------------------+----------------------------------------+
	//	| 0x00000000 S_OK                | The call was successful.               |
	//	+--------------------------------+----------------------------------------+
	//	| 0x80070103 ERROR_NO_MORE_ITEMS | No more data is available.             |
	//	+--------------------------------+----------------------------------------+
	//	| 0x80070057 E_INVALIDARG        | The pszMDBackupName parameter is NULL. |
	//	+--------------------------------+----------------------------------------+
	//
	// The opnum field value for this method is 30.
	EnumBackups(context.Context, *EnumBackupsRequest, ...dcerpc.CallOption) (*EnumBackupsResponse, error)

	// The DeleteBackup method deletes a metabase backup.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+------------------------------------+-------------------------------------------------------+
	//	|               RETURN               |                                                       |
	//	|             VALUE/CODE             |                      DESCRIPTION                      |
	//	|                                    |                                                       |
	//	+------------------------------------+-------------------------------------------------------+
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x00000000 S_OK                    | The call was successful.                              |
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG            | One or more arguments are invalid.                    |
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND    | The system cannot find the file specified.            |
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY | There is not enough memory to complete the operation. |
	//	+------------------------------------+-------------------------------------------------------+
	//
	// The opnum field value for this method is 31.
	DeleteBackup(context.Context, *DeleteBackupRequest, ...dcerpc.CallOption) (*DeleteBackupResponse, error)

	// The UnmarshalInterface method returns a pointer to the IMSAdminBaseW interface.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
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
	// The opnum field value for this method is 32.
	UnmarshalInterface(context.Context, *UnmarshalInterfaceRequest, ...dcerpc.CallOption) (*UnmarshalInterfaceResponse, error)

	// The R_GetServerGuid method returns a GUID for this DCOM object.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
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
	// The opnum field value for this method is 33.
	GetServerGUID(context.Context, *GetServerGUIDRequest, ...dcerpc.CallOption) (*GetServerGUIDResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) IMSAdminBaseWClient
}

type xxx_DefaultIMSAdminBaseWClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultIMSAdminBaseWClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultIMSAdminBaseWClient) AddKey(ctx context.Context, in *AddKeyRequest, opts ...dcerpc.CallOption) (*AddKeyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &AddKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) DeleteKey(ctx context.Context, in *DeleteKeyRequest, opts ...dcerpc.CallOption) (*DeleteKeyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &DeleteKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) DeleteChildKeys(ctx context.Context, in *DeleteChildKeysRequest, opts ...dcerpc.CallOption) (*DeleteChildKeysResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &DeleteChildKeysResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) EnumKeys(ctx context.Context, in *EnumKeysRequest, opts ...dcerpc.CallOption) (*EnumKeysResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &EnumKeysResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) CopyKey(ctx context.Context, in *CopyKeyRequest, opts ...dcerpc.CallOption) (*CopyKeyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &CopyKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) RenameKey(ctx context.Context, in *RenameKeyRequest, opts ...dcerpc.CallOption) (*RenameKeyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &RenameKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) SetData(ctx context.Context, in *SetDataRequest, opts ...dcerpc.CallOption) (*SetDataResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &SetDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) GetData(ctx context.Context, in *GetDataRequest, opts ...dcerpc.CallOption) (*GetDataResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) DeleteData(ctx context.Context, in *DeleteDataRequest, opts ...dcerpc.CallOption) (*DeleteDataResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &DeleteDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) EnumData(ctx context.Context, in *EnumDataRequest, opts ...dcerpc.CallOption) (*EnumDataResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &EnumDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) GetAllData(ctx context.Context, in *GetAllDataRequest, opts ...dcerpc.CallOption) (*GetAllDataResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetAllDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) DeleteAllData(ctx context.Context, in *DeleteAllDataRequest, opts ...dcerpc.CallOption) (*DeleteAllDataResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &DeleteAllDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) CopyData(ctx context.Context, in *CopyDataRequest, opts ...dcerpc.CallOption) (*CopyDataResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &CopyDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) GetDataPaths(ctx context.Context, in *GetDataPathsRequest, opts ...dcerpc.CallOption) (*GetDataPathsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetDataPathsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) OpenKey(ctx context.Context, in *OpenKeyRequest, opts ...dcerpc.CallOption) (*OpenKeyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &OpenKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) CloseKey(ctx context.Context, in *CloseKeyRequest, opts ...dcerpc.CallOption) (*CloseKeyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &CloseKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) ChangePermissions(ctx context.Context, in *ChangePermissionsRequest, opts ...dcerpc.CallOption) (*ChangePermissionsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &ChangePermissionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) SaveData(ctx context.Context, in *SaveDataRequest, opts ...dcerpc.CallOption) (*SaveDataResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &SaveDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) GetHandleInfo(ctx context.Context, in *GetHandleInfoRequest, opts ...dcerpc.CallOption) (*GetHandleInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetHandleInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) GetSystemChangeNumber(ctx context.Context, in *GetSystemChangeNumberRequest, opts ...dcerpc.CallOption) (*GetSystemChangeNumberResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetSystemChangeNumberResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) GetDataSetNumber(ctx context.Context, in *GetDataSetNumberRequest, opts ...dcerpc.CallOption) (*GetDataSetNumberResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetDataSetNumberResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) SetLastChangeTime(ctx context.Context, in *SetLastChangeTimeRequest, opts ...dcerpc.CallOption) (*SetLastChangeTimeResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &SetLastChangeTimeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) GetLastChangeTime(ctx context.Context, in *GetLastChangeTimeRequest, opts ...dcerpc.CallOption) (*GetLastChangeTimeResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetLastChangeTimeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) KeyExchangePhase1(ctx context.Context, in *KeyExchangePhase1Request, opts ...dcerpc.CallOption) (*KeyExchangePhase1Response, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &KeyExchangePhase1Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) KeyExchangePhase2(ctx context.Context, in *KeyExchangePhase2Request, opts ...dcerpc.CallOption) (*KeyExchangePhase2Response, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &KeyExchangePhase2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) Backup(ctx context.Context, in *BackupRequest, opts ...dcerpc.CallOption) (*BackupResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &BackupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) Restore(ctx context.Context, in *RestoreRequest, opts ...dcerpc.CallOption) (*RestoreResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &RestoreResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) EnumBackups(ctx context.Context, in *EnumBackupsRequest, opts ...dcerpc.CallOption) (*EnumBackupsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &EnumBackupsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) DeleteBackup(ctx context.Context, in *DeleteBackupRequest, opts ...dcerpc.CallOption) (*DeleteBackupResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &DeleteBackupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) UnmarshalInterface(ctx context.Context, in *UnmarshalInterfaceRequest, opts ...dcerpc.CallOption) (*UnmarshalInterfaceResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &UnmarshalInterfaceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) GetServerGUID(ctx context.Context, in *GetServerGUIDRequest, opts ...dcerpc.CallOption) (*GetServerGUIDResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetServerGUIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIMSAdminBaseWClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultIMSAdminBaseWClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultIMSAdminBaseWClient) IPID(ctx context.Context, ipid *dcom.IPID) IMSAdminBaseWClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultIMSAdminBaseWClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewIMSAdminBaseWClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (IMSAdminBaseWClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(IMSAdminBaseWSyntaxV0_0))...)
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
	return &xxx_DefaultIMSAdminBaseWClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_AddKeyOperation structure represents the AddKey operation
type xxx_AddKeyOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Handle uint32         `idl:"name:hMDHandle" json:"handle"`
	Path   string         `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_AddKeyOperation) OpNum() int { return 3 }

func (o *xxx_AddKeyOperation) OpName() string { return "/IMSAdminBaseW/v0/AddKey" }

func (o *xxx_AddKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.WriteData(o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_pszMDPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_pszMDPath); err != nil {
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

func (o *xxx_AddKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.ReadData(&o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszMDPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_pszMDPath := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_pszMDPath, _ptr_pszMDPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AddKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// AddKeyRequest structure represents the AddKey operation request
type AddKeyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// hMDHandle: An unsigned 32-bit integer value containing an open metabase handle specifying
	// the node in the metabase where the new key is to be added.
	Handle uint32 `idl:"name:hMDHandle" json:"handle"`
	// pszMDPath: A pointer to a Unicode string that contains the new node's path, relative
	// to the path of the hMDHandle parameter.
	Path string `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
}

func (o *AddKeyRequest) xxx_ToOp(ctx context.Context, op *xxx_AddKeyOperation) *xxx_AddKeyOperation {
	if op == nil {
		op = &xxx_AddKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Handle = o.Handle
	op.Path = o.Path
	return op
}

func (o *AddKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_AddKeyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Handle = op.Handle
	o.Path = op.Path
}
func (o *AddKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AddKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddKeyResponse structure represents the AddKey operation response
type AddKeyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AddKey return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AddKeyResponse) xxx_ToOp(ctx context.Context, op *xxx_AddKeyOperation) *xxx_AddKeyOperation {
	if op == nil {
		op = &xxx_AddKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *AddKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_AddKeyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *AddKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AddKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteKeyOperation structure represents the DeleteKey operation
type xxx_DeleteKeyOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Handle uint32         `idl:"name:hMDHandle" json:"handle"`
	Path   string         `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteKeyOperation) OpNum() int { return 4 }

func (o *xxx_DeleteKeyOperation) OpName() string { return "/IMSAdminBaseW/v0/DeleteKey" }

func (o *xxx_DeleteKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.WriteData(o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_pszMDPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_pszMDPath); err != nil {
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

func (o *xxx_DeleteKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.ReadData(&o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszMDPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_pszMDPath := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_pszMDPath, _ptr_pszMDPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteKeyRequest structure represents the DeleteKey operation request
type DeleteKeyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// hMDHandle: An unsigned 32-bit integer value containing an open metabase handle specifying
	// a node in the metabase where the key is to be deleted.
	Handle uint32 `idl:"name:hMDHandle" json:"handle"`
	// pszMDPath: A pointer to a Unicode string that contains the path of the node to be
	// deleted, relative to the path of the hMDHandle parameter. This parameter MUST NOT
	// be NULL.
	Path string `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
}

func (o *DeleteKeyRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteKeyOperation) *xxx_DeleteKeyOperation {
	if op == nil {
		op = &xxx_DeleteKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Handle = o.Handle
	op.Path = o.Path
	return op
}

func (o *DeleteKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteKeyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Handle = op.Handle
	o.Path = op.Path
}
func (o *DeleteKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteKeyResponse structure represents the DeleteKey operation response
type DeleteKeyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteKey return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteKeyResponse) xxx_ToOp(ctx context.Context, op *xxx_DeleteKeyOperation) *xxx_DeleteKeyOperation {
	if op == nil {
		op = &xxx_DeleteKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DeleteKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteKeyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteChildKeysOperation structure represents the DeleteChildKeys operation
type xxx_DeleteChildKeysOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Handle uint32         `idl:"name:hMDHandle" json:"handle"`
	Path   string         `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteChildKeysOperation) OpNum() int { return 5 }

func (o *xxx_DeleteChildKeysOperation) OpName() string { return "/IMSAdminBaseW/v0/DeleteChildKeys" }

func (o *xxx_DeleteChildKeysOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteChildKeysOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.WriteData(o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_pszMDPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_pszMDPath); err != nil {
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

func (o *xxx_DeleteChildKeysOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.ReadData(&o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszMDPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_pszMDPath := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_pszMDPath, _ptr_pszMDPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteChildKeysOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteChildKeysOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteChildKeysOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteChildKeysRequest structure represents the DeleteChildKeys operation request
type DeleteChildKeysRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// hMDHandle: An unsigned 32-bit integer value containing an open metabase handle specifying
	// the node in the metabase where the child key is to be deleted.
	Handle uint32 `idl:"name:hMDHandle" json:"handle"`
	// pszMDPath: A pointer to a Unicode string that contains the path of the node whose
	// subnodes are to be deleted, relative to the path of the hMDHandle parameter.
	Path string `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
}

func (o *DeleteChildKeysRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteChildKeysOperation) *xxx_DeleteChildKeysOperation {
	if op == nil {
		op = &xxx_DeleteChildKeysOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Handle = o.Handle
	op.Path = o.Path
	return op
}

func (o *DeleteChildKeysRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteChildKeysOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Handle = op.Handle
	o.Path = op.Path
}
func (o *DeleteChildKeysRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteChildKeysRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteChildKeysOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteChildKeysResponse structure represents the DeleteChildKeys operation response
type DeleteChildKeysResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteChildKeys return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteChildKeysResponse) xxx_ToOp(ctx context.Context, op *xxx_DeleteChildKeysOperation) *xxx_DeleteChildKeysOperation {
	if op == nil {
		op = &xxx_DeleteChildKeysOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DeleteChildKeysResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteChildKeysOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteChildKeysResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteChildKeysResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteChildKeysOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumKeysOperation structure represents the EnumKeys operation
type xxx_EnumKeysOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	Handle          uint32         `idl:"name:hMDHandle" json:"handle"`
	Path            string         `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	Name            string         `idl:"name:pszMDName;size_is:(256)" json:"name"`
	EnumObjectIndex uint32         `idl:"name:dwMDEnumObjectIndex" json:"enum_object_index"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumKeysOperation) OpNum() int { return 6 }

func (o *xxx_EnumKeysOperation) OpName() string { return "/IMSAdminBaseW/v0/EnumKeys" }

func (o *xxx_EnumKeysOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumKeysOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.WriteData(o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_pszMDPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_pszMDPath); err != nil {
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
	// dwMDEnumObjectIndex {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.EnumObjectIndex); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumKeysOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.ReadData(&o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszMDPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_pszMDPath := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_pszMDPath, _ptr_pszMDPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwMDEnumObjectIndex {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.EnumObjectIndex); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumKeysOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	// cannot evaluate expression 256
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumKeysOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pszMDName {out} (1:{alias=LPWSTR}*(1)[dim:0,size_is=256,string](wchar))
	{
		dimSize1 := uint64(256)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		_Name_buf := utf16.Encode([]rune(o.Name))
		if uint64(len(_Name_buf)) > sizeInfo[0] {
			_Name_buf = _Name_buf[:sizeInfo[0]]
		}
		for i1 := range _Name_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_Name_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_Name_buf); uint64(i1) < sizeInfo[0]; i1++ {
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

func (o *xxx_EnumKeysOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pszMDName {out} (1:{alias=LPWSTR,pointer=ref}*(1)[dim:0,size_is=256,string](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _Name_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Name_buf", sizeInfo[0])
		}
		_Name_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Name_buf {
			i1 := i1
			if err := w.ReadData(&_Name_buf[i1]); err != nil {
				return err
			}
		}
		o.Name = strings.TrimRight(string(utf16.Decode(_Name_buf)), ndr.ZeroString)
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumKeysRequest structure represents the EnumKeys operation request
type EnumKeysRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// hMDHandle: An unsigned 32-bit integer value containing an open metabase handle specifying
	// the key to be enumerated.
	Handle uint32 `idl:"name:hMDHandle" json:"handle"`
	// pszMDPath: A pointer to a Unicode string that contains the path of the node whose
	// subnodes are to be enumerated, relative to the path of the hMDHandle parameter.
	Path string `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	// dwMDEnumObjectIndex: An integer value specifying the index of the subnode to be retrieved.
	EnumObjectIndex uint32 `idl:"name:dwMDEnumObjectIndex" json:"enum_object_index"`
}

func (o *EnumKeysRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumKeysOperation) *xxx_EnumKeysOperation {
	if op == nil {
		op = &xxx_EnumKeysOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Handle = o.Handle
	op.Path = o.Path
	op.EnumObjectIndex = o.EnumObjectIndex
	return op
}

func (o *EnumKeysRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumKeysOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Handle = op.Handle
	o.Path = op.Path
	o.EnumObjectIndex = op.EnumObjectIndex
}
func (o *EnumKeysRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumKeysRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumKeysOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumKeysResponse structure represents the EnumKeys operation response
type EnumKeysResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pszMDName: A pointer to a string buffer that receives the names of the enumerated
	// metabase subnodes.
	Name string `idl:"name:pszMDName;size_is:(256)" json:"name"`
	// Return: The EnumKeys return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumKeysResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumKeysOperation) *xxx_EnumKeysOperation {
	if op == nil {
		op = &xxx_EnumKeysOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Name = o.Name
	op.Return = o.Return
	return op
}

func (o *EnumKeysResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumKeysOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Name = op.Name
	o.Return = op.Return
}
func (o *EnumKeysResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumKeysResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumKeysOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CopyKeyOperation structure represents the CopyKey operation
type xxx_CopyKeyOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	SourceHandle      uint32         `idl:"name:hMDSourceHandle" json:"source_handle"`
	SourcePath        string         `idl:"name:pszMDSourcePath;string;pointer:unique" json:"source_path"`
	DestinationHandle uint32         `idl:"name:hMDDestHandle" json:"destination_handle"`
	DestinationPath   string         `idl:"name:pszMDDestPath;string;pointer:unique" json:"destination_path"`
	OverwriteFlag     bool           `idl:"name:bMDOverwriteFlag" json:"overwrite_flag"`
	CopyFlag          bool           `idl:"name:bMDCopyFlag" json:"copy_flag"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CopyKeyOperation) OpNum() int { return 7 }

func (o *xxx_CopyKeyOperation) OpName() string { return "/IMSAdminBaseW/v0/CopyKey" }

func (o *xxx_CopyKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CopyKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// hMDSourceHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.WriteData(o.SourceHandle); err != nil {
			return err
		}
	}
	// pszMDSourcePath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.SourcePath != "" {
			_ptr_pszMDSourcePath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.SourcePath); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.SourcePath, _ptr_pszMDSourcePath); err != nil {
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
	// hMDDestHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.WriteData(o.DestinationHandle); err != nil {
			return err
		}
	}
	// pszMDDestPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.DestinationPath != "" {
			_ptr_pszMDDestPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.DestinationPath); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.DestinationPath, _ptr_pszMDDestPath); err != nil {
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
	// bMDOverwriteFlag {in} (1:{alias=BOOL}(int32))
	{
		if !o.OverwriteFlag {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	// bMDCopyFlag {in} (1:{alias=BOOL}(int32))
	{
		if !o.CopyFlag {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CopyKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// hMDSourceHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.ReadData(&o.SourceHandle); err != nil {
			return err
		}
	}
	// pszMDSourcePath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszMDSourcePath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.SourcePath); err != nil {
				return err
			}
			return nil
		})
		_s_pszMDSourcePath := func(ptr interface{}) { o.SourcePath = *ptr.(*string) }
		if err := w.ReadPointer(&o.SourcePath, _s_pszMDSourcePath, _ptr_pszMDSourcePath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// hMDDestHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.ReadData(&o.DestinationHandle); err != nil {
			return err
		}
	}
	// pszMDDestPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszMDDestPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.DestinationPath); err != nil {
				return err
			}
			return nil
		})
		_s_pszMDDestPath := func(ptr interface{}) { o.DestinationPath = *ptr.(*string) }
		if err := w.ReadPointer(&o.DestinationPath, _s_pszMDDestPath, _ptr_pszMDDestPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bMDOverwriteFlag {in} (1:{alias=BOOL}(int32))
	{
		var _bOverwriteFlag int32
		if err := w.ReadData(&_bOverwriteFlag); err != nil {
			return err
		}
		o.OverwriteFlag = _bOverwriteFlag != 0
	}
	// bMDCopyFlag {in} (1:{alias=BOOL}(int32))
	{
		var _bCopyFlag int32
		if err := w.ReadData(&_bCopyFlag); err != nil {
			return err
		}
		o.CopyFlag = _bCopyFlag != 0
	}
	return nil
}

func (o *xxx_CopyKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CopyKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CopyKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CopyKeyRequest structure represents the CopyKey operation request
type CopyKeyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// hMDSourceHandle: An unsigned 32-bit integer value containing an open metabase handle
	// specifying the source node to be copied or moved.
	SourceHandle uint32 `idl:"name:hMDSourceHandle" json:"source_handle"`
	// pszMDSourcePath: A pointer to a Unicode string that contains the path of the node
	// to be copied or moved relative to the path of the hMDSourceHandle parameter.
	SourcePath string `idl:"name:pszMDSourcePath;string;pointer:unique" json:"source_path"`
	// hMDDestHandle: An unsigned 32-bit integer value containing an open metabase handle
	// specifying the destination node of the moved or copied metabase key.
	DestinationHandle uint32 `idl:"name:hMDDestHandle" json:"destination_handle"`
	// pszMDDestPath: A pointer to a string that contains the path of the new or moved node,
	// relative to the hMDDestHandle parameter.
	DestinationPath string `idl:"name:pszMDDestPath;string;pointer:unique" json:"destination_path"`
	// bMDOverwriteFlag: A Boolean value that determine the behavior if a node with the
	// same name as source is already a child of destination node. If TRUE, the existing
	// node and all its data and children are deleted prior to copying or moving the source.
	// If FALSE, the existing node, data, and children remain, and the source is merged
	// with that data. In cases of data conflicts, the source data overwrites the destination
	// data.
	OverwriteFlag bool `idl:"name:bMDOverwriteFlag" json:"overwrite_flag"`
	// bMDCopyFlag: A Boolean value that specifies whether to copy or move the specified
	// node. If TRUE, the node is copied. If FALSE, the node is moved, and the source node
	// is deleted from its original location.
	CopyFlag bool `idl:"name:bMDCopyFlag" json:"copy_flag"`
}

func (o *CopyKeyRequest) xxx_ToOp(ctx context.Context, op *xxx_CopyKeyOperation) *xxx_CopyKeyOperation {
	if op == nil {
		op = &xxx_CopyKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.SourceHandle = o.SourceHandle
	op.SourcePath = o.SourcePath
	op.DestinationHandle = o.DestinationHandle
	op.DestinationPath = o.DestinationPath
	op.OverwriteFlag = o.OverwriteFlag
	op.CopyFlag = o.CopyFlag
	return op
}

func (o *CopyKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_CopyKeyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SourceHandle = op.SourceHandle
	o.SourcePath = op.SourcePath
	o.DestinationHandle = op.DestinationHandle
	o.DestinationPath = op.DestinationPath
	o.OverwriteFlag = op.OverwriteFlag
	o.CopyFlag = op.CopyFlag
}
func (o *CopyKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CopyKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CopyKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CopyKeyResponse structure represents the CopyKey operation response
type CopyKeyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CopyKey return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CopyKeyResponse) xxx_ToOp(ctx context.Context, op *xxx_CopyKeyOperation) *xxx_CopyKeyOperation {
	if op == nil {
		op = &xxx_CopyKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CopyKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_CopyKeyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CopyKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CopyKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CopyKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RenameKeyOperation structure represents the RenameKey operation
type xxx_RenameKeyOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Handle  uint32         `idl:"name:hMDHandle" json:"handle"`
	Path    string         `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	NewName string         `idl:"name:pszMDNewName;string;pointer:unique" json:"new_name"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RenameKeyOperation) OpNum() int { return 8 }

func (o *xxx_RenameKeyOperation) OpName() string { return "/IMSAdminBaseW/v0/RenameKey" }

func (o *xxx_RenameKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.WriteData(o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_pszMDPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_pszMDPath); err != nil {
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
	// pszMDNewName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.NewName != "" {
			_ptr_pszMDNewName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.NewName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.NewName, _ptr_pszMDNewName); err != nil {
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

func (o *xxx_RenameKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.ReadData(&o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszMDPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_pszMDPath := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_pszMDPath, _ptr_pszMDPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszMDNewName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszMDNewName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.NewName); err != nil {
				return err
			}
			return nil
		})
		_s_pszMDNewName := func(ptr interface{}) { o.NewName = *ptr.(*string) }
		if err := w.ReadPointer(&o.NewName, _s_pszMDNewName, _ptr_pszMDNewName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RenameKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RenameKeyRequest structure represents the RenameKey operation request
type RenameKeyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// hMDHandle: An unsigned 32-bit integer value containing an open metabase handle specifying
	// the key to be renamed.
	Handle uint32 `idl:"name:hMDHandle" json:"handle"`
	// pszMDPath: A pointer to a Unicode string that contains the path of the node to be
	// renamed, relative to the path of the hMDHandle parameter.
	Path string `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	// pszMDNewName: A pointer to a string that contains the new name for the node.
	NewName string `idl:"name:pszMDNewName;string;pointer:unique" json:"new_name"`
}

func (o *RenameKeyRequest) xxx_ToOp(ctx context.Context, op *xxx_RenameKeyOperation) *xxx_RenameKeyOperation {
	if op == nil {
		op = &xxx_RenameKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Handle = o.Handle
	op.Path = o.Path
	op.NewName = o.NewName
	return op
}

func (o *RenameKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_RenameKeyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Handle = op.Handle
	o.Path = op.Path
	o.NewName = op.NewName
}
func (o *RenameKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RenameKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RenameKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RenameKeyResponse structure represents the RenameKey operation response
type RenameKeyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RenameKey return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RenameKeyResponse) xxx_ToOp(ctx context.Context, op *xxx_RenameKeyOperation) *xxx_RenameKeyOperation {
	if op == nil {
		op = &xxx_RenameKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RenameKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_RenameKeyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RenameKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RenameKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RenameKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetDataOperation structure represents the R_SetData operation
type xxx_SetDataOperation struct {
	This   *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat       `idl:"name:That" json:"that"`
	Handle uint32               `idl:"name:hMDHandle" json:"handle"`
	Path   string               `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	Data   *imsa.MetadataRecord `idl:"name:pmdrMDData" json:"data"`
	Return int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_SetDataOperation) OpNum() int { return 9 }

func (o *xxx_SetDataOperation) OpName() string { return "/IMSAdminBaseW/v0/R_SetData" }

func (o *xxx_SetDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.WriteData(o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_pszMDPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_pszMDPath); err != nil {
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
	// pmdrMDData {in} (1:{pointer=ref}*(1))(2:{alias=METADATA_RECORD}(struct))
	{
		if o.Data != nil {
			if err := o.Data.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&imsa.MetadataRecord{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.ReadData(&o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszMDPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_pszMDPath := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_pszMDPath, _ptr_pszMDPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pmdrMDData {in} (1:{pointer=ref}*(1))(2:{alias=METADATA_RECORD}(struct))
	{
		if o.Data == nil {
			o.Data = &imsa.MetadataRecord{}
		}
		if err := o.Data.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetDataRequest structure represents the R_SetData operation request
type SetDataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// hMDHandle: An unsigned 32-bit integer value specifying a handle to a node in the
	// metabase with write permissions as returned by the OpenKey method.
	Handle uint32 `idl:"name:hMDHandle" json:"handle"`
	// pszMDPath: A pointer to a Unicode string that contains the path of the node that
	// stores the entry, relative to the path of the hMDHandle parameter.
	Path string `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	// pmdrMDData: A pointer to a METADATA_RECORD structure that contains the data to set.
	Data *imsa.MetadataRecord `idl:"name:pmdrMDData" json:"data"`
}

func (o *SetDataRequest) xxx_ToOp(ctx context.Context, op *xxx_SetDataOperation) *xxx_SetDataOperation {
	if op == nil {
		op = &xxx_SetDataOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Handle = o.Handle
	op.Path = o.Path
	op.Data = o.Data
	return op
}

func (o *SetDataRequest) xxx_FromOp(ctx context.Context, op *xxx_SetDataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Handle = op.Handle
	o.Path = op.Path
	o.Data = op.Data
}
func (o *SetDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetDataResponse structure represents the R_SetData operation response
type SetDataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The R_SetData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetDataResponse) xxx_ToOp(ctx context.Context, op *xxx_SetDataOperation) *xxx_SetDataOperation {
	if op == nil {
		op = &xxx_SetDataOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetDataResponse) xxx_FromOp(ctx context.Context, op *xxx_SetDataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDataOperation structure represents the R_GetData operation
type xxx_GetDataOperation struct {
	This               *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat       `idl:"name:That" json:"that"`
	Handle             uint32               `idl:"name:hMDHandle" json:"handle"`
	Path               string               `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	Data               *imsa.MetadataRecord `idl:"name:pmdrMDData" json:"data"`
	RequiredDataLength uint32               `idl:"name:pdwMDRequiredDataLen" json:"required_data_length"`
	DataBlob           *imsa.IISCryptoBlob  `idl:"name:ppDataBlob" json:"data_blob"`
	Return             int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDataOperation) OpNum() int { return 10 }

func (o *xxx_GetDataOperation) OpName() string { return "/IMSAdminBaseW/v0/R_GetData" }

func (o *xxx_GetDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.WriteData(o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_pszMDPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_pszMDPath); err != nil {
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
	// pmdrMDData {in, out} (1:{pointer=ref}*(1))(2:{alias=METADATA_RECORD}(struct))
	{
		if o.Data != nil {
			if err := o.Data.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&imsa.MetadataRecord{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.ReadData(&o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszMDPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_pszMDPath := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_pszMDPath, _ptr_pszMDPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pmdrMDData {in, out} (1:{pointer=ref}*(1))(2:{alias=METADATA_RECORD}(struct))
	{
		if o.Data == nil {
			o.Data = &imsa.MetadataRecord{}
		}
		if err := o.Data.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pmdrMDData {in, out} (1:{pointer=ref}*(1))(2:{alias=METADATA_RECORD}(struct))
	{
		if o.Data != nil {
			if err := o.Data.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&imsa.MetadataRecord{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pdwMDRequiredDataLen {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequiredDataLength); err != nil {
			return err
		}
	}
	// ppDataBlob {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IIS_CRYPTO_BLOB}(struct))
	{
		if o.DataBlob != nil {
			_ptr_ppDataBlob := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DataBlob != nil {
					if err := o.DataBlob.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&imsa.IISCryptoBlob{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DataBlob, _ptr_ppDataBlob); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pmdrMDData {in, out} (1:{pointer=ref}*(1))(2:{alias=METADATA_RECORD}(struct))
	{
		if o.Data == nil {
			o.Data = &imsa.MetadataRecord{}
		}
		if err := o.Data.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pdwMDRequiredDataLen {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequiredDataLength); err != nil {
			return err
		}
	}
	// ppDataBlob {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IIS_CRYPTO_BLOB}(struct))
	{
		_ptr_ppDataBlob := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DataBlob == nil {
				o.DataBlob = &imsa.IISCryptoBlob{}
			}
			if err := o.DataBlob.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppDataBlob := func(ptr interface{}) { o.DataBlob = *ptr.(**imsa.IISCryptoBlob) }
		if err := w.ReadPointer(&o.DataBlob, _s_ppDataBlob, _ptr_ppDataBlob); err != nil {
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

// GetDataRequest structure represents the R_GetData operation request
type GetDataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// hMDHandle: An unsigned 32-bit integer value containing an open metabase handle specifying
	// the key to be queried.
	Handle uint32 `idl:"name:hMDHandle" json:"handle"`
	// pszMDPath: A pointer to a Unicode string that contains the path of the node containing
	// the data, relative to the path of the hMDHandle parameter.
	Path string `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	// pmdrMDData: A pointer to a METADATA_RECORD structure that describes the requested
	// data.
	Data *imsa.MetadataRecord `idl:"name:pmdrMDData" json:"data"`
}

func (o *GetDataRequest) xxx_ToOp(ctx context.Context, op *xxx_GetDataOperation) *xxx_GetDataOperation {
	if op == nil {
		op = &xxx_GetDataOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Handle = o.Handle
	op.Path = o.Path
	op.Data = o.Data
	return op
}

func (o *GetDataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Handle = op.Handle
	o.Path = op.Path
	o.Data = op.Data
}
func (o *GetDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDataResponse structure represents the R_GetData operation response
type GetDataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pmdrMDData: A pointer to a METADATA_RECORD structure that describes the requested
	// data.
	Data *imsa.MetadataRecord `idl:"name:pmdrMDData" json:"data"`
	// pdwMDRequiredDataLen: A pointer to an integer value that contains the buffer length
	// required, in bytes, to contain the decrypted data referenced by the ppDataBlob parameter.
	RequiredDataLength uint32 `idl:"name:pdwMDRequiredDataLen" json:"required_data_length"`
	// ppDataBlob: An IIS_CRYPTO_BLOB structure containing the requested data value as encrypted
	// opaque data.
	DataBlob *imsa.IISCryptoBlob `idl:"name:ppDataBlob" json:"data_blob"`
	// Return: The R_GetData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDataResponse) xxx_ToOp(ctx context.Context, op *xxx_GetDataOperation) *xxx_GetDataOperation {
	if op == nil {
		op = &xxx_GetDataOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Data = o.Data
	op.RequiredDataLength = o.RequiredDataLength
	op.DataBlob = o.DataBlob
	op.Return = o.Return
	return op
}

func (o *GetDataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Data = op.Data
	o.RequiredDataLength = op.RequiredDataLength
	o.DataBlob = op.DataBlob
	o.Return = op.Return
}
func (o *GetDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteDataOperation structure represents the DeleteData operation
type xxx_DeleteDataOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Handle   uint32         `idl:"name:hMDHandle" json:"handle"`
	Path     string         `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	ID       uint32         `idl:"name:dwMDIdentifier" json:"id"`
	DataType uint32         `idl:"name:dwMDDataType" json:"data_type"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteDataOperation) OpNum() int { return 11 }

func (o *xxx_DeleteDataOperation) OpName() string { return "/IMSAdminBaseW/v0/DeleteData" }

func (o *xxx_DeleteDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.WriteData(o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_pszMDPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_pszMDPath); err != nil {
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
	// dwMDIdentifier {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ID); err != nil {
			return err
		}
	}
	// dwMDDataType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.DataType); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.ReadData(&o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszMDPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_pszMDPath := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_pszMDPath, _ptr_pszMDPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwMDIdentifier {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ID); err != nil {
			return err
		}
	}
	// dwMDDataType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DataType); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteDataRequest structure represents the DeleteData operation request
type DeleteDataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// hMDHandle: An unsigned 32-bit integer value containing an open metabase handle specifying
	// the node in the metabase where the key data is to be deleted.
	Handle uint32 `idl:"name:hMDHandle" json:"handle"`
	// pszMDPath: A pointer to a Unicode string that contains the path of the node whose
	// data is to be deleted, relative to the path of the hMDHandle parameter.
	Path string `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	// dwMDIdentifier: An integer value specifying the data identifier.
	ID uint32 `idl:"name:dwMDIdentifier" json:"id"`
	// dwMDDataType: An integer value specifying a data type. If this parameter is not set
	// to ALL_METADATA, the data item will be removed only if its data type matches the
	// specified type.
	//
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	|                              |                                                                                  |
	//	|            VALUE             |                                     MEANING                                      |
	//	|                              |                                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| ALL_METADATA 0x00000000      | Specifies all data, regardless of type.                                          |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| BINARY_METADATA 0x00000003   | Specifies binary data in any form.                                               |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| DWORD_METADATA 0x00000001    | Specifies all DWORD (unsigned 32-bit integer) data.                              |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| EXPANDSZ_METADATA 0x00000004 | Specifies all data consisting of a string that includes the terminating null     |
	//	|                              | character, which contains unexpanded environment variables.                      |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| MULTISZ_METADATA 0x00000005  | Specifies all data represented as an array of strings, where each string         |
	//	|                              | contains two occurrences of the terminating null character.                      |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| STRING_METADATA 0x00000002   | Specifies all data consisting of an ASCII string that includes the terminating   |
	//	|                              | null character.                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	DataType uint32 `idl:"name:dwMDDataType" json:"data_type"`
}

func (o *DeleteDataRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteDataOperation) *xxx_DeleteDataOperation {
	if op == nil {
		op = &xxx_DeleteDataOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Handle = o.Handle
	op.Path = o.Path
	op.ID = o.ID
	op.DataType = o.DataType
	return op
}

func (o *DeleteDataRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteDataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Handle = op.Handle
	o.Path = op.Path
	o.ID = op.ID
	o.DataType = op.DataType
}
func (o *DeleteDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteDataResponse structure represents the DeleteData operation response
type DeleteDataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteDataResponse) xxx_ToOp(ctx context.Context, op *xxx_DeleteDataOperation) *xxx_DeleteDataOperation {
	if op == nil {
		op = &xxx_DeleteDataOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DeleteDataResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteDataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumDataOperation structure represents the R_EnumData operation
type xxx_EnumDataOperation struct {
	This               *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat       `idl:"name:That" json:"that"`
	Handle             uint32               `idl:"name:hMDHandle" json:"handle"`
	Path               string               `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	Data               *imsa.MetadataRecord `idl:"name:pmdrMDData" json:"data"`
	EnumDataIndex      uint32               `idl:"name:dwMDEnumDataIndex" json:"enum_data_index"`
	RequiredDataLength uint32               `idl:"name:pdwMDRequiredDataLen" json:"required_data_length"`
	DataBlob           *imsa.IISCryptoBlob  `idl:"name:ppDataBlob" json:"data_blob"`
	Return             int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumDataOperation) OpNum() int { return 12 }

func (o *xxx_EnumDataOperation) OpName() string { return "/IMSAdminBaseW/v0/R_EnumData" }

func (o *xxx_EnumDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.WriteData(o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_pszMDPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_pszMDPath); err != nil {
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
	// pmdrMDData {in, out} (1:{pointer=ref}*(1))(2:{alias=METADATA_RECORD}(struct))
	{
		if o.Data != nil {
			if err := o.Data.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&imsa.MetadataRecord{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwMDEnumDataIndex {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.EnumDataIndex); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.ReadData(&o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszMDPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_pszMDPath := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_pszMDPath, _ptr_pszMDPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pmdrMDData {in, out} (1:{pointer=ref}*(1))(2:{alias=METADATA_RECORD}(struct))
	{
		if o.Data == nil {
			o.Data = &imsa.MetadataRecord{}
		}
		if err := o.Data.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwMDEnumDataIndex {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.EnumDataIndex); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pmdrMDData {in, out} (1:{pointer=ref}*(1))(2:{alias=METADATA_RECORD}(struct))
	{
		if o.Data != nil {
			if err := o.Data.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&imsa.MetadataRecord{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pdwMDRequiredDataLen {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequiredDataLength); err != nil {
			return err
		}
	}
	// ppDataBlob {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IIS_CRYPTO_BLOB}(struct))
	{
		if o.DataBlob != nil {
			_ptr_ppDataBlob := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DataBlob != nil {
					if err := o.DataBlob.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&imsa.IISCryptoBlob{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DataBlob, _ptr_ppDataBlob); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pmdrMDData {in, out} (1:{pointer=ref}*(1))(2:{alias=METADATA_RECORD}(struct))
	{
		if o.Data == nil {
			o.Data = &imsa.MetadataRecord{}
		}
		if err := o.Data.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pdwMDRequiredDataLen {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequiredDataLength); err != nil {
			return err
		}
	}
	// ppDataBlob {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IIS_CRYPTO_BLOB}(struct))
	{
		_ptr_ppDataBlob := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DataBlob == nil {
				o.DataBlob = &imsa.IISCryptoBlob{}
			}
			if err := o.DataBlob.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppDataBlob := func(ptr interface{}) { o.DataBlob = *ptr.(**imsa.IISCryptoBlob) }
		if err := w.ReadPointer(&o.DataBlob, _s_ppDataBlob, _ptr_ppDataBlob); err != nil {
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

// EnumDataRequest structure represents the R_EnumData operation request
type EnumDataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// hMDHandle: An unsigned 32-bit integer value containing an open metabase handle specifying
	// the key to be enumerated.
	Handle uint32 `idl:"name:hMDHandle" json:"handle"`
	// pszMDPath: A pointer to a Unicode string that contains the path of the node to be
	// enumerated, relative to the path of the hMDHandle parameter.
	Path string `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	// pmdrMDData: A pointer to a METADATA_RECORD structure that specifies the retrieved
	// data.
	Data *imsa.MetadataRecord `idl:"name:pmdrMDData" json:"data"`
	// dwMDEnumDataIndex: An integer value specifying the index of the entry retrieved.
	EnumDataIndex uint32 `idl:"name:dwMDEnumDataIndex" json:"enum_data_index"`
}

func (o *EnumDataRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumDataOperation) *xxx_EnumDataOperation {
	if op == nil {
		op = &xxx_EnumDataOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Handle = o.Handle
	op.Path = o.Path
	op.Data = o.Data
	op.EnumDataIndex = o.EnumDataIndex
	return op
}

func (o *EnumDataRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumDataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Handle = op.Handle
	o.Path = op.Path
	o.Data = op.Data
	o.EnumDataIndex = op.EnumDataIndex
}
func (o *EnumDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumDataResponse structure represents the R_EnumData operation response
type EnumDataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pmdrMDData: A pointer to a METADATA_RECORD structure that specifies the retrieved
	// data.
	Data *imsa.MetadataRecord `idl:"name:pmdrMDData" json:"data"`
	// pdwMDRequiredDataLen: Pointer to a DWORD that receives the required buffer size if
	// the method returns ERROR_INSUFFICIENT_BUFFER as specified in [MS-ERREF].
	RequiredDataLength uint32 `idl:"name:pdwMDRequiredDataLen" json:"required_data_length"`
	// ppDataBlob: An IIS_CRYPTO_BLOB structure containing the data value as encrypted opaque
	// data.
	DataBlob *imsa.IISCryptoBlob `idl:"name:ppDataBlob" json:"data_blob"`
	// Return: The R_EnumData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumDataResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumDataOperation) *xxx_EnumDataOperation {
	if op == nil {
		op = &xxx_EnumDataOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Data = o.Data
	op.RequiredDataLength = o.RequiredDataLength
	op.DataBlob = o.DataBlob
	op.Return = o.Return
	return op
}

func (o *EnumDataResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumDataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Data = op.Data
	o.RequiredDataLength = op.RequiredDataLength
	o.DataBlob = op.DataBlob
	o.Return = op.Return
}
func (o *EnumDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAllDataOperation structure represents the R_GetAllData operation
type xxx_GetAllDataOperation struct {
	This               *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat      `idl:"name:That" json:"that"`
	Handle             uint32              `idl:"name:hMDHandle" json:"handle"`
	Path               string              `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	Attributes         uint32              `idl:"name:dwMDAttributes" json:"attributes"`
	UserType           uint32              `idl:"name:dwMDUserType" json:"user_type"`
	DataType           uint32              `idl:"name:dwMDDataType" json:"data_type"`
	DataEntriesLength  uint32              `idl:"name:pdwMDNumDataEntries" json:"data_entries_length"`
	DataSetNumber      uint32              `idl:"name:pdwMDDataSetNumber" json:"data_set_number"`
	BufferSize         uint32              `idl:"name:dwMDBufferSize" json:"buffer_size"`
	RequiredBufferSize uint32              `idl:"name:pdwMDRequiredBufferSize" json:"required_buffer_size"`
	DataBlob           *imsa.IISCryptoBlob `idl:"name:ppDataBlob" json:"data_blob"`
	Return             int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAllDataOperation) OpNum() int { return 13 }

func (o *xxx_GetAllDataOperation) OpName() string { return "/IMSAdminBaseW/v0/R_GetAllData" }

func (o *xxx_GetAllDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.WriteData(o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_pszMDPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_pszMDPath); err != nil {
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
	// dwMDAttributes {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Attributes); err != nil {
			return err
		}
	}
	// dwMDUserType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.UserType); err != nil {
			return err
		}
	}
	// dwMDDataType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.DataType); err != nil {
			return err
		}
	}
	// dwMDBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.ReadData(&o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszMDPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_pszMDPath := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_pszMDPath, _ptr_pszMDPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwMDAttributes {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Attributes); err != nil {
			return err
		}
	}
	// dwMDUserType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.UserType); err != nil {
			return err
		}
	}
	// dwMDDataType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DataType); err != nil {
			return err
		}
	}
	// dwMDBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pdwMDNumDataEntries {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.DataEntriesLength); err != nil {
			return err
		}
	}
	// pdwMDDataSetNumber {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.DataSetNumber); err != nil {
			return err
		}
	}
	// pdwMDRequiredBufferSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequiredBufferSize); err != nil {
			return err
		}
	}
	// ppDataBlob {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IIS_CRYPTO_BLOB}(struct))
	{
		if o.DataBlob != nil {
			_ptr_ppDataBlob := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DataBlob != nil {
					if err := o.DataBlob.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&imsa.IISCryptoBlob{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DataBlob, _ptr_ppDataBlob); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pdwMDNumDataEntries {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DataEntriesLength); err != nil {
			return err
		}
	}
	// pdwMDDataSetNumber {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DataSetNumber); err != nil {
			return err
		}
	}
	// pdwMDRequiredBufferSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequiredBufferSize); err != nil {
			return err
		}
	}
	// ppDataBlob {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IIS_CRYPTO_BLOB}(struct))
	{
		_ptr_ppDataBlob := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DataBlob == nil {
				o.DataBlob = &imsa.IISCryptoBlob{}
			}
			if err := o.DataBlob.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppDataBlob := func(ptr interface{}) { o.DataBlob = *ptr.(**imsa.IISCryptoBlob) }
		if err := w.ReadPointer(&o.DataBlob, _s_ppDataBlob, _ptr_ppDataBlob); err != nil {
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

// GetAllDataRequest structure represents the R_GetAllData operation request
type GetAllDataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// hMDHandle: An unsigned 32-bit integer value containing an open metabase handle specifying
	// the key to be queried.
	Handle uint32 `idl:"name:hMDHandle" json:"handle"`
	// pszMDPath: A pointer to a Unicode string that contains the path of the node with
	// which the data to be returned is associated, relative to the path of the hMDHandle
	// parameter.
	Path string `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	// dwMDAttributes: Flags used to specify the data, as listed in the METADATA_RECORD
	// structure.
	Attributes uint32 `idl:"name:dwMDAttributes" json:"attributes"`
	// dwMDUserType: An integer value specifying the data to return based on user type.
	//
	//	+-----------------------------+-------------------------------------------------------------------------------+
	//	|                             |                                                                               |
	//	|            VALUE            |                                    MEANING                                    |
	//	|                             |                                                                               |
	//	+-----------------------------+-------------------------------------------------------------------------------+
	//	+-----------------------------+-------------------------------------------------------------------------------+
	//	| ALL_METADATA 0x00000000     | Returns all data, regardless of user type.                                    |
	//	+-----------------------------+-------------------------------------------------------------------------------+
	//	| ASP_MD_UT_APP 0x00000065    | Returns data specific to ASP application configuration.                       |
	//	+-----------------------------+-------------------------------------------------------------------------------+
	//	| IIS_MD_UT_FILE 0x00000002   | Returns data specific to a file, such as access permissions or logon methods. |
	//	+-----------------------------+-------------------------------------------------------------------------------+
	//	| IIS_MD_UT_SERVER 0x00000001 | Returns data specific to the server, such as ports in use and IP addresses.   |
	//	+-----------------------------+-------------------------------------------------------------------------------+
	//	| IIS_MD_UT_WAM 0x00000064    | Returns data specific to WAM.                                                 |
	//	+-----------------------------+-------------------------------------------------------------------------------+
	UserType uint32 `idl:"name:dwMDUserType" json:"user_type"`
	// dwMDDataType: An integer value specifying a data type. If this parameter is not set
	// to ALL_METADATA, the data item will be returned only if its data type matches the
	// specified type.
	//
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	|                              |                                                                                  |
	//	|            VALUE             |                                     MEANING                                      |
	//	|                              |                                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| ALL_METADATA 0x00000000      | Specifies all data, regardless of type.                                          |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| BINARY_METADATA 0x00000003   | Specifies binary data in any form.                                               |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| DWORD_METADATA 0x00000001    | Specifies all DWORD (unsigned 32-bit integer) data.                              |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| EXPANDSZ_METADATA 0x00000004 | Specifies all data that consists of a null-terminated string containing          |
	//	|                              | environment variables that are not expanded.                                     |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| MULTISZ_METADATA 0x00000005  | Specifies all data represented as an array of null-terminated strings,           |
	//	|                              | terminated by two null characters.                                               |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| STRING_METADATA 0x00000002   | Specifies all data consisting of a null-terminated ASCII string.                 |
	//	+------------------------------+----------------------------------------------------------------------------------+
	DataType uint32 `idl:"name:dwMDDataType" json:"data_type"`
	// dwMDBufferSize: An integer value specifying the size, in bytes, required to hold
	// the decrypted data returned by the ppDataBlob parameter.
	BufferSize uint32 `idl:"name:dwMDBufferSize" json:"buffer_size"`
}

func (o *GetAllDataRequest) xxx_ToOp(ctx context.Context, op *xxx_GetAllDataOperation) *xxx_GetAllDataOperation {
	if op == nil {
		op = &xxx_GetAllDataOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Handle = o.Handle
	op.Path = o.Path
	op.Attributes = o.Attributes
	op.UserType = o.UserType
	op.DataType = o.DataType
	op.BufferSize = o.BufferSize
	return op
}

func (o *GetAllDataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAllDataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Handle = op.Handle
	o.Path = op.Path
	o.Attributes = op.Attributes
	o.UserType = op.UserType
	o.DataType = op.DataType
	o.BufferSize = op.BufferSize
}
func (o *GetAllDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAllDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAllDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAllDataResponse structure represents the R_GetAllData operation response
type GetAllDataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pdwMDNumDataEntries: A pointer to an integer value that contains the number of entries
	// in the array of METADATA_GETALL_RECORD structures returned in the ppDataBlob parameter.
	DataEntriesLength uint32 `idl:"name:pdwMDNumDataEntries" json:"data_entries_length"`
	// pdwMDDataSetNumber: A pointer to an integer value used to identify the dataset number
	// for the metabase node whose data is being retrieved. The dataset number is obtained
	// by the GetDataSetNumber (section 3.1.4.18) method.
	DataSetNumber uint32 `idl:"name:pdwMDDataSetNumber" json:"data_set_number"`
	// pdwMDRequiredBufferSize: A pointer to an integer value that contains the buffer length
	// required, in bytes, to contain the decrypted data referenced by the ppDataBlob parameter.
	RequiredBufferSize uint32 `idl:"name:pdwMDRequiredBufferSize" json:"required_buffer_size"`
	// ppDataBlob: An IIS_CRYPTO_BLOB structure containing the requested values as encrypted
	// opaque data. The encrypted data returned in IIS_CRYPTO_BLOB is a contiguous buffer
	// containing an array of METADATA_GETALL_RECORD structures followed by the data referenced
	// by the METADATA_GETALL_RECORD structures.
	DataBlob *imsa.IISCryptoBlob `idl:"name:ppDataBlob" json:"data_blob"`
	// Return: The R_GetAllData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAllDataResponse) xxx_ToOp(ctx context.Context, op *xxx_GetAllDataOperation) *xxx_GetAllDataOperation {
	if op == nil {
		op = &xxx_GetAllDataOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.DataEntriesLength = o.DataEntriesLength
	op.DataSetNumber = o.DataSetNumber
	op.RequiredBufferSize = o.RequiredBufferSize
	op.DataBlob = o.DataBlob
	op.Return = o.Return
	return op
}

func (o *GetAllDataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAllDataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DataEntriesLength = op.DataEntriesLength
	o.DataSetNumber = op.DataSetNumber
	o.RequiredBufferSize = op.RequiredBufferSize
	o.DataBlob = op.DataBlob
	o.Return = op.Return
}
func (o *GetAllDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAllDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAllDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteAllDataOperation structure represents the DeleteAllData operation
type xxx_DeleteAllDataOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Handle   uint32         `idl:"name:hMDHandle" json:"handle"`
	Path     string         `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	UserType uint32         `idl:"name:dwMDUserType" json:"user_type"`
	DataType uint32         `idl:"name:dwMDDataType" json:"data_type"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteAllDataOperation) OpNum() int { return 14 }

func (o *xxx_DeleteAllDataOperation) OpName() string { return "/IMSAdminBaseW/v0/DeleteAllData" }

func (o *xxx_DeleteAllDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteAllDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.WriteData(o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_pszMDPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_pszMDPath); err != nil {
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
	// dwMDUserType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.UserType); err != nil {
			return err
		}
	}
	// dwMDDataType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.DataType); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteAllDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.ReadData(&o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszMDPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_pszMDPath := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_pszMDPath, _ptr_pszMDPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwMDUserType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.UserType); err != nil {
			return err
		}
	}
	// dwMDDataType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DataType); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteAllDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteAllDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteAllDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteAllDataRequest structure represents the DeleteAllData operation request
type DeleteAllDataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// hMDHandle: An unsigned 32-bit integer value containing an open metabase handle specifying
	// the node in the metabase where the key data is to be deleted.
	Handle uint32 `idl:"name:hMDHandle" json:"handle"`
	// pszMDPath: A pointer to a Unicode string that contains the path of the node with
	// which the data to be deleted is associated, relative to the path of the hMDHandle
	// parameter.
	Path string `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	// dwMDUserType: An integer value specifying the data to delete based on user type.
	//
	//	+-----------------------------+---------------------------------------------------------------------------------+
	//	|                             |                                                                                 |
	//	|            VALUE            |                                     MEANING                                     |
	//	|                             |                                                                                 |
	//	+-----------------------------+---------------------------------------------------------------------------------+
	//	+-----------------------------+---------------------------------------------------------------------------------+
	//	| ALL_METADATA 0x00000000     | Specifies all data, regardless of type.                                         |
	//	+-----------------------------+---------------------------------------------------------------------------------+
	//	| ASP_MD_UT_APP 0x00000065    | Specifies data specific to ASP application configuration.                       |
	//	+-----------------------------+---------------------------------------------------------------------------------+
	//	| IIS_MD_UT_FILE 0x00000002   | Specifies data specific to a file, such as access permissions or logon methods. |
	//	+-----------------------------+---------------------------------------------------------------------------------+
	//	| IIS_MD_UT_SERVER 0x00000001 | Specifies data specific to the server, such as ports in use and IP addresses.   |
	//	+-----------------------------+---------------------------------------------------------------------------------+
	//	| IIS_MD_UT_WAM 0x00000064    | Specifies data specific to WAM.                                                 |
	//	+-----------------------------+---------------------------------------------------------------------------------+
	UserType uint32 `idl:"name:dwMDUserType" json:"user_type"`
	// dwMDDataType: An integer value specifying a data type. If this parameter is not set
	// to ALL_METADATA, the data item will be removed only if its data type matches the
	// specified type.
	//
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	|                              |                                                                                  |
	//	|            VALUE             |                                     MEANING                                      |
	//	|                              |                                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| ALL_METADATA 0x00000000      | Specifies all data, regardless of type.                                          |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| BINARY_METADATA 0x00000003   | Specifies binary data in any form.                                               |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| DWORD_METADATA 0x00000001    | Specifies all DWORD (unsigned 32-bit integer) data.                              |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| EXPANDSZ_METADATA 0x00000004 | Specifies all data consisting of a string that includes the terminating null     |
	//	|                              | character, which contains unexpanded environment variables.                      |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| MULTISZ_METADATA 0x00000005  | Specifies all data represented as an array of strings, where each string         |
	//	|                              | contains two occurrences of the terminating null character.                      |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| STRING_METADATA 0x00000002   | Specifies all data consisting of an ASCII string that includes the terminating   |
	//	|                              | null character.                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	DataType uint32 `idl:"name:dwMDDataType" json:"data_type"`
}

func (o *DeleteAllDataRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteAllDataOperation) *xxx_DeleteAllDataOperation {
	if op == nil {
		op = &xxx_DeleteAllDataOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Handle = o.Handle
	op.Path = o.Path
	op.UserType = o.UserType
	op.DataType = o.DataType
	return op
}

func (o *DeleteAllDataRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteAllDataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Handle = op.Handle
	o.Path = op.Path
	o.UserType = op.UserType
	o.DataType = op.DataType
}
func (o *DeleteAllDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteAllDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteAllDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteAllDataResponse structure represents the DeleteAllData operation response
type DeleteAllDataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteAllData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteAllDataResponse) xxx_ToOp(ctx context.Context, op *xxx_DeleteAllDataOperation) *xxx_DeleteAllDataOperation {
	if op == nil {
		op = &xxx_DeleteAllDataOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DeleteAllDataResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteAllDataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteAllDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteAllDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteAllDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CopyDataOperation structure represents the CopyData operation
type xxx_CopyDataOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	SourceHandle      uint32         `idl:"name:hMDSourceHandle" json:"source_handle"`
	SourcePath        string         `idl:"name:pszMDSourcePath;string;pointer:unique" json:"source_path"`
	DestinationHandle uint32         `idl:"name:hMDDestHandle" json:"destination_handle"`
	DestinationPath   string         `idl:"name:pszMDDestPath;string;pointer:unique" json:"destination_path"`
	Attributes        uint32         `idl:"name:dwMDAttributes" json:"attributes"`
	UserType          uint32         `idl:"name:dwMDUserType" json:"user_type"`
	DataType          uint32         `idl:"name:dwMDDataType" json:"data_type"`
	CopyFlag          bool           `idl:"name:bMDCopyFlag" json:"copy_flag"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CopyDataOperation) OpNum() int { return 15 }

func (o *xxx_CopyDataOperation) OpName() string { return "/IMSAdminBaseW/v0/CopyData" }

func (o *xxx_CopyDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CopyDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// hMDSourceHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.WriteData(o.SourceHandle); err != nil {
			return err
		}
	}
	// pszMDSourcePath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.SourcePath != "" {
			_ptr_pszMDSourcePath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.SourcePath); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.SourcePath, _ptr_pszMDSourcePath); err != nil {
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
	// hMDDestHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.WriteData(o.DestinationHandle); err != nil {
			return err
		}
	}
	// pszMDDestPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.DestinationPath != "" {
			_ptr_pszMDDestPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.DestinationPath); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.DestinationPath, _ptr_pszMDDestPath); err != nil {
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
	// dwMDAttributes {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Attributes); err != nil {
			return err
		}
	}
	// dwMDUserType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.UserType); err != nil {
			return err
		}
	}
	// dwMDDataType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.DataType); err != nil {
			return err
		}
	}
	// bMDCopyFlag {in} (1:{alias=BOOL}(int32))
	{
		if !o.CopyFlag {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CopyDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// hMDSourceHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.ReadData(&o.SourceHandle); err != nil {
			return err
		}
	}
	// pszMDSourcePath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszMDSourcePath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.SourcePath); err != nil {
				return err
			}
			return nil
		})
		_s_pszMDSourcePath := func(ptr interface{}) { o.SourcePath = *ptr.(*string) }
		if err := w.ReadPointer(&o.SourcePath, _s_pszMDSourcePath, _ptr_pszMDSourcePath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// hMDDestHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.ReadData(&o.DestinationHandle); err != nil {
			return err
		}
	}
	// pszMDDestPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszMDDestPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.DestinationPath); err != nil {
				return err
			}
			return nil
		})
		_s_pszMDDestPath := func(ptr interface{}) { o.DestinationPath = *ptr.(*string) }
		if err := w.ReadPointer(&o.DestinationPath, _s_pszMDDestPath, _ptr_pszMDDestPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwMDAttributes {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Attributes); err != nil {
			return err
		}
	}
	// dwMDUserType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.UserType); err != nil {
			return err
		}
	}
	// dwMDDataType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DataType); err != nil {
			return err
		}
	}
	// bMDCopyFlag {in} (1:{alias=BOOL}(int32))
	{
		var _bCopyFlag int32
		if err := w.ReadData(&_bCopyFlag); err != nil {
			return err
		}
		o.CopyFlag = _bCopyFlag != 0
	}
	return nil
}

func (o *xxx_CopyDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CopyDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CopyDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CopyDataRequest structure represents the CopyData operation request
type CopyDataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// hMDSourceHandle: An unsigned 32-bit integer value containing an open metabase handle
	// specifying the source node from which the data is to be copied or moved.
	SourceHandle uint32 `idl:"name:hMDSourceHandle" json:"source_handle"`
	// pszMDSourcePath: A pointer to a Unicode string that contains the path of the node
	// with which the source data is associated, relative to the path of the hMDSourceHandle
	// parameter.
	SourcePath string `idl:"name:pszMDSourcePath;string;pointer:unique" json:"source_path"`
	// hMDDestHandle: An unsigned 32-bit integer value containing an open metabase handle
	// specifying the destination node to which the data is to be copied or moved.
	DestinationHandle uint32 `idl:"name:hMDDestHandle" json:"destination_handle"`
	// pszMDDestPath: A pointer to a Unicode string that contains the path of the node for
	// data to be copied to or moved to, relative to the path of the hMDDestHandle parameter.
	DestinationPath string `idl:"name:pszMDDestPath;string;pointer:unique" json:"destination_path"`
	// dwMDAttributes: Flags used to filter the data, as specified in the METADATA_RECORD
	// structure.
	Attributes uint32 `idl:"name:dwMDAttributes" json:"attributes"`
	// dwMDUserType: An integer value specifying the data to copy based on the user type.
	//
	//	+-----------------------------+---------------------------------------------------------------------------------+
	//	|                             |                                                                                 |
	//	|            VALUE            |                                     MEANING                                     |
	//	|                             |                                                                                 |
	//	+-----------------------------+---------------------------------------------------------------------------------+
	//	+-----------------------------+---------------------------------------------------------------------------------+
	//	| ALL_METADATA 0x00000000     | Specifies all data, regardless of user type.                                    |
	//	+-----------------------------+---------------------------------------------------------------------------------+
	//	| ASP_MD_UT_APP 0x00000065    | Specifies data specific to ASP application configuration.                       |
	//	+-----------------------------+---------------------------------------------------------------------------------+
	//	| IIS_MD_UT_FILE 0x00000002   | Specifies data specific to a file, such as access permissions or logon methods. |
	//	+-----------------------------+---------------------------------------------------------------------------------+
	//	| IIS_MD_UT_SERVER 0x00000001 | Specifies data specific to the server, such as ports in use and IP addresses.   |
	//	+-----------------------------+---------------------------------------------------------------------------------+
	//	| IIS_MD_UT_WAM 0x00000064    | Specifies data specific to WAM.                                                 |
	//	+-----------------------------+---------------------------------------------------------------------------------+
	UserType uint32 `idl:"name:dwMDUserType" json:"user_type"`
	// dwMDDataType: An integer value specifying a data type. If this parameter is not set
	// to ALL_METADATA, the data item will be copied only if its data type matches the specified
	// type.
	//
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	|                              |                                                                                  |
	//	|            VALUE             |                                     MEANING                                      |
	//	|                              |                                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| ALL_METADATA 0x00000000      | Specifies all data, regardless of type.                                          |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| BINARY_METADATA 0x00000003   | Specifies binary data in any form.                                               |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| DWORD_METADATA 0x00000001    | Specifies all DWORD (unsigned 32-bit integer) data.                              |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| EXPANDSZ_METADATA 0x00000004 | Specifies all data consisting of a string that includes the terminating null     |
	//	|                              | character, which contains unexpanded environment variables.                      |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| MULTISZ_METADATA 0x00000005  | Specifies all data represented as an array of strings, where each string         |
	//	|                              | contains two occurrences of the terminating null character.                      |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| STRING_METADATA 0x00000002   | Specifies all data consisting of an ASCII string that includes the terminating   |
	//	|                              | null character.                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	DataType uint32 `idl:"name:dwMDDataType" json:"data_type"`
	// bMDCopyFlag: A Boolean value that specifies whether to copy or move the data. If
	// this parameter is set to TRUE, the data is copied. If it is FALSE, the data is moved.
	CopyFlag bool `idl:"name:bMDCopyFlag" json:"copy_flag"`
}

func (o *CopyDataRequest) xxx_ToOp(ctx context.Context, op *xxx_CopyDataOperation) *xxx_CopyDataOperation {
	if op == nil {
		op = &xxx_CopyDataOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.SourceHandle = o.SourceHandle
	op.SourcePath = o.SourcePath
	op.DestinationHandle = o.DestinationHandle
	op.DestinationPath = o.DestinationPath
	op.Attributes = o.Attributes
	op.UserType = o.UserType
	op.DataType = o.DataType
	op.CopyFlag = o.CopyFlag
	return op
}

func (o *CopyDataRequest) xxx_FromOp(ctx context.Context, op *xxx_CopyDataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SourceHandle = op.SourceHandle
	o.SourcePath = op.SourcePath
	o.DestinationHandle = op.DestinationHandle
	o.DestinationPath = op.DestinationPath
	o.Attributes = op.Attributes
	o.UserType = op.UserType
	o.DataType = op.DataType
	o.CopyFlag = op.CopyFlag
}
func (o *CopyDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CopyDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CopyDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CopyDataResponse structure represents the CopyData operation response
type CopyDataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CopyData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CopyDataResponse) xxx_ToOp(ctx context.Context, op *xxx_CopyDataOperation) *xxx_CopyDataOperation {
	if op == nil {
		op = &xxx_CopyDataOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CopyDataResponse) xxx_FromOp(ctx context.Context, op *xxx_CopyDataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CopyDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CopyDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CopyDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDataPathsOperation structure represents the GetDataPaths operation
type xxx_GetDataPathsOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	Handle             uint32         `idl:"name:hMDHandle" json:"handle"`
	Path               string         `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	ID                 uint32         `idl:"name:dwMDIdentifier" json:"id"`
	DataType           uint32         `idl:"name:dwMDDataType" json:"data_type"`
	BufferSize         uint32         `idl:"name:dwMDBufferSize" json:"buffer_size"`
	Buffer             string         `idl:"name:pszBuffer;size_is:(dwMDBufferSize)" json:"buffer"`
	RequiredBufferSize uint32         `idl:"name:pdwMDRequiredBufferSize" json:"required_buffer_size"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDataPathsOperation) OpNum() int { return 16 }

func (o *xxx_GetDataPathsOperation) OpName() string { return "/IMSAdminBaseW/v0/GetDataPaths" }

func (o *xxx_GetDataPathsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDataPathsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.WriteData(o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_pszMDPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_pszMDPath); err != nil {
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
	// dwMDIdentifier {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ID); err != nil {
			return err
		}
	}
	// dwMDDataType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.DataType); err != nil {
			return err
		}
	}
	// dwMDBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDataPathsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.ReadData(&o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszMDPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_pszMDPath := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_pszMDPath, _ptr_pszMDPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwMDIdentifier {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ID); err != nil {
			return err
		}
	}
	// dwMDDataType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DataType); err != nil {
			return err
		}
	}
	// dwMDBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDataPathsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDataPathsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pszBuffer {out} (1:{pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,size_is=dwMDBufferSize,string](wchar))
	{
		dimSize1 := uint64(o.BufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
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
	}
	// pdwMDRequiredBufferSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequiredBufferSize); err != nil {
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

func (o *xxx_GetDataPathsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pszBuffer {out} (1:{pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,size_is=dwMDBufferSize,string](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
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
	}
	// pdwMDRequiredBufferSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequiredBufferSize); err != nil {
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

// GetDataPathsRequest structure represents the GetDataPaths operation request
type GetDataPathsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// hMDHandle: An unsigned 32-bit integer value containing an open metabase handle specifying
	// the key to be queried.
	Handle uint32 `idl:"name:hMDHandle" json:"handle"`
	// pszMDPath: A pointer to a Unicode string that contains the path of the node to be
	// queried, relative to the hMDHandle parameter.
	Path string `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	// dwMDIdentifier: An integer value identifying the data to be queried.
	ID uint32 `idl:"name:dwMDIdentifier" json:"id"`
	// dwMDDataType: An integer value specifying a data type. If this parameter is not set
	// to ALL_METADATA, the data item will be returned only if its data type matches the
	// specified type.
	//
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	|                              |                                                                                  |
	//	|            VALUE             |                                     MEANING                                      |
	//	|                              |                                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| ALL_METADATA 0x00000000      | Specifies all data, regardless of type.                                          |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| BINARY_METADATA 0x00000003   | Specifies binary data in any form.                                               |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| DWORD_METADATA 0x00000001    | Specifies all DWORD (unsigned 32-bit integer) data.                              |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| EXPANDSZ_METADATA 0x00000004 | Specifies all data consisting of a string that includes the terminating null     |
	//	|                              | character, which contains unexpanded environment variables.                      |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| MULTISZ_METADATA 0x00000005  | Specifies all data represented as an array of strings, where each string         |
	//	|                              | contains two occurrences of the terminating null character.                      |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| STRING_METADATA 0x00000002   | Specifies all data consisting of an ASCII string that includes the terminating   |
	//	|                              | null character.                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	DataType uint32 `idl:"name:dwMDDataType" json:"data_type"`
	// dwMDBufferSize: An integer value specifying the size, in WCHARs, of the pszBuffer
	// parameter.
	BufferSize uint32 `idl:"name:dwMDBufferSize" json:"buffer_size"`
}

func (o *GetDataPathsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetDataPathsOperation) *xxx_GetDataPathsOperation {
	if op == nil {
		op = &xxx_GetDataPathsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Handle = o.Handle
	op.Path = o.Path
	op.ID = o.ID
	op.DataType = o.DataType
	op.BufferSize = o.BufferSize
	return op
}

func (o *GetDataPathsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDataPathsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Handle = op.Handle
	o.Path = op.Path
	o.ID = op.ID
	o.DataType = op.DataType
	o.BufferSize = op.BufferSize
}
func (o *GetDataPathsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetDataPathsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDataPathsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDataPathsResponse structure represents the GetDataPaths operation response
type GetDataPathsResponse struct {
	// XXX: dwMDBufferSize is an implicit input depedency for output parameters
	BufferSize uint32 `idl:"name:dwMDBufferSize" json:"buffer_size"`
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pszBuffer: A pointer to a buffer that contains the retrieved data. If the method
	// call is successful, the buffer will contain a contiguous sequence of null-terminated
	// strings in "multi-string" format. Each string in the sequence is a metabase path
	// at which data matching the dwMDIdentifier and dwMDDataType fields were found.
	Buffer string `idl:"name:pszBuffer;size_is:(dwMDBufferSize)" json:"buffer"`
	// pdwMDRequiredBufferSize: A pointer to an integer value that contains the buffer length
	// required, in WCHARs.
	RequiredBufferSize uint32 `idl:"name:pdwMDRequiredBufferSize" json:"required_buffer_size"`
	// Return: The GetDataPaths return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDataPathsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetDataPathsOperation) *xxx_GetDataPathsOperation {
	if op == nil {
		op = &xxx_GetDataPathsOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.BufferSize == uint32(0) {
		op.BufferSize = o.BufferSize
	}

	op.That = o.That
	op.Buffer = o.Buffer
	op.RequiredBufferSize = o.RequiredBufferSize
	op.Return = o.Return
	return op
}

func (o *GetDataPathsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDataPathsOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.BufferSize = op.BufferSize

	o.That = op.That
	o.Buffer = op.Buffer
	o.RequiredBufferSize = op.RequiredBufferSize
	o.Return = op.Return
}
func (o *GetDataPathsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetDataPathsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDataPathsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenKeyOperation structure represents the OpenKey operation
type xxx_OpenKeyOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	Handle          uint32         `idl:"name:hMDHandle" json:"handle"`
	Path            string         `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	AccessRequested uint32         `idl:"name:dwMDAccessRequested" json:"access_requested"`
	Timeout         uint32         `idl:"name:dwMDTimeOut" json:"timeout"`
	MDNewHandle     uint32         `idl:"name:phMDNewHandle" json:"md_new_handle"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenKeyOperation) OpNum() int { return 17 }

func (o *xxx_OpenKeyOperation) OpName() string { return "/IMSAdminBaseW/v0/OpenKey" }

func (o *xxx_OpenKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.WriteData(o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_pszMDPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_pszMDPath); err != nil {
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
	// dwMDAccessRequested {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.AccessRequested); err != nil {
			return err
		}
	}
	// dwMDTimeOut {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Timeout); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.ReadData(&o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszMDPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_pszMDPath := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_pszMDPath, _ptr_pszMDPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwMDAccessRequested {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.AccessRequested); err != nil {
			return err
		}
	}
	// dwMDTimeOut {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Timeout); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// phMDNewHandle {out} (1:{pointer=ref}*(1))(2:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.WriteData(o.MDNewHandle); err != nil {
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

func (o *xxx_OpenKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// phMDNewHandle {out} (1:{pointer=ref}*(1))(2:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.ReadData(&o.MDNewHandle); err != nil {
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

// OpenKeyRequest structure represents the OpenKey operation request
type OpenKeyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// hMDHandle: An unsigned 32-bit integer value containing a handle to a node in the
	// metabase with read permissions as returned by the OpenKey method or the metabase
	// master root handle (0x00000000).
	Handle uint32 `idl:"name:hMDHandle" json:"handle"`
	// pszMDPath: A pointer to a Unicode string that contains the path of the node to be
	// opened, relative to the hMDHandle parameter.
	Path string `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	// dwMDAccessRequested: A set of bit flags specifying the requested permissions for
	// the handle. This parameter MUST be set to at least one of the following values.
	//
	//	+--------------------------------------+----------------------------+
	//	|                                      |                            |
	//	|                VALUE                 |          MEANING           |
	//	|                                      |                            |
	//	+--------------------------------------+----------------------------+
	//	+--------------------------------------+----------------------------+
	//	| METADATA_PERMISSION_READ 0x00000001  | Open the node for reading. |
	//	+--------------------------------------+----------------------------+
	//	| METADATA_PERMISSION_WRITE 0x00000002 | Open the node for writing. |
	//	+--------------------------------------+----------------------------+
	AccessRequested uint32 `idl:"name:dwMDAccessRequested" json:"access_requested"`
	// dwMDTimeOut: An unsigned 32-bit integer value specifying the time, in milliseconds,
	// for the method to wait on a successful open operation.
	Timeout uint32 `idl:"name:dwMDTimeOut" json:"timeout"`
}

func (o *OpenKeyRequest) xxx_ToOp(ctx context.Context, op *xxx_OpenKeyOperation) *xxx_OpenKeyOperation {
	if op == nil {
		op = &xxx_OpenKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Handle = o.Handle
	op.Path = o.Path
	op.AccessRequested = o.AccessRequested
	op.Timeout = o.Timeout
	return op
}

func (o *OpenKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenKeyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Handle = op.Handle
	o.Path = op.Path
	o.AccessRequested = op.AccessRequested
	o.Timeout = op.Timeout
}
func (o *OpenKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenKeyResponse structure represents the OpenKey operation response
type OpenKeyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// phMDNewHandle: A pointer to the newly opened metadata handle (see DWORD).
	MDNewHandle uint32 `idl:"name:phMDNewHandle" json:"md_new_handle"`
	// Return: The OpenKey return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenKeyResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenKeyOperation) *xxx_OpenKeyOperation {
	if op == nil {
		op = &xxx_OpenKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.MDNewHandle = o.MDNewHandle
	op.Return = o.Return
	return op
}

func (o *OpenKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenKeyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MDNewHandle = op.MDNewHandle
	o.Return = op.Return
}
func (o *OpenKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseKeyOperation structure represents the CloseKey operation
type xxx_CloseKeyOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Handle uint32         `idl:"name:hMDHandle" json:"handle"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseKeyOperation) OpNum() int { return 18 }

func (o *xxx_CloseKeyOperation) OpName() string { return "/IMSAdminBaseW/v0/CloseKey" }

func (o *xxx_CloseKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.WriteData(o.Handle); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.ReadData(&o.Handle); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CloseKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CloseKeyRequest structure represents the CloseKey operation request
type CloseKeyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// hMDHandle: An unsigned 32-bit integer value containing the handle to close, as returned
	// by the OpenKey method.
	Handle uint32 `idl:"name:hMDHandle" json:"handle"`
}

func (o *CloseKeyRequest) xxx_ToOp(ctx context.Context, op *xxx_CloseKeyOperation) *xxx_CloseKeyOperation {
	if op == nil {
		op = &xxx_CloseKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Handle = o.Handle
	return op
}

func (o *CloseKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseKeyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Handle = op.Handle
}
func (o *CloseKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CloseKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseKeyResponse structure represents the CloseKey operation response
type CloseKeyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CloseKey return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CloseKeyResponse) xxx_ToOp(ctx context.Context, op *xxx_CloseKeyOperation) *xxx_CloseKeyOperation {
	if op == nil {
		op = &xxx_CloseKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CloseKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseKeyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CloseKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CloseKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ChangePermissionsOperation structure represents the ChangePermissions operation
type xxx_ChangePermissionsOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	Handle          uint32         `idl:"name:hMDHandle" json:"handle"`
	Timeout         uint32         `idl:"name:dwMDTimeOut" json:"timeout"`
	AccessRequested uint32         `idl:"name:dwMDAccessRequested" json:"access_requested"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ChangePermissionsOperation) OpNum() int { return 19 }

func (o *xxx_ChangePermissionsOperation) OpName() string {
	return "/IMSAdminBaseW/v0/ChangePermissions"
}

func (o *xxx_ChangePermissionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangePermissionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.WriteData(o.Handle); err != nil {
			return err
		}
	}
	// dwMDTimeOut {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Timeout); err != nil {
			return err
		}
	}
	// dwMDAccessRequested {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.AccessRequested); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangePermissionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.ReadData(&o.Handle); err != nil {
			return err
		}
	}
	// dwMDTimeOut {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Timeout); err != nil {
			return err
		}
	}
	// dwMDAccessRequested {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.AccessRequested); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangePermissionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangePermissionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ChangePermissionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ChangePermissionsRequest structure represents the ChangePermissions operation request
type ChangePermissionsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// hMDHandle: An unsigned 32-bit integer value containing the handle to change the permissions
	// for, as returned by the OpenKey method.
	Handle uint32 `idl:"name:hMDHandle" json:"handle"`
	// dwMDTimeOut: An integer value specifying the time, in milliseconds, for the method
	// to wait on a successful permission change operation.
	Timeout uint32 `idl:"name:dwMDTimeOut" json:"timeout"`
	// dwMDAccessRequested: A set of bit flags specifying the requested permissions for
	// the handle. This parameter MUST be set to at least one of the following values.
	//
	//	+--------------------------------------+----------------------------+
	//	|                                      |                            |
	//	|                VALUE                 |          MEANING           |
	//	|                                      |                            |
	//	+--------------------------------------+----------------------------+
	//	+--------------------------------------+----------------------------+
	//	| METADATA_PERMISSION_READ 0x00000001  | Open the node for reading. |
	//	+--------------------------------------+----------------------------+
	//	| METADATA_PERMISSION_WRITE 0x00000002 | Open the node for writing. |
	//	+--------------------------------------+----------------------------+
	AccessRequested uint32 `idl:"name:dwMDAccessRequested" json:"access_requested"`
}

func (o *ChangePermissionsRequest) xxx_ToOp(ctx context.Context, op *xxx_ChangePermissionsOperation) *xxx_ChangePermissionsOperation {
	if op == nil {
		op = &xxx_ChangePermissionsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Handle = o.Handle
	op.Timeout = o.Timeout
	op.AccessRequested = o.AccessRequested
	return op
}

func (o *ChangePermissionsRequest) xxx_FromOp(ctx context.Context, op *xxx_ChangePermissionsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Handle = op.Handle
	o.Timeout = op.Timeout
	o.AccessRequested = op.AccessRequested
}
func (o *ChangePermissionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ChangePermissionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ChangePermissionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ChangePermissionsResponse structure represents the ChangePermissions operation response
type ChangePermissionsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ChangePermissions return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ChangePermissionsResponse) xxx_ToOp(ctx context.Context, op *xxx_ChangePermissionsOperation) *xxx_ChangePermissionsOperation {
	if op == nil {
		op = &xxx_ChangePermissionsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ChangePermissionsResponse) xxx_FromOp(ctx context.Context, op *xxx_ChangePermissionsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ChangePermissionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ChangePermissionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ChangePermissionsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SaveDataOperation structure represents the SaveData operation
type xxx_SaveDataOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SaveDataOperation) OpNum() int { return 20 }

func (o *xxx_SaveDataOperation) OpName() string { return "/IMSAdminBaseW/v0/SaveData" }

func (o *xxx_SaveDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SaveDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SaveDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_SaveDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SaveDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SaveDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SaveDataRequest structure represents the SaveData operation request
type SaveDataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *SaveDataRequest) xxx_ToOp(ctx context.Context, op *xxx_SaveDataOperation) *xxx_SaveDataOperation {
	if op == nil {
		op = &xxx_SaveDataOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *SaveDataRequest) xxx_FromOp(ctx context.Context, op *xxx_SaveDataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *SaveDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SaveDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SaveDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SaveDataResponse structure represents the SaveData operation response
type SaveDataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SaveData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SaveDataResponse) xxx_ToOp(ctx context.Context, op *xxx_SaveDataOperation) *xxx_SaveDataOperation {
	if op == nil {
		op = &xxx_SaveDataOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SaveDataResponse) xxx_FromOp(ctx context.Context, op *xxx_SaveDataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SaveDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SaveDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SaveDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetHandleInfoOperation structure represents the GetHandleInfo operation
type xxx_GetHandleInfoOperation struct {
	This   *dcom.ORPCThis           `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat           `idl:"name:That" json:"that"`
	Handle uint32                   `idl:"name:hMDHandle" json:"handle"`
	Info   *imsa.MetadataHandleInfo `idl:"name:pmdhiInfo" json:"info"`
	Return int32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_GetHandleInfoOperation) OpNum() int { return 21 }

func (o *xxx_GetHandleInfoOperation) OpName() string { return "/IMSAdminBaseW/v0/GetHandleInfo" }

func (o *xxx_GetHandleInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetHandleInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.WriteData(o.Handle); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetHandleInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.ReadData(&o.Handle); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetHandleInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetHandleInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pmdhiInfo {out} (1:{pointer=ref}*(1))(2:{alias=METADATA_HANDLE_INFO}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&imsa.MetadataHandleInfo{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetHandleInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pmdhiInfo {out} (1:{pointer=ref}*(1))(2:{alias=METADATA_HANDLE_INFO}(struct))
	{
		if o.Info == nil {
			o.Info = &imsa.MetadataHandleInfo{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
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

// GetHandleInfoRequest structure represents the GetHandleInfo operation request
type GetHandleInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// hMDHandle: An unsigned 32-bit integer value containing a handle to a node in the
	// metabase as returned by the OpenKey method.
	Handle uint32 `idl:"name:hMDHandle" json:"handle"`
}

func (o *GetHandleInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetHandleInfoOperation) *xxx_GetHandleInfoOperation {
	if op == nil {
		op = &xxx_GetHandleInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Handle = o.Handle
	return op
}

func (o *GetHandleInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetHandleInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Handle = op.Handle
}
func (o *GetHandleInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetHandleInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetHandleInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetHandleInfoResponse structure represents the GetHandleInfo operation response
type GetHandleInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pmdhiInfo: A pointer to a METADATA_HANDLE_INFO structure containing information about
	// the specified handle.
	Info *imsa.MetadataHandleInfo `idl:"name:pmdhiInfo" json:"info"`
	// Return: The GetHandleInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetHandleInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetHandleInfoOperation) *xxx_GetHandleInfoOperation {
	if op == nil {
		op = &xxx_GetHandleInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *GetHandleInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetHandleInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Info = op.Info
	o.Return = op.Return
}
func (o *GetHandleInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetHandleInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetHandleInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSystemChangeNumberOperation structure represents the GetSystemChangeNumber operation
type xxx_GetSystemChangeNumberOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	SystemChangeNumber uint32         `idl:"name:pdwSystemChangeNumber" json:"system_change_number"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSystemChangeNumberOperation) OpNum() int { return 22 }

func (o *xxx_GetSystemChangeNumberOperation) OpName() string {
	return "/IMSAdminBaseW/v0/GetSystemChangeNumber"
}

func (o *xxx_GetSystemChangeNumberOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSystemChangeNumberOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSystemChangeNumberOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSystemChangeNumberOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSystemChangeNumberOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pdwSystemChangeNumber {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SystemChangeNumber); err != nil {
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

func (o *xxx_GetSystemChangeNumberOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pdwSystemChangeNumber {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SystemChangeNumber); err != nil {
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

// GetSystemChangeNumberRequest structure represents the GetSystemChangeNumber operation request
type GetSystemChangeNumberRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSystemChangeNumberRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSystemChangeNumberOperation) *xxx_GetSystemChangeNumberOperation {
	if op == nil {
		op = &xxx_GetSystemChangeNumberOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetSystemChangeNumberRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSystemChangeNumberOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSystemChangeNumberRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSystemChangeNumberRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSystemChangeNumberOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSystemChangeNumberResponse structure represents the GetSystemChangeNumber operation response
type GetSystemChangeNumberResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pdwSystemChangeNumber: A pointer to an unsigned 32-bit integer value containing the
	// system change number. This number is increased each time the metabase is updated.
	SystemChangeNumber uint32 `idl:"name:pdwSystemChangeNumber" json:"system_change_number"`
	// Return: The GetSystemChangeNumber return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSystemChangeNumberResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSystemChangeNumberOperation) *xxx_GetSystemChangeNumberOperation {
	if op == nil {
		op = &xxx_GetSystemChangeNumberOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.SystemChangeNumber = o.SystemChangeNumber
	op.Return = o.Return
	return op
}

func (o *GetSystemChangeNumberResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSystemChangeNumberOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SystemChangeNumber = op.SystemChangeNumber
	o.Return = op.Return
}
func (o *GetSystemChangeNumberResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSystemChangeNumberResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSystemChangeNumberOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDataSetNumberOperation structure represents the GetDataSetNumber operation
type xxx_GetDataSetNumberOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	Handle        uint32         `idl:"name:hMDHandle" json:"handle"`
	Path          string         `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	DataSetNumber uint32         `idl:"name:pdwMDDataSetNumber" json:"data_set_number"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDataSetNumberOperation) OpNum() int { return 23 }

func (o *xxx_GetDataSetNumberOperation) OpName() string { return "/IMSAdminBaseW/v0/GetDataSetNumber" }

func (o *xxx_GetDataSetNumberOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDataSetNumberOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.WriteData(o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_pszMDPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_pszMDPath); err != nil {
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

func (o *xxx_GetDataSetNumberOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.ReadData(&o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszMDPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_pszMDPath := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_pszMDPath, _ptr_pszMDPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDataSetNumberOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDataSetNumberOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pdwMDDataSetNumber {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.DataSetNumber); err != nil {
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

func (o *xxx_GetDataSetNumberOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pdwMDDataSetNumber {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DataSetNumber); err != nil {
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

// GetDataSetNumberRequest structure represents the GetDataSetNumber operation request
type GetDataSetNumberRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// hMDHandle: An unsigned 32-bit integer value containing an open metabase handle specifying
	// the key to be queried.
	Handle uint32 `idl:"name:hMDHandle" json:"handle"`
	// pszMDPath: A pointer to a Unicode string that contains the path of the node to have
	// its dataset number retrieved, relative to the path of the hMDHandle parameter.
	Path string `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
}

func (o *GetDataSetNumberRequest) xxx_ToOp(ctx context.Context, op *xxx_GetDataSetNumberOperation) *xxx_GetDataSetNumberOperation {
	if op == nil {
		op = &xxx_GetDataSetNumberOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Handle = o.Handle
	op.Path = o.Path
	return op
}

func (o *GetDataSetNumberRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDataSetNumberOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Handle = op.Handle
	o.Path = op.Path
}
func (o *GetDataSetNumberRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetDataSetNumberRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDataSetNumberOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDataSetNumberResponse structure represents the GetDataSetNumber operation response
type GetDataSetNumberResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pdwMDDataSetNumber: A pointer to an integer value that returns the number associated
	// with this dataset. This value can be used to identify datasets common to multiple
	// nodes.
	DataSetNumber uint32 `idl:"name:pdwMDDataSetNumber" json:"data_set_number"`
	// Return: The GetDataSetNumber return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDataSetNumberResponse) xxx_ToOp(ctx context.Context, op *xxx_GetDataSetNumberOperation) *xxx_GetDataSetNumberOperation {
	if op == nil {
		op = &xxx_GetDataSetNumberOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.DataSetNumber = o.DataSetNumber
	op.Return = o.Return
	return op
}

func (o *GetDataSetNumberResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDataSetNumberOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DataSetNumber = op.DataSetNumber
	o.Return = op.Return
}
func (o *GetDataSetNumberResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetDataSetNumberResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDataSetNumberOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetLastChangeTimeOperation structure represents the SetLastChangeTime operation
type xxx_SetLastChangeTimeOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	Handle         uint32         `idl:"name:hMDHandle" json:"handle"`
	Path           string         `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	LastChangeTime *dtyp.Filetime `idl:"name:pftMDLastChangeTime" json:"last_change_time"`
	LocalTime      bool           `idl:"name:bLocalTime" json:"local_time"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetLastChangeTimeOperation) OpNum() int { return 24 }

func (o *xxx_SetLastChangeTimeOperation) OpName() string {
	return "/IMSAdminBaseW/v0/SetLastChangeTime"
}

func (o *xxx_SetLastChangeTimeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLastChangeTimeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.WriteData(o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_pszMDPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_pszMDPath); err != nil {
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
	// pftMDLastChangeTime {in} (1:{alias=PFILETIME}*(1))(2:{alias=FILETIME}(struct))
	{
		if o.LastChangeTime != nil {
			if err := o.LastChangeTime.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// bLocalTime {in} (1:{alias=BOOL}(int32))
	{
		if !o.LocalTime {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_SetLastChangeTimeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.ReadData(&o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszMDPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_pszMDPath := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_pszMDPath, _ptr_pszMDPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pftMDLastChangeTime {in} (1:{alias=PFILETIME,pointer=ref}*(1))(2:{alias=FILETIME}(struct))
	{
		if o.LastChangeTime == nil {
			o.LastChangeTime = &dtyp.Filetime{}
		}
		if err := o.LastChangeTime.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// bLocalTime {in} (1:{alias=BOOL}(int32))
	{
		var _bLocalTime int32
		if err := w.ReadData(&_bLocalTime); err != nil {
			return err
		}
		o.LocalTime = _bLocalTime != 0
	}
	return nil
}

func (o *xxx_SetLastChangeTimeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLastChangeTimeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetLastChangeTimeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetLastChangeTimeRequest structure represents the SetLastChangeTime operation request
type SetLastChangeTimeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// hMDHandle: An unsigned 32-bit integer value containing a handle to a node in the
	// metabase as returned by the OpenKey method.
	Handle uint32 `idl:"name:hMDHandle" json:"handle"`
	// pszMDPath: A pointer to a Unicode string containing the path of the node to be set,
	// relative to the path of the hMDHandle parameter.
	Path string `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	// pftMDLastChangeTime: A pointer to a FILETIME structure that contains the last change
	// time to set for the node.
	LastChangeTime *dtyp.Filetime `idl:"name:pftMDLastChangeTime" json:"last_change_time"`
	// bLocalTime: A Boolean value indicating whether the time value specified in the pftMDLastChangeTime
	// parameter is local time (TRUE) or UTC time (FALSE).
	LocalTime bool `idl:"name:bLocalTime" json:"local_time"`
}

func (o *SetLastChangeTimeRequest) xxx_ToOp(ctx context.Context, op *xxx_SetLastChangeTimeOperation) *xxx_SetLastChangeTimeOperation {
	if op == nil {
		op = &xxx_SetLastChangeTimeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Handle = o.Handle
	op.Path = o.Path
	op.LastChangeTime = o.LastChangeTime
	op.LocalTime = o.LocalTime
	return op
}

func (o *SetLastChangeTimeRequest) xxx_FromOp(ctx context.Context, op *xxx_SetLastChangeTimeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Handle = op.Handle
	o.Path = op.Path
	o.LastChangeTime = op.LastChangeTime
	o.LocalTime = op.LocalTime
}
func (o *SetLastChangeTimeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetLastChangeTimeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetLastChangeTimeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetLastChangeTimeResponse structure represents the SetLastChangeTime operation response
type SetLastChangeTimeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetLastChangeTime return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetLastChangeTimeResponse) xxx_ToOp(ctx context.Context, op *xxx_SetLastChangeTimeOperation) *xxx_SetLastChangeTimeOperation {
	if op == nil {
		op = &xxx_SetLastChangeTimeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetLastChangeTimeResponse) xxx_FromOp(ctx context.Context, op *xxx_SetLastChangeTimeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetLastChangeTimeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetLastChangeTimeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetLastChangeTimeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetLastChangeTimeOperation structure represents the GetLastChangeTime operation
type xxx_GetLastChangeTimeOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	Handle         uint32         `idl:"name:hMDHandle" json:"handle"`
	Path           string         `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	LastChangeTime *dtyp.Filetime `idl:"name:pftMDLastChangeTime" json:"last_change_time"`
	LocalTime      bool           `idl:"name:bLocalTime" json:"local_time"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetLastChangeTimeOperation) OpNum() int { return 25 }

func (o *xxx_GetLastChangeTimeOperation) OpName() string {
	return "/IMSAdminBaseW/v0/GetLastChangeTime"
}

func (o *xxx_GetLastChangeTimeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLastChangeTimeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.WriteData(o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_pszMDPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_pszMDPath); err != nil {
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
	// bLocalTime {in} (1:{alias=BOOL}(int32))
	{
		if !o.LocalTime {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetLastChangeTimeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// hMDHandle {in} (1:{alias=METADATA_HANDLE}(uint32))
	{
		if err := w.ReadData(&o.Handle); err != nil {
			return err
		}
	}
	// pszMDPath {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszMDPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_pszMDPath := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_pszMDPath, _ptr_pszMDPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bLocalTime {in} (1:{alias=BOOL}(int32))
	{
		var _bLocalTime int32
		if err := w.ReadData(&_bLocalTime); err != nil {
			return err
		}
		o.LocalTime = _bLocalTime != 0
	}
	return nil
}

func (o *xxx_GetLastChangeTimeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLastChangeTimeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pftMDLastChangeTime {out} (1:{alias=PFILETIME}*(1))(2:{alias=FILETIME}(struct))
	{
		if o.LastChangeTime != nil {
			if err := o.LastChangeTime.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetLastChangeTimeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pftMDLastChangeTime {out} (1:{alias=PFILETIME,pointer=ref}*(1))(2:{alias=FILETIME}(struct))
	{
		if o.LastChangeTime == nil {
			o.LastChangeTime = &dtyp.Filetime{}
		}
		if err := o.LastChangeTime.UnmarshalNDR(ctx, w); err != nil {
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

// GetLastChangeTimeRequest structure represents the GetLastChangeTime operation request
type GetLastChangeTimeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// hMDHandle: An unsigned 32-bit integer value containing an open metabase handle specifying
	// the key to be queried.
	Handle uint32 `idl:"name:hMDHandle" json:"handle"`
	// pszMDPath: A pointer to a Unicode string containing the path of the node to be queried,
	// relative to the path of the hMDHandle parameter.
	Path string `idl:"name:pszMDPath;string;pointer:unique" json:"path"`
	// bLocalTime: A Boolean value indicating whether the time value returned in the pftMDLastChangeTime
	// parameter is specified as local time (TRUE) or UTC time (FALSE).
	LocalTime bool `idl:"name:bLocalTime" json:"local_time"`
}

func (o *GetLastChangeTimeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetLastChangeTimeOperation) *xxx_GetLastChangeTimeOperation {
	if op == nil {
		op = &xxx_GetLastChangeTimeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Handle = o.Handle
	op.Path = o.Path
	op.LocalTime = o.LocalTime
	return op
}

func (o *GetLastChangeTimeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetLastChangeTimeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Handle = op.Handle
	o.Path = op.Path
	o.LocalTime = op.LocalTime
}
func (o *GetLastChangeTimeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetLastChangeTimeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLastChangeTimeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetLastChangeTimeResponse structure represents the GetLastChangeTime operation response
type GetLastChangeTimeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pftMDLastChangeTime: A pointer to a FILETIME structure that returns the last change
	// time for the node.
	LastChangeTime *dtyp.Filetime `idl:"name:pftMDLastChangeTime" json:"last_change_time"`
	// Return: The GetLastChangeTime return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetLastChangeTimeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetLastChangeTimeOperation) *xxx_GetLastChangeTimeOperation {
	if op == nil {
		op = &xxx_GetLastChangeTimeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.LastChangeTime = o.LastChangeTime
	op.Return = o.Return
	return op
}

func (o *GetLastChangeTimeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetLastChangeTimeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.LastChangeTime = op.LastChangeTime
	o.Return = op.Return
}
func (o *GetLastChangeTimeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetLastChangeTimeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLastChangeTimeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_KeyExchangePhase1Operation structure represents the R_KeyExchangePhase1 operation
type xxx_KeyExchangePhase1Operation struct {
	This                     *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That                     *dcom.ORPCThat      `idl:"name:That" json:"that"`
	ClientKeyExchangeKeyBlob *imsa.IISCryptoBlob `idl:"name:pClientKeyExchangeKeyBlob;pointer:unique" json:"client_key_exchange_key_blob"`
	ClientSignatureKeyBlob   *imsa.IISCryptoBlob `idl:"name:pClientSignatureKeyBlob;pointer:unique" json:"client_signature_key_blob"`
	ServerKeyExchangeKeyBlob *imsa.IISCryptoBlob `idl:"name:ppServerKeyExchangeKeyBlob" json:"server_key_exchange_key_blob"`
	ServerSignatureKeyBlob   *imsa.IISCryptoBlob `idl:"name:ppServerSignatureKeyBlob" json:"server_signature_key_blob"`
	ServerSessionKeyBlob     *imsa.IISCryptoBlob `idl:"name:ppServerSessionKeyBlob" json:"server_session_key_blob"`
	Return                   int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_KeyExchangePhase1Operation) OpNum() int { return 26 }

func (o *xxx_KeyExchangePhase1Operation) OpName() string {
	return "/IMSAdminBaseW/v0/R_KeyExchangePhase1"
}

func (o *xxx_KeyExchangePhase1Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_KeyExchangePhase1Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pClientKeyExchangeKeyBlob {in} (1:{pointer=unique}*(1))(2:{alias=IIS_CRYPTO_BLOB}(struct))
	{
		if o.ClientKeyExchangeKeyBlob != nil {
			_ptr_pClientKeyExchangeKeyBlob := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ClientKeyExchangeKeyBlob != nil {
					if err := o.ClientKeyExchangeKeyBlob.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&imsa.IISCryptoBlob{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ClientKeyExchangeKeyBlob, _ptr_pClientKeyExchangeKeyBlob); err != nil {
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
	// pClientSignatureKeyBlob {in} (1:{pointer=unique}*(1))(2:{alias=IIS_CRYPTO_BLOB}(struct))
	{
		if o.ClientSignatureKeyBlob != nil {
			_ptr_pClientSignatureKeyBlob := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ClientSignatureKeyBlob != nil {
					if err := o.ClientSignatureKeyBlob.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&imsa.IISCryptoBlob{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ClientSignatureKeyBlob, _ptr_pClientSignatureKeyBlob); err != nil {
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

func (o *xxx_KeyExchangePhase1Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pClientKeyExchangeKeyBlob {in} (1:{pointer=unique}*(1))(2:{alias=IIS_CRYPTO_BLOB}(struct))
	{
		_ptr_pClientKeyExchangeKeyBlob := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ClientKeyExchangeKeyBlob == nil {
				o.ClientKeyExchangeKeyBlob = &imsa.IISCryptoBlob{}
			}
			if err := o.ClientKeyExchangeKeyBlob.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pClientKeyExchangeKeyBlob := func(ptr interface{}) { o.ClientKeyExchangeKeyBlob = *ptr.(**imsa.IISCryptoBlob) }
		if err := w.ReadPointer(&o.ClientKeyExchangeKeyBlob, _s_pClientKeyExchangeKeyBlob, _ptr_pClientKeyExchangeKeyBlob); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pClientSignatureKeyBlob {in} (1:{pointer=unique}*(1))(2:{alias=IIS_CRYPTO_BLOB}(struct))
	{
		_ptr_pClientSignatureKeyBlob := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ClientSignatureKeyBlob == nil {
				o.ClientSignatureKeyBlob = &imsa.IISCryptoBlob{}
			}
			if err := o.ClientSignatureKeyBlob.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pClientSignatureKeyBlob := func(ptr interface{}) { o.ClientSignatureKeyBlob = *ptr.(**imsa.IISCryptoBlob) }
		if err := w.ReadPointer(&o.ClientSignatureKeyBlob, _s_pClientSignatureKeyBlob, _ptr_pClientSignatureKeyBlob); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_KeyExchangePhase1Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_KeyExchangePhase1Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppServerKeyExchangeKeyBlob {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IIS_CRYPTO_BLOB}(struct))
	{
		if o.ServerKeyExchangeKeyBlob != nil {
			_ptr_ppServerKeyExchangeKeyBlob := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ServerKeyExchangeKeyBlob != nil {
					if err := o.ServerKeyExchangeKeyBlob.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&imsa.IISCryptoBlob{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerKeyExchangeKeyBlob, _ptr_ppServerKeyExchangeKeyBlob); err != nil {
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
	// ppServerSignatureKeyBlob {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IIS_CRYPTO_BLOB}(struct))
	{
		if o.ServerSignatureKeyBlob != nil {
			_ptr_ppServerSignatureKeyBlob := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ServerSignatureKeyBlob != nil {
					if err := o.ServerSignatureKeyBlob.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&imsa.IISCryptoBlob{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerSignatureKeyBlob, _ptr_ppServerSignatureKeyBlob); err != nil {
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
	// ppServerSessionKeyBlob {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IIS_CRYPTO_BLOB}(struct))
	{
		if o.ServerSessionKeyBlob != nil {
			_ptr_ppServerSessionKeyBlob := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ServerSessionKeyBlob != nil {
					if err := o.ServerSessionKeyBlob.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&imsa.IISCryptoBlob{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerSessionKeyBlob, _ptr_ppServerSessionKeyBlob); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_KeyExchangePhase1Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppServerKeyExchangeKeyBlob {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IIS_CRYPTO_BLOB}(struct))
	{
		_ptr_ppServerKeyExchangeKeyBlob := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ServerKeyExchangeKeyBlob == nil {
				o.ServerKeyExchangeKeyBlob = &imsa.IISCryptoBlob{}
			}
			if err := o.ServerKeyExchangeKeyBlob.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppServerKeyExchangeKeyBlob := func(ptr interface{}) { o.ServerKeyExchangeKeyBlob = *ptr.(**imsa.IISCryptoBlob) }
		if err := w.ReadPointer(&o.ServerKeyExchangeKeyBlob, _s_ppServerKeyExchangeKeyBlob, _ptr_ppServerKeyExchangeKeyBlob); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppServerSignatureKeyBlob {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IIS_CRYPTO_BLOB}(struct))
	{
		_ptr_ppServerSignatureKeyBlob := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ServerSignatureKeyBlob == nil {
				o.ServerSignatureKeyBlob = &imsa.IISCryptoBlob{}
			}
			if err := o.ServerSignatureKeyBlob.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppServerSignatureKeyBlob := func(ptr interface{}) { o.ServerSignatureKeyBlob = *ptr.(**imsa.IISCryptoBlob) }
		if err := w.ReadPointer(&o.ServerSignatureKeyBlob, _s_ppServerSignatureKeyBlob, _ptr_ppServerSignatureKeyBlob); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppServerSessionKeyBlob {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IIS_CRYPTO_BLOB}(struct))
	{
		_ptr_ppServerSessionKeyBlob := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ServerSessionKeyBlob == nil {
				o.ServerSessionKeyBlob = &imsa.IISCryptoBlob{}
			}
			if err := o.ServerSessionKeyBlob.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppServerSessionKeyBlob := func(ptr interface{}) { o.ServerSessionKeyBlob = *ptr.(**imsa.IISCryptoBlob) }
		if err := w.ReadPointer(&o.ServerSessionKeyBlob, _s_ppServerSessionKeyBlob, _ptr_ppServerSessionKeyBlob); err != nil {
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

// KeyExchangePhase1Request structure represents the R_KeyExchangePhase1 operation request
type KeyExchangePhase1Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pClientKeyExchangeKeyBlob: A pointer to an IIS_CRYPTO_BLOB structure containing the
	// encrypted client key used to decrypt client data.
	ClientKeyExchangeKeyBlob *imsa.IISCryptoBlob `idl:"name:pClientKeyExchangeKeyBlob;pointer:unique" json:"client_key_exchange_key_blob"`
	// pClientSignatureKeyBlob: A pointer to an IIS_CRYPTO_BLOB structure containing the
	// encrypted client signature key used for data verification.
	ClientSignatureKeyBlob *imsa.IISCryptoBlob `idl:"name:pClientSignatureKeyBlob;pointer:unique" json:"client_signature_key_blob"`
}

func (o *KeyExchangePhase1Request) xxx_ToOp(ctx context.Context, op *xxx_KeyExchangePhase1Operation) *xxx_KeyExchangePhase1Operation {
	if op == nil {
		op = &xxx_KeyExchangePhase1Operation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ClientKeyExchangeKeyBlob = o.ClientKeyExchangeKeyBlob
	op.ClientSignatureKeyBlob = o.ClientSignatureKeyBlob
	return op
}

func (o *KeyExchangePhase1Request) xxx_FromOp(ctx context.Context, op *xxx_KeyExchangePhase1Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ClientKeyExchangeKeyBlob = op.ClientKeyExchangeKeyBlob
	o.ClientSignatureKeyBlob = op.ClientSignatureKeyBlob
}
func (o *KeyExchangePhase1Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *KeyExchangePhase1Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_KeyExchangePhase1Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// KeyExchangePhase1Response structure represents the R_KeyExchangePhase1 operation response
type KeyExchangePhase1Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppServerKeyExchangeKeyBlob: A pointer to a set of IIS_CRYPTO_BLOB structures containing
	// encrypted server keys used by the client to decrypt server data.
	ServerKeyExchangeKeyBlob *imsa.IISCryptoBlob `idl:"name:ppServerKeyExchangeKeyBlob" json:"server_key_exchange_key_blob"`
	// ppServerSignatureKeyBlob: A pointer to a set of IIS_CRYPTO_BLOB structures containing
	// encrypted server signature keys used for data verification.
	ServerSignatureKeyBlob *imsa.IISCryptoBlob `idl:"name:ppServerSignatureKeyBlob" json:"server_signature_key_blob"`
	// ppServerSessionKeyBlob: A pointer to a set of IIS_CRYPTO_BLOB structures containing
	// encrypted server session keys.
	ServerSessionKeyBlob *imsa.IISCryptoBlob `idl:"name:ppServerSessionKeyBlob" json:"server_session_key_blob"`
	// Return: The R_KeyExchangePhase1 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *KeyExchangePhase1Response) xxx_ToOp(ctx context.Context, op *xxx_KeyExchangePhase1Operation) *xxx_KeyExchangePhase1Operation {
	if op == nil {
		op = &xxx_KeyExchangePhase1Operation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ServerKeyExchangeKeyBlob = o.ServerKeyExchangeKeyBlob
	op.ServerSignatureKeyBlob = o.ServerSignatureKeyBlob
	op.ServerSessionKeyBlob = o.ServerSessionKeyBlob
	op.Return = o.Return
	return op
}

func (o *KeyExchangePhase1Response) xxx_FromOp(ctx context.Context, op *xxx_KeyExchangePhase1Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ServerKeyExchangeKeyBlob = op.ServerKeyExchangeKeyBlob
	o.ServerSignatureKeyBlob = op.ServerSignatureKeyBlob
	o.ServerSessionKeyBlob = op.ServerSessionKeyBlob
	o.Return = op.Return
}
func (o *KeyExchangePhase1Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *KeyExchangePhase1Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_KeyExchangePhase1Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_KeyExchangePhase2Operation structure represents the R_KeyExchangePhase2 operation
type xxx_KeyExchangePhase2Operation struct {
	This                 *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat      `idl:"name:That" json:"that"`
	ClientSessionKeyBlob *imsa.IISCryptoBlob `idl:"name:pClientSessionKeyBlob;pointer:unique" json:"client_session_key_blob"`
	ClientHashBlob       *imsa.IISCryptoBlob `idl:"name:pClientHashBlob;pointer:unique" json:"client_hash_blob"`
	ServerHashBlob       *imsa.IISCryptoBlob `idl:"name:ppServerHashBlob" json:"server_hash_blob"`
	Return               int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_KeyExchangePhase2Operation) OpNum() int { return 27 }

func (o *xxx_KeyExchangePhase2Operation) OpName() string {
	return "/IMSAdminBaseW/v0/R_KeyExchangePhase2"
}

func (o *xxx_KeyExchangePhase2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_KeyExchangePhase2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pClientSessionKeyBlob {in} (1:{pointer=unique}*(1))(2:{alias=IIS_CRYPTO_BLOB}(struct))
	{
		if o.ClientSessionKeyBlob != nil {
			_ptr_pClientSessionKeyBlob := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ClientSessionKeyBlob != nil {
					if err := o.ClientSessionKeyBlob.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&imsa.IISCryptoBlob{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ClientSessionKeyBlob, _ptr_pClientSessionKeyBlob); err != nil {
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
	// pClientHashBlob {in} (1:{pointer=unique}*(1))(2:{alias=IIS_CRYPTO_BLOB}(struct))
	{
		if o.ClientHashBlob != nil {
			_ptr_pClientHashBlob := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ClientHashBlob != nil {
					if err := o.ClientHashBlob.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&imsa.IISCryptoBlob{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ClientHashBlob, _ptr_pClientHashBlob); err != nil {
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

func (o *xxx_KeyExchangePhase2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pClientSessionKeyBlob {in} (1:{pointer=unique}*(1))(2:{alias=IIS_CRYPTO_BLOB}(struct))
	{
		_ptr_pClientSessionKeyBlob := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ClientSessionKeyBlob == nil {
				o.ClientSessionKeyBlob = &imsa.IISCryptoBlob{}
			}
			if err := o.ClientSessionKeyBlob.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pClientSessionKeyBlob := func(ptr interface{}) { o.ClientSessionKeyBlob = *ptr.(**imsa.IISCryptoBlob) }
		if err := w.ReadPointer(&o.ClientSessionKeyBlob, _s_pClientSessionKeyBlob, _ptr_pClientSessionKeyBlob); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pClientHashBlob {in} (1:{pointer=unique}*(1))(2:{alias=IIS_CRYPTO_BLOB}(struct))
	{
		_ptr_pClientHashBlob := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ClientHashBlob == nil {
				o.ClientHashBlob = &imsa.IISCryptoBlob{}
			}
			if err := o.ClientHashBlob.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pClientHashBlob := func(ptr interface{}) { o.ClientHashBlob = *ptr.(**imsa.IISCryptoBlob) }
		if err := w.ReadPointer(&o.ClientHashBlob, _s_pClientHashBlob, _ptr_pClientHashBlob); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_KeyExchangePhase2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_KeyExchangePhase2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppServerHashBlob {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IIS_CRYPTO_BLOB}(struct))
	{
		if o.ServerHashBlob != nil {
			_ptr_ppServerHashBlob := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ServerHashBlob != nil {
					if err := o.ServerHashBlob.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&imsa.IISCryptoBlob{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerHashBlob, _ptr_ppServerHashBlob); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_KeyExchangePhase2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppServerHashBlob {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IIS_CRYPTO_BLOB}(struct))
	{
		_ptr_ppServerHashBlob := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ServerHashBlob == nil {
				o.ServerHashBlob = &imsa.IISCryptoBlob{}
			}
			if err := o.ServerHashBlob.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppServerHashBlob := func(ptr interface{}) { o.ServerHashBlob = *ptr.(**imsa.IISCryptoBlob) }
		if err := w.ReadPointer(&o.ServerHashBlob, _s_ppServerHashBlob, _ptr_ppServerHashBlob); err != nil {
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

// KeyExchangePhase2Request structure represents the R_KeyExchangePhase2 operation request
type KeyExchangePhase2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pClientSessionKeyBlob: A pointer to an IIS_CRYPTO_BLOB structure containing the encrypted
	// client session key.
	ClientSessionKeyBlob *imsa.IISCryptoBlob `idl:"name:pClientSessionKeyBlob;pointer:unique" json:"client_session_key_blob"`
	// pClientHashBlob: A pointer to an IIS_CRYPTO_BLOB structure containing the encrypted
	// client hash key.
	ClientHashBlob *imsa.IISCryptoBlob `idl:"name:pClientHashBlob;pointer:unique" json:"client_hash_blob"`
}

func (o *KeyExchangePhase2Request) xxx_ToOp(ctx context.Context, op *xxx_KeyExchangePhase2Operation) *xxx_KeyExchangePhase2Operation {
	if op == nil {
		op = &xxx_KeyExchangePhase2Operation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ClientSessionKeyBlob = o.ClientSessionKeyBlob
	op.ClientHashBlob = o.ClientHashBlob
	return op
}

func (o *KeyExchangePhase2Request) xxx_FromOp(ctx context.Context, op *xxx_KeyExchangePhase2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ClientSessionKeyBlob = op.ClientSessionKeyBlob
	o.ClientHashBlob = op.ClientHashBlob
}
func (o *KeyExchangePhase2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *KeyExchangePhase2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_KeyExchangePhase2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// KeyExchangePhase2Response structure represents the R_KeyExchangePhase2 operation response
type KeyExchangePhase2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppServerHashBlob: A pointer to a set of IIS_CRYPTO_BLOB structures containing the
	// encrypted session hash keys.
	ServerHashBlob *imsa.IISCryptoBlob `idl:"name:ppServerHashBlob" json:"server_hash_blob"`
	// Return: The R_KeyExchangePhase2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *KeyExchangePhase2Response) xxx_ToOp(ctx context.Context, op *xxx_KeyExchangePhase2Operation) *xxx_KeyExchangePhase2Operation {
	if op == nil {
		op = &xxx_KeyExchangePhase2Operation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ServerHashBlob = o.ServerHashBlob
	op.Return = o.Return
	return op
}

func (o *KeyExchangePhase2Response) xxx_FromOp(ctx context.Context, op *xxx_KeyExchangePhase2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ServerHashBlob = op.ServerHashBlob
	o.Return = op.Return
}
func (o *KeyExchangePhase2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *KeyExchangePhase2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_KeyExchangePhase2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BackupOperation structure represents the Backup operation
type xxx_BackupOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	BackupName string         `idl:"name:pszMDBackupName;string;pointer:unique" json:"backup_name"`
	Version    uint32         `idl:"name:dwMDVersion" json:"version"`
	Flags      uint32         `idl:"name:dwMDFlags" json:"flags"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_BackupOperation) OpNum() int { return 28 }

func (o *xxx_BackupOperation) OpName() string { return "/IMSAdminBaseW/v0/Backup" }

func (o *xxx_BackupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pszMDBackupName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.BackupName != "" {
			_ptr_pszMDBackupName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.BackupName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.BackupName, _ptr_pszMDBackupName); err != nil {
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
	// dwMDVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Version); err != nil {
			return err
		}
	}
	// dwMDFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pszMDBackupName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszMDBackupName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.BackupName); err != nil {
				return err
			}
			return nil
		})
		_s_pszMDBackupName := func(ptr interface{}) { o.BackupName = *ptr.(*string) }
		if err := w.ReadPointer(&o.BackupName, _s_pszMDBackupName, _ptr_pszMDBackupName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwMDVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Version); err != nil {
			return err
		}
	}
	// dwMDFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_BackupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// BackupRequest structure represents the Backup operation request
type BackupRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pszMDBackupName: A string of up to 100 Unicode characters that names the backup.
	BackupName string `idl:"name:pszMDBackupName;string;pointer:unique" json:"backup_name"`
	// dwMDVersion: An integer value specifying the version number to be used for the backup.
	//
	//	+--------------------------------------+-------------------------------------------------------------------------------+
	//	|                                      |                                                                               |
	//	|                VALUE                 |                                    MEANING                                    |
	//	|                                      |                                                                               |
	//	+--------------------------------------+-------------------------------------------------------------------------------+
	//	+--------------------------------------+-------------------------------------------------------------------------------+
	//	| MD_BACKUP_HIGHEST_VERSION 0xFFFFFFFE | Overwrite the highest existing backup version with the specified backup name. |
	//	+--------------------------------------+-------------------------------------------------------------------------------+
	//	| MD_BACKUP_NEXT_VERSION 0xFFFFFFFF    | Use the next backup version number available with the specified backup name.  |
	//	+--------------------------------------+-------------------------------------------------------------------------------+
	Version uint32 `idl:"name:dwMDVersion" json:"version"`
	// dwMDFlags: An integer value containing the bit flags describing the type of backup
	// operation to be performed. The flags can be one or more of the following values.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|                                   |                                                                                  |
	//	|               VALUE               |                                     MEANING                                      |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| MD_BACKUP_FORCE_BACKUP 0x00000004 | Force the backup even if the SaveData operation specified by                     |
	//	|                                   | MD_BACKUP_SAVE_FIRST fails.                                                      |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| MD_BACKUP_OVERWRITE 0x00000001    | Back up even if a backup of the same name and version exists, overwriting it if  |
	//	|                                   | necessary.                                                                       |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| MD_BACKUP_SAVE_FIRST 0x00000002   | Perform a SaveData operation before the backup.                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:dwMDFlags" json:"flags"`
}

func (o *BackupRequest) xxx_ToOp(ctx context.Context, op *xxx_BackupOperation) *xxx_BackupOperation {
	if op == nil {
		op = &xxx_BackupOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.BackupName = o.BackupName
	op.Version = o.Version
	op.Flags = o.Flags
	return op
}

func (o *BackupRequest) xxx_FromOp(ctx context.Context, op *xxx_BackupOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.BackupName = op.BackupName
	o.Version = op.Version
	o.Flags = op.Flags
}
func (o *BackupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BackupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BackupResponse structure represents the Backup operation response
type BackupResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Backup return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *BackupResponse) xxx_ToOp(ctx context.Context, op *xxx_BackupOperation) *xxx_BackupOperation {
	if op == nil {
		op = &xxx_BackupOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *BackupResponse) xxx_FromOp(ctx context.Context, op *xxx_BackupOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *BackupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BackupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RestoreOperation structure represents the Restore operation
type xxx_RestoreOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	BackupName string         `idl:"name:pszMDBackupName;string;pointer:unique" json:"backup_name"`
	Version    uint32         `idl:"name:dwMDVersion" json:"version"`
	Flags      uint32         `idl:"name:dwMDFlags" json:"flags"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RestoreOperation) OpNum() int { return 29 }

func (o *xxx_RestoreOperation) OpName() string { return "/IMSAdminBaseW/v0/Restore" }

func (o *xxx_RestoreOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestoreOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pszMDBackupName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.BackupName != "" {
			_ptr_pszMDBackupName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.BackupName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.BackupName, _ptr_pszMDBackupName); err != nil {
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
	// dwMDVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Version); err != nil {
			return err
		}
	}
	// dwMDFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestoreOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pszMDBackupName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszMDBackupName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.BackupName); err != nil {
				return err
			}
			return nil
		})
		_s_pszMDBackupName := func(ptr interface{}) { o.BackupName = *ptr.(*string) }
		if err := w.ReadPointer(&o.BackupName, _s_pszMDBackupName, _ptr_pszMDBackupName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwMDVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Version); err != nil {
			return err
		}
	}
	// dwMDFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestoreOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestoreOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RestoreOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RestoreRequest structure represents the Restore operation request
type RestoreRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pszMDBackupName: A string of up to 100 Unicode characters that identifies the backup
	// to be restored.
	BackupName string `idl:"name:pszMDBackupName;string;pointer:unique" json:"backup_name"`
	// dwMDVersion: An integer value specifying either the version number of the backup
	// to be restored or the following constant.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                                      |                                                                                  |
	//	|                VALUE                 |                                     MEANING                                      |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MD_BACKUP_HIGHEST_VERSION 0xFFFFFFFE | Restore from the highest existing backup version in the specified backup         |
	//	|                                      | location.                                                                        |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	Version uint32 `idl:"name:dwMDVersion" json:"version"`
	// dwMDFlags: This parameter is reserved and SHOULD always be set to zero.
	Flags uint32 `idl:"name:dwMDFlags" json:"flags"`
}

func (o *RestoreRequest) xxx_ToOp(ctx context.Context, op *xxx_RestoreOperation) *xxx_RestoreOperation {
	if op == nil {
		op = &xxx_RestoreOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.BackupName = o.BackupName
	op.Version = o.Version
	op.Flags = o.Flags
	return op
}

func (o *RestoreRequest) xxx_FromOp(ctx context.Context, op *xxx_RestoreOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.BackupName = op.BackupName
	o.Version = op.Version
	o.Flags = op.Flags
}
func (o *RestoreRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RestoreRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RestoreOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RestoreResponse structure represents the Restore operation response
type RestoreResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Restore return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RestoreResponse) xxx_ToOp(ctx context.Context, op *xxx_RestoreOperation) *xxx_RestoreOperation {
	if op == nil {
		op = &xxx_RestoreOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RestoreResponse) xxx_FromOp(ctx context.Context, op *xxx_RestoreOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RestoreResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RestoreResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RestoreOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumBackupsOperation structure represents the EnumBackups operation
type xxx_EnumBackupsOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	BackupName string         `idl:"name:pszMDBackupName;size_is:(100)" json:"backup_name"`
	Version    uint32         `idl:"name:pdwMDVersion" json:"version"`
	BackupTime *dtyp.Filetime `idl:"name:pftMDBackupTime" json:"backup_time"`
	EnumIndex  uint32         `idl:"name:dwMDEnumIndex" json:"enum_index"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumBackupsOperation) OpNum() int { return 30 }

func (o *xxx_EnumBackupsOperation) OpName() string { return "/IMSAdminBaseW/v0/EnumBackups" }

func (o *xxx_EnumBackupsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	// cannot evaluate expression 100
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumBackupsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pszMDBackupName {in, out} (1:{alias=LPWSTR}*(1)[dim:0,size_is=100,string](wchar))
	{
		dimSize1 := uint64(100)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		_BackupName_buf := utf16.Encode([]rune(o.BackupName))
		if uint64(len(_BackupName_buf)) > sizeInfo[0] {
			_BackupName_buf = _BackupName_buf[:sizeInfo[0]]
		}
		for i1 := range _BackupName_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_BackupName_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_BackupName_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	// dwMDEnumIndex {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.EnumIndex); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumBackupsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pszMDBackupName {in, out} (1:{alias=LPWSTR,pointer=ref}*(1)[dim:0,size_is=100,string](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _BackupName_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _BackupName_buf", sizeInfo[0])
		}
		_BackupName_buf = make([]uint16, sizeInfo[0])
		for i1 := range _BackupName_buf {
			i1 := i1
			if err := w.ReadData(&_BackupName_buf[i1]); err != nil {
				return err
			}
		}
		o.BackupName = strings.TrimRight(string(utf16.Decode(_BackupName_buf)), ndr.ZeroString)
	}
	// dwMDEnumIndex {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.EnumIndex); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumBackupsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	// cannot evaluate expression 100
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumBackupsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pszMDBackupName {in, out} (1:{alias=LPWSTR}*(1)[dim:0,size_is=100,string](wchar))
	{
		dimSize1 := uint64(100)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		_BackupName_buf := utf16.Encode([]rune(o.BackupName))
		if uint64(len(_BackupName_buf)) > sizeInfo[0] {
			_BackupName_buf = _BackupName_buf[:sizeInfo[0]]
		}
		for i1 := range _BackupName_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_BackupName_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_BackupName_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	// pdwMDVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Version); err != nil {
			return err
		}
	}
	// pftMDBackupTime {out} (1:{alias=PFILETIME}*(1))(2:{alias=FILETIME}(struct))
	{
		if o.BackupTime != nil {
			if err := o.BackupTime.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_EnumBackupsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pszMDBackupName {in, out} (1:{alias=LPWSTR,pointer=ref}*(1)[dim:0,size_is=100,string](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _BackupName_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _BackupName_buf", sizeInfo[0])
		}
		_BackupName_buf = make([]uint16, sizeInfo[0])
		for i1 := range _BackupName_buf {
			i1 := i1
			if err := w.ReadData(&_BackupName_buf[i1]); err != nil {
				return err
			}
		}
		o.BackupName = strings.TrimRight(string(utf16.Decode(_BackupName_buf)), ndr.ZeroString)
	}
	// pdwMDVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Version); err != nil {
			return err
		}
	}
	// pftMDBackupTime {out} (1:{alias=PFILETIME,pointer=ref}*(1))(2:{alias=FILETIME}(struct))
	{
		if o.BackupTime == nil {
			o.BackupTime = &dtyp.Filetime{}
		}
		if err := o.BackupTime.UnmarshalNDR(ctx, w); err != nil {
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

// EnumBackupsRequest structure represents the EnumBackups operation request
type EnumBackupsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pszMDBackupName: A buffer of size MD_BACKUP_MAX_LEN. On input, the buffer can contain
	// either a string of Unicode characters that names the backup set to be enumerated
	// or an empty string.
	BackupName string `idl:"name:pszMDBackupName;size_is:(100)" json:"backup_name"`
	// dwMDEnumIndex: An integer value specifying the index number of the backup to be enumerated.
	EnumIndex uint32 `idl:"name:dwMDEnumIndex" json:"enum_index"`
}

func (o *EnumBackupsRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumBackupsOperation) *xxx_EnumBackupsOperation {
	if op == nil {
		op = &xxx_EnumBackupsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.BackupName = o.BackupName
	op.EnumIndex = o.EnumIndex
	return op
}

func (o *EnumBackupsRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumBackupsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.BackupName = op.BackupName
	o.EnumIndex = op.EnumIndex
}
func (o *EnumBackupsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumBackupsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumBackupsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumBackupsResponse structure represents the EnumBackups operation response
type EnumBackupsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pszMDBackupName: A buffer of size MD_BACKUP_MAX_LEN. On input, the buffer can contain
	// either a string of Unicode characters that names the backup set to be enumerated
	// or an empty string.
	BackupName string `idl:"name:pszMDBackupName;size_is:(100)" json:"backup_name"`
	// pdwMDVersion: An integer value containing the version number of the backup.
	Version uint32 `idl:"name:pdwMDVersion" json:"version"`
	// pftMDBackupTime: A FILETIME structure containing the Coordinated Universal Time (UTC)
	// date and time when this backup was created.
	BackupTime *dtyp.Filetime `idl:"name:pftMDBackupTime" json:"backup_time"`
	// Return: The EnumBackups return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumBackupsResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumBackupsOperation) *xxx_EnumBackupsOperation {
	if op == nil {
		op = &xxx_EnumBackupsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.BackupName = o.BackupName
	op.Version = o.Version
	op.BackupTime = o.BackupTime
	op.Return = o.Return
	return op
}

func (o *EnumBackupsResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumBackupsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.BackupName = op.BackupName
	o.Version = op.Version
	o.BackupTime = op.BackupTime
	o.Return = op.Return
}
func (o *EnumBackupsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumBackupsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumBackupsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteBackupOperation structure represents the DeleteBackup operation
type xxx_DeleteBackupOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	BackupName string         `idl:"name:pszMDBackupName;string;pointer:unique" json:"backup_name"`
	Version    uint32         `idl:"name:dwMDVersion" json:"version"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteBackupOperation) OpNum() int { return 31 }

func (o *xxx_DeleteBackupOperation) OpName() string { return "/IMSAdminBaseW/v0/DeleteBackup" }

func (o *xxx_DeleteBackupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteBackupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pszMDBackupName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.BackupName != "" {
			_ptr_pszMDBackupName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.BackupName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.BackupName, _ptr_pszMDBackupName); err != nil {
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
	// dwMDVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Version); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteBackupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pszMDBackupName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszMDBackupName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.BackupName); err != nil {
				return err
			}
			return nil
		})
		_s_pszMDBackupName := func(ptr interface{}) { o.BackupName = *ptr.(*string) }
		if err := w.ReadPointer(&o.BackupName, _s_pszMDBackupName, _ptr_pszMDBackupName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwMDVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Version); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteBackupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteBackupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteBackupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteBackupRequest structure represents the DeleteBackup operation request
type DeleteBackupRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pszMDBackupName: A string of up to 100 Unicode characters that names the backup.
	BackupName string `idl:"name:pszMDBackupName;string;pointer:unique" json:"backup_name"`
	// dwMDVersion: Either an integer value specifying the version number of the backup
	// to be deleted or the following constant.
	//
	//	+--------------------------------------+-------------------------------------------------------------+
	//	|                                      |                                                             |
	//	|                VALUE                 |                           MEANING                           |
	//	|                                      |                                                             |
	//	+--------------------------------------+-------------------------------------------------------------+
	//	+--------------------------------------+-------------------------------------------------------------+
	//	| MD_BACKUP_HIGHEST_VERSION 0xFFFFFFFE | Delete the existing backup with the highest version number. |
	//	+--------------------------------------+-------------------------------------------------------------+
	Version uint32 `idl:"name:dwMDVersion" json:"version"`
}

func (o *DeleteBackupRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteBackupOperation) *xxx_DeleteBackupOperation {
	if op == nil {
		op = &xxx_DeleteBackupOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.BackupName = o.BackupName
	op.Version = o.Version
	return op
}

func (o *DeleteBackupRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteBackupOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.BackupName = op.BackupName
	o.Version = op.Version
}
func (o *DeleteBackupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteBackupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteBackupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteBackupResponse structure represents the DeleteBackup operation response
type DeleteBackupResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteBackup return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteBackupResponse) xxx_ToOp(ctx context.Context, op *xxx_DeleteBackupOperation) *xxx_DeleteBackupOperation {
	if op == nil {
		op = &xxx_DeleteBackupOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DeleteBackupResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteBackupOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteBackupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteBackupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteBackupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_UnmarshalInterfaceOperation structure represents the UnmarshalInterface operation
type xxx_UnmarshalInterfaceOperation struct {
	This      *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat      `idl:"name:That" json:"that"`
	Interface *imsa.IMSAdminBaseW `idl:"name:piadmbwInterface" json:"interface"`
	Return    int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_UnmarshalInterfaceOperation) OpNum() int { return 32 }

func (o *xxx_UnmarshalInterfaceOperation) OpName() string {
	return "/IMSAdminBaseW/v0/UnmarshalInterface"
}

func (o *xxx_UnmarshalInterfaceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnmarshalInterfaceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_UnmarshalInterfaceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_UnmarshalInterfaceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnmarshalInterfaceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// piadmbwInterface {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSAdminBaseW}(interface))
	{
		if o.Interface != nil {
			_ptr_piadmbwInterface := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Interface != nil {
					if err := o.Interface.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&imsa.IMSAdminBaseW{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Interface, _ptr_piadmbwInterface); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnmarshalInterfaceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// piadmbwInterface {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSAdminBaseW}(interface))
	{
		_ptr_piadmbwInterface := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Interface == nil {
				o.Interface = &imsa.IMSAdminBaseW{}
			}
			if err := o.Interface.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_piadmbwInterface := func(ptr interface{}) { o.Interface = *ptr.(**imsa.IMSAdminBaseW) }
		if err := w.ReadPointer(&o.Interface, _s_piadmbwInterface, _ptr_piadmbwInterface); err != nil {
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

// UnmarshalInterfaceRequest structure represents the UnmarshalInterface operation request
type UnmarshalInterfaceRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *UnmarshalInterfaceRequest) xxx_ToOp(ctx context.Context, op *xxx_UnmarshalInterfaceOperation) *xxx_UnmarshalInterfaceOperation {
	if op == nil {
		op = &xxx_UnmarshalInterfaceOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *UnmarshalInterfaceRequest) xxx_FromOp(ctx context.Context, op *xxx_UnmarshalInterfaceOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *UnmarshalInterfaceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *UnmarshalInterfaceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UnmarshalInterfaceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// UnmarshalInterfaceResponse structure represents the UnmarshalInterface operation response
type UnmarshalInterfaceResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// piadmbwInterface: The address of the pointer that contains the pointer to IMSAdminBaseW.
	Interface *imsa.IMSAdminBaseW `idl:"name:piadmbwInterface" json:"interface"`
	// Return: The UnmarshalInterface return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *UnmarshalInterfaceResponse) xxx_ToOp(ctx context.Context, op *xxx_UnmarshalInterfaceOperation) *xxx_UnmarshalInterfaceOperation {
	if op == nil {
		op = &xxx_UnmarshalInterfaceOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Interface = o.Interface
	op.Return = o.Return
	return op
}

func (o *UnmarshalInterfaceResponse) xxx_FromOp(ctx context.Context, op *xxx_UnmarshalInterfaceOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Interface = op.Interface
	o.Return = op.Return
}
func (o *UnmarshalInterfaceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *UnmarshalInterfaceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UnmarshalInterfaceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetServerGUIDOperation structure represents the R_GetServerGuid operation
type xxx_GetServerGUIDOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	ServerGUID *dtyp.GUID     `idl:"name:pServerGuid" json:"server_guid"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetServerGUIDOperation) OpNum() int { return 33 }

func (o *xxx_GetServerGUIDOperation) OpName() string { return "/IMSAdminBaseW/v0/R_GetServerGuid" }

func (o *xxx_GetServerGUIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerGUIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetServerGUIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetServerGUIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerGUIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pServerGuid {out} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ServerGUID != nil {
			if err := o.ServerGUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetServerGUIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pServerGuid {out} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ServerGUID == nil {
			o.ServerGUID = &dtyp.GUID{}
		}
		if err := o.ServerGUID.UnmarshalNDR(ctx, w); err != nil {
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

// GetServerGUIDRequest structure represents the R_GetServerGuid operation request
type GetServerGUIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetServerGUIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetServerGUIDOperation) *xxx_GetServerGUIDOperation {
	if op == nil {
		op = &xxx_GetServerGUIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetServerGUIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetServerGUIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetServerGUIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetServerGUIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetServerGUIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetServerGUIDResponse structure represents the R_GetServerGuid operation response
type GetServerGUIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pServerGuid: A GUID uniquely identifying this DCOM object.
	ServerGUID *dtyp.GUID `idl:"name:pServerGuid" json:"server_guid"`
	// Return: The R_GetServerGuid return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetServerGUIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetServerGUIDOperation) *xxx_GetServerGUIDOperation {
	if op == nil {
		op = &xxx_GetServerGUIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ServerGUID = o.ServerGUID
	op.Return = o.Return
	return op
}

func (o *GetServerGUIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetServerGUIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ServerGUID = op.ServerGUID
	o.Return = op.Return
}
func (o *GetServerGUIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetServerGUIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetServerGUIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
