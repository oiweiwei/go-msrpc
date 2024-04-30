package ifsrmfilescreenmanager

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
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
	_ = idispatch.GoPackage
)

// IFsrmFileScreenManager server interface.
type FileScreenManagerServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// ActionVariables operation.
	GetActionVariables(context.Context, *GetActionVariablesRequest) (*GetActionVariablesResponse, error)

	// ActionVariableDescriptions operation.
	GetActionVariableDescriptions(context.Context, *GetActionVariableDescriptionsRequest) (*GetActionVariableDescriptionsResponse, error)

	// The CreateFileScreen method creates a Non-Persisted File Screen Instance (section
	// 3.2.1.3.1.2) on the specified path.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+----------------------------------+---------------------------------------------------------------------------------+
	//	|              RETURN              |                                                                                 |
	//	|            VALUE/CODE            |                                   DESCRIPTION                                   |
	//	|                                  |                                                                                 |
	//	+----------------------------------+---------------------------------------------------------------------------------+
	//	+----------------------------------+---------------------------------------------------------------------------------+
	//	| 0x80045304 FSRM_E_PATH_NOT_FOUND | The specified path could not be found.                                          |
	//	+----------------------------------+---------------------------------------------------------------------------------+
	//	| 0x80045306 FSRM_E_INVALID_PATH   | The content of the path parameter exceeds the maximum length of 260 characters. |
	//	+----------------------------------+---------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG          | The fileScreen parameter is NULL.                                               |
	//	+----------------------------------+---------------------------------------------------------------------------------+
	CreateFileScreen(context.Context, *CreateFileScreenRequest) (*CreateFileScreenResponse, error)

	// The GetFileScreen method returns the file screen from the List of Persisted File
	// Screens (section 3.2.1.3) for the specified path.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN              |                                                                                  |
	//	|            VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                  |                                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045304 FSRM_E_PATH_NOT_FOUND | The file screen for the specified path could not be found.                       |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045306 FSRM_E_INVALID_PATH   | The content of the path parameter exceeds the maximum length of 260 characters.  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG          | This code is returned for the following reasons: The path parameter is NULL. The |
	//	|                                  | fileScreen parameter is NULL.                                                    |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	GetFileScreen(context.Context, *GetFileScreenRequest) (*GetFileScreenResponse, error)

	// The EnumFileScreens method returns all the file screens from the List of Persisted
	// File Screens (section 3.2.1.3) that fall under the specified path.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND     | A file screen has not been applied to the specified directories.                 |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG         | The fileScreens parameter is NULL.                                               |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045311 FSRM_E_NOT_SUPPORTED | This options parameter contains invalid FsrmEnumOptions (section 2.2.1.2.5)      |
	//	|                                 | values.                                                                          |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	EnumFileScreens(context.Context, *EnumFileScreensRequest) (*EnumFileScreensResponse, error)

	// The CreateFileScreenException method creates a Non-Persisted File Screen Exception
	// Instance (section 3.2.1.3.2.2) on the specified path.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+--------------------------------+---------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                 |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                   |
	//	|                                |                                                                                 |
	//	+--------------------------------+---------------------------------------------------------------------------------+
	//	+--------------------------------+---------------------------------------------------------------------------------+
	//	| 0x80045306 FSRM_E_INVALID_PATH | The content of the path parameter exceeds the maximum length of 260 characters. |
	//	+--------------------------------+---------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG        | The fileScreenException parameter is NULL.                                      |
	//	+--------------------------------+---------------------------------------------------------------------------------+
	CreateFileScreenException(context.Context, *CreateFileScreenExceptionRequest) (*CreateFileScreenExceptionResponse, error)

	// The GetFileScreenException method returns the file screen exception from the List
	// of Persisted File Screen Exceptions (section 3.2.1.3) for the specified path.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN              |                                                                                  |
	//	|            VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                  |                                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND      | A file screen exception has not been applied to the specified directory.         |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045304 FSRM_E_PATH_NOT_FOUND | The file screen exception for the specified path could not be found.             |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045306 FSRM_E_INVALID_PATH   | The content of the path parameter exceeds the maximum length of 260 characters.  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG          | This code is returned for the following reasons: The path parameter is NULL. The |
	//	|                                  | fileScreenException parameter is NULL.                                           |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	GetFileScreenException(context.Context, *GetFileScreenExceptionRequest) (*GetFileScreenExceptionResponse, error)

	// The EnumFileScreenExceptions method returns all the file screen exceptions from the
	// List of Persisted File Screen Exceptions (section 3.2.1.3) that fall under the specified
	// path.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND     | A file screen exception has not been applied to the specified directory.         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG         | The fileScreenExceptions parameter is NULL.                                      |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045311 FSRM_E_NOT_SUPPORTED | This options parameter contains invalid FsrmEnumOptions (section 2.2.1.2.5)      |
	//	|                                 | values.                                                                          |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	EnumFileScreenExceptions(context.Context, *EnumFileScreenExceptionsRequest) (*EnumFileScreenExceptionsResponse, error)

	// The CreateFileScreenCollection method creates an empty collection. This creates a
	// location where callers can add file screens.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+-----------------------------------+
	//	|         RETURN          |                                   |
	//	|       VALUE/CODE        |            DESCRIPTION            |
	//	|                         |                                   |
	//	+-------------------------+-----------------------------------+
	//	+-------------------------+-----------------------------------+
	//	| 0x80070057 E_INVALIDARG | The collection parameter is NULL. |
	//	+-------------------------+-----------------------------------+
	CreateFileScreenCollection(context.Context, *CreateFileScreenCollectionRequest) (*CreateFileScreenCollectionResponse, error)
}

func RegisterFileScreenManagerServer(conn dcerpc.Conn, o FileScreenManagerServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewFileScreenManagerServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(FileScreenManagerSyntaxV0_0))...)
}

func NewFileScreenManagerServerHandle(o FileScreenManagerServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return FileScreenManagerServerHandle(ctx, o, opNum, r)
	}
}

func FileScreenManagerServerHandle(ctx context.Context, o FileScreenManagerServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // ActionVariables
		in := &GetActionVariablesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetActionVariables(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // ActionVariableDescriptions
		in := &GetActionVariableDescriptionsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetActionVariableDescriptions(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // CreateFileScreen
		in := &CreateFileScreenRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateFileScreen(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // GetFileScreen
		in := &GetFileScreenRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFileScreen(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // EnumFileScreens
		in := &EnumFileScreensRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumFileScreens(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // CreateFileScreenException
		in := &CreateFileScreenExceptionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateFileScreenException(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // GetFileScreenException
		in := &GetFileScreenExceptionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFileScreenException(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // EnumFileScreenExceptions
		in := &EnumFileScreenExceptionsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumFileScreenExceptions(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // CreateFileScreenCollection
		in := &CreateFileScreenCollectionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateFileScreenCollection(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
