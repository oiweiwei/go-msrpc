package iupdatesession3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iupdatesession2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdatesession2/v0"
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
	_ = iupdatesession2.GoPackage
)

// IUpdateSession3 server interface.
type UpdateSession3Server interface {

	// IUpdateSession2 base class.
	iupdatesession2.UpdateSession2Server

	// The IUpdateSession3::CreateUpdateServiceManager (Opnum 16) method retrieves an instance
	// of the IUpdateServiceManager2 interface.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return a newly created object that implements IUpdateServiceManager2.
	CreateUpdateServiceManager(context.Context, *CreateUpdateServiceManagerRequest) (*CreateUpdateServiceManagerResponse, error)

	// The IUpdateSearcher::QueryHistory (opnum 19) method retrieves a collection of history
	// events.
	//
	// The IUpdateSession3::QueryHistory (Opnum 17) method retrieves relevant update history
	// entries.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	QueryHistory(context.Context, *QueryHistoryRequest) (*QueryHistoryResponse, error)
}

func RegisterUpdateSession3Server(conn dcerpc.Conn, o UpdateSession3Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUpdateSession3ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(UpdateSession3SyntaxV0_0))...)
}

func NewUpdateSession3ServerHandle(o UpdateSession3Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return UpdateSession3ServerHandle(ctx, o, opNum, r)
	}
}

func UpdateSession3ServerHandle(ctx context.Context, o UpdateSession3Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 17 {
		// IUpdateSession2 base method.
		return iupdatesession2.UpdateSession2ServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 17: // CreateUpdateServiceManager
		op := &xxx_CreateUpdateServiceManagerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateUpdateServiceManagerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateUpdateServiceManager(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // QueryHistory
		op := &xxx_QueryHistoryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryHistoryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryHistory(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IUpdateSession3
type UnimplementedUpdateSession3Server struct {
	iupdatesession2.UnimplementedUpdateSession2Server
}

func (UnimplementedUpdateSession3Server) CreateUpdateServiceManager(context.Context, *CreateUpdateServiceManagerRequest) (*CreateUpdateServiceManagerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSession3Server) QueryHistory(context.Context, *QueryHistoryRequest) (*QueryHistoryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ UpdateSession3Server = (*UnimplementedUpdateSession3Server)(nil)
