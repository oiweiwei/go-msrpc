package icatalogtableread

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

// ICatalogTableRead server interface.
type CatalogTableReadServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// This method is called by a client to read entries from a catalog table according
	// to a query.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically, with the exception of E_DETAILEDERRORS (0x80110802).
	ReadTable(context.Context, *ReadTableRequest) (*ReadTableResponse, error)
}

func RegisterCatalogTableReadServer(conn dcerpc.Conn, o CatalogTableReadServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewCatalogTableReadServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(CatalogTableReadSyntaxV0_0))...)
}

func NewCatalogTableReadServerHandle(o CatalogTableReadServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return CatalogTableReadServerHandle(ctx, o, opNum, r)
	}
}

func CatalogTableReadServerHandle(ctx context.Context, o CatalogTableReadServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // ReadTable
		op := &xxx_ReadTableOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReadTableRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReadTable(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented ICatalogTableRead
type UnimplementedCatalogTableReadServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedCatalogTableReadServer) ReadTable(context.Context, *ReadTableRequest) (*ReadTableResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ CatalogTableReadServer = (*UnimplementedCatalogTableReadServer)(nil)
