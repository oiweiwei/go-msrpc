package icatalogtablewrite

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

// ICatalogTableWrite server interface.
type CatalogTableWriteServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// This method is called by a client to write entries to a catalog table.
	//
	// Return Values:  This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically, with the exception of E_DETAILEDERRORS (0x80110802).
	WriteTable(context.Context, *WriteTableRequest) (*WriteTableResponse, error)
}

func RegisterCatalogTableWriteServer(conn dcerpc.Conn, o CatalogTableWriteServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewCatalogTableWriteServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(CatalogTableWriteSyntaxV0_0))...)
}

func NewCatalogTableWriteServerHandle(o CatalogTableWriteServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return CatalogTableWriteServerHandle(ctx, o, opNum, r)
	}
}

func CatalogTableWriteServerHandle(ctx context.Context, o CatalogTableWriteServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // WriteTable
		op := &xxx_WriteTableOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &WriteTableRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.WriteTable(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented ICatalogTableWrite
type UnimplementedCatalogTableWriteServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedCatalogTableWriteServer) WriteTable(context.Context, *WriteTableRequest) (*WriteTableResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ CatalogTableWriteServer = (*UnimplementedCatalogTableWriteServer)(nil)
