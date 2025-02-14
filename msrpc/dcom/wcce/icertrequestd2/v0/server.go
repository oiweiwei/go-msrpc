package icertrequestd2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	icertrequestd "github.com/oiweiwei/go-msrpc/msrpc/dcom/wcce/icertrequestd/v0"
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
	_ = icertrequestd.GoPackage
)

// ICertRequestD2 server interface.
type CertRequestD2Server interface {

	// ICertRequestD base class.
	icertrequestd.CertRequestDServer

	// The Request2 method requests a certificate from the CA. It is similar to the ICertRequestD::Request
	// method, but it has an additional parameter, pwszSerialNumber, which is specified
	// as follows.
	//
	// Return Values: Identical to the return value of the ICertRequestD::Request method.
	Request2(context.Context, *Request2Request) (*Request2Response, error)

	// The GetCAProperty method retrieves a property value from the CA.
	//
	// Return Values: For a successful invocation, the CA MUST return 0; otherwise, the
	// CA MUST return a nonzero value.
	GetCAProperty(context.Context, *GetCAPropertyRequest) (*GetCAPropertyResponse, error)

	// The GetCAPropertyInfo method retrieves a set of property structures from the CA.
	// The list of properties is specified in section 3.2.1.4.3.2.
	//
	// Return Values: For a successful invocation, the CA MUST return 0. Otherwise, the
	// CA MUST return a nonzero value.
	GetCAPropertyInfo(context.Context, *GetCAPropertyInfoRequest) (*GetCAPropertyInfoResponse, error)

	// The Ping2 method pings the CA.
	//
	// Return Values: For a successful invocation, the CA MUST return 0; otherwise, the
	// CA MUST return a nonzero value.
	Ping2(context.Context, *Ping2Request) (*Ping2Response, error)
}

func RegisterCertRequestD2Server(conn dcerpc.Conn, o CertRequestD2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewCertRequestD2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(CertRequestD2SyntaxV0_0))...)
}

func NewCertRequestD2ServerHandle(o CertRequestD2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return CertRequestD2ServerHandle(ctx, o, opNum, r)
	}
}

func CertRequestD2ServerHandle(ctx context.Context, o CertRequestD2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 6 {
		// ICertRequestD base method.
		return icertrequestd.CertRequestDServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 6: // Request2
		op := &xxx_Request2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &Request2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Request2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // GetCAProperty
		op := &xxx_GetCAPropertyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCAPropertyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCAProperty(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // GetCAPropertyInfo
		op := &xxx_GetCAPropertyInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCAPropertyInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCAPropertyInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // Ping2
		op := &xxx_Ping2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &Ping2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Ping2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented ICertRequestD2
type UnimplementedCertRequestD2Server struct {
	icertrequestd.UnimplementedCertRequestDServer
}

func (UnimplementedCertRequestD2Server) Request2(context.Context, *Request2Request) (*Request2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertRequestD2Server) GetCAProperty(context.Context, *GetCAPropertyRequest) (*GetCAPropertyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertRequestD2Server) GetCAPropertyInfo(context.Context, *GetCAPropertyInfoRequest) (*GetCAPropertyInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertRequestD2Server) Ping2(context.Context, *Ping2Request) (*Ping2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ CertRequestD2Server = (*UnimplementedCertRequestD2Server)(nil)
