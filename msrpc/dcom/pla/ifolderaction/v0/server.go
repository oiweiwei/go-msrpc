package ifolderaction

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

// IFolderAction server interface.
type FolderActionServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Age operation.
	GetAge(context.Context, *GetAgeRequest) (*GetAgeResponse, error)

	// Age operation.
	SetAge(context.Context, *SetAgeRequest) (*SetAgeResponse, error)

	// Size operation.
	GetSize(context.Context, *GetSizeRequest) (*GetSizeResponse, error)

	// Size operation.
	SetSize(context.Context, *SetSizeRequest) (*SetSizeResponse, error)

	// Actions operation.
	GetActions(context.Context, *GetActionsRequest) (*GetActionsResponse, error)

	// Actions operation.
	SetActions(context.Context, *SetActionsRequest) (*SetActionsResponse, error)

	// SendCabTo operation.
	GetSendCabTo(context.Context, *GetSendCabToRequest) (*GetSendCabToResponse, error)

	// SendCabTo operation.
	SetSendCabTo(context.Context, *SetSendCabToRequest) (*SetSendCabToResponse, error)
}

func RegisterFolderActionServer(conn dcerpc.Conn, o FolderActionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewFolderActionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(FolderActionSyntaxV0_0))...)
}

func NewFolderActionServerHandle(o FolderActionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return FolderActionServerHandle(ctx, o, opNum, r)
	}
}

func FolderActionServerHandle(ctx context.Context, o FolderActionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Age
		in := &GetAgeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAge(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // Age
		in := &SetAgeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetAge(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // Size
		in := &GetSizeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSize(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // Size
		in := &SetSizeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSize(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // Actions
		in := &GetActionsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetActions(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // Actions
		in := &SetActionsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetActions(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // SendCabTo
		in := &GetSendCabToRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSendCabTo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // SendCabTo
		in := &SetSendCabToRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSendCabTo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
