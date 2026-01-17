package iupdate3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iupdate2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdate2/v0"
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
	_ = iupdate2.GoPackage
)

// IUpdate3 server interface.
type Update3Server interface {

	// IUpdate2 base class.
	iupdate2.Update2Server

	// The IUpdate3::BrowseOnly (opnum 57) method retrieves whether the update is browse-only.
	//
	// The IWindowsDriverUpdate3::BrowseOnly (opnum 65) method retrieves whether the update
	// is browse-only. If an update is browse-only, then it is not recommended for the update
	// to be installed automatically.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the BrowseOnly ADM element.
	GetBrowseOnly(context.Context, *GetBrowseOnlyRequest) (*GetBrowseOnlyResponse, error)
}

func RegisterUpdate3Server(conn dcerpc.Conn, o Update3Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUpdate3ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Update3SyntaxV0_0))...)
}

func NewUpdate3ServerHandle(o Update3Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Update3ServerHandle(ctx, o, opNum, r)
	}
}

func Update3ServerHandle(ctx context.Context, o Update3Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 56 {
		// IUpdate2 base method.
		return iupdate2.Update2ServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 56: // BrowseOnly
		op := &xxx_GetBrowseOnlyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetBrowseOnlyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetBrowseOnly(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IUpdate3
type UnimplementedUpdate3Server struct {
	iupdate2.UnimplementedUpdate2Server
}

func (UnimplementedUpdate3Server) GetBrowseOnly(context.Context, *GetBrowseOnlyRequest) (*GetBrowseOnlyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ Update3Server = (*UnimplementedUpdate3Server)(nil)
