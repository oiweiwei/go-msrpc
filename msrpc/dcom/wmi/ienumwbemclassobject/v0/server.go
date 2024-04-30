package ienumwbemclassobject

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

// IEnumWbemClassObject server interface.
type EnumClassObjectServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// When the IEnumWbemClassObject::Reset method is invoked, the server MUST reset the
	// enumeration sequence to the beginning of the collection of CIM objects.
	//
	// This method has no parameters.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method. If the IEnumWbemClassObject::Reset
	// method is invoked on an enumerator that does not support reset capability, the server
	// MUST return WBEM_E_INVALID_OPERATION.
	Reset(context.Context, *ResetRequest) (*ResetResponse, error)

	// Next operation.
	Next(context.Context, *NextRequest) (*NextResponse, error)

	// The IEnumWbemClassObject::NextAsync method is the asynchronous version of the IEnumWbemClassObject::Next
	// method. It provides controlled asynchronous retrieval of CIM objects to a sink. The
	// server MUST asynchronously get one or more CIM objects, starting at the current position
	// in an enumeration, and MUST move the current position by the number of CIM objects.
	// When IEnumWbemClassObject is created, the current position MUST be set on the first
	// CIM object of the collection. The order of the CIM objects that are stored in the
	// enumerator is arbitrary.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	NextAsync(context.Context, *NextAsyncRequest) (*NextAsyncResponse, error)

	// The IEnumWbemClassObject::Clone method makes a logical copy of the entire enumerator.
	// The cloned enumerator MUST have the same current position as the source enumerator.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	Clone(context.Context, *CloneRequest) (*CloneResponse, error)

	// When the IEnumWbemClassObject::Skip method is invoked, the server MUST move the current
	// position in an enumeration ahead by a specified number of CIM objects.
	//
	// The IEnumWbemClassObject::Skip method opnum equals 7.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	Skip(context.Context, *SkipRequest) (*SkipResponse, error)
}

func RegisterEnumClassObjectServer(conn dcerpc.Conn, o EnumClassObjectServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewEnumClassObjectServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(EnumClassObjectSyntaxV0_0))...)
}

func NewEnumClassObjectServerHandle(o EnumClassObjectServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return EnumClassObjectServerHandle(ctx, o, opNum, r)
	}
}

func EnumClassObjectServerHandle(ctx context.Context, o EnumClassObjectServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Reset
		in := &ResetRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Reset(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // Next
		in := &NextRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Next(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // NextAsync
		in := &NextAsyncRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.NextAsync(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // Clone
		in := &CloneRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Clone(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // Skip
		in := &SkipRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Skip(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
