package idatafactory

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
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
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/adtg"
)

var (
	// IDataFactory interface identifier 0eac4842-8763-11cf-a743-00aa00a3f00d
	DataFactoryIID = &dcom.IID{Data1: 0x0eac4842, Data2: 0x8763, Data3: 0x11cf, Data4: []byte{0xa7, 0x43, 0x00, 0xaa, 0x00, 0xa3, 0xf0, 0x0d}}
	// Syntax UUID
	DataFactorySyntaxUUID = &uuid.UUID{TimeLow: 0xeac4842, TimeMid: 0x8763, TimeHiAndVersion: 0x11cf, ClockSeqHiAndReserved: 0xa7, ClockSeqLow: 0x43, Node: [6]uint8{0x0, 0xaa, 0x0, 0xa3, 0xf0, 0xd}}
	// Syntax ID
	DataFactorySyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: DataFactorySyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IDataFactory interface.
type DataFactoryClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// Query operation.
	Query(context.Context, *QueryRequest, ...dcerpc.CallOption) (*QueryResponse, error)

	// SubmitChanges operation.
	SubmitChanges(context.Context, *SubmitChangesRequest, ...dcerpc.CallOption) (*SubmitChangesResponse, error)

	// ConvertToString operation.
	ConvertToString(context.Context, *ConvertToStringRequest, ...dcerpc.CallOption) (*ConvertToStringResponse, error)

	// CreateRecordSet operation.
	CreateRecordSet(context.Context, *CreateRecordSetRequest, ...dcerpc.CallOption) (*CreateRecordSetResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) DataFactoryClient
}

type xxx_DefaultDataFactoryClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultDataFactoryClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultDataFactoryClient) Query(ctx context.Context, in *QueryRequest, opts ...dcerpc.CallOption) (*QueryResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataFactoryClient) SubmitChanges(ctx context.Context, in *SubmitChangesRequest, opts ...dcerpc.CallOption) (*SubmitChangesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &SubmitChangesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataFactoryClient) ConvertToString(ctx context.Context, in *ConvertToStringRequest, opts ...dcerpc.CallOption) (*ConvertToStringResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &ConvertToStringResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataFactoryClient) CreateRecordSet(ctx context.Context, in *CreateRecordSetRequest, opts ...dcerpc.CallOption) (*CreateRecordSetResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &CreateRecordSetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataFactoryClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultDataFactoryClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultDataFactoryClient) IPID(ctx context.Context, ipid *dcom.IPID) DataFactoryClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultDataFactoryClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewDataFactoryClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (DataFactoryClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(DataFactorySyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iunknown.NewUnknownClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultDataFactoryClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_QueryOperation structure represents the Query operation
type xxx_QueryOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	Connection     *oaut.String   `idl:"name:bstrConnection" json:"connection"`
	Query          *oaut.String   `idl:"name:bstrQuery" json:"query"`
	MarshalOptions int32          `idl:"name:lMarshalOptions" json:"marshal_options"`
	RecordSet      *oaut.Dispatch `idl:"name:ppRecordset" json:"record_set"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryOperation) OpNum() int { return 3 }

func (o *xxx_QueryOperation) OpName() string { return "/IDataFactory/v0/Query" }

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
	// bstrConnection {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Connection != nil {
			_ptr_bstrConnection := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Connection != nil {
					if err := o.Connection.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Connection, _ptr_bstrConnection); err != nil {
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
	// bstrQuery {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Query != nil {
			_ptr_bstrQuery := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Query != nil {
					if err := o.Query.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Query, _ptr_bstrQuery); err != nil {
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
	// lMarshalOptions {in, default_value={0}} (1:(int32))
	{
		if err := w.WriteData(o.MarshalOptions); err != nil {
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
	// bstrConnection {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrConnection := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Connection == nil {
				o.Connection = &oaut.String{}
			}
			if err := o.Connection.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrConnection := func(ptr interface{}) { o.Connection = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Connection, _s_bstrConnection, _ptr_bstrConnection); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrQuery {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrQuery := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Query == nil {
				o.Query = &oaut.String{}
			}
			if err := o.Query.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrQuery := func(ptr interface{}) { o.Query = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Query, _s_bstrQuery, _ptr_bstrQuery); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lMarshalOptions {in, default_value={0}} (1:(int32))
	{
		if err := w.ReadData(&o.MarshalOptions); err != nil {
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
	// ppRecordset {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IDispatch}(interface))
	{
		if o.RecordSet != nil {
			_ptr_ppRecordset := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RecordSet != nil {
					if err := o.RecordSet.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Dispatch{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RecordSet, _ptr_ppRecordset); err != nil {
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
	// ppRecordset {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IDispatch}(interface))
	{
		_ptr_ppRecordset := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.RecordSet == nil {
				o.RecordSet = &oaut.Dispatch{}
			}
			if err := o.RecordSet.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppRecordset := func(ptr interface{}) { o.RecordSet = *ptr.(**oaut.Dispatch) }
		if err := w.ReadPointer(&o.RecordSet, _s_ppRecordset, _ptr_ppRecordset); err != nil {
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
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	Connection     *oaut.String   `idl:"name:bstrConnection" json:"connection"`
	Query          *oaut.String   `idl:"name:bstrQuery" json:"query"`
	MarshalOptions int32          `idl:"name:lMarshalOptions" json:"marshal_options"`
}

func (o *QueryRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryOperation) *xxx_QueryOperation {
	if op == nil {
		op = &xxx_QueryOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Connection = o.Connection
	op.Query = o.Query
	op.MarshalOptions = o.MarshalOptions
	return op
}

func (o *QueryRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Connection = op.Connection
	o.Query = op.Query
	o.MarshalOptions = op.MarshalOptions
}
func (o *QueryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
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
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	RecordSet *oaut.Dispatch `idl:"name:ppRecordset" json:"record_set"`
	// Return: The Query return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryOperation) *xxx_QueryOperation {
	if op == nil {
		op = &xxx_QueryOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.RecordSet = o.RecordSet
	op.Return = o.Return
	return op
}

func (o *QueryResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.RecordSet = op.RecordSet
	o.Return = op.Return
}
func (o *QueryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SubmitChangesOperation structure represents the SubmitChanges operation
type xxx_SubmitChangesOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	Connection *oaut.String   `idl:"name:bstrConnection" json:"connection"`
	RecordSet  *oaut.Dispatch `idl:"name:pRecordset" json:"record_set"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SubmitChangesOperation) OpNum() int { return 4 }

func (o *xxx_SubmitChangesOperation) OpName() string { return "/IDataFactory/v0/SubmitChanges" }

func (o *xxx_SubmitChangesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SubmitChangesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrConnection {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Connection != nil {
			_ptr_bstrConnection := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Connection != nil {
					if err := o.Connection.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Connection, _ptr_bstrConnection); err != nil {
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
	// pRecordset {in} (1:{pointer=ref}*(1))(2:{alias=IDispatch}(interface))
	{
		if o.RecordSet != nil {
			_ptr_pRecordset := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RecordSet != nil {
					if err := o.RecordSet.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Dispatch{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RecordSet, _ptr_pRecordset); err != nil {
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

func (o *xxx_SubmitChangesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrConnection {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrConnection := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Connection == nil {
				o.Connection = &oaut.String{}
			}
			if err := o.Connection.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrConnection := func(ptr interface{}) { o.Connection = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Connection, _s_bstrConnection, _ptr_bstrConnection); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pRecordset {in} (1:{pointer=ref}*(1))(2:{alias=IDispatch}(interface))
	{
		_ptr_pRecordset := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.RecordSet == nil {
				o.RecordSet = &oaut.Dispatch{}
			}
			if err := o.RecordSet.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pRecordset := func(ptr interface{}) { o.RecordSet = *ptr.(**oaut.Dispatch) }
		if err := w.ReadPointer(&o.RecordSet, _s_pRecordset, _ptr_pRecordset); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SubmitChangesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SubmitChangesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SubmitChangesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SubmitChangesRequest structure represents the SubmitChanges operation request
type SubmitChangesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	Connection *oaut.String   `idl:"name:bstrConnection" json:"connection"`
	RecordSet  *oaut.Dispatch `idl:"name:pRecordset" json:"record_set"`
}

func (o *SubmitChangesRequest) xxx_ToOp(ctx context.Context, op *xxx_SubmitChangesOperation) *xxx_SubmitChangesOperation {
	if op == nil {
		op = &xxx_SubmitChangesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Connection = o.Connection
	op.RecordSet = o.RecordSet
	return op
}

func (o *SubmitChangesRequest) xxx_FromOp(ctx context.Context, op *xxx_SubmitChangesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Connection = op.Connection
	o.RecordSet = op.RecordSet
}
func (o *SubmitChangesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SubmitChangesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SubmitChangesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SubmitChangesResponse structure represents the SubmitChanges operation response
type SubmitChangesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SubmitChanges return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SubmitChangesResponse) xxx_ToOp(ctx context.Context, op *xxx_SubmitChangesOperation) *xxx_SubmitChangesOperation {
	if op == nil {
		op = &xxx_SubmitChangesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SubmitChangesResponse) xxx_FromOp(ctx context.Context, op *xxx_SubmitChangesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SubmitChangesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SubmitChangesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SubmitChangesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ConvertToStringOperation structure represents the ConvertToString operation
type xxx_ConvertToStringOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Object *dcom.Unknown  `idl:"name:punkObject" json:"object"`
	Inline *oaut.String   `idl:"name:pbstrInline" json:"inline"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ConvertToStringOperation) OpNum() int { return 5 }

func (o *xxx_ConvertToStringOperation) OpName() string { return "/IDataFactory/v0/ConvertToString" }

func (o *xxx_ConvertToStringOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConvertToStringOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// punkObject {in} (1:{pointer=ref}*(1))(2:{alias=IUnknown}(interface))
	{
		if o.Object != nil {
			_ptr_punkObject := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Object != nil {
					if err := o.Object.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.Unknown{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Object, _ptr_punkObject); err != nil {
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

func (o *xxx_ConvertToStringOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// punkObject {in} (1:{pointer=ref}*(1))(2:{alias=IUnknown}(interface))
	{
		_ptr_punkObject := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Object == nil {
				o.Object = &dcom.Unknown{}
			}
			if err := o.Object.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_punkObject := func(ptr interface{}) { o.Object = *ptr.(**dcom.Unknown) }
		if err := w.ReadPointer(&o.Object, _s_punkObject, _ptr_punkObject); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConvertToStringOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConvertToStringOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrInline {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Inline != nil {
			_ptr_pbstrInline := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Inline != nil {
					if err := o.Inline.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Inline, _ptr_pbstrInline); err != nil {
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

func (o *xxx_ConvertToStringOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrInline {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrInline := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Inline == nil {
				o.Inline = &oaut.String{}
			}
			if err := o.Inline.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrInline := func(ptr interface{}) { o.Inline = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Inline, _s_pbstrInline, _ptr_pbstrInline); err != nil {
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

// ConvertToStringRequest structure represents the ConvertToString operation request
type ConvertToStringRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	Object *dcom.Unknown  `idl:"name:punkObject" json:"object"`
}

func (o *ConvertToStringRequest) xxx_ToOp(ctx context.Context, op *xxx_ConvertToStringOperation) *xxx_ConvertToStringOperation {
	if op == nil {
		op = &xxx_ConvertToStringOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Object = o.Object
	return op
}

func (o *ConvertToStringRequest) xxx_FromOp(ctx context.Context, op *xxx_ConvertToStringOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Object = op.Object
}
func (o *ConvertToStringRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ConvertToStringRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConvertToStringOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ConvertToStringResponse structure represents the ConvertToString operation response
type ConvertToStringResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Inline *oaut.String   `idl:"name:pbstrInline" json:"inline"`
	// Return: The ConvertToString return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ConvertToStringResponse) xxx_ToOp(ctx context.Context, op *xxx_ConvertToStringOperation) *xxx_ConvertToStringOperation {
	if op == nil {
		op = &xxx_ConvertToStringOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Inline = o.Inline
	op.Return = o.Return
	return op
}

func (o *ConvertToStringResponse) xxx_FromOp(ctx context.Context, op *xxx_ConvertToStringOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Inline = op.Inline
	o.Return = op.Return
}
func (o *ConvertToStringResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ConvertToStringResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConvertToStringOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateRecordSetOperation structure represents the CreateRecordSet operation
type xxx_CreateRecordSetOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	VarColumnInfos *oaut.Variant  `idl:"name:varColumnInfos" json:"var_column_infos"`
	Dispatch       *oaut.Dispatch `idl:"name:ppDispatch" json:"dispatch"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateRecordSetOperation) OpNum() int { return 6 }

func (o *xxx_CreateRecordSetOperation) OpName() string { return "/IDataFactory/v0/CreateRecordSet" }

func (o *xxx_CreateRecordSetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateRecordSetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// varColumnInfos {in} (1:{alias=VARIANT}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.VarColumnInfos != nil {
			if err := o.VarColumnInfos.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateRecordSetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// varColumnInfos {in} (1:{alias=VARIANT,pointer=ref}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.VarColumnInfos == nil {
			o.VarColumnInfos = &oaut.Variant{}
		}
		if err := o.VarColumnInfos.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateRecordSetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateRecordSetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppDispatch {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IDispatch}(interface))
	{
		if o.Dispatch != nil {
			_ptr_ppDispatch := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Dispatch != nil {
					if err := o.Dispatch.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Dispatch{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Dispatch, _ptr_ppDispatch); err != nil {
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

func (o *xxx_CreateRecordSetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppDispatch {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IDispatch}(interface))
	{
		_ptr_ppDispatch := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Dispatch == nil {
				o.Dispatch = &oaut.Dispatch{}
			}
			if err := o.Dispatch.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppDispatch := func(ptr interface{}) { o.Dispatch = *ptr.(**oaut.Dispatch) }
		if err := w.ReadPointer(&o.Dispatch, _s_ppDispatch, _ptr_ppDispatch); err != nil {
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

// CreateRecordSetRequest structure represents the CreateRecordSet operation request
type CreateRecordSetRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	VarColumnInfos *oaut.Variant  `idl:"name:varColumnInfos" json:"var_column_infos"`
}

func (o *CreateRecordSetRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateRecordSetOperation) *xxx_CreateRecordSetOperation {
	if op == nil {
		op = &xxx_CreateRecordSetOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.VarColumnInfos = o.VarColumnInfos
	return op
}

func (o *CreateRecordSetRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateRecordSetOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VarColumnInfos = op.VarColumnInfos
}
func (o *CreateRecordSetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateRecordSetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateRecordSetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateRecordSetResponse structure represents the CreateRecordSet operation response
type CreateRecordSetResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Dispatch *oaut.Dispatch `idl:"name:ppDispatch" json:"dispatch"`
	// Return: The CreateRecordSet return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateRecordSetResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateRecordSetOperation) *xxx_CreateRecordSetOperation {
	if op == nil {
		op = &xxx_CreateRecordSetOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Dispatch = o.Dispatch
	op.Return = o.Return
	return op
}

func (o *CreateRecordSetResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateRecordSetOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Dispatch = op.Dispatch
	o.Return = op.Return
}
func (o *CreateRecordSetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateRecordSetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateRecordSetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
