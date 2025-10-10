package intmsobjectmanagement1

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

// INtmsObjectManagement1 server interface.
type NTMSObjectManagement1Server interface {

	// IUnknown base class.
	iunknown.UnknownServer

	GetNTMSObjectSecurity(context.Context, *GetNTMSObjectSecurityRequest) (*GetNTMSObjectSecurityResponse, error)

	SetNTMSObjectSecurity(context.Context, *SetNTMSObjectSecurityRequest) (*SetNTMSObjectSecurityResponse, error)

	GetNTMSObjectAttributeA(context.Context, *GetNTMSObjectAttributeARequest) (*GetNTMSObjectAttributeAResponse, error)

	GetNTMSObjectAttributeW(context.Context, *GetNTMSObjectAttributeWRequest) (*GetNTMSObjectAttributeWResponse, error)

	SetNTMSObjectAttributeA(context.Context, *SetNTMSObjectAttributeARequest) (*SetNTMSObjectAttributeAResponse, error)

	SetNTMSObjectAttributeW(context.Context, *SetNTMSObjectAttributeWRequest) (*SetNTMSObjectAttributeWResponse, error)

	EnumerateNTMSObject(context.Context, *EnumerateNTMSObjectRequest) (*EnumerateNTMSObjectResponse, error)

	DisableNTMSObject(context.Context, *DisableNTMSObjectRequest) (*DisableNTMSObjectResponse, error)

	EnableNTMSObject(context.Context, *EnableNTMSObjectRequest) (*EnableNTMSObjectResponse, error)
}

func RegisterNTMSObjectManagement1Server(conn dcerpc.Conn, o NTMSObjectManagement1Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewNTMSObjectManagement1ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(NTMSObjectManagement1SyntaxV0_0))...)
}

func NewNTMSObjectManagement1ServerHandle(o NTMSObjectManagement1Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return NTMSObjectManagement1ServerHandle(ctx, o, opNum, r)
	}
}

func NTMSObjectManagement1ServerHandle(ctx context.Context, o NTMSObjectManagement1Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // GetNtmsObjectSecurity
		op := &xxx_GetNTMSObjectSecurityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNTMSObjectSecurityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNTMSObjectSecurity(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // SetNtmsObjectSecurity
		op := &xxx_SetNTMSObjectSecurityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetNTMSObjectSecurityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetNTMSObjectSecurity(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // GetNtmsObjectAttributeA
		op := &xxx_GetNTMSObjectAttributeAOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNTMSObjectAttributeARequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNTMSObjectAttributeA(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // GetNtmsObjectAttributeW
		op := &xxx_GetNTMSObjectAttributeWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNTMSObjectAttributeWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNTMSObjectAttributeW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // SetNtmsObjectAttributeA
		op := &xxx_SetNTMSObjectAttributeAOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetNTMSObjectAttributeARequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetNTMSObjectAttributeA(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // SetNtmsObjectAttributeW
		op := &xxx_SetNTMSObjectAttributeWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetNTMSObjectAttributeWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetNTMSObjectAttributeW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // EnumerateNtmsObject
		op := &xxx_EnumerateNTMSObjectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumerateNTMSObjectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumerateNTMSObject(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // DisableNtmsObject
		op := &xxx_DisableNTMSObjectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DisableNTMSObjectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DisableNTMSObject(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // EnableNtmsObject
		op := &xxx_EnableNTMSObjectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnableNTMSObjectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnableNTMSObject(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented INtmsObjectManagement1
type UnimplementedNTMSObjectManagement1Server struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedNTMSObjectManagement1Server) GetNTMSObjectSecurity(context.Context, *GetNTMSObjectSecurityRequest) (*GetNTMSObjectSecurityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSObjectManagement1Server) SetNTMSObjectSecurity(context.Context, *SetNTMSObjectSecurityRequest) (*SetNTMSObjectSecurityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSObjectManagement1Server) GetNTMSObjectAttributeA(context.Context, *GetNTMSObjectAttributeARequest) (*GetNTMSObjectAttributeAResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSObjectManagement1Server) GetNTMSObjectAttributeW(context.Context, *GetNTMSObjectAttributeWRequest) (*GetNTMSObjectAttributeWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSObjectManagement1Server) SetNTMSObjectAttributeA(context.Context, *SetNTMSObjectAttributeARequest) (*SetNTMSObjectAttributeAResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSObjectManagement1Server) SetNTMSObjectAttributeW(context.Context, *SetNTMSObjectAttributeWRequest) (*SetNTMSObjectAttributeWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSObjectManagement1Server) EnumerateNTMSObject(context.Context, *EnumerateNTMSObjectRequest) (*EnumerateNTMSObjectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSObjectManagement1Server) DisableNTMSObject(context.Context, *DisableNTMSObjectRequest) (*DisableNTMSObjectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSObjectManagement1Server) EnableNTMSObject(context.Context, *EnableNTMSObjectRequest) (*EnableNTMSObjectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ NTMSObjectManagement1Server = (*UnimplementedNTMSObjectManagement1Server)(nil)
