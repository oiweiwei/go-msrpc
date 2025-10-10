package intmslibrarycontrol2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	intmslibrarycontrol1 "github.com/oiweiwei/go-msrpc/msrpc/dcom/rsmp/intmslibrarycontrol1/v0"
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
	_ = intmslibrarycontrol1.GoPackage
)

// INtmsLibraryControl2 server interface.
type LibraryControl2Server interface {

	// INtmsLibraryControl1 base class.
	intmslibrarycontrol1.LibraryControl1Server

	IdentifyNTMSSlot(context.Context, *IdentifyNTMSSlotRequest) (*IdentifyNTMSSlotResponse, error)
}

func RegisterLibraryControl2Server(conn dcerpc.Conn, o LibraryControl2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewLibraryControl2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(LibraryControl2SyntaxV0_0))...)
}

func NewLibraryControl2ServerHandle(o LibraryControl2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return LibraryControl2ServerHandle(ctx, o, opNum, r)
	}
}

func LibraryControl2ServerHandle(ctx context.Context, o LibraryControl2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 20 {
		// INtmsLibraryControl1 base method.
		return intmslibrarycontrol1.LibraryControl1ServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 20: // IdentifyNtmsSlot
		op := &xxx_IdentifyNTMSSlotOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IdentifyNTMSSlotRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IdentifyNTMSSlot(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented INtmsLibraryControl2
type UnimplementedLibraryControl2Server struct {
	intmslibrarycontrol1.UnimplementedLibraryControl1Server
}

func (UnimplementedLibraryControl2Server) IdentifyNTMSSlot(context.Context, *IdentifyNTMSSlotRequest) (*IdentifyNTMSSlotResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ LibraryControl2Server = (*UnimplementedLibraryControl2Server)(nil)
