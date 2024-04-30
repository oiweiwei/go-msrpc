package dsaop

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

// dsaop server interface.
type DsaopServer interface {

	// The IDL_DSAPrepareScript method prepares the DC to run a maintenance script.
	//
	// Return Values: 0 if successful, or a Windows error code if a failure occurs.
	PrepareScript(context.Context, *PrepareScriptRequest) (*PrepareScriptResponse, error)

	// The IDL_DSAExecuteScript method executes a maintenance script.
	//
	// Return Values: 0 if successful, or a Windows error code if a failure occurs.
	ExecuteScript(context.Context, *ExecuteScriptRequest) (*ExecuteScriptResponse, error)
}

func RegisterDsaopServer(conn dcerpc.Conn, o DsaopServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewDsaopServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(DsaopSyntaxV1_0))...)
}

func NewDsaopServerHandle(o DsaopServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return DsaopServerHandle(ctx, o, opNum, r)
	}
}

func DsaopServerHandle(ctx context.Context, o DsaopServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // IDL_DSAPrepareScript
		in := &PrepareScriptRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PrepareScript(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // IDL_DSAExecuteScript
		in := &ExecuteScriptRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ExecuteScript(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
