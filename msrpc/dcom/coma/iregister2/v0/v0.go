package iregister2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	coma "github.com/oiweiwei/go-msrpc/msrpc/dcom/coma"
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
	_ = coma.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/coma"
)

var (
	// IRegister2 interface identifier 971668dc-c3fe-4ea1-9643-0c7230f494a1
	Register2IID = &dcom.IID{Data1: 0x971668dc, Data2: 0xc3fe, Data3: 0x4ea1, Data4: []byte{0x96, 0x43, 0x0c, 0x72, 0x30, 0xf4, 0x94, 0xa1}}
	// Syntax UUID
	Register2SyntaxUUID = &uuid.UUID{TimeLow: 0x971668dc, TimeMid: 0xc3fe, TimeHiAndVersion: 0x4ea1, ClockSeqHiAndReserved: 0x96, ClockSeqLow: 0x43, Node: [6]uint8{0xc, 0x72, 0x30, 0xf4, 0x94, 0xa1}}
	// Syntax ID
	Register2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: Register2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IRegister2 interface.
type Register2Client interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// This method is called by a client to create a component full configuration for an
	// existing component in an existing conglomeration in the global partition.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	CreateFullConfiguration(context.Context, *CreateFullConfigurationRequest, ...dcerpc.CallOption) (*CreateFullConfigurationResponse, error)

	// This method is called by a client to create a component legacy configuration for
	// an existing component in an existing conglomeration in the global partition.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	CreateLegacyConfiguration(context.Context, *CreateLegacyConfigurationRequest, ...dcerpc.CallOption) (*CreateLegacyConfigurationResponse, error)

	// This method is called by a client to convert an existing component legacy configuration
	// for a component into a component full configuration for that component.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result (as specified in [MS-ERREF] section 2.1) on failure. All failure results MUST
	// be treated identically.
	PromoteLegacyConfiguration(context.Context, *PromoteLegacyConfigurationRequest, ...dcerpc.CallOption) (*PromoteLegacyConfigurationResponse, error)

	// Opnum6NotUsedOnWire operation.
	// Opnum6NotUsedOnWire

	// Opnum7NotUsedOnWire operation.
	// Opnum7NotUsedOnWire

	// This method is called by a client to register the components in one or more modules
	// and to create component full configurations for those modules in an existing conglomeration.
	// This method supports conglomerations in any partition.
	//
	// Alternatively, this method can be called to verify modules without actually registering
	// the components. As a side effect, this method returns implementation-specific detailed
	// results of the registration or verification operation for informational purposes.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1,  on failure. All failure results
	// MUST be treated identically.
	RegisterModule2(context.Context, *RegisterModule2Request, ...dcerpc.CallOption) (*RegisterModule2Response, error)

	// Opnum9NotUsedOnWire operation.
	// Opnum9NotUsedOnWire

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) Register2Client
}

type xxx_DefaultRegister2Client struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultRegister2Client) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultRegister2Client) CreateFullConfiguration(ctx context.Context, in *CreateFullConfigurationRequest, opts ...dcerpc.CallOption) (*CreateFullConfigurationResponse, error) {
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
	out := &CreateFullConfigurationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRegister2Client) CreateLegacyConfiguration(ctx context.Context, in *CreateLegacyConfigurationRequest, opts ...dcerpc.CallOption) (*CreateLegacyConfigurationResponse, error) {
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
	out := &CreateLegacyConfigurationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRegister2Client) PromoteLegacyConfiguration(ctx context.Context, in *PromoteLegacyConfigurationRequest, opts ...dcerpc.CallOption) (*PromoteLegacyConfigurationResponse, error) {
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
	out := &PromoteLegacyConfigurationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRegister2Client) RegisterModule2(ctx context.Context, in *RegisterModule2Request, opts ...dcerpc.CallOption) (*RegisterModule2Response, error) {
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
	out := &RegisterModule2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRegister2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultRegister2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultRegister2Client) IPID(ctx context.Context, ipid *dcom.IPID) Register2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultRegister2Client{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewRegister2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Register2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(Register2SyntaxV0_0))...)
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
	return &xxx_DefaultRegister2Client{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_CreateFullConfigurationOperation structure represents the CreateFullConfiguration operation
type xxx_CreateFullConfigurationOperation struct {
	This                   *dcom.ORPCThis     `idl:"name:This" json:"this"`
	That                   *dcom.ORPCThat     `idl:"name:That" json:"that"`
	ConglomerationIDOrName string             `idl:"name:pwszConglomerationIdOrName;string" json:"conglomeration_id_or_name"`
	ClassIDOrProgID        string             `idl:"name:pwszCLSIDOrProgId;string" json:"class_id_or_prog_id"`
	CtComponentType        coma.ComponentType `idl:"name:ctComponentType" json:"ct_component_type"`
	Return                 int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateFullConfigurationOperation) OpNum() int { return 3 }

func (o *xxx_CreateFullConfigurationOperation) OpName() string {
	return "/IRegister2/v0/CreateFullConfiguration"
}

func (o *xxx_CreateFullConfigurationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateFullConfigurationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszConglomerationIdOrName {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ConglomerationIDOrName); err != nil {
			return err
		}
	}
	// pwszCLSIDOrProgId {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ClassIDOrProgID); err != nil {
			return err
		}
	}
	// ctComponentType {in} (1:{alias=eComponentType}(enum))
	{
		if err := w.WriteEnum(uint16(o.CtComponentType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateFullConfigurationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszConglomerationIdOrName {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ConglomerationIDOrName); err != nil {
			return err
		}
	}
	// pwszCLSIDOrProgId {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClassIDOrProgID); err != nil {
			return err
		}
	}
	// ctComponentType {in} (1:{alias=eComponentType}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.CtComponentType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateFullConfigurationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateFullConfigurationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateFullConfigurationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CreateFullConfigurationRequest structure represents the CreateFullConfiguration operation request
type CreateFullConfigurationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszConglomerationIdOrName: A string containing either the Curly Braced GUID String
	// Syntax ([MS-DTYP] section 2.3.4.3) representation of a conglomeration identifier
	// or the Name property (see section 3.1.1.3.3) of a conglomeration.
	ConglomerationIDOrName string `idl:"name:pwszConglomerationIdOrName;string" json:"conglomeration_id_or_name"`
	// pwszCLSIDOrProgId: A string containing either the Curly Braced GUID String Syntax
	// ([MS-DTYP] section 2.3.4.3) representation of a CLSID or the ProgID property (see
	// section 3.1.1.3.1) of a component.
	ClassIDOrProgID string `idl:"name:pwszCLSIDOrProgId;string" json:"class_id_or_prog_id"`
	// ctComponentType: An eComponentType (section 2.2.5) value to select the bitness of
	// the component, when there might be multiple bitnesses.
	CtComponentType coma.ComponentType `idl:"name:ctComponentType" json:"ct_component_type"`
}

func (o *CreateFullConfigurationRequest) xxx_ToOp(ctx context.Context) *xxx_CreateFullConfigurationOperation {
	if o == nil {
		return &xxx_CreateFullConfigurationOperation{}
	}
	return &xxx_CreateFullConfigurationOperation{
		This:                   o.This,
		ConglomerationIDOrName: o.ConglomerationIDOrName,
		ClassIDOrProgID:        o.ClassIDOrProgID,
		CtComponentType:        o.CtComponentType,
	}
}

func (o *CreateFullConfigurationRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateFullConfigurationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ConglomerationIDOrName = op.ConglomerationIDOrName
	o.ClassIDOrProgID = op.ClassIDOrProgID
	o.CtComponentType = op.CtComponentType
}
func (o *CreateFullConfigurationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateFullConfigurationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateFullConfigurationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateFullConfigurationResponse structure represents the CreateFullConfiguration operation response
type CreateFullConfigurationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CreateFullConfiguration return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateFullConfigurationResponse) xxx_ToOp(ctx context.Context) *xxx_CreateFullConfigurationOperation {
	if o == nil {
		return &xxx_CreateFullConfigurationOperation{}
	}
	return &xxx_CreateFullConfigurationOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *CreateFullConfigurationResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateFullConfigurationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CreateFullConfigurationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateFullConfigurationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateFullConfigurationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateLegacyConfigurationOperation structure represents the CreateLegacyConfiguration operation
type xxx_CreateLegacyConfigurationOperation struct {
	This                   *dcom.ORPCThis     `idl:"name:This" json:"this"`
	That                   *dcom.ORPCThat     `idl:"name:That" json:"that"`
	ConglomerationIDOrName string             `idl:"name:pwszConglomerationIdOrName;string" json:"conglomeration_id_or_name"`
	ClassIDOrProgID        string             `idl:"name:pwszCLSIDOrProgId;string" json:"class_id_or_prog_id"`
	CtComponentType        coma.ComponentType `idl:"name:ctComponentType" json:"ct_component_type"`
	Return                 int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateLegacyConfigurationOperation) OpNum() int { return 4 }

func (o *xxx_CreateLegacyConfigurationOperation) OpName() string {
	return "/IRegister2/v0/CreateLegacyConfiguration"
}

func (o *xxx_CreateLegacyConfigurationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateLegacyConfigurationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszConglomerationIdOrName {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ConglomerationIDOrName); err != nil {
			return err
		}
	}
	// pwszCLSIDOrProgId {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ClassIDOrProgID); err != nil {
			return err
		}
	}
	// ctComponentType {in} (1:{alias=eComponentType}(enum))
	{
		if err := w.WriteEnum(uint16(o.CtComponentType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateLegacyConfigurationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszConglomerationIdOrName {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ConglomerationIDOrName); err != nil {
			return err
		}
	}
	// pwszCLSIDOrProgId {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClassIDOrProgID); err != nil {
			return err
		}
	}
	// ctComponentType {in} (1:{alias=eComponentType}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.CtComponentType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateLegacyConfigurationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateLegacyConfigurationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateLegacyConfigurationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CreateLegacyConfigurationRequest structure represents the CreateLegacyConfiguration operation request
type CreateLegacyConfigurationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszConglomerationIdOrName: A string containing either the Curly Braced GUID String
	// Syntax ([MS-DTYP] section 2.3.4.3) representation of a conglomeration identifier
	// or the Name property (see section 3.1.1.3.3) of a conglomeration.
	ConglomerationIDOrName string `idl:"name:pwszConglomerationIdOrName;string" json:"conglomeration_id_or_name"`
	// pwszCLSIDOrProgId: A string containing either the Curly Braced GUID String Syntax
	// ([MS-DTYP] section 2.3.4.3) representation of a CLSID or the ProgID property (see
	// section 3.1.1.3.1) of a component.
	ClassIDOrProgID string `idl:"name:pwszCLSIDOrProgId;string" json:"class_id_or_prog_id"`
	// ctComponentType: An eComponentType (section 2.2.5) value to select the bitness of
	// the component when there might be multiple bitnesses. This value MUST NOT be eCT_UNKNOWN
	// or eCT_ANY.
	CtComponentType coma.ComponentType `idl:"name:ctComponentType" json:"ct_component_type"`
}

func (o *CreateLegacyConfigurationRequest) xxx_ToOp(ctx context.Context) *xxx_CreateLegacyConfigurationOperation {
	if o == nil {
		return &xxx_CreateLegacyConfigurationOperation{}
	}
	return &xxx_CreateLegacyConfigurationOperation{
		This:                   o.This,
		ConglomerationIDOrName: o.ConglomerationIDOrName,
		ClassIDOrProgID:        o.ClassIDOrProgID,
		CtComponentType:        o.CtComponentType,
	}
}

func (o *CreateLegacyConfigurationRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateLegacyConfigurationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ConglomerationIDOrName = op.ConglomerationIDOrName
	o.ClassIDOrProgID = op.ClassIDOrProgID
	o.CtComponentType = op.CtComponentType
}
func (o *CreateLegacyConfigurationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateLegacyConfigurationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateLegacyConfigurationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateLegacyConfigurationResponse structure represents the CreateLegacyConfiguration operation response
type CreateLegacyConfigurationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CreateLegacyConfiguration return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateLegacyConfigurationResponse) xxx_ToOp(ctx context.Context) *xxx_CreateLegacyConfigurationOperation {
	if o == nil {
		return &xxx_CreateLegacyConfigurationOperation{}
	}
	return &xxx_CreateLegacyConfigurationOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *CreateLegacyConfigurationResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateLegacyConfigurationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CreateLegacyConfigurationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateLegacyConfigurationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateLegacyConfigurationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PromoteLegacyConfigurationOperation structure represents the PromoteLegacyConfiguration operation
type xxx_PromoteLegacyConfigurationOperation struct {
	This                   *dcom.ORPCThis     `idl:"name:This" json:"this"`
	That                   *dcom.ORPCThat     `idl:"name:That" json:"that"`
	ConglomerationIDOrName string             `idl:"name:pwszConglomerationIdOrName;string" json:"conglomeration_id_or_name"`
	ClassIDOrProgID        string             `idl:"name:pwszCLSIDOrProgId;string" json:"class_id_or_prog_id"`
	CtComponentType        coma.ComponentType `idl:"name:ctComponentType" json:"ct_component_type"`
	Return                 int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_PromoteLegacyConfigurationOperation) OpNum() int { return 5 }

func (o *xxx_PromoteLegacyConfigurationOperation) OpName() string {
	return "/IRegister2/v0/PromoteLegacyConfiguration"
}

func (o *xxx_PromoteLegacyConfigurationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PromoteLegacyConfigurationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszConglomerationIdOrName {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ConglomerationIDOrName); err != nil {
			return err
		}
	}
	// pwszCLSIDOrProgId {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ClassIDOrProgID); err != nil {
			return err
		}
	}
	// ctComponentType {in} (1:{alias=eComponentType}(enum))
	{
		if err := w.WriteEnum(uint16(o.CtComponentType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PromoteLegacyConfigurationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszConglomerationIdOrName {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ConglomerationIDOrName); err != nil {
			return err
		}
	}
	// pwszCLSIDOrProgId {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClassIDOrProgID); err != nil {
			return err
		}
	}
	// ctComponentType {in} (1:{alias=eComponentType}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.CtComponentType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PromoteLegacyConfigurationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PromoteLegacyConfigurationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PromoteLegacyConfigurationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// PromoteLegacyConfigurationRequest structure represents the PromoteLegacyConfiguration operation request
type PromoteLegacyConfigurationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszConglomerationIdOrName:  A string containing either the Curly Braced GUID String
	// Syntax ([MS-DTYP] section 2.3.4.3) representation of a conglomeration identifier
	// or the Name property (see section 3.1.1.3.3) of a conglomeration.
	ConglomerationIDOrName string `idl:"name:pwszConglomerationIdOrName;string" json:"conglomeration_id_or_name"`
	// pwszCLSIDOrProgId:  A string containing either the Curly Braced GUID String Syntax
	// ([MS-DTYP] section 2.3.4.3) representation of a CLSID or the ProgID property (see
	// section 3.1.1.3.1) of a component.
	ClassIDOrProgID string `idl:"name:pwszCLSIDOrProgId;string" json:"class_id_or_prog_id"`
	// ctComponentType:  An eComponentType (section 2.2.5) value to select the bitness
	// of the component, when there might be multiple bitnesses.
	CtComponentType coma.ComponentType `idl:"name:ctComponentType" json:"ct_component_type"`
}

func (o *PromoteLegacyConfigurationRequest) xxx_ToOp(ctx context.Context) *xxx_PromoteLegacyConfigurationOperation {
	if o == nil {
		return &xxx_PromoteLegacyConfigurationOperation{}
	}
	return &xxx_PromoteLegacyConfigurationOperation{
		This:                   o.This,
		ConglomerationIDOrName: o.ConglomerationIDOrName,
		ClassIDOrProgID:        o.ClassIDOrProgID,
		CtComponentType:        o.CtComponentType,
	}
}

func (o *PromoteLegacyConfigurationRequest) xxx_FromOp(ctx context.Context, op *xxx_PromoteLegacyConfigurationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ConglomerationIDOrName = op.ConglomerationIDOrName
	o.ClassIDOrProgID = op.ClassIDOrProgID
	o.CtComponentType = op.CtComponentType
}
func (o *PromoteLegacyConfigurationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PromoteLegacyConfigurationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PromoteLegacyConfigurationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PromoteLegacyConfigurationResponse structure represents the PromoteLegacyConfiguration operation response
type PromoteLegacyConfigurationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The PromoteLegacyConfiguration return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PromoteLegacyConfigurationResponse) xxx_ToOp(ctx context.Context) *xxx_PromoteLegacyConfigurationOperation {
	if o == nil {
		return &xxx_PromoteLegacyConfigurationOperation{}
	}
	return &xxx_PromoteLegacyConfigurationOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *PromoteLegacyConfigurationResponse) xxx_FromOp(ctx context.Context, op *xxx_PromoteLegacyConfigurationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *PromoteLegacyConfigurationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PromoteLegacyConfigurationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PromoteLegacyConfigurationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RegisterModule2Operation structure represents the RegisterModule2 operation
type xxx_RegisterModule2Operation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	ConglomerationID  *dtyp.GUID     `idl:"name:ConglomerationIdentifier" json:"conglomeration_id"`
	PartitionID       *dtyp.GUID     `idl:"name:PartitionIdentifier" json:"partition_id"`
	Modules           []string       `idl:"name:ppModules;size_is:(cModules, );string" json:"modules"`
	ModulesCount      uint32         `idl:"name:cModules" json:"modules_count"`
	Flags             uint32         `idl:"name:dwFlags" json:"flags"`
	RequestedClassIDs []*dtyp.GUID   `idl:"name:pRequestedCLSIDs;size_is:(cRequested);pointer:unique" json:"requested_class_ids"`
	RequestedCount    uint32         `idl:"name:cRequested" json:"requested_count"`
	ModuleFlags       []uint32       `idl:"name:ppModuleFlags;size_is:(, cModules)" json:"module_flags"`
	ResultsCount      uint32         `idl:"name:pcResults" json:"results_count"`
	ResultClassIDs    []*dtyp.GUID   `idl:"name:ppResultCLSIDs;size_is:(, pcResults)" json:"result_class_ids"`
	ResultNames       []string       `idl:"name:ppResultNames;size_is:(, pcResults);string" json:"result_names"`
	ResultFlags       []uint32       `idl:"name:ppResultFlags;size_is:(, pcResults)" json:"result_flags"`
	ResultHRs         []int32        `idl:"name:ppResultHRs;size_is:(, pcResults)" json:"result_h_rs"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RegisterModule2Operation) OpNum() int { return 8 }

func (o *xxx_RegisterModule2Operation) OpName() string { return "/IRegister2/v0/RegisterModule2" }

func (o *xxx_RegisterModule2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Modules != nil && o.ModulesCount == 0 {
		o.ModulesCount = uint32(len(o.Modules))
	}
	if o.RequestedClassIDs != nil && o.RequestedCount == 0 {
		o.RequestedCount = uint32(len(o.RequestedClassIDs))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterModule2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ConglomerationIdentifier {in} (1:{alias=GUID}(struct))
	{
		if o.ConglomerationID != nil {
			if err := o.ConglomerationID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// PartitionIdentifier {in} (1:{alias=GUID}(struct))
	{
		if o.PartitionID != nil {
			if err := o.PartitionID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ppModules {in} (1:{string, pointer=ref}*(1))(2:{alias=LPWSTR}[dim:0,size_is=cModules]*(1)[dim:0,string,null](wchar))
	{
		dimSize1 := uint64(o.ModulesCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Modules {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Modules[i1] != "" {
				_ptr_ppModules := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if err := ndr.WriteUTF16NString(ctx, w, o.Modules[i1]); err != nil {
						return err
					}
					return nil
				})
				if err := w.WritePointer(&o.Modules[i1], _ptr_ppModules); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Modules); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// cModules {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ModulesCount); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pRequestedCLSIDs {in} (1:{pointer=unique}*(1))(2:{alias=GUID}[dim:0,size_is=cRequested](struct))
	{
		if o.RequestedClassIDs != nil || o.RequestedCount > 0 {
			_ptr_pRequestedCLSIDs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.RequestedCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.RequestedClassIDs {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.RequestedClassIDs[i1] != nil {
						if err := o.RequestedClassIDs[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.RequestedClassIDs); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RequestedClassIDs, _ptr_pRequestedCLSIDs); err != nil {
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
	// cRequested {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequestedCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterModule2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ConglomerationIdentifier {in} (1:{alias=GUID}(struct))
	{
		if o.ConglomerationID == nil {
			o.ConglomerationID = &dtyp.GUID{}
		}
		if err := o.ConglomerationID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// PartitionIdentifier {in} (1:{alias=GUID}(struct))
	{
		if o.PartitionID == nil {
			o.PartitionID = &dtyp.GUID{}
		}
		if err := o.PartitionID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ppModules {in} (1:{string, pointer=ref}*(1))(2:{alias=LPWSTR}[dim:0,size_is=cModules]*(1)[dim:0,string,null](wchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Modules", sizeInfo[0])
		}
		o.Modules = make([]string, sizeInfo[0])
		for i1 := range o.Modules {
			i1 := i1
			_ptr_ppModules := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if err := ndr.ReadUTF16NString(ctx, w, &o.Modules[i1]); err != nil {
					return err
				}
				return nil
			})
			_s_ppModules := func(ptr interface{}) { o.Modules[i1] = *ptr.(*string) }
			if err := w.ReadPointer(&o.Modules[i1], _s_ppModules, _ptr_ppModules); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cModules {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ModulesCount); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pRequestedCLSIDs {in} (1:{pointer=unique}*(1))(2:{alias=GUID}[dim:0,size_is=cRequested](struct))
	{
		_ptr_pRequestedCLSIDs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.RequestedClassIDs", sizeInfo[0])
			}
			o.RequestedClassIDs = make([]*dtyp.GUID, sizeInfo[0])
			for i1 := range o.RequestedClassIDs {
				i1 := i1
				if o.RequestedClassIDs[i1] == nil {
					o.RequestedClassIDs[i1] = &dtyp.GUID{}
				}
				if err := o.RequestedClassIDs[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pRequestedCLSIDs := func(ptr interface{}) { o.RequestedClassIDs = *ptr.(*[]*dtyp.GUID) }
		if err := w.ReadPointer(&o.RequestedClassIDs, _s_pRequestedCLSIDs, _ptr_pRequestedCLSIDs); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cRequested {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequestedCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterModule2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ResultClassIDs != nil && o.ResultsCount == 0 {
		o.ResultsCount = uint32(len(o.ResultClassIDs))
	}
	if o.ResultNames != nil && o.ResultsCount == 0 {
		o.ResultsCount = uint32(len(o.ResultNames))
	}
	if o.ResultFlags != nil && o.ResultsCount == 0 {
		o.ResultsCount = uint32(len(o.ResultFlags))
	}
	if o.ResultHRs != nil && o.ResultsCount == 0 {
		o.ResultsCount = uint32(len(o.ResultHRs))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterModule2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppModuleFlags {out} (1:{pointer=ref}*(2)*(1))(2:{alias=DWORD}[dim:0,size_is=cModules](uint32))
	{
		if o.ModuleFlags != nil || o.ModulesCount > 0 {
			_ptr_ppModuleFlags := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ModulesCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ModuleFlags {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.ModuleFlags[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.ModuleFlags); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint32(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ModuleFlags, _ptr_ppModuleFlags); err != nil {
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
	// pcResults {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ResultsCount); err != nil {
			return err
		}
	}
	// ppResultCLSIDs {out} (1:{pointer=ref}*(2)*(1))(2:{alias=GUID}[dim:0,size_is=pcResults](struct))
	{
		if o.ResultClassIDs != nil || o.ResultsCount > 0 {
			_ptr_ppResultCLSIDs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ResultsCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ResultClassIDs {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.ResultClassIDs[i1] != nil {
						if err := o.ResultClassIDs[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.ResultClassIDs); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResultClassIDs, _ptr_ppResultCLSIDs); err != nil {
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
	// ppResultNames {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pcResults]*(1)[dim:0,string,null](wchar))
	{
		if o.ResultNames != nil || o.ResultsCount > 0 {
			_ptr_ppResultNames := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ResultsCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ResultNames {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.ResultNames[i1] != "" {
						_ptr_ppResultNames := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
							if err := ndr.WriteUTF16NString(ctx, w, o.ResultNames[i1]); err != nil {
								return err
							}
							return nil
						})
						if err := w.WritePointer(&o.ResultNames[i1], _ptr_ppResultNames); err != nil {
							return err
						}
					} else {
						if err := w.WritePointer(nil); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.ResultNames); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResultNames, _ptr_ppResultNames); err != nil {
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
	// ppResultFlags {out} (1:{pointer=ref}*(2)*(1))(2:{alias=DWORD}[dim:0,size_is=pcResults](uint32))
	{
		if o.ResultFlags != nil || o.ResultsCount > 0 {
			_ptr_ppResultFlags := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ResultsCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ResultFlags {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.ResultFlags[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.ResultFlags); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint32(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResultFlags, _ptr_ppResultFlags); err != nil {
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
	// ppResultHRs {out} (1:{pointer=ref}*(2)*(1))(2:{alias=LONG}[dim:0,size_is=pcResults](int32))
	{
		if o.ResultHRs != nil || o.ResultsCount > 0 {
			_ptr_ppResultHRs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ResultsCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ResultHRs {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.ResultHRs[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.ResultHRs); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(int32(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResultHRs, _ptr_ppResultHRs); err != nil {
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

func (o *xxx_RegisterModule2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppModuleFlags {out} (1:{pointer=ref}*(2)*(1))(2:{alias=DWORD}[dim:0,size_is=cModules](uint32))
	{
		_ptr_ppModuleFlags := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ModuleFlags", sizeInfo[0])
			}
			o.ModuleFlags = make([]uint32, sizeInfo[0])
			for i1 := range o.ModuleFlags {
				i1 := i1
				if err := w.ReadData(&o.ModuleFlags[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppModuleFlags := func(ptr interface{}) { o.ModuleFlags = *ptr.(*[]uint32) }
		if err := w.ReadPointer(&o.ModuleFlags, _s_ppModuleFlags, _ptr_ppModuleFlags); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcResults {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ResultsCount); err != nil {
			return err
		}
	}
	// ppResultCLSIDs {out} (1:{pointer=ref}*(2)*(1))(2:{alias=GUID}[dim:0,size_is=pcResults](struct))
	{
		_ptr_ppResultCLSIDs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ResultClassIDs", sizeInfo[0])
			}
			o.ResultClassIDs = make([]*dtyp.GUID, sizeInfo[0])
			for i1 := range o.ResultClassIDs {
				i1 := i1
				if o.ResultClassIDs[i1] == nil {
					o.ResultClassIDs[i1] = &dtyp.GUID{}
				}
				if err := o.ResultClassIDs[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppResultCLSIDs := func(ptr interface{}) { o.ResultClassIDs = *ptr.(*[]*dtyp.GUID) }
		if err := w.ReadPointer(&o.ResultClassIDs, _s_ppResultCLSIDs, _ptr_ppResultCLSIDs); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppResultNames {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pcResults]*(1)[dim:0,string,null](wchar))
	{
		_ptr_ppResultNames := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ResultNames", sizeInfo[0])
			}
			o.ResultNames = make([]string, sizeInfo[0])
			for i1 := range o.ResultNames {
				i1 := i1
				_ptr_ppResultNames := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
					if err := ndr.ReadUTF16NString(ctx, w, &o.ResultNames[i1]); err != nil {
						return err
					}
					return nil
				})
				_s_ppResultNames := func(ptr interface{}) { o.ResultNames[i1] = *ptr.(*string) }
				if err := w.ReadPointer(&o.ResultNames[i1], _s_ppResultNames, _ptr_ppResultNames); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppResultNames := func(ptr interface{}) { o.ResultNames = *ptr.(*[]string) }
		if err := w.ReadPointer(&o.ResultNames, _s_ppResultNames, _ptr_ppResultNames); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppResultFlags {out} (1:{pointer=ref}*(2)*(1))(2:{alias=DWORD}[dim:0,size_is=pcResults](uint32))
	{
		_ptr_ppResultFlags := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ResultFlags", sizeInfo[0])
			}
			o.ResultFlags = make([]uint32, sizeInfo[0])
			for i1 := range o.ResultFlags {
				i1 := i1
				if err := w.ReadData(&o.ResultFlags[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppResultFlags := func(ptr interface{}) { o.ResultFlags = *ptr.(*[]uint32) }
		if err := w.ReadPointer(&o.ResultFlags, _s_ppResultFlags, _ptr_ppResultFlags); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppResultHRs {out} (1:{pointer=ref}*(2)*(1))(2:{alias=LONG}[dim:0,size_is=pcResults](int32))
	{
		_ptr_ppResultHRs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ResultHRs", sizeInfo[0])
			}
			o.ResultHRs = make([]int32, sizeInfo[0])
			for i1 := range o.ResultHRs {
				i1 := i1
				if err := w.ReadData(&o.ResultHRs[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppResultHRs := func(ptr interface{}) { o.ResultHRs = *ptr.(*[]int32) }
		if err := w.ReadPointer(&o.ResultHRs, _s_ppResultHRs, _ptr_ppResultHRs); err != nil {
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

// RegisterModule2Request structure represents the RegisterModule2 operation request
type RegisterModule2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ConglomerationIdentifier:  The conglomeration identifier of an existing conglomeration
	// on the server, in which the component full configurations are to be created, or against
	// which the modules are to be verified.
	ConglomerationID *dtyp.GUID `idl:"name:ConglomerationIdentifier" json:"conglomeration_id"`
	// PartitionIdentifier:  The partition identifier of the partition in which the conglomeration
	// identified by ConglomerationIdentifier is contained, or against which the modules
	// are to be verified.
	PartitionID *dtyp.GUID `idl:"name:PartitionIdentifier" json:"partition_id"`
	// ppModules:  An array of one or more strings, each of which is a path in UNC to a
	// file that the server will recognize as a module.
	Modules []string `idl:"name:ppModules;size_is:(cModules, );string" json:"modules"`
	// cModules:  The number of elements in ppModules.
	ModulesCount uint32 `idl:"name:cModules" json:"modules_count"`
	// dwFlags:  A combination of zero or more of the following flags.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|                                   |                                                                                  |
	//	|               FLAG                |                                     MEANING                                      |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| fREGISTER_VERIFYONLY 0x00000020   | The server SHOULD verify the modules, but MUST NOT actually register any         |
	//	|                                   | components.                                                                      |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| fREGISTER_EVENTCLASSES 0x00000400 | The server MUST configure the components registered by this operation as event   |
	//	|                                   | classes.                                                                         |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
	// pRequestedCLSIDs:  Either an array of one or more CLSIDs of components to be registered
	// (or verified), or NULL to specify that all components in the modules are to be registered
	// (or verified).
	RequestedClassIDs []*dtyp.GUID `idl:"name:pRequestedCLSIDs;size_is:(cRequested);pointer:unique" json:"requested_class_ids"`
	// cRequested:  The number of elements in pRequestedCLSIDs.
	RequestedCount uint32 `idl:"name:cRequested" json:"requested_count"`
}

func (o *RegisterModule2Request) xxx_ToOp(ctx context.Context) *xxx_RegisterModule2Operation {
	if o == nil {
		return &xxx_RegisterModule2Operation{}
	}
	return &xxx_RegisterModule2Operation{
		This:              o.This,
		ConglomerationID:  o.ConglomerationID,
		PartitionID:       o.PartitionID,
		Modules:           o.Modules,
		ModulesCount:      o.ModulesCount,
		Flags:             o.Flags,
		RequestedClassIDs: o.RequestedClassIDs,
		RequestedCount:    o.RequestedCount,
	}
}

func (o *RegisterModule2Request) xxx_FromOp(ctx context.Context, op *xxx_RegisterModule2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ConglomerationID = op.ConglomerationID
	o.PartitionID = op.PartitionID
	o.Modules = op.Modules
	o.ModulesCount = op.ModulesCount
	o.Flags = op.Flags
	o.RequestedClassIDs = op.RequestedClassIDs
	o.RequestedCount = op.RequestedCount
}
func (o *RegisterModule2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RegisterModule2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterModule2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RegisterModule2Response structure represents the RegisterModule2 operation response
type RegisterModule2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppModuleFlags:  A pointer to a variable that, upon successful completion, SHOULD
	// be set to an array of fModuleStatus (section 2.2.3) values representing the detailed
	// results of registration for the modules located by the paths in ppModules, in the
	// same order.
	ModuleFlags []uint32 `idl:"name:ppModuleFlags;size_is:(, cModules)" json:"module_flags"`
	// pcResults:  A pointer to a variable that, upon successful completion, MUST be set
	// to the number of result items, as specified later.
	ResultsCount uint32 `idl:"name:pcResults" json:"results_count"`
	// ppResultCLSIDs:  A pointer to a variable that, upon successful completion, MUST
	// be set to an array of GUID values, each being the CLSID of a result item, as specified
	// later.
	ResultClassIDs []*dtyp.GUID `idl:"name:ppResultCLSIDs;size_is:(, pcResults)" json:"result_class_ids"`
	// ppResultNames:  A pointer to a variable that, upon successful completion, MUST be
	// set to an array of string values, each being an implementation-specific<325> name
	// of a result item, as specified later, in the same order as ppResultCLSIDs.
	ResultNames []string `idl:"name:ppResultNames;size_is:(, pcResults);string" json:"result_names"`
	// ppResultFlags:  A pointer to a variable that, upon successful completion, MUST be
	// set to an array of fComponentStatus (section 2.2.4) values, each representing detailed
	// results for a result item, as specified later, in the same order as ppResultCLSIDs.
	ResultFlags []uint32 `idl:"name:ppResultFlags;size_is:(, pcResults)" json:"result_flags"`
	// ppResultHRs:  A pointer to a variable that, upon successful completion, MUST be
	// set to an array of values, each representing an HRESULT ([MS-ERREF] section 2.1)
	// for a result item, as specified later, in the same order as ppResultCLSIDs.
	ResultHRs []int32 `idl:"name:ppResultHRs;size_is:(, pcResults)" json:"result_h_rs"`
	// Return: The RegisterModule2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RegisterModule2Response) xxx_ToOp(ctx context.Context) *xxx_RegisterModule2Operation {
	if o == nil {
		return &xxx_RegisterModule2Operation{}
	}
	return &xxx_RegisterModule2Operation{
		That:           o.That,
		ModuleFlags:    o.ModuleFlags,
		ResultsCount:   o.ResultsCount,
		ResultClassIDs: o.ResultClassIDs,
		ResultNames:    o.ResultNames,
		ResultFlags:    o.ResultFlags,
		ResultHRs:      o.ResultHRs,
		Return:         o.Return,
	}
}

func (o *RegisterModule2Response) xxx_FromOp(ctx context.Context, op *xxx_RegisterModule2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ModuleFlags = op.ModuleFlags
	o.ResultsCount = op.ResultsCount
	o.ResultClassIDs = op.ResultClassIDs
	o.ResultNames = op.ResultNames
	o.ResultFlags = op.ResultFlags
	o.ResultHRs = op.ResultHRs
	o.Return = op.Return
}
func (o *RegisterModule2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RegisterModule2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterModule2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
