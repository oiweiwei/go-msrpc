package intmsobjectmanagement2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	intmsobjectmanagement1 "github.com/oiweiwei/go-msrpc/msrpc/dcom/rsmp/intmsobjectmanagement1/v0"
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
	_ = intmsobjectmanagement1.GoPackage
)

// INtmsObjectManagement2 server interface.
type NTMSObjectManagement2Server interface {

	// INtmsObjectManagement1 base class.
	intmsobjectmanagement1.NTMSObjectManagement1Server

	EnumerateNTMSObjectR(context.Context, *EnumerateNTMSObjectRRequest) (*EnumerateNTMSObjectRResponse, error)

	GetNTMSUIOptionsA(context.Context, *GetNTMSUIOptionsARequest) (*GetNTMSUIOptionsAResponse, error)

	GetNTMSUIOptionsW(context.Context, *GetNTMSUIOptionsWRequest) (*GetNTMSUIOptionsWResponse, error)

	SetNTMSUIOptionsA(context.Context, *SetNTMSUIOptionsARequest) (*SetNTMSUIOptionsAResponse, error)

	SetNTMSUIOptionsW(context.Context, *SetNTMSUIOptionsWRequest) (*SetNTMSUIOptionsWResponse, error)
}

func RegisterNTMSObjectManagement2Server(conn dcerpc.Conn, o NTMSObjectManagement2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewNTMSObjectManagement2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(NTMSObjectManagement2SyntaxV0_0))...)
}

func NewNTMSObjectManagement2ServerHandle(o NTMSObjectManagement2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return NTMSObjectManagement2ServerHandle(ctx, o, opNum, r)
	}
}

func NTMSObjectManagement2ServerHandle(ctx context.Context, o NTMSObjectManagement2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 9 {
		// INtmsObjectManagement1 base method.
		return intmsobjectmanagement1.NTMSObjectManagement1ServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 9: // EnumerateNtmsObjectR
		op := &xxx_EnumerateNTMSObjectROperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumerateNTMSObjectRRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumerateNTMSObjectR(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // GetNtmsUIOptionsA
		op := &xxx_GetNTMSUIOptionsAOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNTMSUIOptionsARequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNTMSUIOptionsA(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // GetNtmsUIOptionsW
		op := &xxx_GetNTMSUIOptionsWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNTMSUIOptionsWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNTMSUIOptionsW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // SetNtmsUIOptionsA
		op := &xxx_SetNTMSUIOptionsAOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetNTMSUIOptionsARequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetNTMSUIOptionsA(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // SetNtmsUIOptionsW
		op := &xxx_SetNTMSUIOptionsWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetNTMSUIOptionsWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetNTMSUIOptionsW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented INtmsObjectManagement2
type UnimplementedNTMSObjectManagement2Server struct {
	intmsobjectmanagement1.UnimplementedNTMSObjectManagement1Server
}

func (UnimplementedNTMSObjectManagement2Server) EnumerateNTMSObjectR(context.Context, *EnumerateNTMSObjectRRequest) (*EnumerateNTMSObjectRResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSObjectManagement2Server) GetNTMSUIOptionsA(context.Context, *GetNTMSUIOptionsARequest) (*GetNTMSUIOptionsAResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSObjectManagement2Server) GetNTMSUIOptionsW(context.Context, *GetNTMSUIOptionsWRequest) (*GetNTMSUIOptionsWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSObjectManagement2Server) SetNTMSUIOptionsA(context.Context, *SetNTMSUIOptionsARequest) (*SetNTMSUIOptionsAResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSObjectManagement2Server) SetNTMSUIOptionsW(context.Context, *SetNTMSUIOptionsWRequest) (*SetNTMSUIOptionsWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ NTMSObjectManagement2Server = (*UnimplementedNTMSObjectManagement2Server)(nil)
