package ivdsdisk2

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

// IVdsDisk2 server interface.
type Disk2Server interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The SetSANMode method sets the SAN mode of a disk to either offline or online. A
	// disk that is offline exposes no volume devices for partitions or volumes with extents
	// on that disk. A disk can be REAONLY or READWRITE independent of the offline or online
	// setting.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	//
	// ERROR_SUCCESS (0x00000000)
	SetSANMode(context.Context, *SetSANModeRequest) (*SetSANModeResponse, error)
}

func RegisterDisk2Server(conn dcerpc.Conn, o Disk2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewDisk2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Disk2SyntaxV0_0))...)
}

func NewDisk2ServerHandle(o Disk2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Disk2ServerHandle(ctx, o, opNum, r)
	}
}

func Disk2ServerHandle(ctx context.Context, o Disk2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // SetSANMode
		op := &xxx_SetSANModeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSANModeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSANMode(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IVdsDisk2
type UnimplementedDisk2Server struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedDisk2Server) SetSANMode(context.Context, *SetSANModeRequest) (*SetSANModeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ Disk2Server = (*UnimplementedDisk2Server)(nil)
