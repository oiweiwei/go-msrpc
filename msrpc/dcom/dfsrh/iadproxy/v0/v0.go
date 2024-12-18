package iadproxy

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
	GoPackage = "dcom/dfsrh"
)

var (
	// IADProxy interface identifier 4bb8ab1d-9ef9-4100-8eb6-dd4b4e418b72
	IADProxyIID = &dcom.IID{Data1: 0x4bb8ab1d, Data2: 0x9ef9, Data3: 0x4100, Data4: []byte{0x8e, 0xb6, 0xdd, 0x4b, 0x4e, 0x41, 0x8b, 0x72}}
	// Syntax UUID
	IADProxySyntaxUUID = &uuid.UUID{TimeLow: 0x4bb8ab1d, TimeMid: 0x9ef9, TimeHiAndVersion: 0x4100, ClockSeqHiAndReserved: 0x8e, ClockSeqLow: 0xb6, Node: [6]uint8{0xdd, 0x4b, 0x4e, 0x41, 0x8b, 0x72}}
	// Syntax ID
	IADProxySyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: IADProxySyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IADProxy interface.
type IADProxyClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The CreateObject method MUST execute an LDAP command under machine security credentials,
	// or for a cluster, under the specified network name credentials in order to create
	// an Active Directory object that has a specific distinguished name and attributes.<52>
	//
	// Return Values: The method MUST return:
	//
	// * 0 on success.
	//
	// * For LDAP protocol failures:
	//
	// * If the LDAP error is LDAP_OPERATIONS_ERROR, dfsrHelperLdapErrorBase + the server-side
	// error code.
	//
	// * For all other LDAP errors, dfsrHelperLdapErrorBase + the LDAP return code. For
	// more information, see [LDAP-ERR] ( https://go.microsoft.com/fwlink/?LinkId=89933
	// ).
	//
	// * For all other failures, an implementation-specific nonzero HRESULT ( ../ms-dtyp/a9046ed2-bfb2-4d56-a719-2824afce59ac
	// ) error code, as specified in [MS-ERREF] ( ../ms-erref/1bc92ddf-b79e-413c-bbaa-99a5281a6c90
	// ) section 2.1 ( ../ms-erref/0642cb2f-2075-4469-918c-4441e69c548a ) , between 0x80000000
	// and 0xFFFFFFFF. For protocol purposes, all nonzero values MUST be treated as equivalent
	// failures.
	CreateObject(context.Context, *CreateObjectRequest, ...dcerpc.CallOption) (*CreateObjectResponse, error)

	// The DeleteObject method MUST execute an LDAP command under machine security credentials
	// to delete an Active Directory object with a specified distinguished name.<46>
	//
	// The DeleteObject method executes an LDAP command to delete an Active Directory object
	// that has a specified distinguished name and attributes. The command MUST be executed
	// under the machine security credentials, or for a cluster, under the specified network
	// name credentials.<54>
	//
	// Return Values: The method MUST return:
	//
	// * A value of 0 when:
	//
	// * The method call is successful.
	//
	// * The LDAP error is LDAP_NO_SUCH_OBJECT.
	//
	// * For other LDAP protocol failures:
	//
	// * A value of dfsrHelperLdapErrorBase + the server-side error code  if the LDAP error
	// is LDAP_OPERATIONS_ERROR.
	//
	// * A value of dfsrHelperLdapErrorBase + the LDAP return code for all other LDAP errors.
	// For more information, see [LDAP-ERR] ( https://go.microsoft.com/fwlink/?LinkId=89933
	// ).
	//
	// * For all other failures, an implementation-specific nonzero HRESULT ( ../ms-dtyp/a9046ed2-bfb2-4d56-a719-2824afce59ac
	// ) error code, as specified in [MS-ERREF] ( ../ms-erref/1bc92ddf-b79e-413c-bbaa-99a5281a6c90
	// ) section 2.1 ( ../ms-erref/0642cb2f-2075-4469-918c-4441e69c548a ) , between 0x80000000
	// and 0xFFFFFFFF. For protocol purposes, all nonzero values MUST be treated as equivalent
	// failures.
	DeleteObject(context.Context, *DeleteObjectRequest, ...dcerpc.CallOption) (*DeleteObjectResponse, error)

	// The ModifyObject method executes an LDAP command to add, delete, or modify attributes
	// of a specified Active Directory object. The command MUST be executed under machine
	// security credentials, or for a cluster, under the specified network name credentials
	// in order to modify an Active Directory object that has a specific distinguished name
	// and attributes.<56>
	//
	// The ModifyObject method MUST execute an LDAP command under machine security credentials
	// to add, delete, or modify attributes of an Active Directory object that has a specified
	// distinguished name.<48>
	//
	// Return Values: The method MUST return:
	//
	// * Zero on success.
	//
	// * For LDAP protocol failures:
	//
	// * If the LDAP error is LDAP_OPERATIONS_ERROR, dfsrHelperLdapErrorBase + the server-side
	// error code.
	//
	// * For all other LDAP errors, dfsrHelperLdapErrorBase + the LDAP return code. For
	// more information, see [LDAP-ERR] ( https://go.microsoft.com/fwlink/?LinkId=89933
	// ).
	//
	// * For all other failures, an implementation-specific nonzero HRESULT ( ../ms-dtyp/a9046ed2-bfb2-4d56-a719-2824afce59ac
	// ) error code, as specified in [MS-ERREF] ( ../ms-erref/1bc92ddf-b79e-413c-bbaa-99a5281a6c90
	// ) section 2.1 ( ../ms-erref/0642cb2f-2075-4469-918c-4441e69c548a ) , between 0x80000000
	// and 0xFFFFFFFF. For protocol purposes, all nonzero values MUST be treated as equivalent
	// failures.
	ModifyObject(context.Context, *ModifyObjectRequest, ...dcerpc.CallOption) (*ModifyObjectResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) IADProxyClient
}

type xxx_DefaultIADProxyClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultIADProxyClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultIADProxyClient) CreateObject(ctx context.Context, in *CreateObjectRequest, opts ...dcerpc.CallOption) (*CreateObjectResponse, error) {
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
	out := &CreateObjectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIADProxyClient) DeleteObject(ctx context.Context, in *DeleteObjectRequest, opts ...dcerpc.CallOption) (*DeleteObjectResponse, error) {
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
	out := &DeleteObjectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIADProxyClient) ModifyObject(ctx context.Context, in *ModifyObjectRequest, opts ...dcerpc.CallOption) (*ModifyObjectResponse, error) {
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
	out := &ModifyObjectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIADProxyClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultIADProxyClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultIADProxyClient) IPID(ctx context.Context, ipid *dcom.IPID) IADProxyClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultIADProxyClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewIADProxyClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (IADProxyClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(IADProxySyntaxV0_0))...)
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
	return &xxx_DefaultIADProxyClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_CreateObjectOperation structure represents the CreateObject operation
type xxx_CreateObjectOperation struct {
	This                           *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That                           *dcom.ORPCThat  `idl:"name:That" json:"that"`
	DomainControllerName           *oaut.String    `idl:"name:domainControllerName" json:"domain_controller_name"`
	DistinguishedName              *oaut.String    `idl:"name:distinguishedName" json:"distinguished_name"`
	Attributes                     *oaut.SafeArray `idl:"name:attributes" json:"attributes"`
	VerifyNameDomainControllerName *oaut.String    `idl:"name:verifyNameDomainControllerName" json:"verify_name_domain_controller_name"`
	Return                         int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateObjectOperation) OpNum() int { return 3 }

func (o *xxx_CreateObjectOperation) OpName() string { return "/IADProxy/v0/CreateObject" }

func (o *xxx_CreateObjectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateObjectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// domainControllerName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.DomainControllerName != nil {
			_ptr_domainControllerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DomainControllerName != nil {
					if err := o.DomainControllerName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DomainControllerName, _ptr_domainControllerName); err != nil {
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
	// distinguishedName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.DistinguishedName != nil {
			_ptr_distinguishedName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DistinguishedName != nil {
					if err := o.DistinguishedName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DistinguishedName, _ptr_distinguishedName); err != nil {
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
	// attributes {in} (1:{pointer=ref}*(3)*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Attributes != nil {
			_ptr_attributes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Attributes != nil {
					_ptr_attributes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Attributes != nil {
							if err := o.Attributes.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Attributes, _ptr_attributes); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Attributes, _ptr_attributes); err != nil {
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
	// verifyNameDomainControllerName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.VerifyNameDomainControllerName != nil {
			_ptr_verifyNameDomainControllerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.VerifyNameDomainControllerName != nil {
					if err := o.VerifyNameDomainControllerName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.VerifyNameDomainControllerName, _ptr_verifyNameDomainControllerName); err != nil {
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

func (o *xxx_CreateObjectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// domainControllerName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_domainControllerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DomainControllerName == nil {
				o.DomainControllerName = &oaut.String{}
			}
			if err := o.DomainControllerName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_domainControllerName := func(ptr interface{}) { o.DomainControllerName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.DomainControllerName, _s_domainControllerName, _ptr_domainControllerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// distinguishedName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_distinguishedName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DistinguishedName == nil {
				o.DistinguishedName = &oaut.String{}
			}
			if err := o.DistinguishedName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_distinguishedName := func(ptr interface{}) { o.DistinguishedName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.DistinguishedName, _s_distinguishedName, _ptr_distinguishedName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// attributes {in} (1:{pointer=ref}*(3)*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_attributes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_attributes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Attributes == nil {
					o.Attributes = &oaut.SafeArray{}
				}
				if err := o.Attributes.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_attributes := func(ptr interface{}) { o.Attributes = *ptr.(**oaut.SafeArray) }
			if err := w.ReadPointer(&o.Attributes, _s_attributes, _ptr_attributes); err != nil {
				return err
			}
			return nil
		})
		_s_attributes := func(ptr interface{}) { o.Attributes = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Attributes, _s_attributes, _ptr_attributes); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// verifyNameDomainControllerName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_verifyNameDomainControllerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.VerifyNameDomainControllerName == nil {
				o.VerifyNameDomainControllerName = &oaut.String{}
			}
			if err := o.VerifyNameDomainControllerName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_verifyNameDomainControllerName := func(ptr interface{}) { o.VerifyNameDomainControllerName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.VerifyNameDomainControllerName, _s_verifyNameDomainControllerName, _ptr_verifyNameDomainControllerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateObjectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateObjectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateObjectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CreateObjectRequest structure represents the CreateObject operation request
type CreateObjectRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// domainControllerName: MUST be the FQDN of the domain controller to which the method
	// sends the LDAP request. The format of the distinguished name MUST be as specified
	// in [RFC2251] section 4.1.3.
	DomainControllerName *oaut.String `idl:"name:domainControllerName" json:"domain_controller_name"`
	// distinguishedName: MUST be the distinguished name of the Active Directory object
	// that is being created. The distinguished name of any object in Active Directory MAY
	// be used. The format of the distinguished name is specified in [RFC2251] section 4.1.3.
	// <53>
	DistinguishedName *oaut.String `idl:"name:distinguishedName" json:"distinguished_name"`
	// attributes: MUST be the safe array of attributes that are to be created for the new
	// object.
	Attributes *oaut.SafeArray `idl:"name:attributes" json:"attributes"`
	// verifyNameDomainControllerName: If the attributes of the object refer to an object
	// in another domain, the client MUST specify the domain controller in that domain that
	// will be used to verify the reference. The LDAP_SERVER_VERIFY_NAME_OID control MUST
	// be added to the LDAP command. The LDAP_SERVER_VERIFY_NAME_OID control is specified
	// in [MS-ADTS] section 3.1.1.3.4.1.16. If this parameter does not specify a domain
	// controller that can be contacted to validate these references, the method MUST fail
	// and return an LDAP protocol failure.
	VerifyNameDomainControllerName *oaut.String `idl:"name:verifyNameDomainControllerName" json:"verify_name_domain_controller_name"`
}

func (o *CreateObjectRequest) xxx_ToOp(ctx context.Context) *xxx_CreateObjectOperation {
	if o == nil {
		return &xxx_CreateObjectOperation{}
	}
	return &xxx_CreateObjectOperation{
		This:                           o.This,
		DomainControllerName:           o.DomainControllerName,
		DistinguishedName:              o.DistinguishedName,
		Attributes:                     o.Attributes,
		VerifyNameDomainControllerName: o.VerifyNameDomainControllerName,
	}
}

func (o *CreateObjectRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateObjectOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DomainControllerName = op.DomainControllerName
	o.DistinguishedName = op.DistinguishedName
	o.Attributes = op.Attributes
	o.VerifyNameDomainControllerName = op.VerifyNameDomainControllerName
}
func (o *CreateObjectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateObjectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateObjectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateObjectResponse structure represents the CreateObject operation response
type CreateObjectResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CreateObject return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateObjectResponse) xxx_ToOp(ctx context.Context) *xxx_CreateObjectOperation {
	if o == nil {
		return &xxx_CreateObjectOperation{}
	}
	return &xxx_CreateObjectOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *CreateObjectResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateObjectOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CreateObjectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateObjectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateObjectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteObjectOperation structure represents the DeleteObject operation
type xxx_DeleteObjectOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	DomainControllerName *oaut.String   `idl:"name:domainControllerName" json:"domain_controller_name"`
	DistinguishedName    *oaut.String   `idl:"name:distinguishedName" json:"distinguished_name"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteObjectOperation) OpNum() int { return 4 }

func (o *xxx_DeleteObjectOperation) OpName() string { return "/IADProxy/v0/DeleteObject" }

func (o *xxx_DeleteObjectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteObjectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// domainControllerName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.DomainControllerName != nil {
			_ptr_domainControllerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DomainControllerName != nil {
					if err := o.DomainControllerName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DomainControllerName, _ptr_domainControllerName); err != nil {
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
	// distinguishedName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.DistinguishedName != nil {
			_ptr_distinguishedName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DistinguishedName != nil {
					if err := o.DistinguishedName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DistinguishedName, _ptr_distinguishedName); err != nil {
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

func (o *xxx_DeleteObjectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// domainControllerName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_domainControllerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DomainControllerName == nil {
				o.DomainControllerName = &oaut.String{}
			}
			if err := o.DomainControllerName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_domainControllerName := func(ptr interface{}) { o.DomainControllerName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.DomainControllerName, _s_domainControllerName, _ptr_domainControllerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// distinguishedName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_distinguishedName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DistinguishedName == nil {
				o.DistinguishedName = &oaut.String{}
			}
			if err := o.DistinguishedName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_distinguishedName := func(ptr interface{}) { o.DistinguishedName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.DistinguishedName, _s_distinguishedName, _ptr_distinguishedName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteObjectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteObjectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteObjectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteObjectRequest structure represents the DeleteObject operation request
type DeleteObjectRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// domainControllerName: MUST be the FQDN of the domain controller to which the method
	// sends the LDAP command.
	//
	// domainControllerName: MUST be the FQDN of the domain controller to which the method
	// sends the LDAP request.
	DomainControllerName *oaut.String `idl:"name:domainControllerName" json:"domain_controller_name"`
	// distinguishedName: MUST be the distinguished name of the Active Directory object
	// that is being deleted. The distinguished name of any object in Active Directory can
	// be used. The format of the distinguished name is specified in [RFC2251] section 4.1.3.
	//
	// distinguishedName: MUST be the distinguished name of the Active Directory object
	// that is being deleted. The distinguished name of any object in Active Directory MAY
	// be used. The format of the distinguished name is specified in [RFC2251] section 4.1.3.
	// <55>
	DistinguishedName *oaut.String `idl:"name:distinguishedName" json:"distinguished_name"`
}

func (o *DeleteObjectRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteObjectOperation {
	if o == nil {
		return &xxx_DeleteObjectOperation{}
	}
	return &xxx_DeleteObjectOperation{
		This:                 o.This,
		DomainControllerName: o.DomainControllerName,
		DistinguishedName:    o.DistinguishedName,
	}
}

func (o *DeleteObjectRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteObjectOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DomainControllerName = op.DomainControllerName
	o.DistinguishedName = op.DistinguishedName
}
func (o *DeleteObjectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteObjectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteObjectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteObjectResponse structure represents the DeleteObject operation response
type DeleteObjectResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteObject return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteObjectResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteObjectOperation {
	if o == nil {
		return &xxx_DeleteObjectOperation{}
	}
	return &xxx_DeleteObjectOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *DeleteObjectResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteObjectOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteObjectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteObjectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteObjectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ModifyObjectOperation structure represents the ModifyObject operation
type xxx_ModifyObjectOperation struct {
	This                 *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat  `idl:"name:That" json:"that"`
	DomainControllerName *oaut.String    `idl:"name:domainControllerName" json:"domain_controller_name"`
	DistinguishedName    *oaut.String    `idl:"name:distinguishedName" json:"distinguished_name"`
	Attributes           *oaut.SafeArray `idl:"name:attributes" json:"attributes"`
	Return               int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_ModifyObjectOperation) OpNum() int { return 5 }

func (o *xxx_ModifyObjectOperation) OpName() string { return "/IADProxy/v0/ModifyObject" }

func (o *xxx_ModifyObjectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyObjectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// domainControllerName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.DomainControllerName != nil {
			_ptr_domainControllerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DomainControllerName != nil {
					if err := o.DomainControllerName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DomainControllerName, _ptr_domainControllerName); err != nil {
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
	// distinguishedName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.DistinguishedName != nil {
			_ptr_distinguishedName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DistinguishedName != nil {
					if err := o.DistinguishedName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DistinguishedName, _ptr_distinguishedName); err != nil {
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
	// attributes {in} (1:{pointer=ref}*(3)*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Attributes != nil {
			_ptr_attributes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Attributes != nil {
					_ptr_attributes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Attributes != nil {
							if err := o.Attributes.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Attributes, _ptr_attributes); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Attributes, _ptr_attributes); err != nil {
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

func (o *xxx_ModifyObjectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// domainControllerName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_domainControllerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DomainControllerName == nil {
				o.DomainControllerName = &oaut.String{}
			}
			if err := o.DomainControllerName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_domainControllerName := func(ptr interface{}) { o.DomainControllerName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.DomainControllerName, _s_domainControllerName, _ptr_domainControllerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// distinguishedName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_distinguishedName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DistinguishedName == nil {
				o.DistinguishedName = &oaut.String{}
			}
			if err := o.DistinguishedName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_distinguishedName := func(ptr interface{}) { o.DistinguishedName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.DistinguishedName, _s_distinguishedName, _ptr_distinguishedName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// attributes {in} (1:{pointer=ref}*(3)*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_attributes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_attributes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Attributes == nil {
					o.Attributes = &oaut.SafeArray{}
				}
				if err := o.Attributes.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_attributes := func(ptr interface{}) { o.Attributes = *ptr.(**oaut.SafeArray) }
			if err := w.ReadPointer(&o.Attributes, _s_attributes, _ptr_attributes); err != nil {
				return err
			}
			return nil
		})
		_s_attributes := func(ptr interface{}) { o.Attributes = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Attributes, _s_attributes, _ptr_attributes); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyObjectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyObjectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ModifyObjectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ModifyObjectRequest structure represents the ModifyObject operation request
type ModifyObjectRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// domainControllerName: MUST be the FQDN of the domain controller to which the method
	// sends the LDAP request.
	//
	// domainControllerName: MUST be the FQDN of the domain controller to which the method
	// sends the LDAP request. The format of the distinguished name is specified in [RFC2251]
	// section 4.1.3.
	DomainControllerName *oaut.String `idl:"name:domainControllerName" json:"domain_controller_name"`
	// distinguishedName: MUST be the distinguished name of the Active Directory object
	// that is being modified. The distinguished name of any object in Active Directory
	// MAY be used. The format of the distinguished name is specified in [RFC2251] section
	// 4.1.3.  <57>
	//
	// distinguishedName: MUST be the distinguished name of the Active Directory object
	// being modified. The distinguished name of any object in Active Directory MAY be used.
	// <49>
	DistinguishedName *oaut.String `idl:"name:distinguishedName" json:"distinguished_name"`
	// attributes: MUST be the safe array of attributes that are to be added, modified,
	// or deleted.
	Attributes *oaut.SafeArray `idl:"name:attributes" json:"attributes"`
}

func (o *ModifyObjectRequest) xxx_ToOp(ctx context.Context) *xxx_ModifyObjectOperation {
	if o == nil {
		return &xxx_ModifyObjectOperation{}
	}
	return &xxx_ModifyObjectOperation{
		This:                 o.This,
		DomainControllerName: o.DomainControllerName,
		DistinguishedName:    o.DistinguishedName,
		Attributes:           o.Attributes,
	}
}

func (o *ModifyObjectRequest) xxx_FromOp(ctx context.Context, op *xxx_ModifyObjectOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DomainControllerName = op.DomainControllerName
	o.DistinguishedName = op.DistinguishedName
	o.Attributes = op.Attributes
}
func (o *ModifyObjectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ModifyObjectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyObjectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ModifyObjectResponse structure represents the ModifyObject operation response
type ModifyObjectResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ModifyObject return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ModifyObjectResponse) xxx_ToOp(ctx context.Context) *xxx_ModifyObjectOperation {
	if o == nil {
		return &xxx_ModifyObjectOperation{}
	}
	return &xxx_ModifyObjectOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *ModifyObjectResponse) xxx_FromOp(ctx context.Context, op *xxx_ModifyObjectOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ModifyObjectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ModifyObjectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyObjectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
