package iiisapplicationadmin

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

// IIISApplicationAdmin server interface.
type IISApplicationAdminServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The CreateApplication method creates a new application at the specified metabase
	// path.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+----------------------------+------------------------------------+
	//	|           RETURN           |                                    |
	//	|         VALUE/CODE         |            DESCRIPTION             |
	//	|                            |                                    |
	//	+----------------------------+------------------------------------+
	//	+----------------------------+------------------------------------+
	//	| 0x00000000 S_OK            | The call was successful.           |
	//	+----------------------------+------------------------------------+
	//	| 0x80070057 E_INVALIDARG    | One or more arguments are invalid. |
	//	+----------------------------+------------------------------------+
	//	| 0x80070490 ERROR_NOT_FOUND | Element not found.                 |
	//	+----------------------------+------------------------------------+
	//
	// The opnum field value for this method is 3.
	CreateApplication(context.Context, *CreateApplicationRequest) (*CreateApplicationResponse, error)

	// The DeleteApplication method deletes the application from the specified metabase
	// path.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
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
	//	| 0x80070057 E_INVALIDARG         | One or more arguments are invalid.         |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070003 ERROR_PATH_NOT_FOUND | The system cannot find the path specified. |
	//	+---------------------------------+--------------------------------------------+
	//
	// The opnum field value for this method is 4.
	DeleteApplication(context.Context, *DeleteApplicationRequest) (*DeleteApplicationResponse, error)

	// The CreateApplicationPool method creates a new application pool.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
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
	// The opnum field value for this method is 5.
	CreateApplicationPool(context.Context, *CreateApplicationPoolRequest) (*CreateApplicationPoolResponse, error)

	// The DeleteApplicationPool method deletes an application pool.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+----------------------------+----------------------------------------------------------------------------+
	//	|           RETURN           |                                                                            |
	//	|         VALUE/CODE         |                                DESCRIPTION                                 |
	//	|                            |                                                                            |
	//	+----------------------------+----------------------------------------------------------------------------+
	//	+----------------------------+----------------------------------------------------------------------------+
	//	| 0x00000000 S_OK            | The call was successful.                                                   |
	//	+----------------------------+----------------------------------------------------------------------------+
	//	| 0x80070490 ERROR_NOT_FOUND | Element not found.                                                         |
	//	+----------------------------+----------------------------------------------------------------------------+
	//	| 0x800710D3 ERROR_NOT_EMPTY | The library, drive, or media pool must be empty to perform this operation. |
	//	+----------------------------+----------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 6.
	DeleteApplicationPool(context.Context, *DeleteApplicationPoolRequest) (*DeleteApplicationPoolResponse, error)

	// The EnumerateApplicationsInPool method returns the metabase paths for the applications
	// associated with the application pool.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
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
	// The opnum field value for this method is 7.
	EnumerateApplicationsInPool(context.Context, *EnumerateApplicationsInPoolRequest) (*EnumerateApplicationsInPoolResponse, error)

	// The RecycleApplicationPool method restarts an application pool.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
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
	//
	// The opnum field value for this method is 8.
	RecycleApplicationPool(context.Context, *RecycleApplicationPoolRequest) (*RecycleApplicationPoolResponse, error)

	// The GetProcessMode method retrieves the application execution mode for the IIS server.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
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
	// The opnum field value for this method is 9.
	GetProcessMode(context.Context, *GetProcessModeRequest) (*GetProcessModeResponse, error)
}

func RegisterIISApplicationAdminServer(conn dcerpc.Conn, o IISApplicationAdminServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewIISApplicationAdminServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(IISApplicationAdminSyntaxV0_0))...)
}

func NewIISApplicationAdminServerHandle(o IISApplicationAdminServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return IISApplicationAdminServerHandle(ctx, o, opNum, r)
	}
}

func IISApplicationAdminServerHandle(ctx context.Context, o IISApplicationAdminServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // CreateApplication
		in := &CreateApplicationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateApplication(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // DeleteApplication
		in := &DeleteApplicationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteApplication(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // CreateApplicationPool
		in := &CreateApplicationPoolRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateApplicationPool(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // DeleteApplicationPool
		in := &DeleteApplicationPoolRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteApplicationPool(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // EnumerateApplicationsInPool
		in := &EnumerateApplicationsInPoolRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumerateApplicationsInPool(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // RecycleApplicationPool
		in := &RecycleApplicationPoolRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RecycleApplicationPool(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // GetProcessMode
		in := &GetProcessModeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetProcessMode(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
