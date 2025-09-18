package istream

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

// IStream server interface.
type StreamServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Read operation.
	Read(context.Context, *ReadRequest) (*ReadResponse, error)

	// Write operation.
	Write(context.Context, *WriteRequest) (*WriteResponse, error)

	// Seek operation.
	Seek(context.Context, *SeekRequest) (*SeekResponse, error)

	// SetSize operation.
	SetSize(context.Context, *SetSizeRequest) (*SetSizeResponse, error)

	// CopyTo operation.
	CopyTo(context.Context, *CopyToRequest) (*CopyToResponse, error)

	// Commit operation.
	Commit(context.Context, *CommitRequest) (*CommitResponse, error)

	// Revert operation.
	Revert(context.Context, *RevertRequest) (*RevertResponse, error)

	// LockRegion operation.
	LockRegion(context.Context, *LockRegionRequest) (*LockRegionResponse, error)

	// UnlockRegion operation.
	UnlockRegion(context.Context, *UnlockRegionRequest) (*UnlockRegionResponse, error)

	// Stat operation.
	Stat(context.Context, *StatRequest) (*StatResponse, error)

	// Clone operation.
	Clone(context.Context, *CloneRequest) (*CloneResponse, error)
}

func RegisterStreamServer(conn dcerpc.Conn, o StreamServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewStreamServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(StreamSyntaxV0_0))...)
}

func NewStreamServerHandle(o StreamServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return StreamServerHandle(ctx, o, opNum, r)
	}
}

func StreamServerHandle(ctx context.Context, o StreamServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Read
		op := &xxx_ReadOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReadRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Read(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // Write
		op := &xxx_WriteOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &WriteRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Write(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // Seek
		op := &xxx_SeekOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SeekRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Seek(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // SetSize
		op := &xxx_SetSizeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSizeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSize(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // CopyTo
		op := &xxx_CopyToOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CopyToRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CopyTo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // Commit
		op := &xxx_CommitOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CommitRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Commit(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // Revert
		op := &xxx_RevertOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RevertRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Revert(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // LockRegion
		op := &xxx_LockRegionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &LockRegionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.LockRegion(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // UnlockRegion
		op := &xxx_UnlockRegionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UnlockRegionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.UnlockRegion(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // Stat
		op := &xxx_StatOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StatRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Stat(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // Clone
		op := &xxx_CloneOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloneRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Clone(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IStream
type UnimplementedStreamServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedStreamServer) Read(context.Context, *ReadRequest) (*ReadResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedStreamServer) Write(context.Context, *WriteRequest) (*WriteResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedStreamServer) Seek(context.Context, *SeekRequest) (*SeekResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedStreamServer) SetSize(context.Context, *SetSizeRequest) (*SetSizeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedStreamServer) CopyTo(context.Context, *CopyToRequest) (*CopyToResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedStreamServer) Commit(context.Context, *CommitRequest) (*CommitResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedStreamServer) Revert(context.Context, *RevertRequest) (*RevertResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedStreamServer) LockRegion(context.Context, *LockRegionRequest) (*LockRegionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedStreamServer) UnlockRegion(context.Context, *UnlockRegionRequest) (*UnlockRegionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedStreamServer) Stat(context.Context, *StatRequest) (*StatResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedStreamServer) Clone(context.Context, *CloneRequest) (*CloneResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ StreamServer = (*UnimplementedStreamServer)(nil)
