package iautomaticupdates2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iautomaticupdates "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iautomaticupdates/v0"
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
	_ = iautomaticupdates.GoPackage
)

// IAutomaticUpdates2 server interface.
type AutomaticUpdates2Server interface {

	// IAutomaticUpdates base class.
	iautomaticupdates.AutomaticUpdatesServer

	// The IAutomaticUpdates2::Results (opnum 15) method retrieves an IAutomaticUpdatesResults
	// instance.
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
	// This method SHOULD return the value of the Results ADM element.
	GetResults(context.Context, *GetResultsRequest) (*GetResultsResponse, error)
}

func RegisterAutomaticUpdates2Server(conn dcerpc.Conn, o AutomaticUpdates2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAutomaticUpdates2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AutomaticUpdates2SyntaxV0_0))...)
}

func NewAutomaticUpdates2ServerHandle(o AutomaticUpdates2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AutomaticUpdates2ServerHandle(ctx, o, opNum, r)
	}
}

func AutomaticUpdates2ServerHandle(ctx context.Context, o AutomaticUpdates2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 14 {
		// IAutomaticUpdates base method.
		return iautomaticupdates.AutomaticUpdatesServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 14: // Results
		op := &xxx_GetResultsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetResultsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetResults(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAutomaticUpdates2
type UnimplementedAutomaticUpdates2Server struct {
	iautomaticupdates.UnimplementedAutomaticUpdatesServer
}

func (UnimplementedAutomaticUpdates2Server) GetResults(context.Context, *GetResultsRequest) (*GetResultsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AutomaticUpdates2Server = (*UnimplementedAutomaticUpdates2Server)(nil)
