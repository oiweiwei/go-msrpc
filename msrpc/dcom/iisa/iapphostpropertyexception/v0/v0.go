package iapphostpropertyexception

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iapphostconfigexception "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostconfigexception/v0"
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
	_ = iapphostconfigexception.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/iisa"
)

var (
	// IAppHostPropertyException interface identifier eafe4895-a929-41ea-b14d-613e23f62b71
	AppHostPropertyExceptionIID = &dcom.IID{Data1: 0xeafe4895, Data2: 0xa929, Data3: 0x41ea, Data4: []byte{0xb1, 0x4d, 0x61, 0x3e, 0x23, 0xf6, 0x2b, 0x71}}
	// Syntax UUID
	AppHostPropertyExceptionSyntaxUUID = &uuid.UUID{TimeLow: 0xeafe4895, TimeMid: 0xa929, TimeHiAndVersion: 0x41ea, ClockSeqHiAndReserved: 0xb1, ClockSeqLow: 0x4d, Node: [6]uint8{0x61, 0x3e, 0x23, 0xf6, 0x2b, 0x71}}
	// Syntax ID
	AppHostPropertyExceptionSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: AppHostPropertyExceptionSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IAppHostPropertyException interface.
type AppHostPropertyExceptionClient interface {

	// IAppHostConfigException retrieval method.
	AppHostConfigException() iapphostconfigexception.AppHostConfigExceptionClient

	// InvalidValue operation.
	GetInvalidValue(context.Context, *GetInvalidValueRequest, ...dcerpc.CallOption) (*GetInvalidValueResponse, error)

	// ValidationFailureReason operation.
	GetValidationFailureReason(context.Context, *GetValidationFailureReasonRequest, ...dcerpc.CallOption) (*GetValidationFailureReasonResponse, error)

	// ValidationFailureParameters operation.
	GetValidationFailureParameters(context.Context, *GetValidationFailureParametersRequest, ...dcerpc.CallOption) (*GetValidationFailureParametersResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) AppHostPropertyExceptionClient
}

type xxx_DefaultAppHostPropertyExceptionClient struct {
	iapphostconfigexception.AppHostConfigExceptionClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultAppHostPropertyExceptionClient) AppHostConfigException() iapphostconfigexception.AppHostConfigExceptionClient {
	return o.AppHostConfigExceptionClient
}

func (o *xxx_DefaultAppHostPropertyExceptionClient) GetInvalidValue(ctx context.Context, in *GetInvalidValueRequest, opts ...dcerpc.CallOption) (*GetInvalidValueResponse, error) {
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
	out := &GetInvalidValueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostPropertyExceptionClient) GetValidationFailureReason(ctx context.Context, in *GetValidationFailureReasonRequest, opts ...dcerpc.CallOption) (*GetValidationFailureReasonResponse, error) {
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
	out := &GetValidationFailureReasonResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostPropertyExceptionClient) GetValidationFailureParameters(ctx context.Context, in *GetValidationFailureParametersRequest, opts ...dcerpc.CallOption) (*GetValidationFailureParametersResponse, error) {
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
	out := &GetValidationFailureParametersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostPropertyExceptionClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultAppHostPropertyExceptionClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultAppHostPropertyExceptionClient) IPID(ctx context.Context, ipid *dcom.IPID) AppHostPropertyExceptionClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultAppHostPropertyExceptionClient{
		AppHostConfigExceptionClient: o.AppHostConfigExceptionClient.IPID(ctx, ipid),
		cc:                           o.cc,
		ipid:                         ipid,
	}
}

func NewAppHostPropertyExceptionClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (AppHostPropertyExceptionClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(AppHostPropertyExceptionSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iapphostconfigexception.NewAppHostConfigExceptionClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultAppHostPropertyExceptionClient{
		AppHostConfigExceptionClient: base,
		cc:                           cc,
		ipid:                         ipid,
	}, nil
}

// xxx_GetInvalidValueOperation structure represents the InvalidValue operation
type xxx_GetInvalidValueOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Value  *oaut.String   `idl:"name:pbstrValue" json:"value"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetInvalidValueOperation) OpNum() int { return 10 }

func (o *xxx_GetInvalidValueOperation) OpName() string {
	return "/IAppHostPropertyException/v0/InvalidValue"
}

func (o *xxx_GetInvalidValueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInvalidValueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetInvalidValueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetInvalidValueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInvalidValueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrValue {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Value != nil {
			_ptr_pbstrValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Value != nil {
					if err := o.Value.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Value, _ptr_pbstrValue); err != nil {
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

func (o *xxx_GetInvalidValueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrValue {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Value == nil {
				o.Value = &oaut.String{}
			}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrValue := func(ptr interface{}) { o.Value = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Value, _s_pbstrValue, _ptr_pbstrValue); err != nil {
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

// GetInvalidValueRequest structure represents the InvalidValue operation request
type GetInvalidValueRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetInvalidValueRequest) xxx_ToOp(ctx context.Context, op *xxx_GetInvalidValueOperation) *xxx_GetInvalidValueOperation {
	if op == nil {
		op = &xxx_GetInvalidValueOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *GetInvalidValueRequest) xxx_FromOp(ctx context.Context, op *xxx_GetInvalidValueOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetInvalidValueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetInvalidValueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetInvalidValueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetInvalidValueResponse structure represents the InvalidValue operation response
type GetInvalidValueResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat `idl:"name:That" json:"that"`
	Value *oaut.String   `idl:"name:pbstrValue" json:"value"`
	// Return: The InvalidValue return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetInvalidValueResponse) xxx_ToOp(ctx context.Context, op *xxx_GetInvalidValueOperation) *xxx_GetInvalidValueOperation {
	if op == nil {
		op = &xxx_GetInvalidValueOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Value = op.Value
	o.Return = op.Return
	return op
}

func (o *GetInvalidValueResponse) xxx_FromOp(ctx context.Context, op *xxx_GetInvalidValueOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Value = op.Value
	o.Return = op.Return
}
func (o *GetInvalidValueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetInvalidValueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetInvalidValueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetValidationFailureReasonOperation structure represents the ValidationFailureReason operation
type xxx_GetValidationFailureReasonOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	ValidationReason *oaut.String   `idl:"name:pbstrValidationReason" json:"validation_reason"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetValidationFailureReasonOperation) OpNum() int { return 11 }

func (o *xxx_GetValidationFailureReasonOperation) OpName() string {
	return "/IAppHostPropertyException/v0/ValidationFailureReason"
}

func (o *xxx_GetValidationFailureReasonOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetValidationFailureReasonOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetValidationFailureReasonOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetValidationFailureReasonOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetValidationFailureReasonOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrValidationReason {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ValidationReason != nil {
			_ptr_pbstrValidationReason := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ValidationReason != nil {
					if err := o.ValidationReason.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ValidationReason, _ptr_pbstrValidationReason); err != nil {
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

func (o *xxx_GetValidationFailureReasonOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrValidationReason {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrValidationReason := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ValidationReason == nil {
				o.ValidationReason = &oaut.String{}
			}
			if err := o.ValidationReason.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrValidationReason := func(ptr interface{}) { o.ValidationReason = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ValidationReason, _s_pbstrValidationReason, _ptr_pbstrValidationReason); err != nil {
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

// GetValidationFailureReasonRequest structure represents the ValidationFailureReason operation request
type GetValidationFailureReasonRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetValidationFailureReasonRequest) xxx_ToOp(ctx context.Context, op *xxx_GetValidationFailureReasonOperation) *xxx_GetValidationFailureReasonOperation {
	if op == nil {
		op = &xxx_GetValidationFailureReasonOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *GetValidationFailureReasonRequest) xxx_FromOp(ctx context.Context, op *xxx_GetValidationFailureReasonOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetValidationFailureReasonRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetValidationFailureReasonRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetValidationFailureReasonOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetValidationFailureReasonResponse structure represents the ValidationFailureReason operation response
type GetValidationFailureReasonResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	ValidationReason *oaut.String   `idl:"name:pbstrValidationReason" json:"validation_reason"`
	// Return: The ValidationFailureReason return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetValidationFailureReasonResponse) xxx_ToOp(ctx context.Context, op *xxx_GetValidationFailureReasonOperation) *xxx_GetValidationFailureReasonOperation {
	if op == nil {
		op = &xxx_GetValidationFailureReasonOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.ValidationReason = op.ValidationReason
	o.Return = op.Return
	return op
}

func (o *GetValidationFailureReasonResponse) xxx_FromOp(ctx context.Context, op *xxx_GetValidationFailureReasonOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ValidationReason = op.ValidationReason
	o.Return = op.Return
}
func (o *GetValidationFailureReasonResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetValidationFailureReasonResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetValidationFailureReasonOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetValidationFailureParametersOperation structure represents the ValidationFailureParameters operation
type xxx_GetValidationFailureParametersOperation struct {
	This           *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ParameterArray *oaut.SafeArray `idl:"name:pParameterArray" json:"parameter_array"`
	Return         int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetValidationFailureParametersOperation) OpNum() int { return 12 }

func (o *xxx_GetValidationFailureParametersOperation) OpName() string {
	return "/IAppHostPropertyException/v0/ValidationFailureParameters"
}

func (o *xxx_GetValidationFailureParametersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetValidationFailureParametersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetValidationFailureParametersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetValidationFailureParametersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetValidationFailureParametersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pParameterArray {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.ParameterArray != nil {
			_ptr_pParameterArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ParameterArray != nil {
					if err := o.ParameterArray.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ParameterArray, _ptr_pParameterArray); err != nil {
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

func (o *xxx_GetValidationFailureParametersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pParameterArray {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_pParameterArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ParameterArray == nil {
				o.ParameterArray = &oaut.SafeArray{}
			}
			if err := o.ParameterArray.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pParameterArray := func(ptr interface{}) { o.ParameterArray = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.ParameterArray, _s_pParameterArray, _ptr_pParameterArray); err != nil {
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

// GetValidationFailureParametersRequest structure represents the ValidationFailureParameters operation request
type GetValidationFailureParametersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetValidationFailureParametersRequest) xxx_ToOp(ctx context.Context, op *xxx_GetValidationFailureParametersOperation) *xxx_GetValidationFailureParametersOperation {
	if op == nil {
		op = &xxx_GetValidationFailureParametersOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *GetValidationFailureParametersRequest) xxx_FromOp(ctx context.Context, op *xxx_GetValidationFailureParametersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetValidationFailureParametersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetValidationFailureParametersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetValidationFailureParametersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetValidationFailureParametersResponse structure represents the ValidationFailureParameters operation response
type GetValidationFailureParametersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That           *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ParameterArray *oaut.SafeArray `idl:"name:pParameterArray" json:"parameter_array"`
	// Return: The ValidationFailureParameters return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetValidationFailureParametersResponse) xxx_ToOp(ctx context.Context, op *xxx_GetValidationFailureParametersOperation) *xxx_GetValidationFailureParametersOperation {
	if op == nil {
		op = &xxx_GetValidationFailureParametersOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.ParameterArray = op.ParameterArray
	o.Return = op.Return
	return op
}

func (o *GetValidationFailureParametersResponse) xxx_FromOp(ctx context.Context, op *xxx_GetValidationFailureParametersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ParameterArray = op.ParameterArray
	o.Return = op.Return
}
func (o *GetValidationFailureParametersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetValidationFailureParametersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetValidationFailureParametersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
