package iconnectionpoint

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
	// IConnectionPoint interface identifier b196b286-bab4-101a-b69c-00aa00341d07
	ConnectionPointIID = &dcom.IID{Data1: 0xb196b286, Data2: 0xbab4, Data3: 0x101a, Data4: []byte{0xb6, 0x9c, 0x00, 0xaa, 0x00, 0x34, 0x1d, 0x07}}
	// Syntax UUID
	ConnectionPointSyntaxUUID = &uuid.UUID{TimeLow: 0xb196b286, TimeMid: 0xbab4, TimeHiAndVersion: 0x101a, ClockSeqHiAndReserved: 0xb6, ClockSeqLow: 0x9c, Node: [6]uint8{0x0, 0xaa, 0x0, 0x34, 0x1d, 0x7}}
	// Syntax ID
	ConnectionPointSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ConnectionPointSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IConnectionPoint interface.
type ConnectionPointClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The GetConnectionInterface method is received by the server in an RPC_REQUEST packet.
	// In response, the server MUST return a pointer to the IID of the interface that client
	// applications MUST implement for the MSMQEvent object to be able to notify them when
	// an event is raised.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.<92>
	GetConnectionInterface(context.Context, *GetConnectionInterfaceRequest, ...dcerpc.CallOption) (*GetConnectionInterfaceResponse, error)

	// The GetConnectionPointContainer method is received by the server in an RPC_REQUEST
	// packet. In response, the server MUST return a pointer to a pointer to an IConnectionPointContainer
	// interface for the MSMQEvent object.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.<93>
	GetConnectionPointContainer(context.Context, *GetConnectionPointContainerRequest, ...dcerpc.CallOption) (*GetConnectionPointContainerResponse, error)

	// The Advise method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST register a client's callback object to receive event notifications.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.<94>
	Advise(context.Context, *AdviseRequest, ...dcerpc.CallOption) (*AdviseResponse, error)

	// The Unadvise method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST deregister a client's callback object to cease receiving event notifications.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.<95>
	Unadvise(context.Context, *UnadviseRequest, ...dcerpc.CallOption) (*UnadviseResponse, error)

	// The EnumConnections method is received by the server in an RPC_REQUEST packet. In
	// response, the server MUST return a pointer to an IEnumConnections interface pointer,
	// as defined in [MSDN-EC], for the client to enumerate all the currently registered
	// callback objects.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.<96>
	EnumConnections(context.Context, *EnumConnectionsRequest, ...dcerpc.CallOption) (*EnumConnectionsResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ConnectionPointClient
}

type xxx_DefaultConnectionPointClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultConnectionPointClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultConnectionPointClient) GetConnectionInterface(ctx context.Context, in *GetConnectionInterfaceRequest, opts ...dcerpc.CallOption) (*GetConnectionInterfaceResponse, error) {
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
	out := &GetConnectionInterfaceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConnectionPointClient) GetConnectionPointContainer(ctx context.Context, in *GetConnectionPointContainerRequest, opts ...dcerpc.CallOption) (*GetConnectionPointContainerResponse, error) {
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
	out := &GetConnectionPointContainerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConnectionPointClient) Advise(ctx context.Context, in *AdviseRequest, opts ...dcerpc.CallOption) (*AdviseResponse, error) {
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
	out := &AdviseResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConnectionPointClient) Unadvise(ctx context.Context, in *UnadviseRequest, opts ...dcerpc.CallOption) (*UnadviseResponse, error) {
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
	out := &UnadviseResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConnectionPointClient) EnumConnections(ctx context.Context, in *EnumConnectionsRequest, opts ...dcerpc.CallOption) (*EnumConnectionsResponse, error) {
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
	out := &EnumConnectionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultConnectionPointClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultConnectionPointClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultConnectionPointClient) IPID(ctx context.Context, ipid *dcom.IPID) ConnectionPointClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultConnectionPointClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewConnectionPointClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ConnectionPointClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ConnectionPointSyntaxV0_0))...)
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
	return &xxx_DefaultConnectionPointClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetConnectionInterfaceOperation structure represents the GetConnectionInterface operation
type xxx_GetConnectionInterfaceOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	IID    *dcom.IID      `idl:"name:pIID" json:"iid"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetConnectionInterfaceOperation) OpNum() int { return 3 }

func (o *xxx_GetConnectionInterfaceOperation) OpName() string {
	return "/IConnectionPoint/v0/GetConnectionInterface"
}

func (o *xxx_GetConnectionInterfaceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConnectionInterfaceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetConnectionInterfaceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetConnectionInterfaceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConnectionInterfaceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pIID {out} (1:{pointer=ref}*(1))(2:{alias=IID, names=GUID}(struct))
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConnectionInterfaceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pIID {out} (1:{pointer=ref}*(1))(2:{alias=IID, names=GUID}(struct))
	{
		if o.IID == nil {
			o.IID = &dcom.IID{}
		}
		if err := o.IID.UnmarshalNDR(ctx, w); err != nil {
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

// GetConnectionInterfaceRequest structure represents the GetConnectionInterface operation request
type GetConnectionInterfaceRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetConnectionInterfaceRequest) xxx_ToOp(ctx context.Context, op *xxx_GetConnectionInterfaceOperation) *xxx_GetConnectionInterfaceOperation {
	if op == nil {
		op = &xxx_GetConnectionInterfaceOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetConnectionInterfaceRequest) xxx_FromOp(ctx context.Context, op *xxx_GetConnectionInterfaceOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetConnectionInterfaceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetConnectionInterfaceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetConnectionInterfaceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetConnectionInterfaceResponse structure represents the GetConnectionInterface operation response
type GetConnectionInterfaceResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pIID: A pointer to an IID that upon successful completion will contain the IID value
	// of the _DMSMQEventEvents interface.
	IID *dcom.IID `idl:"name:pIID" json:"iid"`
	// Return: The GetConnectionInterface return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetConnectionInterfaceResponse) xxx_ToOp(ctx context.Context, op *xxx_GetConnectionInterfaceOperation) *xxx_GetConnectionInterfaceOperation {
	if op == nil {
		op = &xxx_GetConnectionInterfaceOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.IID = o.IID
	op.Return = o.Return
	return op
}

func (o *GetConnectionInterfaceResponse) xxx_FromOp(ctx context.Context, op *xxx_GetConnectionInterfaceOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.IID = op.IID
	o.Return = op.Return
}
func (o *GetConnectionInterfaceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetConnectionInterfaceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetConnectionInterfaceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetConnectionPointContainerOperation structure represents the GetConnectionPointContainer operation
type xxx_GetConnectionPointContainerOperation struct {
	This   *dcom.ORPCThis                 `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat                 `idl:"name:That" json:"that"`
	CPC    *mqac.ConnectionPointContainer `idl:"name:ppCPC" json:"cpc"`
	Return int32                          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetConnectionPointContainerOperation) OpNum() int { return 4 }

func (o *xxx_GetConnectionPointContainerOperation) OpName() string {
	return "/IConnectionPoint/v0/GetConnectionPointContainer"
}

func (o *xxx_GetConnectionPointContainerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConnectionPointContainerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetConnectionPointContainerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetConnectionPointContainerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConnectionPointContainerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppCPC {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IConnectionPointContainer}(interface))
	{
		if o.CPC != nil {
			_ptr_ppCPC := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CPC != nil {
					if err := o.CPC.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.ConnectionPointContainer{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CPC, _ptr_ppCPC); err != nil {
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

func (o *xxx_GetConnectionPointContainerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppCPC {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IConnectionPointContainer}(interface))
	{
		_ptr_ppCPC := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CPC == nil {
				o.CPC = &mqac.ConnectionPointContainer{}
			}
			if err := o.CPC.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppCPC := func(ptr interface{}) { o.CPC = *ptr.(**mqac.ConnectionPointContainer) }
		if err := w.ReadPointer(&o.CPC, _s_ppCPC, _ptr_ppCPC); err != nil {
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

// GetConnectionPointContainerRequest structure represents the GetConnectionPointContainer operation request
type GetConnectionPointContainerRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetConnectionPointContainerRequest) xxx_ToOp(ctx context.Context, op *xxx_GetConnectionPointContainerOperation) *xxx_GetConnectionPointContainerOperation {
	if op == nil {
		op = &xxx_GetConnectionPointContainerOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetConnectionPointContainerRequest) xxx_FromOp(ctx context.Context, op *xxx_GetConnectionPointContainerOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetConnectionPointContainerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetConnectionPointContainerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetConnectionPointContainerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetConnectionPointContainerResponse structure represents the GetConnectionPointContainer operation response
type GetConnectionPointContainerResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppCPC: A pointer to an IConnectionPointContainer interface pointer for the MSMQEvent
	// object.
	CPC *mqac.ConnectionPointContainer `idl:"name:ppCPC" json:"cpc"`
	// Return: The GetConnectionPointContainer return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetConnectionPointContainerResponse) xxx_ToOp(ctx context.Context, op *xxx_GetConnectionPointContainerOperation) *xxx_GetConnectionPointContainerOperation {
	if op == nil {
		op = &xxx_GetConnectionPointContainerOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.CPC = o.CPC
	op.Return = o.Return
	return op
}

func (o *GetConnectionPointContainerResponse) xxx_FromOp(ctx context.Context, op *xxx_GetConnectionPointContainerOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.CPC = op.CPC
	o.Return = op.Return
}
func (o *GetConnectionPointContainerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetConnectionPointContainerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetConnectionPointContainerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AdviseOperation structure represents the Advise operation
type xxx_AdviseOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Sink   *dcom.Unknown  `idl:"name:pUnkSink" json:"sink"`
	Cookie uint32         `idl:"name:pdwCookie" json:"cookie"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_AdviseOperation) OpNum() int { return 5 }

func (o *xxx_AdviseOperation) OpName() string { return "/IConnectionPoint/v0/Advise" }

func (o *xxx_AdviseOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AdviseOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pUnkSink {in} (1:{pointer=ref}*(1))(2:{alias=IUnknown}(interface))
	{
		if o.Sink != nil {
			_ptr_pUnkSink := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Sink != nil {
					if err := o.Sink.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.Unknown{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Sink, _ptr_pUnkSink); err != nil {
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

func (o *xxx_AdviseOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pUnkSink {in} (1:{pointer=ref}*(1))(2:{alias=IUnknown}(interface))
	{
		_ptr_pUnkSink := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Sink == nil {
				o.Sink = &dcom.Unknown{}
			}
			if err := o.Sink.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pUnkSink := func(ptr interface{}) { o.Sink = *ptr.(**dcom.Unknown) }
		if err := w.ReadPointer(&o.Sink, _s_pUnkSink, _ptr_pUnkSink); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AdviseOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AdviseOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pdwCookie {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Cookie); err != nil {
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

func (o *xxx_AdviseOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pdwCookie {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Cookie); err != nil {
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

// AdviseRequest structure represents the Advise operation request
type AdviseRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pUnkSink: A pointer to an IUnknown (as specified in [MS-DCOM] section 3.1.1.5.8)
	// interface for the client object.
	Sink *dcom.Unknown `idl:"name:pUnkSink" json:"sink"`
}

func (o *AdviseRequest) xxx_ToOp(ctx context.Context, op *xxx_AdviseOperation) *xxx_AdviseOperation {
	if op == nil {
		op = &xxx_AdviseOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Sink = o.Sink
	return op
}

func (o *AdviseRequest) xxx_FromOp(ctx context.Context, op *xxx_AdviseOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Sink = op.Sink
}
func (o *AdviseRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AdviseRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AdviseOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AdviseResponse structure represents the Advise operation response
type AdviseResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pdwCookie: A pointer to a unique DWORD value that uniquely identifies the new element
	// in the ConnectionCollection instance variable.
	Cookie uint32 `idl:"name:pdwCookie" json:"cookie"`
	// Return: The Advise return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AdviseResponse) xxx_ToOp(ctx context.Context, op *xxx_AdviseOperation) *xxx_AdviseOperation {
	if op == nil {
		op = &xxx_AdviseOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Cookie = o.Cookie
	op.Return = o.Return
	return op
}

func (o *AdviseResponse) xxx_FromOp(ctx context.Context, op *xxx_AdviseOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Cookie = op.Cookie
	o.Return = op.Return
}
func (o *AdviseResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AdviseResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AdviseOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_UnadviseOperation structure represents the Unadvise operation
type xxx_UnadviseOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Cookie uint32         `idl:"name:dwCookie" json:"cookie"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_UnadviseOperation) OpNum() int { return 6 }

func (o *xxx_UnadviseOperation) OpName() string { return "/IConnectionPoint/v0/Unadvise" }

func (o *xxx_UnadviseOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnadviseOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// dwCookie {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Cookie); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnadviseOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// dwCookie {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Cookie); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnadviseOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnadviseOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_UnadviseOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// UnadviseRequest structure represents the Unadvise operation request
type UnadviseRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// dwCookie: A DWORD value that uniquely identifies the interface pointer for the callback
	// that is to be deregistered. This corresponds to the value of the "cookie" that was
	// returned in the call to IConnectionPoint::Advise.
	Cookie uint32 `idl:"name:dwCookie" json:"cookie"`
}

func (o *UnadviseRequest) xxx_ToOp(ctx context.Context, op *xxx_UnadviseOperation) *xxx_UnadviseOperation {
	if op == nil {
		op = &xxx_UnadviseOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Cookie = o.Cookie
	return op
}

func (o *UnadviseRequest) xxx_FromOp(ctx context.Context, op *xxx_UnadviseOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Cookie = op.Cookie
}
func (o *UnadviseRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *UnadviseRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UnadviseOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// UnadviseResponse structure represents the Unadvise operation response
type UnadviseResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Unadvise return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *UnadviseResponse) xxx_ToOp(ctx context.Context, op *xxx_UnadviseOperation) *xxx_UnadviseOperation {
	if op == nil {
		op = &xxx_UnadviseOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *UnadviseResponse) xxx_FromOp(ctx context.Context, op *xxx_UnadviseOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *UnadviseResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *UnadviseResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UnadviseOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumConnectionsOperation structure represents the EnumConnections operation
type xxx_EnumConnectionsOperation struct {
	This   *dcom.ORPCThis        `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat        `idl:"name:That" json:"that"`
	Enum   *mqac.EnumConnections `idl:"name:ppEnum" json:"enum"`
	Return int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumConnectionsOperation) OpNum() int { return 7 }

func (o *xxx_EnumConnectionsOperation) OpName() string { return "/IConnectionPoint/v0/EnumConnections" }

func (o *xxx_EnumConnectionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumConnectionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_EnumConnectionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_EnumConnectionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumConnectionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumConnections}(interface))
	{
		if o.Enum != nil {
			_ptr_ppEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Enum != nil {
					if err := o.Enum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.EnumConnections{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_EnumConnectionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumConnections}(interface))
	{
		_ptr_ppEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Enum == nil {
				o.Enum = &mqac.EnumConnections{}
			}
			if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppEnum := func(ptr interface{}) { o.Enum = *ptr.(**mqac.EnumConnections) }
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

// EnumConnectionsRequest structure represents the EnumConnections operation request
type EnumConnectionsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *EnumConnectionsRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumConnectionsOperation) *xxx_EnumConnectionsOperation {
	if op == nil {
		op = &xxx_EnumConnectionsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *EnumConnectionsRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumConnectionsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *EnumConnectionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumConnectionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumConnectionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumConnectionsResponse structure represents the EnumConnections operation response
type EnumConnectionsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppEnum: A pointer to an IEnumConnections interface pointer that upon successful completion
	// will allow the user to enumerate all the currently registered callback objects.
	Enum *mqac.EnumConnections `idl:"name:ppEnum" json:"enum"`
	// Return: The EnumConnections return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumConnectionsResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumConnectionsOperation) *xxx_EnumConnectionsOperation {
	if op == nil {
		op = &xxx_EnumConnectionsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Enum = o.Enum
	op.Return = o.Return
	return op
}

func (o *EnumConnectionsResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumConnectionsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enum = op.Enum
	o.Return = op.Return
}
func (o *EnumConnectionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumConnectionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumConnectionsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
