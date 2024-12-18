package ieventsystem

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
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
	_ = dcom.GoPackage
	_ = idispatch.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/comev"
)

var (
	// IEventSystem interface identifier 4e14fb9f-2e22-11d1-9964-00c04fbbb345
	EventSystemIID = &dcom.IID{Data1: 0x4e14fb9f, Data2: 0x2e22, Data3: 0x11d1, Data4: []byte{0x99, 0x64, 0x00, 0xc0, 0x4f, 0xbb, 0xb3, 0x45}}
	// Syntax UUID
	EventSystemSyntaxUUID = &uuid.UUID{TimeLow: 0x4e14fb9f, TimeMid: 0x2e22, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x99, ClockSeqLow: 0x64, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xbb, 0xb3, 0x45}}
	// Syntax ID
	EventSystemSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: EventSystemSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IEventSystem interface.
type EventSystemClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

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
	Query(context.Context, *QueryRequest, ...dcerpc.CallOption) (*QueryResponse, error)

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
	Store(context.Context, *StoreRequest, ...dcerpc.CallOption) (*StoreResponse, error)

	// Remove operation.
	Remove(context.Context, *RemoveRequest, ...dcerpc.CallOption) (*RemoveResponse, error)

	// The get_EventObjectChangeEventClassID method extracts the server-specific EventClassID
	// for server-specific event class or subscription change notifications.
	//
	// Return Values: An HRESULT specifying success or failure. All success codes MUST be
	// treated the same, and all failure codes MUST be treated the same.
	GetEventObjectChangeEventClassID(context.Context, *GetEventObjectChangeEventClassIDRequest, ...dcerpc.CallOption) (*GetEventObjectChangeEventClassIDResponse, error)

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
	QueryS(context.Context, *QuerySRequest, ...dcerpc.CallOption) (*QuerySResponse, error)

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
	RemoveS(context.Context, *RemoveSRequest, ...dcerpc.CallOption) (*RemoveSResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) EventSystemClient
}

type xxx_DefaultEventSystemClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultEventSystemClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultEventSystemClient) Query(ctx context.Context, in *QueryRequest, opts ...dcerpc.CallOption) (*QueryResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSystemClient) Store(ctx context.Context, in *StoreRequest, opts ...dcerpc.CallOption) (*StoreResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &StoreResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSystemClient) Remove(ctx context.Context, in *RemoveRequest, opts ...dcerpc.CallOption) (*RemoveResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoveResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSystemClient) GetEventObjectChangeEventClassID(ctx context.Context, in *GetEventObjectChangeEventClassIDRequest, opts ...dcerpc.CallOption) (*GetEventObjectChangeEventClassIDResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetEventObjectChangeEventClassIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSystemClient) QueryS(ctx context.Context, in *QuerySRequest, opts ...dcerpc.CallOption) (*QuerySResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QuerySResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSystemClient) RemoveS(ctx context.Context, in *RemoveSRequest, opts ...dcerpc.CallOption) (*RemoveSResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoveSResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSystemClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultEventSystemClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultEventSystemClient) IPID(ctx context.Context, ipid *dcom.IPID) EventSystemClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultEventSystemClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewEventSystemClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (EventSystemClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(EventSystemSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := idispatch.NewDispatchClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultEventSystemClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_QueryOperation structure represents the Query operation
type xxx_QueryOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	ProgID        *oaut.String   `idl:"name:progID" json:"prog_id"`
	QueryCriteria *oaut.String   `idl:"name:queryCriteria" json:"query_criteria"`
	ErrorIndex    int32          `idl:"name:errorIndex" json:"error_index"`
	Interface     *dcom.Unknown  `idl:"name:ppInterface" json:"interface"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryOperation) OpNum() int { return 7 }

func (o *xxx_QueryOperation) OpName() string { return "/IEventSystem/v0/Query" }

func (o *xxx_QueryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// progID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ProgID != nil {
			_ptr_progID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ProgID != nil {
					if err := o.ProgID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ProgID, _ptr_progID); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// queryCriteria {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.QueryCriteria != nil {
			_ptr_queryCriteria := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.QueryCriteria != nil {
					if err := o.QueryCriteria.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QueryCriteria, _ptr_queryCriteria); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// progID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_progID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ProgID == nil {
				o.ProgID = &oaut.String{}
			}
			if err := o.ProgID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_progID := func(ptr interface{}) { o.ProgID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ProgID, _s_progID, _ptr_progID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// queryCriteria {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_queryCriteria := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.QueryCriteria == nil {
				o.QueryCriteria = &oaut.String{}
			}
			if err := o.QueryCriteria.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_queryCriteria := func(ptr interface{}) { o.QueryCriteria = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.QueryCriteria, _s_queryCriteria, _ptr_queryCriteria); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// errorIndex {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.ErrorIndex); err != nil {
			return err
		}
	}
	// ppInterface {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IUnknown}(interface))
	{
		if o.Interface != nil {
			_ptr_ppInterface := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Interface != nil {
					if err := o.Interface.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.Unknown{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Interface, _ptr_ppInterface); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// errorIndex {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.ErrorIndex); err != nil {
			return err
		}
	}
	// ppInterface {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IUnknown}(interface))
	{
		_ptr_ppInterface := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Interface == nil {
				o.Interface = &dcom.Unknown{}
			}
			if err := o.Interface.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppInterface := func(ptr interface{}) { o.Interface = *ptr.(**dcom.Unknown) }
		if err := w.ReadPointer(&o.Interface, _s_ppInterface, _ptr_ppInterface); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryRequest structure represents the Query operation request
type QueryRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// progID: A string that identifies the type of collection. The value MUST be one of
	// the following.
	//
	//	+-------------------------------------------+----------------------------------------------------------------+
	//	|                                           |                                                                |
	//	|                   VALUE                   |                            MEANING                             |
	//	|                                           |                                                                |
	//	+-------------------------------------------+----------------------------------------------------------------+
	//	+-------------------------------------------+----------------------------------------------------------------+
	//	| "EventSystem.EventClassCollection"        | The store for event classes (as specified in section 3.1.1.1). |
	//	+-------------------------------------------+----------------------------------------------------------------+
	//	| "EventSystem.EventSubscriptionCollection" | The store for subscriptions (as specified in section 3.1.1.2). |
	//	+-------------------------------------------+----------------------------------------------------------------+
	ProgID *oaut.String `idl:"name:progID" json:"prog_id"`
	// queryCriteria: The actual query string. The syntax for this string MUST conform to
	// section 2.2.1.
	QueryCriteria *oaut.String `idl:"name:queryCriteria" json:"query_criteria"`
}

func (o *QueryRequest) xxx_ToOp(ctx context.Context) *xxx_QueryOperation {
	if o == nil {
		return &xxx_QueryOperation{}
	}
	return &xxx_QueryOperation{
		This:          o.This,
		ProgID:        o.ProgID,
		QueryCriteria: o.QueryCriteria,
	}
}

func (o *QueryRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ProgID = op.ProgID
	o.QueryCriteria = op.QueryCriteria
}
func (o *QueryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryResponse structure represents the Query operation response
type QueryResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// errorIndex: The zero-based character index in the queryCriteria parameter where an
	// error has occurred. This can occur if the syntax of the query string is incorrect,
	// in which case errorIndex specifies the index at which the problematic syntax is present
	// in the queryCriteria parameter.
	ErrorIndex int32 `idl:"name:errorIndex" json:"error_index"`
	// ppInterface: If the method returns a success HRESULT, this MUST contain an interface
	// pointer that represents the collection of the event classes or subscriptions based
	// on the criteria specified in the queryCriteria parameter.
	Interface *dcom.Unknown `idl:"name:ppInterface" json:"interface"`
	// Return: The Query return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryResponse) xxx_ToOp(ctx context.Context) *xxx_QueryOperation {
	if o == nil {
		return &xxx_QueryOperation{}
	}
	return &xxx_QueryOperation{
		That:       o.That,
		ErrorIndex: o.ErrorIndex,
		Interface:  o.Interface,
		Return:     o.Return,
	}
}

func (o *QueryResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ErrorIndex = op.ErrorIndex
	o.Interface = op.Interface
	o.Return = op.Return
}
func (o *QueryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_StoreOperation structure represents the Store operation
type xxx_StoreOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	ProgID    *oaut.String   `idl:"name:ProgID" json:"prog_id"`
	Interface *dcom.Unknown  `idl:"name:pInterface" json:"interface"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_StoreOperation) OpNum() int { return 8 }

func (o *xxx_StoreOperation) OpName() string { return "/IEventSystem/v0/Store" }

func (o *xxx_StoreOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StoreOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ProgID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ProgID != nil {
			_ptr_ProgID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ProgID != nil {
					if err := o.ProgID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ProgID, _ptr_ProgID); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pInterface {in} (1:{pointer=ref}*(1))(2:{alias=IUnknown}(interface))
	{
		if o.Interface != nil {
			_ptr_pInterface := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Interface != nil {
					if err := o.Interface.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.Unknown{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Interface, _ptr_pInterface); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StoreOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ProgID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_ProgID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ProgID == nil {
				o.ProgID = &oaut.String{}
			}
			if err := o.ProgID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ProgID := func(ptr interface{}) { o.ProgID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ProgID, _s_ProgID, _ptr_ProgID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pInterface {in} (1:{pointer=ref}*(1))(2:{alias=IUnknown}(interface))
	{
		_ptr_pInterface := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Interface == nil {
				o.Interface = &dcom.Unknown{}
			}
			if err := o.Interface.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pInterface := func(ptr interface{}) { o.Interface = *ptr.(**dcom.Unknown) }
		if err := w.ReadPointer(&o.Interface, _s_pInterface, _ptr_pInterface); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StoreOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StoreOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StoreOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// StoreRequest structure represents the Store operation request
type StoreRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ProgID: A string that uniquely identifies the kind of object that the client is trying
	// to store. It MUST be one of the following values.
	//
	//	+---------------------------------+---------------------------------------------------------------+
	//	|                                 |                                                               |
	//	|              VALUE              |                            MEANING                            |
	//	|                                 |                                                               |
	//	+---------------------------------+---------------------------------------------------------------+
	//	+---------------------------------+---------------------------------------------------------------+
	//	| "EventSystem.EventClass"        | The store for event classes, as specified in section 3.1.1.1. |
	//	+---------------------------------+---------------------------------------------------------------+
	//	| "EventSystem.EventSubscription" | The store for subscriptions, as specified in section 3.1.1.2. |
	//	+---------------------------------+---------------------------------------------------------------+
	ProgID *oaut.String `idl:"name:ProgID" json:"prog_id"`
	// pInterface: An interface pointer to a DCOM object that was created by performing
	// DCOM activation on the server by the client by using either the CLSID_EventClass
	// (as specified in section 1.9), which represents the CLSID for event class, or CLSID_Subscription
	// (as specified in section 1.9), which represents the subscriber.
	Interface *dcom.Unknown `idl:"name:pInterface" json:"interface"`
}

func (o *StoreRequest) xxx_ToOp(ctx context.Context) *xxx_StoreOperation {
	if o == nil {
		return &xxx_StoreOperation{}
	}
	return &xxx_StoreOperation{
		This:      o.This,
		ProgID:    o.ProgID,
		Interface: o.Interface,
	}
}

func (o *StoreRequest) xxx_FromOp(ctx context.Context, op *xxx_StoreOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ProgID = op.ProgID
	o.Interface = op.Interface
}
func (o *StoreRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *StoreRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StoreOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// StoreResponse structure represents the Store operation response
type StoreResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Store return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *StoreResponse) xxx_ToOp(ctx context.Context) *xxx_StoreOperation {
	if o == nil {
		return &xxx_StoreOperation{}
	}
	return &xxx_StoreOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *StoreResponse) xxx_FromOp(ctx context.Context, op *xxx_StoreOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *StoreResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *StoreResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StoreOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoveOperation structure represents the Remove operation
type xxx_RemoveOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	ProgID        *oaut.String   `idl:"name:progID" json:"prog_id"`
	QueryCriteria *oaut.String   `idl:"name:queryCriteria" json:"query_criteria"`
	ErrorIndex    int32          `idl:"name:errorIndex" json:"error_index"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoveOperation) OpNum() int { return 9 }

func (o *xxx_RemoveOperation) OpName() string { return "/IEventSystem/v0/Remove" }

func (o *xxx_RemoveOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// progID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ProgID != nil {
			_ptr_progID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ProgID != nil {
					if err := o.ProgID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ProgID, _ptr_progID); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// queryCriteria {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.QueryCriteria != nil {
			_ptr_queryCriteria := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.QueryCriteria != nil {
					if err := o.QueryCriteria.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QueryCriteria, _ptr_queryCriteria); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// progID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_progID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ProgID == nil {
				o.ProgID = &oaut.String{}
			}
			if err := o.ProgID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_progID := func(ptr interface{}) { o.ProgID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ProgID, _s_progID, _ptr_progID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// queryCriteria {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_queryCriteria := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.QueryCriteria == nil {
				o.QueryCriteria = &oaut.String{}
			}
			if err := o.QueryCriteria.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_queryCriteria := func(ptr interface{}) { o.QueryCriteria = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.QueryCriteria, _s_queryCriteria, _ptr_queryCriteria); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// errorIndex {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.ErrorIndex); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// errorIndex {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.ErrorIndex); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RemoveRequest structure represents the Remove operation request
type RemoveRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	ProgID        *oaut.String   `idl:"name:progID" json:"prog_id"`
	QueryCriteria *oaut.String   `idl:"name:queryCriteria" json:"query_criteria"`
}

func (o *RemoveRequest) xxx_ToOp(ctx context.Context) *xxx_RemoveOperation {
	if o == nil {
		return &xxx_RemoveOperation{}
	}
	return &xxx_RemoveOperation{
		This:          o.This,
		ProgID:        o.ProgID,
		QueryCriteria: o.QueryCriteria,
	}
}

func (o *RemoveRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoveOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ProgID = op.ProgID
	o.QueryCriteria = op.QueryCriteria
}
func (o *RemoveRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RemoveRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoveResponse structure represents the Remove operation response
type RemoveResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	ErrorIndex int32          `idl:"name:errorIndex" json:"error_index"`
	// Return: The Remove return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemoveResponse) xxx_ToOp(ctx context.Context) *xxx_RemoveOperation {
	if o == nil {
		return &xxx_RemoveOperation{}
	}
	return &xxx_RemoveOperation{
		That:       o.That,
		ErrorIndex: o.ErrorIndex,
		Return:     o.Return,
	}
}

func (o *RemoveResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoveOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ErrorIndex = op.ErrorIndex
	o.Return = op.Return
}
func (o *RemoveResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RemoveResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetEventObjectChangeEventClassIDOperation structure represents the EventObjectChangeEventClassID operation
type xxx_GetEventObjectChangeEventClassIDOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	EventClassID *oaut.String   `idl:"name:pbstrEventClassID" json:"event_class_id"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetEventObjectChangeEventClassIDOperation) OpNum() int { return 10 }

func (o *xxx_GetEventObjectChangeEventClassIDOperation) OpName() string {
	return "/IEventSystem/v0/EventObjectChangeEventClassID"
}

func (o *xxx_GetEventObjectChangeEventClassIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventObjectChangeEventClassIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventObjectChangeEventClassIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventObjectChangeEventClassIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventObjectChangeEventClassIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pbstrEventClassID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.EventClassID != nil {
			_ptr_pbstrEventClassID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EventClassID != nil {
					if err := o.EventClassID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EventClassID, _ptr_pbstrEventClassID); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventObjectChangeEventClassIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbstrEventClassID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrEventClassID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EventClassID == nil {
				o.EventClassID = &oaut.String{}
			}
			if err := o.EventClassID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrEventClassID := func(ptr interface{}) { o.EventClassID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.EventClassID, _s_pbstrEventClassID, _ptr_pbstrEventClassID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetEventObjectChangeEventClassIDRequest structure represents the EventObjectChangeEventClassID operation request
type GetEventObjectChangeEventClassIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetEventObjectChangeEventClassIDRequest) xxx_ToOp(ctx context.Context) *xxx_GetEventObjectChangeEventClassIDOperation {
	if o == nil {
		return &xxx_GetEventObjectChangeEventClassIDOperation{}
	}
	return &xxx_GetEventObjectChangeEventClassIDOperation{
		This: o.This,
	}
}

func (o *GetEventObjectChangeEventClassIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetEventObjectChangeEventClassIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetEventObjectChangeEventClassIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetEventObjectChangeEventClassIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEventObjectChangeEventClassIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetEventObjectChangeEventClassIDResponse structure represents the EventObjectChangeEventClassID operation response
type GetEventObjectChangeEventClassIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbstrEventClassID: If the method call returns a success HRESULT, this MUST contain
	// the returned unique identifier representing the EventClassID for the server specific
	// EventClass/Subscription change notifications. This MUST be a GUID specified as a
	// string as specified in section 2.2.3.
	EventClassID *oaut.String `idl:"name:pbstrEventClassID" json:"event_class_id"`
	// Return: The EventObjectChangeEventClassID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetEventObjectChangeEventClassIDResponse) xxx_ToOp(ctx context.Context) *xxx_GetEventObjectChangeEventClassIDOperation {
	if o == nil {
		return &xxx_GetEventObjectChangeEventClassIDOperation{}
	}
	return &xxx_GetEventObjectChangeEventClassIDOperation{
		That:         o.That,
		EventClassID: o.EventClassID,
		Return:       o.Return,
	}
}

func (o *GetEventObjectChangeEventClassIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetEventObjectChangeEventClassIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.EventClassID = op.EventClassID
	o.Return = op.Return
}
func (o *GetEventObjectChangeEventClassIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetEventObjectChangeEventClassIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEventObjectChangeEventClassIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QuerySOperation structure represents the QueryS operation
type xxx_QuerySOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	ProgID        *oaut.String   `idl:"name:progID" json:"prog_id"`
	QueryCriteria *oaut.String   `idl:"name:queryCriteria" json:"query_criteria"`
	Interface     *dcom.Unknown  `idl:"name:ppInterface" json:"interface"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_QuerySOperation) OpNum() int { return 11 }

func (o *xxx_QuerySOperation) OpName() string { return "/IEventSystem/v0/QueryS" }

func (o *xxx_QuerySOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// progID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ProgID != nil {
			_ptr_progID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ProgID != nil {
					if err := o.ProgID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ProgID, _ptr_progID); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// queryCriteria {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.QueryCriteria != nil {
			_ptr_queryCriteria := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.QueryCriteria != nil {
					if err := o.QueryCriteria.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QueryCriteria, _ptr_queryCriteria); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// progID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_progID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ProgID == nil {
				o.ProgID = &oaut.String{}
			}
			if err := o.ProgID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_progID := func(ptr interface{}) { o.ProgID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ProgID, _s_progID, _ptr_progID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// queryCriteria {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_queryCriteria := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.QueryCriteria == nil {
				o.QueryCriteria = &oaut.String{}
			}
			if err := o.QueryCriteria.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_queryCriteria := func(ptr interface{}) { o.QueryCriteria = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.QueryCriteria, _s_queryCriteria, _ptr_queryCriteria); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ppInterface {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IUnknown}(interface))
	{
		if o.Interface != nil {
			_ptr_ppInterface := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Interface != nil {
					if err := o.Interface.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.Unknown{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Interface, _ptr_ppInterface); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppInterface {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IUnknown}(interface))
	{
		_ptr_ppInterface := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Interface == nil {
				o.Interface = &dcom.Unknown{}
			}
			if err := o.Interface.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppInterface := func(ptr interface{}) { o.Interface = *ptr.(**dcom.Unknown) }
		if err := w.ReadPointer(&o.Interface, _s_ppInterface, _ptr_ppInterface); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QuerySRequest structure represents the QueryS operation request
type QuerySRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// progID: A string that uniquely identifies the type of collection. The value MUST
	// be one of the following.
	//
	//	+-------------------------------------------+----------------------------------------------------------------+
	//	|                                           |                                                                |
	//	|                   VALUE                   |                            MEANING                             |
	//	|                                           |                                                                |
	//	+-------------------------------------------+----------------------------------------------------------------+
	//	+-------------------------------------------+----------------------------------------------------------------+
	//	| "EventSystem.EventClassCollection"        | The store for event classes (as specified in section 3.1.1.1). |
	//	+-------------------------------------------+----------------------------------------------------------------+
	//	| "EventSystem.EventSubscriptionCollection" | The store for subscriptions (as specified in section 3.1.1.2). |
	//	+-------------------------------------------+----------------------------------------------------------------+
	ProgID *oaut.String `idl:"name:progID" json:"prog_id"`
	// queryCriteria: The actual query string. The syntax for this string MUST conform to
	// section 2.2.1.
	QueryCriteria *oaut.String `idl:"name:queryCriteria" json:"query_criteria"`
}

func (o *QuerySRequest) xxx_ToOp(ctx context.Context) *xxx_QuerySOperation {
	if o == nil {
		return &xxx_QuerySOperation{}
	}
	return &xxx_QuerySOperation{
		This:          o.This,
		ProgID:        o.ProgID,
		QueryCriteria: o.QueryCriteria,
	}
}

func (o *QuerySRequest) xxx_FromOp(ctx context.Context, op *xxx_QuerySOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ProgID = op.ProgID
	o.QueryCriteria = op.QueryCriteria
}
func (o *QuerySRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QuerySRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QuerySOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QuerySResponse structure represents the QueryS operation response
type QuerySResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppInterface: If the method returns success, this MUST contain an interface pointer
	// that represents the collection of the event classes or subscriptions based on the
	// criteria specified in the queryCriteria parameter.
	Interface *dcom.Unknown `idl:"name:ppInterface" json:"interface"`
	// Return: The QueryS return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QuerySResponse) xxx_ToOp(ctx context.Context) *xxx_QuerySOperation {
	if o == nil {
		return &xxx_QuerySOperation{}
	}
	return &xxx_QuerySOperation{
		That:      o.That,
		Interface: o.Interface,
		Return:    o.Return,
	}
}

func (o *QuerySResponse) xxx_FromOp(ctx context.Context, op *xxx_QuerySOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Interface = op.Interface
	o.Return = op.Return
}
func (o *QuerySResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QuerySResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QuerySOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoveSOperation structure represents the RemoveS operation
type xxx_RemoveSOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	ProgID        *oaut.String   `idl:"name:progID" json:"prog_id"`
	QueryCriteria *oaut.String   `idl:"name:queryCriteria" json:"query_criteria"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoveSOperation) OpNum() int { return 12 }

func (o *xxx_RemoveSOperation) OpName() string { return "/IEventSystem/v0/RemoveS" }

func (o *xxx_RemoveSOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveSOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// progID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ProgID != nil {
			_ptr_progID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ProgID != nil {
					if err := o.ProgID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ProgID, _ptr_progID); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// queryCriteria {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.QueryCriteria != nil {
			_ptr_queryCriteria := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.QueryCriteria != nil {
					if err := o.QueryCriteria.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QueryCriteria, _ptr_queryCriteria); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveSOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// progID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_progID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ProgID == nil {
				o.ProgID = &oaut.String{}
			}
			if err := o.ProgID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_progID := func(ptr interface{}) { o.ProgID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ProgID, _s_progID, _ptr_progID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// queryCriteria {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_queryCriteria := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.QueryCriteria == nil {
				o.QueryCriteria = &oaut.String{}
			}
			if err := o.QueryCriteria.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_queryCriteria := func(ptr interface{}) { o.QueryCriteria = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.QueryCriteria, _s_queryCriteria, _ptr_queryCriteria); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveSOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveSOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveSOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RemoveSRequest structure represents the RemoveS operation request
type RemoveSRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// progID: A string that uniquely identifies the type of collection. The value MUST
	// be one of the following.
	//
	//	+-------------------------------------------+----------------------------------------------------------------+
	//	|                                           |                                                                |
	//	|                   VALUE                   |                            MEANING                             |
	//	|                                           |                                                                |
	//	+-------------------------------------------+----------------------------------------------------------------+
	//	+-------------------------------------------+----------------------------------------------------------------+
	//	| "EventSystem.EventClassCollection"        | The store for event classes (as specified in section 3.1.1.1). |
	//	+-------------------------------------------+----------------------------------------------------------------+
	//	| "EventSystem.EventSubscriptionCollection" | The store for subscriptions (as specified in section 3.1.1.2). |
	//	+-------------------------------------------+----------------------------------------------------------------+
	ProgID *oaut.String `idl:"name:progID" json:"prog_id"`
	// queryCriteria: The actual query string. The syntax for this string MUST conform to
	// section 2.2.1.
	QueryCriteria *oaut.String `idl:"name:queryCriteria" json:"query_criteria"`
}

func (o *RemoveSRequest) xxx_ToOp(ctx context.Context) *xxx_RemoveSOperation {
	if o == nil {
		return &xxx_RemoveSOperation{}
	}
	return &xxx_RemoveSOperation{
		This:          o.This,
		ProgID:        o.ProgID,
		QueryCriteria: o.QueryCriteria,
	}
}

func (o *RemoveSRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoveSOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ProgID = op.ProgID
	o.QueryCriteria = op.QueryCriteria
}
func (o *RemoveSRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RemoveSRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveSOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoveSResponse structure represents the RemoveS operation response
type RemoveSResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RemoveS return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemoveSResponse) xxx_ToOp(ctx context.Context) *xxx_RemoveSOperation {
	if o == nil {
		return &xxx_RemoveSOperation{}
	}
	return &xxx_RemoveSOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *RemoveSResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoveSOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RemoveSResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RemoveSResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveSOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
