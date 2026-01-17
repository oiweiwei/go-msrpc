package icatalogsession

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/coma"
)

var (
	// ICatalogSession interface identifier 182c40fa-32e4-11d0-818b-00a0c9231c29
	CatalogSessionIID = &dcom.IID{Data1: 0x182c40fa, Data2: 0x32e4, Data3: 0x11d0, Data4: []byte{0x81, 0x8b, 0x00, 0xa0, 0xc9, 0x23, 0x1c, 0x29}}
	// Syntax UUID
	CatalogSessionSyntaxUUID = &uuid.UUID{TimeLow: 0x182c40fa, TimeMid: 0x32e4, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0x81, ClockSeqLow: 0x8b, Node: [6]uint8{0x0, 0xa0, 0xc9, 0x23, 0x1c, 0x29}}
	// Syntax ID
	CatalogSessionSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: CatalogSessionSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// ICatalogSession interface.
type CatalogSessionClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// Opnum3NotUsedOnWire operation.
	// Opnum3NotUsedOnWire

	// Opnum4NotUsedOnWire operation.
	// Opnum4NotUsedOnWire

	// Opnum5NotUsedOnWire operation.
	// Opnum5NotUsedOnWire

	// Opnum6NotUsedOnWire operation.
	// Opnum6NotUsedOnWire

	// This method is called by a client to perform Catalog Version Negotiation (section
	// 3.1.4.1).
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success and a failure
	// result, as specified in [MS-ERREF] section 2.1 on failure. All failure results MUST
	// be treated identically.
	InitializeSession(context.Context, *InitializeSessionRequest, ...dcerpc.CallOption) (*InitializeSessionResponse, error)

	// This method is called by a client to perform capability negotiation for the Multiple-partition
	// Support Capability Negotiation (section 3.1.4.3).
	//
	// Return Values:  This method MUST return S_OK (0x00000000) on success and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	//
	// A server that does not support catalog version 4.00 or catalog version 5.00 SHOULD
	// immediately return E_NOTIMPL (0x80004001) instead of implementing this method.
	//
	// Otherwise, the server MUST attempt to set the value referenced by plMultiplePartitionSupport
	// to the previously specified value that indicates its support of multiple partitions,
	// and fail the call if it cannot set the value.
	GetServerInformation(context.Context, *GetServerInformationRequest, ...dcerpc.CallOption) (*GetServerInformationResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) CatalogSessionClient
}

type xxx_DefaultCatalogSessionClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultCatalogSessionClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultCatalogSessionClient) InitializeSession(ctx context.Context, in *InitializeSessionRequest, opts ...dcerpc.CallOption) (*InitializeSessionResponse, error) {
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
	out := &InitializeSessionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCatalogSessionClient) GetServerInformation(ctx context.Context, in *GetServerInformationRequest, opts ...dcerpc.CallOption) (*GetServerInformationResponse, error) {
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
	out := &GetServerInformationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCatalogSessionClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultCatalogSessionClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultCatalogSessionClient) IPID(ctx context.Context, ipid *dcom.IPID) CatalogSessionClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultCatalogSessionClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewCatalogSessionClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (CatalogSessionClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(CatalogSessionSyntaxV0_0))...)
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
	return &xxx_DefaultCatalogSessionClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_InitializeSessionOperation structure represents the InitializeSession operation
type xxx_InitializeSessionOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	VerLower   float32        `idl:"name:flVerLower" json:"ver_lower"`
	VerUpper   float32        `idl:"name:flVerUpper" json:"ver_upper"`
	_          int32          `idl:"name:reserved"`
	VerSession float32        `idl:"name:pflVerSession" json:"ver_session"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_InitializeSessionOperation) OpNum() int { return 7 }

func (o *xxx_InitializeSessionOperation) OpName() string {
	return "/ICatalogSession/v0/InitializeSession"
}

func (o *xxx_InitializeSessionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeSessionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// flVerLower {in} (1:(float32))
	{
		if err := w.WriteData(o.VerLower); err != nil {
			return err
		}
	}
	// flVerUpper {in} (1:(float32))
	{
		if err := w.WriteData(o.VerUpper); err != nil {
			return err
		}
	}
	// reserved {in} (1:(int32))
	{
		// reserved reserved
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeSessionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// flVerLower {in} (1:(float32))
	{
		if err := w.ReadData(&o.VerLower); err != nil {
			return err
		}
	}
	// flVerUpper {in} (1:(float32))
	{
		if err := w.ReadData(&o.VerUpper); err != nil {
			return err
		}
	}
	// reserved {in} (1:(int32))
	{
		// reserved reserved
		var _reserved int32
		if err := w.ReadData(&_reserved); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeSessionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeSessionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pflVerSession {out} (1:{pointer=ref}*(1)(float32))
	{
		if err := w.WriteData(o.VerSession); err != nil {
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

func (o *xxx_InitializeSessionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pflVerSession {out} (1:{pointer=ref}*(1)(float32))
	{
		if err := w.ReadData(&o.VerSession); err != nil {
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

// InitializeSessionRequest structure represents the InitializeSession operation request
type InitializeSessionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// flVerLower:  The lowest catalog version supported by the client.
	VerLower float32 `idl:"name:flVerLower" json:"ver_lower"`
	// flVerUpper:  The highest catalog version supported by the client.
	VerUpper float32 `idl:"name:flVerUpper" json:"ver_upper"`
}

func (o *InitializeSessionRequest) xxx_ToOp(ctx context.Context, op *xxx_InitializeSessionOperation) *xxx_InitializeSessionOperation {
	if op == nil {
		op = &xxx_InitializeSessionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.VerLower = o.VerLower
	op.VerUpper = o.VerUpper
	return op
}

func (o *InitializeSessionRequest) xxx_FromOp(ctx context.Context, op *xxx_InitializeSessionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VerLower = op.VerLower
	o.VerUpper = op.VerUpper
}
func (o *InitializeSessionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *InitializeSessionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InitializeSessionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// InitializeSessionResponse structure represents the InitializeSession operation response
type InitializeSessionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pflVerSession: A pointer to a variable that, upon successful completion, MUST be
	// set to the negotiated catalog version.
	VerSession float32 `idl:"name:pflVerSession" json:"ver_session"`
	// Return: The InitializeSession return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *InitializeSessionResponse) xxx_ToOp(ctx context.Context, op *xxx_InitializeSessionOperation) *xxx_InitializeSessionOperation {
	if op == nil {
		op = &xxx_InitializeSessionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.VerSession = o.VerSession
	op.Return = o.Return
	return op
}

func (o *InitializeSessionResponse) xxx_FromOp(ctx context.Context, op *xxx_InitializeSessionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.VerSession = op.VerSession
	o.Return = op.Return
}
func (o *InitializeSessionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *InitializeSessionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InitializeSessionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetServerInformationOperation structure represents the GetServerInformation operation
type xxx_GetServerInformationOperation struct {
	This                     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                     *dcom.ORPCThat `idl:"name:That" json:"that"`
	_                        int32          `idl:"name:plReserved1"`
	_                        int32          `idl:"name:plReserved2"`
	_                        int32          `idl:"name:plReserved3"`
	MultiplePartitionSupport int32          `idl:"name:plMultiplePartitionSupport" json:"multiple_partition_support"`
	_                        int32          `idl:"name:plReserved4"`
	_                        int32          `idl:"name:plReserved5"`
	Return                   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetServerInformationOperation) OpNum() int { return 8 }

func (o *xxx_GetServerInformationOperation) OpName() string {
	return "/ICatalogSession/v0/GetServerInformation"
}

func (o *xxx_GetServerInformationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerInformationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetServerInformationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetServerInformationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerInformationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plReserved1 {out} (1:{pointer=ref}*(1)(int32))
	{
		// reserved plReserved1
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	}
	// plReserved2 {out} (1:{pointer=ref}*(1)(int32))
	{
		// reserved plReserved2
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	}
	// plReserved3 {out} (1:{pointer=ref}*(1)(int32))
	{
		// reserved plReserved3
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	}
	// plMultiplePartitionSupport {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.MultiplePartitionSupport); err != nil {
			return err
		}
	}
	// plReserved4 {out} (1:{pointer=ref}*(1)(int32))
	{
		// reserved plReserved4
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	}
	// plReserved5 {out} (1:{pointer=ref}*(1)(int32))
	{
		// reserved plReserved5
		if err := w.WriteData(int32(0)); err != nil {
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

func (o *xxx_GetServerInformationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plReserved1 {out} (1:{pointer=ref}*(1)(int32))
	{
		// reserved plReserved1
		var _plReserved1 int32
		if err := w.ReadData(&_plReserved1); err != nil {
			return err
		}
	}
	// plReserved2 {out} (1:{pointer=ref}*(1)(int32))
	{
		// reserved plReserved2
		var _plReserved2 int32
		if err := w.ReadData(&_plReserved2); err != nil {
			return err
		}
	}
	// plReserved3 {out} (1:{pointer=ref}*(1)(int32))
	{
		// reserved plReserved3
		var _plReserved3 int32
		if err := w.ReadData(&_plReserved3); err != nil {
			return err
		}
	}
	// plMultiplePartitionSupport {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.MultiplePartitionSupport); err != nil {
			return err
		}
	}
	// plReserved4 {out} (1:{pointer=ref}*(1)(int32))
	{
		// reserved plReserved4
		var _plReserved4 int32
		if err := w.ReadData(&_plReserved4); err != nil {
			return err
		}
	}
	// plReserved5 {out} (1:{pointer=ref}*(1)(int32))
	{
		// reserved plReserved5
		var _plReserved5 int32
		if err := w.ReadData(&_plReserved5); err != nil {
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

// GetServerInformationRequest structure represents the GetServerInformation operation request
type GetServerInformationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetServerInformationRequest) xxx_ToOp(ctx context.Context, op *xxx_GetServerInformationOperation) *xxx_GetServerInformationOperation {
	if op == nil {
		op = &xxx_GetServerInformationOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetServerInformationRequest) xxx_FromOp(ctx context.Context, op *xxx_GetServerInformationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetServerInformationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetServerInformationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetServerInformationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetServerInformationResponse structure represents the GetServerInformation operation response
type GetServerInformationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// plMultiplePartitionSupport: A pointer to a value that, upon successful completion,
	// MUST be set to one of the following values indicating support of multiple partitions.
	//
	//	+-------+----------------------------------------------------------------------------------+
	//	|       |                                                                                  |
	//	| VALUE |                                     MEANING                                      |
	//	|       |                                                                                  |
	//	+-------+----------------------------------------------------------------------------------+
	//	+-------+----------------------------------------------------------------------------------+
	//	|     1 | The server does not support multiple partitions.                                 |
	//	+-------+----------------------------------------------------------------------------------+
	//	|     2 | The server supports multiple partitions.                                         |
	//	+-------+----------------------------------------------------------------------------------+
	//	|     3 | The server supports multiple partitions and is also capable of managing the      |
	//	|       | domain-controlled PartitionRoles (section 3.1.1.3.17), PartitionRoleMembers      |
	//	|       | (section 3.1.1.3.18), and PartitionUsers (section 3.1.1.3.16) tables for other   |
	//	|       | servers. This value SHOULD be treated the same as 2, because it does not affect  |
	//	|       | interoperability.                                                                |
	//	+-------+----------------------------------------------------------------------------------+
	MultiplePartitionSupport int32 `idl:"name:plMultiplePartitionSupport" json:"multiple_partition_support"`
	// Return: The GetServerInformation return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetServerInformationResponse) xxx_ToOp(ctx context.Context, op *xxx_GetServerInformationOperation) *xxx_GetServerInformationOperation {
	if op == nil {
		op = &xxx_GetServerInformationOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.MultiplePartitionSupport = o.MultiplePartitionSupport
	op.Return = o.Return
	return op
}

func (o *GetServerInformationResponse) xxx_FromOp(ctx context.Context, op *xxx_GetServerInformationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MultiplePartitionSupport = op.MultiplePartitionSupport
	o.Return = op.Return
}
func (o *GetServerInformationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetServerInformationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetServerInformationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
