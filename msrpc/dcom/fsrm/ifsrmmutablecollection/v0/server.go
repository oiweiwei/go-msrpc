package ifsrmmutablecollection

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
		op := &xxx_AddOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Add(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // Remove
		op := &xxx_RemoveOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Remove(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // RemoveById
		op := &xxx_RemoveByIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveByIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveByID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // Clone
		op := &xxx_CloneOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloneRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Clone(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmMutableCollection
type UnimplementedMutableCollectionServer struct {
	ifsrmcollection.UnimplementedCollectionServer
}

func (UnimplementedMutableCollectionServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMutableCollectionServer) Remove(context.Context, *RemoveRequest) (*RemoveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMutableCollectionServer) RemoveByID(context.Context, *RemoveByIDRequest) (*RemoveByIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMutableCollectionServer) Clone(context.Context, *CloneRequest) (*CloneResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ MutableCollectionServer = (*UnimplementedMutableCollectionServer)(nil)
