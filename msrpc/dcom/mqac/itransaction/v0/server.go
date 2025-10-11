package itransaction

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

// ITransaction server interface.
type ITransactionServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Commit operation.
	Commit(context.Context, *CommitRequest) (*CommitResponse, error)

	// Abort operation.
	Abort(context.Context, *AbortRequest) (*AbortResponse, error)

	// The GetTransactionInfo method retrieves information about the transaction represented
	// by its TransactionImpl (section 3.8) instance.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.<32>
	//
	// When the server processes this call, it MUST follow these guidelines:
	//
	// * If the pinfo output variable is NULL:
	//
	// * Return an error HRESULT, and take no further action.
	//
	// * If the isInternal instance variable is True:
	//
	// * Set the uow property of the pinfo output variable to the value of the TransactionIdentifier
	// instance variable.
	//
	// * Set all the remaining properties of the pinfo output variable to 0.
	//
	// * Else:
	//
	// * Set the uow property of the pinfo output variable to the value of the TransactionIdentifier
	// instance variable.
	//
	// * Send a request to the distributed transaction manager ( 3b7be3f7-651c-4f9c-930b-a9a7c4355ad8#gt_4553803e-9d8d-407c-ad7d-9e65e01d6eb3
	// ) to retrieve the transaction details by using the TransactionIdentifier instance
	// variable. For more details about this request, refer to [MS-DTCO] ( ../ms-dtco/c367c571-33f3-44ac-85cb-4b9ebbb2779d
	// ) section 2.2.8.3.1.1 ( ../ms-dtco/3a6ed3f4-9395-4046-ba10-c334e2337473 ).
	//
	// * Set the remaining properties of the pinfo output variable to the values obtained
	// from the preceding request.
	GetTransactionInfo(context.Context, *GetTransactionInfoRequest) (*GetTransactionInfoResponse, error)
}

func RegisterITransactionServer(conn dcerpc.Conn, o ITransactionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewITransactionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ITransactionSyntaxV0_0))...)
}

func NewITransactionServerHandle(o ITransactionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ITransactionServerHandle(ctx, o, opNum, r)
	}
}

func ITransactionServerHandle(ctx context.Context, o ITransactionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Commit
		op := &xxx_CommitOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CommitRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Commit(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // Abort
		op := &xxx_AbortOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AbortRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Abort(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // GetTransactionInfo
		op := &xxx_GetTransactionInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTransactionInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTransactionInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented ITransaction
type UnimplementedITransactionServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedITransactionServer) Commit(context.Context, *CommitRequest) (*CommitResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedITransactionServer) Abort(context.Context, *AbortRequest) (*AbortResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedITransactionServer) GetTransactionInfo(context.Context, *GetTransactionInfoRequest) (*GetTransactionInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ITransactionServer = (*UnimplementedITransactionServer)(nil)
