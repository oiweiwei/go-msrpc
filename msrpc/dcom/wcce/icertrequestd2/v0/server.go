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
		in := &Request2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Request2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // GetCAProperty
		in := &GetCAPropertyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetCAProperty(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // GetCAPropertyInfo
		in := &GetCAPropertyInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetCAPropertyInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // Ping2
		in := &Ping2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Ping2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
