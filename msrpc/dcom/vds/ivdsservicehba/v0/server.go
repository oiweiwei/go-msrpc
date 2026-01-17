package ivdsservicehba

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

// IVdsServiceHba server interface.
type ServiceHBAServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The QueryHbaPorts method returns an IEnumVdsObject enumeration object that contains
	// a list of the HBA ports that are known to VDS on the system.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryHBAPorts(context.Context, *QueryHBAPortsRequest) (*QueryHBAPortsResponse, error)
}

func RegisterServiceHBAServer(conn dcerpc.Conn, o ServiceHBAServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewServiceHBAServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ServiceHBASyntaxV0_0))...)
}

func NewServiceHBAServerHandle(o ServiceHBAServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ServiceHBAServerHandle(ctx, o, opNum, r)
	}
}

func ServiceHBAServerHandle(ctx context.Context, o ServiceHBAServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // QueryHbaPorts
		op := &xxx_QueryHBAPortsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryHBAPortsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryHBAPorts(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IVdsServiceHba
type UnimplementedServiceHBAServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedServiceHBAServer) QueryHBAPorts(context.Context, *QueryHBAPortsRequest) (*QueryHBAPortsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ServiceHBAServer = (*UnimplementedServiceHBAServer)(nil)
