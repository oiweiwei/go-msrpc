package ipersistmoniker

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
	urlmon "github.com/oiweiwei/go-msrpc/msrpc/dcom/urlmon"
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
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
	_ = dtyp.GoPackage
	_ = urlmon.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/urlmon"
)

var (
	// IPersistMoniker interface identifier 79eac9c9-baf9-11ce-8c82-00aa004ba90b
	PersistMonikerIID = &dcom.IID{Data1: 0x79eac9c9, Data2: 0xbaf9, Data3: 0x11ce, Data4: []byte{0x8c, 0x82, 0x00, 0xaa, 0x00, 0x4b, 0xa9, 0x0b}}
	// Syntax UUID
	PersistMonikerSyntaxUUID = &uuid.UUID{TimeLow: 0x79eac9c9, TimeMid: 0xbaf9, TimeHiAndVersion: 0x11ce, ClockSeqHiAndReserved: 0x8c, ClockSeqLow: 0x82, Node: [6]uint8{0x0, 0xaa, 0x0, 0x4b, 0xa9, 0xb}}
	// Syntax ID
	PersistMonikerSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: PersistMonikerSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IPersistMoniker interface.
type PersistMonikerClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// GetClassID operation.
	GetClassID(context.Context, *GetClassIDRequest, ...dcerpc.CallOption) (*GetClassIDResponse, error)

	// IsDirty operation.
	IsDirty(context.Context, *IsDirtyRequest, ...dcerpc.CallOption) (*IsDirtyResponse, error)

	// Load operation.
	Load(context.Context, *LoadRequest, ...dcerpc.CallOption) (*LoadResponse, error)

	// Save operation.
	Save(context.Context, *SaveRequest, ...dcerpc.CallOption) (*SaveResponse, error)

	// SaveCompleted operation.
	SaveCompleted(context.Context, *SaveCompletedRequest, ...dcerpc.CallOption) (*SaveCompletedResponse, error)

	// GetCurMoniker operation.
	GetCurMoniker(context.Context, *GetCurMonikerRequest, ...dcerpc.CallOption) (*GetCurMonikerResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) PersistMonikerClient
}

type xxx_DefaultPersistMonikerClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultPersistMonikerClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultPersistMonikerClient) GetClassID(ctx context.Context, in *GetClassIDRequest, opts ...dcerpc.CallOption) (*GetClassIDResponse, error) {
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
	out := &GetClassIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPersistMonikerClient) IsDirty(ctx context.Context, in *IsDirtyRequest, opts ...dcerpc.CallOption) (*IsDirtyResponse, error) {
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
	out := &IsDirtyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPersistMonikerClient) Load(ctx context.Context, in *LoadRequest, opts ...dcerpc.CallOption) (*LoadResponse, error) {
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
	out := &LoadResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPersistMonikerClient) Save(ctx context.Context, in *SaveRequest, opts ...dcerpc.CallOption) (*SaveResponse, error) {
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
	out := &SaveResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPersistMonikerClient) SaveCompleted(ctx context.Context, in *SaveCompletedRequest, opts ...dcerpc.CallOption) (*SaveCompletedResponse, error) {
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
	out := &SaveCompletedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPersistMonikerClient) GetCurMoniker(ctx context.Context, in *GetCurMonikerRequest, opts ...dcerpc.CallOption) (*GetCurMonikerResponse, error) {
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
	out := &GetCurMonikerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPersistMonikerClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultPersistMonikerClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultPersistMonikerClient) IPID(ctx context.Context, ipid *dcom.IPID) PersistMonikerClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultPersistMonikerClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewPersistMonikerClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (PersistMonikerClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(PersistMonikerSyntaxV0_0))...)
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
	return &xxx_DefaultPersistMonikerClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetClassIDOperation structure represents the GetClassID operation
type xxx_GetClassIDOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	ClassID *dtyp.GUID     `idl:"name:pClassID" json:"class_id"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetClassIDOperation) OpNum() int { return 3 }

func (o *xxx_GetClassIDOperation) OpName() string { return "/IPersistMoniker/v0/GetClassID" }

func (o *xxx_GetClassIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetClassIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetClassIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pClassID {out} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ClassID != nil {
			if err := o.ClassID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
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

func (o *xxx_GetClassIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pClassID {out} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ClassID == nil {
			o.ClassID = &dtyp.GUID{}
		}
		if err := o.ClassID.UnmarshalNDR(ctx, w); err != nil {
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

// GetClassIDRequest structure represents the GetClassID operation request
type GetClassIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetClassIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetClassIDOperation) *xxx_GetClassIDOperation {
	if op == nil {
		op = &xxx_GetClassIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetClassIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetClassIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetClassIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetClassIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClassIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetClassIDResponse structure represents the GetClassID operation response
type GetClassIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	ClassID *dtyp.GUID     `idl:"name:pClassID" json:"class_id"`
	// Return: The GetClassID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetClassIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetClassIDOperation) *xxx_GetClassIDOperation {
	if op == nil {
		op = &xxx_GetClassIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ClassID = o.ClassID
	op.Return = o.Return
	return op
}

func (o *GetClassIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetClassIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ClassID = op.ClassID
	o.Return = op.Return
}
func (o *GetClassIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetClassIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClassIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IsDirtyOperation structure represents the IsDirty operation
type xxx_IsDirtyOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_IsDirtyOperation) OpNum() int { return 4 }

func (o *xxx_IsDirtyOperation) OpName() string { return "/IPersistMoniker/v0/IsDirty" }

func (o *xxx_IsDirtyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsDirtyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_IsDirtyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_IsDirtyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsDirtyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_IsDirtyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// IsDirtyRequest structure represents the IsDirty operation request
type IsDirtyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *IsDirtyRequest) xxx_ToOp(ctx context.Context, op *xxx_IsDirtyOperation) *xxx_IsDirtyOperation {
	if op == nil {
		op = &xxx_IsDirtyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *IsDirtyRequest) xxx_FromOp(ctx context.Context, op *xxx_IsDirtyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *IsDirtyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *IsDirtyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsDirtyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsDirtyResponse structure represents the IsDirty operation response
type IsDirtyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The IsDirty return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *IsDirtyResponse) xxx_ToOp(ctx context.Context, op *xxx_IsDirtyOperation) *xxx_IsDirtyOperation {
	if op == nil {
		op = &xxx_IsDirtyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *IsDirtyResponse) xxx_FromOp(ctx context.Context, op *xxx_IsDirtyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *IsDirtyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *IsDirtyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsDirtyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_LoadOperation structure represents the Load operation
type xxx_LoadOperation struct {
	This           *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat      `idl:"name:That" json:"that"`
	FullyAvailable bool                `idl:"name:fFullyAvailable" json:"fully_available"`
	Name           *urlmon.Moniker     `idl:"name:pimkName" json:"name"`
	BindContext    *urlmon.BindContext `idl:"name:pibc" json:"bind_context"`
	Mode           uint32              `idl:"name:grfMode" json:"mode"`
	Return         int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_LoadOperation) OpNum() int { return 5 }

func (o *xxx_LoadOperation) OpName() string { return "/IPersistMoniker/v0/Load" }

func (o *xxx_LoadOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LoadOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// fFullyAvailable {in} (1:{alias=BOOL}(int32))
	{
		if !o.FullyAvailable {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	// pimkName {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		if o.Name != nil {
			_ptr_pimkName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Name != nil {
					if err := o.Name.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Moniker{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Name, _ptr_pimkName); err != nil {
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
	// pibc {in} (1:{pointer=ref}*(1))(2:{alias=IBindCtx}(interface))
	{
		if o.BindContext != nil {
			_ptr_pibc := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.BindContext != nil {
					if err := o.BindContext.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.BindContext{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.BindContext, _ptr_pibc); err != nil {
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
	// grfMode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Mode); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LoadOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// fFullyAvailable {in} (1:{alias=BOOL}(int32))
	{
		var _bFullyAvailable int32
		if err := w.ReadData(&_bFullyAvailable); err != nil {
			return err
		}
		o.FullyAvailable = _bFullyAvailable != 0
	}
	// pimkName {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		_ptr_pimkName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Name == nil {
				o.Name = &urlmon.Moniker{}
			}
			if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pimkName := func(ptr interface{}) { o.Name = *ptr.(**urlmon.Moniker) }
		if err := w.ReadPointer(&o.Name, _s_pimkName, _ptr_pimkName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pibc {in} (1:{pointer=ref}*(1))(2:{alias=IBindCtx}(interface))
	{
		_ptr_pibc := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.BindContext == nil {
				o.BindContext = &urlmon.BindContext{}
			}
			if err := o.BindContext.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pibc := func(ptr interface{}) { o.BindContext = *ptr.(**urlmon.BindContext) }
		if err := w.ReadPointer(&o.BindContext, _s_pibc, _ptr_pibc); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// grfMode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Mode); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LoadOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LoadOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_LoadOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// LoadRequest structure represents the Load operation request
type LoadRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis      `idl:"name:This" json:"this"`
	FullyAvailable bool                `idl:"name:fFullyAvailable" json:"fully_available"`
	Name           *urlmon.Moniker     `idl:"name:pimkName" json:"name"`
	BindContext    *urlmon.BindContext `idl:"name:pibc" json:"bind_context"`
	Mode           uint32              `idl:"name:grfMode" json:"mode"`
}

func (o *LoadRequest) xxx_ToOp(ctx context.Context, op *xxx_LoadOperation) *xxx_LoadOperation {
	if op == nil {
		op = &xxx_LoadOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.FullyAvailable = o.FullyAvailable
	op.Name = o.Name
	op.BindContext = o.BindContext
	op.Mode = o.Mode
	return op
}

func (o *LoadRequest) xxx_FromOp(ctx context.Context, op *xxx_LoadOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.FullyAvailable = op.FullyAvailable
	o.Name = op.Name
	o.BindContext = op.BindContext
	o.Mode = op.Mode
}
func (o *LoadRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *LoadRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LoadOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LoadResponse structure represents the Load operation response
type LoadResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Load return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *LoadResponse) xxx_ToOp(ctx context.Context, op *xxx_LoadOperation) *xxx_LoadOperation {
	if op == nil {
		op = &xxx_LoadOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *LoadResponse) xxx_FromOp(ctx context.Context, op *xxx_LoadOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *LoadResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *LoadResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LoadOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SaveOperation structure represents the Save operation
type xxx_SaveOperation struct {
	This        *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat      `idl:"name:That" json:"that"`
	Name        *urlmon.Moniker     `idl:"name:pimkName" json:"name"`
	BindContext *urlmon.BindContext `idl:"name:pbc" json:"bind_context"`
	Remember    bool                `idl:"name:fRemember" json:"remember"`
	Return      int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_SaveOperation) OpNum() int { return 6 }

func (o *xxx_SaveOperation) OpName() string { return "/IPersistMoniker/v0/Save" }

func (o *xxx_SaveOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SaveOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pimkName {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		if o.Name != nil {
			_ptr_pimkName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Name != nil {
					if err := o.Name.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Moniker{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Name, _ptr_pimkName); err != nil {
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
	// pbc {in} (1:{pointer=ref}*(1))(2:{alias=IBindCtx}(interface))
	{
		if o.BindContext != nil {
			_ptr_pbc := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.BindContext != nil {
					if err := o.BindContext.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.BindContext{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.BindContext, _ptr_pbc); err != nil {
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
	// fRemember {in} (1:{alias=BOOL}(int32))
	{
		if !o.Remember {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_SaveOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pimkName {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		_ptr_pimkName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Name == nil {
				o.Name = &urlmon.Moniker{}
			}
			if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pimkName := func(ptr interface{}) { o.Name = *ptr.(**urlmon.Moniker) }
		if err := w.ReadPointer(&o.Name, _s_pimkName, _ptr_pimkName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbc {in} (1:{pointer=ref}*(1))(2:{alias=IBindCtx}(interface))
	{
		_ptr_pbc := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.BindContext == nil {
				o.BindContext = &urlmon.BindContext{}
			}
			if err := o.BindContext.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbc := func(ptr interface{}) { o.BindContext = *ptr.(**urlmon.BindContext) }
		if err := w.ReadPointer(&o.BindContext, _s_pbc, _ptr_pbc); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// fRemember {in} (1:{alias=BOOL}(int32))
	{
		var _bRemember int32
		if err := w.ReadData(&_bRemember); err != nil {
			return err
		}
		o.Remember = _bRemember != 0
	}
	return nil
}

func (o *xxx_SaveOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SaveOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SaveOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SaveRequest structure represents the Save operation request
type SaveRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis      `idl:"name:This" json:"this"`
	Name        *urlmon.Moniker     `idl:"name:pimkName" json:"name"`
	BindContext *urlmon.BindContext `idl:"name:pbc" json:"bind_context"`
	Remember    bool                `idl:"name:fRemember" json:"remember"`
}

func (o *SaveRequest) xxx_ToOp(ctx context.Context, op *xxx_SaveOperation) *xxx_SaveOperation {
	if op == nil {
		op = &xxx_SaveOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Name = o.Name
	op.BindContext = o.BindContext
	op.Remember = o.Remember
	return op
}

func (o *SaveRequest) xxx_FromOp(ctx context.Context, op *xxx_SaveOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Name = op.Name
	o.BindContext = op.BindContext
	o.Remember = op.Remember
}
func (o *SaveRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SaveRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SaveOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SaveResponse structure represents the Save operation response
type SaveResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Save return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SaveResponse) xxx_ToOp(ctx context.Context, op *xxx_SaveOperation) *xxx_SaveOperation {
	if op == nil {
		op = &xxx_SaveOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SaveResponse) xxx_FromOp(ctx context.Context, op *xxx_SaveOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SaveResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SaveResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SaveOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SaveCompletedOperation structure represents the SaveCompleted operation
type xxx_SaveCompletedOperation struct {
	This        *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat      `idl:"name:That" json:"that"`
	Name        *urlmon.Moniker     `idl:"name:pimkName" json:"name"`
	BindContext *urlmon.BindContext `idl:"name:pibc" json:"bind_context"`
	Return      int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_SaveCompletedOperation) OpNum() int { return 7 }

func (o *xxx_SaveCompletedOperation) OpName() string { return "/IPersistMoniker/v0/SaveCompleted" }

func (o *xxx_SaveCompletedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SaveCompletedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pimkName {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		if o.Name != nil {
			_ptr_pimkName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Name != nil {
					if err := o.Name.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Moniker{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Name, _ptr_pimkName); err != nil {
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
	// pibc {in} (1:{pointer=ref}*(1))(2:{alias=IBindCtx}(interface))
	{
		if o.BindContext != nil {
			_ptr_pibc := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.BindContext != nil {
					if err := o.BindContext.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.BindContext{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.BindContext, _ptr_pibc); err != nil {
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

func (o *xxx_SaveCompletedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pimkName {in} (1:{pointer=ref}*(1))(2:{alias=IMoniker}(interface))
	{
		_ptr_pimkName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Name == nil {
				o.Name = &urlmon.Moniker{}
			}
			if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pimkName := func(ptr interface{}) { o.Name = *ptr.(**urlmon.Moniker) }
		if err := w.ReadPointer(&o.Name, _s_pimkName, _ptr_pimkName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pibc {in} (1:{pointer=ref}*(1))(2:{alias=IBindCtx}(interface))
	{
		_ptr_pibc := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.BindContext == nil {
				o.BindContext = &urlmon.BindContext{}
			}
			if err := o.BindContext.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pibc := func(ptr interface{}) { o.BindContext = *ptr.(**urlmon.BindContext) }
		if err := w.ReadPointer(&o.BindContext, _s_pibc, _ptr_pibc); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SaveCompletedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SaveCompletedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SaveCompletedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SaveCompletedRequest structure represents the SaveCompleted operation request
type SaveCompletedRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis      `idl:"name:This" json:"this"`
	Name        *urlmon.Moniker     `idl:"name:pimkName" json:"name"`
	BindContext *urlmon.BindContext `idl:"name:pibc" json:"bind_context"`
}

func (o *SaveCompletedRequest) xxx_ToOp(ctx context.Context, op *xxx_SaveCompletedOperation) *xxx_SaveCompletedOperation {
	if op == nil {
		op = &xxx_SaveCompletedOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Name = o.Name
	op.BindContext = o.BindContext
	return op
}

func (o *SaveCompletedRequest) xxx_FromOp(ctx context.Context, op *xxx_SaveCompletedOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Name = op.Name
	o.BindContext = op.BindContext
}
func (o *SaveCompletedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SaveCompletedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SaveCompletedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SaveCompletedResponse structure represents the SaveCompleted operation response
type SaveCompletedResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SaveCompleted return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SaveCompletedResponse) xxx_ToOp(ctx context.Context, op *xxx_SaveCompletedOperation) *xxx_SaveCompletedOperation {
	if op == nil {
		op = &xxx_SaveCompletedOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SaveCompletedResponse) xxx_FromOp(ctx context.Context, op *xxx_SaveCompletedOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SaveCompletedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SaveCompletedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SaveCompletedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetCurMonikerOperation structure represents the GetCurMoniker operation
type xxx_GetCurMonikerOperation struct {
	This   *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Name   *urlmon.Moniker `idl:"name:ppimkName" json:"name"`
	Return int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCurMonikerOperation) OpNum() int { return 8 }

func (o *xxx_GetCurMonikerOperation) OpName() string { return "/IPersistMoniker/v0/GetCurMoniker" }

func (o *xxx_GetCurMonikerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCurMonikerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetCurMonikerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetCurMonikerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCurMonikerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppimkName {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IMoniker}(interface))
	{
		if o.Name != nil {
			_ptr_ppimkName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Name != nil {
					if err := o.Name.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Moniker{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Name, _ptr_ppimkName); err != nil {
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

func (o *xxx_GetCurMonikerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppimkName {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IMoniker}(interface))
	{
		_ptr_ppimkName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Name == nil {
				o.Name = &urlmon.Moniker{}
			}
			if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppimkName := func(ptr interface{}) { o.Name = *ptr.(**urlmon.Moniker) }
		if err := w.ReadPointer(&o.Name, _s_ppimkName, _ptr_ppimkName); err != nil {
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

// GetCurMonikerRequest structure represents the GetCurMoniker operation request
type GetCurMonikerRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetCurMonikerRequest) xxx_ToOp(ctx context.Context, op *xxx_GetCurMonikerOperation) *xxx_GetCurMonikerOperation {
	if op == nil {
		op = &xxx_GetCurMonikerOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetCurMonikerRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCurMonikerOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetCurMonikerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetCurMonikerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCurMonikerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCurMonikerResponse structure represents the GetCurMoniker operation response
type GetCurMonikerResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Name *urlmon.Moniker `idl:"name:ppimkName" json:"name"`
	// Return: The GetCurMoniker return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCurMonikerResponse) xxx_ToOp(ctx context.Context, op *xxx_GetCurMonikerOperation) *xxx_GetCurMonikerOperation {
	if op == nil {
		op = &xxx_GetCurMonikerOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Name = o.Name
	op.Return = o.Return
	return op
}

func (o *GetCurMonikerResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCurMonikerOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Name = op.Name
	o.Return = op.Return
}
func (o *GetCurMonikerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetCurMonikerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCurMonikerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
