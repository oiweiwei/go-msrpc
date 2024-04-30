package ifsrmderivedobjectsresult

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
	_ = fsrm.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmDerivedObjectsResult interface identifier 39322a2d-38ee-4d0d-8095-421a80849a82
	DerivedObjectsResultIID = &dcom.IID{Data1: 0x39322a2d, Data2: 0x38ee, Data3: 0x4d0d, Data4: []byte{0x80, 0x95, 0x42, 0x1a, 0x80, 0x84, 0x9a, 0x82}}
	// Syntax UUID
	DerivedObjectsResultSyntaxUUID = &uuid.UUID{TimeLow: 0x39322a2d, TimeMid: 0x38ee, TimeHiAndVersion: 0x4d0d, ClockSeqHiAndReserved: 0x80, ClockSeqLow: 0x95, Node: [6]uint8{0x42, 0x1a, 0x80, 0x84, 0x9a, 0x82}}
	// Syntax ID
	DerivedObjectsResultSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: DerivedObjectsResultSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmDerivedObjectsResult interface.
type DerivedObjectsResultClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// The DerivedObjects (get) method returns the collection of derived objects for the
	// calling template.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+---------------------------------------+
	//	|         RETURN          |                                       |
	//	|       VALUE/CODE        |              DESCRIPTION              |
	//	|                         |                                       |
	//	+-------------------------+---------------------------------------+
	//	+-------------------------+---------------------------------------+
	//	| 0x80070057 E_INVALIDARG | The derivedObjects parameter is NULL. |
	//	+-------------------------+---------------------------------------+
	GetDerivedObjects(context.Context, *GetDerivedObjectsRequest, ...dcerpc.CallOption) (*GetDerivedObjectsResponse, error)

	// The Results (get) method returns the collection HRESULTS received when committing
	// derived objects that were updated as a result of the source template's call to CommitAndUpdateDerived.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+--------------------------------+
	//	|         RETURN          |                                |
	//	|       VALUE/CODE        |          DESCRIPTION           |
	//	|                         |                                |
	//	+-------------------------+--------------------------------+
	//	+-------------------------+--------------------------------+
	//	| 0x80070057 E_INVALIDARG | The results parameter is NULL. |
	//	+-------------------------+--------------------------------+
	GetResults(context.Context, *GetResultsRequest, ...dcerpc.CallOption) (*GetResultsResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) DerivedObjectsResultClient
}

type xxx_DefaultDerivedObjectsResultClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultDerivedObjectsResultClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultDerivedObjectsResultClient) GetDerivedObjects(ctx context.Context, in *GetDerivedObjectsRequest, opts ...dcerpc.CallOption) (*GetDerivedObjectsResponse, error) {
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
	out := &GetDerivedObjectsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDerivedObjectsResultClient) GetResults(ctx context.Context, in *GetResultsRequest, opts ...dcerpc.CallOption) (*GetResultsResponse, error) {
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
	out := &GetResultsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDerivedObjectsResultClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultDerivedObjectsResultClient) IPID(ctx context.Context, ipid *dcom.IPID) DerivedObjectsResultClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultDerivedObjectsResultClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}
func NewDerivedObjectsResultClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (DerivedObjectsResultClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(DerivedObjectsResultSyntaxV0_0))...)
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
	return &xxx_DefaultDerivedObjectsResultClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetDerivedObjectsOperation structure represents the DerivedObjects operation
type xxx_GetDerivedObjectsOperation struct {
	This           *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat   `idl:"name:That" json:"that"`
	DerivedObjects *fsrm.Collection `idl:"name:derivedObjects" json:"derived_objects"`
	Return         int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDerivedObjectsOperation) OpNum() int { return 7 }

func (o *xxx_GetDerivedObjectsOperation) OpName() string {
	return "/IFsrmDerivedObjectsResult/v0/DerivedObjects"
}

func (o *xxx_GetDerivedObjectsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDerivedObjectsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetDerivedObjectsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetDerivedObjectsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDerivedObjectsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// derivedObjects {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCollection}(interface))
	{
		if o.DerivedObjects != nil {
			_ptr_derivedObjects := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DerivedObjects != nil {
					if err := o.DerivedObjects.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.Collection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DerivedObjects, _ptr_derivedObjects); err != nil {
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

func (o *xxx_GetDerivedObjectsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// derivedObjects {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCollection}(interface))
	{
		_ptr_derivedObjects := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DerivedObjects == nil {
				o.DerivedObjects = &fsrm.Collection{}
			}
			if err := o.DerivedObjects.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_derivedObjects := func(ptr interface{}) { o.DerivedObjects = *ptr.(**fsrm.Collection) }
		if err := w.ReadPointer(&o.DerivedObjects, _s_derivedObjects, _ptr_derivedObjects); err != nil {
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

// GetDerivedObjectsRequest structure represents the DerivedObjects operation request
type GetDerivedObjectsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetDerivedObjectsRequest) xxx_ToOp(ctx context.Context) *xxx_GetDerivedObjectsOperation {
	if o == nil {
		return &xxx_GetDerivedObjectsOperation{}
	}
	return &xxx_GetDerivedObjectsOperation{
		This: o.This,
	}
}

func (o *GetDerivedObjectsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDerivedObjectsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetDerivedObjectsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetDerivedObjectsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDerivedObjectsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDerivedObjectsResponse structure represents the DerivedObjects operation response
type GetDerivedObjectsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// derivedObjects: Pointer to an IFsrmCollection interface pointer (section 3.2.4.2.1)
	// that upon completion contains interface pointers for the derived objects that were
	// updated as a result of the source template's call to CommitAndUpdateDerived. A caller
	// MUST release the collection interface received when it is done with it.
	DerivedObjects *fsrm.Collection `idl:"name:derivedObjects" json:"derived_objects"`
	// Return: The DerivedObjects return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDerivedObjectsResponse) xxx_ToOp(ctx context.Context) *xxx_GetDerivedObjectsOperation {
	if o == nil {
		return &xxx_GetDerivedObjectsOperation{}
	}
	return &xxx_GetDerivedObjectsOperation{
		That:           o.That,
		DerivedObjects: o.DerivedObjects,
		Return:         o.Return,
	}
}

func (o *GetDerivedObjectsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDerivedObjectsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DerivedObjects = op.DerivedObjects
	o.Return = op.Return
}
func (o *GetDerivedObjectsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetDerivedObjectsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDerivedObjectsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetResultsOperation structure represents the Results operation
type xxx_GetResultsOperation struct {
	This    *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Results *fsrm.Collection `idl:"name:results" json:"results"`
	Return  int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetResultsOperation) OpNum() int { return 8 }

func (o *xxx_GetResultsOperation) OpName() string { return "/IFsrmDerivedObjectsResult/v0/Results" }

func (o *xxx_GetResultsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResultsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetResultsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetResultsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResultsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// results {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCollection}(interface))
	{
		if o.Results != nil {
			_ptr_results := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Results != nil {
					if err := o.Results.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.Collection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Results, _ptr_results); err != nil {
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

func (o *xxx_GetResultsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// results {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCollection}(interface))
	{
		_ptr_results := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Results == nil {
				o.Results = &fsrm.Collection{}
			}
			if err := o.Results.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_results := func(ptr interface{}) { o.Results = *ptr.(**fsrm.Collection) }
		if err := w.ReadPointer(&o.Results, _s_results, _ptr_results); err != nil {
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

// GetResultsRequest structure represents the Results operation request
type GetResultsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetResultsRequest) xxx_ToOp(ctx context.Context) *xxx_GetResultsOperation {
	if o == nil {
		return &xxx_GetResultsOperation{}
	}
	return &xxx_GetResultsOperation{
		This: o.This,
	}
}

func (o *GetResultsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetResultsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetResultsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetResultsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResultsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetResultsResponse structure represents the Results operation response
type GetResultsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// results: Pointer to an IFsrmCollection interface pointer (section 3.2.4.2.1) that
	// upon completion contains HRESULTS for the committing of derived objects that were
	// updated as a result of the source template's call to CommitAndUpdateDerived. A caller
	// MUST release the collection interface received when it is done with it.
	Results *fsrm.Collection `idl:"name:results" json:"results"`
	// Return: The Results return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetResultsResponse) xxx_ToOp(ctx context.Context) *xxx_GetResultsOperation {
	if o == nil {
		return &xxx_GetResultsOperation{}
	}
	return &xxx_GetResultsOperation{
		That:    o.That,
		Results: o.Results,
		Return:  o.Return,
	}
}

func (o *GetResultsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetResultsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Results = op.Results
	o.Return = op.Return
}
func (o *GetResultsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetResultsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResultsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
