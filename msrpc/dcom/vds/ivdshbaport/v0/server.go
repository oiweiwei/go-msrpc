package ivdshbaport

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

// IVdsHbaPort server interface.
type HBAPortServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// GetProperties operation.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)

	// The SetAllPathStatuses method sets the statuses of all paths that originate from
	// the HBA port to a specified status.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure.<76> For the HRESULT values predefined by the Virtual Disk
	// Service Remote Protocol, see section 2.2.3.
	SetAllPathStatuses(context.Context, *SetAllPathStatusesRequest) (*SetAllPathStatusesResponse, error)
}

func RegisterHBAPortServer(conn dcerpc.Conn, o HBAPortServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewHBAPortServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(HBAPortSyntaxV0_0))...)
}

func NewHBAPortServerHandle(o HBAPortServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return HBAPortServerHandle(ctx, o, opNum, r)
	}
}

func HBAPortServerHandle(ctx context.Context, o HBAPortServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetProperties
		in := &GetPropertiesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetProperties(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // SetAllPathStatuses
		in := &SetAllPathStatusesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetAllPathStatuses(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
