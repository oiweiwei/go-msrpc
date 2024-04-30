package ivdsvolumeshrink

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

// IVdsVolumeShrink server interface.
type VolumeShrinkServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The QueryMaxReclaimableBytes method retrieves the maximum number of bytes that can
	// be reclaimed from the current volume.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryMaxReclaimableBytes(context.Context, *QueryMaxReclaimableBytesRequest) (*QueryMaxReclaimableBytesResponse, error)

	// Shrink operation.
	Shrink(context.Context, *ShrinkRequest) (*ShrinkResponse, error)
}

func RegisterVolumeShrinkServer(conn dcerpc.Conn, o VolumeShrinkServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewVolumeShrinkServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(VolumeShrinkSyntaxV0_0))...)
}

func NewVolumeShrinkServerHandle(o VolumeShrinkServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return VolumeShrinkServerHandle(ctx, o, opNum, r)
	}
}

func VolumeShrinkServerHandle(ctx context.Context, o VolumeShrinkServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // QueryMaxReclaimableBytes
		in := &QueryMaxReclaimableBytesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryMaxReclaimableBytes(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // Shrink
		in := &ShrinkRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Shrink(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
