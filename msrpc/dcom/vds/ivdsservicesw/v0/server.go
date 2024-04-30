package ivdsservicesw

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

// IVdsServiceSw server interface.
type ServiceSwServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The GetDiskObject method<75> returns the disk for the given PnP Device ID string.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT, as specified in
	// [MS-ERREF], to indicate success or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetDiskObject(context.Context, *GetDiskObjectRequest) (*GetDiskObjectResponse, error)
}

func RegisterServiceSwServer(conn dcerpc.Conn, o ServiceSwServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewServiceSwServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ServiceSwSyntaxV0_0))...)
}

func NewServiceSwServerHandle(o ServiceSwServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ServiceSwServerHandle(ctx, o, opNum, r)
	}
}

func ServiceSwServerHandle(ctx context.Context, o ServiceSwServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetDiskObject
		in := &GetDiskObjectRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDiskObject(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
