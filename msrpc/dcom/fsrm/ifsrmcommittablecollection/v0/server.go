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
		in := &CommitRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Commit(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
