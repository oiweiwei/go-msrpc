package idatafactory2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	idatafactory "github.com/oiweiwei/go-msrpc/msrpc/dcom/adtg/idatafactory/v0"
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
	_ = idatafactory.GoPackage
)

// IDataFactory2 server interface.
type DataFactory2Server interface {

	// IDataFactory base class.
	idatafactory.DataFactoryServer

	// Execute21 operation.
	Execute21(context.Context, *Execute21Request) (*Execute21Response, error)

	// Synchronize21 operation.
	Synchronize21(context.Context, *Synchronize21Request) (*Synchronize21Response, error)
}

func RegisterDataFactory2Server(conn dcerpc.Conn, o DataFactory2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewDataFactory2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(DataFactory2SyntaxV0_0))...)
}

func NewDataFactory2ServerHandle(o DataFactory2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return DataFactory2ServerHandle(ctx, o, opNum, r)
	}
}

func DataFactory2ServerHandle(ctx context.Context, o DataFactory2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDataFactory base method.
		return idatafactory.DataFactoryServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Execute21
		op := &xxx_Execute21Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &Execute21Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Execute21(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // Synchronize21
		op := &xxx_Synchronize21Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &Synchronize21Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Synchronize21(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IDataFactory2
type UnimplementedDataFactory2Server struct {
	idatafactory.UnimplementedDataFactoryServer
}

func (UnimplementedDataFactory2Server) Execute21(context.Context, *Execute21Request) (*Execute21Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataFactory2Server) Synchronize21(context.Context, *Synchronize21Request) (*Synchronize21Response, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ DataFactory2Server = (*UnimplementedDataFactory2Server)(nil)
