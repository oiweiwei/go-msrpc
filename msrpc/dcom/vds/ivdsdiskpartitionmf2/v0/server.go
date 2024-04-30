package ivdsdiskpartitionmf2

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

// IVdsDiskPartitionMF2 server interface.
type DiskPartitionMF2Server interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The FormatPartitionEx2 method formats an existing OEM, ESP, or unknown partition.
	// This method is only supported on OEM, ESP, recovery, and unknown partitions.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	FormatPartitionEx2(context.Context, *FormatPartitionEx2Request) (*FormatPartitionEx2Response, error)
}

func RegisterDiskPartitionMF2Server(conn dcerpc.Conn, o DiskPartitionMF2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewDiskPartitionMF2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(DiskPartitionMF2SyntaxV0_0))...)
}

func NewDiskPartitionMF2ServerHandle(o DiskPartitionMF2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return DiskPartitionMF2ServerHandle(ctx, o, opNum, r)
	}
}

func DiskPartitionMF2ServerHandle(ctx context.Context, o DiskPartitionMF2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // FormatPartitionEx2
		in := &FormatPartitionEx2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FormatPartitionEx2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
