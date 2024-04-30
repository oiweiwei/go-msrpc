package idatafactory2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
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
		in := &Execute21Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Execute21(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // Synchronize21
		in := &Synchronize21Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Synchronize21(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
