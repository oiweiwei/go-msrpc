package iimport2

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

// IImport2 server interface.
type Import2Server interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// This method is called by a client to set the import target partition (as specified
	// in section 3.1.1.5).
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result (as specified in [MS-ERREF], section 2.1) on failure. All failure results
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
	SetPartition(context.Context, *SetPartitionRequest) (*SetPartitionResponse, error)

	// Opnum4NotUsedOnWire operation.
	// Opnum4NotUsedOnWire

	// Opnum5NotUsedOnWire operation.
	// Opnum5NotUsedOnWire
}

func RegisterImport2Server(conn dcerpc.Conn, o Import2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewImport2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Import2SyntaxV0_0))...)
}

func NewImport2ServerHandle(o Import2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Import2ServerHandle(ctx, o, opNum, r)
	}
}

func Import2ServerHandle(ctx context.Context, o Import2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // SetPartition
		op := &xxx_SetPartitionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetPartitionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetPartition(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // Opnum4NotUsedOnWire
		// Opnum4NotUsedOnWire
		return nil, nil
	case 5: // Opnum5NotUsedOnWire
		// Opnum5NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}

// Unimplemented IImport2
type UnimplementedImport2Server struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedImport2Server) SetPartition(context.Context, *SetPartitionRequest) (*SetPartitionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ Import2Server = (*UnimplementedImport2Server)(nil)
