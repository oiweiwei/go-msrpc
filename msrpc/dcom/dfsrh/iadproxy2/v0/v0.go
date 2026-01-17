package iadproxy2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iadproxy "github.com/oiweiwei/go-msrpc/msrpc/dcom/dfsrh/iadproxy/v0"
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
	_ = dcom.GoPackage
	_ = iadproxy.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/dfsrh"
)

var (
	// IADProxy2 interface identifier c4b0c7d9-abe0-4733-a1e1-9fdedf260c7a
	IADProxy2IID = &dcom.IID{Data1: 0xc4b0c7d9, Data2: 0xabe0, Data3: 0x4733, Data4: []byte{0xa1, 0xe1, 0x9f, 0xde, 0xdf, 0x26, 0x0c, 0x7a}}
	// Syntax UUID
	IADProxy2SyntaxUUID = &uuid.UUID{TimeLow: 0xc4b0c7d9, TimeMid: 0xabe0, TimeHiAndVersion: 0x4733, ClockSeqHiAndReserved: 0xa1, ClockSeqLow: 0xe1, Node: [6]uint8{0x9f, 0xde, 0xdf, 0x26, 0xc, 0x7a}}
	// Syntax ID
	IADProxy2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: IADProxy2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IADProxy2 interface.
type IADProxy2Client interface {

	// IADProxy retrieval method.
	IADProxy() iadproxy.IADProxyClient

	// CreateObject2 operation.
	CreateObject2(context.Context, *CreateObject2Request, ...dcerpc.CallOption) (*CreateObject2Response, error)

	// DeleteObject2 operation.
	DeleteObject2(context.Context, *DeleteObject2Request, ...dcerpc.CallOption) (*DeleteObject2Response, error)

	// ModifyObject2 operation.
	ModifyObject2(context.Context, *ModifyObject2Request, ...dcerpc.CallOption) (*ModifyObject2Response, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) IADProxy2Client
}

type xxx_DefaultIADProxy2Client struct {
	iadproxy.IADProxyClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultIADProxy2Client) IADProxy() iadproxy.IADProxyClient {
	return o.IADProxyClient
}

func (o *xxx_DefaultIADProxy2Client) CreateObject2(ctx context.Context, in *CreateObject2Request, opts ...dcerpc.CallOption) (*CreateObject2Response, error) {
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
	out := &CreateObject2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIADProxy2Client) DeleteObject2(ctx context.Context, in *DeleteObject2Request, opts ...dcerpc.CallOption) (*DeleteObject2Response, error) {
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
	out := &DeleteObject2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIADProxy2Client) ModifyObject2(ctx context.Context, in *ModifyObject2Request, opts ...dcerpc.CallOption) (*ModifyObject2Response, error) {
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
	out := &ModifyObject2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIADProxy2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultIADProxy2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultIADProxy2Client) IPID(ctx context.Context, ipid *dcom.IPID) IADProxy2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultIADProxy2Client{
		IADProxyClient: o.IADProxyClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewIADProxy2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (IADProxy2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(IADProxy2SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iadproxy.NewIADProxyClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultIADProxy2Client{
		IADProxyClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_CreateObject2Operation structure represents the CreateObject2 operation
type xxx_CreateObject2Operation struct {
	This                           *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That                           *dcom.ORPCThat  `idl:"name:That" json:"that"`
	DomainControllerName           *oaut.String    `idl:"name:domainControllerName" json:"domain_controller_name"`
	DistinguishedName              *oaut.String    `idl:"name:distinguishedName" json:"distinguished_name"`
	Attributes                     *oaut.SafeArray `idl:"name:attributes" json:"attributes"`
	VerifyNameDomainControllerName *oaut.String    `idl:"name:verifyNameDomainControllerName" json:"verify_name_domain_controller_name"`
	NetworkNameResourceName        *oaut.String    `idl:"name:networkNameResourceName" json:"network_name_resource_name"`
	Return                         int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateObject2Operation) OpNum() int { return 6 }

func (o *xxx_CreateObject2Operation) OpName() string { return "/IADProxy2/v0/CreateObject2" }

func (o *xxx_CreateObject2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateObject2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// networkNameResourceName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.NetworkNameResourceName != nil {
			_ptr_networkNameResourceName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NetworkNameResourceName != nil {
					if err := o.NetworkNameResourceName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NetworkNameResourceName, _ptr_networkNameResourceName); err != nil {
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

func (o *xxx_CreateObject2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// networkNameResourceName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_networkNameResourceName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NetworkNameResourceName == nil {
				o.NetworkNameResourceName = &oaut.String{}
			}
			if err := o.NetworkNameResourceName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_networkNameResourceName := func(ptr interface{}) { o.NetworkNameResourceName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.NetworkNameResourceName, _s_networkNameResourceName, _ptr_networkNameResourceName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateObject2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateObject2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateObject2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CreateObject2Request structure represents the CreateObject2 operation request
type CreateObject2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                           *dcom.ORPCThis  `idl:"name:This" json:"this"`
	DomainControllerName           *oaut.String    `idl:"name:domainControllerName" json:"domain_controller_name"`
	DistinguishedName              *oaut.String    `idl:"name:distinguishedName" json:"distinguished_name"`
	Attributes                     *oaut.SafeArray `idl:"name:attributes" json:"attributes"`
	VerifyNameDomainControllerName *oaut.String    `idl:"name:verifyNameDomainControllerName" json:"verify_name_domain_controller_name"`
	NetworkNameResourceName        *oaut.String    `idl:"name:networkNameResourceName" json:"network_name_resource_name"`
}

func (o *CreateObject2Request) xxx_ToOp(ctx context.Context, op *xxx_CreateObject2Operation) *xxx_CreateObject2Operation {
	if op == nil {
		op = &xxx_CreateObject2Operation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.DomainControllerName = o.DomainControllerName
	op.DistinguishedName = o.DistinguishedName
	op.Attributes = o.Attributes
	op.VerifyNameDomainControllerName = o.VerifyNameDomainControllerName
	op.NetworkNameResourceName = o.NetworkNameResourceName
	return op
}

func (o *CreateObject2Request) xxx_FromOp(ctx context.Context, op *xxx_CreateObject2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DomainControllerName = op.DomainControllerName
	o.DistinguishedName = op.DistinguishedName
	o.Attributes = op.Attributes
	o.VerifyNameDomainControllerName = op.VerifyNameDomainControllerName
	o.NetworkNameResourceName = op.NetworkNameResourceName
}
func (o *CreateObject2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateObject2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateObject2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateObject2Response structure represents the CreateObject2 operation response
type CreateObject2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CreateObject2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateObject2Response) xxx_ToOp(ctx context.Context, op *xxx_CreateObject2Operation) *xxx_CreateObject2Operation {
	if op == nil {
		op = &xxx_CreateObject2Operation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CreateObject2Response) xxx_FromOp(ctx context.Context, op *xxx_CreateObject2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CreateObject2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateObject2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateObject2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteObject2Operation structure represents the DeleteObject2 operation
type xxx_DeleteObject2Operation struct {
	This                    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                    *dcom.ORPCThat `idl:"name:That" json:"that"`
	DomainControllerName    *oaut.String   `idl:"name:domainControllerName" json:"domain_controller_name"`
	DistinguishedName       *oaut.String   `idl:"name:distinguishedName" json:"distinguished_name"`
	NetworkNameResourceName *oaut.String   `idl:"name:networkNameResourceName" json:"network_name_resource_name"`
	Return                  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteObject2Operation) OpNum() int { return 7 }

func (o *xxx_DeleteObject2Operation) OpName() string { return "/IADProxy2/v0/DeleteObject2" }

func (o *xxx_DeleteObject2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteObject2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// networkNameResourceName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.NetworkNameResourceName != nil {
			_ptr_networkNameResourceName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NetworkNameResourceName != nil {
					if err := o.NetworkNameResourceName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NetworkNameResourceName, _ptr_networkNameResourceName); err != nil {
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

func (o *xxx_DeleteObject2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// networkNameResourceName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_networkNameResourceName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NetworkNameResourceName == nil {
				o.NetworkNameResourceName = &oaut.String{}
			}
			if err := o.NetworkNameResourceName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_networkNameResourceName := func(ptr interface{}) { o.NetworkNameResourceName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.NetworkNameResourceName, _s_networkNameResourceName, _ptr_networkNameResourceName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteObject2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteObject2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteObject2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteObject2Request structure represents the DeleteObject2 operation request
type DeleteObject2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                    *dcom.ORPCThis `idl:"name:This" json:"this"`
	DomainControllerName    *oaut.String   `idl:"name:domainControllerName" json:"domain_controller_name"`
	DistinguishedName       *oaut.String   `idl:"name:distinguishedName" json:"distinguished_name"`
	NetworkNameResourceName *oaut.String   `idl:"name:networkNameResourceName" json:"network_name_resource_name"`
}

func (o *DeleteObject2Request) xxx_ToOp(ctx context.Context, op *xxx_DeleteObject2Operation) *xxx_DeleteObject2Operation {
	if op == nil {
		op = &xxx_DeleteObject2Operation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.DomainControllerName = o.DomainControllerName
	op.DistinguishedName = o.DistinguishedName
	op.NetworkNameResourceName = o.NetworkNameResourceName
	return op
}

func (o *DeleteObject2Request) xxx_FromOp(ctx context.Context, op *xxx_DeleteObject2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DomainControllerName = op.DomainControllerName
	o.DistinguishedName = op.DistinguishedName
	o.NetworkNameResourceName = op.NetworkNameResourceName
}
func (o *DeleteObject2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteObject2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteObject2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteObject2Response structure represents the DeleteObject2 operation response
type DeleteObject2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteObject2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteObject2Response) xxx_ToOp(ctx context.Context, op *xxx_DeleteObject2Operation) *xxx_DeleteObject2Operation {
	if op == nil {
		op = &xxx_DeleteObject2Operation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DeleteObject2Response) xxx_FromOp(ctx context.Context, op *xxx_DeleteObject2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteObject2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteObject2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteObject2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ModifyObject2Operation structure represents the ModifyObject2 operation
type xxx_ModifyObject2Operation struct {
	This                    *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That                    *dcom.ORPCThat  `idl:"name:That" json:"that"`
	DomainControllerName    *oaut.String    `idl:"name:domainControllerName" json:"domain_controller_name"`
	DistinguishedName       *oaut.String    `idl:"name:distinguishedName" json:"distinguished_name"`
	Attributes              *oaut.SafeArray `idl:"name:attributes" json:"attributes"`
	NetworkNameResourceName *oaut.String    `idl:"name:networkNameResourceName" json:"network_name_resource_name"`
	Return                  int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_ModifyObject2Operation) OpNum() int { return 8 }

func (o *xxx_ModifyObject2Operation) OpName() string { return "/IADProxy2/v0/ModifyObject2" }

func (o *xxx_ModifyObject2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyObject2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// networkNameResourceName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.NetworkNameResourceName != nil {
			_ptr_networkNameResourceName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NetworkNameResourceName != nil {
					if err := o.NetworkNameResourceName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NetworkNameResourceName, _ptr_networkNameResourceName); err != nil {
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

func (o *xxx_ModifyObject2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// networkNameResourceName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_networkNameResourceName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NetworkNameResourceName == nil {
				o.NetworkNameResourceName = &oaut.String{}
			}
			if err := o.NetworkNameResourceName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_networkNameResourceName := func(ptr interface{}) { o.NetworkNameResourceName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.NetworkNameResourceName, _s_networkNameResourceName, _ptr_networkNameResourceName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyObject2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyObject2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ModifyObject2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ModifyObject2Request structure represents the ModifyObject2 operation request
type ModifyObject2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                    *dcom.ORPCThis  `idl:"name:This" json:"this"`
	DomainControllerName    *oaut.String    `idl:"name:domainControllerName" json:"domain_controller_name"`
	DistinguishedName       *oaut.String    `idl:"name:distinguishedName" json:"distinguished_name"`
	Attributes              *oaut.SafeArray `idl:"name:attributes" json:"attributes"`
	NetworkNameResourceName *oaut.String    `idl:"name:networkNameResourceName" json:"network_name_resource_name"`
}

func (o *ModifyObject2Request) xxx_ToOp(ctx context.Context, op *xxx_ModifyObject2Operation) *xxx_ModifyObject2Operation {
	if op == nil {
		op = &xxx_ModifyObject2Operation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.DomainControllerName = o.DomainControllerName
	op.DistinguishedName = o.DistinguishedName
	op.Attributes = o.Attributes
	op.NetworkNameResourceName = o.NetworkNameResourceName
	return op
}

func (o *ModifyObject2Request) xxx_FromOp(ctx context.Context, op *xxx_ModifyObject2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DomainControllerName = op.DomainControllerName
	o.DistinguishedName = op.DistinguishedName
	o.Attributes = op.Attributes
	o.NetworkNameResourceName = op.NetworkNameResourceName
}
func (o *ModifyObject2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ModifyObject2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyObject2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ModifyObject2Response structure represents the ModifyObject2 operation response
type ModifyObject2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ModifyObject2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ModifyObject2Response) xxx_ToOp(ctx context.Context, op *xxx_ModifyObject2Operation) *xxx_ModifyObject2Operation {
	if op == nil {
		op = &xxx_ModifyObject2Operation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ModifyObject2Response) xxx_FromOp(ctx context.Context, op *xxx_ModifyObject2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ModifyObject2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ModifyObject2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyObject2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
