package ienumvdsobject

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

// IEnumVdsObject server interface.
type EnumObjectServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The Next method returns a specified number of objects in the enumeration. It begins
	// from the current point.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Next(context.Context, *NextRequest) (*NextResponse, error)

	// The Skip method skips a specified number of objects in the enumeration.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Skip(context.Context, *SkipRequest) (*SkipResponse, error)

	// The Reset method resets the enumerator to the beginning of the collection.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Reset(context.Context, *ResetRequest) (*ResetResponse, error)

	// The Clone method creates a new enumeration that has the same state as the current
	// enumeration.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Clone(context.Context, *CloneRequest) (*CloneResponse, error)
}

func RegisterEnumObjectServer(conn dcerpc.Conn, o EnumObjectServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewEnumObjectServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(EnumObjectSyntaxV0_0))...)
}

func NewEnumObjectServerHandle(o EnumObjectServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return EnumObjectServerHandle(ctx, o, opNum, r)
	}
}

func EnumObjectServerHandle(ctx context.Context, o EnumObjectServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Next
		in := &NextRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Next(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // Skip
		in := &SkipRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Skip(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // Reset
		in := &ResetRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Reset(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // Clone
		in := &CloneRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Clone(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
