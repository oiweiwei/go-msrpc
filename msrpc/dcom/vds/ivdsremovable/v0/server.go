package ivdsremovable

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

// IVdsRemovable server interface.
type RemovableServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The QueryMedia method identifies the media in the drive.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	//
	// ERROR_SUCCESS (0x00000000)
	QueryMedia(context.Context, *QueryMediaRequest) (*QueryMediaResponse, error)

	// The Eject method ejects the media in the drive.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Eject(context.Context, *EjectRequest) (*EjectResponse, error)
}

func RegisterRemovableServer(conn dcerpc.Conn, o RemovableServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRemovableServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RemovableSyntaxV0_0))...)
}

func NewRemovableServerHandle(o RemovableServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RemovableServerHandle(ctx, o, opNum, r)
	}
}

func RemovableServerHandle(ctx context.Context, o RemovableServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // QueryMedia
		op := &xxx_QueryMediaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryMediaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryMedia(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // Eject
		op := &xxx_EjectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EjectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Eject(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IVdsRemovable
type UnimplementedRemovableServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedRemovableServer) QueryMedia(context.Context, *QueryMediaRequest) (*QueryMediaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemovableServer) Eject(context.Context, *EjectRequest) (*EjectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ RemovableServer = (*UnimplementedRemovableServer)(nil)
