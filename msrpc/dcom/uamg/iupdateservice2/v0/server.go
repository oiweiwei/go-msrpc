package iupdateservice2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iupdateservice "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdateservice/v0"
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
	_ = iupdateservice.GoPackage
)

// IUpdateService2 server interface.
type UpdateService2Server interface {

	// IUpdateService base class.
	iupdateservice.UpdateServiceServer

	// The IUpdateService2::IsDefaultAUService (opnum 21) method retrieves whether this
	// update service is the default update service for the automatic update agent.
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
	// This method SHOULD return the value of the IsDefaultAUService ADM element.
	GetIsDefaultAUService(context.Context, *GetIsDefaultAUServiceRequest) (*GetIsDefaultAUServiceResponse, error)
}

func RegisterUpdateService2Server(conn dcerpc.Conn, o UpdateService2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUpdateService2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(UpdateService2SyntaxV0_0))...)
}

func NewUpdateService2ServerHandle(o UpdateService2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return UpdateService2ServerHandle(ctx, o, opNum, r)
	}
}

func UpdateService2ServerHandle(ctx context.Context, o UpdateService2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 20 {
		// IUpdateService base method.
		return iupdateservice.UpdateServiceServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 20: // IsDefaultAUService
		op := &xxx_GetIsDefaultAUServiceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsDefaultAUServiceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsDefaultAUService(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IUpdateService2
type UnimplementedUpdateService2Server struct {
	iupdateservice.UnimplementedUpdateServiceServer
}

func (UnimplementedUpdateService2Server) GetIsDefaultAUService(context.Context, *GetIsDefaultAUServiceRequest) (*GetIsDefaultAUServiceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ UpdateService2Server = (*UnimplementedUpdateService2Server)(nil)
