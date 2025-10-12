package conv

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
)

// conv server interface.
type ConvServer interface {

	// conv_who_are_you operation.
	WhoAreYou(context.Context, *WhoAreYouRequest) (*WhoAreYouResponse, error)

	// conv_who_are_you2 operation.
	WhoAreYou2(context.Context, *WhoAreYou2Request) (*WhoAreYou2Response, error)

	// conv_are_you_there operation.
	AreYouThere(context.Context, *AreYouThereRequest) (*AreYouThereResponse, error)

	// conv_who_are_you_auth operation.
	WhoAreYouAuth(context.Context, *WhoAreYouAuthRequest) (*WhoAreYouAuthResponse, error)

	// conv_who_are_you_auth_more operation.
	WhoAreYouAuthMore(context.Context, *WhoAreYouAuthMoreRequest) (*WhoAreYouAuthMoreResponse, error)
}

func RegisterConvServer(conn dcerpc.Conn, o ConvServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewConvServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ConvSyntaxV3_0))...)
}

func NewConvServerHandle(o ConvServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ConvServerHandle(ctx, o, opNum, r)
	}
}

func ConvServerHandle(ctx context.Context, o ConvServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // conv_who_are_you
		op := &xxx_WhoAreYouOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &WhoAreYouRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.WhoAreYou(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // conv_who_are_you2
		op := &xxx_WhoAreYou2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &WhoAreYou2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.WhoAreYou2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // conv_are_you_there
		op := &xxx_AreYouThereOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AreYouThereRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AreYouThere(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // conv_who_are_you_auth
		op := &xxx_WhoAreYouAuthOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &WhoAreYouAuthRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.WhoAreYouAuth(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // conv_who_are_you_auth_more
		op := &xxx_WhoAreYouAuthMoreOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &WhoAreYouAuthMoreRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.WhoAreYouAuthMore(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented conv
type UnimplementedConvServer struct {
}

func (UnimplementedConvServer) WhoAreYou(context.Context, *WhoAreYouRequest) (*WhoAreYouResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConvServer) WhoAreYou2(context.Context, *WhoAreYou2Request) (*WhoAreYou2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConvServer) AreYouThere(context.Context, *AreYouThereRequest) (*AreYouThereResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConvServer) WhoAreYouAuth(context.Context, *WhoAreYouAuthRequest) (*WhoAreYouAuthResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConvServer) WhoAreYouAuthMore(context.Context, *WhoAreYouAuthMoreRequest) (*WhoAreYouAuthMoreResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ConvServer = (*UnimplementedConvServer)(nil)
