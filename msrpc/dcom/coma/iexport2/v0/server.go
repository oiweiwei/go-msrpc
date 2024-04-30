package iexport2

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

// IExport2 server interface.
type Export2Server interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// This method is called by a client to export all conglomerations in a partition at
	// once to an installer package file.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF], section 2.1, on failure. All failure results
	// MUST be treated identically.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	ExportPartition(context.Context, *ExportPartitionRequest) (*ExportPartitionResponse, error)
}

func RegisterExport2Server(conn dcerpc.Conn, o Export2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewExport2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Export2SyntaxV0_0))...)
}

func NewExport2ServerHandle(o Export2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Export2ServerHandle(ctx, o, opNum, r)
	}
}

func Export2ServerHandle(ctx context.Context, o Export2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // ExportPartition
		in := &ExportPartitionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ExportPartition(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
