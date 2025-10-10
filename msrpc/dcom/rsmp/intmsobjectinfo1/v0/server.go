package intmsobjectinfo1

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
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
	_ = iunknown.GoPackage
)

// INtmsObjectInfo1 server interface.
type NTMSObjectInfo1Server interface {

	// IUnknown base class.
	iunknown.UnknownServer

	GetNTMSServerObjectInformationA(context.Context, *GetNTMSServerObjectInformationARequest) (*GetNTMSServerObjectInformationAResponse, error)

	GetNTMSServerObjectInformationW(context.Context, *GetNTMSServerObjectInformationWRequest) (*GetNTMSServerObjectInformationWResponse, error)

	SetNTMSObjectInformationA(context.Context, *SetNTMSObjectInformationARequest) (*SetNTMSObjectInformationAResponse, error)

	SetNTMSObjectInformationW(context.Context, *SetNTMSObjectInformationWRequest) (*SetNTMSObjectInformationWResponse, error)

	CreateNTMSMediaA(context.Context, *CreateNTMSMediaARequest) (*CreateNTMSMediaAResponse, error)

	CreateNTMSMediaW(context.Context, *CreateNTMSMediaWRequest) (*CreateNTMSMediaWResponse, error)
}

func RegisterNTMSObjectInfo1Server(conn dcerpc.Conn, o NTMSObjectInfo1Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewNTMSObjectInfo1ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(NTMSObjectInfo1SyntaxV0_0))...)
}

func NewNTMSObjectInfo1ServerHandle(o NTMSObjectInfo1Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return NTMSObjectInfo1ServerHandle(ctx, o, opNum, r)
	}
}

func NTMSObjectInfo1ServerHandle(ctx context.Context, o NTMSObjectInfo1Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // GetNtmsServerObjectInformationA
		op := &xxx_GetNTMSServerObjectInformationAOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNTMSServerObjectInformationARequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNTMSServerObjectInformationA(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // GetNtmsServerObjectInformationW
		op := &xxx_GetNTMSServerObjectInformationWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNTMSServerObjectInformationWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNTMSServerObjectInformationW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // SetNtmsObjectInformationA
		op := &xxx_SetNTMSObjectInformationAOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetNTMSObjectInformationARequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetNTMSObjectInformationA(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // SetNtmsObjectInformationW
		op := &xxx_SetNTMSObjectInformationWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetNTMSObjectInformationWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetNTMSObjectInformationW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // CreateNtmsMediaA
		op := &xxx_CreateNTMSMediaAOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateNTMSMediaARequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateNTMSMediaA(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // CreateNtmsMediaW
		op := &xxx_CreateNTMSMediaWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateNTMSMediaWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateNTMSMediaW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented INtmsObjectInfo1
type UnimplementedNTMSObjectInfo1Server struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedNTMSObjectInfo1Server) GetNTMSServerObjectInformationA(context.Context, *GetNTMSServerObjectInformationARequest) (*GetNTMSServerObjectInformationAResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSObjectInfo1Server) GetNTMSServerObjectInformationW(context.Context, *GetNTMSServerObjectInformationWRequest) (*GetNTMSServerObjectInformationWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSObjectInfo1Server) SetNTMSObjectInformationA(context.Context, *SetNTMSObjectInformationARequest) (*SetNTMSObjectInformationAResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSObjectInfo1Server) SetNTMSObjectInformationW(context.Context, *SetNTMSObjectInformationWRequest) (*SetNTMSObjectInformationWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSObjectInfo1Server) CreateNTMSMediaA(context.Context, *CreateNTMSMediaARequest) (*CreateNTMSMediaAResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSObjectInfo1Server) CreateNTMSMediaW(context.Context, *CreateNTMSMediaWRequest) (*CreateNTMSMediaWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ NTMSObjectInfo1Server = (*UnimplementedNTMSObjectInfo1Server)(nil)
