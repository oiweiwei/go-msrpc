package initshutdown

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
)

// InitShutdown server interface.
type InitShutdownServer interface {

	// The BaseInitiateShutdown method is used to initiate the shutdown of the remote computer.<8>
	//
	// Return Values: The method returns ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it returns a nonzero error code.<9>
	BaseInitiateShutdown(context.Context, *BaseInitiateShutdownRequest) (*BaseInitiateShutdownResponse, error)

	// The BaseAbortShutdown method is used to terminate the shutdown of the remote computer
	// within the waiting period.<10>
	//
	// Return Values: The method returns ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it returns a nonzero error code.
	BaseAbortShutdown(context.Context, *BaseAbortShutdownRequest) (*BaseAbortShutdownResponse, error)

	// The BaseInitiateShutdownEx method is used to initiate the shutdown of the remote
	// computer.<11>
	//
	// Return Values: The method returns ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it returns a nonzero error code.
	BaseInitiateShutdownEx(context.Context, *BaseInitiateShutdownExRequest) (*BaseInitiateShutdownExResponse, error)
}

func RegisterInitShutdownServer(conn dcerpc.Conn, o InitShutdownServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewInitShutdownServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(InitShutdownSyntaxV1_0))...)
}

func NewInitShutdownServerHandle(o InitShutdownServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return InitShutdownServerHandle(ctx, o, opNum, r)
	}
}

func InitShutdownServerHandle(ctx context.Context, o InitShutdownServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // BaseInitiateShutdown
		op := &xxx_BaseInitiateShutdownOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BaseInitiateShutdownRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BaseInitiateShutdown(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // BaseAbortShutdown
		op := &xxx_BaseAbortShutdownOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BaseAbortShutdownRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BaseAbortShutdown(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // BaseInitiateShutdownEx
		op := &xxx_BaseInitiateShutdownExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BaseInitiateShutdownExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BaseInitiateShutdownEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented InitShutdown
type UnimplementedInitShutdownServer struct {
}

func (UnimplementedInitShutdownServer) BaseInitiateShutdown(context.Context, *BaseInitiateShutdownRequest) (*BaseInitiateShutdownResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedInitShutdownServer) BaseAbortShutdown(context.Context, *BaseAbortShutdownRequest) (*BaseAbortShutdownResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedInitShutdownServer) BaseInitiateShutdownEx(context.Context, *BaseInitiateShutdownExRequest) (*BaseInitiateShutdownExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ InitShutdownServer = (*UnimplementedInitShutdownServer)(nil)
