package ifsrmmutablecollection

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	ifsrmcollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmcollection/v0"
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
	_ = ifsrmcollection.GoPackage
)

// IFsrmMutableCollection server interface.
type MutableCollectionServer interface {

	// IFsrmCollection base class.
	ifsrmcollection.CollectionServer

	// The Add method adds the specified object to the collection of Objects Being Enumerated.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN              |                                                                                  |
	//	|            VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                  |                                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045303 FSRM_E_ALREADY_EXISTS | The object pointed to by the IDispatch pointer that is contained in the VARIANT  |
	//	|                                  | structure already exists in the collection of Objects Being Enumerated.          |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG          | The item parameter is not a value type.                                          |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	Add(context.Context, *AddRequest) (*AddResponse, error)

	// The Remove method removes an object from the collection of Objects Being Enumerated.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	|         RETURN          |                                                                                  |
	//	|       VALUE/CODE        |                                   DESCRIPTION                                    |
	//	|                         |                                                                                  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | The index is out of range; the index is less than one or greater than the size   |
	//	|                         | of the collection.                                                               |
	//	+-------------------------+----------------------------------------------------------------------------------+
	Remove(context.Context, *RemoveRequest) (*RemoveResponse, error)

	// The RemoveById method removes from the collection of Objects Being Enumerated the
	// object whose ID matches the specified id.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-----------------------------+------------------------------------------------------------------+
	//	|           RETURN            |                                                                  |
	//	|         VALUE/CODE          |                           DESCRIPTION                            |
	//	|                             |                                                                  |
	//	+-----------------------------+------------------------------------------------------------------+
	//	+-----------------------------+------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND | An object with the specified ID was not found in the collection. |
	//	+-----------------------------+------------------------------------------------------------------+
	RemoveByID(context.Context, *RemoveByIDRequest) (*RemoveByIDResponse, error)

	// The Clone method returns a copy of the collection of Objects Being Enumerated.
	//
	// Return Values: The method MUST return a nonzero error code. Upon receiving this message,
	// the server MUST return E_NOTIMPL.
	Clone(context.Context, *CloneRequest) (*CloneResponse, error)
}

func RegisterMutableCollectionServer(conn dcerpc.Conn, o MutableCollectionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewMutableCollectionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(MutableCollectionSyntaxV0_0))...)
}

func NewMutableCollectionServerHandle(o MutableCollectionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return MutableCollectionServerHandle(ctx, o, opNum, r)
	}
}

func MutableCollectionServerHandle(ctx context.Context, o MutableCollectionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 14 {
		// IFsrmCollection base method.
		return ifsrmcollection.CollectionServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 14: // Add
		in := &AddRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Add(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // Remove
		in := &RemoveRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Remove(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // RemoveById
		in := &RemoveByIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveByID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // Clone
		in := &CloneRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Clone(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
