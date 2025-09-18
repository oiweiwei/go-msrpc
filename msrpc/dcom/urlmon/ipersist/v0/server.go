package ipersist

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

// IPersist server interface.
type PersistServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// GetClassID operation.
	GetClassID(context.Context, *GetClassIDRequest) (*GetClassIDResponse, error)
}

func RegisterPersistServer(conn dcerpc.Conn, o PersistServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewPersistServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(PersistSyntaxV0_0))...)
}

func NewPersistServerHandle(o PersistServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return PersistServerHandle(ctx, o, opNum, r)
	}
}

func PersistServerHandle(ctx context.Context, o PersistServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetClassID
		op := &xxx_GetClassIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetClassIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetClassID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IPersist
type UnimplementedPersistServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedPersistServer) GetClassID(context.Context, *GetClassIDRequest) (*GetClassIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ PersistServer = (*UnimplementedPersistServer)(nil)
