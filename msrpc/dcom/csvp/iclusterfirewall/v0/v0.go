package iclusterfirewall

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	csvp "github.com/oiweiwei/go-msrpc/msrpc/dcom/csvp"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
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
	_ = (*errors.Error)(nil)
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
	_ = dtyp.GoPackage
	_ = csvp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/csvp"
)

var (
	// IClusterFirewall interface identifier f1d6c29c-8fbe-4691-8724-f6d8deaeafc8
	ClusterFirewallIID = &dcom.IID{Data1: 0xf1d6c29c, Data2: 0x8fbe, Data3: 0x4691, Data4: []byte{0x87, 0x24, 0xf6, 0xd8, 0xde, 0xae, 0xaf, 0xc8}}
	// Syntax UUID
	ClusterFirewallSyntaxUUID = &uuid.UUID{TimeLow: 0xf1d6c29c, TimeMid: 0x8fbe, TimeHiAndVersion: 0x4691, ClockSeqHiAndReserved: 0x87, ClockSeqLow: 0x24, Node: [6]uint8{0xf6, 0xd8, 0xde, 0xae, 0xaf, 0xc8}}
	// Syntax ID
	ClusterFirewallSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ClusterFirewallSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IClusterFirewall interface.
type ClusterFirewallClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The InitializeAdapterConfiguration method initializes the server Firewall State to
	// process subsequent calls of GetNextAdapterFirewallConfiguration.
	//
	// This method is called at least once before GetNextAdapterFirewallConfiguration.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 3.
	InitializeAdapterConfiguration(context.Context, *InitializeAdapterConfigurationRequest, ...dcerpc.CallOption) (*InitializeAdapterConfigurationResponse, error)

	// The GetNextAdapterFirewallConfiguration method returns information about a specific
	// network adapter attached to the system.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	|         RETURN          |                                                                                  |
	//	|       VALUE/CODE        |                                   DESCRIPTION                                    |
	//	|                         |                                                                                  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK         | The call was successful.                                                         |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | The value the client specified in idx is greater than or equal                   |
	//	|                         | to the cRetAdapters value returned by the previous call to                       |
	//	|                         | InitializeAdapterConfiguration.                                                  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x8000FFFF E_UNEXPECTED | InitializeAdapterConfiguration has not yet been called.                          |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 4.
	//
	// The server returns the following information to the client:
	//
	// * The output parameters set to the values specified previously.
	GetNextAdapterFirewallConfiguration(context.Context, *GetNextAdapterFirewallConfigurationRequest, ...dcerpc.CallOption) (*GetNextAdapterFirewallConfigurationResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ClusterFirewallClient
}

type xxx_DefaultClusterFirewallClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultClusterFirewallClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultClusterFirewallClient) InitializeAdapterConfiguration(ctx context.Context, in *InitializeAdapterConfigurationRequest, opts ...dcerpc.CallOption) (*InitializeAdapterConfigurationResponse, error) {
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
	out := &InitializeAdapterConfigurationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterFirewallClient) GetNextAdapterFirewallConfiguration(ctx context.Context, in *GetNextAdapterFirewallConfigurationRequest, opts ...dcerpc.CallOption) (*GetNextAdapterFirewallConfigurationResponse, error) {
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
	out := &GetNextAdapterFirewallConfigurationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterFirewallClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultClusterFirewallClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultClusterFirewallClient) IPID(ctx context.Context, ipid *dcom.IPID) ClusterFirewallClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultClusterFirewallClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewClusterFirewallClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ClusterFirewallClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ClusterFirewallSyntaxV0_0))...)
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
	return &xxx_DefaultClusterFirewallClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_InitializeAdapterConfigurationOperation structure represents the InitializeAdapterConfiguration operation
type xxx_InitializeAdapterConfigurationOperation struct {
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReturnAdaptersCount uint32         `idl:"name:cRetAdapters" json:"return_adapters_count"`
	Return              int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_InitializeAdapterConfigurationOperation) OpNum() int { return 3 }

func (o *xxx_InitializeAdapterConfigurationOperation) OpName() string {
	return "/IClusterFirewall/v0/InitializeAdapterConfiguration"
}

func (o *xxx_InitializeAdapterConfigurationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeAdapterConfigurationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_InitializeAdapterConfigurationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_InitializeAdapterConfigurationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeAdapterConfigurationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// cRetAdapters {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.ReturnAdaptersCount); err != nil {
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

func (o *xxx_InitializeAdapterConfigurationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// cRetAdapters {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.ReturnAdaptersCount); err != nil {
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

// InitializeAdapterConfigurationRequest structure represents the InitializeAdapterConfiguration operation request
type InitializeAdapterConfigurationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *InitializeAdapterConfigurationRequest) xxx_ToOp(ctx context.Context) *xxx_InitializeAdapterConfigurationOperation {
	if o == nil {
		return &xxx_InitializeAdapterConfigurationOperation{}
	}
	return &xxx_InitializeAdapterConfigurationOperation{
		This: o.This,
	}
}

func (o *InitializeAdapterConfigurationRequest) xxx_FromOp(ctx context.Context, op *xxx_InitializeAdapterConfigurationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *InitializeAdapterConfigurationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *InitializeAdapterConfigurationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InitializeAdapterConfigurationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// InitializeAdapterConfigurationResponse structure represents the InitializeAdapterConfiguration operation response
type InitializeAdapterConfigurationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// cRetAdapters: A pointer to an unsigned 32-bit integer indicating the number of adapters
	// in the network adapter index of the Firewall State. Upon successful completion of
	// this method, the server MUST set this value. If the method fails, the client MUST
	// ignore this value.
	ReturnAdaptersCount uint32 `idl:"name:cRetAdapters" json:"return_adapters_count"`
	// Return: The InitializeAdapterConfiguration return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *InitializeAdapterConfigurationResponse) xxx_ToOp(ctx context.Context) *xxx_InitializeAdapterConfigurationOperation {
	if o == nil {
		return &xxx_InitializeAdapterConfigurationOperation{}
	}
	return &xxx_InitializeAdapterConfigurationOperation{
		That:                o.That,
		ReturnAdaptersCount: o.ReturnAdaptersCount,
		Return:              o.Return,
	}
}

func (o *InitializeAdapterConfigurationResponse) xxx_FromOp(ctx context.Context, op *xxx_InitializeAdapterConfigurationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnAdaptersCount = op.ReturnAdaptersCount
	o.Return = op.Return
}
func (o *InitializeAdapterConfigurationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *InitializeAdapterConfigurationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InitializeAdapterConfigurationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNextAdapterFirewallConfigurationOperation structure represents the GetNextAdapterFirewallConfiguration operation
type xxx_GetNextAdapterFirewallConfigurationOperation struct {
	This                   *dcom.ORPCThis             `idl:"name:This" json:"this"`
	That                   *dcom.ORPCThat             `idl:"name:That" json:"that"`
	Index                  uint32                     `idl:"name:idx" json:"index"`
	AdapterID              *dtyp.GUID                 `idl:"name:adapterId" json:"adapter_id"`
	AdapterProfile         csvp.ClusterNetworkProfile `idl:"name:adapterProfile" json:"adapter_profile"`
	ServerRulesEnabled     bool                       `idl:"name:serverRulesEnabled" json:"server_rules_enabled"`
	ManagementRulesEnabled bool                       `idl:"name:managementRulesEnabled" json:"management_rules_enabled"`
	CommonRulesEnabled     bool                       `idl:"name:commonRulesEnabled" json:"common_rules_enabled"`
	Return                 int32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNextAdapterFirewallConfigurationOperation) OpNum() int { return 4 }

func (o *xxx_GetNextAdapterFirewallConfigurationOperation) OpName() string {
	return "/IClusterFirewall/v0/GetNextAdapterFirewallConfiguration"
}

func (o *xxx_GetNextAdapterFirewallConfigurationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNextAdapterFirewallConfigurationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// idx {in} (1:(uint32))
	{
		if err := w.WriteData(o.Index); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNextAdapterFirewallConfigurationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// idx {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Index); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNextAdapterFirewallConfigurationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNextAdapterFirewallConfigurationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// adapterId {out} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.AdapterID != nil {
			if err := o.AdapterID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// adapterProfile {out} (1:{pointer=ref}*(1))(2:{alias=CLUSTER_NETWORK_PROFILE}(enum))
	{
		if err := w.WriteData(uint16(o.AdapterProfile)); err != nil {
			return err
		}
	}
	// serverRulesEnabled {out} (1:{pointer=ref}*(1))(2:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.ServerRulesEnabled); err != nil {
			return err
		}
	}
	// managementRulesEnabled {out} (1:{pointer=ref}*(1))(2:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.ManagementRulesEnabled); err != nil {
			return err
		}
	}
	// commonRulesEnabled {out} (1:{pointer=ref}*(1))(2:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.CommonRulesEnabled); err != nil {
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

func (o *xxx_GetNextAdapterFirewallConfigurationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// adapterId {out} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.AdapterID == nil {
			o.AdapterID = &dtyp.GUID{}
		}
		if err := o.AdapterID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// adapterProfile {out} (1:{pointer=ref}*(1))(2:{alias=CLUSTER_NETWORK_PROFILE}(enum))
	{
		if err := w.ReadData((*uint16)(&o.AdapterProfile)); err != nil {
			return err
		}
	}
	// serverRulesEnabled {out} (1:{pointer=ref}*(1))(2:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.ServerRulesEnabled); err != nil {
			return err
		}
	}
	// managementRulesEnabled {out} (1:{pointer=ref}*(1))(2:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.ManagementRulesEnabled); err != nil {
			return err
		}
	}
	// commonRulesEnabled {out} (1:{pointer=ref}*(1))(2:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.CommonRulesEnabled); err != nil {
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

// GetNextAdapterFirewallConfigurationRequest structure represents the GetNextAdapterFirewallConfiguration operation request
type GetNextAdapterFirewallConfigurationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// idx: A 32-bit unsigned integer that indicates the index of the adapter information
	// to retrieve. The server MUST fail this method with error 0x80070057 (E_INVALIDARG)
	// if idx is greater than or equal to the cRetAdapters value returned by the previous
	// call to InitializeAdapterConfiguration (Opnum 3).
	Index uint32 `idl:"name:idx" json:"index"`
}

func (o *GetNextAdapterFirewallConfigurationRequest) xxx_ToOp(ctx context.Context) *xxx_GetNextAdapterFirewallConfigurationOperation {
	if o == nil {
		return &xxx_GetNextAdapterFirewallConfigurationOperation{}
	}
	return &xxx_GetNextAdapterFirewallConfigurationOperation{
		This:  o.This,
		Index: o.Index,
	}
}

func (o *GetNextAdapterFirewallConfigurationRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNextAdapterFirewallConfigurationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Index = op.Index
}
func (o *GetNextAdapterFirewallConfigurationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetNextAdapterFirewallConfigurationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNextAdapterFirewallConfigurationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNextAdapterFirewallConfigurationResponse structure represents the GetNextAdapterFirewallConfiguration operation response
type GetNextAdapterFirewallConfigurationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// adapterId: A GUID that uniquely identifies the network adapter on the system. Upon
	// successful completion of this method, the server MUST set this value. If the method
	// fails, the client MUST ignore this value.
	AdapterID *dtyp.GUID `idl:"name:adapterId" json:"adapter_id"`
	// adapterProfile: The firewall profile assigned to the network adapter. Upon successful
	// completion of this method, the server MUST set this value to one of the specified
	// values of CLUSTER_NETWORK_PROFILE. If the method fails, the client MUST ignore this
	// value.
	AdapterProfile csvp.ClusterNetworkProfile `idl:"name:adapterProfile" json:"adapter_profile"`
	// serverRulesEnabled: An output parameter that indicates whether the server is suitable
	// for server-to-server failover cluster communication. Upon successful completion of
	// this method, the server MUST set this value to TRUE if the server is suitable or
	// to FALSE if the server is not suitable. When the server firewall enforces policies
	// specified in [MS-FASP], the server sets this value to TRUE if the group of rules
	// with the localized name "Failover Clusters" is enabled. If the method fails, the
	// client MUST ignore this value.
	//
	//	+----------------+------------------------------------------------------------------+
	//	|                |                                                                  |
	//	|     VALUE      |                             MEANING                              |
	//	|                |                                                                  |
	//	+----------------+------------------------------------------------------------------+
	//	+----------------+------------------------------------------------------------------+
	//	| TRUE -128 — -1 | Firewall settings allow the traffic specified previously.        |
	//	+----------------+------------------------------------------------------------------+
	//	| FALSE 0        | Firewall settings do not allow the traffic specified previously. |
	//	+----------------+------------------------------------------------------------------+
	//	| TRUE 1 — 128   | Firewall settings allow the traffic specified previously.        |
	//	+----------------+------------------------------------------------------------------+
	ServerRulesEnabled bool `idl:"name:serverRulesEnabled" json:"server_rules_enabled"`
	// managementRulesEnabled: An output parameter that indicates whether the server is
	// compatible with the failover cluster management components. Upon successful completion
	// of this method, the server MUST set this value to TRUE if the server is compatible
	// or to FALSE if the server is not compatible. When the server firewall enforces policies
	// specified in [MS-FASP], the server SHOULD set this value to TRUE if the group of
	// rules with the localized name "Failover Cluster Manager" is enabled. If the method
	// fails, the client MUST ignore this value.
	//
	//	+----------------+------------------------------------------------------------------+
	//	|                |                                                                  |
	//	|     VALUE      |                             MEANING                              |
	//	|                |                                                                  |
	//	+----------------+------------------------------------------------------------------+
	//	+----------------+------------------------------------------------------------------+
	//	| TRUE -128 — -1 | Firewall settings allow the traffic specified previously.        |
	//	+----------------+------------------------------------------------------------------+
	//	| FALSE 0        | Firewall settings do not allow the traffic specified previously. |
	//	+----------------+------------------------------------------------------------------+
	//	| TRUE 1 — 128   | Firewall settings allow the traffic specified previously.        |
	//	+----------------+------------------------------------------------------------------+
	ManagementRulesEnabled bool `idl:"name:managementRulesEnabled" json:"management_rules_enabled"`
	// commonRulesEnabled: An output parameter that indicates whether the server is compatible
	// with the failover cluster components common to failover cluster management and server-to-server
	// failover cluster communications. Upon successful completion of this method, the server
	// MUST set this value to TRUE if the server is compatible or to FALSE if the server
	// is not compatible. When the server firewall enforces policies specified in [MS-FASP],
	// the server SHOULD set this value to TRUE if the group of rules with the localized
	// name "Failover Cluster Common" is enabled. If the method fails, the client MUST ignore
	// this value.
	//
	//	+----------------+------------------------------------------------------------------+
	//	|                |                                                                  |
	//	|     VALUE      |                             MEANING                              |
	//	|                |                                                                  |
	//	+----------------+------------------------------------------------------------------+
	//	+----------------+------------------------------------------------------------------+
	//	| TRUE -128 — -1 | Firewall settings allow the traffic specified previously.        |
	//	+----------------+------------------------------------------------------------------+
	//	| FALSE 0        | Firewall settings do not allow the traffic specified previously. |
	//	+----------------+------------------------------------------------------------------+
	//	| TRUE 1 — 128   | Firewall settings allow the traffic specified previously.        |
	//	+----------------+------------------------------------------------------------------+
	CommonRulesEnabled bool `idl:"name:commonRulesEnabled" json:"common_rules_enabled"`
	// Return: The GetNextAdapterFirewallConfiguration return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNextAdapterFirewallConfigurationResponse) xxx_ToOp(ctx context.Context) *xxx_GetNextAdapterFirewallConfigurationOperation {
	if o == nil {
		return &xxx_GetNextAdapterFirewallConfigurationOperation{}
	}
	return &xxx_GetNextAdapterFirewallConfigurationOperation{
		That:                   o.That,
		AdapterID:              o.AdapterID,
		AdapterProfile:         o.AdapterProfile,
		ServerRulesEnabled:     o.ServerRulesEnabled,
		ManagementRulesEnabled: o.ManagementRulesEnabled,
		CommonRulesEnabled:     o.CommonRulesEnabled,
		Return:                 o.Return,
	}
}

func (o *GetNextAdapterFirewallConfigurationResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNextAdapterFirewallConfigurationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.AdapterID = op.AdapterID
	o.AdapterProfile = op.AdapterProfile
	o.ServerRulesEnabled = op.ServerRulesEnabled
	o.ManagementRulesEnabled = op.ManagementRulesEnabled
	o.CommonRulesEnabled = op.CommonRulesEnabled
	o.Return = op.Return
}
func (o *GetNextAdapterFirewallConfigurationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetNextAdapterFirewallConfigurationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNextAdapterFirewallConfigurationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
