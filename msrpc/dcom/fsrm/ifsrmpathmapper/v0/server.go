package ifsrmpathmapper

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

// IFsrmPathMapper server interface.
type PathMapperServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The GetSharePathsForLocalPath method returns all the network share paths that point
	// to the specified local path.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045306 FSRM_E_INVALID_PATH | The local path to return network shares does not exist or exceeds the maximum    |
	//	|                                | length of 260 characters.                                                        |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG        | This code is returned for the following reasons: The localPath parameter is      |
	//	|                                | empty or NULL. The sharePaths parameter is NULL.                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	GetSharePathsForLocalPath(context.Context, *GetSharePathsForLocalPathRequest) (*GetSharePathsForLocalPathResponse, error)
}

func RegisterPathMapperServer(conn dcerpc.Conn, o PathMapperServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewPathMapperServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(PathMapperSyntaxV0_0))...)
}

func NewPathMapperServerHandle(o PathMapperServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return PathMapperServerHandle(ctx, o, opNum, r)
	}
}

func PathMapperServerHandle(ctx context.Context, o PathMapperServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // GetSharePathsForLocalPath
		op := &xxx_GetSharePathsForLocalPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSharePathsForLocalPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSharePathsForLocalPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmPathMapper
type UnimplementedPathMapperServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedPathMapperServer) GetSharePathsForLocalPath(context.Context, *GetSharePathsForLocalPathRequest) (*GetSharePathsForLocalPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ PathMapperServer = (*UnimplementedPathMapperServer)(nil)
