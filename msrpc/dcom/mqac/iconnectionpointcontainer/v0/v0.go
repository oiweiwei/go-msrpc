package iconnectionpointcontainer

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
	mqac "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac"
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
	_ = iunknown.GoPackage
	_ = mqac.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/mqac"
)

var (
	// IConnectionPointContainer interface identifier b196b284-bab4-101a-b69c-00aa00341d07
	ConnectionPointContainerIID = &dcom.IID{Data1: 0xb196b284, Data2: 0xbab4, Data3: 0x101a, Data4: []byte{0xb6, 0x9c, 0x00, 0xaa, 0x00, 0x34, 0x1d, 0x07}}
	// Syntax UUID
	ConnectionPointContainerSyntaxUUID = &uuid.UUID{TimeLow: 0xb196b284, TimeMid: 0xbab4, TimeHiAndVersion: 0x101a, ClockSeqHiAndReserved: 0xb6, ClockSeqLow: 0x9c, Node: [6]uint8{0x0, 0xaa, 0x0, 0x34, 0x1d, 0x7}}
	// Syntax ID
	ConnectionPointContainerSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ConnectionPointContainerSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IConnectionPointContainer interface.
type ConnectionPointContainerClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// EnumConnectionPoints operation.
	EnumConnectionPoints(context.Context, *EnumConnectionPointsRequest, ...dcerpc.CallOption) (*EnumConnectionPointsResponse, error)

	// FindConnectionPoint operation.
	FindConnectionPoint(context.Context, *FindConnectionPointRequest, ...dcerpc.CallOption) (*FindConnectionPointResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ConnectionPointContainerClient
}

type xxx_DefaultConnectionPointContainerClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultConnectionPointContainerClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultConnectionPointContainerClient) EnumConnectionPoints(ctx context.Context, in *EnumConnectionPointsRequest, opts ...dcerpc.CallOption) (*EnumConnectionPointsResponse, error) {
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
	out := &EnumConnectionPointsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConnectionPointContainerClient) FindConnectionPoint(ctx context.Context, in *FindConnectionPointRequest, opts ...dcerpc.CallOption) (*FindConnectionPointResponse, error) {
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
	out := &FindConnectionPointResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConnectionPointContainerClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultConnectionPointContainerClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultConnectionPointContainerClient) IPID(ctx context.Context, ipid *dcom.IPID) ConnectionPointContainerClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultConnectionPointContainerClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewConnectionPointContainerClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ConnectionPointContainerClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ConnectionPointContainerSyntaxV0_0))...)
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
	return &xxx_DefaultConnectionPointContainerClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_EnumConnectionPointsOperation structure represents the EnumConnectionPoints operation
type xxx_EnumConnectionPointsOperation struct {
	This   *dcom.ORPCThis             `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat             `idl:"name:That" json:"that"`
	Enum   *mqac.EnumConnectionPoints `idl:"name:ppEnum" json:"enum"`
	Return int32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumConnectionPointsOperation) OpNum() int { return 3 }

func (o *xxx_EnumConnectionPointsOperation) OpName() string {
	return "/IConnectionPointContainer/v0/EnumConnectionPoints"
}

func (o *xxx_EnumConnectionPointsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumConnectionPointsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_EnumConnectionPointsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_EnumConnectionPointsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumConnectionPointsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumConnectionPoints}(interface))
	{
		if o.Enum != nil {
			_ptr_ppEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Enum != nil {
					if err := o.Enum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.EnumConnectionPoints{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Enum, _ptr_ppEnum); err != nil {
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

func (o *xxx_EnumConnectionPointsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumConnectionPoints}(interface))
	{
		_ptr_ppEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Enum == nil {
				o.Enum = &mqac.EnumConnectionPoints{}
			}
			if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppEnum := func(ptr interface{}) { o.Enum = *ptr.(**mqac.EnumConnectionPoints) }
		if err := w.ReadPointer(&o.Enum, _s_ppEnum, _ptr_ppEnum); err != nil {
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

// EnumConnectionPointsRequest structure represents the EnumConnectionPoints operation request
type EnumConnectionPointsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *EnumConnectionPointsRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumConnectionPointsOperation) *xxx_EnumConnectionPointsOperation {
	if op == nil {
		op = &xxx_EnumConnectionPointsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *EnumConnectionPointsRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumConnectionPointsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *EnumConnectionPointsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumConnectionPointsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumConnectionPointsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumConnectionPointsResponse structure represents the EnumConnectionPoints operation response
type EnumConnectionPointsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat             `idl:"name:That" json:"that"`
	Enum *mqac.EnumConnectionPoints `idl:"name:ppEnum" json:"enum"`
	// Return: The EnumConnectionPoints return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumConnectionPointsResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumConnectionPointsOperation) *xxx_EnumConnectionPointsOperation {
	if op == nil {
		op = &xxx_EnumConnectionPointsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Enum = o.Enum
	op.Return = o.Return
	return op
}

func (o *EnumConnectionPointsResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumConnectionPointsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enum = op.Enum
	o.Return = op.Return
}
func (o *EnumConnectionPointsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumConnectionPointsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumConnectionPointsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FindConnectionPointOperation structure represents the FindConnectionPoint operation
type xxx_FindConnectionPointOperation struct {
	This            *dcom.ORPCThis        `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat        `idl:"name:That" json:"that"`
	IID             *dcom.IID             `idl:"name:riid" json:"iid"`
	ConnectionPoint *mqac.ConnectionPoint `idl:"name:ppCP" json:"connection_point"`
	Return          int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_FindConnectionPointOperation) OpNum() int { return 4 }

func (o *xxx_FindConnectionPointOperation) OpName() string {
	return "/IConnectionPointContainer/v0/FindConnectionPoint"
}

func (o *xxx_FindConnectionPointOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FindConnectionPointOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// riid {in} (1:{alias=REFIID}*(1))(2:{alias=IID, names=GUID}(struct))
	{
		if o.IID != nil {
			if err := o.IID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.IID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_FindConnectionPointOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// riid {in} (1:{alias=REFIID,pointer=ref}*(1))(2:{alias=IID, names=GUID}(struct))
	{
		if o.IID == nil {
			o.IID = &dcom.IID{}
		}
		if err := o.IID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FindConnectionPointOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FindConnectionPointOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppCP {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IConnectionPoint}(interface))
	{
		if o.ConnectionPoint != nil {
			_ptr_ppCP := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConnectionPoint != nil {
					if err := o.ConnectionPoint.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.ConnectionPoint{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConnectionPoint, _ptr_ppCP); err != nil {
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

func (o *xxx_FindConnectionPointOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppCP {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IConnectionPoint}(interface))
	{
		_ptr_ppCP := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConnectionPoint == nil {
				o.ConnectionPoint = &mqac.ConnectionPoint{}
			}
			if err := o.ConnectionPoint.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppCP := func(ptr interface{}) { o.ConnectionPoint = *ptr.(**mqac.ConnectionPoint) }
		if err := w.ReadPointer(&o.ConnectionPoint, _s_ppCP, _ptr_ppCP); err != nil {
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

// FindConnectionPointRequest structure represents the FindConnectionPoint operation request
type FindConnectionPointRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	IID  *dcom.IID      `idl:"name:riid" json:"iid"`
}

func (o *FindConnectionPointRequest) xxx_ToOp(ctx context.Context, op *xxx_FindConnectionPointOperation) *xxx_FindConnectionPointOperation {
	if op == nil {
		op = &xxx_FindConnectionPointOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.IID = o.IID
	return op
}

func (o *FindConnectionPointRequest) xxx_FromOp(ctx context.Context, op *xxx_FindConnectionPointOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.IID = op.IID
}
func (o *FindConnectionPointRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *FindConnectionPointRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FindConnectionPointOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FindConnectionPointResponse structure represents the FindConnectionPoint operation response
type FindConnectionPointResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That            *dcom.ORPCThat        `idl:"name:That" json:"that"`
	ConnectionPoint *mqac.ConnectionPoint `idl:"name:ppCP" json:"connection_point"`
	// Return: The FindConnectionPoint return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *FindConnectionPointResponse) xxx_ToOp(ctx context.Context, op *xxx_FindConnectionPointOperation) *xxx_FindConnectionPointOperation {
	if op == nil {
		op = &xxx_FindConnectionPointOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ConnectionPoint = o.ConnectionPoint
	op.Return = o.Return
	return op
}

func (o *FindConnectionPointResponse) xxx_FromOp(ctx context.Context, op *xxx_FindConnectionPointOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ConnectionPoint = op.ConnectionPoint
	o.Return = op.Return
}
func (o *FindConnectionPointResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *FindConnectionPointResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FindConnectionPointOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
