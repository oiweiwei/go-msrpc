package idatafactory

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

// IDataFactory server interface.
type DataFactoryServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Query operation.
	Query(context.Context, *QueryRequest) (*QueryResponse, error)

	// SubmitChanges operation.
	SubmitChanges(context.Context, *SubmitChangesRequest) (*SubmitChangesResponse, error)

	// ConvertToString operation.
	ConvertToString(context.Context, *ConvertToStringRequest) (*ConvertToStringResponse, error)

	// CreateRecordSet operation.
	CreateRecordSet(context.Context, *CreateRecordSetRequest) (*CreateRecordSetResponse, error)
}

func RegisterDataFactoryServer(conn dcerpc.Conn, o DataFactoryServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewDataFactoryServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(DataFactorySyntaxV0_0))...)
}

func NewDataFactoryServerHandle(o DataFactoryServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return DataFactoryServerHandle(ctx, o, opNum, r)
	}
}

func DataFactoryServerHandle(ctx context.Context, o DataFactoryServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Query
		in := &QueryRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Query(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // SubmitChanges
		in := &SubmitChangesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SubmitChanges(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // ConvertToString
		in := &ConvertToStringRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ConvertToString(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // CreateRecordSet
		in := &CreateRecordSetRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateRecordSet(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
