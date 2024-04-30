package iwbemcallresult

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
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
	wmi "github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi"
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
	_ = wmi.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wmi"
)

var (
	// IWbemCallResult interface identifier 44aca675-e8fc-11d0-a07c-00c04fb68820
	CallResultIID = &dcom.IID{Data1: 0x44aca675, Data2: 0xe8fc, Data3: 0x11d0, Data4: []byte{0xa0, 0x7c, 0x00, 0xc0, 0x4f, 0xb6, 0x88, 0x20}}
	// Syntax UUID
	CallResultSyntaxUUID = &uuid.UUID{TimeLow: 0x44aca675, TimeMid: 0xe8fc, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xa0, ClockSeqLow: 0x7c, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xb6, 0x88, 0x20}}
	// Syntax ID
	CallResultSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: CallResultSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IWbemCallResult interface.
type CallResultClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// When the IWbemCallResult::GetResultObject method is called, the server MUST attempt
	// to retrieve a CIM object from a previous semisynchronous operation call to the IWbemServices::GetObject
	// method or the IWbemServices::ExecMethod method. The entry in WbemCallResultTable
	// with WbemCallResultPointer pointing to IWbemCallResult is used to identify the previous
	// semisynchronous call.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	GetResultObject(context.Context, *GetResultObjectRequest, ...dcerpc.CallOption) (*GetResultObjectResponse, error)

	// When the IWbemCallResult::GetResultString method is called, the server MUST return
	// the assigned CIM path of a CIM instance that was created by the IWbemServices::PutInstance
	// method that returned IWbemCallResult in the ppCallResult parameter.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	GetResultString(context.Context, *GetResultStringRequest, ...dcerpc.CallOption) (*GetResultStringResponse, error)

	// When the IWbemCallResult::GetResultServices method is called, the server MUST retrieve
	// a pointer to the IWbemServices interface that results from a semisynchronous call
	// to the IWbemServices::OpenNamespace method.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	GetResultServices(context.Context, *GetResultServicesRequest, ...dcerpc.CallOption) (*GetResultServicesResponse, error)

	// When the IWbemCallResult::GetCallStatus method is invoked, the server MUST return
	// the status of the current outstanding semisynchronous call.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	GetCallStatus(context.Context, *GetCallStatusRequest, ...dcerpc.CallOption) (*GetCallStatusResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) CallResultClient
}

type xxx_DefaultCallResultClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultCallResultClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultCallResultClient) GetResultObject(ctx context.Context, in *GetResultObjectRequest, opts ...dcerpc.CallOption) (*GetResultObjectResponse, error) {
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
	out := &GetResultObjectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCallResultClient) GetResultString(ctx context.Context, in *GetResultStringRequest, opts ...dcerpc.CallOption) (*GetResultStringResponse, error) {
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
	out := &GetResultStringResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCallResultClient) GetResultServices(ctx context.Context, in *GetResultServicesRequest, opts ...dcerpc.CallOption) (*GetResultServicesResponse, error) {
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
	out := &GetResultServicesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCallResultClient) GetCallStatus(ctx context.Context, in *GetCallStatusRequest, opts ...dcerpc.CallOption) (*GetCallStatusResponse, error) {
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
	out := &GetCallStatusResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCallResultClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultCallResultClient) IPID(ctx context.Context, ipid *dcom.IPID) CallResultClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultCallResultClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}
func NewCallResultClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (CallResultClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(CallResultSyntaxV0_0))...)
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
	return &xxx_DefaultCallResultClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetResultObjectOperation structure represents the GetResultObject operation
type xxx_GetResultObjectOperation struct {
	This         *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Timeout      int32            `idl:"name:lTimeout" json:"timeout"`
	ResultObject *wmi.ClassObject `idl:"name:ppResultObject" json:"result_object"`
	Return       int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetResultObjectOperation) OpNum() int { return 3 }

func (o *xxx_GetResultObjectOperation) OpName() string { return "/IWbemCallResult/v0/GetResultObject" }

func (o *xxx_GetResultObjectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResultObjectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lTimeout {in} (1:(int32))
	{
		if err := w.WriteData(o.Timeout); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResultObjectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lTimeout {in} (1:(int32))
	{
		if err := w.ReadData(&o.Timeout); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResultObjectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResultObjectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppResultObject {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IWbemClassObject}(interface))
	{
		if o.ResultObject != nil {
			_ptr_ppResultObject := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ResultObject != nil {
					if err := o.ResultObject.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.ClassObject{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResultObject, _ptr_ppResultObject); err != nil {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResultObjectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppResultObject {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IWbemClassObject}(interface))
	{
		_ptr_ppResultObject := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ResultObject == nil {
				o.ResultObject = &wmi.ClassObject{}
			}
			if err := o.ResultObject.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppResultObject := func(ptr interface{}) { o.ResultObject = *ptr.(**wmi.ClassObject) }
		if err := w.ReadPointer(&o.ResultObject, _s_ppResultObject, _ptr_ppResultObject); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetResultObjectRequest structure represents the GetResultObject operation request
type GetResultObjectRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lTimeout: MUST be the maximum amount of time, in milliseconds, that the call to the
	// IWbemCallResult::GetResultObject method allows to pass before it times out. If the
	// constant WBEM_INFINITE (0xFFFFFFFF) is used, the GetResultObject method call MUST
	// wait until the operation succeeds. If this parameter is set to 0 and the result object
	// is available at the time of the method call, the object MUST be returned in ppResultObject
	// and WBEM_S_NO_ERROR MUST also be returned. If this parameter is set to 0 but the
	// result object is not available at the time of the method call, WBEM_S_TIMEDOUT MUST
	// be returned.
	Timeout int32 `idl:"name:lTimeout" json:"timeout"`
}

func (o *GetResultObjectRequest) xxx_ToOp(ctx context.Context) *xxx_GetResultObjectOperation {
	if o == nil {
		return &xxx_GetResultObjectOperation{}
	}
	return &xxx_GetResultObjectOperation{
		This:    o.This,
		Timeout: o.Timeout,
	}
}

func (o *GetResultObjectRequest) xxx_FromOp(ctx context.Context, op *xxx_GetResultObjectOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Timeout = op.Timeout
}
func (o *GetResultObjectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetResultObjectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResultObjectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetResultObjectResponse structure represents the GetResultObject operation response
type GetResultObjectResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppResultObject: A pointer to a variable that receives a logical copy of the CIM object
	// when the semisynchronous operation is complete. A new CIM object MUST NOT be returned
	// on error. When sent by the client, this parameter value MUST NOT be NULL. Upon return
	// by the server, this parameter value can be NULL if there is a failure or if there
	// are no results. The caller of this method MUST call IWbemClassObject::Release on
	// the returned object when the object is no longer required.
	ResultObject *wmi.ClassObject `idl:"name:ppResultObject" json:"result_object"`
	// Return: The GetResultObject return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetResultObjectResponse) xxx_ToOp(ctx context.Context) *xxx_GetResultObjectOperation {
	if o == nil {
		return &xxx_GetResultObjectOperation{}
	}
	return &xxx_GetResultObjectOperation{
		That:         o.That,
		ResultObject: o.ResultObject,
		Return:       o.Return,
	}
}

func (o *GetResultObjectResponse) xxx_FromOp(ctx context.Context, op *xxx_GetResultObjectOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ResultObject = op.ResultObject
	o.Return = op.Return
}
func (o *GetResultObjectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetResultObjectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResultObjectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetResultStringOperation structure represents the GetResultString operation
type xxx_GetResultStringOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	Timeout      int32          `idl:"name:lTimeout" json:"timeout"`
	ResultString *oaut.String   `idl:"name:pstrResultString" json:"result_string"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetResultStringOperation) OpNum() int { return 4 }

func (o *xxx_GetResultStringOperation) OpName() string { return "/IWbemCallResult/v0/GetResultString" }

func (o *xxx_GetResultStringOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResultStringOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lTimeout {in} (1:(int32))
	{
		if err := w.WriteData(o.Timeout); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResultStringOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lTimeout {in} (1:(int32))
	{
		if err := w.ReadData(&o.Timeout); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResultStringOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResultStringOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pstrResultString {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ResultString != nil {
			_ptr_pstrResultString := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ResultString != nil {
					if err := o.ResultString.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResultString, _ptr_pstrResultString); err != nil {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResultStringOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pstrResultString {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pstrResultString := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ResultString == nil {
				o.ResultString = &oaut.String{}
			}
			if err := o.ResultString.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pstrResultString := func(ptr interface{}) { o.ResultString = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ResultString, _s_pstrResultString, _ptr_pstrResultString); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetResultStringRequest structure represents the GetResultString operation request
type GetResultStringRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lTimeout: MUST be a maximum amount of time, in milliseconds, that the call to GetResultString
	// allows to pass before timing out. If the constant WBEM_INFINITE (0xFFFFFFFF) is used,
	// the GetResultString method call MUST wait until the operation succeeds. This parameter
	// MUST NOT be NULL.
	Timeout int32 `idl:"name:lTimeout" json:"timeout"`
}

func (o *GetResultStringRequest) xxx_ToOp(ctx context.Context) *xxx_GetResultStringOperation {
	if o == nil {
		return &xxx_GetResultStringOperation{}
	}
	return &xxx_GetResultStringOperation{
		This:    o.This,
		Timeout: o.Timeout,
	}
}

func (o *GetResultStringRequest) xxx_FromOp(ctx context.Context, op *xxx_GetResultStringOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Timeout = op.Timeout
}
func (o *GetResultStringRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetResultStringRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResultStringOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetResultStringResponse structure represents the GetResultString operation response
type GetResultStringResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pstrResultString: MUST be a pointer to a BSTR value, which MUST contain the CIM path
	// of the CIM object instance, which MUST lead to the CIM instance that was created
	// using IWbemServices::PutInstance. In case of failure of the semisynchronous operation,
	// the returned string MUST be NULL. When sent by the client, this pointer parameter
	// MUST NOT be NULL. If the original operation does not return a string, the returned
	// string MUST be NULL.
	ResultString *oaut.String `idl:"name:pstrResultString" json:"result_string"`
	// Return: The GetResultString return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetResultStringResponse) xxx_ToOp(ctx context.Context) *xxx_GetResultStringOperation {
	if o == nil {
		return &xxx_GetResultStringOperation{}
	}
	return &xxx_GetResultStringOperation{
		That:         o.That,
		ResultString: o.ResultString,
		Return:       o.Return,
	}
}

func (o *GetResultStringResponse) xxx_FromOp(ctx context.Context, op *xxx_GetResultStringOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ResultString = op.ResultString
	o.Return = op.Return
}
func (o *GetResultStringResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetResultStringResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResultStringOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetResultServicesOperation structure represents the GetResultServices operation
type xxx_GetResultServicesOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Timeout  int32          `idl:"name:lTimeout" json:"timeout"`
	Services *wmi.Services  `idl:"name:ppServices" json:"services"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetResultServicesOperation) OpNum() int { return 5 }

func (o *xxx_GetResultServicesOperation) OpName() string {
	return "/IWbemCallResult/v0/GetResultServices"
}

func (o *xxx_GetResultServicesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResultServicesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lTimeout {in} (1:(int32))
	{
		if err := w.WriteData(o.Timeout); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResultServicesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lTimeout {in} (1:(int32))
	{
		if err := w.ReadData(&o.Timeout); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResultServicesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResultServicesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppServices {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IWbemServices}(interface))
	{
		if o.Services != nil {
			_ptr_ppServices := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Services != nil {
					if err := o.Services.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.Services{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Services, _ptr_ppServices); err != nil {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResultServicesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppServices {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IWbemServices}(interface))
	{
		_ptr_ppServices := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Services == nil {
				o.Services = &wmi.Services{}
			}
			if err := o.Services.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppServices := func(ptr interface{}) { o.Services = *ptr.(**wmi.Services) }
		if err := w.ReadPointer(&o.Services, _s_ppServices, _ptr_ppServices); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetResultServicesRequest structure represents the GetResultServices operation request
type GetResultServicesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lTimeout: MUST be the time, in milliseconds, that the call to GetResultServices allows
	// to pass before timing out. If the constant WBEM_INFINITE (0xFFFFFFFF) is used, the
	// Skip method call MUST wait until the operation succeeds.
	Timeout int32 `idl:"name:lTimeout" json:"timeout"`
}

func (o *GetResultServicesRequest) xxx_ToOp(ctx context.Context) *xxx_GetResultServicesOperation {
	if o == nil {
		return &xxx_GetResultServicesOperation{}
	}
	return &xxx_GetResultServicesOperation{
		This:    o.This,
		Timeout: o.Timeout,
	}
}

func (o *GetResultServicesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetResultServicesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Timeout = op.Timeout
}
func (o *GetResultServicesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetResultServicesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResultServicesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetResultServicesResponse structure represents the GetResultServices operation response
type GetResultServicesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppServices: MUST be a pointer to the IWbemServices interface that is requested by
	// the original call to IWbemServices::OpenNamespace when that interface becomes available.
	// If the semisynchronous operation fails, the returned parameter MUST be NULL. When
	// sent by the client, this pointer parameter MUST NOT be NULL. If the original operation
	// does not return an interface pointer, the returned parameter MUST be NULL.
	Services *wmi.Services `idl:"name:ppServices" json:"services"`
	// Return: The GetResultServices return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetResultServicesResponse) xxx_ToOp(ctx context.Context) *xxx_GetResultServicesOperation {
	if o == nil {
		return &xxx_GetResultServicesOperation{}
	}
	return &xxx_GetResultServicesOperation{
		That:     o.That,
		Services: o.Services,
		Return:   o.Return,
	}
}

func (o *GetResultServicesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetResultServicesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Services = op.Services
	o.Return = op.Return
}
func (o *GetResultServicesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetResultServicesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResultServicesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetCallStatusOperation structure represents the GetCallStatus operation
type xxx_GetCallStatusOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Timeout int32          `idl:"name:lTimeout" json:"timeout"`
	Status  int32          `idl:"name:plStatus" json:"status"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCallStatusOperation) OpNum() int { return 6 }

func (o *xxx_GetCallStatusOperation) OpName() string { return "/IWbemCallResult/v0/GetCallStatus" }

func (o *xxx_GetCallStatusOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCallStatusOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lTimeout {in} (1:(int32))
	{
		if err := w.WriteData(o.Timeout); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCallStatusOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lTimeout {in} (1:(int32))
	{
		if err := w.ReadData(&o.Timeout); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCallStatusOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCallStatusOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plStatus {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCallStatusOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plStatus {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetCallStatusRequest structure represents the GetCallStatus operation request
type GetCallStatusRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lTimeout: MUST be the maximum amount of time, in milliseconds, that the call to GetCallStatus
	// allows to pass before timing out. If the constant WBEM_INFINITE (0xFFFFFFFF) is used,
	// the Skip method call waits until the operation succeeds.
	Timeout int32 `idl:"name:lTimeout" json:"timeout"`
}

func (o *GetCallStatusRequest) xxx_ToOp(ctx context.Context) *xxx_GetCallStatusOperation {
	if o == nil {
		return &xxx_GetCallStatusOperation{}
	}
	return &xxx_GetCallStatusOperation{
		This:    o.This,
		Timeout: o.Timeout,
	}
}

func (o *GetCallStatusRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCallStatusOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Timeout = op.Timeout
}
func (o *GetCallStatusRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetCallStatusRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCallStatusOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCallStatusResponse structure represents the GetCallStatus operation response
type GetCallStatusResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// plStatus:  MUST be the status of a call to an IWbemServices method if the WBEM_S_NO_ERROR
	// code is returned for this method. When sent by the client, this parameter MUST NOT
	// be NULL. Upon return by the server, this parameter can be NULL if there is a failure
	// or if there are no results.
	Status int32 `idl:"name:plStatus" json:"status"`
	// Return: The GetCallStatus return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCallStatusResponse) xxx_ToOp(ctx context.Context) *xxx_GetCallStatusOperation {
	if o == nil {
		return &xxx_GetCallStatusOperation{}
	}
	return &xxx_GetCallStatusOperation{
		That:   o.That,
		Status: o.Status,
		Return: o.Return,
	}
}

func (o *GetCallStatusResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCallStatusOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Status = op.Status
	o.Return = op.Return
}
func (o *GetCallStatusResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetCallStatusResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCallStatusOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
