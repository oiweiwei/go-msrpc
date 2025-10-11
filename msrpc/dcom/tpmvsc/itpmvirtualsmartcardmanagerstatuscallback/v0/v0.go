package itpmvirtualsmartcardmanagerstatuscallback

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
	tpmvsc "github.com/oiweiwei/go-msrpc/msrpc/dcom/tpmvsc"
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
	_ = tpmvsc.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/tpmvsc"
)

var (
	// ITpmVirtualSmartCardManagerStatusCallback interface identifier 1a1bb35f-abb8-451c-a1ae-33d98f1bef4a
	VirtualSmartCardManagerStatusCallbackIID = &dcom.IID{Data1: 0x1a1bb35f, Data2: 0xabb8, Data3: 0x451c, Data4: []byte{0xa1, 0xae, 0x33, 0xd9, 0x8f, 0x1b, 0xef, 0x4a}}
	// Syntax UUID
	VirtualSmartCardManagerStatusCallbackSyntaxUUID = &uuid.UUID{TimeLow: 0x1a1bb35f, TimeMid: 0xabb8, TimeHiAndVersion: 0x451c, ClockSeqHiAndReserved: 0xa1, ClockSeqLow: 0xae, Node: [6]uint8{0x33, 0xd9, 0x8f, 0x1b, 0xef, 0x4a}}
	// Syntax ID
	VirtualSmartCardManagerStatusCallbackSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: VirtualSmartCardManagerStatusCallbackSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// ITpmVirtualSmartCardManagerStatusCallback interface.
type VirtualSmartCardManagerStatusCallbackClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// This method is called by the target to indicate the progress of a TPMVSC management
	// request on the target. The association to a specific ITpmVirtualSmartCardManager
	// method invocation is made by the causality ID in the underlying DCOM transport, as
	// specified in [MS-DCOM] section 3.2.4.2.
	//
	// Return Values: The server MUST return 0 unless it has been instructed to abort the
	// TPMVSC management request as specified in section 3.2.6.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The server SHOULD report the status code to the higher-layer protocol or application
	// that called the associated ITpmVirtualSmartCardManager method.
	ReportProgress(context.Context, *ReportProgressRequest, ...dcerpc.CallOption) (*ReportProgressResponse, error)

	// This method is called by the target to indicate that an error was encountered during
	// the execution of a TPMVSC management request on the target. The association to a
	// specific ITpmVirtualSmartCardManager method invocation is made by the causality ID
	// in the underlying DCOM transport, as specified in [MS-DCOM] section 3.2.4.2.
	//
	// Return Values: The server MUST return 0 unless it has been instructed to abort the
	// TPMVSC management request as specified in section 3.2.6.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The server SHOULD report the error code to the higher-layer protocol or application
	// that called the associated ITpmVirtualSmartCardManager method.
	ReportError(context.Context, *ReportErrorRequest, ...dcerpc.CallOption) (*ReportErrorResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) VirtualSmartCardManagerStatusCallbackClient
}

type xxx_DefaultVirtualSmartCardManagerStatusCallbackClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultVirtualSmartCardManagerStatusCallbackClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultVirtualSmartCardManagerStatusCallbackClient) ReportProgress(ctx context.Context, in *ReportProgressRequest, opts ...dcerpc.CallOption) (*ReportProgressResponse, error) {
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
	out := &ReportProgressResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVirtualSmartCardManagerStatusCallbackClient) ReportError(ctx context.Context, in *ReportErrorRequest, opts ...dcerpc.CallOption) (*ReportErrorResponse, error) {
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
	out := &ReportErrorResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVirtualSmartCardManagerStatusCallbackClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultVirtualSmartCardManagerStatusCallbackClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultVirtualSmartCardManagerStatusCallbackClient) IPID(ctx context.Context, ipid *dcom.IPID) VirtualSmartCardManagerStatusCallbackClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultVirtualSmartCardManagerStatusCallbackClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewVirtualSmartCardManagerStatusCallbackClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (VirtualSmartCardManagerStatusCallbackClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(VirtualSmartCardManagerStatusCallbackSyntaxV0_0))...)
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
	return &xxx_DefaultVirtualSmartCardManagerStatusCallbackClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_ReportProgressOperation structure represents the ReportProgress operation
type xxx_ReportProgressOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Status tpmvsc.Status  `idl:"name:Status" json:"status"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ReportProgressOperation) OpNum() int { return 3 }

func (o *xxx_ReportProgressOperation) OpName() string {
	return "/ITpmVirtualSmartCardManagerStatusCallback/v0/ReportProgress"
}

func (o *xxx_ReportProgressOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportProgressOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Status {in} (1:{v1_enum, alias=TPMVSCMGR_STATUS}(enum))
	{
		if err := w.WriteEnum(uint32(o.Status)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportProgressOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Status {in} (1:{v1_enum, alias=TPMVSCMGR_STATUS}(enum))
	{
		if err := w.ReadEnum((*uint32)(&o.Status)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportProgressOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportProgressOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ReportProgressOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ReportProgressRequest structure represents the ReportProgress operation request
type ReportProgressRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// Status: A TPMVSCMGR_STATUS, defined in section 2.2.1.2.
	Status tpmvsc.Status `idl:"name:Status" json:"status"`
}

func (o *ReportProgressRequest) xxx_ToOp(ctx context.Context, op *xxx_ReportProgressOperation) *xxx_ReportProgressOperation {
	if op == nil {
		op = &xxx_ReportProgressOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Status = o.Status
	return op
}

func (o *ReportProgressRequest) xxx_FromOp(ctx context.Context, op *xxx_ReportProgressOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Status = op.Status
}
func (o *ReportProgressRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ReportProgressRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReportProgressOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReportProgressResponse structure represents the ReportProgress operation response
type ReportProgressResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ReportProgress return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReportProgressResponse) xxx_ToOp(ctx context.Context, op *xxx_ReportProgressOperation) *xxx_ReportProgressOperation {
	if op == nil {
		op = &xxx_ReportProgressOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ReportProgressResponse) xxx_FromOp(ctx context.Context, op *xxx_ReportProgressOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ReportProgressResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ReportProgressResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReportProgressOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReportErrorOperation structure represents the ReportError operation
type xxx_ReportErrorOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Error  tpmvsc.Error   `idl:"name:Error" json:"error"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ReportErrorOperation) OpNum() int { return 4 }

func (o *xxx_ReportErrorOperation) OpName() string {
	return "/ITpmVirtualSmartCardManagerStatusCallback/v0/ReportError"
}

func (o *xxx_ReportErrorOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportErrorOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Error {in} (1:{v1_enum, alias=TPMVSCMGR_ERROR}(enum))
	{
		if err := w.WriteEnum(uint32(o.Error)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportErrorOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Error {in} (1:{v1_enum, alias=TPMVSCMGR_ERROR}(enum))
	{
		if err := w.ReadEnum((*uint32)(&o.Error)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportErrorOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReportErrorOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ReportErrorOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ReportErrorRequest structure represents the ReportError operation request
type ReportErrorRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// Error: A TPMVSCMGR_ERROR, defined in section 2.2.1.1.
	Error tpmvsc.Error `idl:"name:Error" json:"error"`
}

func (o *ReportErrorRequest) xxx_ToOp(ctx context.Context, op *xxx_ReportErrorOperation) *xxx_ReportErrorOperation {
	if op == nil {
		op = &xxx_ReportErrorOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Error = o.Error
	return op
}

func (o *ReportErrorRequest) xxx_FromOp(ctx context.Context, op *xxx_ReportErrorOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Error = op.Error
}
func (o *ReportErrorRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ReportErrorRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReportErrorOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReportErrorResponse structure represents the ReportError operation response
type ReportErrorResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ReportError return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReportErrorResponse) xxx_ToOp(ctx context.Context, op *xxx_ReportErrorOperation) *xxx_ReportErrorOperation {
	if op == nil {
		op = &xxx_ReportErrorOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ReportErrorResponse) xxx_FromOp(ctx context.Context, op *xxx_ReportErrorOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ReportErrorResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ReportErrorResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReportErrorOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
