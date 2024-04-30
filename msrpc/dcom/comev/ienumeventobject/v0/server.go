package ienumeventobject

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

// IEnumEventObject server interface.
type EnumEventObjectServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The Clone method clones the underlying collection into another DCOM object implementing
	// the IEnumEventObject interface.
	//
	// Return Values: An HRESULT specifying success or failure. All success codes MUST be
	// treated the same, and all failure codes MUST be treated the same.
	Clone(context.Context, *CloneRequest) (*CloneResponse, error)

	// The Next method gets up to a specified number of items from the collection, if they
	// are available, starting at the current enumerator position.
	//
	// Return Values: An HRESULT specifying success or failure. All success codes other
	// than S_FALSE MUST be treated the same, and all failure codes MUST be treated the
	// same.
	//
	//	+--------------------+------------------------+
	//	|       RETURN       |                        |
	//	|     VALUE/CODE     |      DESCRIPTION       |
	//	|                    |                        |
	//	+--------------------+------------------------+
	//	+--------------------+------------------------+
	//	| 0x00000001 S_FALSE | End of the collection. |
	//	+--------------------+------------------------+
	Next(context.Context, *NextRequest) (*NextResponse, error)

	// The Reset method resets the enumerator back to the first element in the collection.
	//
	// This method has no parameters.
	//
	// Return Values: An HRESULT specifying success or failure. All success codes other
	// than S_FALSE MUST be treated the same, and all failure codes MUST be treated the
	// same.
	//
	//	+--------------------+-------------------------------------------------------------------------------+
	//	|       RETURN       |                                                                               |
	//	|     VALUE/CODE     |                                  DESCRIPTION                                  |
	//	|                    |                                                                               |
	//	+--------------------+-------------------------------------------------------------------------------+
	//	+--------------------+-------------------------------------------------------------------------------+
	//	| 0x00000001 S_FALSE | The enumeration sequence was reset, but there are no items in the enumerator. |
	//	+--------------------+-------------------------------------------------------------------------------+
	Reset(context.Context, *ResetRequest) (*ResetResponse, error)

	// The Skip method skips ahead in the collection by the number of elements specified.
	//
	// Return Values: An HRESULT specifying success or failure. All success codes other
	// than S_FALSE MUST be treated the same, and all failure codes MUST be treated the
	// same.
	//
	//	+--------------------+--------------------------------------------------------------------------+
	//	|       RETURN       |                                                                          |
	//	|     VALUE/CODE     |                               DESCRIPTION                                |
	//	|                    |                                                                          |
	//	+--------------------+--------------------------------------------------------------------------+
	//	+--------------------+--------------------------------------------------------------------------+
	//	| 0x00000001 S_FALSE | The number of elements skipped was not the same as the number requested. |
	//	+--------------------+--------------------------------------------------------------------------+
	Skip(context.Context, *SkipRequest) (*SkipResponse, error)
}

func RegisterEnumEventObjectServer(conn dcerpc.Conn, o EnumEventObjectServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewEnumEventObjectServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(EnumEventObjectSyntaxV0_0))...)
}

func NewEnumEventObjectServerHandle(o EnumEventObjectServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return EnumEventObjectServerHandle(ctx, o, opNum, r)
	}
}

func EnumEventObjectServerHandle(ctx context.Context, o EnumEventObjectServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Clone
		in := &CloneRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Clone(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // Next
		in := &NextRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Next(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // Reset
		in := &ResetRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Reset(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // Skip
		in := &SkipRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Skip(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
