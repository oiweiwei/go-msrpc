package iremotenetworkconfig

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
)

var (
	// import guard
	GoPackage = "dcom/rrasm"
)

var (
	// IRemoteNetworkConfig interface identifier 66a2db1b-d706-11d0-a37b-00c04fc9da04
	RemoteNetworkConfigIID = &dcom.IID{Data1: 0x66a2db1b, Data2: 0xd706, Data3: 0x11d0, Data4: []byte{0xa3, 0x7b, 0x00, 0xc0, 0x4f, 0xc9, 0xda, 0x04}}
	// Syntax UUID
	RemoteNetworkConfigSyntaxUUID = &uuid.UUID{TimeLow: 0x66a2db1b, TimeMid: 0xd706, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xa3, ClockSeqLow: 0x7b, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xc9, 0xda, 0x4}}
	// Syntax ID
	RemoteNetworkConfigSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: RemoteNetworkConfigSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IRemoteNetworkConfig interface.
type RemoteNetworkConfigClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// UpgradeRouterConfig operation.
	UpgradeRouterConfig(context.Context, *UpgradeRouterConfigRequest, ...dcerpc.CallOption) (*UpgradeRouterConfigResponse, error)

	// SetUserConfig operation.
	SetUserConfig(context.Context, *SetUserConfigRequest, ...dcerpc.CallOption) (*SetUserConfigResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) RemoteNetworkConfigClient
}

type xxx_DefaultRemoteNetworkConfigClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultRemoteNetworkConfigClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultRemoteNetworkConfigClient) UpgradeRouterConfig(ctx context.Context, in *UpgradeRouterConfigRequest, opts ...dcerpc.CallOption) (*UpgradeRouterConfigResponse, error) {
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
	out := &UpgradeRouterConfigResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteNetworkConfigClient) SetUserConfig(ctx context.Context, in *SetUserConfigRequest, opts ...dcerpc.CallOption) (*SetUserConfigResponse, error) {
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
	out := &SetUserConfigResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteNetworkConfigClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultRemoteNetworkConfigClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultRemoteNetworkConfigClient) IPID(ctx context.Context, ipid *dcom.IPID) RemoteNetworkConfigClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultRemoteNetworkConfigClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewRemoteNetworkConfigClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (RemoteNetworkConfigClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(RemoteNetworkConfigSyntaxV0_0))...)
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
	return &xxx_DefaultRemoteNetworkConfigClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_UpgradeRouterConfigOperation structure represents the UpgradeRouterConfig operation
type xxx_UpgradeRouterConfigOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_UpgradeRouterConfigOperation) OpNum() int { return 3 }

func (o *xxx_UpgradeRouterConfigOperation) OpName() string {
	return "/IRemoteNetworkConfig/v0/UpgradeRouterConfig"
}

func (o *xxx_UpgradeRouterConfigOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpgradeRouterConfigOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_UpgradeRouterConfigOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_UpgradeRouterConfigOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpgradeRouterConfigOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_UpgradeRouterConfigOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// UpgradeRouterConfigRequest structure represents the UpgradeRouterConfig operation request
type UpgradeRouterConfigRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *UpgradeRouterConfigRequest) xxx_ToOp(ctx context.Context, op *xxx_UpgradeRouterConfigOperation) *xxx_UpgradeRouterConfigOperation {
	if op == nil {
		op = &xxx_UpgradeRouterConfigOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *UpgradeRouterConfigRequest) xxx_FromOp(ctx context.Context, op *xxx_UpgradeRouterConfigOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *UpgradeRouterConfigRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *UpgradeRouterConfigRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UpgradeRouterConfigOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// UpgradeRouterConfigResponse structure represents the UpgradeRouterConfig operation response
type UpgradeRouterConfigResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The UpgradeRouterConfig return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *UpgradeRouterConfigResponse) xxx_ToOp(ctx context.Context, op *xxx_UpgradeRouterConfigOperation) *xxx_UpgradeRouterConfigOperation {
	if op == nil {
		op = &xxx_UpgradeRouterConfigOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *UpgradeRouterConfigResponse) xxx_FromOp(ctx context.Context, op *xxx_UpgradeRouterConfigOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *UpgradeRouterConfigResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *UpgradeRouterConfigResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UpgradeRouterConfigOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetUserConfigOperation structure represents the SetUserConfig operation
type xxx_SetUserConfigOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Service  string         `idl:"name:pszService" json:"service"`
	NewGroup string         `idl:"name:pszNewGroup" json:"new_group"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetUserConfigOperation) OpNum() int { return 4 }

func (o *xxx_SetUserConfigOperation) OpName() string { return "/IRemoteNetworkConfig/v0/SetUserConfig" }

func (o *xxx_SetUserConfigOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetUserConfigOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pszService {in} (1:{string, alias=LPCOLESTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Service); err != nil {
			return err
		}
	}
	// pszNewGroup {in} (1:{string, alias=LPCOLESTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.NewGroup); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetUserConfigOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pszService {in} (1:{string, alias=LPCOLESTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Service); err != nil {
			return err
		}
	}
	// pszNewGroup {in} (1:{string, alias=LPCOLESTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.NewGroup); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetUserConfigOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetUserConfigOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetUserConfigOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetUserConfigRequest structure represents the SetUserConfig operation request
type SetUserConfigRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	Service  string         `idl:"name:pszService" json:"service"`
	NewGroup string         `idl:"name:pszNewGroup" json:"new_group"`
}

func (o *SetUserConfigRequest) xxx_ToOp(ctx context.Context, op *xxx_SetUserConfigOperation) *xxx_SetUserConfigOperation {
	if op == nil {
		op = &xxx_SetUserConfigOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Service = o.Service
	op.NewGroup = o.NewGroup
	return op
}

func (o *SetUserConfigRequest) xxx_FromOp(ctx context.Context, op *xxx_SetUserConfigOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Service = op.Service
	o.NewGroup = op.NewGroup
}
func (o *SetUserConfigRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetUserConfigRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetUserConfigOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetUserConfigResponse structure represents the SetUserConfig operation response
type SetUserConfigResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetUserConfig return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetUserConfigResponse) xxx_ToOp(ctx context.Context, op *xxx_SetUserConfigOperation) *xxx_SetUserConfigOperation {
	if op == nil {
		op = &xxx_SetUserConfigOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetUserConfigResponse) xxx_FromOp(ctx context.Context, op *xxx_SetUserConfigOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetUserConfigResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetUserConfigResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetUserConfigOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
