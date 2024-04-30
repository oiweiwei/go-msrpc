package icatalogtableinfo

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

// ICatalogTableInfo server interface.
type CatalogTableInfoServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// This method is called by a client to retrieve the table metadata (section 3.1.1.2.1)
	// for a catalog table.
	//
	// Return Values:  This method MUST return S_OK (0x00000000) on success and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	GetClientTableInfo(context.Context, *GetClientTableInfoRequest) (*GetClientTableInfoResponse, error)
}

func RegisterCatalogTableInfoServer(conn dcerpc.Conn, o CatalogTableInfoServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewCatalogTableInfoServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(CatalogTableInfoSyntaxV0_0))...)
}

func NewCatalogTableInfoServerHandle(o CatalogTableInfoServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return CatalogTableInfoServerHandle(ctx, o, opNum, r)
	}
}

func CatalogTableInfoServerHandle(ctx context.Context, o CatalogTableInfoServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetClientTableInfo
		in := &GetClientTableInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetClientTableInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
