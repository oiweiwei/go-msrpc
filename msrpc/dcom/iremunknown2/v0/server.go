package iremunknown2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iremunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iremunknown/v0"
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
	_ = iremunknown.GoPackage
)

// IRemUnknown2 server interface.
type RemoteUnknown2Server interface {

	// IRemUnknown base class.
	iremunknown.RemoteUnknownServer

	// The RemQueryInterface2 (Opnum 6) method acquires standard object references (see
	// section 2.2.18.1) to additional interfaces on the object, marshaled as an MInterfacePointer
	// structure.
	RemoteQueryInterface2(context.Context, *RemoteQueryInterface2Request) (*RemoteQueryInterface2Response, error)
}

func RegisterRemoteUnknown2Server(conn dcerpc.Conn, o RemoteUnknown2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRemoteUnknown2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RemoteUnknown2SyntaxV0_0))...)
}

func NewRemoteUnknown2ServerHandle(o RemoteUnknown2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RemoteUnknown2ServerHandle(ctx, o, opNum, r)
	}
}

func RemoteUnknown2ServerHandle(ctx context.Context, o RemoteUnknown2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 6 {
		// IRemUnknown base method.
		return iremunknown.RemoteUnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 6: // RemQueryInterface2
		in := &RemoteQueryInterface2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoteQueryInterface2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
