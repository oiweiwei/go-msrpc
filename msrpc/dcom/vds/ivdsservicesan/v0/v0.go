package ivdsservicesan

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
	vds "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds"
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
	_ = vds.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/vds"
)

var (
	// IVdsServiceSAN interface identifier fc5d23e8-a88b-41a5-8de0-2d2f73c5a630
	ServiceSANIID = &dcom.IID{Data1: 0xfc5d23e8, Data2: 0xa88b, Data3: 0x41a5, Data4: []byte{0x8d, 0xe0, 0x2d, 0x2f, 0x73, 0xc5, 0xa6, 0x30}}
	// Syntax UUID
	ServiceSANSyntaxUUID = &uuid.UUID{TimeLow: 0xfc5d23e8, TimeMid: 0xa88b, TimeHiAndVersion: 0x41a5, ClockSeqHiAndReserved: 0x8d, ClockSeqLow: 0xe0, Node: [6]uint8{0x2d, 0x2f, 0x73, 0xc5, 0xa6, 0x30}}
	// Syntax ID
	ServiceSANSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ServiceSANSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsServiceSAN interface.
type ServiceSANClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The GetSANPolicy method returns the current SAN policy setting.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT, as specified in
	// [MS-ERREF], to indicate success or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetSANPolicy(context.Context, *GetSANPolicyRequest, ...dcerpc.CallOption) (*GetSANPolicyResponse, error)

	// The SetSANPolicy method sets the SAN policy value.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT, as specified in
	// [MS-ERREF], to indicate success or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	SetSANPolicy(context.Context, *SetSANPolicyRequest, ...dcerpc.CallOption) (*SetSANPolicyResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ServiceSANClient
}

type xxx_DefaultServiceSANClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultServiceSANClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultServiceSANClient) GetSANPolicy(ctx context.Context, in *GetSANPolicyRequest, opts ...dcerpc.CallOption) (*GetSANPolicyResponse, error) {
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
	out := &GetSANPolicyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServiceSANClient) SetSANPolicy(ctx context.Context, in *SetSANPolicyRequest, opts ...dcerpc.CallOption) (*SetSANPolicyResponse, error) {
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
	out := &SetSANPolicyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServiceSANClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultServiceSANClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultServiceSANClient) IPID(ctx context.Context, ipid *dcom.IPID) ServiceSANClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultServiceSANClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewServiceSANClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ServiceSANClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ServiceSANSyntaxV0_0))...)
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
	return &xxx_DefaultServiceSANClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetSANPolicyOperation structure represents the GetSANPolicy operation
type xxx_GetSANPolicyOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	SANPolicy vds.SANPolicy  `idl:"name:pSanPolicy" json:"san_policy"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSANPolicyOperation) OpNum() int { return 3 }

func (o *xxx_GetSANPolicyOperation) OpName() string { return "/IVdsServiceSAN/v0/GetSANPolicy" }

func (o *xxx_GetSANPolicyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSANPolicyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSANPolicyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSANPolicyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSANPolicyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pSanPolicy {out} (1:{pointer=ref}*(1))(2:{alias=VDS_SAN_POLICY}(enum))
	{
		if err := w.WriteEnum(uint16(o.SANPolicy)); err != nil {
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

func (o *xxx_GetSANPolicyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pSanPolicy {out} (1:{pointer=ref}*(1))(2:{alias=VDS_SAN_POLICY}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.SANPolicy)); err != nil {
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

// GetSANPolicyRequest structure represents the GetSANPolicy operation request
type GetSANPolicyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSANPolicyRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSANPolicyOperation) *xxx_GetSANPolicyOperation {
	if op == nil {
		op = &xxx_GetSANPolicyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetSANPolicyRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSANPolicyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSANPolicyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSANPolicyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSANPolicyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSANPolicyResponse structure represents the GetSANPolicy operation response
type GetSANPolicyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pSanPolicy: A pointer to a VDS_SAN_POLICY (section 2.2.2.2.1.1) structure that, if
	// the operation is successfully completed, receives the SAN policy setting's current
	// value.
	SANPolicy vds.SANPolicy `idl:"name:pSanPolicy" json:"san_policy"`
	// Return: The GetSANPolicy return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSANPolicyResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSANPolicyOperation) *xxx_GetSANPolicyOperation {
	if op == nil {
		op = &xxx_GetSANPolicyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.SANPolicy = o.SANPolicy
	op.Return = o.Return
	return op
}

func (o *GetSANPolicyResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSANPolicyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SANPolicy = op.SANPolicy
	o.Return = op.Return
}
func (o *GetSANPolicyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSANPolicyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSANPolicyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSANPolicyOperation structure represents the SetSANPolicy operation
type xxx_SetSANPolicyOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	SANPolicy vds.SANPolicy  `idl:"name:SanPolicy" json:"san_policy"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSANPolicyOperation) OpNum() int { return 4 }

func (o *xxx_SetSANPolicyOperation) OpName() string { return "/IVdsServiceSAN/v0/SetSANPolicy" }

func (o *xxx_SetSANPolicyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSANPolicyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// SanPolicy {in} (1:{alias=VDS_SAN_POLICY}(enum))
	{
		if err := w.WriteEnum(uint16(o.SANPolicy)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSANPolicyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// SanPolicy {in} (1:{alias=VDS_SAN_POLICY}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.SANPolicy)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSANPolicyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSANPolicyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSANPolicyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetSANPolicyRequest structure represents the SetSANPolicy operation request
type SetSANPolicyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// SanPolicy: A VDS_SAN_POLICY (section 2.2.2.2.1.1) structure that, if the operation
	// is successfully completed, is used to set the new value for the SAN policy.
	SANPolicy vds.SANPolicy `idl:"name:SanPolicy" json:"san_policy"`
}

func (o *SetSANPolicyRequest) xxx_ToOp(ctx context.Context, op *xxx_SetSANPolicyOperation) *xxx_SetSANPolicyOperation {
	if op == nil {
		op = &xxx_SetSANPolicyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.SANPolicy = o.SANPolicy
	return op
}

func (o *SetSANPolicyRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSANPolicyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SANPolicy = op.SANPolicy
}
func (o *SetSANPolicyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetSANPolicyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSANPolicyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSANPolicyResponse structure represents the SetSANPolicy operation response
type SetSANPolicyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetSANPolicy return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSANPolicyResponse) xxx_ToOp(ctx context.Context, op *xxx_SetSANPolicyOperation) *xxx_SetSANPolicyOperation {
	if op == nil {
		op = &xxx_SetSANPolicyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetSANPolicyResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSANPolicyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSANPolicyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetSANPolicyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSANPolicyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
