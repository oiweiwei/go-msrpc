package ivdsiscsiinitiatorportal

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
	_ = (*errors.Error)(nil)
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
	_ = vds.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/vds"
)

var (
	// IVdsIscsiInitiatorPortal interface identifier 38a0a9ab-7cc8-4693-ac07-1f28bd03c3da
	ISCSIInitiatorPortalIID = &dcom.IID{Data1: 0x38a0a9ab, Data2: 0x7cc8, Data3: 0x4693, Data4: []byte{0xac, 0x07, 0x1f, 0x28, 0xbd, 0x03, 0xc3, 0xda}}
	// Syntax UUID
	ISCSIInitiatorPortalSyntaxUUID = &uuid.UUID{TimeLow: 0x38a0a9ab, TimeMid: 0x7cc8, TimeHiAndVersion: 0x4693, ClockSeqHiAndReserved: 0xac, ClockSeqLow: 0x7, Node: [6]uint8{0x1f, 0x28, 0xbd, 0x3, 0xc3, 0xda}}
	// Syntax ID
	ISCSIInitiatorPortalSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ISCSIInitiatorPortalSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsIscsiInitiatorPortal interface.
type ISCSIInitiatorPortalClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// GetProperties operation.
	GetProperties(context.Context, *GetPropertiesRequest, ...dcerpc.CallOption) (*GetPropertiesResponse, error)

	// The GetInitiatorAdapter method returns the initiator adapter to the initiator portal
	// it belongs to.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetInitiatorAdapter(context.Context, *GetInitiatorAdapterRequest, ...dcerpc.CallOption) (*GetInitiatorAdapterResponse, error)

	// Opnum05NotUsedOnWire operation.
	// Opnum05NotUsedOnWire

	// Opnum06NotUsedOnWire operation.
	// Opnum06NotUsedOnWire

	// Opnum07NotUsedOnWire operation.
	// Opnum07NotUsedOnWire

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ISCSIInitiatorPortalClient
}

type xxx_DefaultISCSIInitiatorPortalClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultISCSIInitiatorPortalClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultISCSIInitiatorPortalClient) GetProperties(ctx context.Context, in *GetPropertiesRequest, opts ...dcerpc.CallOption) (*GetPropertiesResponse, error) {
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
	out := &GetPropertiesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultISCSIInitiatorPortalClient) GetInitiatorAdapter(ctx context.Context, in *GetInitiatorAdapterRequest, opts ...dcerpc.CallOption) (*GetInitiatorAdapterResponse, error) {
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
	out := &GetInitiatorAdapterResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultISCSIInitiatorPortalClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultISCSIInitiatorPortalClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultISCSIInitiatorPortalClient) IPID(ctx context.Context, ipid *dcom.IPID) ISCSIInitiatorPortalClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultISCSIInitiatorPortalClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewISCSIInitiatorPortalClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ISCSIInitiatorPortalClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ISCSIInitiatorPortalSyntaxV0_0))...)
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
	return &xxx_DefaultISCSIInitiatorPortalClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetPropertiesOperation structure represents the GetProperties operation
type xxx_GetPropertiesOperation struct {
	This                    *dcom.ORPCThis                    `idl:"name:This" json:"this"`
	That                    *dcom.ORPCThat                    `idl:"name:That" json:"that"`
	InitiatorPortalProperty *vds.ISCSIInitiatorPortalProperty `idl:"name:pInitiatorPortalProp" json:"initiator_portal_property"`
	Return                  int32                             `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPropertiesOperation) OpNum() int { return 3 }

func (o *xxx_GetPropertiesOperation) OpName() string {
	return "/IVdsIscsiInitiatorPortal/v0/GetProperties"
}

func (o *xxx_GetPropertiesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPropertiesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPropertiesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pInitiatorPortalProp {out} (1:{pointer=ref}*(1))(2:{alias=VDS_ISCSI_INITIATOR_PORTAL_PROP}(struct))
	{
		if o.InitiatorPortalProperty != nil {
			if err := o.InitiatorPortalProperty.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&vds.ISCSIInitiatorPortalProperty{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetPropertiesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pInitiatorPortalProp {out} (1:{pointer=ref}*(1))(2:{alias=VDS_ISCSI_INITIATOR_PORTAL_PROP}(struct))
	{
		if o.InitiatorPortalProperty == nil {
			o.InitiatorPortalProperty = &vds.ISCSIInitiatorPortalProperty{}
		}
		if err := o.InitiatorPortalProperty.UnmarshalNDR(ctx, w); err != nil {
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

// GetPropertiesRequest structure represents the GetProperties operation request
type GetPropertiesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPropertiesRequest) xxx_ToOp(ctx context.Context) *xxx_GetPropertiesOperation {
	if o == nil {
		return &xxx_GetPropertiesOperation{}
	}
	return &xxx_GetPropertiesOperation{
		This: o.This,
	}
}

func (o *GetPropertiesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPropertiesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetPropertiesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPropertiesResponse structure represents the GetProperties operation response
type GetPropertiesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That                    *dcom.ORPCThat                    `idl:"name:That" json:"that"`
	InitiatorPortalProperty *vds.ISCSIInitiatorPortalProperty `idl:"name:pInitiatorPortalProp" json:"initiator_portal_property"`
	// Return: The GetProperties return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPropertiesResponse) xxx_ToOp(ctx context.Context) *xxx_GetPropertiesOperation {
	if o == nil {
		return &xxx_GetPropertiesOperation{}
	}
	return &xxx_GetPropertiesOperation{
		That:                    o.That,
		InitiatorPortalProperty: o.InitiatorPortalProperty,
		Return:                  o.Return,
	}
}

func (o *GetPropertiesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.InitiatorPortalProperty = op.InitiatorPortalProperty
	o.Return = op.Return
}
func (o *GetPropertiesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetPropertiesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetInitiatorAdapterOperation structure represents the GetInitiatorAdapter operation
type xxx_GetInitiatorAdapterOperation struct {
	This             *dcom.ORPCThis             `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat             `idl:"name:That" json:"that"`
	InitiatorAdapter *vds.ISCSIInitiatorAdapter `idl:"name:ppInitiatorAdapter" json:"initiator_adapter"`
	Return           int32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_GetInitiatorAdapterOperation) OpNum() int { return 4 }

func (o *xxx_GetInitiatorAdapterOperation) OpName() string {
	return "/IVdsIscsiInitiatorPortal/v0/GetInitiatorAdapter"
}

func (o *xxx_GetInitiatorAdapterOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInitiatorAdapterOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetInitiatorAdapterOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetInitiatorAdapterOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInitiatorAdapterOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppInitiatorAdapter {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsIscsiInitiatorAdapter}(interface))
	{
		if o.InitiatorAdapter != nil {
			_ptr_ppInitiatorAdapter := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.InitiatorAdapter != nil {
					if err := o.InitiatorAdapter.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&vds.ISCSIInitiatorAdapter{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InitiatorAdapter, _ptr_ppInitiatorAdapter); err != nil {
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

func (o *xxx_GetInitiatorAdapterOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppInitiatorAdapter {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsIscsiInitiatorAdapter}(interface))
	{
		_ptr_ppInitiatorAdapter := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.InitiatorAdapter == nil {
				o.InitiatorAdapter = &vds.ISCSIInitiatorAdapter{}
			}
			if err := o.InitiatorAdapter.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppInitiatorAdapter := func(ptr interface{}) { o.InitiatorAdapter = *ptr.(**vds.ISCSIInitiatorAdapter) }
		if err := w.ReadPointer(&o.InitiatorAdapter, _s_ppInitiatorAdapter, _ptr_ppInitiatorAdapter); err != nil {
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

// GetInitiatorAdapterRequest structure represents the GetInitiatorAdapter operation request
type GetInitiatorAdapterRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetInitiatorAdapterRequest) xxx_ToOp(ctx context.Context) *xxx_GetInitiatorAdapterOperation {
	if o == nil {
		return &xxx_GetInitiatorAdapterOperation{}
	}
	return &xxx_GetInitiatorAdapterOperation{
		This: o.This,
	}
}

func (o *GetInitiatorAdapterRequest) xxx_FromOp(ctx context.Context, op *xxx_GetInitiatorAdapterOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetInitiatorAdapterRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetInitiatorAdapterRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetInitiatorAdapterOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetInitiatorAdapterResponse structure represents the GetInitiatorAdapter operation response
type GetInitiatorAdapterResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppInitiatorAdapter: A pointer to an IVdsIscsiInitiatorAdapter interface that, if
	// the operation is successfully completed, receives the IVdsIscsiInitiatorAdapter interface
	// of the initiator adapter object that the initiator portal belongs to. Callers MUST
	// release the interface when they are done with it.
	InitiatorAdapter *vds.ISCSIInitiatorAdapter `idl:"name:ppInitiatorAdapter" json:"initiator_adapter"`
	// Return: The GetInitiatorAdapter return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetInitiatorAdapterResponse) xxx_ToOp(ctx context.Context) *xxx_GetInitiatorAdapterOperation {
	if o == nil {
		return &xxx_GetInitiatorAdapterOperation{}
	}
	return &xxx_GetInitiatorAdapterOperation{
		That:             o.That,
		InitiatorAdapter: o.InitiatorAdapter,
		Return:           o.Return,
	}
}

func (o *GetInitiatorAdapterResponse) xxx_FromOp(ctx context.Context, op *xxx_GetInitiatorAdapterOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.InitiatorAdapter = op.InitiatorAdapter
	o.Return = op.Return
}
func (o *GetInitiatorAdapterResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetInitiatorAdapterResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetInitiatorAdapterOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
