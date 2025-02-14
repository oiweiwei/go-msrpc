package ifsrmcommittablecollection

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	ifsrmmutablecollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmmutablecollection/v0"
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
	_ = ifsrmmutablecollection.GoPackage
)

// IFsrmCommittableCollection server interface.
type CommittableCollectionServer interface {

	// IFsrmMutableCollection base class.
	ifsrmmutablecollection.MutableCollectionServer

	// Commit operation.
	Commit(context.Context, *CommitRequest) (*CommitResponse, error)
}

func RegisterCommittableCollectionServer(conn dcerpc.Conn, o CommittableCollectionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewCommittableCollectionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(CommittableCollectionSyntaxV0_0))...)
}

func NewCommittableCollectionServerHandle(o CommittableCollectionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return CommittableCollectionServerHandle(ctx, o, opNum, r)
	}
}

func CommittableCollectionServerHandle(ctx context.Context, o CommittableCollectionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 18 {
		// IFsrmMutableCollection base method.
		return ifsrmmutablecollection.MutableCollectionServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 18: // Commit
		op := &xxx_CommitOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CommitRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Commit(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmCommittableCollection
type UnimplementedCommittableCollectionServer struct {
	ifsrmmutablecollection.UnimplementedMutableCollectionServer
}

func (UnimplementedCommittableCollectionServer) Commit(context.Context, *CommitRequest) (*CommitResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ CommittableCollectionServer = (*UnimplementedCommittableCollectionServer)(nil)
