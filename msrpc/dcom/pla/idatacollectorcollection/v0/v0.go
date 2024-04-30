package idatacollectorcollection

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
	pla "github.com/oiweiwei/go-msrpc/msrpc/dcom/pla"
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
	_ = pla.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/pla"
)

var (
	// IDataCollectorCollection interface identifier 03837502-098b-11d8-9414-505054503030
	DataCollectorCollectionIID = &dcom.IID{Data1: 0x03837502, Data2: 0x098b, Data3: 0x11d8, Data4: []byte{0x94, 0x14, 0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	// Syntax UUID
	DataCollectorCollectionSyntaxUUID = &uuid.UUID{TimeLow: 0x3837502, TimeMid: 0x98b, TimeHiAndVersion: 0x11d8, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x14, Node: [6]uint8{0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	// Syntax ID
	DataCollectorCollectionSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: DataCollectorCollectionSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IDataCollectorCollection interface.
type DataCollectorCollectionClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// Count operation.
	GetCount(context.Context, *GetCountRequest, ...dcerpc.CallOption) (*GetCountResponse, error)

	// Item operation.
	GetItem(context.Context, *GetItemRequest, ...dcerpc.CallOption) (*GetItemResponse, error)

	// _NewEnum operation.
	Get_NewEnum(context.Context, *Get_NewEnumRequest, ...dcerpc.CallOption) (*Get_NewEnumResponse, error)

	// Add operation.
	Add(context.Context, *AddRequest, ...dcerpc.CallOption) (*AddResponse, error)

	// Remove operation.
	Remove(context.Context, *RemoveRequest, ...dcerpc.CallOption) (*RemoveResponse, error)

	// Clear operation.
	Clear(context.Context, *ClearRequest, ...dcerpc.CallOption) (*ClearResponse, error)

	// AddRange operation.
	AddRange(context.Context, *AddRangeRequest, ...dcerpc.CallOption) (*AddRangeResponse, error)

	// The CreateDataCollectorFromXml method creates a data collector using XML.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	CreateDataCollectorFromXML(context.Context, *CreateDataCollectorFromXMLRequest, ...dcerpc.CallOption) (*CreateDataCollectorFromXMLResponse, error)

	// The CreateDataCollector method creates a data collector of the specified type.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	CreateDataCollector(context.Context, *CreateDataCollectorRequest, ...dcerpc.CallOption) (*CreateDataCollectorResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) DataCollectorCollectionClient
}

type xxx_DefaultDataCollectorCollectionClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultDataCollectorCollectionClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultDataCollectorCollectionClient) GetCount(ctx context.Context, in *GetCountRequest, opts ...dcerpc.CallOption) (*GetCountResponse, error) {
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
	out := &GetCountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorCollectionClient) GetItem(ctx context.Context, in *GetItemRequest, opts ...dcerpc.CallOption) (*GetItemResponse, error) {
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
	out := &GetItemResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorCollectionClient) Get_NewEnum(ctx context.Context, in *Get_NewEnumRequest, opts ...dcerpc.CallOption) (*Get_NewEnumResponse, error) {
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
	out := &Get_NewEnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorCollectionClient) Add(ctx context.Context, in *AddRequest, opts ...dcerpc.CallOption) (*AddResponse, error) {
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
	out := &AddResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorCollectionClient) Remove(ctx context.Context, in *RemoveRequest, opts ...dcerpc.CallOption) (*RemoveResponse, error) {
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

func (o *xxx_DefaultDataCollectorCollectionClient) Clear(ctx context.Context, in *ClearRequest, opts ...dcerpc.CallOption) (*ClearResponse, error) {
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
	out := &ClearResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorCollectionClient) AddRange(ctx context.Context, in *AddRangeRequest, opts ...dcerpc.CallOption) (*AddRangeResponse, error) {
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
	out := &AddRangeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorCollectionClient) CreateDataCollectorFromXML(ctx context.Context, in *CreateDataCollectorFromXMLRequest, opts ...dcerpc.CallOption) (*CreateDataCollectorFromXMLResponse, error) {
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
	out := &CreateDataCollectorFromXMLResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorCollectionClient) CreateDataCollector(ctx context.Context, in *CreateDataCollectorRequest, opts ...dcerpc.CallOption) (*CreateDataCollectorResponse, error) {
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
	out := &CreateDataCollectorResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDataCollectorCollectionClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultDataCollectorCollectionClient) IPID(ctx context.Context, ipid *dcom.IPID) DataCollectorCollectionClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultDataCollectorCollectionClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}
func NewDataCollectorCollectionClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (DataCollectorCollectionClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(DataCollectorCollectionSyntaxV0_0))...)
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
	return &xxx_DefaultDataCollectorCollectionClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetCountOperation structure represents the Count operation
type xxx_GetCountOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReturnValue int32          `idl:"name:retVal" json:"return_value"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCountOperation) OpNum() int { return 7 }

func (o *xxx_GetCountOperation) OpName() string { return "/IDataCollectorCollection/v0/Count" }

func (o *xxx_GetCountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetCountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetCountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// retVal {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.ReturnValue); err != nil {
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

func (o *xxx_GetCountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// retVal {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.ReturnValue); err != nil {
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

// GetCountRequest structure represents the Count operation request
type GetCountRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetCountRequest) xxx_ToOp(ctx context.Context) *xxx_GetCountOperation {
	if o == nil {
		return &xxx_GetCountOperation{}
	}
	return &xxx_GetCountOperation{
		This: o.This,
	}
}

func (o *GetCountRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCountOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetCountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetCountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCountResponse structure represents the Count operation response
type GetCountResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReturnValue int32          `idl:"name:retVal" json:"return_value"`
	// Return: The Count return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCountResponse) xxx_ToOp(ctx context.Context) *xxx_GetCountOperation {
	if o == nil {
		return &xxx_GetCountOperation{}
	}
	return &xxx_GetCountOperation{
		That:        o.That,
		ReturnValue: o.ReturnValue,
		Return:      o.Return,
	}
}

func (o *GetCountResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCountOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *GetCountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetCountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetItemOperation structure represents the Item operation
type xxx_GetItemOperation struct {
	This      *dcom.ORPCThis     `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat     `idl:"name:That" json:"that"`
	Index     *oaut.Variant      `idl:"name:index" json:"index"`
	Collector *pla.DataCollector `idl:"name:collector" json:"collector"`
	Return    int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_GetItemOperation) OpNum() int { return 8 }

func (o *xxx_GetItemOperation) OpName() string { return "/IDataCollectorCollection/v0/Item" }

func (o *xxx_GetItemOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetItemOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// index {in} (1:{alias=VARIANT}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.Index != nil {
			if err := o.Index.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetItemOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// index {in} (1:{alias=VARIANT,pointer=ref}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.Index == nil {
			o.Index = &oaut.Variant{}
		}
		if err := o.Index.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetItemOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetItemOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// collector {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IDataCollector}(interface))
	{
		if o.Collector != nil {
			_ptr_collector := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Collector != nil {
					if err := o.Collector.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&pla.DataCollector{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Collector, _ptr_collector); err != nil {
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

func (o *xxx_GetItemOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// collector {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IDataCollector}(interface))
	{
		_ptr_collector := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Collector == nil {
				o.Collector = &pla.DataCollector{}
			}
			if err := o.Collector.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_collector := func(ptr interface{}) { o.Collector = *ptr.(**pla.DataCollector) }
		if err := w.ReadPointer(&o.Collector, _s_collector, _ptr_collector); err != nil {
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

// GetItemRequest structure represents the Item operation request
type GetItemRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis `idl:"name:This" json:"this"`
	Index *oaut.Variant  `idl:"name:index" json:"index"`
}

func (o *GetItemRequest) xxx_ToOp(ctx context.Context) *xxx_GetItemOperation {
	if o == nil {
		return &xxx_GetItemOperation{}
	}
	return &xxx_GetItemOperation{
		This:  o.This,
		Index: o.Index,
	}
}

func (o *GetItemRequest) xxx_FromOp(ctx context.Context, op *xxx_GetItemOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Index = op.Index
}
func (o *GetItemRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetItemRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetItemOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetItemResponse structure represents the Item operation response
type GetItemResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat     `idl:"name:That" json:"that"`
	Collector *pla.DataCollector `idl:"name:collector" json:"collector"`
	// Return: The Item return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetItemResponse) xxx_ToOp(ctx context.Context) *xxx_GetItemOperation {
	if o == nil {
		return &xxx_GetItemOperation{}
	}
	return &xxx_GetItemOperation{
		That:      o.That,
		Collector: o.Collector,
		Return:    o.Return,
	}
}

func (o *GetItemResponse) xxx_FromOp(ctx context.Context, op *xxx_GetItemOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Collector = op.Collector
	o.Return = op.Return
}
func (o *GetItemResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetItemResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetItemOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_Get_NewEnumOperation structure represents the _NewEnum operation
type xxx_Get_NewEnumOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReturnValue *dcom.Unknown  `idl:"name:retVal" json:"return_value"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_Get_NewEnumOperation) OpNum() int { return 9 }

func (o *xxx_Get_NewEnumOperation) OpName() string { return "/IDataCollectorCollection/v0/_NewEnum" }

func (o *xxx_Get_NewEnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Get_NewEnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_Get_NewEnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_Get_NewEnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Get_NewEnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// retVal {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IUnknown}(interface))
	{
		if o.ReturnValue != nil {
			_ptr_retVal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReturnValue != nil {
					if err := o.ReturnValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.Unknown{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReturnValue, _ptr_retVal); err != nil {
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

func (o *xxx_Get_NewEnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// retVal {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IUnknown}(interface))
	{
		_ptr_retVal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReturnValue == nil {
				o.ReturnValue = &dcom.Unknown{}
			}
			if err := o.ReturnValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_retVal := func(ptr interface{}) { o.ReturnValue = *ptr.(**dcom.Unknown) }
		if err := w.ReadPointer(&o.ReturnValue, _s_retVal, _ptr_retVal); err != nil {
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

// Get_NewEnumRequest structure represents the _NewEnum operation request
type Get_NewEnumRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *Get_NewEnumRequest) xxx_ToOp(ctx context.Context) *xxx_Get_NewEnumOperation {
	if o == nil {
		return &xxx_Get_NewEnumOperation{}
	}
	return &xxx_Get_NewEnumOperation{
		This: o.This,
	}
}

func (o *Get_NewEnumRequest) xxx_FromOp(ctx context.Context, op *xxx_Get_NewEnumOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *Get_NewEnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *Get_NewEnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Get_NewEnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// Get_NewEnumResponse structure represents the _NewEnum operation response
type Get_NewEnumResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReturnValue *dcom.Unknown  `idl:"name:retVal" json:"return_value"`
	// Return: The _NewEnum return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *Get_NewEnumResponse) xxx_ToOp(ctx context.Context) *xxx_Get_NewEnumOperation {
	if o == nil {
		return &xxx_Get_NewEnumOperation{}
	}
	return &xxx_Get_NewEnumOperation{
		That:        o.That,
		ReturnValue: o.ReturnValue,
		Return:      o.Return,
	}
}

func (o *Get_NewEnumResponse) xxx_FromOp(ctx context.Context, op *xxx_Get_NewEnumOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *Get_NewEnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *Get_NewEnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Get_NewEnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddOperation structure represents the Add operation
type xxx_AddOperation struct {
	This      *dcom.ORPCThis     `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat     `idl:"name:That" json:"that"`
	Collector *pla.DataCollector `idl:"name:collector" json:"collector"`
	Return    int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_AddOperation) OpNum() int { return 10 }

func (o *xxx_AddOperation) OpName() string { return "/IDataCollectorCollection/v0/Add" }

func (o *xxx_AddOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// collector {in} (1:{pointer=ref}*(1))(2:{alias=IDataCollector}(interface))
	{
		if o.Collector != nil {
			_ptr_collector := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Collector != nil {
					if err := o.Collector.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&pla.DataCollector{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Collector, _ptr_collector); err != nil {
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

func (o *xxx_AddOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// collector {in} (1:{pointer=ref}*(1))(2:{alias=IDataCollector}(interface))
	{
		_ptr_collector := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Collector == nil {
				o.Collector = &pla.DataCollector{}
			}
			if err := o.Collector.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_collector := func(ptr interface{}) { o.Collector = *ptr.(**pla.DataCollector) }
		if err := w.ReadPointer(&o.Collector, _s_collector, _ptr_collector); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AddOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// AddRequest structure represents the Add operation request
type AddRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This      *dcom.ORPCThis     `idl:"name:This" json:"this"`
	Collector *pla.DataCollector `idl:"name:collector" json:"collector"`
}

func (o *AddRequest) xxx_ToOp(ctx context.Context) *xxx_AddOperation {
	if o == nil {
		return &xxx_AddOperation{}
	}
	return &xxx_AddOperation{
		This:      o.This,
		Collector: o.Collector,
	}
}

func (o *AddRequest) xxx_FromOp(ctx context.Context, op *xxx_AddOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Collector = op.Collector
}
func (o *AddRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AddRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddResponse structure represents the Add operation response
type AddResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Add return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AddResponse) xxx_ToOp(ctx context.Context) *xxx_AddOperation {
	if o == nil {
		return &xxx_AddOperation{}
	}
	return &xxx_AddOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *AddResponse) xxx_FromOp(ctx context.Context, op *xxx_AddOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *AddResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AddResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoveOperation structure represents the Remove operation
type xxx_RemoveOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Collector *oaut.Variant  `idl:"name:collector" json:"collector"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoveOperation) OpNum() int { return 11 }

func (o *xxx_RemoveOperation) OpName() string { return "/IDataCollectorCollection/v0/Remove" }

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
	// collector {in} (1:{alias=VARIANT}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.Collector != nil {
			if err := o.Collector.MarshalNDR(ctx, w); err != nil {
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
	// collector {in} (1:{alias=VARIANT,pointer=ref}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.Collector == nil {
			o.Collector = &oaut.Variant{}
		}
		if err := o.Collector.UnmarshalNDR(ctx, w); err != nil {
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
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	Collector *oaut.Variant  `idl:"name:collector" json:"collector"`
}

func (o *RemoveRequest) xxx_ToOp(ctx context.Context) *xxx_RemoveOperation {
	if o == nil {
		return &xxx_RemoveOperation{}
	}
	return &xxx_RemoveOperation{
		This:      o.This,
		Collector: o.Collector,
	}
}

func (o *RemoveRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoveOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Collector = op.Collector
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
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Remove return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemoveResponse) xxx_ToOp(ctx context.Context) *xxx_RemoveOperation {
	if o == nil {
		return &xxx_RemoveOperation{}
	}
	return &xxx_RemoveOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *RemoveResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoveOperation) {
	if o == nil {
		return
	}
	o.That = op.That
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

// xxx_ClearOperation structure represents the Clear operation
type xxx_ClearOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ClearOperation) OpNum() int { return 12 }

func (o *xxx_ClearOperation) OpName() string { return "/IDataCollectorCollection/v0/Clear" }

func (o *xxx_ClearOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ClearOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_ClearOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ClearOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ClearRequest structure represents the Clear operation request
type ClearRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *ClearRequest) xxx_ToOp(ctx context.Context) *xxx_ClearOperation {
	if o == nil {
		return &xxx_ClearOperation{}
	}
	return &xxx_ClearOperation{
		This: o.This,
	}
}

func (o *ClearRequest) xxx_FromOp(ctx context.Context, op *xxx_ClearOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *ClearRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ClearRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClearOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ClearResponse structure represents the Clear operation response
type ClearResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Clear return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ClearResponse) xxx_ToOp(ctx context.Context) *xxx_ClearOperation {
	if o == nil {
		return &xxx_ClearOperation{}
	}
	return &xxx_ClearOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *ClearResponse) xxx_FromOp(ctx context.Context, op *xxx_ClearOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ClearResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ClearResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClearOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddRangeOperation structure represents the AddRange operation
type xxx_AddRangeOperation struct {
	This       *dcom.ORPCThis               `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat               `idl:"name:That" json:"that"`
	Collectors *pla.DataCollectorCollection `idl:"name:collectors" json:"collectors"`
	Return     int32                        `idl:"name:Return" json:"return"`
}

func (o *xxx_AddRangeOperation) OpNum() int { return 13 }

func (o *xxx_AddRangeOperation) OpName() string { return "/IDataCollectorCollection/v0/AddRange" }

func (o *xxx_AddRangeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddRangeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// collectors {in} (1:{pointer=ref}*(1))(2:{alias=IDataCollectorCollection}(interface))
	{
		if o.Collectors != nil {
			_ptr_collectors := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Collectors != nil {
					if err := o.Collectors.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&pla.DataCollectorCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Collectors, _ptr_collectors); err != nil {
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

func (o *xxx_AddRangeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// collectors {in} (1:{pointer=ref}*(1))(2:{alias=IDataCollectorCollection}(interface))
	{
		_ptr_collectors := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Collectors == nil {
				o.Collectors = &pla.DataCollectorCollection{}
			}
			if err := o.Collectors.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_collectors := func(ptr interface{}) { o.Collectors = *ptr.(**pla.DataCollectorCollection) }
		if err := w.ReadPointer(&o.Collectors, _s_collectors, _ptr_collectors); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddRangeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddRangeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AddRangeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// AddRangeRequest structure represents the AddRange operation request
type AddRangeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis               `idl:"name:This" json:"this"`
	Collectors *pla.DataCollectorCollection `idl:"name:collectors" json:"collectors"`
}

func (o *AddRangeRequest) xxx_ToOp(ctx context.Context) *xxx_AddRangeOperation {
	if o == nil {
		return &xxx_AddRangeOperation{}
	}
	return &xxx_AddRangeOperation{
		This:       o.This,
		Collectors: o.Collectors,
	}
}

func (o *AddRangeRequest) xxx_FromOp(ctx context.Context, op *xxx_AddRangeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Collectors = op.Collectors
}
func (o *AddRangeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AddRangeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddRangeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddRangeResponse structure represents the AddRange operation response
type AddRangeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AddRange return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AddRangeResponse) xxx_ToOp(ctx context.Context) *xxx_AddRangeOperation {
	if o == nil {
		return &xxx_AddRangeOperation{}
	}
	return &xxx_AddRangeOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *AddRangeResponse) xxx_FromOp(ctx context.Context, op *xxx_AddRangeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *AddRangeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AddRangeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddRangeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateDataCollectorFromXMLOperation structure represents the CreateDataCollectorFromXml operation
type xxx_CreateDataCollectorFromXMLOperation struct {
	This       *dcom.ORPCThis     `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat     `idl:"name:That" json:"that"`
	XML        *oaut.String       `idl:"name:bstrXml" json:"xml"`
	Validation *pla.ValueMap      `idl:"name:pValidation" json:"validation"`
	Collector  *pla.DataCollector `idl:"name:pCollector" json:"collector"`
	Return     int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateDataCollectorFromXMLOperation) OpNum() int { return 14 }

func (o *xxx_CreateDataCollectorFromXMLOperation) OpName() string {
	return "/IDataCollectorCollection/v0/CreateDataCollectorFromXml"
}

func (o *xxx_CreateDataCollectorFromXMLOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateDataCollectorFromXMLOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.XML != nil {
			_ptr_bstrXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.XML != nil {
					if err := o.XML.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.XML, _ptr_bstrXml); err != nil {
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

func (o *xxx_CreateDataCollectorFromXMLOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrXml {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.XML == nil {
				o.XML = &oaut.String{}
			}
			if err := o.XML.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrXml := func(ptr interface{}) { o.XML = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.XML, _s_bstrXml, _ptr_bstrXml); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateDataCollectorFromXMLOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateDataCollectorFromXMLOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pValidation {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IValueMap}(interface))
	{
		if o.Validation != nil {
			_ptr_pValidation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Validation != nil {
					if err := o.Validation.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&pla.ValueMap{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Validation, _ptr_pValidation); err != nil {
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
	// pCollector {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IDataCollector}(interface))
	{
		if o.Collector != nil {
			_ptr_pCollector := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Collector != nil {
					if err := o.Collector.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&pla.DataCollector{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Collector, _ptr_pCollector); err != nil {
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

func (o *xxx_CreateDataCollectorFromXMLOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pValidation {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IValueMap}(interface))
	{
		_ptr_pValidation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Validation == nil {
				o.Validation = &pla.ValueMap{}
			}
			if err := o.Validation.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pValidation := func(ptr interface{}) { o.Validation = *ptr.(**pla.ValueMap) }
		if err := w.ReadPointer(&o.Validation, _s_pValidation, _ptr_pValidation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pCollector {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IDataCollector}(interface))
	{
		_ptr_pCollector := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Collector == nil {
				o.Collector = &pla.DataCollector{}
			}
			if err := o.Collector.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pCollector := func(ptr interface{}) { o.Collector = *ptr.(**pla.DataCollector) }
		if err := w.ReadPointer(&o.Collector, _s_pCollector, _ptr_pCollector); err != nil {
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

// CreateDataCollectorFromXMLRequest structure represents the CreateDataCollectorFromXml operation request
type CreateDataCollectorFromXMLRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrXml: Supplies a string that contains the XML specifying the data collector to
	// create. The possible data collector definitions can be as follows: IPerformanceCounterDataCollector
	// (section 3.2.4.6), IConfigurationDataCollector (section 3.2.4.7), IAlertDataCollector
	// (section 3.2.4.8), ITraceDataCollector (section 3.2.4.9), and IApiTracingDataCollector
	// (section 3.2.4.10). The XML for each of those respective data collector types is
	// in their corresponding sections; the overall XML of the data collector set, which
	// includes the XML for each type of data collector, specified in section 3.2.4.19.
	// The bStrXml parameter value is the set of XML elements corresponding to a single
	// type of data collector.
	XML *oaut.String `idl:"name:bstrXml" json:"xml"`
}

func (o *CreateDataCollectorFromXMLRequest) xxx_ToOp(ctx context.Context) *xxx_CreateDataCollectorFromXMLOperation {
	if o == nil {
		return &xxx_CreateDataCollectorFromXMLOperation{}
	}
	return &xxx_CreateDataCollectorFromXMLOperation{
		This: o.This,
		XML:  o.XML,
	}
}

func (o *CreateDataCollectorFromXMLRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateDataCollectorFromXMLOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.XML = op.XML
}
func (o *CreateDataCollectorFromXMLRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateDataCollectorFromXMLRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateDataCollectorFromXMLOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateDataCollectorFromXMLResponse structure represents the CreateDataCollectorFromXml operation response
type CreateDataCollectorFromXMLResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pValidation: Receives a validation value map with a list of properties from the input
	// bstrXml for this data collector set (and its encapsulated objects) that are invalid
	// or ignored. For each property of the data collector set and its associated objects,
	// passed in by the client, that could not be set, the server MUST return in an IValueMap.
	// Each IValueMapItem in the IValueMap represents a property of the data collector set
	// and its encapsulated objects that could not be set by the server. The Names property
	// of the IValueMapItem represents the property name, while the Values property of the
	// IValueMap represents the HRESULT describing the specific property corresponding to
	// that property. The ValueMapType property of the IValueMap is plaValidation; more
	// information can be found in section 2.2.11. Note that the client MAY choose to ignore
	// any warnings or errors that are returned by the server; however, if it does so, the
	// data collector set might not be executed by the server as the client expects.
	Validation *pla.ValueMap `idl:"name:pValidation" json:"validation"`
	// pCollector: Receives the newly created data collector.
	Collector *pla.DataCollector `idl:"name:pCollector" json:"collector"`
	// Return: The CreateDataCollectorFromXml return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateDataCollectorFromXMLResponse) xxx_ToOp(ctx context.Context) *xxx_CreateDataCollectorFromXMLOperation {
	if o == nil {
		return &xxx_CreateDataCollectorFromXMLOperation{}
	}
	return &xxx_CreateDataCollectorFromXMLOperation{
		That:       o.That,
		Validation: o.Validation,
		Collector:  o.Collector,
		Return:     o.Return,
	}
}

func (o *CreateDataCollectorFromXMLResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateDataCollectorFromXMLOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Validation = op.Validation
	o.Collector = op.Collector
	o.Return = op.Return
}
func (o *CreateDataCollectorFromXMLResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateDataCollectorFromXMLResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateDataCollectorFromXMLOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateDataCollectorOperation structure represents the CreateDataCollector operation
type xxx_CreateDataCollectorOperation struct {
	This      *dcom.ORPCThis        `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat        `idl:"name:That" json:"that"`
	Type      pla.DataCollectorType `idl:"name:Type" json:"type"`
	Collector *pla.DataCollector    `idl:"name:Collector" json:"collector"`
	Return    int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateDataCollectorOperation) OpNum() int { return 15 }

func (o *xxx_CreateDataCollectorOperation) OpName() string {
	return "/IDataCollectorCollection/v0/CreateDataCollector"
}

func (o *xxx_CreateDataCollectorOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateDataCollectorOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Type {in} (1:{alias=DataCollectorType}(enum))
	{
		if err := w.WriteData(uint16(o.Type)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateDataCollectorOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Type {in} (1:{alias=DataCollectorType}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Type)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateDataCollectorOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateDataCollectorOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Collector {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IDataCollector}(interface))
	{
		if o.Collector != nil {
			_ptr_Collector := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Collector != nil {
					if err := o.Collector.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&pla.DataCollector{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Collector, _ptr_Collector); err != nil {
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

func (o *xxx_CreateDataCollectorOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Collector {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IDataCollector}(interface))
	{
		_ptr_Collector := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Collector == nil {
				o.Collector = &pla.DataCollector{}
			}
			if err := o.Collector.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Collector := func(ptr interface{}) { o.Collector = *ptr.(**pla.DataCollector) }
		if err := w.ReadPointer(&o.Collector, _s_Collector, _ptr_Collector); err != nil {
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

// CreateDataCollectorRequest structure represents the CreateDataCollector operation request
type CreateDataCollectorRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// Type: Supplies the type of data collector to create. For possible data collector
	// types, see the DataCollectorType enumeration in section 2.2.2.5.
	Type pla.DataCollectorType `idl:"name:Type" json:"type"`
}

func (o *CreateDataCollectorRequest) xxx_ToOp(ctx context.Context) *xxx_CreateDataCollectorOperation {
	if o == nil {
		return &xxx_CreateDataCollectorOperation{}
	}
	return &xxx_CreateDataCollectorOperation{
		This: o.This,
		Type: o.Type,
	}
}

func (o *CreateDataCollectorRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateDataCollectorOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Type = op.Type
}
func (o *CreateDataCollectorRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateDataCollectorRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateDataCollectorOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateDataCollectorResponse structure represents the CreateDataCollector operation response
type CreateDataCollectorResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Collector: Receives the newly created data collector.
	Collector *pla.DataCollector `idl:"name:Collector" json:"collector"`
	// Return: The CreateDataCollector return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateDataCollectorResponse) xxx_ToOp(ctx context.Context) *xxx_CreateDataCollectorOperation {
	if o == nil {
		return &xxx_CreateDataCollectorOperation{}
	}
	return &xxx_CreateDataCollectorOperation{
		That:      o.That,
		Collector: o.Collector,
		Return:    o.Return,
	}
}

func (o *CreateDataCollectorResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateDataCollectorOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Collector = op.Collector
	o.Return = op.Return
}
func (o *CreateDataCollectorResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateDataCollectorResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateDataCollectorOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
