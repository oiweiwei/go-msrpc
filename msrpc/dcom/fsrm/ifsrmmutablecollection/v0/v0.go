package ifsrmmutablecollection

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
	ifsrmcollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmcollection/v0"
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
	_ = (*errors.Error)(nil)
	_ = dcom.GoPackage
	_ = ifsrmcollection.GoPackage
	_ = oaut.GoPackage
	_ = fsrm.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmMutableCollection interface identifier 1bb617b8-3886-49dc-af82-a6c90fa35dda
	MutableCollectionIID = &dcom.IID{Data1: 0x1bb617b8, Data2: 0x3886, Data3: 0x49dc, Data4: []byte{0xaf, 0x82, 0xa6, 0xc9, 0x0f, 0xa3, 0x5d, 0xda}}
	// Syntax UUID
	MutableCollectionSyntaxUUID = &uuid.UUID{TimeLow: 0x1bb617b8, TimeMid: 0x3886, TimeHiAndVersion: 0x49dc, ClockSeqHiAndReserved: 0xaf, ClockSeqLow: 0x82, Node: [6]uint8{0xa6, 0xc9, 0xf, 0xa3, 0x5d, 0xda}}
	// Syntax ID
	MutableCollectionSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: MutableCollectionSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmMutableCollection interface.
type MutableCollectionClient interface {

	// IFsrmCollection retrieval method.
	Collection() ifsrmcollection.CollectionClient

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
	Add(context.Context, *AddRequest, ...dcerpc.CallOption) (*AddResponse, error)

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
	Remove(context.Context, *RemoveRequest, ...dcerpc.CallOption) (*RemoveResponse, error)

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
	RemoveByID(context.Context, *RemoveByIDRequest, ...dcerpc.CallOption) (*RemoveByIDResponse, error)

	// The Clone method returns a copy of the collection of Objects Being Enumerated.
	//
	// Return Values: The method MUST return a nonzero error code. Upon receiving this message,
	// the server MUST return E_NOTIMPL.
	Clone(context.Context, *CloneRequest, ...dcerpc.CallOption) (*CloneResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) MutableCollectionClient
}

type xxx_DefaultMutableCollectionClient struct {
	ifsrmcollection.CollectionClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultMutableCollectionClient) Collection() ifsrmcollection.CollectionClient {
	return o.CollectionClient
}

func (o *xxx_DefaultMutableCollectionClient) Add(ctx context.Context, in *AddRequest, opts ...dcerpc.CallOption) (*AddResponse, error) {
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

func (o *xxx_DefaultMutableCollectionClient) Remove(ctx context.Context, in *RemoveRequest, opts ...dcerpc.CallOption) (*RemoveResponse, error) {
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

func (o *xxx_DefaultMutableCollectionClient) RemoveByID(ctx context.Context, in *RemoveByIDRequest, opts ...dcerpc.CallOption) (*RemoveByIDResponse, error) {
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
	out := &RemoveByIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMutableCollectionClient) Clone(ctx context.Context, in *CloneRequest, opts ...dcerpc.CallOption) (*CloneResponse, error) {
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
	out := &CloneResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMutableCollectionClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultMutableCollectionClient) IPID(ctx context.Context, ipid *dcom.IPID) MutableCollectionClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultMutableCollectionClient{
		CollectionClient: o.CollectionClient.IPID(ctx, ipid),
		cc:               o.cc,
		ipid:             ipid,
	}
}
func NewMutableCollectionClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (MutableCollectionClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(MutableCollectionSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := ifsrmcollection.NewCollectionClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultMutableCollectionClient{
		CollectionClient: base,
		cc:               cc,
		ipid:             ipid,
	}, nil
}

// xxx_AddOperation structure represents the Add operation
type xxx_AddOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Item   *oaut.Variant  `idl:"name:item" json:"item"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_AddOperation) OpNum() int { return 14 }

func (o *xxx_AddOperation) OpName() string { return "/IFsrmMutableCollection/v0/Add" }

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
	// item {in} (1:{alias=VARIANT}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.Item != nil {
			if err := o.Item.MarshalNDR(ctx, w); err != nil {
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
	// item {in} (1:{alias=VARIANT,pointer=ref}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.Item == nil {
			o.Item = &oaut.Variant{}
		}
		if err := o.Item.UnmarshalNDR(ctx, w); err != nil {
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
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// item: A VARIANT structure that contains the IDispatch interface of the object to
	// add to the collection of Objects Being Enumerated.
	Item *oaut.Variant `idl:"name:item" json:"item"`
}

func (o *AddRequest) xxx_ToOp(ctx context.Context) *xxx_AddOperation {
	if o == nil {
		return &xxx_AddOperation{}
	}
	return &xxx_AddOperation{
		This: o.This,
		Item: o.Item,
	}
}

func (o *AddRequest) xxx_FromOp(ctx context.Context, op *xxx_AddOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Item = op.Item
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
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Index  int32          `idl:"name:index" json:"index"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoveOperation) OpNum() int { return 15 }

func (o *xxx_RemoveOperation) OpName() string { return "/IFsrmMutableCollection/v0/Remove" }

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
	// index {in} (1:(int32))
	{
		if err := w.WriteData(o.Index); err != nil {
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
	// index {in} (1:(int32))
	{
		if err := w.ReadData(&o.Index); err != nil {
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
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// index: Contains the position of the object to remove from the collection of Objects
	// Being Enumerated.
	Index int32 `idl:"name:index" json:"index"`
}

func (o *RemoveRequest) xxx_ToOp(ctx context.Context) *xxx_RemoveOperation {
	if o == nil {
		return &xxx_RemoveOperation{}
	}
	return &xxx_RemoveOperation{
		This:  o.This,
		Index: o.Index,
	}
}

func (o *RemoveRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoveOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Index = op.Index
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

// xxx_RemoveByIDOperation structure represents the RemoveById operation
type xxx_RemoveByIDOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	ID     *fsrm.ObjectID `idl:"name:id" json:"id"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoveByIDOperation) OpNum() int { return 16 }

func (o *xxx_RemoveByIDOperation) OpName() string { return "/IFsrmMutableCollection/v0/RemoveById" }

func (o *xxx_RemoveByIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveByIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RemoveByIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_RemoveByIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveByIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RemoveByIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RemoveByIDRequest structure represents the RemoveById operation request
type RemoveByIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// id: The ID to match for identifying the object to be removed.
	ID *fsrm.ObjectID `idl:"name:id" json:"id"`
}

func (o *RemoveByIDRequest) xxx_ToOp(ctx context.Context) *xxx_RemoveByIDOperation {
	if o == nil {
		return &xxx_RemoveByIDOperation{}
	}
	return &xxx_RemoveByIDOperation{
		This: o.This,
		ID:   o.ID,
	}
}

func (o *RemoveByIDRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoveByIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ID = op.ID
}
func (o *RemoveByIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RemoveByIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveByIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoveByIDResponse structure represents the RemoveById operation response
type RemoveByIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RemoveById return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemoveByIDResponse) xxx_ToOp(ctx context.Context) *xxx_RemoveByIDOperation {
	if o == nil {
		return &xxx_RemoveByIDOperation{}
	}
	return &xxx_RemoveByIDOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *RemoveByIDResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoveByIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RemoveByIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RemoveByIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveByIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloneOperation structure represents the Clone operation
type xxx_CloneOperation struct {
	This       *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Collection *fsrm.MutableCollection `idl:"name:collection" json:"collection"`
	Return     int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_CloneOperation) OpNum() int { return 17 }

func (o *xxx_CloneOperation) OpName() string { return "/IFsrmMutableCollection/v0/Clone" }

func (o *xxx_CloneOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloneOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CloneOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_CloneOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloneOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// collection {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmMutableCollection}(interface))
	{
		if o.Collection != nil {
			_ptr_collection := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Collection != nil {
					if err := o.Collection.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.MutableCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Collection, _ptr_collection); err != nil {
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

func (o *xxx_CloneOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// collection {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmMutableCollection}(interface))
	{
		_ptr_collection := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Collection == nil {
				o.Collection = &fsrm.MutableCollection{}
			}
			if err := o.Collection.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_collection := func(ptr interface{}) { o.Collection = *ptr.(**fsrm.MutableCollection) }
		if err := w.ReadPointer(&o.Collection, _s_collection, _ptr_collection); err != nil {
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

// CloneRequest structure represents the Clone operation request
type CloneRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *CloneRequest) xxx_ToOp(ctx context.Context) *xxx_CloneOperation {
	if o == nil {
		return &xxx_CloneOperation{}
	}
	return &xxx_CloneOperation{
		This: o.This,
	}
}

func (o *CloneRequest) xxx_FromOp(ctx context.Context, op *xxx_CloneOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *CloneRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CloneRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloneOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloneResponse structure represents the Clone operation response
type CloneResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// collection: Pointer to an IFsrmMutableCollection interface pointer (section 3.2.4.2.2)
	// that upon successful completion contains the IFsrmMutableCollection pointer of a
	// copy of this collection of Objects Being Enumerated.
	Collection *fsrm.MutableCollection `idl:"name:collection" json:"collection"`
	// Return: The Clone return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CloneResponse) xxx_ToOp(ctx context.Context) *xxx_CloneOperation {
	if o == nil {
		return &xxx_CloneOperation{}
	}
	return &xxx_CloneOperation{
		That:       o.That,
		Collection: o.Collection,
		Return:     o.Return,
	}
}

func (o *CloneResponse) xxx_FromOp(ctx context.Context, op *xxx_CloneOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Collection = op.Collection
	o.Return = op.Return
}
func (o *CloneResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CloneResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloneOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
