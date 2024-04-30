package ipchcollection

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

// IPCHCollection server interface.
type PCHCollectionServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// _NewEnum operation.
	Get_NewEnum(context.Context, *Get_NewEnumRequest) (*Get_NewEnumResponse, error)

	// The Item method retrieves an element.
	//
	// Return Values: A signed 32-bit value indicating return status. This method MUST return
	// zero to indicate success, or an HRESULT error value (as specified in [MS-ERREF] section
	// 2.1.1) to indicate failure.
	//
	//	+--------------------------+----------------------------------------------------+
	//	|          RETURN          |                                                    |
	//	|        VALUE/CODE        |                    DESCRIPTION                     |
	//	|                          |                                                    |
	//	+--------------------------+----------------------------------------------------+
	//	+--------------------------+----------------------------------------------------+
	//	| 0x00000000 S_OK          | The call was successful.                           |
	//	+--------------------------+----------------------------------------------------+
	//	| 0x80004003 E_POINTER     | The method failed due to an invalid pointer.       |
	//	+--------------------------+----------------------------------------------------+
	//	| 0x8007000E E_OUTOFMEMORY | The method was unable to allocate required memory. |
	//	+--------------------------+----------------------------------------------------+
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol [MS-RPCE].
	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)

	// The Count method retrieves the number of elements in the collection.
	//
	// Return Values: A signed 32-bit value indicating return status. This method MUST return
	// zero to indicate success, or an HRESULT error value (as specified in [MS-ERREF] section
	// 2.1.1) to indicate failure.
	//
	//	+--------------------------+----------------------------------------------------+
	//	|          RETURN          |                                                    |
	//	|        VALUE/CODE        |                    DESCRIPTION                     |
	//	|                          |                                                    |
	//	+--------------------------+----------------------------------------------------+
	//	+--------------------------+----------------------------------------------------+
	//	| 0x00000000 S_OK          | The call was successful.                           |
	//	+--------------------------+----------------------------------------------------+
	//	| 0x80004003 E_POINTER     | The method failed due to an invalid pointer.       |
	//	+--------------------------+----------------------------------------------------+
	//	| 0x8007000E E_OUTOFMEMORY | The method was unable to allocate required memory. |
	//	+--------------------------+----------------------------------------------------+
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol [MS-RPCE].
	GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error)
}

func RegisterPCHCollectionServer(conn dcerpc.Conn, o PCHCollectionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewPCHCollectionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(PCHCollectionSyntaxV0_0))...)
}

func NewPCHCollectionServerHandle(o PCHCollectionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return PCHCollectionServerHandle(ctx, o, opNum, r)
	}
}

func PCHCollectionServerHandle(ctx context.Context, o PCHCollectionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
	case 9: // Count
		in := &GetCountRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetCount(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
