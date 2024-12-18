package itypelib

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
	_ = oaut.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/oaut"
)

var (
	// ITypeLib interface identifier 00020402-0000-0000-c000-000000000046
	TypeLibIID = &dcom.IID{Data1: 0x00020402, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
	// Syntax UUID
	TypeLibSyntaxUUID = &uuid.UUID{TimeLow: 0x20402, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xc0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x46}}
	// Syntax ID
	TypeLibSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: TypeLibSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// ITypeLib interface.
type TypeLibClient interface {

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

	// The GetTypeInfoType method retrieves the TYPEKIND value associated with an automation
	// type description.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1 and the entire *HRESULT* DWORD does not match a
	// value in the following table, a fatal failure occurred.
	//
	// * If the severity bit is set to 1 and the entire *HRESULT* DWORD matches a value
	// in the following table, a failure occurred.
	//
	// Return value/code
	//
	// # Description
	//
	// 0x8002802B
	//
	// TYPE_E_ELEMENTNOTFOUND
	//
	// The value of index did not specify the ordinal position of an element in the type
	// information table. See [MS-ERREF].
	GetTypeInfoType(context.Context, *GetTypeInfoTypeRequest, ...dcerpc.CallOption) (*GetTypeInfoTypeResponse, error)

	// The GetTypeInfoOfGuid method retrieves the automation type description with the specified
	// GUID from the server's type information table.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1 and the entire *HRESULT* DWORD does not match a
	// value in the following table, a fatal failure occurred.
	//
	// * If the severity bit is set to 1 and the entire *HRESULT* DWORD matches a value
	// in the following table, a failure occurred.
	//
	// Return value/code
	//
	// # Description
	//
	// 0x8002802B
	//
	// TYPE_E_ELEMENTNOTFOUND
	//
	// The value of guid did not correspond to any entry in the Type information table,
	// or the value of guid was IID_NULL. See [MS-ERREF].
	GetTypeInfoOfGUID(context.Context, *GetTypeInfoOfGUIDRequest, ...dcerpc.CallOption) (*GetTypeInfoOfGUIDResponse, error)

	// The GetLibAttr method retrieves a structure that contains the attributes declared
	// with the Type library.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	GetLibAttribute(context.Context, *GetLibAttributeRequest, ...dcerpc.CallOption) (*GetLibAttributeResponse, error)

	// The GetTypeComp method retrieves a reference to the ITypeComp server instance associated
	// with the ITypeInfo server.
	//
	// The GetTypeComp method retrieves a reference to the ITypeComp server instance associated
	// with the ITypeLib server.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the sevierty bit is set to 1, the method failed and encountered a fatal error.
	GetTypeComp(context.Context, *GetTypeCompRequest, ...dcerpc.CallOption) (*GetTypeCompResponse, error)

	// The GetDocumentation method retrieves the documentation resources associated with
	// a type member.
	//
	// The GetDocumentation method retrieves the documentation resources associated with
	// the automation type library.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// If memid is MEMBERID_NIL, the values of pBstrName, pBstrDocString, pdwHelpContext,
	// and pBstrHelpFile MUST correspond to the attributes declared with the type, as specified
	// in section 2.2.49.3. Otherwise, they MUST correspond to the attributes declared with
	// the specified member of the type.
	GetDocumentation(context.Context, *GetDocumentationRequest, ...dcerpc.CallOption) (*GetDocumentationResponse, error)

	// The IsName method indicates whether the specified string matches the name of a type
	// or type member that is contained in the automation type library or its binding context.
	//
	// Return Values: The method MUST return the information in an HRESULT data structure,
	// which is defined in [MS-ERREF] section 2.1. The severity bit in the structure identifies
	// the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	IsName(context.Context, *IsNameRequest, ...dcerpc.CallOption) (*IsNameResponse, error)

	// The FindName method retrieves references to types, or type members, contained in
	// the automation type library whose names match a specified string.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	FindName(context.Context, *FindNameRequest, ...dcerpc.CallOption) (*FindNameResponse, error)

	// Opnum12NotUsedOnWire operation.
	// Opnum12NotUsedOnWire

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) TypeLibClient
}

type xxx_DefaultTypeLibClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultTypeLibClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultTypeLibClient) GetTypeInfoCount(ctx context.Context, in *GetTypeInfoCountRequest, opts ...dcerpc.CallOption) (*GetTypeInfoCountResponse, error) {
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
	out := &GetTypeInfoCountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeLibClient) GetTypeInfo(ctx context.Context, in *GetTypeInfoRequest, opts ...dcerpc.CallOption) (*GetTypeInfoResponse, error) {
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
	out := &GetTypeInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeLibClient) GetTypeInfoType(ctx context.Context, in *GetTypeInfoTypeRequest, opts ...dcerpc.CallOption) (*GetTypeInfoTypeResponse, error) {
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
	out := &GetTypeInfoTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeLibClient) GetTypeInfoOfGUID(ctx context.Context, in *GetTypeInfoOfGUIDRequest, opts ...dcerpc.CallOption) (*GetTypeInfoOfGUIDResponse, error) {
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
	out := &GetTypeInfoOfGUIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeLibClient) GetLibAttribute(ctx context.Context, in *GetLibAttributeRequest, opts ...dcerpc.CallOption) (*GetLibAttributeResponse, error) {
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
	out := &GetLibAttributeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeLibClient) GetTypeComp(ctx context.Context, in *GetTypeCompRequest, opts ...dcerpc.CallOption) (*GetTypeCompResponse, error) {
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
	out := &GetTypeCompResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeLibClient) GetDocumentation(ctx context.Context, in *GetDocumentationRequest, opts ...dcerpc.CallOption) (*GetDocumentationResponse, error) {
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
	out := &GetDocumentationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeLibClient) IsName(ctx context.Context, in *IsNameRequest, opts ...dcerpc.CallOption) (*IsNameResponse, error) {
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
	out := &IsNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeLibClient) FindName(ctx context.Context, in *FindNameRequest, opts ...dcerpc.CallOption) (*FindNameResponse, error) {
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
	out := &FindNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeLibClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultTypeLibClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultTypeLibClient) IPID(ctx context.Context, ipid *dcom.IPID) TypeLibClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultTypeLibClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewTypeLibClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (TypeLibClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(TypeLibSyntaxV0_0))...)
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
	return &xxx_DefaultTypeLibClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetTypeInfoCountOperation structure represents the GetTypeInfoCount operation
type xxx_GetTypeInfoCountOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	TypeInfoCount uint32         `idl:"name:pcTInfo" json:"type_info_count"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTypeInfoCountOperation) OpNum() int { return 3 }

func (o *xxx_GetTypeInfoCountOperation) OpName() string { return "/ITypeLib/v0/GetTypeInfoCount" }

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
	// pcTInfo {out} (1:{pointer=ref}*(1))(2:{alias=UINT}(uint32))
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
	// pcTInfo {out} (1:{pointer=ref}*(1))(2:{alias=UINT}(uint32))
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

func (o *GetTypeInfoCountRequest) xxx_ToOp(ctx context.Context) *xxx_GetTypeInfoCountOperation {
	if o == nil {
		return &xxx_GetTypeInfoCountOperation{}
	}
	return &xxx_GetTypeInfoCountOperation{
		This: o.This,
	}
}

func (o *GetTypeInfoCountRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTypeInfoCountOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetTypeInfoCountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
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
	// pcTInfo: MUST be set to the number of automation type descriptions contained in the
	// type information table of the automation type library.
	TypeInfoCount uint32 `idl:"name:pcTInfo" json:"type_info_count"`
	// Return: The GetTypeInfoCount return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTypeInfoCountResponse) xxx_ToOp(ctx context.Context) *xxx_GetTypeInfoCountOperation {
	if o == nil {
		return &xxx_GetTypeInfoCountOperation{}
	}
	return &xxx_GetTypeInfoCountOperation{
		That:          o.That,
		TypeInfoCount: o.TypeInfoCount,
		Return:        o.Return,
	}
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
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
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
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Index    uint32         `idl:"name:index" json:"index"`
	TypeInfo *oaut.TypeInfo `idl:"name:ppTInfo" json:"type_info"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTypeInfoOperation) OpNum() int { return 4 }

func (o *xxx_GetTypeInfoOperation) OpName() string { return "/ITypeLib/v0/GetTypeInfo" }

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
	// index {in} (1:{alias=UINT}(uint32))
	{
		if err := w.WriteData(o.Index); err != nil {
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
	// index {in} (1:{alias=UINT}(uint32))
	{
		if err := w.ReadData(&o.Index); err != nil {
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
	// index: MUST equal the ordinal position of the specified automation type description
	// within the type information table.
	Index uint32 `idl:"name:index" json:"index"`
}

func (o *GetTypeInfoRequest) xxx_ToOp(ctx context.Context) *xxx_GetTypeInfoOperation {
	if o == nil {
		return &xxx_GetTypeInfoOperation{}
	}
	return &xxx_GetTypeInfoOperation{
		This:  o.This,
		Index: o.Index,
	}
}

func (o *GetTypeInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTypeInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Index = op.Index
}
func (o *GetTypeInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
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

func (o *GetTypeInfoResponse) xxx_ToOp(ctx context.Context) *xxx_GetTypeInfoOperation {
	if o == nil {
		return &xxx_GetTypeInfoOperation{}
	}
	return &xxx_GetTypeInfoOperation{
		That:     o.That,
		TypeInfo: o.TypeInfo,
		Return:   o.Return,
	}
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
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetTypeInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTypeInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTypeInfoTypeOperation structure represents the GetTypeInfoType operation
type xxx_GetTypeInfoTypeOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Index  uint32         `idl:"name:index" json:"index"`
	Kind   oaut.TypeKind  `idl:"name:pTKind" json:"kind"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTypeInfoTypeOperation) OpNum() int { return 5 }

func (o *xxx_GetTypeInfoTypeOperation) OpName() string { return "/ITypeLib/v0/GetTypeInfoType" }

func (o *xxx_GetTypeInfoTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeInfoTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// index {in} (1:{alias=UINT}(uint32))
	{
		if err := w.WriteData(o.Index); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeInfoTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// index {in} (1:{alias=UINT}(uint32))
	{
		if err := w.ReadData(&o.Index); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeInfoTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeInfoTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pTKind {out} (1:{pointer=ref}*(1))(2:{v1_enum, alias=TYPEKIND}(enum))
	{
		if err := w.WriteData(uint32(o.Kind)); err != nil {
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

func (o *xxx_GetTypeInfoTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pTKind {out} (1:{pointer=ref}*(1))(2:{v1_enum, alias=TYPEKIND}(enum))
	{
		if err := w.ReadData((*uint32)(&o.Kind)); err != nil {
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

// GetTypeInfoTypeRequest structure represents the GetTypeInfoType operation request
type GetTypeInfoTypeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// index: MUST equal the ordinal position of the specified automation type description
	// within the type information table.
	Index uint32 `idl:"name:index" json:"index"`
}

func (o *GetTypeInfoTypeRequest) xxx_ToOp(ctx context.Context) *xxx_GetTypeInfoTypeOperation {
	if o == nil {
		return &xxx_GetTypeInfoTypeOperation{}
	}
	return &xxx_GetTypeInfoTypeOperation{
		This:  o.This,
		Index: o.Index,
	}
}

func (o *GetTypeInfoTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTypeInfoTypeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Index = op.Index
}
func (o *GetTypeInfoTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetTypeInfoTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTypeInfoTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTypeInfoTypeResponse structure represents the GetTypeInfoType operation response
type GetTypeInfoTypeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pTKind: MUST be set to the TYPEKIND value associated with the type description, as
	// specified in 2.2.17.
	Kind oaut.TypeKind `idl:"name:pTKind" json:"kind"`
	// Return: The GetTypeInfoType return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTypeInfoTypeResponse) xxx_ToOp(ctx context.Context) *xxx_GetTypeInfoTypeOperation {
	if o == nil {
		return &xxx_GetTypeInfoTypeOperation{}
	}
	return &xxx_GetTypeInfoTypeOperation{
		That:   o.That,
		Kind:   o.Kind,
		Return: o.Return,
	}
}

func (o *GetTypeInfoTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTypeInfoTypeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Kind = op.Kind
	o.Return = op.Return
}
func (o *GetTypeInfoTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetTypeInfoTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTypeInfoTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTypeInfoOfGUIDOperation structure represents the GetTypeInfoOfGuid operation
type xxx_GetTypeInfoOfGUIDOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	GUID     *dtyp.GUID     `idl:"name:guid" json:"guid"`
	TypeInfo *oaut.TypeInfo `idl:"name:ppTInfo" json:"type_info"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTypeInfoOfGUIDOperation) OpNum() int { return 6 }

func (o *xxx_GetTypeInfoOfGUIDOperation) OpName() string { return "/ITypeLib/v0/GetTypeInfoOfGuid" }

func (o *xxx_GetTypeInfoOfGUIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeInfoOfGUIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// guid {in} (1:{alias=REFGUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID != nil {
			if err := o.GUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetTypeInfoOfGUIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// guid {in} (1:{alias=REFGUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID == nil {
			o.GUID = &dtyp.GUID{}
		}
		if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeInfoOfGUIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeInfoOfGUIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetTypeInfoOfGUIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// GetTypeInfoOfGUIDRequest structure represents the GetTypeInfoOfGuid operation request
type GetTypeInfoOfGUIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// guid: MUST be a GUID.
	GUID *dtyp.GUID `idl:"name:guid" json:"guid"`
}

func (o *GetTypeInfoOfGUIDRequest) xxx_ToOp(ctx context.Context) *xxx_GetTypeInfoOfGUIDOperation {
	if o == nil {
		return &xxx_GetTypeInfoOfGUIDOperation{}
	}
	return &xxx_GetTypeInfoOfGUIDOperation{
		This: o.This,
		GUID: o.GUID,
	}
}

func (o *GetTypeInfoOfGUIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTypeInfoOfGUIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.GUID = op.GUID
}
func (o *GetTypeInfoOfGUIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetTypeInfoOfGUIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTypeInfoOfGUIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTypeInfoOfGUIDResponse structure represents the GetTypeInfoOfGuid operation response
type GetTypeInfoOfGUIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppTInfo: MUST be set to an ITypeInfo server instance that represents the automation
	// type description associated with the specified GUID in the type information table
	// (see section 3.7) or to NULL. MUST be NULL if the value of guid is IID_NULL, or is
	// not associated with an automation type description.
	TypeInfo *oaut.TypeInfo `idl:"name:ppTInfo" json:"type_info"`
	// Return: The GetTypeInfoOfGuid return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTypeInfoOfGUIDResponse) xxx_ToOp(ctx context.Context) *xxx_GetTypeInfoOfGUIDOperation {
	if o == nil {
		return &xxx_GetTypeInfoOfGUIDOperation{}
	}
	return &xxx_GetTypeInfoOfGUIDOperation{
		That:     o.That,
		TypeInfo: o.TypeInfo,
		Return:   o.Return,
	}
}

func (o *GetTypeInfoOfGUIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTypeInfoOfGUIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TypeInfo = op.TypeInfo
	o.Return = op.Return
}
func (o *GetTypeInfoOfGUIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetTypeInfoOfGUIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTypeInfoOfGUIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetLibAttributeOperation structure represents the GetLibAttr operation
type xxx_GetLibAttributeOperation struct {
	This         *dcom.ORPCThis         `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat         `idl:"name:That" json:"that"`
	LibAttribute *oaut.TypeLibAttribute `idl:"name:ppTLibAttr" json:"lib_attribute"`
	_            uint32                 `idl:"name:pReserved"`
	Return       int32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_GetLibAttributeOperation) OpNum() int { return 7 }

func (o *xxx_GetLibAttributeOperation) OpName() string { return "/ITypeLib/v0/GetLibAttr" }

func (o *xxx_GetLibAttributeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLibAttributeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetLibAttributeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetLibAttributeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLibAttributeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppTLibAttr {out} (1:{pointer=ref}*(2))(2:{alias=LPTLIBATTR}*(1))(3:{alias=TLIBATTR}(struct))
	{
		if o.LibAttribute != nil {
			_ptr_ppTLibAttr := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.LibAttribute != nil {
					if err := o.LibAttribute.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.TypeLibAttribute{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.LibAttribute, _ptr_ppTLibAttr); err != nil {
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
	// pReserved {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		// reserved pReserved
		if err := w.WriteData(uint32(0)); err != nil {
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

func (o *xxx_GetLibAttributeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppTLibAttr {out} (1:{pointer=ref}*(2))(2:{alias=LPTLIBATTR,pointer=ref}*(1))(3:{alias=TLIBATTR}(struct))
	{
		_ptr_ppTLibAttr := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.LibAttribute == nil {
				o.LibAttribute = &oaut.TypeLibAttribute{}
			}
			if err := o.LibAttribute.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppTLibAttr := func(ptr interface{}) { o.LibAttribute = *ptr.(**oaut.TypeLibAttribute) }
		if err := w.ReadPointer(&o.LibAttribute, _s_ppTLibAttr, _ptr_ppTLibAttr); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pReserved {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		// reserved pReserved
		var _pReserved uint32
		if err := w.ReadData(&_pReserved); err != nil {
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

// GetLibAttributeRequest structure represents the GetLibAttr operation request
type GetLibAttributeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetLibAttributeRequest) xxx_ToOp(ctx context.Context) *xxx_GetLibAttributeOperation {
	if o == nil {
		return &xxx_GetLibAttributeOperation{}
	}
	return &xxx_GetLibAttributeOperation{
		This: o.This,
	}
}

func (o *GetLibAttributeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetLibAttributeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetLibAttributeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetLibAttributeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLibAttributeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetLibAttributeResponse structure represents the GetLibAttr operation response
type GetLibAttributeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppTLibAttr: MUST be set to a TLIBATTR (section 2.2.45) structure that specifies the
	// attributes declared with the Type library.
	LibAttribute *oaut.TypeLibAttribute `idl:"name:ppTLibAttr" json:"lib_attribute"`
	// Return: The GetLibAttr return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetLibAttributeResponse) xxx_ToOp(ctx context.Context) *xxx_GetLibAttributeOperation {
	if o == nil {
		return &xxx_GetLibAttributeOperation{}
	}
	return &xxx_GetLibAttributeOperation{
		That:         o.That,
		LibAttribute: o.LibAttribute,
		Return:       o.Return,
	}
}

func (o *GetLibAttributeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetLibAttributeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.LibAttribute = op.LibAttribute
	o.Return = op.Return
}
func (o *GetLibAttributeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetLibAttributeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLibAttributeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTypeCompOperation structure represents the GetTypeComp operation
type xxx_GetTypeCompOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Comp   *oaut.TypeComp `idl:"name:ppTComp" json:"comp"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTypeCompOperation) OpNum() int { return 8 }

func (o *xxx_GetTypeCompOperation) OpName() string { return "/ITypeLib/v0/GetTypeComp" }

func (o *xxx_GetTypeCompOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeCompOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetTypeCompOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetTypeCompOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeCompOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppTComp {out} (1:{pointer=ref}*(2)*(1))(2:{alias=ITypeComp}(interface))
	{
		if o.Comp != nil {
			_ptr_ppTComp := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Comp != nil {
					if err := o.Comp.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.TypeComp{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Comp, _ptr_ppTComp); err != nil {
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

func (o *xxx_GetTypeCompOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppTComp {out} (1:{pointer=ref}*(2)*(1))(2:{alias=ITypeComp}(interface))
	{
		_ptr_ppTComp := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Comp == nil {
				o.Comp = &oaut.TypeComp{}
			}
			if err := o.Comp.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppTComp := func(ptr interface{}) { o.Comp = *ptr.(**oaut.TypeComp) }
		if err := w.ReadPointer(&o.Comp, _s_ppTComp, _ptr_ppTComp); err != nil {
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

// GetTypeCompRequest structure represents the GetTypeComp operation request
type GetTypeCompRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetTypeCompRequest) xxx_ToOp(ctx context.Context) *xxx_GetTypeCompOperation {
	if o == nil {
		return &xxx_GetTypeCompOperation{}
	}
	return &xxx_GetTypeCompOperation{
		This: o.This,
	}
}

func (o *GetTypeCompRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTypeCompOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetTypeCompRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetTypeCompRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTypeCompOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTypeCompResponse structure represents the GetTypeComp operation response
type GetTypeCompResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppTComp: MUST be set to a reference to the ITypeComp server instance associated with
	// the ITypeInfo server (see section 3.5).
	//
	// ppTComp: MUST be set to a reference to the ITypeComp server instance associated with
	// the automation type library, or to NULL if the automation type library does not have
	// an associated ITypeComp server instance.
	Comp *oaut.TypeComp `idl:"name:ppTComp" json:"comp"`
	// Return: The GetTypeComp return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTypeCompResponse) xxx_ToOp(ctx context.Context) *xxx_GetTypeCompOperation {
	if o == nil {
		return &xxx_GetTypeCompOperation{}
	}
	return &xxx_GetTypeCompOperation{
		That:   o.That,
		Comp:   o.Comp,
		Return: o.Return,
	}
}

func (o *GetTypeCompResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTypeCompOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Comp = op.Comp
	o.Return = op.Return
}
func (o *GetTypeCompResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetTypeCompResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTypeCompOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDocumentationOperation structure represents the GetDocumentation operation
type xxx_GetDocumentationOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	Index        int32          `idl:"name:index" json:"index"`
	PointerFlags uint32         `idl:"name:refPtrFlags" json:"pointer_flags"`
	Name         *oaut.String   `idl:"name:pBstrName" json:"name"`
	DocString    *oaut.String   `idl:"name:pBstrDocString" json:"doc_string"`
	HelpContext  uint32         `idl:"name:pdwHelpContext" json:"help_context"`
	HelpFile     *oaut.String   `idl:"name:pBstrHelpFile" json:"help_file"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDocumentationOperation) OpNum() int { return 9 }

func (o *xxx_GetDocumentationOperation) OpName() string { return "/ITypeLib/v0/GetDocumentation" }

func (o *xxx_GetDocumentationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDocumentationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// index {in} (1:{alias=INT}(int32))
	{
		if err := w.WriteData(o.Index); err != nil {
			return err
		}
	}
	// refPtrFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PointerFlags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDocumentationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// index {in} (1:{alias=INT}(int32))
	{
		if err := w.ReadData(&o.Index); err != nil {
			return err
		}
	}
	// refPtrFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PointerFlags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDocumentationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDocumentationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pBstrName {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Name != nil {
			_ptr_pBstrName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Name != nil {
					if err := o.Name.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Name, _ptr_pBstrName); err != nil {
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
	// pBstrDocString {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.DocString != nil {
			_ptr_pBstrDocString := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DocString != nil {
					if err := o.DocString.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DocString, _ptr_pBstrDocString); err != nil {
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
	// pdwHelpContext {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.HelpContext); err != nil {
			return err
		}
	}
	// pBstrHelpFile {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.HelpFile != nil {
			_ptr_pBstrHelpFile := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.HelpFile != nil {
					if err := o.HelpFile.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.HelpFile, _ptr_pBstrHelpFile); err != nil {
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

func (o *xxx_GetDocumentationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pBstrName {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pBstrName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Name == nil {
				o.Name = &oaut.String{}
			}
			if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pBstrName := func(ptr interface{}) { o.Name = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Name, _s_pBstrName, _ptr_pBstrName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pBstrDocString {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pBstrDocString := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DocString == nil {
				o.DocString = &oaut.String{}
			}
			if err := o.DocString.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pBstrDocString := func(ptr interface{}) { o.DocString = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.DocString, _s_pBstrDocString, _ptr_pBstrDocString); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pdwHelpContext {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.HelpContext); err != nil {
			return err
		}
	}
	// pBstrHelpFile {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pBstrHelpFile := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.HelpFile == nil {
				o.HelpFile = &oaut.String{}
			}
			if err := o.HelpFile.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pBstrHelpFile := func(ptr interface{}) { o.HelpFile = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.HelpFile, _s_pBstrHelpFile, _ptr_pBstrHelpFile); err != nil {
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

// GetDocumentationRequest structure represents the GetDocumentation operation request
type GetDocumentationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// index: MUST equal the ordinal position of an automation type description in the type
	// information table, or –1. If index is –1, the values of pBstrName, pBstrDocString,
	// pdwHelpContext, and pBstrHelpFile MUST correspond to the attributes declared with
	// the Type library, as specified in section 2.2.49.2. Otherwise, they MUST correspond
	// to the attributes declared with the specified type.
	Index int32 `idl:"name:index" json:"index"`
	// refPtrFlags: MUST be a combination of the bit flags specified in the following table,
	// or 0.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|                                    |                                                                                  |
	//	|               VALUE                |                                     MEANING                                      |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| TYPEINFO_NameArg 0x00000001        | MUST specify that the client is interested in the actual pBstrName [out]         |
	//	|                                    | argument.                                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| TYPEINFO_DocStringArg 0x00000002   | MUST specify that the client is interested in the actual pBstrDocString [out]    |
	//	|                                    | argument.                                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| TYPEINFO_HelpContextArg 0x00000004 | MUST specify that the client is interested in the actual pdwHelpContext [out]    |
	//	|                                    | argument.                                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| TYPEINFO_HelpFileArg 0x00000008    | MUST specify that the client is interested in the actual pBstrHelpFile [out]     |
	//	|                                    | argument.                                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// refPtrFlags: MUST be a combination of 0, or the bit flags specified in the following
	// table.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|                                   |                                                                                  |
	//	|               VALUE               |                                     MEANING                                      |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| TYPELIB_NameArg 0x00000001        | MUST specify that the client is interested in the actual pBstrName [out]         |
	//	|                                   | argument.                                                                        |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| TYPELIB_DocStringArg 0x00000002   | MUST specify that the client is interested in the actual pBstrDocString [out]    |
	//	|                                   | argument.                                                                        |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| TYPELIB_HelpContextArg 0x00000004 | MUST specify that the client is interested in the actual pdwHelpContext [out]    |
	//	|                                   | argument.                                                                        |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| TYPELIB_HelpFileArg 0x00000008    | MUST specify that the client is interested in the actual pBstrHelpFile [out]     |
	//	|                                   | argument.                                                                        |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	PointerFlags uint32 `idl:"name:refPtrFlags" json:"pointer_flags"`
}

func (o *GetDocumentationRequest) xxx_ToOp(ctx context.Context) *xxx_GetDocumentationOperation {
	if o == nil {
		return &xxx_GetDocumentationOperation{}
	}
	return &xxx_GetDocumentationOperation{
		This:         o.This,
		Index:        o.Index,
		PointerFlags: o.PointerFlags,
	}
}

func (o *GetDocumentationRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDocumentationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Index = op.Index
	o.PointerFlags = op.PointerFlags
}
func (o *GetDocumentationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetDocumentationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDocumentationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDocumentationResponse structure represents the GetDocumentation operation response
type GetDocumentationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pBstrName: If the TYPEINFO_NameArg bit flag is set in refPtrFlags,pBstrName MUST
	// be set to a BSTR that contains the name of the type or specified type member. Othernwise,
	// pBstrName MUST be set to a NULL BSTR.
	//
	// pBstrName: MUST be set to a BSTR that contains the name of the specified type or
	// Type library if the TYPELIB_NameArg bit flag is set in refPtrFlags. MUST be set to
	// a NULL BSTR otherwise.
	Name *oaut.String `idl:"name:pBstrName" json:"name"`
	// pBstrDocString: MUST be set to the documentation string that was associated with
	// the type or specified type member using the [helpstring] attribute (see section 2.2.49.2),
	// if the TYPEINFO_DocStringArg bit flag is set in refPtrFlags. MAY be set to an implementation-specific
	// string<59> if no [helpstring] attribute is specified. MUST be set to a NULL BSTR
	// otherwise.
	//
	// pBstrDocString: MUST be set to the documentation string that was associated with
	// the specified type or Type library using the [helpstring] attribute (see section
	// 2.2.49.2), if the TYPELIB_DocStringArg bit flag is set in refPtrFlags. MAY be set
	// to an implementation-specific string<61> if no [helpstring] attribute is specified.
	// MUST be set to a NULL BSTR otherwise.
	DocString *oaut.String `idl:"name:pBstrDocString" json:"doc_string"`
	// pdwHelpContext: MUST be set to the value that was associated with the type or specified
	// type member using the [helpcontext] attribute (see section 2.2.49.2), if the TYPEINFO_HelpContextArg
	// bit flag is set in refPtrFlags. MUST be set to 0 otherwise.
	//
	// pdwHelpContext: MUST be set to the value that was associated with the specified type
	// or Type library using the [helpcontext] attribute (see section 2.2.49.2), if the
	// TYPELIB_HelpContextArg bit flag is set in refPtrFlags. MUST be set to 0 otherwise.
	HelpContext uint32 `idl:"name:pdwHelpContext" json:"help_context"`
	// pBstrHelpFile: MUST be set to the documentation string that was associated with the
	// type or specified type member using the [helpfile] attribute (see section 2.2.49.2),
	// if the TYPEINFO_HelpFileArg bit flag is set in refPtrFlags. MUST be set to a NULL
	// BSTR otherwise.
	//
	// pBstrHelpFile: MUST be set to the documentation string that was associated with the
	// specified type or Type library using the [helpfile] attribute (see section 2.2.49.2),
	// if the TYPELIB_HelpFileArg bit flag is set in refPtrFlags. MUST be set to a NULL
	// BSTR otherwise.
	HelpFile *oaut.String `idl:"name:pBstrHelpFile" json:"help_file"`
	// Return: The GetDocumentation return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDocumentationResponse) xxx_ToOp(ctx context.Context) *xxx_GetDocumentationOperation {
	if o == nil {
		return &xxx_GetDocumentationOperation{}
	}
	return &xxx_GetDocumentationOperation{
		That:        o.That,
		Name:        o.Name,
		DocString:   o.DocString,
		HelpContext: o.HelpContext,
		HelpFile:    o.HelpFile,
		Return:      o.Return,
	}
}

func (o *GetDocumentationResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDocumentationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Name = op.Name
	o.DocString = op.DocString
	o.HelpContext = op.HelpContext
	o.HelpFile = op.HelpFile
	o.Return = op.Return
}
func (o *GetDocumentationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetDocumentationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDocumentationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IsNameOperation structure represents the IsName operation
type xxx_IsNameOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	NameBuffer    string         `idl:"name:szNameBuf" json:"name_buffer"`
	HashValue     uint32         `idl:"name:lHashVal" json:"hash_value"`
	Name          bool           `idl:"name:pfName" json:"name"`
	NameInLibrary *oaut.String   `idl:"name:pBstrNameInLibrary" json:"name_in_library"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_IsNameOperation) OpNum() int { return 10 }

func (o *xxx_IsNameOperation) OpName() string { return "/ITypeLib/v0/IsName" }

func (o *xxx_IsNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// szNameBuf {in} (1:{string, alias=LPOLESTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.NameBuffer); err != nil {
			return err
		}
	}
	// lHashVal {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.HashValue); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// szNameBuf {in} (1:{string, alias=LPOLESTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.NameBuffer); err != nil {
			return err
		}
	}
	// lHashVal {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.HashValue); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pfName {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		if !o.Name {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	// pBstrNameInLibrary {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.NameInLibrary != nil {
			_ptr_pBstrNameInLibrary := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NameInLibrary != nil {
					if err := o.NameInLibrary.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NameInLibrary, _ptr_pBstrNameInLibrary); err != nil {
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

func (o *xxx_IsNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pfName {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		var _bName int32
		if err := w.ReadData(&_bName); err != nil {
			return err
		}
		o.Name = _bName != 0
	}
	// pBstrNameInLibrary {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pBstrNameInLibrary := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NameInLibrary == nil {
				o.NameInLibrary = &oaut.String{}
			}
			if err := o.NameInLibrary.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pBstrNameInLibrary := func(ptr interface{}) { o.NameInLibrary = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.NameInLibrary, _s_pBstrNameInLibrary, _ptr_pBstrNameInLibrary); err != nil {
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

// IsNameRequest structure represents the IsName operation request
type IsNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// szNameBuf: MUST be set to a string to be tested if it matches the name of a type
	// or type member.
	NameBuffer string `idl:"name:szNameBuf" json:"name_buffer"`
	// lHashVal: MUST be either the hash value that corresponds to the value of szNameBuf
	// (as specified in section 2.2.51) or 0.
	HashValue uint32 `idl:"name:lHashVal" json:"hash_value"`
}

func (o *IsNameRequest) xxx_ToOp(ctx context.Context) *xxx_IsNameOperation {
	if o == nil {
		return &xxx_IsNameOperation{}
	}
	return &xxx_IsNameOperation{
		This:       o.This,
		NameBuffer: o.NameBuffer,
		HashValue:  o.HashValue,
	}
}

func (o *IsNameRequest) xxx_FromOp(ctx context.Context, op *xxx_IsNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.NameBuffer = op.NameBuffer
	o.HashValue = op.HashValue
}
func (o *IsNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *IsNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsNameResponse structure represents the IsName operation response
type IsNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pfName: MUST be set to TRUE if the specified string matches the name of a type or
	// member that is contained in the automation type library (see section 3.11.4.9) or
	// its binding context (see section 3.5.4.1.1.1) according to the string-matching criteria,
	// as specified in section 2.2.50. Otherwise, MUST be set to FALSE.
	Name bool `idl:"name:pfName" json:"name"`
	// pBstrNameInLibrary: MUST be set to a string whose value matches the value of szNameBuf
	// according to the string-matching rules (as specified in section 2.2.50), if pfName
	// is TRUE. MUST be set to a NULL BSTR if pfName is FALSE.
	NameInLibrary *oaut.String `idl:"name:pBstrNameInLibrary" json:"name_in_library"`
	// Return: The IsName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *IsNameResponse) xxx_ToOp(ctx context.Context) *xxx_IsNameOperation {
	if o == nil {
		return &xxx_IsNameOperation{}
	}
	return &xxx_IsNameOperation{
		That:          o.That,
		Name:          o.Name,
		NameInLibrary: o.NameInLibrary,
		Return:        o.Return,
	}
}

func (o *IsNameResponse) xxx_FromOp(ctx context.Context, op *xxx_IsNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Name = op.Name
	o.NameInLibrary = op.NameInLibrary
	o.Return = op.Return
}
func (o *IsNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *IsNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FindNameOperation structure represents the FindName operation
type xxx_FindNameOperation struct {
	This          *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat   `idl:"name:That" json:"that"`
	NameBuffer    string           `idl:"name:szNameBuf" json:"name_buffer"`
	HashValue     uint32           `idl:"name:lHashVal" json:"hash_value"`
	TypeInfo      []*oaut.TypeInfo `idl:"name:ppTInfo;size_is:(pcFound);length_is:(pcFound)" json:"type_info"`
	MemberIDs     []int32          `idl:"name:rgMemId;size_is:(pcFound);length_is:(pcFound)" json:"member_ids"`
	FoundCount    uint16           `idl:"name:pcFound" json:"found_count"`
	NameInLibrary *oaut.String     `idl:"name:pBstrNameInLibrary" json:"name_in_library"`
	Return        int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_FindNameOperation) OpNum() int { return 11 }

func (o *xxx_FindNameOperation) OpName() string { return "/ITypeLib/v0/FindName" }

func (o *xxx_FindNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FindNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// szNameBuf {in} (1:{string, alias=LPOLESTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.NameBuffer); err != nil {
			return err
		}
	}
	// lHashVal {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.HashValue); err != nil {
			return err
		}
	}
	// pcFound {in, out} (1:{pointer=ref}*(1))(2:{alias=USHORT}(uint16))
	{
		if err := w.WriteData(o.FoundCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FindNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// szNameBuf {in} (1:{string, alias=LPOLESTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.NameBuffer); err != nil {
			return err
		}
	}
	// lHashVal {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.HashValue); err != nil {
			return err
		}
	}
	// pcFound {in, out} (1:{pointer=ref}*(1))(2:{alias=USHORT}(uint16))
	{
		if err := w.ReadData(&o.FoundCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FindNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.TypeInfo != nil && o.FoundCount == 0 {
		o.FoundCount = uint16(len(o.TypeInfo))
	}
	if o.TypeInfo != nil && o.FoundCount == 0 {
		o.FoundCount = uint16(len(o.TypeInfo))
	}
	if o.MemberIDs != nil && o.FoundCount == 0 {
		o.FoundCount = uint16(len(o.MemberIDs))
	}
	if o.MemberIDs != nil && o.FoundCount == 0 {
		o.FoundCount = uint16(len(o.MemberIDs))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FindNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppTInfo {out} (1:{pointer=ref}*(1)[dim:0,size_is=pcFound,length_is=pcFound]*(1))(2:{alias=ITypeInfo}(interface))
	{
		dimSize1 := uint64(o.FoundCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.FoundCount)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.TypeInfo {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.TypeInfo[i1] != nil {
				_ptr_ppTInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if o.TypeInfo[i1] != nil {
						if err := o.TypeInfo[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&oaut.TypeInfo{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.TypeInfo[i1], _ptr_ppTInfo); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.TypeInfo); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// rgMemId {out} (1:{pointer=ref}*(1))(2:{alias=MEMBERID, names=LONG}[dim:0,size_is=pcFound,length_is=pcFound](int32))
	{
		dimSize1 := uint64(o.FoundCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.FoundCount)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.MemberIDs {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.MemberIDs[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.MemberIDs); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		}
	}
	// pcFound {in, out} (1:{pointer=ref}*(1))(2:{alias=USHORT}(uint16))
	{
		if err := w.WriteData(o.FoundCount); err != nil {
			return err
		}
	}
	// pBstrNameInLibrary {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.NameInLibrary != nil {
			_ptr_pBstrNameInLibrary := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NameInLibrary != nil {
					if err := o.NameInLibrary.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NameInLibrary, _ptr_pBstrNameInLibrary); err != nil {
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

func (o *xxx_FindNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppTInfo {out} (1:{pointer=ref}*(1)[dim:0,size_is=pcFound,length_is=pcFound]*(1))(2:{alias=ITypeInfo}(interface))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.TypeInfo", sizeInfo[0])
		}
		o.TypeInfo = make([]*oaut.TypeInfo, sizeInfo[0])
		for i1 := range o.TypeInfo {
			i1 := i1
			_ptr_ppTInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.TypeInfo[i1] == nil {
					o.TypeInfo[i1] = &oaut.TypeInfo{}
				}
				if err := o.TypeInfo[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppTInfo := func(ptr interface{}) { o.TypeInfo[i1] = *ptr.(**oaut.TypeInfo) }
			if err := w.ReadPointer(&o.TypeInfo[i1], _s_ppTInfo, _ptr_ppTInfo); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// rgMemId {out} (1:{pointer=ref}*(1))(2:{alias=MEMBERID, names=LONG}[dim:0,size_is=pcFound,length_is=pcFound](int32))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.MemberIDs", sizeInfo[0])
		}
		o.MemberIDs = make([]int32, sizeInfo[0])
		for i1 := range o.MemberIDs {
			i1 := i1
			if err := w.ReadData(&o.MemberIDs[i1]); err != nil {
				return err
			}
		}
	}
	// pcFound {in, out} (1:{pointer=ref}*(1))(2:{alias=USHORT}(uint16))
	{
		if err := w.ReadData(&o.FoundCount); err != nil {
			return err
		}
	}
	// pBstrNameInLibrary {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pBstrNameInLibrary := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NameInLibrary == nil {
				o.NameInLibrary = &oaut.String{}
			}
			if err := o.NameInLibrary.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pBstrNameInLibrary := func(ptr interface{}) { o.NameInLibrary = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.NameInLibrary, _s_pBstrNameInLibrary, _ptr_pBstrNameInLibrary); err != nil {
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

// FindNameRequest structure represents the FindName operation request
type FindNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// szNameBuf: MUST be a string.
	NameBuffer string `idl:"name:szNameBuf" json:"name_buffer"`
	// lHashVal: MUST be either the hash value corresponding to the value of szNameBuf (as
	// specified in section 2.2.51), or 0.
	HashValue uint32 `idl:"name:lHashVal" json:"hash_value"`
	// pcFound: The client MUST set pcFound to the maximum number of matches that can be
	// returned. The server MUST set pcFound to the number of elements in the ppTInfo and
	// rgMemId arrays.
	FoundCount uint16 `idl:"name:pcFound" json:"found_count"`
}

func (o *FindNameRequest) xxx_ToOp(ctx context.Context) *xxx_FindNameOperation {
	if o == nil {
		return &xxx_FindNameOperation{}
	}
	return &xxx_FindNameOperation{
		This:       o.This,
		NameBuffer: o.NameBuffer,
		HashValue:  o.HashValue,
		FoundCount: o.FoundCount,
	}
}

func (o *FindNameRequest) xxx_FromOp(ctx context.Context, op *xxx_FindNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.NameBuffer = op.NameBuffer
	o.HashValue = op.HashValue
	o.FoundCount = op.FoundCount
}
func (o *FindNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *FindNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FindNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FindNameResponse structure represents the FindName operation response
type FindNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppTInfo: MUST be set to an array of references to ITypeInfo server instances (see
	// section 3.7.4). Each entry of ppTInfo MUST correspond to a type whose name matches
	// the value of szNameBuf according to the string matching criteria (as specified in
	// section 2.2.50) or that contains a member whose name matches the value of szNameBuf.
	//
	// The array MUST be empty if there are no types or method or data members of types
	// defined in the automation scope whose names match the value of szNameBuf. Otherwise,
	// the array MUST contain one entry for each named nonparameter element defined in the
	// automation scope whose name matches szNameBuf. The array MAY contain entries for
	// matching types or type members that are referenced, but not defined in the automation
	// scope. <62>
	//
	// If szNameBuf matches the name of a dual interface or one of its members, the corresponding
	// entry in ppTInfo MUST refer to the partner dispinterface and MUST NOT refer to the
	// partner interface.
	TypeInfo []*oaut.TypeInfo `idl:"name:ppTInfo;size_is:(pcFound);length_is:(pcFound)" json:"type_info"`
	// rgMemId: MUST be set to an array of MEMBERIDs (see section 2.2.35) corresponding
	// to the ITypeInfo instances in the ppTInfo array. For each entry in the ppTInfo array,
	// the corresponding value in the rgMemId array MUST specify the MEMBERID of the type
	// member whose name matches the value of szNameBuf, or MEMBERID_NIL to specify that
	// the name of the type matches the value of szNameBuf.
	MemberIDs []int32 `idl:"name:rgMemId;size_is:(pcFound);length_is:(pcFound)" json:"member_ids"`
	// pcFound: The client MUST set pcFound to the maximum number of matches that can be
	// returned. The server MUST set pcFound to the number of elements in the ppTInfo and
	// rgMemId arrays.
	FoundCount uint16 `idl:"name:pcFound" json:"found_count"`
	// pBstrNameInLibrary: MUST be set to a string whose value matches the value of szNameBuf
	// according to the string-matching rules (as specified in section 2.2.50), if the ppTInfo
	// array is not empty. MUST be set to a NULL BSTR otherwise.
	NameInLibrary *oaut.String `idl:"name:pBstrNameInLibrary" json:"name_in_library"`
	// Return: The FindName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *FindNameResponse) xxx_ToOp(ctx context.Context) *xxx_FindNameOperation {
	if o == nil {
		return &xxx_FindNameOperation{}
	}
	return &xxx_FindNameOperation{
		That:          o.That,
		TypeInfo:      o.TypeInfo,
		MemberIDs:     o.MemberIDs,
		FoundCount:    o.FoundCount,
		NameInLibrary: o.NameInLibrary,
		Return:        o.Return,
	}
}

func (o *FindNameResponse) xxx_FromOp(ctx context.Context, op *xxx_FindNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TypeInfo = op.TypeInfo
	o.MemberIDs = op.MemberIDs
	o.FoundCount = op.FoundCount
	o.NameInLibrary = op.NameInLibrary
	o.Return = op.Return
}
func (o *FindNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *FindNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FindNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
