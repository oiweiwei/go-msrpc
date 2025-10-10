package imsmqapplication3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	imsmqapplication2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqapplication2/v0"
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
	_ = imsmqapplication2.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/mqac"
)

var (
	// IMSMQApplication3 interface identifier eba96b1f-2168-11d3-898c-00e02c074f6b
	ImsmqApplication3IID = &dcom.IID{Data1: 0xeba96b1f, Data2: 0x2168, Data3: 0x11d3, Data4: []byte{0x89, 0x8c, 0x00, 0xe0, 0x2c, 0x07, 0x4f, 0x6b}}
	// Syntax UUID
	ImsmqApplication3SyntaxUUID = &uuid.UUID{TimeLow: 0xeba96b1f, TimeMid: 0x2168, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x8c, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	// Syntax ID
	ImsmqApplication3SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ImsmqApplication3SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IMSMQApplication3 interface.
type ImsmqApplication3Client interface {

	// IMSMQApplication2 retrieval method.
	ImsmqApplication2() imsmqapplication2.ImsmqApplication2Client

	// ActiveQueues operation.
	ActiveQueues(context.Context, *ActiveQueuesRequest, ...dcerpc.CallOption) (*ActiveQueuesResponse, error)

	// PrivateQueues operation.
	GetPrivateQueues(context.Context, *GetPrivateQueuesRequest, ...dcerpc.CallOption) (*GetPrivateQueuesResponse, error)

	// DirectoryServiceServer operation.
	GetDirectoryServiceServer(context.Context, *GetDirectoryServiceServerRequest, ...dcerpc.CallOption) (*GetDirectoryServiceServerResponse, error)

	// IsConnected operation.
	GetIsConnected(context.Context, *GetIsConnectedRequest, ...dcerpc.CallOption) (*GetIsConnectedResponse, error)

	// BytesInAllQueues operation.
	GetBytesInAllQueues(context.Context, *GetBytesInAllQueuesRequest, ...dcerpc.CallOption) (*GetBytesInAllQueuesResponse, error)

	// Machine operation.
	SetMachine(context.Context, *SetMachineRequest, ...dcerpc.CallOption) (*SetMachineResponse, error)

	// Machine operation.
	GetMachine(context.Context, *GetMachineRequest, ...dcerpc.CallOption) (*GetMachineResponse, error)

	// Connect operation.
	Connect(context.Context, *ConnectRequest, ...dcerpc.CallOption) (*ConnectResponse, error)

	// Disconnect operation.
	Disconnect(context.Context, *DisconnectRequest, ...dcerpc.CallOption) (*DisconnectResponse, error)

	// Tidy operation.
	Tidy(context.Context, *TidyRequest, ...dcerpc.CallOption) (*TidyResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ImsmqApplication3Client
}

type xxx_DefaultImsmqApplication3Client struct {
	imsmqapplication2.ImsmqApplication2Client
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultImsmqApplication3Client) ImsmqApplication2() imsmqapplication2.ImsmqApplication2Client {
	return o.ImsmqApplication2Client
}

func (o *xxx_DefaultImsmqApplication3Client) ActiveQueues(ctx context.Context, in *ActiveQueuesRequest, opts ...dcerpc.CallOption) (*ActiveQueuesResponse, error) {
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
	out := &ActiveQueuesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqApplication3Client) GetPrivateQueues(ctx context.Context, in *GetPrivateQueuesRequest, opts ...dcerpc.CallOption) (*GetPrivateQueuesResponse, error) {
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
	out := &GetPrivateQueuesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqApplication3Client) GetDirectoryServiceServer(ctx context.Context, in *GetDirectoryServiceServerRequest, opts ...dcerpc.CallOption) (*GetDirectoryServiceServerResponse, error) {
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
	out := &GetDirectoryServiceServerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqApplication3Client) GetIsConnected(ctx context.Context, in *GetIsConnectedRequest, opts ...dcerpc.CallOption) (*GetIsConnectedResponse, error) {
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
	out := &GetIsConnectedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqApplication3Client) GetBytesInAllQueues(ctx context.Context, in *GetBytesInAllQueuesRequest, opts ...dcerpc.CallOption) (*GetBytesInAllQueuesResponse, error) {
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
	out := &GetBytesInAllQueuesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqApplication3Client) SetMachine(ctx context.Context, in *SetMachineRequest, opts ...dcerpc.CallOption) (*SetMachineResponse, error) {
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
	out := &SetMachineResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqApplication3Client) GetMachine(ctx context.Context, in *GetMachineRequest, opts ...dcerpc.CallOption) (*GetMachineResponse, error) {
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
	out := &GetMachineResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqApplication3Client) Connect(ctx context.Context, in *ConnectRequest, opts ...dcerpc.CallOption) (*ConnectResponse, error) {
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
	out := &ConnectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqApplication3Client) Disconnect(ctx context.Context, in *DisconnectRequest, opts ...dcerpc.CallOption) (*DisconnectResponse, error) {
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
	out := &DisconnectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqApplication3Client) Tidy(ctx context.Context, in *TidyRequest, opts ...dcerpc.CallOption) (*TidyResponse, error) {
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
	out := &TidyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqApplication3Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultImsmqApplication3Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultImsmqApplication3Client) IPID(ctx context.Context, ipid *dcom.IPID) ImsmqApplication3Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultImsmqApplication3Client{
		ImsmqApplication2Client: o.ImsmqApplication2Client.IPID(ctx, ipid),
		cc:                      o.cc,
		ipid:                    ipid,
	}
}

func NewImsmqApplication3Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ImsmqApplication3Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ImsmqApplication3SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := imsmqapplication2.NewImsmqApplication2Client(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultImsmqApplication3Client{
		ImsmqApplication2Client: base,
		cc:                      cc,
		ipid:                    ipid,
	}, nil
}

// xxx_ActiveQueuesOperation structure represents the ActiveQueues operation
type xxx_ActiveQueuesOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	ActiveQueues *oaut.Variant  `idl:"name:pvActiveQueues" json:"active_queues"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ActiveQueuesOperation) OpNum() int { return 15 }

func (o *xxx_ActiveQueuesOperation) OpName() string { return "/IMSMQApplication3/v0/ActiveQueues" }

func (o *xxx_ActiveQueuesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ActiveQueuesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ActiveQueuesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_ActiveQueuesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ActiveQueuesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pvActiveQueues {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.ActiveQueues != nil {
			_ptr_pvActiveQueues := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ActiveQueues != nil {
					if err := o.ActiveQueues.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ActiveQueues, _ptr_pvActiveQueues); err != nil {
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

func (o *xxx_ActiveQueuesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pvActiveQueues {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pvActiveQueues := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ActiveQueues == nil {
				o.ActiveQueues = &oaut.Variant{}
			}
			if err := o.ActiveQueues.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pvActiveQueues := func(ptr interface{}) { o.ActiveQueues = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.ActiveQueues, _s_pvActiveQueues, _ptr_pvActiveQueues); err != nil {
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

// ActiveQueuesRequest structure represents the ActiveQueues operation request
type ActiveQueuesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *ActiveQueuesRequest) xxx_ToOp(ctx context.Context, op *xxx_ActiveQueuesOperation) *xxx_ActiveQueuesOperation {
	if op == nil {
		op = &xxx_ActiveQueuesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *ActiveQueuesRequest) xxx_FromOp(ctx context.Context, op *xxx_ActiveQueuesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *ActiveQueuesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ActiveQueuesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ActiveQueuesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ActiveQueuesResponse structure represents the ActiveQueues operation response
type ActiveQueuesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	ActiveQueues *oaut.Variant  `idl:"name:pvActiveQueues" json:"active_queues"`
	// Return: The ActiveQueues return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ActiveQueuesResponse) xxx_ToOp(ctx context.Context, op *xxx_ActiveQueuesOperation) *xxx_ActiveQueuesOperation {
	if op == nil {
		op = &xxx_ActiveQueuesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ActiveQueues = o.ActiveQueues
	op.Return = o.Return
	return op
}

func (o *ActiveQueuesResponse) xxx_FromOp(ctx context.Context, op *xxx_ActiveQueuesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ActiveQueues = op.ActiveQueues
	o.Return = op.Return
}
func (o *ActiveQueuesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ActiveQueuesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ActiveQueuesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPrivateQueuesOperation structure represents the PrivateQueues operation
type xxx_GetPrivateQueuesOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	PrivateQueues *oaut.Variant  `idl:"name:pvPrivateQueues" json:"private_queues"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPrivateQueuesOperation) OpNum() int { return 16 }

func (o *xxx_GetPrivateQueuesOperation) OpName() string { return "/IMSMQApplication3/v0/PrivateQueues" }

func (o *xxx_GetPrivateQueuesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPrivateQueuesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPrivateQueuesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPrivateQueuesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPrivateQueuesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pvPrivateQueues {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.PrivateQueues != nil {
			_ptr_pvPrivateQueues := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PrivateQueues != nil {
					if err := o.PrivateQueues.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PrivateQueues, _ptr_pvPrivateQueues); err != nil {
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

func (o *xxx_GetPrivateQueuesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pvPrivateQueues {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pvPrivateQueues := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PrivateQueues == nil {
				o.PrivateQueues = &oaut.Variant{}
			}
			if err := o.PrivateQueues.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pvPrivateQueues := func(ptr interface{}) { o.PrivateQueues = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.PrivateQueues, _s_pvPrivateQueues, _ptr_pvPrivateQueues); err != nil {
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

// GetPrivateQueuesRequest structure represents the PrivateQueues operation request
type GetPrivateQueuesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPrivateQueuesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPrivateQueuesOperation) *xxx_GetPrivateQueuesOperation {
	if op == nil {
		op = &xxx_GetPrivateQueuesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetPrivateQueuesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPrivateQueuesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPrivateQueuesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPrivateQueuesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPrivateQueuesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPrivateQueuesResponse structure represents the PrivateQueues operation response
type GetPrivateQueuesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	PrivateQueues *oaut.Variant  `idl:"name:pvPrivateQueues" json:"private_queues"`
	// Return: The PrivateQueues return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPrivateQueuesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPrivateQueuesOperation) *xxx_GetPrivateQueuesOperation {
	if op == nil {
		op = &xxx_GetPrivateQueuesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PrivateQueues = o.PrivateQueues
	op.Return = o.Return
	return op
}

func (o *GetPrivateQueuesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPrivateQueuesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PrivateQueues = op.PrivateQueues
	o.Return = op.Return
}
func (o *GetPrivateQueuesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPrivateQueuesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPrivateQueuesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDirectoryServiceServerOperation structure represents the DirectoryServiceServer operation
type xxx_GetDirectoryServiceServerOperation struct {
	This                   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                   *dcom.ORPCThat `idl:"name:That" json:"that"`
	DirectoryServiceServer *oaut.String   `idl:"name:pbstrDirectoryServiceServer" json:"directory_service_server"`
	Return                 int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDirectoryServiceServerOperation) OpNum() int { return 17 }

func (o *xxx_GetDirectoryServiceServerOperation) OpName() string {
	return "/IMSMQApplication3/v0/DirectoryServiceServer"
}

func (o *xxx_GetDirectoryServiceServerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDirectoryServiceServerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetDirectoryServiceServerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetDirectoryServiceServerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDirectoryServiceServerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrDirectoryServiceServer {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.DirectoryServiceServer != nil {
			_ptr_pbstrDirectoryServiceServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DirectoryServiceServer != nil {
					if err := o.DirectoryServiceServer.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DirectoryServiceServer, _ptr_pbstrDirectoryServiceServer); err != nil {
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

func (o *xxx_GetDirectoryServiceServerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrDirectoryServiceServer {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrDirectoryServiceServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DirectoryServiceServer == nil {
				o.DirectoryServiceServer = &oaut.String{}
			}
			if err := o.DirectoryServiceServer.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrDirectoryServiceServer := func(ptr interface{}) { o.DirectoryServiceServer = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.DirectoryServiceServer, _s_pbstrDirectoryServiceServer, _ptr_pbstrDirectoryServiceServer); err != nil {
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

// GetDirectoryServiceServerRequest structure represents the DirectoryServiceServer operation request
type GetDirectoryServiceServerRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetDirectoryServiceServerRequest) xxx_ToOp(ctx context.Context, op *xxx_GetDirectoryServiceServerOperation) *xxx_GetDirectoryServiceServerOperation {
	if op == nil {
		op = &xxx_GetDirectoryServiceServerOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetDirectoryServiceServerRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDirectoryServiceServerOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetDirectoryServiceServerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetDirectoryServiceServerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDirectoryServiceServerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDirectoryServiceServerResponse structure represents the DirectoryServiceServer operation response
type GetDirectoryServiceServerResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That                   *dcom.ORPCThat `idl:"name:That" json:"that"`
	DirectoryServiceServer *oaut.String   `idl:"name:pbstrDirectoryServiceServer" json:"directory_service_server"`
	// Return: The DirectoryServiceServer return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDirectoryServiceServerResponse) xxx_ToOp(ctx context.Context, op *xxx_GetDirectoryServiceServerOperation) *xxx_GetDirectoryServiceServerOperation {
	if op == nil {
		op = &xxx_GetDirectoryServiceServerOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.DirectoryServiceServer = o.DirectoryServiceServer
	op.Return = o.Return
	return op
}

func (o *GetDirectoryServiceServerResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDirectoryServiceServerOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DirectoryServiceServer = op.DirectoryServiceServer
	o.Return = op.Return
}
func (o *GetDirectoryServiceServerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetDirectoryServiceServerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDirectoryServiceServerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetIsConnectedOperation structure represents the IsConnected operation
type xxx_GetIsConnectedOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	IsConnected int16          `idl:"name:pfIsConnected" json:"is_connected"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIsConnectedOperation) OpNum() int { return 18 }

func (o *xxx_GetIsConnectedOperation) OpName() string { return "/IMSMQApplication3/v0/IsConnected" }

func (o *xxx_GetIsConnectedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsConnectedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetIsConnectedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetIsConnectedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsConnectedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pfIsConnected {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.IsConnected); err != nil {
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

func (o *xxx_GetIsConnectedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pfIsConnected {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.IsConnected); err != nil {
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

// GetIsConnectedRequest structure represents the IsConnected operation request
type GetIsConnectedRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetIsConnectedRequest) xxx_ToOp(ctx context.Context, op *xxx_GetIsConnectedOperation) *xxx_GetIsConnectedOperation {
	if op == nil {
		op = &xxx_GetIsConnectedOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetIsConnectedRequest) xxx_FromOp(ctx context.Context, op *xxx_GetIsConnectedOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetIsConnectedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetIsConnectedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsConnectedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIsConnectedResponse structure represents the IsConnected operation response
type GetIsConnectedResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	IsConnected int16          `idl:"name:pfIsConnected" json:"is_connected"`
	// Return: The IsConnected return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIsConnectedResponse) xxx_ToOp(ctx context.Context, op *xxx_GetIsConnectedOperation) *xxx_GetIsConnectedOperation {
	if op == nil {
		op = &xxx_GetIsConnectedOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.IsConnected = o.IsConnected
	op.Return = o.Return
	return op
}

func (o *GetIsConnectedResponse) xxx_FromOp(ctx context.Context, op *xxx_GetIsConnectedOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.IsConnected = op.IsConnected
	o.Return = op.Return
}
func (o *GetIsConnectedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetIsConnectedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsConnectedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetBytesInAllQueuesOperation structure represents the BytesInAllQueues operation
type xxx_GetBytesInAllQueuesOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	BytesInAllQueues *oaut.Variant  `idl:"name:pvBytesInAllQueues" json:"bytes_in_all_queues"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetBytesInAllQueuesOperation) OpNum() int { return 19 }

func (o *xxx_GetBytesInAllQueuesOperation) OpName() string {
	return "/IMSMQApplication3/v0/BytesInAllQueues"
}

func (o *xxx_GetBytesInAllQueuesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBytesInAllQueuesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetBytesInAllQueuesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetBytesInAllQueuesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBytesInAllQueuesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pvBytesInAllQueues {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.BytesInAllQueues != nil {
			_ptr_pvBytesInAllQueues := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.BytesInAllQueues != nil {
					if err := o.BytesInAllQueues.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.BytesInAllQueues, _ptr_pvBytesInAllQueues); err != nil {
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

func (o *xxx_GetBytesInAllQueuesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pvBytesInAllQueues {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pvBytesInAllQueues := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.BytesInAllQueues == nil {
				o.BytesInAllQueues = &oaut.Variant{}
			}
			if err := o.BytesInAllQueues.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pvBytesInAllQueues := func(ptr interface{}) { o.BytesInAllQueues = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.BytesInAllQueues, _s_pvBytesInAllQueues, _ptr_pvBytesInAllQueues); err != nil {
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

// GetBytesInAllQueuesRequest structure represents the BytesInAllQueues operation request
type GetBytesInAllQueuesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetBytesInAllQueuesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetBytesInAllQueuesOperation) *xxx_GetBytesInAllQueuesOperation {
	if op == nil {
		op = &xxx_GetBytesInAllQueuesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetBytesInAllQueuesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetBytesInAllQueuesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetBytesInAllQueuesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetBytesInAllQueuesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetBytesInAllQueuesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetBytesInAllQueuesResponse structure represents the BytesInAllQueues operation response
type GetBytesInAllQueuesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	BytesInAllQueues *oaut.Variant  `idl:"name:pvBytesInAllQueues" json:"bytes_in_all_queues"`
	// Return: The BytesInAllQueues return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetBytesInAllQueuesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetBytesInAllQueuesOperation) *xxx_GetBytesInAllQueuesOperation {
	if op == nil {
		op = &xxx_GetBytesInAllQueuesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.BytesInAllQueues = o.BytesInAllQueues
	op.Return = o.Return
	return op
}

func (o *GetBytesInAllQueuesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetBytesInAllQueuesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.BytesInAllQueues = op.BytesInAllQueues
	o.Return = op.Return
}
func (o *GetBytesInAllQueuesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetBytesInAllQueuesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetBytesInAllQueuesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetMachineOperation structure represents the Machine operation
type xxx_SetMachineOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Machine *oaut.String   `idl:"name:bstrMachine" json:"machine"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetMachineOperation) OpNum() int { return 20 }

func (o *xxx_SetMachineOperation) OpName() string { return "/IMSMQApplication3/v0/Machine" }

func (o *xxx_SetMachineOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMachineOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrMachine {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Machine != nil {
			_ptr_bstrMachine := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Machine != nil {
					if err := o.Machine.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Machine, _ptr_bstrMachine); err != nil {
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

func (o *xxx_SetMachineOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrMachine {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrMachine := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Machine == nil {
				o.Machine = &oaut.String{}
			}
			if err := o.Machine.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrMachine := func(ptr interface{}) { o.Machine = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Machine, _s_bstrMachine, _ptr_bstrMachine); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMachineOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMachineOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetMachineOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetMachineRequest structure represents the Machine operation request
type SetMachineRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Machine *oaut.String   `idl:"name:bstrMachine" json:"machine"`
}

func (o *SetMachineRequest) xxx_ToOp(ctx context.Context, op *xxx_SetMachineOperation) *xxx_SetMachineOperation {
	if op == nil {
		op = &xxx_SetMachineOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Machine = o.Machine
	return op
}

func (o *SetMachineRequest) xxx_FromOp(ctx context.Context, op *xxx_SetMachineOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Machine = op.Machine
}
func (o *SetMachineRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetMachineRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMachineOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetMachineResponse structure represents the Machine operation response
type SetMachineResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Machine return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetMachineResponse) xxx_ToOp(ctx context.Context, op *xxx_SetMachineOperation) *xxx_SetMachineOperation {
	if op == nil {
		op = &xxx_SetMachineOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetMachineResponse) xxx_FromOp(ctx context.Context, op *xxx_SetMachineOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetMachineResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetMachineResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMachineOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMachineOperation structure represents the Machine operation
type xxx_GetMachineOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Machine *oaut.String   `idl:"name:pbstrMachine" json:"machine"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMachineOperation) OpNum() int { return 21 }

func (o *xxx_GetMachineOperation) OpName() string { return "/IMSMQApplication3/v0/Machine" }

func (o *xxx_GetMachineOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMachineOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetMachineOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetMachineOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMachineOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrMachine {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Machine != nil {
			_ptr_pbstrMachine := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Machine != nil {
					if err := o.Machine.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Machine, _ptr_pbstrMachine); err != nil {
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

func (o *xxx_GetMachineOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrMachine {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrMachine := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Machine == nil {
				o.Machine = &oaut.String{}
			}
			if err := o.Machine.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrMachine := func(ptr interface{}) { o.Machine = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Machine, _s_pbstrMachine, _ptr_pbstrMachine); err != nil {
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

// GetMachineRequest structure represents the Machine operation request
type GetMachineRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMachineRequest) xxx_ToOp(ctx context.Context, op *xxx_GetMachineOperation) *xxx_GetMachineOperation {
	if op == nil {
		op = &xxx_GetMachineOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetMachineRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMachineOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMachineRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetMachineRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMachineOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMachineResponse structure represents the Machine operation response
type GetMachineResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Machine *oaut.String   `idl:"name:pbstrMachine" json:"machine"`
	// Return: The Machine return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMachineResponse) xxx_ToOp(ctx context.Context, op *xxx_GetMachineOperation) *xxx_GetMachineOperation {
	if op == nil {
		op = &xxx_GetMachineOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Machine = o.Machine
	op.Return = o.Return
	return op
}

func (o *GetMachineResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMachineOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Machine = op.Machine
	o.Return = op.Return
}
func (o *GetMachineResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetMachineResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMachineOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ConnectOperation structure represents the Connect operation
type xxx_ConnectOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ConnectOperation) OpNum() int { return 22 }

func (o *xxx_ConnectOperation) OpName() string { return "/IMSMQApplication3/v0/Connect" }

func (o *xxx_ConnectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ConnectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_ConnectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ConnectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ConnectRequest structure represents the Connect operation request
type ConnectRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *ConnectRequest) xxx_ToOp(ctx context.Context, op *xxx_ConnectOperation) *xxx_ConnectOperation {
	if op == nil {
		op = &xxx_ConnectOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *ConnectRequest) xxx_FromOp(ctx context.Context, op *xxx_ConnectOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *ConnectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ConnectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ConnectResponse structure represents the Connect operation response
type ConnectResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Connect return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ConnectResponse) xxx_ToOp(ctx context.Context, op *xxx_ConnectOperation) *xxx_ConnectOperation {
	if op == nil {
		op = &xxx_ConnectOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ConnectResponse) xxx_FromOp(ctx context.Context, op *xxx_ConnectOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ConnectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ConnectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DisconnectOperation structure represents the Disconnect operation
type xxx_DisconnectOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DisconnectOperation) OpNum() int { return 23 }

func (o *xxx_DisconnectOperation) OpName() string { return "/IMSMQApplication3/v0/Disconnect" }

func (o *xxx_DisconnectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DisconnectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DisconnectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_DisconnectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DisconnectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DisconnectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DisconnectRequest structure represents the Disconnect operation request
type DisconnectRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *DisconnectRequest) xxx_ToOp(ctx context.Context, op *xxx_DisconnectOperation) *xxx_DisconnectOperation {
	if op == nil {
		op = &xxx_DisconnectOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *DisconnectRequest) xxx_FromOp(ctx context.Context, op *xxx_DisconnectOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *DisconnectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DisconnectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DisconnectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DisconnectResponse structure represents the Disconnect operation response
type DisconnectResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Disconnect return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DisconnectResponse) xxx_ToOp(ctx context.Context, op *xxx_DisconnectOperation) *xxx_DisconnectOperation {
	if op == nil {
		op = &xxx_DisconnectOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DisconnectResponse) xxx_FromOp(ctx context.Context, op *xxx_DisconnectOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DisconnectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DisconnectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DisconnectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_TidyOperation structure represents the Tidy operation
type xxx_TidyOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_TidyOperation) OpNum() int { return 24 }

func (o *xxx_TidyOperation) OpName() string { return "/IMSMQApplication3/v0/Tidy" }

func (o *xxx_TidyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TidyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_TidyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_TidyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TidyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_TidyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// TidyRequest structure represents the Tidy operation request
type TidyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *TidyRequest) xxx_ToOp(ctx context.Context, op *xxx_TidyOperation) *xxx_TidyOperation {
	if op == nil {
		op = &xxx_TidyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *TidyRequest) xxx_FromOp(ctx context.Context, op *xxx_TidyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *TidyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *TidyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_TidyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// TidyResponse structure represents the Tidy operation response
type TidyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Tidy return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *TidyResponse) xxx_ToOp(ctx context.Context, op *xxx_TidyOperation) *xxx_TidyOperation {
	if op == nil {
		op = &xxx_TidyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *TidyResponse) xxx_FromOp(ctx context.Context, op *xxx_TidyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *TidyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *TidyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_TidyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
