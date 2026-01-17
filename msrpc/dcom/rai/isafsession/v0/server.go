package isafsession

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
		op := &xxx_GetSessionIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSessionIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSessionID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // SessionID
		op := &xxx_SetSessionIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSessionIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSessionID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // SessionState
		op := &xxx_GetSessionStateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSessionStateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSessionState(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // SessionState
		op := &xxx_SetSessionStateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSessionStateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSessionState(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // DomainName
		op := &xxx_GetDomainNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDomainNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDomainName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // DomainName
		op := &xxx_SetDomainNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetDomainNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetDomainName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // UserName
		op := &xxx_GetUserNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUserNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUserName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // UserName
		op := &xxx_SetUserNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetUserNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetUserName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented ISAFSession
type UnimplementedSAFSessionServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedSAFSessionServer) GetSessionID(context.Context, *GetSessionIDRequest) (*GetSessionIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSAFSessionServer) SetSessionID(context.Context, *SetSessionIDRequest) (*SetSessionIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSAFSessionServer) GetSessionState(context.Context, *GetSessionStateRequest) (*GetSessionStateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSAFSessionServer) SetSessionState(context.Context, *SetSessionStateRequest) (*SetSessionStateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSAFSessionServer) GetDomainName(context.Context, *GetDomainNameRequest) (*GetDomainNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSAFSessionServer) SetDomainName(context.Context, *SetDomainNameRequest) (*SetDomainNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSAFSessionServer) GetUserName(context.Context, *GetUserNameRequest) (*GetUserNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSAFSessionServer) SetUserName(context.Context, *SetUserNameRequest) (*SetUserNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ SAFSessionServer = (*UnimplementedSAFSessionServer)(nil)
