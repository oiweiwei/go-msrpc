package iservicedcomponentinfo

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

// IServicedComponentInfo server interface.
type ServicedComponentInfoServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The GetComponentInfo method is used to determine the environment of the server object.
	//
	// Return Values: An HRESULT that specifies success or failure. All success HRESULT
	// values MUST be treated as success and all failure HRESULT values MUST be treated
	// as failure.
	GetComponentInfo(context.Context, *GetComponentInfoRequest) (*GetComponentInfoResponse, error)
}

func RegisterServicedComponentInfoServer(conn dcerpc.Conn, o ServicedComponentInfoServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewServicedComponentInfoServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ServicedComponentInfoSyntaxV0_0))...)
}

func NewServicedComponentInfoServerHandle(o ServicedComponentInfoServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ServicedComponentInfoServerHandle(ctx, o, opNum, r)
	}
}

func ServicedComponentInfoServerHandle(ctx context.Context, o ServicedComponentInfoServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetComponentInfo
		in := &GetComponentInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetComponentInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
