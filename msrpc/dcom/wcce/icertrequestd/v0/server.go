package icertrequestd

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

// ICertRequestD server interface.
type CertRequestDServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The Request method initiates the certificate issuance process.
	//
	// Return Values:  The method MUST return zero unless otherwise explicitly stated in
	// this section. The server MUST return errors through the pdwDisposition parameter.
	//
	// This section, and the following sections, describe the processing rules for ICertRequestD::Request
	// and ICertRequestD2::Request2.
	Request(context.Context, *RequestRequest) (*RequestResponse, error)

	// The GetCACert method returns property values on the CA. The main use of this method
	// is to enable clients to diagnose issues and the state of the server. In addition,
	// one of the properties returned by this method is required to support the advanced
	// CA functionality (GETCERT_CAXCHGCERT).
	//
	// Return Values: For a successful invocation, the CA MUST return 0; otherwise, the
	// CA MUST return a nonzero value.
	GetCACert(context.Context, *GetCACertRequest) (*GetCACertResponse, error)

	// The Ping method performs a request response test (ping) to the CA.
	//
	// Return Values: For a successful invocation, the CA MUST return 0; otherwise, the
	// CA MUST return a nonzero value.
	Ping(context.Context, *PingRequest) (*PingResponse, error)
}

func RegisterCertRequestDServer(conn dcerpc.Conn, o CertRequestDServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewCertRequestDServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(CertRequestDSyntaxV0_0))...)
}

func NewCertRequestDServerHandle(o CertRequestDServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return CertRequestDServerHandle(ctx, o, opNum, r)
	}
}

func CertRequestDServerHandle(ctx context.Context, o CertRequestDServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Request
		op := &xxx_RequestOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RequestRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Request(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // GetCACert
		op := &xxx_GetCACertOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCACertRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCACert(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // Ping
		op := &xxx_PingOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PingRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Ping(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented ICertRequestD
type UnimplementedCertRequestDServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedCertRequestDServer) Request(context.Context, *RequestRequest) (*RequestResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertRequestDServer) GetCACert(context.Context, *GetCACertRequest) (*GetCACertResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertRequestDServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ CertRequestDServer = (*UnimplementedCertRequestDServer)(nil)
