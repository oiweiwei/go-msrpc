package ivdsadvanceddisk2

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

// IVdsAdvancedDisk2 server interface.
type AdvancedDisk2Server interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The ChangePartitionType method changes the partition type on the disk at a specified
	// byte offset.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	//
	// ERROR_SUCCESS (0x00000000)
	ChangePartitionType(context.Context, *ChangePartitionTypeRequest) (*ChangePartitionTypeResponse, error)
}

func RegisterAdvancedDisk2Server(conn dcerpc.Conn, o AdvancedDisk2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAdvancedDisk2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AdvancedDisk2SyntaxV0_0))...)
}

func NewAdvancedDisk2ServerHandle(o AdvancedDisk2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AdvancedDisk2ServerHandle(ctx, o, opNum, r)
	}
}

func AdvancedDisk2ServerHandle(ctx context.Context, o AdvancedDisk2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // ChangePartitionType
		op := &xxx_ChangePartitionTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ChangePartitionTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ChangePartitionType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IVdsAdvancedDisk2
type UnimplementedAdvancedDisk2Server struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAdvancedDisk2Server) ChangePartitionType(context.Context, *ChangePartitionTypeRequest) (*ChangePartitionTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AdvancedDisk2Server = (*UnimplementedAdvancedDisk2Server)(nil)
