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
type ObjectInfo1Server interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The GetNtmsServerObjectInformationA method retrieves information about an object,
	// as a sequence of ASCII characters.
	GetNTMSServerObjectInformationA(context.Context, *GetNTMSServerObjectInformationARequest) (*GetNTMSServerObjectInformationAResponse, error)

	// The GetNtmsServerObjectInformationW method retrieves information about an object,
	// as a sequence of Unicode characters.
	GetNTMSServerObjectInformationW(context.Context, *GetNTMSServerObjectInformationWRequest) (*GetNTMSServerObjectInformationWResponse, error)

	// The SetNtmsObjectInformationA method changes the information of an object, with strings
	// encoded using ASCII.
	SetNTMSObjectInformationA(context.Context, *SetNTMSObjectInformationARequest) (*SetNTMSObjectInformationAResponse, error)

	// The SetNtmsObjectInformationW method changes the information of an object, with strings
	// encoded using Unicode.
	SetNTMSObjectInformationW(context.Context, *SetNTMSObjectInformationWRequest) (*SetNTMSObjectInformationWResponse, error)

	// The CreateNtmsMediaA method creates a new offline medium for a media pool, with strings
	// encoded using ASCII.
	CreateNTMSMediaA(context.Context, *CreateNTMSMediaARequest) (*CreateNTMSMediaAResponse, error)

	// The CreateNtmsMediaW method creates a new offline medium for a media pool, with strings
	// encoded using Unicode.
	CreateNTMSMediaW(context.Context, *CreateNTMSMediaWRequest) (*CreateNTMSMediaWResponse, error)
}

func RegisterObjectInfo1Server(conn dcerpc.Conn, o ObjectInfo1Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewObjectInfo1ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ObjectInfo1SyntaxV0_0))...)
}

func NewObjectInfo1ServerHandle(o ObjectInfo1Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ObjectInfo1ServerHandle(ctx, o, opNum, r)
	}
}

func ObjectInfo1ServerHandle(ctx context.Context, o ObjectInfo1Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
type UnimplementedObjectInfo1Server struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedObjectInfo1Server) GetNTMSServerObjectInformationA(context.Context, *GetNTMSServerObjectInformationARequest) (*GetNTMSServerObjectInformationAResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedObjectInfo1Server) GetNTMSServerObjectInformationW(context.Context, *GetNTMSServerObjectInformationWRequest) (*GetNTMSServerObjectInformationWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedObjectInfo1Server) SetNTMSObjectInformationA(context.Context, *SetNTMSObjectInformationARequest) (*SetNTMSObjectInformationAResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedObjectInfo1Server) SetNTMSObjectInformationW(context.Context, *SetNTMSObjectInformationWRequest) (*SetNTMSObjectInformationWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedObjectInfo1Server) CreateNTMSMediaA(context.Context, *CreateNTMSMediaARequest) (*CreateNTMSMediaAResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedObjectInfo1Server) CreateNTMSMediaW(context.Context, *CreateNTMSMediaWRequest) (*CreateNTMSMediaWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ObjectInfo1Server = (*UnimplementedObjectInfo1Server)(nil)
