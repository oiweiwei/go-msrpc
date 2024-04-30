package iapphostsectiongroup

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

// IAppHostSectionGroup server interface.
type AppHostSectionGroupServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Count operation.
	GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error)

	// Item operation.
	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)

	// Sections operation.
	GetSections(context.Context, *GetSectionsRequest) (*GetSectionsResponse, error)

	// AddSectionGroup operation.
	AddSectionGroup(context.Context, *AddSectionGroupRequest) (*AddSectionGroupResponse, error)

	// DeleteSectionGroup operation.
	DeleteSectionGroup(context.Context, *DeleteSectionGroupRequest) (*DeleteSectionGroupResponse, error)

	// Name operation.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// Type operation.
	GetType(context.Context, *GetTypeRequest) (*GetTypeResponse, error)

	// Type operation.
	SetType(context.Context, *SetTypeRequest) (*SetTypeResponse, error)
}

func RegisterAppHostSectionGroupServer(conn dcerpc.Conn, o AppHostSectionGroupServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostSectionGroupServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostSectionGroupSyntaxV0_0))...)
}

func NewAppHostSectionGroupServerHandle(o AppHostSectionGroupServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostSectionGroupServerHandle(ctx, o, opNum, r)
	}
}

func AppHostSectionGroupServerHandle(ctx context.Context, o AppHostSectionGroupServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Count
		in := &GetCountRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetCount(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // Item
		in := &GetItemRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetItem(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // Sections
		in := &GetSectionsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSections(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // AddSectionGroup
		in := &AddSectionGroupRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddSectionGroup(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // DeleteSectionGroup
		in := &DeleteSectionGroupRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteSectionGroup(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // Name
		in := &GetNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // Type
		in := &GetTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // Type
		in := &SetTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
