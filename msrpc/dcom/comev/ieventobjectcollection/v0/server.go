package ieventobjectcollection

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

// IEventObjectCollection server interface.
type EventObjectCollectionServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// _NewEnum operation.
	Get_NewEnum(context.Context, *Get_NewEnumRequest) (*Get_NewEnumResponse, error)

	// The get_Item method gets the item in the collection with a specified ID.
	//
	// Return Values: An HRESULT that specifies success or failure. All success codes MUST
	// be treated the same, and all failure codes MUST be treated the same.
	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)

	// The get_NewEnum method gets an IEnumEventObject-based object for enumerating the
	// underlying collection.
	//
	// The get__NewEnum method gets a DCOM object for enumerating the underlying collection.
	//
	// Return Values: An HRESULT specifying success or failure. All success codes MUST be
	// treated the same, and all failure codes MUST be treated the same.
	GetNewEnum(context.Context, *GetNewEnumRequest) (*GetNewEnumResponse, error)

	// The get_Count method gets the count of the number of items in the collection contained
	// by the enumerator.
	//
	// Return Values: An HRESULT specifying success or failure. All success codes MUST be
	// treated the same, and all failure codes MUST be treated the same.
	GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error)

	// The Add method adds an item to the underlying collection of the enumerator; if the
	// item is already present in the collection, it is replaced by the one being passed
	// to this method.
	//
	// Return Values: An HRESULT specifying success or failure. All success codes MUST be
	// treated the same, and all failure codes MUST be treated the same.
	Add(context.Context, *AddRequest) (*AddResponse, error)

	// Remove operation.
	Remove(context.Context, *RemoveRequest) (*RemoveResponse, error)
}

func RegisterEventObjectCollectionServer(conn dcerpc.Conn, o EventObjectCollectionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewEventObjectCollectionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(EventObjectCollectionSyntaxV0_0))...)
}

func NewEventObjectCollectionServerHandle(o EventObjectCollectionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return EventObjectCollectionServerHandle(ctx, o, opNum, r)
	}
}

func EventObjectCollectionServerHandle(ctx context.Context, o EventObjectCollectionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // _NewEnum
		in := &Get_NewEnumRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Get_NewEnum(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // Item
		in := &GetItemRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetItem(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // NewEnum
		in := &GetNewEnumRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNewEnum(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // Count
		in := &GetCountRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetCount(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // Add
		in := &AddRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Add(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // Remove
		in := &RemoveRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Remove(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
