package imsadminbase2w

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	imsadminbasew "github.com/oiweiwei/go-msrpc/msrpc/dcom/imsa/imsadminbasew/v0"
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
	_ = imsadminbasew.GoPackage
)

// IMSAdminBase2W server interface.
type IMSAdminBase2WServer interface {

	// IMSAdminBaseW base class.
	imsadminbasew.IMSAdminBaseWServer

	// The BackupWithPasswd method backs up the metabase using a supplied password to encrypt
	// all secure data.
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
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 34.
	BackupWithPassword(context.Context, *BackupWithPasswordRequest) (*BackupWithPasswordResponse, error)

	// The RestoreWithPasswd method restores the metabase from a backup, using a supplied
	// password to decrypt the secure data.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                 | The call was successful.                                                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG         | One or more arguments are invalid.                                               |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007052B ERROR_WRONG_PASSWORD | Unable to update the password. The value provided as the current password is     |
	//	|                                 | incorrect.                                                                       |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 35.
	RestoreWithPassword(context.Context, *RestoreWithPasswordRequest) (*RestoreWithPasswordResponse, error)

	// The Export method exports a section of the metabase to a file.
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
	//	| 0x80070032 ERROR_NOT_SUPPORTED  | The request is not supported.              |
	//	+---------------------------------+--------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 36.
	Export(context.Context, *ExportRequest) (*ExportResponse, error)

	// The Import method imports metabase data from an exported file into the metabase.
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
	//	| 0x80070003 ERROR_PATH_NOT_FOUND | The system cannot find the file specified. |
	//	+---------------------------------+--------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 37.
	Import(context.Context, *ImportRequest) (*ImportResponse, error)

	// The RestoreHistory method restores a metabase history entry for a specific history
	// version.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+-------------------------------------+---------------------------------------------------------------+
	//	|               RETURN                |                                                               |
	//	|             VALUE/CODE              |                          DESCRIPTION                          |
	//	|                                     |                                                               |
	//	+-------------------------------------+---------------------------------------------------------------+
	//	+-------------------------------------+---------------------------------------------------------------+
	//	| 0x00000000 S_OK                     | The call was successful.                                      |
	//	+-------------------------------------+---------------------------------------------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND     | The system cannot find the file specified.                    |
	//	+-------------------------------------+---------------------------------------------------------------+
	//	| 0x80070003 ERROR_PATH_NOT_FOUND     | The system cannot find the path specified.                    |
	//	+-------------------------------------+---------------------------------------------------------------+
	//	| 0x8007000E E_OUTOFMEMORY            | Ran out of memory.                                            |
	//	+-------------------------------------+---------------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY  | Not enough storage is available to process this command.      |
	//	+-------------------------------------+---------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG             | One or more arguments are invalid.                            |
	//	+-------------------------------------+---------------------------------------------------------------+
	//	| 0x80070013 ERROR_INVALID_DATA       | One or more arguments are invalid.                            |
	//	+-------------------------------------+---------------------------------------------------------------+
	//	| 0x800703EC ERROR_INVALID_FLAGS      | Invalid flags were passed.                                    |
	//	+-------------------------------------+---------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED      | Access is denied.                                             |
	//	+-------------------------------------+---------------------------------------------------------------+
	//	| 0x800CC802 MD_ERROR_INVALID_VERSION | The version specified in metadata storage was not recognized. |
	//	+-------------------------------------+---------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 38.
	RestoreHistory(context.Context, *RestoreHistoryRequest) (*RestoreHistoryResponse, error)

	// The EnumHistory method returns an enumerated history entry with a supplied index.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                      | The call was successful.                                                         |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_PATH_NOT_FOUND      | The system cannot find the file specified.                                       |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY   | Not enough storage is available to process this command.                         |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000012 ERROR_NO_MORE_ITEMS       | There are no more history versions.                                              |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000007A ERROR_INSUFFICIENT_BUFFER | The data area passed to a system call is too small. In this case the location    |
	//	|                                      | string does not have enough space to return the path to the history location.    |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED       | Access is denied.                                                                |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG              | One or more arguments are invalid.                                               |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 39.
	EnumHistory(context.Context, *EnumHistoryRequest) (*EnumHistoryResponse, error)
}

func RegisterIMSAdminBase2WServer(conn dcerpc.Conn, o IMSAdminBase2WServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewIMSAdminBase2WServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(IMSAdminBase2WSyntaxV0_0))...)
}

func NewIMSAdminBase2WServerHandle(o IMSAdminBase2WServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return IMSAdminBase2WServerHandle(ctx, o, opNum, r)
	}
}

func IMSAdminBase2WServerHandle(ctx context.Context, o IMSAdminBase2WServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 34 {
		// IMSAdminBaseW base method.
		return imsadminbasew.IMSAdminBaseWServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 34: // BackupWithPasswd
		in := &BackupWithPasswordRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.BackupWithPassword(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 35: // RestoreWithPasswd
		in := &RestoreWithPasswordRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RestoreWithPassword(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 36: // Export
		in := &ExportRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Export(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 37: // Import
		in := &ImportRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Import(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 38: // RestoreHistory
		in := &RestoreHistoryRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RestoreHistory(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 39: // EnumHistory
		in := &EnumHistoryRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumHistory(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
