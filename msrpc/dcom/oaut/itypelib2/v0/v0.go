package itypelib2

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
	itypelib "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/itypelib/v0"
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
	_ = itypelib.GoPackage
	_ = dtyp.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/oaut"
)

var (
	// ITypeLib2 interface identifier 00020411-0000-0000-c000-000000000046
	TypeLib2IID = &dcom.IID{Data1: 0x00020411, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
	// Syntax UUID
	TypeLib2SyntaxUUID = &uuid.UUID{TimeLow: 0x20411, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xc0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x46}}
	// Syntax ID
	TypeLib2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: TypeLib2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// ITypeLib2 interface.
type TypeLib2Client interface {

	// ITypeLib retrieval method.
	TypeLib() itypelib.TypeLibClient

	// The GetCustData method retrieves the value of a custom data item associated with
	// the automation type library.
	//
	// The GetCustData method retrieves the value of a custom data item associated with
	// the type.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	GetCustomData(context.Context, *GetCustomDataRequest, ...dcerpc.CallOption) (*GetCustomDataResponse, error)

	// The GetLibStatistics method returns statistics about the unique names in the automation
	// type library.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	GetLibStatistics(context.Context, *GetLibStatisticsRequest, ...dcerpc.CallOption) (*GetLibStatisticsResponse, error)

	// The GetDocumentation2 method retrieves values associated with the automation type
	// library.
	//
	// The GetDocumentation2 method retrieves values associated with a type member.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	GetDocumentation2(context.Context, *GetDocumentation2Request, ...dcerpc.CallOption) (*GetDocumentation2Response, error)

	// The GetAllCustData method retrieves the values of all custom data items associated
	// with the automation type library.
	//
	// The GetAllCustData method retrieves all the custom data items associated with the
	// automation type description.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	GetAllCustomData(context.Context, *GetAllCustomDataRequest, ...dcerpc.CallOption) (*GetAllCustomDataResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) TypeLib2Client
}

type xxx_DefaultTypeLib2Client struct {
	itypelib.TypeLibClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultTypeLib2Client) TypeLib() itypelib.TypeLibClient {
	return o.TypeLibClient
}

func (o *xxx_DefaultTypeLib2Client) GetCustomData(ctx context.Context, in *GetCustomDataRequest, opts ...dcerpc.CallOption) (*GetCustomDataResponse, error) {
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
	out := &GetCustomDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeLib2Client) GetLibStatistics(ctx context.Context, in *GetLibStatisticsRequest, opts ...dcerpc.CallOption) (*GetLibStatisticsResponse, error) {
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
	out := &GetLibStatisticsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeLib2Client) GetDocumentation2(ctx context.Context, in *GetDocumentation2Request, opts ...dcerpc.CallOption) (*GetDocumentation2Response, error) {
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
	out := &GetDocumentation2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeLib2Client) GetAllCustomData(ctx context.Context, in *GetAllCustomDataRequest, opts ...dcerpc.CallOption) (*GetAllCustomDataResponse, error) {
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
	out := &GetAllCustomDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeLib2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultTypeLib2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultTypeLib2Client) IPID(ctx context.Context, ipid *dcom.IPID) TypeLib2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultTypeLib2Client{
		TypeLibClient: o.TypeLibClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewTypeLib2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (TypeLib2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(TypeLib2SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := itypelib.NewTypeLibClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultTypeLib2Client{
		TypeLibClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetCustomDataOperation structure represents the GetCustData operation
type xxx_GetCustomDataOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	GUID     *dtyp.GUID     `idl:"name:guid" json:"guid"`
	VarValue *oaut.Variant  `idl:"name:pVarVal" json:"var_value"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCustomDataOperation) OpNum() int { return 13 }

func (o *xxx_GetCustomDataOperation) OpName() string { return "/ITypeLib2/v0/GetCustData" }

func (o *xxx_GetCustomDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCustomDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetCustomDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetCustomDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCustomDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pVarVal {out} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.VarValue != nil {
			_ptr_pVarVal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.VarValue != nil {
					if err := o.VarValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.VarValue, _ptr_pVarVal); err != nil {
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

func (o *xxx_GetCustomDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pVarVal {out} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pVarVal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.VarValue == nil {
				o.VarValue = &oaut.Variant{}
			}
			if err := o.VarValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pVarVal := func(ptr interface{}) { o.VarValue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.VarValue, _s_pVarVal, _ptr_pVarVal); err != nil {
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

// GetCustomDataRequest structure represents the GetCustData operation request
type GetCustomDataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// guid: MUST be the GUID associated with the custom data item using the [custom] attribute,
	// as specified in section 2.2.49.2.
	//
	// guid: MUST be a GUID associated with the custom data item.
	GUID *dtyp.GUID `idl:"name:guid" json:"guid"`
}

func (o *GetCustomDataRequest) xxx_ToOp(ctx context.Context, op *xxx_GetCustomDataOperation) *xxx_GetCustomDataOperation {
	if op == nil {
		op = &xxx_GetCustomDataOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.GUID = op.GUID
	return op
}

func (o *GetCustomDataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCustomDataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.GUID = op.GUID
}
func (o *GetCustomDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetCustomDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCustomDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCustomDataResponse structure represents the GetCustData operation response
type GetCustomDataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pVarVal: MUST be set to the value of the custom data item, or VT_EMPTY if there is
	// no custom data item associated with the specified GUID.
	//
	// pVarVal: MUST be set to the value associated with the GUID using the [custom] attribute
	// (as specified in section 2.2.49.3), or VT_EMPTY if the type does not have a value
	// associated with the GUID.
	VarValue *oaut.Variant `idl:"name:pVarVal" json:"var_value"`
	// Return: The GetCustData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCustomDataResponse) xxx_ToOp(ctx context.Context, op *xxx_GetCustomDataOperation) *xxx_GetCustomDataOperation {
	if op == nil {
		op = &xxx_GetCustomDataOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.VarValue = op.VarValue
	o.Return = op.Return
	return op
}

func (o *GetCustomDataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCustomDataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.VarValue = op.VarValue
	o.Return = op.Return
}
func (o *GetCustomDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetCustomDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCustomDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetLibStatisticsOperation structure represents the GetLibStatistics operation
type xxx_GetLibStatisticsOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	UniqueNamesCount  uint32         `idl:"name:pcUniqueNames" json:"unique_names_count"`
	UniqueNamesLength uint32         `idl:"name:pcchUniqueNames" json:"unique_names_length"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetLibStatisticsOperation) OpNum() int { return 14 }

func (o *xxx_GetLibStatisticsOperation) OpName() string { return "/ITypeLib2/v0/GetLibStatistics" }

func (o *xxx_GetLibStatisticsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLibStatisticsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetLibStatisticsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetLibStatisticsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLibStatisticsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pcUniqueNames {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.UniqueNamesCount); err != nil {
			return err
		}
	}
	// pcchUniqueNames {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.UniqueNamesLength); err != nil {
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

func (o *xxx_GetLibStatisticsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pcUniqueNames {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.UniqueNamesCount); err != nil {
			return err
		}
	}
	// pcchUniqueNames {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.UniqueNamesLength); err != nil {
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

// GetLibStatisticsRequest structure represents the GetLibStatistics operation request
type GetLibStatisticsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetLibStatisticsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetLibStatisticsOperation) *xxx_GetLibStatisticsOperation {
	if op == nil {
		op = &xxx_GetLibStatisticsOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *GetLibStatisticsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetLibStatisticsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetLibStatisticsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetLibStatisticsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLibStatisticsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetLibStatisticsResponse structure represents the GetLibStatistics operation response
type GetLibStatisticsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pcUniqueNames: MUST be set to the number of unique names in the Type library.
	UniqueNamesCount uint32 `idl:"name:pcUniqueNames" json:"unique_names_count"`
	// pcchUniqueNames: MUST be set to the total length, in characters, of the unique names
	// in the library.
	UniqueNamesLength uint32 `idl:"name:pcchUniqueNames" json:"unique_names_length"`
	// Return: The GetLibStatistics return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetLibStatisticsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetLibStatisticsOperation) *xxx_GetLibStatisticsOperation {
	if op == nil {
		op = &xxx_GetLibStatisticsOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.UniqueNamesCount = op.UniqueNamesCount
	o.UniqueNamesLength = op.UniqueNamesLength
	o.Return = op.Return
	return op
}

func (o *GetLibStatisticsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetLibStatisticsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.UniqueNamesCount = op.UniqueNamesCount
	o.UniqueNamesLength = op.UniqueNamesLength
	o.Return = op.Return
}
func (o *GetLibStatisticsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetLibStatisticsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLibStatisticsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDocumentation2Operation structure represents the GetDocumentation2 operation
type xxx_GetDocumentation2Operation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	Index             int32          `idl:"name:index" json:"index"`
	LocaleID          uint32         `idl:"name:lcid" json:"locale_id"`
	PointerFlags      uint32         `idl:"name:refPtrFlags" json:"pointer_flags"`
	HelpString        *oaut.String   `idl:"name:pbstrHelpString" json:"help_string"`
	HelpStringContext uint32         `idl:"name:pdwHelpStringContext" json:"help_string_context"`
	HelpStringDLL     *oaut.String   `idl:"name:pbstrHelpStringDll" json:"help_string_dll"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDocumentation2Operation) OpNum() int { return 15 }

func (o *xxx_GetDocumentation2Operation) OpName() string { return "/ITypeLib2/v0/GetDocumentation2" }

func (o *xxx_GetDocumentation2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDocumentation2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lcid {in} (1:{alias=LCID, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.LocaleID); err != nil {
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

func (o *xxx_GetDocumentation2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lcid {in} (1:{alias=LCID, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.LocaleID); err != nil {
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

func (o *xxx_GetDocumentation2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDocumentation2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrHelpString {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.HelpString != nil {
			_ptr_pbstrHelpString := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.HelpString != nil {
					if err := o.HelpString.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.HelpString, _ptr_pbstrHelpString); err != nil {
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
	// pdwHelpStringContext {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.HelpStringContext); err != nil {
			return err
		}
	}
	// pbstrHelpStringDll {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.HelpStringDLL != nil {
			_ptr_pbstrHelpStringDll := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.HelpStringDLL != nil {
					if err := o.HelpStringDLL.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.HelpStringDLL, _ptr_pbstrHelpStringDll); err != nil {
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

func (o *xxx_GetDocumentation2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrHelpString {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrHelpString := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.HelpString == nil {
				o.HelpString = &oaut.String{}
			}
			if err := o.HelpString.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrHelpString := func(ptr interface{}) { o.HelpString = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.HelpString, _s_pbstrHelpString, _ptr_pbstrHelpString); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pdwHelpStringContext {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.HelpStringContext); err != nil {
			return err
		}
	}
	// pbstrHelpStringDll {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrHelpStringDll := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.HelpStringDLL == nil {
				o.HelpStringDLL = &oaut.String{}
			}
			if err := o.HelpStringDLL.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrHelpStringDll := func(ptr interface{}) { o.HelpStringDLL = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.HelpStringDLL, _s_pbstrHelpStringDll, _ptr_pbstrHelpStringDll); err != nil {
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

// GetDocumentation2Request structure represents the GetDocumentation2 operation request
type GetDocumentation2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// index: MUST be equal to the index of an automation type description or to –1. If
	// index is –1, the values of pBstrHelpString, pdwHelpStringContext, and pBstrHelpStringDll
	// MUST correspond to the attributes declared with the Type library as specified in
	// section 2.2.49.3. Otherwise, they MUST correspond to the attributes declared with
	// the specified type.
	Index int32 `idl:"name:index" json:"index"`
	// lcid: MUST be the locale ID of the specified type or type library.
	//
	// lcid: MUST be the Locale ID associated with the specified type member.
	LocaleID uint32 `idl:"name:lcid" json:"locale_id"`
	// refPtrFlags: MUST be 0, or a combination of the bit flags specified in the following
	// table.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|                                   |                                                                                  |
	//	|               VALUE               |                                     MEANING                                      |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| TYPELIB_HelpStringArg 0x00000001  | MUST specify that the client is interested in the actual pBstrHelpString [out]   |
	//	|                                   | argument.                                                                        |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| TYPELIB_HelpContextArg 0x00000002 | MUST specify that the client is interested in the actual pdwHelpStringContext    |
	//	|                                   | [out] argument.                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| TYPELIB_HelpFileArg 0x00000004    | MUST specify that the client is interested in the actual pBstrHelpStringDll      |
	//	|                                   | [out] argument.                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	PointerFlags uint32 `idl:"name:refPtrFlags" json:"pointer_flags"`
}

func (o *GetDocumentation2Request) xxx_ToOp(ctx context.Context, op *xxx_GetDocumentation2Operation) *xxx_GetDocumentation2Operation {
	if op == nil {
		op = &xxx_GetDocumentation2Operation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Index = op.Index
	o.LocaleID = op.LocaleID
	o.PointerFlags = op.PointerFlags
	return op
}

func (o *GetDocumentation2Request) xxx_FromOp(ctx context.Context, op *xxx_GetDocumentation2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Index = op.Index
	o.LocaleID = op.LocaleID
	o.PointerFlags = op.PointerFlags
}
func (o *GetDocumentation2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetDocumentation2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDocumentation2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDocumentation2Response structure represents the GetDocumentation2 operation response
type GetDocumentation2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbstrHelpString: MUST be set to an implementation-specific BSTR type<63> if the TYPELIB_HelpStringArg
	// bit flag is set in refPtrFlags. MUST be set to a NULL BSTR otherwise.
	//
	// pbstrHelpString: If the TYPEINFO_HelpStringContextArg and TYPEINFO_HelpStringDllArg
	// bit flags are set in refPtrFlags, pbstrHelpString MUST be set to an implementation-specific
	// BSTR<60> . Otherwise, MUST be set to a NULL BSTR.
	HelpString *oaut.String `idl:"name:pbstrHelpString" json:"help_string"`
	// pdwHelpStringContext: MUST be set to the value that was associated with the specified
	// type or type library using the [helpstringcontext] attribute (see section 2.2.49.2)
	// if the TYPELIB_HelpContextArg bit flag is set in refPtrFlags. MUST be set to 0 otherwise.
	//
	// pdwHelpStringContext: MUST be set to the value that was associated with the specified
	// type or type member using the [helpstringcontext] attribute (see IDL Automation Scope)
	// if the TYPEINFO_HelpStringContextArg bit flag is set in refPtrFlags. MUST be set
	// to 0 otherwise.
	HelpStringContext uint32 `idl:"name:pdwHelpStringContext" json:"help_string_context"`
	// pbstrHelpStringDll: MUST be set to the documentation string that was associated with
	// the specified type or type library using the [helpstringdll] attribute (see section
	// 2.2.49.2) if the TYPELIB_HelpFileArg bit flag is set in refPtrFlags. MUST be set
	// to a NULL BSTR otherwise.
	//
	// pbstrHelpStringDll: MUST be set to the documentation string that was associated with
	// the specified type or type member using the [helpstringdll] attribute (see IDL Automation
	// Scope) if the TYPEINFO_HelpStringDllArg bit flag is set in refPtrFlags. MUST be set
	// to a NULL BSTR otherwise.
	HelpStringDLL *oaut.String `idl:"name:pbstrHelpStringDll" json:"help_string_dll"`
	// Return: The GetDocumentation2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDocumentation2Response) xxx_ToOp(ctx context.Context, op *xxx_GetDocumentation2Operation) *xxx_GetDocumentation2Operation {
	if op == nil {
		op = &xxx_GetDocumentation2Operation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.HelpString = op.HelpString
	o.HelpStringContext = op.HelpStringContext
	o.HelpStringDLL = op.HelpStringDLL
	o.Return = op.Return
	return op
}

func (o *GetDocumentation2Response) xxx_FromOp(ctx context.Context, op *xxx_GetDocumentation2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.HelpString = op.HelpString
	o.HelpStringContext = op.HelpStringContext
	o.HelpStringDLL = op.HelpStringDLL
	o.Return = op.Return
}
func (o *GetDocumentation2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetDocumentation2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDocumentation2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAllCustomDataOperation structure represents the GetAllCustData operation
type xxx_GetAllCustomDataOperation struct {
	This       *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat   `idl:"name:That" json:"that"`
	CustomData *oaut.CustomData `idl:"name:pCustData" json:"custom_data"`
	Return     int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAllCustomDataOperation) OpNum() int { return 16 }

func (o *xxx_GetAllCustomDataOperation) OpName() string { return "/ITypeLib2/v0/GetAllCustData" }

func (o *xxx_GetAllCustomDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllCustomDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAllCustomDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAllCustomDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllCustomDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pCustData {out} (1:{pointer=ref}*(1))(2:{alias=CUSTDATA}(struct))
	{
		if o.CustomData != nil {
			if err := o.CustomData.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.CustomData{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetAllCustomDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pCustData {out} (1:{pointer=ref}*(1))(2:{alias=CUSTDATA}(struct))
	{
		if o.CustomData == nil {
			o.CustomData = &oaut.CustomData{}
		}
		if err := o.CustomData.UnmarshalNDR(ctx, w); err != nil {
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

// GetAllCustomDataRequest structure represents the GetAllCustData operation request
type GetAllCustomDataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetAllCustomDataRequest) xxx_ToOp(ctx context.Context, op *xxx_GetAllCustomDataOperation) *xxx_GetAllCustomDataOperation {
	if op == nil {
		op = &xxx_GetAllCustomDataOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *GetAllCustomDataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAllCustomDataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetAllCustomDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAllCustomDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAllCustomDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAllCustomDataResponse structure represents the GetAllCustData operation response
type GetAllCustomDataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pCustData: MUST be set to a CUSTDATA structure that contains an array of custom data
	// items, as specified in section 2.2.47. The structure's cCustData field MUST be set
	// to 0 and its prgCustData field set to NULL if there are no custom data items associated
	// with the automation type library.
	//
	// pCustData: MUST be set to a CUSTDATA structure that contains an array of custom data
	// items, as specified in section 2.2.47. The structure's cCustData field MUST be set
	// to 0 and its prgCustData field set to NULL, if there are no custom data items associated
	// with the automation type description.
	CustomData *oaut.CustomData `idl:"name:pCustData" json:"custom_data"`
	// Return: The GetAllCustData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAllCustomDataResponse) xxx_ToOp(ctx context.Context, op *xxx_GetAllCustomDataOperation) *xxx_GetAllCustomDataOperation {
	if op == nil {
		op = &xxx_GetAllCustomDataOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.CustomData = op.CustomData
	o.Return = op.Return
	return op
}

func (o *GetAllCustomDataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAllCustomDataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.CustomData = op.CustomData
	o.Return = op.Return
}
func (o *GetAllCustomDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAllCustomDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAllCustomDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
