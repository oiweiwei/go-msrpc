package ivdscreatepartitionex

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

// IVdsCreatePartitionEx server interface.
type CreatePartitionExServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The CreatePartitionEx method creates a partition on a disk at a specified byte offset,
	// with an optional alignment parameter.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	//
	// Note  Creating or deleting partitions on dynamic disks is not supported.
	//
	// ERROR_SUCCESS (0x00000000)
	CreatePartitionEx(context.Context, *CreatePartitionExRequest) (*CreatePartitionExResponse, error)
}

func RegisterCreatePartitionExServer(conn dcerpc.Conn, o CreatePartitionExServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewCreatePartitionExServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(CreatePartitionExSyntaxV0_0))...)
}

func NewCreatePartitionExServerHandle(o CreatePartitionExServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return CreatePartitionExServerHandle(ctx, o, opNum, r)
	}
}

func CreatePartitionExServerHandle(ctx context.Context, o CreatePartitionExServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // CreatePartitionEx
		in := &CreatePartitionExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreatePartitionEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
