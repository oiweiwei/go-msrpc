package ieventsystem

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

// IEventSystem server interface.
type EventSystemServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The Query method is called by a client to query a collection for a collection of
	// event classes or subscriptions.
	//
	// Return Values: An HRESULT that specifies success or failure. All success codes MUST
	// be treated the same, and all failure codes other than EVENT_E_QUERYSYNTAX and EVENT_E_QUERYFIELD
	// MUST be treated the same.
	//
	//	+--------------------------------+------------------------------------------------------------------+
	//	|             RETURN             |                                                                  |
	//	|           VALUE/CODE           |                           DESCRIPTION                            |
	//	|                                |                                                                  |
	//	+--------------------------------+------------------------------------------------------------------+
	//	+--------------------------------+------------------------------------------------------------------+
	//	| 0x80040203 EVENT_E_QUERYSYNTAX | A syntax error occurred while trying to evaluate a query string. |
	//	+--------------------------------+------------------------------------------------------------------+
	//	| 0x80040204 EVENT_E_QUERYFIELD  | An invalid field name was used in a query string.                |
	//	+--------------------------------+------------------------------------------------------------------+
	Query(context.Context, *QueryRequest) (*QueryResponse, error)

	// The Store method is called by a client to store either an event class or a subscription.
	//
	// Return Values: An HRESULT that specifies success or failure. All success codes MUST
	// be treated the same, and all failure codes other than EVENT_E_INVALID_PER_USER_SID
	// MUST be treated the same.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                 RETURN                  |                                                                                  |
	//	|               VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80040207 EVENT_E_INVALID_PER_USER_SID | The owner SID, as defined in [MS-DTYP] section 2.4.2, on a per-user subscription |
	//	|                                         | does not exist.                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	Store(context.Context, *StoreRequest) (*StoreResponse, error)

	// Remove operation.
	Remove(context.Context, *RemoveRequest) (*RemoveResponse, error)

	// The get_EventObjectChangeEventClassID method extracts the server-specific EventClassID
	// for server-specific event class or subscription change notifications.
	//
	// Return Values: An HRESULT specifying success or failure. All success codes MUST be
	// treated the same, and all failure codes MUST be treated the same.
	GetEventObjectChangeEventClassID(context.Context, *GetEventObjectChangeEventClassIDRequest) (*GetEventObjectChangeEventClassIDResponse, error)

	// The QueryS method is called by the client to query a specific event class or subscription.
	//
	// Return Values: An HRESULT that specifies success or failure. All success codes MUST
	// be treated the same, and all failure codes other than EVENT_E_QUERYSYNTAX and EVENT_E_QUERYFIELD
	// MUST be treated the same.
	//
	//	+--------------------------------+------------------------------------------------------------------+
	//	|             RETURN             |                                                                  |
	//	|           VALUE/CODE           |                           DESCRIPTION                            |
	//	|                                |                                                                  |
	//	+--------------------------------+------------------------------------------------------------------+
	//	+--------------------------------+------------------------------------------------------------------+
	//	| 0x80040203 EVENT_E_QUERYSYNTAX | A syntax error occurred while trying to evaluate a query string. |
	//	+--------------------------------+------------------------------------------------------------------+
	//	| 0x80040204 EVENT_E_QUERYFIELD  | An invalid field name was used in a query string.                |
	//	+--------------------------------+------------------------------------------------------------------+
	QueryS(context.Context, *QuerySRequest) (*QuerySResponse, error)

	// The RemoveS method is called by the client to remove an event class or subscription.
	//
	// Return Values: An HRESULT specifying success or failure. All success codes MUST be
	// treated the same, and all failure codes other than EVENT_E_QUERYSYNTAX, EVENT_E_QUERYFIELD,
	// and EVENT_E_NOT_ALL_REMOVED MUST be treated the same.
	//
	//	+------------------------------------+------------------------------------------------------------------+
	//	|               RETURN               |                                                                  |
	//	|             VALUE/CODE             |                           DESCRIPTION                            |
	//	|                                    |                                                                  |
	//	+------------------------------------+------------------------------------------------------------------+
	//	+------------------------------------+------------------------------------------------------------------+
	//	| 0x80040203 EVENT_E_QUERYSYNTAX     | A syntax error occurred while trying to evaluate a query string. |
	//	+------------------------------------+------------------------------------------------------------------+
	//	| 0x80040204 EVENT_E_QUERYFIELD      | An invalid field name was used in a query string.                |
	//	+------------------------------------+------------------------------------------------------------------+
	//	| 0x8004020B EVENT_E_NOT_ALL_REMOVED | Not all of the requested objects could be removed.               |
	//	+------------------------------------+------------------------------------------------------------------+
	RemoveS(context.Context, *RemoveSRequest) (*RemoveSResponse, error)
}

func RegisterEventSystemServer(conn dcerpc.Conn, o EventSystemServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewEventSystemServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(EventSystemSyntaxV0_0))...)
}

func NewEventSystemServerHandle(o EventSystemServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return EventSystemServerHandle(ctx, o, opNum, r)
	}
}

func EventSystemServerHandle(ctx context.Context, o EventSystemServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Query
		in := &QueryRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Query(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // Store
		in := &StoreRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Store(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // Remove
		in := &RemoveRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Remove(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // EventObjectChangeEventClassID
		in := &GetEventObjectChangeEventClassIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetEventObjectChangeEventClassID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // QueryS
		in := &QuerySRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryS(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // RemoveS
		in := &RemoveSRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveS(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
