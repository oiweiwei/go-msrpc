package ifsrmcollection

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	fsrm "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm"
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
	_ = fsrm.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmCollection interface identifier f76fbf3b-8ddd-4b42-b05a-cb1c3ff1fee8
	CollectionIID = &dcom.IID{Data1: 0xf76fbf3b, Data2: 0x8ddd, Data3: 0x4b42, Data4: []byte{0xb0, 0x5a, 0xcb, 0x1c, 0x3f, 0xf1, 0xfe, 0xe8}}
	// Syntax UUID
	CollectionSyntaxUUID = &uuid.UUID{TimeLow: 0xf76fbf3b, TimeMid: 0x8ddd, TimeHiAndVersion: 0x4b42, ClockSeqHiAndReserved: 0xb0, ClockSeqLow: 0x5a, Node: [6]uint8{0xcb, 0x1c, 0x3f, 0xf1, 0xfe, 0xe8}}
	// Syntax ID
	CollectionSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: CollectionSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmCollection interface.
type CollectionClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// _NewEnum operation.
	Get_NewEnum(context.Context, *Get_NewEnumRequest, ...dcerpc.CallOption) (*Get_NewEnumResponse, error)

	// The Item method returns a pointer to the object at the requested position in the
	// collection of Objects Being Enumerated.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN                |                                                                                  |
	//	|             VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80004003 E_POINTER                | The item parameter is NULL.                                                      |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80131502 COR_E_ARGUMENTOUTOFRANGE | The value of the index parameter is greater than the number of Objects Being     |
	//	|                                     | Enumerated in the collection.                                                    |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	GetItem(context.Context, *GetItemRequest, ...dcerpc.CallOption) (*GetItemResponse, error)

	// The Count method returns the number of objects in the collection of Objects Being
	// Enumerated.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+----------------------+------------------------------+
	//	|        RETURN        |                              |
	//	|      VALUE/CODE      |         DESCRIPTION          |
	//	|                      |                              |
	//	+----------------------+------------------------------+
	//	+----------------------+------------------------------+
	//	| 0x80004003 E_POINTER | The count parameter is NULL. |
	//	+----------------------+------------------------------+
	GetCount(context.Context, *GetCountRequest, ...dcerpc.CallOption) (*GetCountResponse, error)

	// The State method returns the state FsrmCollectionState_Complete.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+----------------------+------------------------------+
	//	|        RETURN        |                              |
	//	|      VALUE/CODE      |         DESCRIPTION          |
	//	|                      |                              |
	//	+----------------------+------------------------------+
	//	+----------------------+------------------------------+
	//	| 0x80004003 E_POINTER | The state parameter is NULL. |
	//	+----------------------+------------------------------+
	GetState(context.Context, *GetStateRequest, ...dcerpc.CallOption) (*GetStateResponse, error)

	// Cancel operation.
	Cancel(context.Context, *CancelRequest, ...dcerpc.CallOption) (*CancelResponse, error)

	// WaitForCompletion operation.
	WaitForCompletion(context.Context, *WaitForCompletionRequest, ...dcerpc.CallOption) (*WaitForCompletionResponse, error)

	// The GetById method returns the object from the collection of Objects Being Enumerated
	// (section 3.2.1.11) whose ID matches the specified id.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	|           RETURN            |                                                                                  |
	//	|         VALUE/CODE          |                                   DESCRIPTION                                    |
	//	|                             |                                                                                  |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND | An object with the specified ID was not found in the collection.                 |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG     | This code is returned for the following reasons: The entry parameter is          |
	//	|                             | NULL. The objects in the collection of Objects Being Enumerated are not          |
	//	|                             | one of the following interfaces: IFsrmFileScreen, IFsrmFileScreenException,      |
	//	|                             | IFsrmFileScreenTemplate, IFsrmFileGroup, IFsrmQuota, IFsrmQuotaTemplate,         |
	//	|                             | IFsrmAction, IFsrmReportJob, IFsrmReport, IFsrmClassifcationRule,                |
	//	|                             | IFsrmPropertyDefinition, IFsrmPipelineModuleDefinition or                        |
	//	|                             | IFsrmFileManagementJob.                                                          |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	GetByID(context.Context, *GetByIDRequest, ...dcerpc.CallOption) (*GetByIDResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) CollectionClient
}

type xxx_DefaultCollectionClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultCollectionClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultCollectionClient) Get_NewEnum(ctx context.Context, in *Get_NewEnumRequest, opts ...dcerpc.CallOption) (*Get_NewEnumResponse, error) {
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
	out := &Get_NewEnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCollectionClient) GetItem(ctx context.Context, in *GetItemRequest, opts ...dcerpc.CallOption) (*GetItemResponse, error) {
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
	out := &GetItemResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCollectionClient) GetCount(ctx context.Context, in *GetCountRequest, opts ...dcerpc.CallOption) (*GetCountResponse, error) {
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
	out := &GetCountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCollectionClient) GetState(ctx context.Context, in *GetStateRequest, opts ...dcerpc.CallOption) (*GetStateResponse, error) {
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
	out := &GetStateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCollectionClient) Cancel(ctx context.Context, in *CancelRequest, opts ...dcerpc.CallOption) (*CancelResponse, error) {
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
	out := &CancelResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCollectionClient) WaitForCompletion(ctx context.Context, in *WaitForCompletionRequest, opts ...dcerpc.CallOption) (*WaitForCompletionResponse, error) {
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
	out := &WaitForCompletionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCollectionClient) GetByID(ctx context.Context, in *GetByIDRequest, opts ...dcerpc.CallOption) (*GetByIDResponse, error) {
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
	out := &GetByIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCollectionClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultCollectionClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultCollectionClient) IPID(ctx context.Context, ipid *dcom.IPID) CollectionClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultCollectionClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewCollectionClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (CollectionClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(CollectionSyntaxV0_0))...)
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
	return &xxx_DefaultCollectionClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_Get_NewEnumOperation structure represents the _NewEnum operation
type xxx_Get_NewEnumOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Unknown *dcom.Unknown  `idl:"name:unknown" json:"unknown"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_Get_NewEnumOperation) OpNum() int { return 7 }

func (o *xxx_Get_NewEnumOperation) OpName() string { return "/IFsrmCollection/v0/_NewEnum" }

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
	// unknown {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IUnknown}(interface))
	{
		if o.Unknown != nil {
			_ptr_unknown := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Unknown != nil {
					if err := o.Unknown.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.Unknown{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Unknown, _ptr_unknown); err != nil {
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
	// unknown {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IUnknown}(interface))
	{
		_ptr_unknown := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Unknown == nil {
				o.Unknown = &dcom.Unknown{}
			}
			if err := o.Unknown.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_unknown := func(ptr interface{}) { o.Unknown = *ptr.(**dcom.Unknown) }
		if err := w.ReadPointer(&o.Unknown, _s_unknown, _ptr_unknown); err != nil {
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

func (o *Get_NewEnumRequest) xxx_ToOp(ctx context.Context, op *xxx_Get_NewEnumOperation) *xxx_Get_NewEnumOperation {
	if op == nil {
		op = &xxx_Get_NewEnumOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *Get_NewEnumRequest) xxx_FromOp(ctx context.Context, op *xxx_Get_NewEnumOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *Get_NewEnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
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
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Unknown *dcom.Unknown  `idl:"name:unknown" json:"unknown"`
	// Return: The _NewEnum return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *Get_NewEnumResponse) xxx_ToOp(ctx context.Context, op *xxx_Get_NewEnumOperation) *xxx_Get_NewEnumOperation {
	if op == nil {
		op = &xxx_Get_NewEnumOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Unknown = o.Unknown
	op.Return = o.Return
	return op
}

func (o *Get_NewEnumResponse) xxx_FromOp(ctx context.Context, op *xxx_Get_NewEnumOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Unknown = op.Unknown
	o.Return = op.Return
}
func (o *Get_NewEnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *Get_NewEnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Get_NewEnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetItemOperation structure represents the Item operation
type xxx_GetItemOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Index  int32          `idl:"name:index" json:"index"`
	Item   *oaut.Variant  `idl:"name:item" json:"item"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetItemOperation) OpNum() int { return 8 }

func (o *xxx_GetItemOperation) OpName() string { return "/IFsrmCollection/v0/Item" }

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
	// index {in} (1:(int32))
	{
		if err := w.WriteData(o.Index); err != nil {
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
	// index {in} (1:(int32))
	{
		if err := w.ReadData(&o.Index); err != nil {
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
	// item {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Item != nil {
			_ptr_item := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Item != nil {
					if err := o.Item.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Item, _ptr_item); err != nil {
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
	// item {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_item := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Item == nil {
				o.Item = &oaut.Variant{}
			}
			if err := o.Item.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_item := func(ptr interface{}) { o.Item = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Item, _s_item, _ptr_item); err != nil {
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
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// index: The position of the item in the collection of Objects Being Enumerated to
	// return.
	Index int32 `idl:"name:index" json:"index"`
}

func (o *GetItemRequest) xxx_ToOp(ctx context.Context, op *xxx_GetItemOperation) *xxx_GetItemOperation {
	if op == nil {
		op = &xxx_GetItemOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Index = o.Index
	return op
}

func (o *GetItemRequest) xxx_FromOp(ctx context.Context, op *xxx_GetItemOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Index = op.Index
}
func (o *GetItemRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
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
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// item: Pointer to a VARIANT structure. Upon successful completion, receives the pointer
	// to the IDispatch interface for the object.
	Item *oaut.Variant `idl:"name:item" json:"item"`
	// Return: The Item return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetItemResponse) xxx_ToOp(ctx context.Context, op *xxx_GetItemOperation) *xxx_GetItemOperation {
	if op == nil {
		op = &xxx_GetItemOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Item = o.Item
	op.Return = o.Return
	return op
}

func (o *GetItemResponse) xxx_FromOp(ctx context.Context, op *xxx_GetItemOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Item = op.Item
	o.Return = op.Return
}
func (o *GetItemResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetItemResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetItemOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetCountOperation structure represents the Count operation
type xxx_GetCountOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Count  int32          `idl:"name:count" json:"count"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCountOperation) OpNum() int { return 9 }

func (o *xxx_GetCountOperation) OpName() string { return "/IFsrmCollection/v0/Count" }

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
	// count {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.Count); err != nil {
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
	// count {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.Count); err != nil {
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

func (o *GetCountRequest) xxx_ToOp(ctx context.Context, op *xxx_GetCountOperation) *xxx_GetCountOperation {
	if op == nil {
		op = &xxx_GetCountOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetCountRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCountOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetCountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
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
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// count: Pointer to a variable that upon successful completion receives the number
	// of objects in the collection of Objects Being Enumerated.
	Count int32 `idl:"name:count" json:"count"`
	// Return: The Count return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCountResponse) xxx_ToOp(ctx context.Context, op *xxx_GetCountOperation) *xxx_GetCountOperation {
	if op == nil {
		op = &xxx_GetCountOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Count = o.Count
	op.Return = o.Return
	return op
}

func (o *GetCountResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCountOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Count = op.Count
	o.Return = op.Return
}
func (o *GetCountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetCountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetStateOperation structure represents the State operation
type xxx_GetStateOperation struct {
	This   *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat       `idl:"name:That" json:"that"`
	State  fsrm.CollectionState `idl:"name:state" json:"state"`
	Return int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_GetStateOperation) OpNum() int { return 10 }

func (o *xxx_GetStateOperation) OpName() string { return "/IFsrmCollection/v0/State" }

func (o *xxx_GetStateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetStateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetStateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetStateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetStateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// state {out, retval} (1:{pointer=ref}*(1))(2:{alias=FsrmCollectionState}(enum))
	{
		if err := w.WriteEnum(uint16(o.State)); err != nil {
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

func (o *xxx_GetStateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// state {out, retval} (1:{pointer=ref}*(1))(2:{alias=FsrmCollectionState}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.State)); err != nil {
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

// GetStateRequest structure represents the State operation request
type GetStateRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetStateRequest) xxx_ToOp(ctx context.Context, op *xxx_GetStateOperation) *xxx_GetStateOperation {
	if op == nil {
		op = &xxx_GetStateOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetStateRequest) xxx_FromOp(ctx context.Context, op *xxx_GetStateOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetStateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetStateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetStateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetStateResponse structure represents the State operation response
type GetStateResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// state: Pointer to a variable that upon completion contains the state FsrmCollectionState_Complete.
	State fsrm.CollectionState `idl:"name:state" json:"state"`
	// Return: The State return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetStateResponse) xxx_ToOp(ctx context.Context, op *xxx_GetStateOperation) *xxx_GetStateOperation {
	if op == nil {
		op = &xxx_GetStateOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.State = o.State
	op.Return = o.Return
	return op
}

func (o *GetStateResponse) xxx_FromOp(ctx context.Context, op *xxx_GetStateOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.State = op.State
	o.Return = op.Return
}
func (o *GetStateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetStateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetStateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CancelOperation structure represents the Cancel operation
type xxx_CancelOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CancelOperation) OpNum() int { return 11 }

func (o *xxx_CancelOperation) OpName() string { return "/IFsrmCollection/v0/Cancel" }

func (o *xxx_CancelOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CancelOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_CancelOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CancelOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CancelRequest structure represents the Cancel operation request
type CancelRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *CancelRequest) xxx_ToOp(ctx context.Context, op *xxx_CancelOperation) *xxx_CancelOperation {
	if op == nil {
		op = &xxx_CancelOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *CancelRequest) xxx_FromOp(ctx context.Context, op *xxx_CancelOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *CancelRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CancelRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CancelOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CancelResponse structure represents the Cancel operation response
type CancelResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Cancel return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CancelResponse) xxx_ToOp(ctx context.Context, op *xxx_CancelOperation) *xxx_CancelOperation {
	if op == nil {
		op = &xxx_CancelOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CancelResponse) xxx_FromOp(ctx context.Context, op *xxx_CancelOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CancelResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CancelResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CancelOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WaitForCompletionOperation structure represents the WaitForCompletion operation
type xxx_WaitForCompletionOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	WaitSeconds int32          `idl:"name:waitSeconds" json:"wait_seconds"`
	Completed   int16          `idl:"name:completed" json:"completed"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_WaitForCompletionOperation) OpNum() int { return 12 }

func (o *xxx_WaitForCompletionOperation) OpName() string {
	return "/IFsrmCollection/v0/WaitForCompletion"
}

func (o *xxx_WaitForCompletionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitForCompletionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// waitSeconds {in} (1:(int32))
	{
		if err := w.WriteData(o.WaitSeconds); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitForCompletionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// waitSeconds {in} (1:(int32))
	{
		if err := w.ReadData(&o.WaitSeconds); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitForCompletionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitForCompletionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// completed {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Completed); err != nil {
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

func (o *xxx_WaitForCompletionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// completed {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Completed); err != nil {
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

// WaitForCompletionRequest structure represents the WaitForCompletion operation request
type WaitForCompletionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	WaitSeconds int32          `idl:"name:waitSeconds" json:"wait_seconds"`
}

func (o *WaitForCompletionRequest) xxx_ToOp(ctx context.Context, op *xxx_WaitForCompletionOperation) *xxx_WaitForCompletionOperation {
	if op == nil {
		op = &xxx_WaitForCompletionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.WaitSeconds = o.WaitSeconds
	return op
}

func (o *WaitForCompletionRequest) xxx_FromOp(ctx context.Context, op *xxx_WaitForCompletionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.WaitSeconds = op.WaitSeconds
}
func (o *WaitForCompletionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WaitForCompletionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WaitForCompletionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WaitForCompletionResponse structure represents the WaitForCompletion operation response
type WaitForCompletionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Completed int16          `idl:"name:completed" json:"completed"`
	// Return: The WaitForCompletion return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *WaitForCompletionResponse) xxx_ToOp(ctx context.Context, op *xxx_WaitForCompletionOperation) *xxx_WaitForCompletionOperation {
	if op == nil {
		op = &xxx_WaitForCompletionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Completed = o.Completed
	op.Return = o.Return
	return op
}

func (o *WaitForCompletionResponse) xxx_FromOp(ctx context.Context, op *xxx_WaitForCompletionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Completed = op.Completed
	o.Return = op.Return
}
func (o *WaitForCompletionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WaitForCompletionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WaitForCompletionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetByIDOperation structure represents the GetById operation
type xxx_GetByIDOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	ID     *fsrm.ObjectID `idl:"name:id" json:"id"`
	Entry  *oaut.Variant  `idl:"name:entry" json:"entry"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetByIDOperation) OpNum() int { return 13 }

func (o *xxx_GetByIDOperation) OpName() string { return "/IFsrmCollection/v0/GetById" }

func (o *xxx_GetByIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetByIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// id {in} (1:{alias=FSRM_OBJECT_ID, names=GUID}(struct))
	{
		if o.ID != nil {
			if err := o.ID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&fsrm.ObjectID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetByIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// id {in} (1:{alias=FSRM_OBJECT_ID, names=GUID}(struct))
	{
		if o.ID == nil {
			o.ID = &fsrm.ObjectID{}
		}
		if err := o.ID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetByIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetByIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// entry {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Entry != nil {
			_ptr_entry := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Entry != nil {
					if err := o.Entry.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Entry, _ptr_entry); err != nil {
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

func (o *xxx_GetByIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// entry {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_entry := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Entry == nil {
				o.Entry = &oaut.Variant{}
			}
			if err := o.Entry.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_entry := func(ptr interface{}) { o.Entry = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Entry, _s_entry, _ptr_entry); err != nil {
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

// GetByIDRequest structure represents the GetById operation request
type GetByIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// id: The ID to use for identifying the object to be returned.
	ID *fsrm.ObjectID `idl:"name:id" json:"id"`
}

func (o *GetByIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetByIDOperation) *xxx_GetByIDOperation {
	if op == nil {
		op = &xxx_GetByIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ID = o.ID
	return op
}

func (o *GetByIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetByIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ID = op.ID
}
func (o *GetByIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetByIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetByIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetByIDResponse structure represents the GetById operation response
type GetByIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// entry: Pointer to a VARIANT structure. Upon successful completion, it SHOULD contain
	// the IDispatch interface of the object from the collection of Objects Being Enumerated
	// whose ID matches the specified id.
	Entry *oaut.Variant `idl:"name:entry" json:"entry"`
	// Return: The GetById return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetByIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetByIDOperation) *xxx_GetByIDOperation {
	if op == nil {
		op = &xxx_GetByIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Entry = o.Entry
	op.Return = o.Return
	return op
}

func (o *GetByIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetByIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Entry = op.Entry
	o.Return = op.Return
}
func (o *GetByIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetByIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetByIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
