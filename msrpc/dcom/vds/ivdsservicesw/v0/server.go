package ivdsservicesw

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
		op := &xxx_GetDiskObjectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDiskObjectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDiskObject(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IVdsServiceSw
type UnimplementedServiceSwServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedServiceSwServer) GetDiskObject(context.Context, *GetDiskObjectRequest) (*GetDiskObjectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ServiceSwServer = (*UnimplementedServiceSwServer)(nil)
