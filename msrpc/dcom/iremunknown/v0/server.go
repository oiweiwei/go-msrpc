package iremunknown

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

// IRemUnknown server interface.
type RemoteUnknownServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// This RemQueryInterface (Opnum 3) method acquires standard object references (see
	// section 2.2.18.1) to additional interfaces on the object.
	RemoteQueryInterface(context.Context, *RemoteQueryInterfaceRequest) (*RemoteQueryInterfaceResponse, error)

	// The RemAddRef (Opnum 4) method requests that a specified number of reference counts
	// be incremented on a specified number of interfaces on the object.
	RemoteAddReference(context.Context, *RemoteAddReferenceRequest) (*RemoteAddReferenceResponse, error)

	// The RemRelease (Opnum 5) method requests that a specified number of reference counts
	// be decremented on a specified number of interfaces on an object.
	RemoteRelease(context.Context, *RemoteReleaseRequest) (*RemoteReleaseResponse, error)
}

func RegisterRemoteUnknownServer(conn dcerpc.Conn, o RemoteUnknownServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRemoteUnknownServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RemoteUnknownSyntaxV0_0))...)
}

func NewRemoteUnknownServerHandle(o RemoteUnknownServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RemoteUnknownServerHandle(ctx, o, opNum, r)
	}
}

func RemoteUnknownServerHandle(ctx context.Context, o RemoteUnknownServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // RemQueryInterface
		in := &RemoteQueryInterfaceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoteQueryInterface(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // RemAddRef
		in := &RemoteAddReferenceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoteAddReference(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // RemRelease
		in := &RemoteReleaseRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoteRelease(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
