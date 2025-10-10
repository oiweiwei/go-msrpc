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
		op := &xxx_GetActionVariablesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetActionVariablesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetActionVariables(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // ActionVariableDescriptions
		op := &xxx_GetActionVariableDescriptionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetActionVariableDescriptionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetActionVariableDescriptions(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // CreateFileScreen
		op := &xxx_CreateFileScreenOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateFileScreenRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateFileScreen(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // GetFileScreen
		op := &xxx_GetFileScreenOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFileScreenRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFileScreen(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // EnumFileScreens
		op := &xxx_EnumFileScreensOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumFileScreensRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumFileScreens(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // CreateFileScreenException
		op := &xxx_CreateFileScreenExceptionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateFileScreenExceptionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateFileScreenException(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // GetFileScreenException
		op := &xxx_GetFileScreenExceptionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFileScreenExceptionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFileScreenException(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // EnumFileScreenExceptions
		op := &xxx_EnumFileScreenExceptionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumFileScreenExceptionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumFileScreenExceptions(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // CreateFileScreenCollection
		op := &xxx_CreateFileScreenCollectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateFileScreenCollectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateFileScreenCollection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmFileScreenManager
type UnimplementedFileScreenManagerServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedFileScreenManagerServer) GetActionVariables(context.Context, *GetActionVariablesRequest) (*GetActionVariablesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileScreenManagerServer) GetActionVariableDescriptions(context.Context, *GetActionVariableDescriptionsRequest) (*GetActionVariableDescriptionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileScreenManagerServer) CreateFileScreen(context.Context, *CreateFileScreenRequest) (*CreateFileScreenResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileScreenManagerServer) GetFileScreen(context.Context, *GetFileScreenRequest) (*GetFileScreenResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileScreenManagerServer) EnumFileScreens(context.Context, *EnumFileScreensRequest) (*EnumFileScreensResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileScreenManagerServer) CreateFileScreenException(context.Context, *CreateFileScreenExceptionRequest) (*CreateFileScreenExceptionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileScreenManagerServer) GetFileScreenException(context.Context, *GetFileScreenExceptionRequest) (*GetFileScreenExceptionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileScreenManagerServer) EnumFileScreenExceptions(context.Context, *EnumFileScreenExceptionsRequest) (*EnumFileScreenExceptionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileScreenManagerServer) CreateFileScreenCollection(context.Context, *CreateFileScreenCollectionRequest) (*CreateFileScreenCollectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ FileScreenManagerServer = (*UnimplementedFileScreenManagerServer)(nil)
