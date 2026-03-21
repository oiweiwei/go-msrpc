package tsvippublic

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

// TSVIPPublic server interface.
type PublicServer interface {

	// The RpcGetSessionIP method retrieves the IP address assigned to the session. This
	// MUST be called by an administrator or the same user who logged onto the session.<214>
	// The method performs access checks as defined in section 3.1.3 and 3.1.4. The method
	// fails if both checks fail.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success; otherwise, it
	// MUST return an implementation-specific negative value.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x00000000 S_OK   | Successful completion. |
	//	+-------------------+------------------------+
	GetSessionIP(context.Context, *GetSessionIPRequest) (*GetSessionIPResponse, error)
}

func RegisterPublicServer(conn dcerpc.Conn, o PublicServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewPublicServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(PublicSyntaxV1_0))...)
}

func NewPublicServerHandle(o PublicServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return PublicServerHandle(ctx, o, opNum, r)
	}
}

func PublicServerHandle(ctx context.Context, o PublicServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // RpcGetSessionIP
		op := &xxx_GetSessionIPOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSessionIPRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSessionIP(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented TSVIPPublic
type UnimplementedPublicServer struct {
}

func (UnimplementedPublicServer) GetSessionIP(context.Context, *GetSessionIPRequest) (*GetSessionIPResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ PublicServer = (*UnimplementedPublicServer)(nil)
