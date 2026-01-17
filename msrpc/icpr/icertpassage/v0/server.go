package icertpassage

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
)

// ICertPassage server interface.
type CertPassageServer interface {

	// The CertServerRequest method processes a certificate enrollment request from the
	// client.<6>
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success. This
	// method's return values MUST have identical syntax and semantics to the return values
	// specified in [MS-WCCE] section 3.2.1.4.2.1.
	//
	// If the ADM element Config.CA.Interface.Flags contains the value IF_NORPCICERTREQUEST,
	// the server SHOULD return an error.<7>
	//
	// If the ADM element Config.CA.Interface.Flags contains the value IF_ENFORCEENCRYPTICERTREQUEST
	// and the RPC_C_AUTHN_LEVEL_PKT_PRIVACY authentication level ([MS-RPCE] section 2.2.1.1.8)
	// is not specified on the RPC connection from the client, the CA MUST refuse to establish
	// a connection with the client by returning E_ACCESSDENIED (0x80000009).
	//
	// Otherwise, the processing rules for the ICertRequestD::Request method ([MS-WCCE]
	// section 3.2.2.6.2.1) apply, except that if the ADM element Config.CA.Interface.Flags
	// contains the value IF_NOREMOTEICERTREQUEST, these values are ignored and the request
	// is processed as though the values were absent.
	CertServerRequest(context.Context, *CertServerRequestRequest) (*CertServerRequestResponse, error)
}

func RegisterCertPassageServer(conn dcerpc.Conn, o CertPassageServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewCertPassageServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(CertPassageSyntaxV0_0))...)
}

func NewCertPassageServerHandle(o CertPassageServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return CertPassageServerHandle(ctx, o, opNum, r)
	}
}

func CertPassageServerHandle(ctx context.Context, o CertPassageServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // CertServerRequest
		op := &xxx_CertServerRequestOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CertServerRequestRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CertServerRequest(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented ICertPassage
type UnimplementedCertPassageServer struct {
}

func (UnimplementedCertPassageServer) CertServerRequest(context.Context, *CertServerRequestRequest) (*CertServerRequestResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ CertPassageServer = (*UnimplementedCertPassageServer)(nil)
