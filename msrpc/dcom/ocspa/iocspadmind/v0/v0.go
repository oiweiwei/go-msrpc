package iocspadmind

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
	ocspa "github.com/oiweiwei/go-msrpc/msrpc/dcom/ocspa"
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
	_ = ocspa.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/ocspa"
)

var (
	// IOCSPAdminD interface identifier 784b693d-95f3-420b-8126-365c098659f2
	OCSPAdminDIID = &dcom.IID{Data1: 0x784b693d, Data2: 0x95f3, Data3: 0x420b, Data4: []byte{0x81, 0x26, 0x36, 0x5c, 0x09, 0x86, 0x59, 0xf2}}
	// Syntax UUID
	OCSPAdminDSyntaxUUID = &uuid.UUID{TimeLow: 0x784b693d, TimeMid: 0x95f3, TimeHiAndVersion: 0x420b, ClockSeqHiAndReserved: 0x81, ClockSeqLow: 0x26, Node: [6]uint8{0x36, 0x5c, 0x9, 0x86, 0x59, 0xf2}}
	// Syntax ID
	OCSPAdminDSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: OCSPAdminDSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IOCSPAdminD interface.
type OCSPAdminDClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// This method retrieves the value of a responder property from the Online Responder
	// Service.
	GetOCSPProperty(context.Context, *GetOCSPPropertyRequest, ...dcerpc.CallOption) (*GetOCSPPropertyResponse, error)

	// This method configures the value of a responder property on the server.
	SetOCSPProperty(context.Context, *SetOCSPPropertyRequest, ...dcerpc.CallOption) (*SetOCSPPropertyResponse, error)

	// The GetCAConfigInformation method retrieves all the properties associated with a
	// particular revocation configuration.
	GetCAConfigInformation(context.Context, *GetCAConfigInformationRequest, ...dcerpc.CallOption) (*GetCAConfigInformationResponse, error)

	// This method sets all the properties for a particular revocation configuration.
	SetCAConfigInformation(context.Context, *SetCAConfigInformationRequest, ...dcerpc.CallOption) (*SetCAConfigInformationResponse, error)

	// The GetSecurity method is used to retrieve the security descriptor associated with
	// the responder.
	GetSecurity(context.Context, *GetSecurityRequest, ...dcerpc.CallOption) (*GetSecurityResponse, error)

	// The SetSecurity method is used to set the Online Responder Security, as defined in
	// the Abstract Data Model.
	SetSecurity(context.Context, *SetSecurityRequest, ...dcerpc.CallOption) (*SetSecurityResponse, error)

	// The GetSigningCertficates method retrieves a list of certificates available at the
	// responder machine that can be used to sign responses to OCSP requests regarding certificates
	// issued by the CA certificate specified.
	GetSigningCertificates(context.Context, *GetSigningCertificatesRequest, ...dcerpc.CallOption) (*GetSigningCertificatesResponse, error)

	// The GetHashAlgorithms method retrieves the list of hash algorithms available at the
	// responder that could be used along with the signing certificate associated with a
	// revocation configuration to sign OCSP responses.
	GetHashAlgorithms(context.Context, *GetHashAlgorithmsRequest, ...dcerpc.CallOption) (*GetHashAlgorithmsResponse, error)

	// The GetMyRoles method retrieves the Online Responder Roles [CIMC-PP] assigned to
	// the user that calls the method.
	GetMyRoles(context.Context, *GetMyRolesRequest, ...dcerpc.CallOption) (*GetMyRolesResponse, error)

	// This method queries the Online Responder Service to find out whether it is running.
	//
	// This method has no parameters.
	//
	// If the Online Responder Service is running, the server MUST return success (0) when
	// this method is invoked.
	Ping(context.Context, *PingRequest, ...dcerpc.CallOption) (*PingResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) OCSPAdminDClient
}

type xxx_DefaultOCSPAdminDClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultOCSPAdminDClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultOCSPAdminDClient) GetOCSPProperty(ctx context.Context, in *GetOCSPPropertyRequest, opts ...dcerpc.CallOption) (*GetOCSPPropertyResponse, error) {
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
	out := &GetOCSPPropertyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultOCSPAdminDClient) SetOCSPProperty(ctx context.Context, in *SetOCSPPropertyRequest, opts ...dcerpc.CallOption) (*SetOCSPPropertyResponse, error) {
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
	out := &SetOCSPPropertyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultOCSPAdminDClient) GetCAConfigInformation(ctx context.Context, in *GetCAConfigInformationRequest, opts ...dcerpc.CallOption) (*GetCAConfigInformationResponse, error) {
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
	out := &GetCAConfigInformationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultOCSPAdminDClient) SetCAConfigInformation(ctx context.Context, in *SetCAConfigInformationRequest, opts ...dcerpc.CallOption) (*SetCAConfigInformationResponse, error) {
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
	out := &SetCAConfigInformationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultOCSPAdminDClient) GetSecurity(ctx context.Context, in *GetSecurityRequest, opts ...dcerpc.CallOption) (*GetSecurityResponse, error) {
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
	out := &GetSecurityResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultOCSPAdminDClient) SetSecurity(ctx context.Context, in *SetSecurityRequest, opts ...dcerpc.CallOption) (*SetSecurityResponse, error) {
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
	out := &SetSecurityResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultOCSPAdminDClient) GetSigningCertificates(ctx context.Context, in *GetSigningCertificatesRequest, opts ...dcerpc.CallOption) (*GetSigningCertificatesResponse, error) {
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
	out := &GetSigningCertificatesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultOCSPAdminDClient) GetHashAlgorithms(ctx context.Context, in *GetHashAlgorithmsRequest, opts ...dcerpc.CallOption) (*GetHashAlgorithmsResponse, error) {
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
	out := &GetHashAlgorithmsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultOCSPAdminDClient) GetMyRoles(ctx context.Context, in *GetMyRolesRequest, opts ...dcerpc.CallOption) (*GetMyRolesResponse, error) {
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
	out := &GetMyRolesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultOCSPAdminDClient) Ping(ctx context.Context, in *PingRequest, opts ...dcerpc.CallOption) (*PingResponse, error) {
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
	out := &PingResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultOCSPAdminDClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultOCSPAdminDClient) IPID(ctx context.Context, ipid *dcom.IPID) OCSPAdminDClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultOCSPAdminDClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}
func NewOCSPAdminDClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (OCSPAdminDClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(OCSPAdminDSyntaxV0_0))...)
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
	return &xxx_DefaultOCSPAdminDClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetOCSPPropertyOperation structure represents the GetOCSPProperty operation
type xxx_GetOCSPPropertyOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	EntryName  *oaut.String   `idl:"name:bstrEntryName;pointer:ref" json:"entry_name"`
	EntryValue *oaut.Variant  `idl:"name:pEntryValue;pointer:ref" json:"entry_value"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetOCSPPropertyOperation) OpNum() int { return 3 }

func (o *xxx_GetOCSPPropertyOperation) OpName() string { return "/IOCSPAdminD/v0/GetOCSPProperty" }

func (o *xxx_GetOCSPPropertyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOCSPPropertyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrEntryName {in} (1:{pointer=ref, alias=BSTR, pointers=[unique]}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.EntryName != nil {
			if err := o.EntryName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetOCSPPropertyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrEntryName {in} (1:{pointer=ref, alias=BSTR, pointers=[unique]}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.EntryName == nil {
			o.EntryName = &oaut.String{}
		}
		if err := o.EntryName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOCSPPropertyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOCSPPropertyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pEntryValue {out} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.EntryValue != nil {
			_ptr_pEntryValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EntryValue != nil {
					if err := o.EntryValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EntryValue, _ptr_pEntryValue); err != nil {
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

func (o *xxx_GetOCSPPropertyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pEntryValue {out} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pEntryValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EntryValue == nil {
				o.EntryValue = &oaut.Variant{}
			}
			if err := o.EntryValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pEntryValue := func(ptr interface{}) { o.EntryValue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.EntryValue, _s_pEntryValue, _ptr_pEntryValue); err != nil {
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

// GetOCSPPropertyRequest structure represents the GetOCSPProperty operation request
type GetOCSPPropertyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrEntryName: A BSTR that specifies the name of the property to retrieve. The Unicode
	// string value SHOULD be one of the values listed in ResponderProperties or one of
	// the following values.
	//
	//	+---------------+----------------------------------------------------------------------------------+
	//	|   PROPERTY    |                                                                                  |
	//	|     NAME      |                                     MEANING                                      |
	//	|               |                                                                                  |
	//	+---------------+----------------------------------------------------------------------------------+
	//	+---------------+----------------------------------------------------------------------------------+
	//	| CAEntries     | A list of strings containing the RevocationConfigurationId corresponding to each |
	//	|               | configured revocation configuration in RevocationConfigurationList.              |
	//	+---------------+----------------------------------------------------------------------------------+
	//	| AllEntries    | A list of all the configured properties in the list ResponderProperties and all  |
	//	|               | the revocation configuration properties for all revocation configurations in     |
	//	|               | RevocationConfigurationList.                                                     |
	//	+---------------+----------------------------------------------------------------------------------+
	EntryName *oaut.String `idl:"name:bstrEntryName;pointer:ref" json:"entry_name"`
}

func (o *GetOCSPPropertyRequest) xxx_ToOp(ctx context.Context) *xxx_GetOCSPPropertyOperation {
	if o == nil {
		return &xxx_GetOCSPPropertyOperation{}
	}
	return &xxx_GetOCSPPropertyOperation{
		This:      o.This,
		EntryName: o.EntryName,
	}
}

func (o *GetOCSPPropertyRequest) xxx_FromOp(ctx context.Context, op *xxx_GetOCSPPropertyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.EntryName = op.EntryName
}
func (o *GetOCSPPropertyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetOCSPPropertyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetOCSPPropertyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetOCSPPropertyResponse structure represents the GetOCSPProperty operation response
type GetOCSPPropertyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pEntryValue: A pointer to a VARIANT. The data returned is the value of the property
	// referenced by bstrEntryName. See the following table for the processing rules that
	// apply to the bstrEntryName values. Other, vendor-defined bstrEntryName values, not
	// defined in the following table, MAY be used, as described in the processing rules
	// that follow the table.
	//
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	|        PROPERTY         |                             PROCESSING RULE FOR DATA                             |
	//	|          NAME           |                                     RETURNED                                     |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| AuditFilter             | The vt member of the VARIANT referenced by pEntryValue MUST be set to VT_I4, and |
	//	|                         | the lVal member MUST be either 0 or a bitwise OR of the following values. Flag   |
	//	|                         | value – Meaning 0x00000000 – Nothing is Audited. 0x00000001 – Audit start/stop   |
	//	|                         | of the service. 0x00000002 – Audit changes to the revocation configurations      |
	//	|                         | on the responder. 0x00000004 – Audit OCSP requests received by the responder.    |
	//	|                         | 0x00000008 – Audit changes to the security descriptor on the responder.          |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| ArrayController         | The vt member of the VARIANT referenced by pEntryValue SHOULD be set to VT_BSTR, |
	//	|                         | and the bstrVal member SHOULD be BSTR for the Unicode string value of the Domain |
	//	|                         | Name System (DNS) name of the machine designated as Array controller for the     |
	//	|                         | array of responder machines.                                                     |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| ArrayMembers            | The vt member of the VARIANT referenced by pEntryValue SHOULD be set to VT_ARRAY |
	//	|                         | | VT_BSTR, and the pArray member SHOULD reference a single dimension safearray.  |
	//	|                         | The number of elements of the safearray referenced by pArray SHOULD be equal     |
	//	|                         | to the number of machines running Online Responder Service with the same         |
	//	|                         | configuration information. For each machine, there SHOULD be an element in the   |
	//	|                         | safearray referenced by pArray containing the BSTR for Unicode string value of   |
	//	|                         | the FQDN of the machine.                                                         |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| NumOfThreads            | The vt member of the VARIANT referenced by pEntryValue MUST be set to VT_I4, and |
	//	|                         | the lVal member MUST be set to the maximum number of simultaneous OCSP requests  |
	//	|                         | [MS-OCSP] that can be served by the Online Responder Service. <5>                |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| MaxNumOfCacheEntries    | The vt member of the VARIANT referenced by pEntryValue MUST be set to VT_I4, and |
	//	|                         | the lVal member MUST be the maximum number of OCSP responses that can be cached  |
	//	|                         | by the responder.                                                                |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| CAEntries               | The vt member of the VARIANT referenced by pEntryValue SHOULD be set to          |
	//	|                         | VT_ARRAY | VT_BSTR, and the pArray member SHOULD reference a single dimension    |
	//	|                         | safearray. The number of elements of the safearray reference by pArray           |
	//	|                         | SHOULD be equal to the number of entries in RevocationConfigurationList.         |
	//	|                         | For each revocation configuration in RevocationConfigurationList, there          |
	//	|                         | SHOULD be an element containing the BSTR for the Unicode string value of the     |
	//	|                         | RevocationConfigurationId.                                                       |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| LogLevel                | The vt member of the VARIANT referenced by pEntryValue SHOULD be set to VT_I4,   |
	//	|                         | and the lVal member SHOULD be set to the integer value that specifies the level  |
	//	|                         | of information to be communicated to the system (application eventlog channel)   |
	//	|                         | as part of operations being performed on the service.<6>                         |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| Debug                   | The vt member of the VARIANT referenced by pEntryValue SHOULD be set to VT_I4,   |
	//	|                         | and the lVal member SHOULD be an integer value that specifies whether tracing    |
	//	|                         | for errors on the responder is enabled or not.<7>                                |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| EnrollPollInterval      | The vt member of the VARIANT referenced by pEntryValue SHOULD be set to VT_I4,   |
	//	|                         | and the lVal member SHOULD be set to the integer value that specifies the        |
	//	|                         | frequency (in number of hours) with which the responder will attempt to enroll   |
	//	|                         | for a signing certificate (for signing OCSP responses).<8>                       |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| RequestFlags            | The vt member of the VARIANT referenced by pEntryValue SHOULD be set to VT_I4,   |
	//	|                         | and the lVal member SHOULD be either 0 or the following value. Flag value –      |
	//	|                         | Meaning 0x00000001:Responder MUST reject OCSP requests that have signatures on   |
	//	|                         | them.                                                                            |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| MaxIncomingMessageSize  | The vt member of the VARIANT referenced by pEntryValue SHOULD be set to VT_I4,   |
	//	|                         | and the lVal member SHOULD be set to the integer value that specifies the        |
	//	|                         | maximum size of the OCSP request [MS-OCSP], in bytes, that is allowed to be      |
	//	|                         | processed on the server.                                                         |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| NumOfBackendConnections | The vt member of the VARIANT referenced by pEntryValue SHOULD be set to VT_I4,   |
	//	|                         | and the lVal member SHOULD be set to the integer value that specifies the        |
	//	|                         | maximum number of connections that can be created by the web server to the       |
	//	|                         | Online Responder Service. <9>                                                    |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| RefreshRate             | The vt member of the VARIANT referenced by pEntryValue SHOULD be set to VT_I4,   |
	//	|                         | and the lVal member SHOULD be set to the integer value that specifies the        |
	//	|                         | frequency (in number of milliseconds) with which the web server will attempt     |
	//	|                         | to contact the Online Responder Service to obtain the latest revocation          |
	//	|                         | configuration information.                                                       |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| MaxAge                  | The vt member of the VARIANT referenced by pEntryValue SHOULD be set to VT_I4,   |
	//	|                         | and the lVal member SHOULD be set to the integer value that specifies the value  |
	//	|                         | for the HTTP max-age cache-control directive [RFC2616] as part of the OCSP       |
	//	|                         | response.                                                                        |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| ISAPIDebug              | The vt member of the VARIANT referenced by pEntryValue SHOULD be set to VT_I4,   |
	//	|                         | and the lVal member SHOULD be set to the integer value that specifies whether    |
	//	|                         | the tracing for errors on the web server is enabled or not.<10>                  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| MaxNumOfRequestEntries  | The vt member of the VARIANT referenced by pEntryValue SHOULD be set to VT_I4,   |
	//	|                         | and the lVal member SHOULD be set to the integer value that specifies the        |
	//	|                         | maximum number of requests that can be included in the requestList field of the  |
	//	|                         | OCSPRequest structure ([RFC2560] section 4.1.1).<11>                             |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| AllEntries              | The vt member of the VARIANT MUST be set to VT_ARRAY | VT_VARIANT, and the       |
	//	|                         | pArray member MUST reference a two-dimensional safearray. The number of elements |
	//	|                         | in the second dimension (signifying the number of columns) of the safearray      |
	//	|                         | referenced by pArray MUST be 2. The number of elements in the first dimension    |
	//	|                         | (signifying the number of rows) of the safearray referenced by pArray MUST       |
	//	|                         | be set to the sum of the number of entries in ResponderProperties and the        |
	//	|                         | number of entries in the RevocationConfigurationList. For each property in       |
	//	|                         | ResponderProperties, the first column of the row MUST be a VARIANT with vt       |
	//	|                         | member as VT_BSTR and the bstrVal member MUST be BSTR for the Unicode string     |
	//	|                         | value of the name of the property. The second column of the row MUST be a        |
	//	|                         | VARIANT with the value defined in this table, corresponding to the name of the   |
	//	|                         | property. For each revocation configuration in RevocationConfigurationList,      |
	//	|                         | the first column of the row MUST be a VARIANT with vt member as VT_BSTR          |
	//	|                         | and the bstrVal member MUST be BSTR for the Unicode string value of              |
	//	|                         | RevocationConfigurationId. The second column of the row MUST be a VARIANT with   |
	//	|                         | the value defined in section 3.2.4.1.3.                                          |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//
	// The following additional processing rules apply:
	//
	// * If the value of bstrEntryName is not the same as one of the values specified in
	// the preceding list or of a vendor-defined property, or if the property with the same
	// name is not yet configured on the responder, the server MUST fail. The error code
	// SHOULD be 0x80070002.
	//
	// * If the value of bstrEntryName corresponds to a vendor-defined property, the server
	// MAY return the value as a VARIANT containing data of the type integer, string, date,
	// or binary object. Otherwise, for bstrEntryName values that do not correspond to the
	// previous list, the server responds as if the property were not yet configured on
	// the responder. <12> ( ff1216b3-6fb4-4be8-b256-fb7055b1e86d#Appendix_A_12 )
	EntryValue *oaut.Variant `idl:"name:pEntryValue;pointer:ref" json:"entry_value"`
	// Return: The GetOCSPProperty return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetOCSPPropertyResponse) xxx_ToOp(ctx context.Context) *xxx_GetOCSPPropertyOperation {
	if o == nil {
		return &xxx_GetOCSPPropertyOperation{}
	}
	return &xxx_GetOCSPPropertyOperation{
		That:       o.That,
		EntryValue: o.EntryValue,
		Return:     o.Return,
	}
}

func (o *GetOCSPPropertyResponse) xxx_FromOp(ctx context.Context, op *xxx_GetOCSPPropertyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.EntryValue = op.EntryValue
	o.Return = op.Return
}
func (o *GetOCSPPropertyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetOCSPPropertyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetOCSPPropertyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetOCSPPropertyOperation structure represents the SetOCSPProperty operation
type xxx_SetOCSPPropertyOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	EntryName  *oaut.String   `idl:"name:bstrEntryName;pointer:ref" json:"entry_name"`
	EntryValue *oaut.Variant  `idl:"name:pEntryValue;pointer:ref" json:"entry_value"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetOCSPPropertyOperation) OpNum() int { return 4 }

func (o *xxx_SetOCSPPropertyOperation) OpName() string { return "/IOCSPAdminD/v0/SetOCSPProperty" }

func (o *xxx_SetOCSPPropertyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOCSPPropertyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrEntryName {in} (1:{pointer=ref, alias=BSTR, pointers=[unique]}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.EntryName != nil {
			if err := o.EntryName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pEntryValue {in} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.EntryValue != nil {
			_ptr_pEntryValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EntryValue != nil {
					if err := o.EntryValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EntryValue, _ptr_pEntryValue); err != nil {
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
	return nil
}

func (o *xxx_SetOCSPPropertyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrEntryName {in} (1:{pointer=ref, alias=BSTR, pointers=[unique]}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.EntryName == nil {
			o.EntryName = &oaut.String{}
		}
		if err := o.EntryName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pEntryValue {in} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pEntryValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EntryValue == nil {
				o.EntryValue = &oaut.Variant{}
			}
			if err := o.EntryValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pEntryValue := func(ptr interface{}) { o.EntryValue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.EntryValue, _s_pEntryValue, _ptr_pEntryValue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOCSPPropertyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOCSPPropertyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetOCSPPropertyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetOCSPPropertyRequest structure represents the SetOCSPProperty operation request
type SetOCSPPropertyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrEntryName: A BSTR that specifies the name of the property to set. The Unicode
	// string value SHOULD be one of the property name values listed in ResponderProperties.
	EntryName *oaut.String `idl:"name:bstrEntryName;pointer:ref" json:"entry_name"`
	// pEntryValue: A pointer to VARIANT data.
	EntryValue *oaut.Variant `idl:"name:pEntryValue;pointer:ref" json:"entry_value"`
}

func (o *SetOCSPPropertyRequest) xxx_ToOp(ctx context.Context) *xxx_SetOCSPPropertyOperation {
	if o == nil {
		return &xxx_SetOCSPPropertyOperation{}
	}
	return &xxx_SetOCSPPropertyOperation{
		This:       o.This,
		EntryName:  o.EntryName,
		EntryValue: o.EntryValue,
	}
}

func (o *SetOCSPPropertyRequest) xxx_FromOp(ctx context.Context, op *xxx_SetOCSPPropertyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.EntryName = op.EntryName
	o.EntryValue = op.EntryValue
}
func (o *SetOCSPPropertyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetOCSPPropertyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetOCSPPropertyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetOCSPPropertyResponse structure represents the SetOCSPProperty operation response
type SetOCSPPropertyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetOCSPProperty return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetOCSPPropertyResponse) xxx_ToOp(ctx context.Context) *xxx_SetOCSPPropertyOperation {
	if o == nil {
		return &xxx_SetOCSPPropertyOperation{}
	}
	return &xxx_SetOCSPPropertyOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetOCSPPropertyResponse) xxx_FromOp(ctx context.Context, op *xxx_SetOCSPPropertyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetOCSPPropertyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetOCSPPropertyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetOCSPPropertyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetCAConfigInformationOperation structure represents the GetCAConfigInformation operation
type xxx_GetCAConfigInformationOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	CAID       *oaut.String   `idl:"name:bstrCAId;pointer:ref" json:"ca_id"`
	EntryValue *oaut.Variant  `idl:"name:pEntryValue;pointer:ref" json:"entry_value"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCAConfigInformationOperation) OpNum() int { return 5 }

func (o *xxx_GetCAConfigInformationOperation) OpName() string {
	return "/IOCSPAdminD/v0/GetCAConfigInformation"
}

func (o *xxx_GetCAConfigInformationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCAConfigInformationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrCAId {in} (1:{pointer=ref, alias=BSTR, pointers=[unique]}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.CAID != nil {
			if err := o.CAID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetCAConfigInformationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrCAId {in} (1:{pointer=ref, alias=BSTR, pointers=[unique]}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.CAID == nil {
			o.CAID = &oaut.String{}
		}
		if err := o.CAID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCAConfigInformationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCAConfigInformationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pEntryValue {out} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.EntryValue != nil {
			_ptr_pEntryValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EntryValue != nil {
					if err := o.EntryValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EntryValue, _ptr_pEntryValue); err != nil {
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

func (o *xxx_GetCAConfigInformationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pEntryValue {out} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pEntryValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EntryValue == nil {
				o.EntryValue = &oaut.Variant{}
			}
			if err := o.EntryValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pEntryValue := func(ptr interface{}) { o.EntryValue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.EntryValue, _s_pEntryValue, _ptr_pEntryValue); err != nil {
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

// GetCAConfigInformationRequest structure represents the GetCAConfigInformation operation request
type GetCAConfigInformationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrCAId: A BSTR that specifies the RevocationConfigurationId for the revocation
	// configuration whose properties are to be retrieved.
	CAID *oaut.String `idl:"name:bstrCAId;pointer:ref" json:"ca_id"`
}

func (o *GetCAConfigInformationRequest) xxx_ToOp(ctx context.Context) *xxx_GetCAConfigInformationOperation {
	if o == nil {
		return &xxx_GetCAConfigInformationOperation{}
	}
	return &xxx_GetCAConfigInformationOperation{
		This: o.This,
		CAID: o.CAID,
	}
}

func (o *GetCAConfigInformationRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCAConfigInformationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.CAID = op.CAID
}
func (o *GetCAConfigInformationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetCAConfigInformationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCAConfigInformationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCAConfigInformationResponse structure represents the GetCAConfigInformation operation response
type GetCAConfigInformationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pEntryValue: A pointer to a VARIANT data type that contains the names and values
	// of all configured revocation configuration properties.
	EntryValue *oaut.Variant `idl:"name:pEntryValue;pointer:ref" json:"entry_value"`
	// Return: The GetCAConfigInformation return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCAConfigInformationResponse) xxx_ToOp(ctx context.Context) *xxx_GetCAConfigInformationOperation {
	if o == nil {
		return &xxx_GetCAConfigInformationOperation{}
	}
	return &xxx_GetCAConfigInformationOperation{
		That:       o.That,
		EntryValue: o.EntryValue,
		Return:     o.Return,
	}
}

func (o *GetCAConfigInformationResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCAConfigInformationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.EntryValue = op.EntryValue
	o.Return = op.Return
}
func (o *GetCAConfigInformationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetCAConfigInformationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCAConfigInformationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetCAConfigInformationOperation structure represents the SetCAConfigInformation operation
type xxx_SetCAConfigInformationOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	CAID       *oaut.String   `idl:"name:bstrCAId;pointer:ref" json:"ca_id"`
	EntryValue *oaut.Variant  `idl:"name:pEntryValue;pointer:ref" json:"entry_value"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetCAConfigInformationOperation) OpNum() int { return 6 }

func (o *xxx_SetCAConfigInformationOperation) OpName() string {
	return "/IOCSPAdminD/v0/SetCAConfigInformation"
}

func (o *xxx_SetCAConfigInformationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCAConfigInformationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrCAId {in} (1:{pointer=ref, alias=BSTR, pointers=[unique]}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.CAID != nil {
			if err := o.CAID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pEntryValue {in} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.EntryValue != nil {
			_ptr_pEntryValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EntryValue != nil {
					if err := o.EntryValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EntryValue, _ptr_pEntryValue); err != nil {
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
	return nil
}

func (o *xxx_SetCAConfigInformationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrCAId {in} (1:{pointer=ref, alias=BSTR, pointers=[unique]}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.CAID == nil {
			o.CAID = &oaut.String{}
		}
		if err := o.CAID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pEntryValue {in} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pEntryValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EntryValue == nil {
				o.EntryValue = &oaut.Variant{}
			}
			if err := o.EntryValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pEntryValue := func(ptr interface{}) { o.EntryValue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.EntryValue, _s_pEntryValue, _ptr_pEntryValue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCAConfigInformationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCAConfigInformationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetCAConfigInformationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetCAConfigInformationRequest structure represents the SetCAConfigInformation operation request
type SetCAConfigInformationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrCAId: This is a BSTR that specifies the RevocationConfigurationId for the revocation
	// configuration whose properties are to be set.
	CAID *oaut.String `idl:"name:bstrCAId;pointer:ref" json:"ca_id"`
	// pEntryValue: This is a pointer to a VARIANT data type that contains the names and
	// values of the properties to set.
	EntryValue *oaut.Variant `idl:"name:pEntryValue;pointer:ref" json:"entry_value"`
}

func (o *SetCAConfigInformationRequest) xxx_ToOp(ctx context.Context) *xxx_SetCAConfigInformationOperation {
	if o == nil {
		return &xxx_SetCAConfigInformationOperation{}
	}
	return &xxx_SetCAConfigInformationOperation{
		This:       o.This,
		CAID:       o.CAID,
		EntryValue: o.EntryValue,
	}
}

func (o *SetCAConfigInformationRequest) xxx_FromOp(ctx context.Context, op *xxx_SetCAConfigInformationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.CAID = op.CAID
	o.EntryValue = op.EntryValue
}
func (o *SetCAConfigInformationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetCAConfigInformationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetCAConfigInformationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetCAConfigInformationResponse structure represents the SetCAConfigInformation operation response
type SetCAConfigInformationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetCAConfigInformation return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetCAConfigInformationResponse) xxx_ToOp(ctx context.Context) *xxx_SetCAConfigInformationOperation {
	if o == nil {
		return &xxx_SetCAConfigInformationOperation{}
	}
	return &xxx_SetCAConfigInformationOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetCAConfigInformationResponse) xxx_FromOp(ctx context.Context, op *xxx_SetCAConfigInformationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetCAConfigInformationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetCAConfigInformationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetCAConfigInformationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSecurityOperation structure represents the GetSecurity operation
type xxx_GetSecurityOperation struct {
	This   *dcom.ORPCThis           `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat           `idl:"name:That" json:"that"`
	SD     *ocspa.CertTransportBlob `idl:"name:pctbSD;pointer:ref" json:"sd"`
	Return int32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSecurityOperation) OpNum() int { return 7 }

func (o *xxx_GetSecurityOperation) OpName() string { return "/IOCSPAdminD/v0/GetSecurity" }

func (o *xxx_GetSecurityOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSecurityOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSecurityOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSecurityOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSecurityOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pctbSD {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.SD != nil {
			if err := o.SD.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ocspa.CertTransportBlob{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetSecurityOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pctbSD {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.SD == nil {
			o.SD = &ocspa.CertTransportBlob{}
		}
		if err := o.SD.UnmarshalNDR(ctx, w); err != nil {
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

// GetSecurityRequest structure represents the GetSecurity operation request
type GetSecurityRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSecurityRequest) xxx_ToOp(ctx context.Context) *xxx_GetSecurityOperation {
	if o == nil {
		return &xxx_GetSecurityOperation{}
	}
	return &xxx_GetSecurityOperation{
		This: o.This,
	}
}

func (o *GetSecurityRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSecurityOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSecurityRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetSecurityRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSecurityOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSecurityResponse structure represents the GetSecurity operation response
type GetSecurityResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pctbSD: This is a pointer to a CERTTRANSBLOB structure that contains the marshaled
	// Security Descriptor. Information on Security Descriptors is documented in [MS-DTYP]
	// section 2.4.6.<17>
	SD *ocspa.CertTransportBlob `idl:"name:pctbSD;pointer:ref" json:"sd"`
	// Return: The GetSecurity return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSecurityResponse) xxx_ToOp(ctx context.Context) *xxx_GetSecurityOperation {
	if o == nil {
		return &xxx_GetSecurityOperation{}
	}
	return &xxx_GetSecurityOperation{
		That:   o.That,
		SD:     o.SD,
		Return: o.Return,
	}
}

func (o *GetSecurityResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSecurityOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SD = op.SD
	o.Return = op.Return
}
func (o *GetSecurityResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetSecurityResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSecurityOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSecurityOperation structure represents the SetSecurity operation
type xxx_SetSecurityOperation struct {
	This   *dcom.ORPCThis           `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat           `idl:"name:That" json:"that"`
	SD     *ocspa.CertTransportBlob `idl:"name:pctbSD;pointer:ref" json:"sd"`
	Return int32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSecurityOperation) OpNum() int { return 8 }

func (o *xxx_SetSecurityOperation) OpName() string { return "/IOCSPAdminD/v0/SetSecurity" }

func (o *xxx_SetSecurityOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSecurityOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pctbSD {in} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.SD != nil {
			if err := o.SD.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ocspa.CertTransportBlob{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSecurityOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pctbSD {in} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.SD == nil {
			o.SD = &ocspa.CertTransportBlob{}
		}
		if err := o.SD.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSecurityOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSecurityOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSecurityOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetSecurityRequest structure represents the SetSecurity operation request
type SetSecurityRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pctbSD: A pointer to the CERTTRANSBLOB structure that contains the marshaled security
	// descriptor. Information on security descriptors is documented in [MS-DTYP] section
	// 2.4.6.
	SD *ocspa.CertTransportBlob `idl:"name:pctbSD;pointer:ref" json:"sd"`
}

func (o *SetSecurityRequest) xxx_ToOp(ctx context.Context) *xxx_SetSecurityOperation {
	if o == nil {
		return &xxx_SetSecurityOperation{}
	}
	return &xxx_SetSecurityOperation{
		This: o.This,
		SD:   o.SD,
	}
}

func (o *SetSecurityRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSecurityOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SD = op.SD
}
func (o *SetSecurityRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetSecurityRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSecurityOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSecurityResponse structure represents the SetSecurity operation response
type SetSecurityResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetSecurity return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSecurityResponse) xxx_ToOp(ctx context.Context) *xxx_SetSecurityOperation {
	if o == nil {
		return &xxx_SetSecurityOperation{}
	}
	return &xxx_SetSecurityOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetSecurityResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSecurityOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSecurityResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetSecurityResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSecurityOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSigningCertificatesOperation structure represents the GetSigningCertificates operation
type xxx_GetSigningCertificatesOperation struct {
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	CAVar               *oaut.Variant  `idl:"name:pCAVar;pointer:ref" json:"ca_var"`
	SigningCertificates *oaut.Variant  `idl:"name:pSigningCertificates;pointer:ref" json:"signing_certificates"`
	Return              int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSigningCertificatesOperation) OpNum() int { return 9 }

func (o *xxx_GetSigningCertificatesOperation) OpName() string {
	return "/IOCSPAdminD/v0/GetSigningCertificates"
}

func (o *xxx_GetSigningCertificatesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSigningCertificatesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pCAVar {in} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.CAVar != nil {
			_ptr_pCAVar := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CAVar != nil {
					if err := o.CAVar.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CAVar, _ptr_pCAVar); err != nil {
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
	return nil
}

func (o *xxx_GetSigningCertificatesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pCAVar {in} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pCAVar := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CAVar == nil {
				o.CAVar = &oaut.Variant{}
			}
			if err := o.CAVar.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pCAVar := func(ptr interface{}) { o.CAVar = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.CAVar, _s_pCAVar, _ptr_pCAVar); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSigningCertificatesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSigningCertificatesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pSigningCertificates {out} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.SigningCertificates != nil {
			_ptr_pSigningCertificates := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SigningCertificates != nil {
					if err := o.SigningCertificates.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SigningCertificates, _ptr_pSigningCertificates); err != nil {
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

func (o *xxx_GetSigningCertificatesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pSigningCertificates {out} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pSigningCertificates := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SigningCertificates == nil {
				o.SigningCertificates = &oaut.Variant{}
			}
			if err := o.SigningCertificates.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pSigningCertificates := func(ptr interface{}) { o.SigningCertificates = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.SigningCertificates, _s_pSigningCertificates, _ptr_pSigningCertificates); err != nil {
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

// GetSigningCertificatesRequest structure represents the GetSigningCertificates operation request
type GetSigningCertificatesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pCAVar: A pointer to a VARIANT data type containing the CA certificate. The vt member of VARIANT SHOULD be set to VT_ARRAY | VT_UI1, and the pArray member SHOULD reference a safearray that contains the ASN.1 DER encoded X.509 certificate data type containing the CA certificate.
	CAVar *oaut.Variant `idl:"name:pCAVar;pointer:ref" json:"ca_var"`
}

func (o *GetSigningCertificatesRequest) xxx_ToOp(ctx context.Context) *xxx_GetSigningCertificatesOperation {
	if o == nil {
		return &xxx_GetSigningCertificatesOperation{}
	}
	return &xxx_GetSigningCertificatesOperation{
		This:  o.This,
		CAVar: o.CAVar,
	}
}

func (o *GetSigningCertificatesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSigningCertificatesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.CAVar = op.CAVar
}
func (o *GetSigningCertificatesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetSigningCertificatesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSigningCertificatesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSigningCertificatesResponse structure represents the GetSigningCertificates operation response
type GetSigningCertificatesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pSigningCertificates: A pointer to VARIANT data type containing the list of certificates. On successful return, the server SHOULD set the vt member of the VARIANT to VT_ARRAY|VT_UI1, and the pArray member SHOULD reference a safearray that contains the ASN.1 DER encoded degenerate PKCS#7 [RFC2315] containing the certificates.
	SigningCertificates *oaut.Variant `idl:"name:pSigningCertificates;pointer:ref" json:"signing_certificates"`
	// Return: The GetSigningCertificates return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSigningCertificatesResponse) xxx_ToOp(ctx context.Context) *xxx_GetSigningCertificatesOperation {
	if o == nil {
		return &xxx_GetSigningCertificatesOperation{}
	}
	return &xxx_GetSigningCertificatesOperation{
		That:                o.That,
		SigningCertificates: o.SigningCertificates,
		Return:              o.Return,
	}
}

func (o *GetSigningCertificatesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSigningCertificatesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SigningCertificates = op.SigningCertificates
	o.Return = op.Return
}
func (o *GetSigningCertificatesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetSigningCertificatesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSigningCertificatesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetHashAlgorithmsOperation structure represents the GetHashAlgorithms operation
type xxx_GetHashAlgorithmsOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	CAID           *oaut.String   `idl:"name:bstrCAId;pointer:ref" json:"ca_id"`
	HashAlgorithms *oaut.Variant  `idl:"name:pHashAlgorithms;pointer:ref" json:"hash_algorithms"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetHashAlgorithmsOperation) OpNum() int { return 10 }

func (o *xxx_GetHashAlgorithmsOperation) OpName() string { return "/IOCSPAdminD/v0/GetHashAlgorithms" }

func (o *xxx_GetHashAlgorithmsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetHashAlgorithmsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrCAId {in} (1:{pointer=ref, alias=BSTR, pointers=[unique]}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.CAID != nil {
			if err := o.CAID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetHashAlgorithmsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrCAId {in} (1:{pointer=ref, alias=BSTR, pointers=[unique]}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.CAID == nil {
			o.CAID = &oaut.String{}
		}
		if err := o.CAID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetHashAlgorithmsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetHashAlgorithmsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pHashAlgorithms {out} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.HashAlgorithms != nil {
			_ptr_pHashAlgorithms := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.HashAlgorithms != nil {
					if err := o.HashAlgorithms.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.HashAlgorithms, _ptr_pHashAlgorithms); err != nil {
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

func (o *xxx_GetHashAlgorithmsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pHashAlgorithms {out} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pHashAlgorithms := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.HashAlgorithms == nil {
				o.HashAlgorithms = &oaut.Variant{}
			}
			if err := o.HashAlgorithms.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pHashAlgorithms := func(ptr interface{}) { o.HashAlgorithms = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.HashAlgorithms, _s_pHashAlgorithms, _ptr_pHashAlgorithms); err != nil {
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

// GetHashAlgorithmsRequest structure represents the GetHashAlgorithms operation request
type GetHashAlgorithmsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrCAId: A BSTR that specifies the RevocationConfigurationId.
	CAID *oaut.String `idl:"name:bstrCAId;pointer:ref" json:"ca_id"`
}

func (o *GetHashAlgorithmsRequest) xxx_ToOp(ctx context.Context) *xxx_GetHashAlgorithmsOperation {
	if o == nil {
		return &xxx_GetHashAlgorithmsOperation{}
	}
	return &xxx_GetHashAlgorithmsOperation{
		This: o.This,
		CAID: o.CAID,
	}
}

func (o *GetHashAlgorithmsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetHashAlgorithmsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.CAID = op.CAID
}
func (o *GetHashAlgorithmsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetHashAlgorithmsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetHashAlgorithmsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetHashAlgorithmsResponse structure represents the GetHashAlgorithms operation response
type GetHashAlgorithmsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pHashAlgorithms:  A pointer to a VARIANT that is of type VT_ARRAY | VT_BSTR. Each element in the array is the name of a hash algorithm that could be used along with the signing certificate associated with a revocation configuration identified by bstrCAId to sign OCSP responses.
	HashAlgorithms *oaut.Variant `idl:"name:pHashAlgorithms;pointer:ref" json:"hash_algorithms"`
	// Return: The GetHashAlgorithms return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetHashAlgorithmsResponse) xxx_ToOp(ctx context.Context) *xxx_GetHashAlgorithmsOperation {
	if o == nil {
		return &xxx_GetHashAlgorithmsOperation{}
	}
	return &xxx_GetHashAlgorithmsOperation{
		That:           o.That,
		HashAlgorithms: o.HashAlgorithms,
		Return:         o.Return,
	}
}

func (o *GetHashAlgorithmsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetHashAlgorithmsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.HashAlgorithms = op.HashAlgorithms
	o.Return = op.Return
}
func (o *GetHashAlgorithmsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetHashAlgorithmsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetHashAlgorithmsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMyRolesOperation structure represents the GetMyRoles operation
type xxx_GetMyRolesOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Roles  int32          `idl:"name:pdwRoles" json:"roles"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMyRolesOperation) OpNum() int { return 11 }

func (o *xxx_GetMyRolesOperation) OpName() string { return "/IOCSPAdminD/v0/GetMyRoles" }

func (o *xxx_GetMyRolesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMyRolesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetMyRolesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetMyRolesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMyRolesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pdwRoles {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Roles); err != nil {
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

func (o *xxx_GetMyRolesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pdwRoles {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Roles); err != nil {
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

// GetMyRolesRequest structure represents the GetMyRoles operation request
type GetMyRolesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMyRolesRequest) xxx_ToOp(ctx context.Context) *xxx_GetMyRolesOperation {
	if o == nil {
		return &xxx_GetMyRolesOperation{}
	}
	return &xxx_GetMyRolesOperation{
		This: o.This,
	}
}

func (o *GetMyRolesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMyRolesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMyRolesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetMyRolesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMyRolesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMyRolesResponse structure represents the GetMyRoles operation response
type GetMyRolesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pdwRoles: Reference to an unsigned integer value that represents the retrieved Online
	// Responder Role for the caller. This can be a bitwise OR of the following values.
	//
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	|                             |                                                                                  |
	//	|            VALUE            |                                     MEANING                                      |
	//	|                             |                                                                                  |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| CA_ACCESS_READ 0x00000100   | The caller can read the configuration information at the responder.              |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| CA_ACCESS_ENROLL 0x00000200 | The caller can request the response status for a particular certificate from the |
	//	|                             | responder.                                                                       |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| CA_ACCESS_ADMIN 0x00000001  | The caller can update the configuration information at the responder.            |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000                  | The caller has no roles.                                                         |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	Roles int32 `idl:"name:pdwRoles" json:"roles"`
	// Return: The GetMyRoles return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMyRolesResponse) xxx_ToOp(ctx context.Context) *xxx_GetMyRolesOperation {
	if o == nil {
		return &xxx_GetMyRolesOperation{}
	}
	return &xxx_GetMyRolesOperation{
		That:   o.That,
		Roles:  o.Roles,
		Return: o.Return,
	}
}

func (o *GetMyRolesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMyRolesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Roles = op.Roles
	o.Return = op.Return
}
func (o *GetMyRolesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetMyRolesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMyRolesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PingOperation structure represents the Ping operation
type xxx_PingOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_PingOperation) OpNum() int { return 12 }

func (o *xxx_PingOperation) OpName() string { return "/IOCSPAdminD/v0/Ping" }

func (o *xxx_PingOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PingOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PingOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_PingOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PingOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PingOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// PingRequest structure represents the Ping operation request
type PingRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *PingRequest) xxx_ToOp(ctx context.Context) *xxx_PingOperation {
	if o == nil {
		return &xxx_PingOperation{}
	}
	return &xxx_PingOperation{
		This: o.This,
	}
}

func (o *PingRequest) xxx_FromOp(ctx context.Context, op *xxx_PingOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *PingRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PingRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PingOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PingResponse structure represents the Ping operation response
type PingResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Ping return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PingResponse) xxx_ToOp(ctx context.Context) *xxx_PingOperation {
	if o == nil {
		return &xxx_PingOperation{}
	}
	return &xxx_PingOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *PingResponse) xxx_FromOp(ctx context.Context, op *xxx_PingOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *PingResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PingResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PingOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
