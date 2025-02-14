package rfri

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
)

// rfri server interface.
type RfriServer interface {

	// The RfrGetNewDSA method returns the name of an NSPI server or a server array.
	//
	// Return Values: The server returns 0 for a successful execution. An error results
	// in an HRESULT or other nonzero error code.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol as specified in [MS-RPCE].
	GetNewDSA(context.Context, *GetNewDSARequest) (*GetNewDSAResponse, error)

	// The RfrGetFQDNFromServerDN method returns the Domain Name System (DNS) FQDN of the
	// server corresponding to the passed DN.
	//
	// Return Values: The server returns 0 for a successful execution. An error results
	// in an HRESULT or other nonzero error code.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol as specified in [MS-RPCE].
	GetFQDNFromServerDN(context.Context, *GetFQDNFromServerDNRequest) (*GetFQDNFromServerDNResponse, error)
}

func RegisterRfriServer(conn dcerpc.Conn, o RfriServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRfriServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RfriSyntaxV1_0))...)
}

func NewRfriServerHandle(o RfriServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RfriServerHandle(ctx, o, opNum, r)
	}
}

func RfriServerHandle(ctx context.Context, o RfriServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // RfrGetNewDSA
		op := &xxx_GetNewDSAOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNewDSARequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNewDSA(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // RfrGetFQDNFromServerDN
		op := &xxx_GetFQDNFromServerDNOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFQDNFromServerDNRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFQDNFromServerDN(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented rfri
type UnimplementedRfriServer struct {
}

func (UnimplementedRfriServer) GetNewDSA(context.Context, *GetNewDSARequest) (*GetNewDSAResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRfriServer) GetFQDNFromServerDN(context.Context, *GetFQDNFromServerDNRequest) (*GetFQDNFromServerDNResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ RfriServer = (*UnimplementedRfriServer)(nil)
