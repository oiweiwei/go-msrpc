package intmsobjectmanagement1

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = iunknown.GoPackage
)

// INtmsObjectManagement1 server interface.
type ObjectManagement1Server interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The GetNtmsObjectSecurity method retrieves the security descriptor of an object.
	GetNTMSObjectSecurity(context.Context, *GetNTMSObjectSecurityRequest) (*GetNTMSObjectSecurityResponse, error)

	// The SetNtmsObjectSecurity method changes the security descriptor of an object.
	SetNTMSObjectSecurity(context.Context, *SetNTMSObjectSecurityRequest) (*SetNTMSObjectSecurityResponse, error)

	// The GetNtmsObjectAttributeA method retrieves private data of an object, with strings
	// encoded using ASCII.
	GetNTMSObjectAttributeA(context.Context, *GetNTMSObjectAttributeARequest) (*GetNTMSObjectAttributeAResponse, error)

	// The GetNtmsObjectAttributeW method retrieves private data from an object, with strings
	// encoded using Unicode.
	GetNTMSObjectAttributeW(context.Context, *GetNTMSObjectAttributeWRequest) (*GetNTMSObjectAttributeWResponse, error)

	// The SetNtmsObjectAttributeA method changes the private data of an object, with strings
	// encoded using ASCII.
	SetNTMSObjectAttributeA(context.Context, *SetNTMSObjectAttributeARequest) (*SetNTMSObjectAttributeAResponse, error)

	// The SetNtmsObjectAttributeW method changes the private data of an object, with strings
	// encoded using Unicode.
	SetNTMSObjectAttributeW(context.Context, *SetNTMSObjectAttributeWRequest) (*SetNTMSObjectAttributeWResponse, error)

	// The EnumerateNtmsObject method enumerates the objects of the container specified
	// by the lpContainerId parameter.
	EnumerateNTMSObject(context.Context, *EnumerateNTMSObjectRequest) (*EnumerateNTMSObjectResponse, error)

	// The DisableNtmsObject method disables an object.
	DisableNTMSObject(context.Context, *DisableNTMSObjectRequest) (*DisableNTMSObjectResponse, error)

	// The EnableNtmsObject method enables an object.
	EnableNTMSObject(context.Context, *EnableNTMSObjectRequest) (*EnableNTMSObjectResponse, error)
}

func RegisterObjectManagement1Server(conn dcerpc.Conn, o ObjectManagement1Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewObjectManagement1ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ObjectManagement1SyntaxV0_0))...)
}

func NewObjectManagement1ServerHandle(o ObjectManagement1Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ObjectManagement1ServerHandle(ctx, o, opNum, r)
	}
}

func ObjectManagement1ServerHandle(ctx context.Context, o ObjectManagement1Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
type UnimplementedObjectManagement1Server struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedObjectManagement1Server) GetNTMSObjectSecurity(context.Context, *GetNTMSObjectSecurityRequest) (*GetNTMSObjectSecurityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedObjectManagement1Server) SetNTMSObjectSecurity(context.Context, *SetNTMSObjectSecurityRequest) (*SetNTMSObjectSecurityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedObjectManagement1Server) GetNTMSObjectAttributeA(context.Context, *GetNTMSObjectAttributeARequest) (*GetNTMSObjectAttributeAResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedObjectManagement1Server) GetNTMSObjectAttributeW(context.Context, *GetNTMSObjectAttributeWRequest) (*GetNTMSObjectAttributeWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedObjectManagement1Server) SetNTMSObjectAttributeA(context.Context, *SetNTMSObjectAttributeARequest) (*SetNTMSObjectAttributeAResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedObjectManagement1Server) SetNTMSObjectAttributeW(context.Context, *SetNTMSObjectAttributeWRequest) (*SetNTMSObjectAttributeWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedObjectManagement1Server) EnumerateNTMSObject(context.Context, *EnumerateNTMSObjectRequest) (*EnumerateNTMSObjectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedObjectManagement1Server) DisableNTMSObject(context.Context, *DisableNTMSObjectRequest) (*DisableNTMSObjectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedObjectManagement1Server) EnableNTMSObject(context.Context, *EnableNTMSObjectRequest) (*EnableNTMSObjectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ObjectManagement1Server = (*UnimplementedObjectManagement1Server)(nil)
