package idispatch

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
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/oaut"
)

var (
	// IDispatch interface identifier 00020400-0000-0000-c000-000000000046
	DispatchIID = &dcom.IID{Data1: 0x00020400, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
	// Syntax UUID
	DispatchSyntaxUUID = &uuid.UUID{TimeLow: 0x20400, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xc0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x46}}
	// Syntax ID
	DispatchSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: DispatchSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IDispatch interface.
type DispatchClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The GetTypeInfoCount method provides the number of automation type descriptions in
	// the type information table.
	//
	// The GetTypeInfoCount method specifies whether the automation server provides Type
	// description information.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	GetTypeInfoCount(context.Context, *GetTypeInfoCountRequest, ...dcerpc.CallOption) (*GetTypeInfoCountResponse, error)

	// The GetTypeInfo method provides access to the Type description information exposed
	// by the automation server.
	//
	// The GetTypeInfo method retrieves the automation type description that has the specified
	// ordinal position within the type information table.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1 and the entire HRESULT ( ../ms-dtyp/a9046ed2-bfb2-4d56-a719-2824afce59ac
	// ) DWORD does not match the value in the following table, a fatal failure occurred.
	//
	// * If the severity bit is set to 1 and the entire HRESULT DWORD matches the value
	// in the following table, a failure occurred.
	//
	// Return value/code
	//
	// # Description
	//
	// 0x8002000B
	//
	// DISP_E_BADINDEX
	//
	// SHOULD be returned when the value of the passed in iTInfo argument was not 0. See
	// [MS-ERREF].
	GetTypeInfo(context.Context, *GetTypeInfoRequest, ...dcerpc.CallOption) (*GetTypeInfoResponse, error)

	// The GetIDsOfNames method maps a single member (method or property) name, and an optional
	// set of argument names, to a corresponding set of integer DISPIDs, which can be used
	// on subsequent calls to IDispatch::Invoke.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, as
	// defined in [MS-ERREF] section 2.1. The severity bit in the structure identifies the
	// following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1 and the entire HRESULT DWORD does not match a value
	// in the following table, a fatal failure occurred.
	//
	// * If the severity bit is set to 1 and the entire HRESULT DWORD matches a value in
	// the following table, a failure occurred.
	//
	// Return value/code
	//
	// # Description
	//
	// 0x80020006
	//
	// DISP_E_UNKNOWNNAME
	//
	// One or more names were not known. The returned array of DISPIDs MUST contain at least
	// one DISPID_UNKNOWN, and there MUST be one DISPID_UNKNOWN for each entry that corresponds
	// to an unknown name. See [MS-ERREF].
	//
	// 0x80020001
	//
	// DISP_E_UNKNOWNINTERFACE
	//
	// The interface identifier passed in riid is not IID_NULL. See [MS-ERREF].
	//
	// Exceptions Thrown: No exceptions are thrown from this method except those that are
	// thrown by the underlying RPC Protocol specified in [MS-RPCE].
	//
	// When GetIDsOfNames is called with more than one name, the first name (rgszNames[0])
	// corresponds to the member name, and subsequent names correspond to the names of member
	// parameters.
	//
	// The same name can map to different DISPIDs, depending on context. For example, a
	// name can have a DISPID when it is used as: a member name with a particular interface,
	// a different ID as a member of a different interface, or a different mapping for each
	// time it appears as a parameter.
	//
	// The implementation of GetIDsOfNames MUST be case-insensitive.
	//
	// An implementation of the OLE Automation Protocol MAY<55> choose to implement a mapping
	// for the parameter names that map to the index of the parameter in the member parameter
	// list.
	GetIDsOfNames(context.Context, *GetIDsOfNamesRequest, ...dcerpc.CallOption) (*GetIDsOfNamesResponse, error)

	// The Invoke method provides access to properties and methods exposed by the automation
	// server.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1 and the entire HRESULT DWORD does not match a value
	// in the following table, a fatal failure occurred.
	//
	// * If the severity bit is set to 1 and the entire HRESULT DWORD matches a value in
	// the following table, a failure occurred.
	//
	// Return value/code
	//
	// # Description
	//
	// 0x80020009
	//
	// DISP_E_EXCEPTION
	//
	// The application needs to raise an exception. In this case, the structure passed in
	// pExcepInfo MUST be filled in with a nonzero error code. See [MS-ERREF].
	//
	// 0x80020004
	//
	// DISP_E_PARAMNOTFOUND
	//
	// One of the parameter DISPIDs does not correspond to a parameter on the method. In
	// this case, pArgErr MUST be set to the first argument that contains the error. See
	// [MS-ERREF].
	//
	// 0x80020005
	//
	// DISP_E_TYPEMISMATCH
	//
	// One or more of the arguments could not be coerced into the type of the corresponding
	// formal parameter. The index within rgvarg of the first parameter with the incorrect
	// type MUST be returned in the pArgErr parameter. For more information, see section
	// 3.1.4.4.4 ( 5c01ab3c-f719-44cc-bb45-d36850cf4c5b ) and [MS-ERREF].
	//
	// 0x8002000E
	//
	// DISP_E_BADPARAMCOUNT
	//
	// The number of elements provided to DISPPARAMS is different from the number of arguments
	// accepted by the method or property. See [MS-ERREF].
	//
	// 0x80020008
	//
	// DISP_E_BADVARTYPE
	//
	// One of the arguments in rgvarg is not a valid variant type. See [MS-ERREF].
	//
	// 0x80020003
	//
	// DISP_E_MEMBERNOTFOUND
	//
	// The requested member does not exist, or the call to Invoke tried to set the value
	// of a read-only property. See [MS-ERREF].
	//
	// 0x80020007
	//
	// DISP_E_NONAMEDARGS
	//
	// At least one named argument was provided for methods with a vararg parameter (see
	// section 3.1.4.4.3 ( be6e35f6-9327-4164-9bde-ffcd0fa0e07d ) ), for which named arguments
	// are illegal. See [MS-ERREF].
	//
	// 0x8002000A
	//
	// DISP_E_OVERFLOW
	//
	// One of the arguments in rgvarg could not be coerced to the type of its corresponding
	// formal parameter. See [MS-ERREF].
	//
	// 0x80020001
	//
	// DISP_E_UNKNOWNINTERFACE
	//
	// The interface identifier passed in riid is not IID_NULL. See [MS-ERREF].
	//
	// 0x8002000F
	//
	// DISP_E_PARAMNOTOPTIONAL
	//
	// A required parameter was omitted. See [MS-ERREF].
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC Protocol, as specified in [MS-RPCE].
	Invoke(context.Context, *InvokeRequest, ...dcerpc.CallOption) (*InvokeResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) DispatchClient
}

// DispatchMethod represents the DISPATCH_METHOD RPC constant
const DispatchMethod = 0x00000001

// DispatchPropertyGet represents the DISPATCH_PROPERTYGET RPC constant
const DispatchPropertyGet = 0x00000002

// DispatchPropertyPut represents the DISPATCH_PROPERTYPUT RPC constant
const DispatchPropertyPut = 0x00000004

// DispatchPropertyPutReference represents the DISPATCH_PROPERTYPUTREF RPC constant
const DispatchPropertyPutReference = 0x00000008

// DispatchZeroVarResult represents the DISPATCH_zeroVarResult RPC constant
const DispatchZeroVarResult = 0x00020000

// DispatchZeroExceptionInfo represents the DISPATCH_zeroExcepInfo RPC constant
const DispatchZeroExceptionInfo = 0x00040000

// DispatchZeroArgError represents the DISPATCH_zeroArgErr RPC constant
const DispatchZeroArgError = 0x00080000

// DispatchIDValue represents the DISPID_VALUE RPC constant
var DispatchIDValue = 0

// DispatchIDUnknown represents the DISPID_UNKNOWN RPC constant
var DispatchIDUnknown = -1

// DispatchIDPropertyPut represents the DISPID_PROPERTYPUT RPC constant
var DispatchIDPropertyPut = -3

// DispatchIDNewenum represents the DISPID_NEWENUM RPC constant
var DispatchIDNewenum = -4

type xxx_DefaultDispatchClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultDispatchClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultDispatchClient) GetTypeInfoCount(ctx context.Context, in *GetTypeInfoCountRequest, opts ...dcerpc.CallOption) (*GetTypeInfoCountResponse, error) {
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
	out := &GetTypeInfoCountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDispatchClient) GetTypeInfo(ctx context.Context, in *GetTypeInfoRequest, opts ...dcerpc.CallOption) (*GetTypeInfoResponse, error) {
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
	out := &GetTypeInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDispatchClient) GetIDsOfNames(ctx context.Context, in *GetIDsOfNamesRequest, opts ...dcerpc.CallOption) (*GetIDsOfNamesResponse, error) {
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
	out := &GetIDsOfNamesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDispatchClient) Invoke(ctx context.Context, in *InvokeRequest, opts ...dcerpc.CallOption) (*InvokeResponse, error) {
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
	out := &InvokeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDispatchClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultDispatchClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultDispatchClient) IPID(ctx context.Context, ipid *dcom.IPID) DispatchClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultDispatchClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewDispatchClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (DispatchClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(DispatchSyntaxV0_0))...)
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
	return &xxx_DefaultDispatchClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetTypeInfoCountOperation structure represents the GetTypeInfoCount operation
type xxx_GetTypeInfoCountOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	TypeInfoCount uint32         `idl:"name:pctinfo" json:"type_info_count"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTypeInfoCountOperation) OpNum() int { return 3 }

func (o *xxx_GetTypeInfoCountOperation) OpName() string { return "/IDispatch/v0/GetTypeInfoCount" }

func (o *xxx_GetTypeInfoCountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeInfoCountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetTypeInfoCountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetTypeInfoCountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeInfoCountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pctinfo {out} (1:{pointer=ref}*(1))(2:{alias=UINT}(uint32))
	{
		if err := w.WriteData(o.TypeInfoCount); err != nil {
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

func (o *xxx_GetTypeInfoCountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pctinfo {out} (1:{pointer=ref}*(1))(2:{alias=UINT}(uint32))
	{
		if err := w.ReadData(&o.TypeInfoCount); err != nil {
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

// GetTypeInfoCountRequest structure represents the GetTypeInfoCount operation request
type GetTypeInfoCountRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetTypeInfoCountRequest) xxx_ToOp(ctx context.Context, op *xxx_GetTypeInfoCountOperation) *xxx_GetTypeInfoCountOperation {
	if op == nil {
		op = &xxx_GetTypeInfoCountOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetTypeInfoCountRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTypeInfoCountOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetTypeInfoCountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetTypeInfoCountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTypeInfoCountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTypeInfoCountResponse structure represents the GetTypeInfoCount operation response
type GetTypeInfoCountResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pctinfo: MUST be set to 0 if the automation server does not provide Type description
	// information. Otherwise, it MUST be set to 1.
	TypeInfoCount uint32 `idl:"name:pctinfo" json:"type_info_count"`
	// Return: The GetTypeInfoCount return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTypeInfoCountResponse) xxx_ToOp(ctx context.Context, op *xxx_GetTypeInfoCountOperation) *xxx_GetTypeInfoCountOperation {
	if op == nil {
		op = &xxx_GetTypeInfoCountOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.TypeInfoCount = o.TypeInfoCount
	op.Return = o.Return
	return op
}

func (o *GetTypeInfoCountResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTypeInfoCountOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TypeInfoCount = op.TypeInfoCount
	o.Return = op.Return
}
func (o *GetTypeInfoCountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetTypeInfoCountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTypeInfoCountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTypeInfoOperation structure represents the GetTypeInfo operation
type xxx_GetTypeInfoOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	TypeInfoIndex uint32         `idl:"name:iTInfo" json:"type_info_index"`
	LocaleID      uint32         `idl:"name:lcid" json:"locale_id"`
	TypeInfo      *oaut.TypeInfo `idl:"name:ppTInfo" json:"type_info"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTypeInfoOperation) OpNum() int { return 4 }

func (o *xxx_GetTypeInfoOperation) OpName() string { return "/IDispatch/v0/GetTypeInfo" }

func (o *xxx_GetTypeInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// iTInfo {in} (1:{alias=UINT}(uint32))
	{
		if err := w.WriteData(o.TypeInfoIndex); err != nil {
			return err
		}
	}
	// lcid {in} (1:{alias=LCID, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.LocaleID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// iTInfo {in} (1:{alias=UINT}(uint32))
	{
		if err := w.ReadData(&o.TypeInfoIndex); err != nil {
			return err
		}
	}
	// lcid {in} (1:{alias=LCID, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.LocaleID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppTInfo {out} (1:{pointer=ref}*(2)*(1))(2:{alias=ITypeInfo}(interface))
	{
		if o.TypeInfo != nil {
			_ptr_ppTInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.TypeInfo != nil {
					if err := o.TypeInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.TypeInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.TypeInfo, _ptr_ppTInfo); err != nil {
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

func (o *xxx_GetTypeInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppTInfo {out} (1:{pointer=ref}*(2)*(1))(2:{alias=ITypeInfo}(interface))
	{
		_ptr_ppTInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.TypeInfo == nil {
				o.TypeInfo = &oaut.TypeInfo{}
			}
			if err := o.TypeInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppTInfo := func(ptr interface{}) { o.TypeInfo = *ptr.(**oaut.TypeInfo) }
		if err := w.ReadPointer(&o.TypeInfo, _s_ppTInfo, _ptr_ppTInfo); err != nil {
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

// GetTypeInfoRequest structure represents the GetTypeInfo operation request
type GetTypeInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// iTInfo: MUST be 0.
	TypeInfoIndex uint32 `idl:"name:iTInfo" json:"type_info_index"`
	// lcid: MUST equal the locale ID for the Type description information to be retrieved.
	LocaleID uint32 `idl:"name:lcid" json:"locale_id"`
}

func (o *GetTypeInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetTypeInfoOperation) *xxx_GetTypeInfoOperation {
	if op == nil {
		op = &xxx_GetTypeInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.TypeInfoIndex = o.TypeInfoIndex
	op.LocaleID = o.LocaleID
	return op
}

func (o *GetTypeInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTypeInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.TypeInfoIndex = op.TypeInfoIndex
	o.LocaleID = op.LocaleID
}
func (o *GetTypeInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetTypeInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTypeInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTypeInfoResponse structure represents the GetTypeInfo operation response
type GetTypeInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppTInfo: MUST be set to reference an instance of an ITypeInfo server that corresponds
	// to the default nonsource interface of the coclass implementing IDispatch (see section
	// 2.2.49.8). MUST refer to the partner dispinterface if the default nonsource interface
	// is a dual interface. MUST be set to NULL if the coclass does not specify a default
	// nonsource interface.
	//
	// ppTInfo: MUST be set to a reference to the ITypeInfo server instance (see section
	// 3.7) with the specified position in the type information table, or to NULL if the
	// value of index is greater than or equal to the number of automation type descriptions
	// in the type information table.
	TypeInfo *oaut.TypeInfo `idl:"name:ppTInfo" json:"type_info"`
	// Return: The GetTypeInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTypeInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetTypeInfoOperation) *xxx_GetTypeInfoOperation {
	if op == nil {
		op = &xxx_GetTypeInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.TypeInfo = o.TypeInfo
	op.Return = o.Return
	return op
}

func (o *GetTypeInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTypeInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TypeInfo = op.TypeInfo
	o.Return = op.Return
}
func (o *GetTypeInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetTypeInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTypeInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetIDsOfNamesOperation structure represents the GetIDsOfNames operation
type xxx_GetIDsOfNamesOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	IID        *dcom.IID      `idl:"name:riid" json:"iid"`
	Names      []string       `idl:"name:rgszNames;size_is:(cNames)" json:"names"`
	NamesCount uint32         `idl:"name:cNames" json:"names_count"`
	LocaleID   uint32         `idl:"name:lcid" json:"locale_id"`
	DispatchID []int32        `idl:"name:rgDispId;size_is:(cNames)" json:"dispatch_id"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIDsOfNamesOperation) OpNum() int { return 5 }

func (o *xxx_GetIDsOfNamesOperation) OpName() string { return "/IDispatch/v0/GetIDsOfNames" }

func (o *xxx_GetIDsOfNamesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Names != nil && o.NamesCount == 0 {
		o.NamesCount = uint32(len(o.Names))
	}
	if o.NamesCount > uint32(16384) {
		return fmt.Errorf("NamesCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIDsOfNamesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// riid {in} (1:{alias=REFIID}*(1))(2:{alias=IID, names=GUID}(struct))
	{
		if o.IID != nil {
			if err := o.IID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.IID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// rgszNames {in} (1:{pointer=ref}*(1))(2:{string, alias=LPOLESTR}[dim:0,size_is=cNames]*(1)[dim:0,string,null](wchar))
	{
		dimSize1 := uint64(o.NamesCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Names {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Names[i1] != "" {
				_ptr_rgszNames := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if err := ndr.WriteUTF16NString(ctx, w, o.Names[i1]); err != nil {
						return err
					}
					return nil
				})
				if err := w.WritePointer(&o.Names[i1], _ptr_rgszNames); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Names); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// cNames {in} (1:{range=(0,16384), alias=UINT}(uint32))
	{
		if err := w.WriteData(o.NamesCount); err != nil {
			return err
		}
	}
	// lcid {in} (1:{alias=LCID, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.LocaleID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIDsOfNamesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// riid {in} (1:{alias=REFIID,pointer=ref}*(1))(2:{alias=IID, names=GUID}(struct))
	{
		if o.IID == nil {
			o.IID = &dcom.IID{}
		}
		if err := o.IID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// rgszNames {in} (1:{pointer=ref}*(1))(2:{string, alias=LPOLESTR}[dim:0,size_is=cNames]*(1)[dim:0,string,null](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Names", sizeInfo[0])
		}
		o.Names = make([]string, sizeInfo[0])
		for i1 := range o.Names {
			i1 := i1
			_ptr_rgszNames := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if err := ndr.ReadUTF16NString(ctx, w, &o.Names[i1]); err != nil {
					return err
				}
				return nil
			})
			_s_rgszNames := func(ptr interface{}) { o.Names[i1] = *ptr.(*string) }
			if err := w.ReadPointer(&o.Names[i1], _s_rgszNames, _ptr_rgszNames); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cNames {in} (1:{range=(0,16384), alias=UINT}(uint32))
	{
		if err := w.ReadData(&o.NamesCount); err != nil {
			return err
		}
	}
	// lcid {in} (1:{alias=LCID, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.LocaleID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIDsOfNamesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIDsOfNamesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// rgDispId {out} (1:{pointer=ref}*(1))(2:{alias=DISPID, names=LONG}[dim:0,size_is=cNames](int32))
	{
		dimSize1 := uint64(o.NamesCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.DispatchID {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.DispatchID[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.DispatchID); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(int32(0)); err != nil {
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

func (o *xxx_GetIDsOfNamesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// rgDispId {out} (1:{pointer=ref}*(1))(2:{alias=DISPID, names=LONG}[dim:0,size_is=cNames](int32))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.DispatchID", sizeInfo[0])
		}
		o.DispatchID = make([]int32, sizeInfo[0])
		for i1 := range o.DispatchID {
			i1 := i1
			if err := w.ReadData(&o.DispatchID[i1]); err != nil {
				return err
			}
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

// GetIDsOfNamesRequest structure represents the GetIDsOfNames operation request
type GetIDsOfNamesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// riid: MUST equal IID_NULL (see section 1.9).
	IID *dcom.IID `idl:"name:riid" json:"iid"`
	// rgszNames: MUST be the array of strings to be mapped. The first string in the array
	// MUST specify the name of a method or property that is supported by the server. Any
	// additional strings MUST contain the names of all arguments for the method or property
	// that is specified by the value in the first string. The mapping MUST be case-insensitive.
	Names []string `idl:"name:rgszNames;size_is:(cNames)" json:"names"`
	// cNames: MUST equal the count of names to be mapped, and MUST<54> be between 0 and
	// 16384.
	NamesCount uint32 `idl:"name:cNames" json:"names_count"`
	// lcid: MUST equal the locale ID in which to interpret the names.
	LocaleID uint32 `idl:"name:lcid" json:"locale_id"`
}

func (o *GetIDsOfNamesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetIDsOfNamesOperation) *xxx_GetIDsOfNamesOperation {
	if op == nil {
		op = &xxx_GetIDsOfNamesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.IID = o.IID
	op.Names = o.Names
	op.NamesCount = o.NamesCount
	op.LocaleID = o.LocaleID
	return op
}

func (o *GetIDsOfNamesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetIDsOfNamesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.IID = op.IID
	o.Names = op.Names
	o.NamesCount = op.NamesCount
	o.LocaleID = op.LocaleID
}
func (o *GetIDsOfNamesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetIDsOfNamesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIDsOfNamesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIDsOfNamesResponse structure represents the GetIDsOfNames operation response
type GetIDsOfNamesResponse struct {
	// XXX: cNames is an implicit input depedency for output parameters
	NamesCount uint32 `idl:"name:cNames" json:"names_count"`

	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// rgDispId: MUST be an array of DISPIDs that are filled in by the server. Each DISPID
	// corresponds, by position, to one of the names passed in rgszNames.
	DispatchID []int32 `idl:"name:rgDispId;size_is:(cNames)" json:"dispatch_id"`
	// Return: The GetIDsOfNames return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIDsOfNamesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetIDsOfNamesOperation) *xxx_GetIDsOfNamesOperation {
	if op == nil {
		op = &xxx_GetIDsOfNamesOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.NamesCount == uint32(0) {
		op.NamesCount = o.NamesCount
	}

	op.That = o.That
	op.DispatchID = o.DispatchID
	op.Return = o.Return
	return op
}

func (o *GetIDsOfNamesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetIDsOfNamesOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.NamesCount = op.NamesCount

	o.That = op.That
	o.DispatchID = op.DispatchID
	o.Return = op.Return
}
func (o *GetIDsOfNamesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetIDsOfNamesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIDsOfNamesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_InvokeOperation structure represents the Invoke operation
type xxx_InvokeOperation struct {
	This              *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat       `idl:"name:That" json:"that"`
	DispatchIDMember  int32                `idl:"name:dispIdMember" json:"dispatch_id_member"`
	IID               *dcom.IID            `idl:"name:riid" json:"iid"`
	LocaleID          uint32               `idl:"name:lcid" json:"locale_id"`
	Flags             uint32               `idl:"name:dwFlags" json:"flags"`
	DispatchParams    *oaut.DispatchParams `idl:"name:pDispParams" json:"dispatch_params"`
	VarResult         *oaut.Variant        `idl:"name:pVarResult" json:"var_result"`
	ExceptionInfo     *oaut.ExceptionInfo  `idl:"name:pExcepInfo" json:"exception_info"`
	ArgError          uint32               `idl:"name:pArgErr" json:"arg_error"`
	VarReferenceCount uint32               `idl:"name:cVarRef" json:"var_reference_count"`
	VarReferenceIndex []uint32             `idl:"name:rgVarRefIdx;size_is:(cVarRef)" json:"var_reference_index"`
	VarReference      []*oaut.Variant      `idl:"name:rgVarRef;size_is:(cVarRef)" json:"var_reference"`
	Return            int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_InvokeOperation) OpNum() int { return 6 }

func (o *xxx_InvokeOperation) OpName() string { return "/IDispatch/v0/Invoke" }

func (o *xxx_InvokeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.VarReferenceIndex != nil && o.VarReferenceCount == 0 {
		o.VarReferenceCount = uint32(len(o.VarReferenceIndex))
	}
	if o.VarReference != nil && o.VarReferenceCount == 0 {
		o.VarReferenceCount = uint32(len(o.VarReference))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InvokeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// dispIdMember {in} (1:{alias=DISPID, names=LONG}(int32))
	{
		if err := w.WriteData(o.DispatchIDMember); err != nil {
			return err
		}
	}
	// riid {in} (1:{alias=REFIID}*(1))(2:{alias=IID, names=GUID}(struct))
	{
		if o.IID != nil {
			if err := o.IID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.IID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lcid {in} (1:{alias=LCID, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.LocaleID); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pDispParams {in} (1:{pointer=ref}*(1))(2:{alias=DISPPARAMS}(struct))
	{
		if o.DispatchParams != nil {
			if err := o.DispatchParams.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.DispatchParams{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// cVarRef {in} (1:{alias=UINT}(uint32))
	{
		if err := w.WriteData(o.VarReferenceCount); err != nil {
			return err
		}
	}
	// rgVarRefIdx {in} (1:{pointer=ref}*(1))(2:{alias=UINT}[dim:0,size_is=cVarRef](uint32))
	{
		dimSize1 := uint64(o.VarReferenceCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.VarReferenceIndex {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.VarReferenceIndex[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.VarReferenceIndex); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint32(0)); err != nil {
				return err
			}
		}
	}
	// rgVarRef {in, out} (1:{pointer=ref}*(1))(2:{alias=VARIANT}[dim:0,size_is=cVarRef]*(1))(3:{alias=_VARIANT}(struct))
	{
		dimSize1 := uint64(o.VarReferenceCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.VarReference {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.VarReference[i1] != nil {
				_ptr_rgVarRef := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if o.VarReference[i1] != nil {
						if err := o.VarReference[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.VarReference[i1], _ptr_rgVarRef); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.VarReference); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InvokeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// dispIdMember {in} (1:{alias=DISPID, names=LONG}(int32))
	{
		if err := w.ReadData(&o.DispatchIDMember); err != nil {
			return err
		}
	}
	// riid {in} (1:{alias=REFIID,pointer=ref}*(1))(2:{alias=IID, names=GUID}(struct))
	{
		if o.IID == nil {
			o.IID = &dcom.IID{}
		}
		if err := o.IID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lcid {in} (1:{alias=LCID, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.LocaleID); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pDispParams {in} (1:{pointer=ref}*(1))(2:{alias=DISPPARAMS}(struct))
	{
		if o.DispatchParams == nil {
			o.DispatchParams = &oaut.DispatchParams{}
		}
		if err := o.DispatchParams.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cVarRef {in} (1:{alias=UINT}(uint32))
	{
		if err := w.ReadData(&o.VarReferenceCount); err != nil {
			return err
		}
	}
	// rgVarRefIdx {in} (1:{pointer=ref}*(1))(2:{alias=UINT}[dim:0,size_is=cVarRef](uint32))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.VarReferenceIndex", sizeInfo[0])
		}
		o.VarReferenceIndex = make([]uint32, sizeInfo[0])
		for i1 := range o.VarReferenceIndex {
			i1 := i1
			if err := w.ReadData(&o.VarReferenceIndex[i1]); err != nil {
				return err
			}
		}
	}
	// rgVarRef {in, out} (1:{pointer=ref}*(1))(2:{alias=VARIANT}[dim:0,size_is=cVarRef]*(1))(3:{alias=_VARIANT}(struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.VarReference", sizeInfo[0])
		}
		o.VarReference = make([]*oaut.Variant, sizeInfo[0])
		for i1 := range o.VarReference {
			i1 := i1
			_ptr_rgVarRef := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.VarReference[i1] == nil {
					o.VarReference[i1] = &oaut.Variant{}
				}
				if err := o.VarReference[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_rgVarRef := func(ptr interface{}) { o.VarReference[i1] = *ptr.(**oaut.Variant) }
			if err := w.ReadPointer(&o.VarReference[i1], _s_rgVarRef, _ptr_rgVarRef); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InvokeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InvokeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pVarResult {out} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.VarResult != nil {
			_ptr_pVarResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.VarResult != nil {
					if err := o.VarResult.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.VarResult, _ptr_pVarResult); err != nil {
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
	// pExcepInfo {out} (1:{pointer=ref}*(1))(2:{alias=EXCEPINFO}(struct))
	{
		if o.ExceptionInfo != nil {
			if err := o.ExceptionInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.ExceptionInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pArgErr {out} (1:{pointer=ref}*(1))(2:{alias=UINT}(uint32))
	{
		if err := w.WriteData(o.ArgError); err != nil {
			return err
		}
	}
	// rgVarRef {in, out} (1:{pointer=ref}*(1))(2:{alias=VARIANT}[dim:0,size_is=cVarRef]*(1))(3:{alias=_VARIANT}(struct))
	{
		dimSize1 := uint64(o.VarReferenceCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.VarReference {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.VarReference[i1] != nil {
				_ptr_rgVarRef := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if o.VarReference[i1] != nil {
						if err := o.VarReference[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.VarReference[i1], _ptr_rgVarRef); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.VarReference); uint64(i1) < sizeInfo[0]; i1++ {
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

func (o *xxx_InvokeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pVarResult {out} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pVarResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.VarResult == nil {
				o.VarResult = &oaut.Variant{}
			}
			if err := o.VarResult.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pVarResult := func(ptr interface{}) { o.VarResult = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.VarResult, _s_pVarResult, _ptr_pVarResult); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pExcepInfo {out} (1:{pointer=ref}*(1))(2:{alias=EXCEPINFO}(struct))
	{
		if o.ExceptionInfo == nil {
			o.ExceptionInfo = &oaut.ExceptionInfo{}
		}
		if err := o.ExceptionInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pArgErr {out} (1:{pointer=ref}*(1))(2:{alias=UINT}(uint32))
	{
		if err := w.ReadData(&o.ArgError); err != nil {
			return err
		}
	}
	// rgVarRef {in, out} (1:{pointer=ref}*(1))(2:{alias=VARIANT}[dim:0,size_is=cVarRef]*(1))(3:{alias=_VARIANT}(struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.VarReference", sizeInfo[0])
		}
		o.VarReference = make([]*oaut.Variant, sizeInfo[0])
		for i1 := range o.VarReference {
			i1 := i1
			_ptr_rgVarRef := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.VarReference[i1] == nil {
					o.VarReference[i1] = &oaut.Variant{}
				}
				if err := o.VarReference[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_rgVarRef := func(ptr interface{}) { o.VarReference[i1] = *ptr.(**oaut.Variant) }
			if err := w.ReadPointer(&o.VarReference[i1], _s_rgVarRef, _ptr_rgVarRef); err != nil {
				return err
			}
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

// InvokeRequest structure represents the Invoke operation request
type InvokeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// dispIdMember: MUST equal the DISPID of the method or property to call.
	DispatchIDMember int32 `idl:"name:dispIdMember" json:"dispatch_id_member"`
	// riid: MUST equal IID_NULL (see section 1.9).
	IID *dcom.IID `idl:"name:riid" json:"iid"`
	// lcid: MUST equal a locale ID supported by the automation server. This value SHOULD
	// be used by the automation server if any of the arguments are strings whose meaning
	// is dependent on a specific locale ID. If no such strings are present in the arguments
	// the server SHOULD ignore this value.
	LocaleID uint32 `idl:"name:lcid" json:"locale_id"`
	// dwFlags:  MUST be a combination of the bit flags specified in the following table.
	//
	// Note  The value MUST specify one and only one of the following bit flags: DISPATCH_METHOD,
	// DISPATCH_PROPERTYGET, DISPATCH_PROPERTYPUT, or DISPATCH_PROPERTYPUTREF.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|                                    |                                                                                  |
	//	|               VALUE                |                                     MEANING                                      |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| DISPATCH_METHOD 0x00000001         | The member is invoked as a method.                                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| DISPATCH_PROPERTYGET 0x00000002    | The member is retrieved as a property or data member.                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| DISPATCH_PROPERTYPUT 0x00000004    | The member is changed as a property or data member.                              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| DISPATCH_PROPERTYPUTREF 0x00000008 | The member is changed by a reference assignment, rather than by a value          |
	//	|                                    | assignment. This flag is valid only when the property accepts a reference to an  |
	//	|                                    | object.                                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| DISPATCH_zeroVarResult 0x00020000  | MUST specify that the client is not interested in the actual pVarResult [out]    |
	//	|                                    | argument. On return the pVarResult argument MUST point to a VT_EMPTY variant,    |
	//	|                                    | with all reserved fields set to 0.                                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| DISPATCH_zeroExcepInfo 0x00040000  | MUST specify that the client is not interested in the actual pExcepInfo [out]    |
	//	|                                    | argument. On return pExcepInfo MUST point to an EXCEPINFO structure, with all    |
	//	|                                    | scalar fields set to 0 and all BSTR fields set to NULL.                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| DISPATCH_zeroArgErr 0x00080000     | MUST specify that the client is not interested in the actual pArgErr [out]       |
	//	|                                    | argument. On return, pArgErr MUST be set to 0.                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
	// pDispParams:  MUST point to a DISPPARAMS structure that defines the arguments passed
	// to the method. Arguments MUST be stored in pDispParams->rgvarg in reverse order,
	// so that the first argument is the one with the highest index in the array. Byref
	// arguments MUST be marked in this array as VT_EMPTY entries, and stored in rgVarRef
	// instead. For more information, see section 2.2.33.
	DispatchParams *oaut.DispatchParams `idl:"name:pDispParams" json:"dispatch_params"`
	// cVarRef: MUST equal the number of byref arguments passed in pDispParams.
	VarReferenceCount uint32 `idl:"name:cVarRef" json:"var_reference_count"`
	// rgVarRefIdx: MUST contain an array of cVarRef unsigned integers that holds the indices
	// of the byref arguments marked as VT_EMPTY entries in pDispParams->rgvarg.
	VarReferenceIndex []uint32 `idl:"name:rgVarRefIdx;size_is:(cVarRef)" json:"var_reference_index"`
	// rgVarRef: MUST contain the byref arguments as set by the client at the time of the
	// call, and by the server on successful return from the call. Arguments in this array
	// MUST also be stored in reverse order, so that the first byref argument has the highest
	// index in the array.
	VarReference []*oaut.Variant `idl:"name:rgVarRef;size_is:(cVarRef)" json:"var_reference"`
}

func (o *InvokeRequest) xxx_ToOp(ctx context.Context, op *xxx_InvokeOperation) *xxx_InvokeOperation {
	if op == nil {
		op = &xxx_InvokeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.DispatchIDMember = o.DispatchIDMember
	op.IID = o.IID
	op.LocaleID = o.LocaleID
	op.Flags = o.Flags
	op.DispatchParams = o.DispatchParams
	op.VarReferenceCount = o.VarReferenceCount
	op.VarReferenceIndex = o.VarReferenceIndex
	op.VarReference = o.VarReference
	return op
}

func (o *InvokeRequest) xxx_FromOp(ctx context.Context, op *xxx_InvokeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DispatchIDMember = op.DispatchIDMember
	o.IID = op.IID
	o.LocaleID = op.LocaleID
	o.Flags = op.Flags
	o.DispatchParams = op.DispatchParams
	o.VarReferenceCount = op.VarReferenceCount
	o.VarReferenceIndex = op.VarReferenceIndex
	o.VarReference = op.VarReference
}
func (o *InvokeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *InvokeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InvokeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// InvokeResponse structure represents the Invoke operation response
type InvokeResponse struct {
	// XXX: cVarRef is an implicit input depedency for output parameters
	VarReferenceCount uint32 `idl:"name:cVarRef" json:"var_reference_count"`

	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pVarResult: MUST point to a VARIANT that will be filled with the result of the method
	// or property call.
	VarResult *oaut.Variant `idl:"name:pVarResult" json:"var_result"`
	// pExcepInfo: If this value is not null and the return value is DISP_E_EXCEPTION, this
	// structure MUST be filled by the automation server. Otherwise, it MUST specify a 0
	// value for the scode and wCode fields, and it MUST be ignored on receipt.
	ExceptionInfo *oaut.ExceptionInfo `idl:"name:pExcepInfo" json:"exception_info"`
	// pArgErr: If this value is not null and the return value is DISP_E_TYPEMISMATCH or
	// DISP_E_PARAMNOTFOUND, this argument MUST equal the index (within pDispParams->rgvarg)
	// of the first argument that has an error. Otherwise, it MUST be ignored on receipt.
	ArgError uint32 `idl:"name:pArgErr" json:"arg_error"`
	// rgVarRef: MUST contain the byref arguments as set by the client at the time of the
	// call, and by the server on successful return from the call. Arguments in this array
	// MUST also be stored in reverse order, so that the first byref argument has the highest
	// index in the array.
	VarReference []*oaut.Variant `idl:"name:rgVarRef;size_is:(cVarRef)" json:"var_reference"`
	// Return: The Invoke return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *InvokeResponse) xxx_ToOp(ctx context.Context, op *xxx_InvokeOperation) *xxx_InvokeOperation {
	if op == nil {
		op = &xxx_InvokeOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.VarReferenceCount == uint32(0) {
		op.VarReferenceCount = o.VarReferenceCount
	}

	op.That = o.That
	op.VarResult = o.VarResult
	op.ExceptionInfo = o.ExceptionInfo
	op.ArgError = o.ArgError
	op.VarReference = o.VarReference
	op.Return = o.Return
	return op
}

func (o *InvokeResponse) xxx_FromOp(ctx context.Context, op *xxx_InvokeOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.VarReferenceCount = op.VarReferenceCount

	o.That = op.That
	o.VarResult = op.VarResult
	o.ExceptionInfo = op.ExceptionInfo
	o.ArgError = op.ArgError
	o.VarReference = op.VarReference
	o.Return = op.Return
}
func (o *InvokeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *InvokeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InvokeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
