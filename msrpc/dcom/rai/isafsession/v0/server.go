package isafsession

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
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
	_ = idispatch.GoPackage
)

// ISAFSession server interface.
type SAFSessionServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// SessionID operation.
	GetSessionID(context.Context, *GetSessionIDRequest) (*GetSessionIDResponse, error)

	// SessionID operation.
	SetSessionID(context.Context, *SetSessionIDRequest) (*SetSessionIDResponse, error)

	// SessionState operation.
	GetSessionState(context.Context, *GetSessionStateRequest) (*GetSessionStateResponse, error)

	// SessionState operation.
	SetSessionState(context.Context, *SetSessionStateRequest) (*SetSessionStateResponse, error)

	// DomainName operation.
	GetDomainName(context.Context, *GetDomainNameRequest) (*GetDomainNameResponse, error)

	// DomainName operation.
	SetDomainName(context.Context, *SetDomainNameRequest) (*SetDomainNameResponse, error)

	// UserName operation.
	GetUserName(context.Context, *GetUserNameRequest) (*GetUserNameResponse, error)

	// UserName operation.
	SetUserName(context.Context, *SetUserNameRequest) (*SetUserNameResponse, error)
}

func RegisterSAFSessionServer(conn dcerpc.Conn, o SAFSessionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewSAFSessionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(SAFSessionSyntaxV0_0))...)
}

func NewSAFSessionServerHandle(o SAFSessionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return SAFSessionServerHandle(ctx, o, opNum, r)
	}
}

func SAFSessionServerHandle(ctx context.Context, o SAFSessionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // SessionID
		in := &GetSessionIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSessionID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // SessionID
		in := &SetSessionIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSessionID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // SessionState
		in := &GetSessionStateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSessionState(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // SessionState
		in := &SetSessionStateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSessionState(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // DomainName
		in := &GetDomainNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDomainName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // DomainName
		in := &SetDomainNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetDomainName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // UserName
		in := &GetUserNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetUserName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // UserName
		in := &SetUserNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetUserName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
