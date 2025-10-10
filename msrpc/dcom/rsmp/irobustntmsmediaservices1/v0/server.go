package irobustntmsmediaservices1

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	intmsmediaservices1 "github.com/oiweiwei/go-msrpc/msrpc/dcom/rsmp/intmsmediaservices1/v0"
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
	_ = intmsmediaservices1.GoPackage
)

// IRobustNtmsMediaServices1 server interface.
type RobustNTMSMediaServices1Server interface {

	// INtmsMediaServices1 base class.
	intmsmediaservices1.NTMSMediaServices1Server

	GetNTMSMediaPoolNameAr(context.Context, *GetNTMSMediaPoolNameArRequest) (*GetNTMSMediaPoolNameArResponse, error)

	GetNTMSMediaPoolNameWr(context.Context, *GetNTMSMediaPoolNameWrRequest) (*GetNTMSMediaPoolNameWrResponse, error)
}

func RegisterRobustNTMSMediaServices1Server(conn dcerpc.Conn, o RobustNTMSMediaServices1Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRobustNTMSMediaServices1ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RobustNTMSMediaServices1SyntaxV0_0))...)
}

func NewRobustNTMSMediaServices1ServerHandle(o RobustNTMSMediaServices1Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RobustNTMSMediaServices1ServerHandle(ctx, o, opNum, r)
	}
}

func RobustNTMSMediaServices1ServerHandle(ctx context.Context, o RobustNTMSMediaServices1Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 18 {
		// INtmsMediaServices1 base method.
		return intmsmediaservices1.NTMSMediaServices1ServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 18: // GetNtmsMediaPoolNameAR
		op := &xxx_GetNTMSMediaPoolNameArOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNTMSMediaPoolNameArRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNTMSMediaPoolNameAr(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // GetNtmsMediaPoolNameWR
		op := &xxx_GetNTMSMediaPoolNameWrOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNTMSMediaPoolNameWrRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNTMSMediaPoolNameWr(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IRobustNtmsMediaServices1
type UnimplementedRobustNTMSMediaServices1Server struct {
	intmsmediaservices1.UnimplementedNTMSMediaServices1Server
}

func (UnimplementedRobustNTMSMediaServices1Server) GetNTMSMediaPoolNameAr(context.Context, *GetNTMSMediaPoolNameArRequest) (*GetNTMSMediaPoolNameArResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRobustNTMSMediaServices1Server) GetNTMSMediaPoolNameWr(context.Context, *GetNTMSMediaPoolNameWrRequest) (*GetNTMSMediaPoolNameWrResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ RobustNTMSMediaServices1Server = (*UnimplementedRobustNTMSMediaServices1Server)(nil)
