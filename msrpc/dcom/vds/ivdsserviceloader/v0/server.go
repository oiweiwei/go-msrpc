package ivdsserviceloader

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

// IVdsServiceLoader server interface.
type ServiceLoaderServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The LoadService method is used by client applications to load the VDS service on
	// a remote machine.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	LoadService(context.Context, *LoadServiceRequest) (*LoadServiceResponse, error)
}

func RegisterServiceLoaderServer(conn dcerpc.Conn, o ServiceLoaderServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewServiceLoaderServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ServiceLoaderSyntaxV0_0))...)
}

func NewServiceLoaderServerHandle(o ServiceLoaderServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ServiceLoaderServerHandle(ctx, o, opNum, r)
	}
}

func ServiceLoaderServerHandle(ctx context.Context, o ServiceLoaderServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // LoadService
		in := &LoadServiceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.LoadService(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
