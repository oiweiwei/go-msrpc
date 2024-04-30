package imanagedobject

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

// IManagedObject server interface.
type ManagedObjectServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The GetSerializedBuffer method converts the given managed object to a binary-formatted
	// string representation that can be used to create a managed object.
	//
	// Return Values: The method MUST return a positive value or 0 to indicate successful
	// completion or a negative value to indicate failure.
	//
	//	+--------------------------+-------------+
	//	|          RETURN          |             |
	//	|        VALUE/CODE        | DESCRIPTION |
	//	|                          |             |
	//	+--------------------------+-------------+
	//	+--------------------------+-------------+
	//	| 0x00000000 ERROR_SUCCESS | Success.    |
	//	+--------------------------+-------------+
	//
	// Exceptions Thrown: No exceptions are thrown from this method beyond those thrown
	// by the underlying RPC protocol.
	GetSerializedBuffer(context.Context, *GetSerializedBufferRequest) (*GetSerializedBufferResponse, error)

	// The IManagedObject::GetObjectIdentity method is used by a CLR instance to determine
	// whether a COM object entering the system is really a managed object that originated
	// in this CLR instance and within the current process division.
	//
	// Return Values: The method MUST return a positive value or 0 to indicate successful
	// completion or a negative value to indicate failure.
	//
	//	+--------------------------+-------------+
	//	|          RETURN          |             |
	//	|        VALUE/CODE        | DESCRIPTION |
	//	|                          |             |
	//	+--------------------------+-------------+
	//	+--------------------------+-------------+
	//	| 0x00000000 ERROR_SUCCESS | Success     |
	//	+--------------------------+-------------+
	//
	// Exceptions Thrown: No exceptions are thrown from this method beyond those thrown
	// by the underlying RPC protocol.
	GetObjectIdentity(context.Context, *GetObjectIdentityRequest) (*GetObjectIdentityResponse, error)
}

func RegisterManagedObjectServer(conn dcerpc.Conn, o ManagedObjectServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewManagedObjectServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ManagedObjectSyntaxV0_0))...)
}

func NewManagedObjectServerHandle(o ManagedObjectServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ManagedObjectServerHandle(ctx, o, opNum, r)
	}
}

func ManagedObjectServerHandle(ctx context.Context, o ManagedObjectServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetSerializedBuffer
		in := &GetSerializedBufferRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSerializedBuffer(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // GetObjectIdentity
		in := &GetObjectIdentityRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetObjectIdentity(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
