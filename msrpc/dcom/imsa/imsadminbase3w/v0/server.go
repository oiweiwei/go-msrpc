package imsadminbase3w

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	imsadminbase2w "github.com/oiweiwei/go-msrpc/msrpc/dcom/imsa/imsadminbase2w/v0"
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
	_ = imsadminbase2w.GoPackage
)

// IMSAdminBase3W server interface.
type IMSAdminBase3WServer interface {

	// IMSAdminBase2W base class.
	imsadminbase2w.IMSAdminBase2WServer

	// The GetChildPaths method returns all child nodes of a specified path from a supplied
	// metadata handle.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+--------------------------------------+----------------------------------------------------------+
	//	|                RETURN                |                                                          |
	//	|              VALUE/CODE              |                       DESCRIPTION                        |
	//	|                                      |                                                          |
	//	+--------------------------------------+----------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------+
	//	| 0x80070000 S_OK                      | The call was successful.                                 |
	//	+--------------------------------------+----------------------------------------------------------+
	//	| 0x80070003 ERROR_PATH_NOT_FOUND      | The system cannot find the path specified.               |
	//	+--------------------------------------+----------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG              | One or more arguments are invalid.                       |
	//	+--------------------------------------+----------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED       | Access is denied.                                        |
	//	+--------------------------------------+----------------------------------------------------------+
	//	| 0x80070008 ERROR_NOT_ENOUGH_MEMORY   | Not enough storage is available to process this command. |
	//	+--------------------------------------+----------------------------------------------------------+
	//	| 0x8007000E E_OUTOFMEMORY             | There was not enough memory to complete the method call. |
	//	+--------------------------------------+----------------------------------------------------------+
	//	| 0x8007007A ERROR_INSUFFICIENT_BUFFER | The data area passed to a system call is too small.      |
	//	+--------------------------------------+----------------------------------------------------------+
	//	| 0x800700A0 ERROR_BAD_ARGUMENTS       | One or more arguments are not correct.                   |
	//	+--------------------------------------+----------------------------------------------------------+
	//	| 0x80004005 E_FAIL                    | An unspecified error occurred.                           |
	//	+--------------------------------------+----------------------------------------------------------+
	//	| 0x80070006 E_HANDLE                  | An invalid handle was passed.                            |
	//	+--------------------------------------+----------------------------------------------------------+
	//	| 0x800CC800 MD_ERROR_NOT_INITIALIZED  | Metadata has not been initialized.                       |
	//	+--------------------------------------+----------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 40.
	GetChildPaths(context.Context, *GetChildPathsRequest) (*GetChildPathsResponse, error)
}

func RegisterIMSAdminBase3WServer(conn dcerpc.Conn, o IMSAdminBase3WServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewIMSAdminBase3WServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(IMSAdminBase3WSyntaxV0_0))...)
}

func NewIMSAdminBase3WServerHandle(o IMSAdminBase3WServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return IMSAdminBase3WServerHandle(ctx, o, opNum, r)
	}
}

func IMSAdminBase3WServerHandle(ctx context.Context, o IMSAdminBase3WServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 40 {
		// IMSAdminBase2W base method.
		return imsadminbase2w.IMSAdminBase2WServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 40: // GetChildPaths
		op := &xxx_GetChildPathsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetChildPathsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetChildPaths(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSAdminBase3W
type UnimplementedIMSAdminBase3WServer struct {
	imsadminbase2w.UnimplementedIMSAdminBase2WServer
}

func (UnimplementedIMSAdminBase3WServer) GetChildPaths(context.Context, *GetChildPathsRequest) (*GetChildPathsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ IMSAdminBase3WServer = (*UnimplementedIMSAdminBase3WServer)(nil)
