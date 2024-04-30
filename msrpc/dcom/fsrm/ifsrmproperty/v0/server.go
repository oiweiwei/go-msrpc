package ifsrmproperty

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

// IFsrmProperty server interface.
type PropertyServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Name operation.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// Value operation.
	GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error)

	// The Sources (get) method retrieves the sources for this Property Definition Instance
	// and returns S_OK upon successful completion.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	GetSources(context.Context, *GetSourcesRequest) (*GetSourcesResponse, error)

	// The PropertyFlags (get) retrieves the property definition instance.flags for the
	// Property Definition Instance in an FsrmPropertyFlags (section 2.2.2.6.1.1) enumeration
	// and returns S_OK upon successful completion.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+------------------------------+
	//	|         RETURN          |                              |
	//	|       VALUE/CODE        |         DESCRIPTION          |
	//	|                         |                              |
	//	+-------------------------+------------------------------+
	//	+-------------------------+------------------------------+
	//	| 0x80070057 E_INVALIDARG | The flags parameter is NULL. |
	//	+-------------------------+------------------------------+
	GetPropertyFlags(context.Context, *GetPropertyFlagsRequest) (*GetPropertyFlagsResponse, error)
}

func RegisterPropertyServer(conn dcerpc.Conn, o PropertyServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewPropertyServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(PropertySyntaxV0_0))...)
}

func NewPropertyServerHandle(o PropertyServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return PropertyServerHandle(ctx, o, opNum, r)
	}
}

func PropertyServerHandle(ctx context.Context, o PropertyServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Name
		in := &GetNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // Value
		in := &GetValueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetValue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // Sources
		in := &GetSourcesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSources(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // PropertyFlags
		in := &GetPropertyFlagsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPropertyFlags(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
