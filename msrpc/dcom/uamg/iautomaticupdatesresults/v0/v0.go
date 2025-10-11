package iautomaticupdatesresults

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
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
	_ = idispatch.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/uamg"
)

var (
	// IAutomaticUpdatesResults interface identifier e7a4d634-7942-4dd9-a111-82228ba33901
	AutomaticUpdatesResultsIID = &dcom.IID{Data1: 0xe7a4d634, Data2: 0x7942, Data3: 0x4dd9, Data4: []byte{0xa1, 0x11, 0x82, 0x22, 0x8b, 0xa3, 0x39, 0x01}}
	// Syntax UUID
	AutomaticUpdatesResultsSyntaxUUID = &uuid.UUID{TimeLow: 0xe7a4d634, TimeMid: 0x7942, TimeHiAndVersion: 0x4dd9, ClockSeqHiAndReserved: 0xa1, ClockSeqLow: 0x11, Node: [6]uint8{0x82, 0x22, 0x8b, 0xa3, 0x39, 0x1}}
	// Syntax ID
	AutomaticUpdatesResultsSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: AutomaticUpdatesResultsSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IAutomaticUpdatesResults interface.
type AutomaticUpdatesResultsClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// The IAutomaticUpdatesResults::LastSearchSuccessDate (opnum 8) method retrieves the
	// date-time of the most recent successful search done by the automatic update agent.
	//
	// Return Values: The method MUST return information in an HRESULT  data structure.
	// The severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return a VARIANT, as specified in [MS-OAUT] section 2.2.29.2,
	// containing the value of the LastSearchSuccessDate ADM element.
	GetLastSearchSuccessDate(context.Context, *GetLastSearchSuccessDateRequest, ...dcerpc.CallOption) (*GetLastSearchSuccessDateResponse, error)

	// The IAutomaticUpdatesResults::LastInstallationSuccessDate (opnum 9) method retrieves
	// the date-time of the most recent successful installation done by the automatic update
	// agent.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return a VARIANT containing the value of the LastInstallationSuccessDate
	// ADM element.
	GetLastInstallationSuccessDate(context.Context, *GetLastInstallationSuccessDateRequest, ...dcerpc.CallOption) (*GetLastInstallationSuccessDateResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) AutomaticUpdatesResultsClient
}

type xxx_DefaultAutomaticUpdatesResultsClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultAutomaticUpdatesResultsClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultAutomaticUpdatesResultsClient) GetLastSearchSuccessDate(ctx context.Context, in *GetLastSearchSuccessDateRequest, opts ...dcerpc.CallOption) (*GetLastSearchSuccessDateResponse, error) {
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
	out := &GetLastSearchSuccessDateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAutomaticUpdatesResultsClient) GetLastInstallationSuccessDate(ctx context.Context, in *GetLastInstallationSuccessDateRequest, opts ...dcerpc.CallOption) (*GetLastInstallationSuccessDateResponse, error) {
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
	out := &GetLastInstallationSuccessDateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAutomaticUpdatesResultsClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultAutomaticUpdatesResultsClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultAutomaticUpdatesResultsClient) IPID(ctx context.Context, ipid *dcom.IPID) AutomaticUpdatesResultsClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultAutomaticUpdatesResultsClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewAutomaticUpdatesResultsClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (AutomaticUpdatesResultsClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(AutomaticUpdatesResultsSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := idispatch.NewDispatchClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultAutomaticUpdatesResultsClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetLastSearchSuccessDateOperation structure represents the LastSearchSuccessDate operation
type xxx_GetLastSearchSuccessDateOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReturnValue *oaut.Variant  `idl:"name:retval" json:"return_value"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetLastSearchSuccessDateOperation) OpNum() int { return 7 }

func (o *xxx_GetLastSearchSuccessDateOperation) OpName() string {
	return "/IAutomaticUpdatesResults/v0/LastSearchSuccessDate"
}

func (o *xxx_GetLastSearchSuccessDateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLastSearchSuccessDateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetLastSearchSuccessDateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetLastSearchSuccessDateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLastSearchSuccessDateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// retval {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.ReturnValue != nil {
			_ptr_retval := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReturnValue != nil {
					if err := o.ReturnValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReturnValue, _ptr_retval); err != nil {
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

func (o *xxx_GetLastSearchSuccessDateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// retval {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_retval := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReturnValue == nil {
				o.ReturnValue = &oaut.Variant{}
			}
			if err := o.ReturnValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_retval := func(ptr interface{}) { o.ReturnValue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.ReturnValue, _s_retval, _ptr_retval); err != nil {
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

// GetLastSearchSuccessDateRequest structure represents the LastSearchSuccessDate operation request
type GetLastSearchSuccessDateRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetLastSearchSuccessDateRequest) xxx_ToOp(ctx context.Context, op *xxx_GetLastSearchSuccessDateOperation) *xxx_GetLastSearchSuccessDateOperation {
	if op == nil {
		op = &xxx_GetLastSearchSuccessDateOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetLastSearchSuccessDateRequest) xxx_FromOp(ctx context.Context, op *xxx_GetLastSearchSuccessDateOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetLastSearchSuccessDateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetLastSearchSuccessDateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLastSearchSuccessDateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetLastSearchSuccessDateResponse structure represents the LastSearchSuccessDate operation response
type GetLastSearchSuccessDateResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// retval: If the date-time for the last successful update search is not known, the
	// vt field (see [MS-OAUT] section 2.2.29.1), MUST be set to VT_EMPTY as specified in
	// [MS-OAUT] section 2.2.7. Otherwise, the vt member MUST be set to VT_DATE (see [MS-OAUT]
	// section 2.2.7, and the date-time of the last successful update search MUST be stored
	// in the date field.
	ReturnValue *oaut.Variant `idl:"name:retval" json:"return_value"`
	// Return: The LastSearchSuccessDate return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetLastSearchSuccessDateResponse) xxx_ToOp(ctx context.Context, op *xxx_GetLastSearchSuccessDateOperation) *xxx_GetLastSearchSuccessDateOperation {
	if op == nil {
		op = &xxx_GetLastSearchSuccessDateOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *GetLastSearchSuccessDateResponse) xxx_FromOp(ctx context.Context, op *xxx_GetLastSearchSuccessDateOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *GetLastSearchSuccessDateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetLastSearchSuccessDateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLastSearchSuccessDateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetLastInstallationSuccessDateOperation structure represents the LastInstallationSuccessDate operation
type xxx_GetLastInstallationSuccessDateOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReturnValue *oaut.Variant  `idl:"name:retval" json:"return_value"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetLastInstallationSuccessDateOperation) OpNum() int { return 8 }

func (o *xxx_GetLastInstallationSuccessDateOperation) OpName() string {
	return "/IAutomaticUpdatesResults/v0/LastInstallationSuccessDate"
}

func (o *xxx_GetLastInstallationSuccessDateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLastInstallationSuccessDateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetLastInstallationSuccessDateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetLastInstallationSuccessDateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLastInstallationSuccessDateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// retval {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.ReturnValue != nil {
			_ptr_retval := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReturnValue != nil {
					if err := o.ReturnValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReturnValue, _ptr_retval); err != nil {
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

func (o *xxx_GetLastInstallationSuccessDateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// retval {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_retval := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReturnValue == nil {
				o.ReturnValue = &oaut.Variant{}
			}
			if err := o.ReturnValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_retval := func(ptr interface{}) { o.ReturnValue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.ReturnValue, _s_retval, _ptr_retval); err != nil {
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

// GetLastInstallationSuccessDateRequest structure represents the LastInstallationSuccessDate operation request
type GetLastInstallationSuccessDateRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetLastInstallationSuccessDateRequest) xxx_ToOp(ctx context.Context, op *xxx_GetLastInstallationSuccessDateOperation) *xxx_GetLastInstallationSuccessDateOperation {
	if op == nil {
		op = &xxx_GetLastInstallationSuccessDateOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetLastInstallationSuccessDateRequest) xxx_FromOp(ctx context.Context, op *xxx_GetLastInstallationSuccessDateOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetLastInstallationSuccessDateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetLastInstallationSuccessDateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLastInstallationSuccessDateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetLastInstallationSuccessDateResponse structure represents the LastInstallationSuccessDate operation response
type GetLastInstallationSuccessDateResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// retval: If the date-time for the last successful update search is not known, the
	// vt field MUST be set to VT_EMPTY. Otherwise, the vt member MUST be set to VT_DATE,
	// and the date-time of the last successful update installation MUST be stored in the
	// date field.
	ReturnValue *oaut.Variant `idl:"name:retval" json:"return_value"`
	// Return: The LastInstallationSuccessDate return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetLastInstallationSuccessDateResponse) xxx_ToOp(ctx context.Context, op *xxx_GetLastInstallationSuccessDateOperation) *xxx_GetLastInstallationSuccessDateOperation {
	if op == nil {
		op = &xxx_GetLastInstallationSuccessDateOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *GetLastInstallationSuccessDateResponse) xxx_FromOp(ctx context.Context, op *xxx_GetLastInstallationSuccessDateOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *GetLastInstallationSuccessDateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetLastInstallationSuccessDateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLastInstallationSuccessDateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
