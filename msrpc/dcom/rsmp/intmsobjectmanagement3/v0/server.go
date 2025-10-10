package intmsobjectmanagement3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	intmsobjectmanagement2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/rsmp/intmsobjectmanagement2/v0"
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
	_ = intmsobjectmanagement2.GoPackage
)

// INtmsObjectManagement3 server interface.
type NTMSObjectManagement3Server interface {

	// INtmsObjectManagement2 base class.
	intmsobjectmanagement2.NTMSObjectManagement2Server

	GetNTMSObjectAttributeAr(context.Context, *GetNTMSObjectAttributeArRequest) (*GetNTMSObjectAttributeArResponse, error)

	GetNTMSObjectAttributeWr(context.Context, *GetNTMSObjectAttributeWrRequest) (*GetNTMSObjectAttributeWrResponse, error)
}

func RegisterNTMSObjectManagement3Server(conn dcerpc.Conn, o NTMSObjectManagement3Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewNTMSObjectManagement3ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(NTMSObjectManagement3SyntaxV0_0))...)
}

func NewNTMSObjectManagement3ServerHandle(o NTMSObjectManagement3Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return NTMSObjectManagement3ServerHandle(ctx, o, opNum, r)
	}
}

func NTMSObjectManagement3ServerHandle(ctx context.Context, o NTMSObjectManagement3Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 14 {
		// INtmsObjectManagement2 base method.
		return intmsobjectmanagement2.NTMSObjectManagement2ServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 14: // GetNtmsObjectAttributeAR
		op := &xxx_GetNTMSObjectAttributeArOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNTMSObjectAttributeArRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNTMSObjectAttributeAr(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // GetNtmsObjectAttributeWR
		op := &xxx_GetNTMSObjectAttributeWrOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNTMSObjectAttributeWrRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNTMSObjectAttributeWr(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented INtmsObjectManagement3
type UnimplementedNTMSObjectManagement3Server struct {
	intmsobjectmanagement2.UnimplementedNTMSObjectManagement2Server
}

func (UnimplementedNTMSObjectManagement3Server) GetNTMSObjectAttributeAr(context.Context, *GetNTMSObjectAttributeArRequest) (*GetNTMSObjectAttributeArResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSObjectManagement3Server) GetNTMSObjectAttributeWr(context.Context, *GetNTMSObjectAttributeWrRequest) (*GetNTMSObjectAttributeWrResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ NTMSObjectManagement3Server = (*UnimplementedNTMSObjectManagement3Server)(nil)
