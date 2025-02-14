package ivdshbaport

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
	// IVdsHbaPort interface identifier 2abd757f-2851-4997-9a13-47d2a885d6ca
	HBAPortIID = &dcom.IID{Data1: 0x2abd757f, Data2: 0x2851, Data3: 0x4997, Data4: []byte{0x9a, 0x13, 0x47, 0xd2, 0xa8, 0x85, 0xd6, 0xca}}
	// Syntax UUID
	HBAPortSyntaxUUID = &uuid.UUID{TimeLow: 0x2abd757f, TimeMid: 0x2851, TimeHiAndVersion: 0x4997, ClockSeqHiAndReserved: 0x9a, ClockSeqLow: 0x13, Node: [6]uint8{0x47, 0xd2, 0xa8, 0x85, 0xd6, 0xca}}
	// Syntax ID
	HBAPortSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: HBAPortSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsHbaPort interface.
type HBAPortClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// GetProperties operation.
	GetProperties(context.Context, *GetPropertiesRequest, ...dcerpc.CallOption) (*GetPropertiesResponse, error)

	// The SetAllPathStatuses method sets the statuses of all paths that originate from
	// the HBA port to a specified status.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure.<76> For the HRESULT values predefined by the Virtual Disk
	// Service Remote Protocol, see section 2.2.3.
	SetAllPathStatuses(context.Context, *SetAllPathStatusesRequest, ...dcerpc.CallOption) (*SetAllPathStatusesResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) HBAPortClient
}

type xxx_DefaultHBAPortClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultHBAPortClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultHBAPortClient) GetProperties(ctx context.Context, in *GetPropertiesRequest, opts ...dcerpc.CallOption) (*GetPropertiesResponse, error) {
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
	out := &GetPropertiesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultHBAPortClient) SetAllPathStatuses(ctx context.Context, in *SetAllPathStatusesRequest, opts ...dcerpc.CallOption) (*SetAllPathStatusesResponse, error) {
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
	out := &SetAllPathStatusesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultHBAPortClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultHBAPortClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultHBAPortClient) IPID(ctx context.Context, ipid *dcom.IPID) HBAPortClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultHBAPortClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewHBAPortClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (HBAPortClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(HBAPortSyntaxV0_0))...)
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
	return &xxx_DefaultHBAPortClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetPropertiesOperation structure represents the GetProperties operation
type xxx_GetPropertiesOperation struct {
	This            *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat       `idl:"name:That" json:"that"`
	HBAPortProperty *vds.HBAPortProperty `idl:"name:pHbaPortProp" json:"hba_port_property"`
	Return          int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPropertiesOperation) OpNum() int { return 3 }

func (o *xxx_GetPropertiesOperation) OpName() string { return "/IVdsHbaPort/v0/GetProperties" }

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
	// pHbaPortProp {out} (1:{pointer=ref}*(1))(2:{alias=VDS_HBAPORT_PROP}(struct))
	{
		if o.HBAPortProperty != nil {
			if err := o.HBAPortProperty.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&vds.HBAPortProperty{}).MarshalNDR(ctx, w); err != nil {
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
	// pHbaPortProp {out} (1:{pointer=ref}*(1))(2:{alias=VDS_HBAPORT_PROP}(struct))
	{
		if o.HBAPortProperty == nil {
			o.HBAPortProperty = &vds.HBAPortProperty{}
		}
		if err := o.HBAPortProperty.UnmarshalNDR(ctx, w); err != nil {
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

func (o *GetPropertiesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPropertiesOperation) *xxx_GetPropertiesOperation {
	if op == nil {
		op = &xxx_GetPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetPropertiesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPropertiesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
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
	That            *dcom.ORPCThat       `idl:"name:That" json:"that"`
	HBAPortProperty *vds.HBAPortProperty `idl:"name:pHbaPortProp" json:"hba_port_property"`
	// Return: The GetProperties return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPropertiesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPropertiesOperation) *xxx_GetPropertiesOperation {
	if op == nil {
		op = &xxx_GetPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.HBAPortProperty = o.HBAPortProperty
	op.Return = o.Return
	return op
}

func (o *GetPropertiesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.HBAPortProperty = op.HBAPortProperty
	o.Return = op.Return
}
func (o *GetPropertiesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPropertiesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetAllPathStatusesOperation structure represents the SetAllPathStatuses operation
type xxx_SetAllPathStatusesOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Status vds.PathStatus `idl:"name:status" json:"status"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetAllPathStatusesOperation) OpNum() int { return 4 }

func (o *xxx_SetAllPathStatusesOperation) OpName() string {
	return "/IVdsHbaPort/v0/SetAllPathStatuses"
}

func (o *xxx_SetAllPathStatusesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAllPathStatusesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// status {in} (1:{alias=VDS_PATH_STATUS}(enum))
	{
		if err := w.WriteEnum(uint16(o.Status)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAllPathStatusesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// status {in} (1:{alias=VDS_PATH_STATUS}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Status)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAllPathStatusesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAllPathStatusesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetAllPathStatusesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetAllPathStatusesRequest structure represents the SetAllPathStatuses operation request
type SetAllPathStatusesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// status: The status, as defined by VDS_PATH_STATUS, to assign to the paths.
	Status vds.PathStatus `idl:"name:status" json:"status"`
}

func (o *SetAllPathStatusesRequest) xxx_ToOp(ctx context.Context, op *xxx_SetAllPathStatusesOperation) *xxx_SetAllPathStatusesOperation {
	if op == nil {
		op = &xxx_SetAllPathStatusesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Status = o.Status
	return op
}

func (o *SetAllPathStatusesRequest) xxx_FromOp(ctx context.Context, op *xxx_SetAllPathStatusesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Status = op.Status
}
func (o *SetAllPathStatusesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetAllPathStatusesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAllPathStatusesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetAllPathStatusesResponse structure represents the SetAllPathStatuses operation response
type SetAllPathStatusesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetAllPathStatuses return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetAllPathStatusesResponse) xxx_ToOp(ctx context.Context, op *xxx_SetAllPathStatusesOperation) *xxx_SetAllPathStatusesOperation {
	if op == nil {
		op = &xxx_SetAllPathStatusesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetAllPathStatusesResponse) xxx_FromOp(ctx context.Context, op *xxx_SetAllPathStatusesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetAllPathStatusesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetAllPathStatusesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAllPathStatusesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
