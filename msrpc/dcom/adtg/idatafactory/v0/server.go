package idatafactory

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
		op := &xxx_QueryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Query(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // SubmitChanges
		op := &xxx_SubmitChangesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SubmitChangesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SubmitChanges(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // ConvertToString
		op := &xxx_ConvertToStringOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ConvertToStringRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ConvertToString(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // CreateRecordSet
		op := &xxx_CreateRecordSetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateRecordSetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateRecordSet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IDataFactory
type UnimplementedDataFactoryServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedDataFactoryServer) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataFactoryServer) SubmitChanges(context.Context, *SubmitChangesRequest) (*SubmitChangesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataFactoryServer) ConvertToString(context.Context, *ConvertToStringRequest) (*ConvertToStringResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataFactoryServer) CreateRecordSet(context.Context, *CreateRecordSetRequest) (*CreateRecordSetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ DataFactoryServer = (*UnimplementedDataFactoryServer)(nil)
