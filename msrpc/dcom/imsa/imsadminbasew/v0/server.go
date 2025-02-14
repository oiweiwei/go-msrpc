package imsadminbasew

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

// IMSAdminBaseW server interface.
type IMSAdminBaseWServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

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
	AddKey(context.Context, *AddKeyRequest) (*AddKeyResponse, error)

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
	DeleteKey(context.Context, *DeleteKeyRequest) (*DeleteKeyResponse, error)

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
	DeleteChildKeys(context.Context, *DeleteChildKeysRequest) (*DeleteChildKeysResponse, error)

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
	EnumKeys(context.Context, *EnumKeysRequest) (*EnumKeysResponse, error)

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
	CopyKey(context.Context, *CopyKeyRequest) (*CopyKeyResponse, error)

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
	RenameKey(context.Context, *RenameKeyRequest) (*RenameKeyResponse, error)

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
	SetData(context.Context, *SetDataRequest) (*SetDataResponse, error)

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
	GetData(context.Context, *GetDataRequest) (*GetDataResponse, error)

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
	DeleteData(context.Context, *DeleteDataRequest) (*DeleteDataResponse, error)

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
	EnumData(context.Context, *EnumDataRequest) (*EnumDataResponse, error)

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
	GetAllData(context.Context, *GetAllDataRequest) (*GetAllDataResponse, error)

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
	DeleteAllData(context.Context, *DeleteAllDataRequest) (*DeleteAllDataResponse, error)

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
	CopyData(context.Context, *CopyDataRequest) (*CopyDataResponse, error)

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
	GetDataPaths(context.Context, *GetDataPathsRequest) (*GetDataPathsResponse, error)

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
	OpenKey(context.Context, *OpenKeyRequest) (*OpenKeyResponse, error)

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
	CloseKey(context.Context, *CloseKeyRequest) (*CloseKeyResponse, error)

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
	ChangePermissions(context.Context, *ChangePermissionsRequest) (*ChangePermissionsResponse, error)

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
	SaveData(context.Context, *SaveDataRequest) (*SaveDataResponse, error)

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
	GetHandleInfo(context.Context, *GetHandleInfoRequest) (*GetHandleInfoResponse, error)

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
	GetSystemChangeNumber(context.Context, *GetSystemChangeNumberRequest) (*GetSystemChangeNumberResponse, error)

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
	GetDataSetNumber(context.Context, *GetDataSetNumberRequest) (*GetDataSetNumberResponse, error)

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
	SetLastChangeTime(context.Context, *SetLastChangeTimeRequest) (*SetLastChangeTimeResponse, error)

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
	GetLastChangeTime(context.Context, *GetLastChangeTimeRequest) (*GetLastChangeTimeResponse, error)

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
	KeyExchangePhase1(context.Context, *KeyExchangePhase1Request) (*KeyExchangePhase1Response, error)

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
	KeyExchangePhase2(context.Context, *KeyExchangePhase2Request) (*KeyExchangePhase2Response, error)

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
	Backup(context.Context, *BackupRequest) (*BackupResponse, error)

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
	Restore(context.Context, *RestoreRequest) (*RestoreResponse, error)

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
	EnumBackups(context.Context, *EnumBackupsRequest) (*EnumBackupsResponse, error)

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
	DeleteBackup(context.Context, *DeleteBackupRequest) (*DeleteBackupResponse, error)

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
	UnmarshalInterface(context.Context, *UnmarshalInterfaceRequest) (*UnmarshalInterfaceResponse, error)

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
	GetServerGUID(context.Context, *GetServerGUIDRequest) (*GetServerGUIDResponse, error)
}

func RegisterIMSAdminBaseWServer(conn dcerpc.Conn, o IMSAdminBaseWServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewIMSAdminBaseWServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(IMSAdminBaseWSyntaxV0_0))...)
}

func NewIMSAdminBaseWServerHandle(o IMSAdminBaseWServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return IMSAdminBaseWServerHandle(ctx, o, opNum, r)
	}
}

func IMSAdminBaseWServerHandle(ctx context.Context, o IMSAdminBaseWServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // AddKey
		op := &xxx_AddKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // DeleteKey
		op := &xxx_DeleteKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // DeleteChildKeys
		op := &xxx_DeleteChildKeysOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteChildKeysRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteChildKeys(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // EnumKeys
		op := &xxx_EnumKeysOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumKeysRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumKeys(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // CopyKey
		op := &xxx_CopyKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CopyKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CopyKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // RenameKey
		op := &xxx_RenameKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RenameKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RenameKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // R_SetData
		op := &xxx_SetDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // R_GetData
		op := &xxx_GetDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // DeleteData
		op := &xxx_DeleteDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // R_EnumData
		op := &xxx_EnumDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // R_GetAllData
		op := &xxx_GetAllDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAllDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAllData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // DeleteAllData
		op := &xxx_DeleteAllDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteAllDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteAllData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // CopyData
		op := &xxx_CopyDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CopyDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CopyData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // GetDataPaths
		op := &xxx_GetDataPathsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDataPathsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDataPaths(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // OpenKey
		op := &xxx_OpenKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // CloseKey
		op := &xxx_CloseKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CloseKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // ChangePermissions
		op := &xxx_ChangePermissionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ChangePermissionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ChangePermissions(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // SaveData
		op := &xxx_SaveDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SaveDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SaveData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // GetHandleInfo
		op := &xxx_GetHandleInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetHandleInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetHandleInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // GetSystemChangeNumber
		op := &xxx_GetSystemChangeNumberOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSystemChangeNumberRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSystemChangeNumber(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // GetDataSetNumber
		op := &xxx_GetDataSetNumberOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDataSetNumberRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDataSetNumber(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // SetLastChangeTime
		op := &xxx_SetLastChangeTimeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetLastChangeTimeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetLastChangeTime(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // GetLastChangeTime
		op := &xxx_GetLastChangeTimeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLastChangeTimeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLastChangeTime(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // R_KeyExchangePhase1
		op := &xxx_KeyExchangePhase1Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &KeyExchangePhase1Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.KeyExchangePhase1(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // R_KeyExchangePhase2
		op := &xxx_KeyExchangePhase2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &KeyExchangePhase2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.KeyExchangePhase2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // Backup
		op := &xxx_BackupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BackupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Backup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // Restore
		op := &xxx_RestoreOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RestoreRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Restore(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // EnumBackups
		op := &xxx_EnumBackupsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumBackupsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumBackups(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 31: // DeleteBackup
		op := &xxx_DeleteBackupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteBackupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteBackup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 32: // UnmarshalInterface
		op := &xxx_UnmarshalInterfaceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UnmarshalInterfaceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.UnmarshalInterface(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // R_GetServerGuid
		op := &xxx_GetServerGUIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetServerGUIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetServerGUID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSAdminBaseW
type UnimplementedIMSAdminBaseWServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedIMSAdminBaseWServer) AddKey(context.Context, *AddKeyRequest) (*AddKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) DeleteKey(context.Context, *DeleteKeyRequest) (*DeleteKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) DeleteChildKeys(context.Context, *DeleteChildKeysRequest) (*DeleteChildKeysResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) EnumKeys(context.Context, *EnumKeysRequest) (*EnumKeysResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) CopyKey(context.Context, *CopyKeyRequest) (*CopyKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) RenameKey(context.Context, *RenameKeyRequest) (*RenameKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) SetData(context.Context, *SetDataRequest) (*SetDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) GetData(context.Context, *GetDataRequest) (*GetDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) DeleteData(context.Context, *DeleteDataRequest) (*DeleteDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) EnumData(context.Context, *EnumDataRequest) (*EnumDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) GetAllData(context.Context, *GetAllDataRequest) (*GetAllDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) DeleteAllData(context.Context, *DeleteAllDataRequest) (*DeleteAllDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) CopyData(context.Context, *CopyDataRequest) (*CopyDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) GetDataPaths(context.Context, *GetDataPathsRequest) (*GetDataPathsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) OpenKey(context.Context, *OpenKeyRequest) (*OpenKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) CloseKey(context.Context, *CloseKeyRequest) (*CloseKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) ChangePermissions(context.Context, *ChangePermissionsRequest) (*ChangePermissionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) SaveData(context.Context, *SaveDataRequest) (*SaveDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) GetHandleInfo(context.Context, *GetHandleInfoRequest) (*GetHandleInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) GetSystemChangeNumber(context.Context, *GetSystemChangeNumberRequest) (*GetSystemChangeNumberResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) GetDataSetNumber(context.Context, *GetDataSetNumberRequest) (*GetDataSetNumberResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) SetLastChangeTime(context.Context, *SetLastChangeTimeRequest) (*SetLastChangeTimeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) GetLastChangeTime(context.Context, *GetLastChangeTimeRequest) (*GetLastChangeTimeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) KeyExchangePhase1(context.Context, *KeyExchangePhase1Request) (*KeyExchangePhase1Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) KeyExchangePhase2(context.Context, *KeyExchangePhase2Request) (*KeyExchangePhase2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) Backup(context.Context, *BackupRequest) (*BackupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) Restore(context.Context, *RestoreRequest) (*RestoreResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) EnumBackups(context.Context, *EnumBackupsRequest) (*EnumBackupsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) DeleteBackup(context.Context, *DeleteBackupRequest) (*DeleteBackupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) UnmarshalInterface(context.Context, *UnmarshalInterfaceRequest) (*UnmarshalInterfaceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIMSAdminBaseWServer) GetServerGUID(context.Context, *GetServerGUIDRequest) (*GetServerGUIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ IMSAdminBaseWServer = (*UnimplementedIMSAdminBaseWServer)(nil)
