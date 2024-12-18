package itypecomp

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
	// ITypeComp interface identifier 00020403-0000-0000-c000-000000000046
	TypeCompIID = &dcom.IID{Data1: 0x00020403, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
	// Syntax UUID
	TypeCompSyntaxUUID = &uuid.UUID{TimeLow: 0x20403, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xc0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x46}}
	// Syntax ID
	TypeCompSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: TypeCompSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// ITypeComp interface.
type TypeCompClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The Bind method retrieves automation type description server instances and related
	// structures that provide access to a method, property or data member whose name corresponds
	// to a specified string.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, which
	// is defined in [MS-ERREF] section 2.1. The severity bit in the structure identifies
	// the following conditions:
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
	// 0x8002802C
	//
	// TYPE_E_AMBIGUOUSNAME
	//
	// The values of szName and wFlags match more than one element in the binding context.
	// See [MS-ERREF].
	//
	// 0x80028CA0
	//
	// TYPE_E_TYPEMISMATCH
	//
	// The value of szName matched a single element in the binding context, but the *INVOKEKIND*
	// value that is associated with that element did not match the value of wFlags. See
	// [MS-ERREF].
	Bind(context.Context, *BindRequest, ...dcerpc.CallOption) (*BindResponse, error)

	// The BindType method retrieves a reference to an automation type description whose
	// name corresponds to a specified string.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	BindType(context.Context, *BindTypeRequest, ...dcerpc.CallOption) (*BindTypeResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) TypeCompClient
}

type xxx_DefaultTypeCompClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultTypeCompClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultTypeCompClient) Bind(ctx context.Context, in *BindRequest, opts ...dcerpc.CallOption) (*BindResponse, error) {
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
	out := &BindResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeCompClient) BindType(ctx context.Context, in *BindTypeRequest, opts ...dcerpc.CallOption) (*BindTypeResponse, error) {
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
	out := &BindTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeCompClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultTypeCompClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultTypeCompClient) IPID(ctx context.Context, ipid *dcom.IPID) TypeCompClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultTypeCompClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewTypeCompClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (TypeCompClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(TypeCompSyntaxV0_0))...)
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
	return &xxx_DefaultTypeCompClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_BindOperation structure represents the Bind operation
type xxx_BindOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Name      string         `idl:"name:szName" json:"name"`
	HashValue uint32         `idl:"name:lHashVal" json:"hash_value"`
	Flags     uint16         `idl:"name:wFlags" json:"flags"`
	TypeInfo  *oaut.TypeInfo `idl:"name:ppTInfo" json:"type_info"`
	DescKind  oaut.DescKind  `idl:"name:pDescKind" json:"desc_kind"`
	FuncDesc  *oaut.FuncDesc `idl:"name:ppFuncDesc" json:"func_desc"`
	VarDesc   *oaut.VarDesc  `idl:"name:ppVarDesc" json:"var_desc"`
	TypeComp  *oaut.TypeComp `idl:"name:ppTypeComp" json:"type_comp"`
	_         uint32         `idl:"name:pReserved"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_BindOperation) OpNum() int { return 3 }

func (o *xxx_BindOperation) OpName() string { return "/ITypeComp/v0/Bind" }

func (o *xxx_BindOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BindOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// szName {in} (1:{string, alias=LPOLESTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
			return err
		}
	}
	// lHashVal {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.HashValue); err != nil {
			return err
		}
	}
	// wFlags {in} (1:{alias=WORD}(uint16))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BindOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// szName {in} (1:{string, alias=LPOLESTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
	}
	// lHashVal {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.HashValue); err != nil {
			return err
		}
	}
	// wFlags {in} (1:{alias=WORD}(uint16))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BindOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BindOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pDescKind {out} (1:{pointer=ref}*(1))(2:{v1_enum, alias=DESCKIND}(enum))
	{
		if err := w.WriteData(uint32(o.DescKind)); err != nil {
			return err
		}
	}
	// ppFuncDesc {out} (1:{pointer=ref}*(2))(2:{alias=LPFUNCDESC}*(1))(3:{alias=FUNCDESC}(struct))
	{
		if o.FuncDesc != nil {
			_ptr_ppFuncDesc := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FuncDesc != nil {
					if err := o.FuncDesc.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.FuncDesc{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FuncDesc, _ptr_ppFuncDesc); err != nil {
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
	// ppVarDesc {out} (1:{pointer=ref}*(2))(2:{alias=LPVARDESC}*(1))(3:{alias=VARDESC}(struct))
	{
		if o.VarDesc != nil {
			_ptr_ppVarDesc := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.VarDesc != nil {
					if err := o.VarDesc.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.VarDesc{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.VarDesc, _ptr_ppVarDesc); err != nil {
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
	// ppTypeComp {out} (1:{pointer=ref}*(2)*(1))(2:{alias=ITypeComp}(interface))
	{
		if o.TypeComp != nil {
			_ptr_ppTypeComp := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.TypeComp != nil {
					if err := o.TypeComp.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.TypeComp{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.TypeComp, _ptr_ppTypeComp); err != nil {
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

func (o *xxx_BindOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pDescKind {out} (1:{pointer=ref}*(1))(2:{v1_enum, alias=DESCKIND}(enum))
	{
		if err := w.ReadData((*uint32)(&o.DescKind)); err != nil {
			return err
		}
	}
	// ppFuncDesc {out} (1:{pointer=ref}*(2))(2:{alias=LPFUNCDESC,pointer=ref}*(1))(3:{alias=FUNCDESC}(struct))
	{
		_ptr_ppFuncDesc := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FuncDesc == nil {
				o.FuncDesc = &oaut.FuncDesc{}
			}
			if err := o.FuncDesc.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppFuncDesc := func(ptr interface{}) { o.FuncDesc = *ptr.(**oaut.FuncDesc) }
		if err := w.ReadPointer(&o.FuncDesc, _s_ppFuncDesc, _ptr_ppFuncDesc); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppVarDesc {out} (1:{pointer=ref}*(2))(2:{alias=LPVARDESC,pointer=ref}*(1))(3:{alias=VARDESC}(struct))
	{
		_ptr_ppVarDesc := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.VarDesc == nil {
				o.VarDesc = &oaut.VarDesc{}
			}
			if err := o.VarDesc.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppVarDesc := func(ptr interface{}) { o.VarDesc = *ptr.(**oaut.VarDesc) }
		if err := w.ReadPointer(&o.VarDesc, _s_ppVarDesc, _ptr_ppVarDesc); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppTypeComp {out} (1:{pointer=ref}*(2)*(1))(2:{alias=ITypeComp}(interface))
	{
		_ptr_ppTypeComp := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.TypeComp == nil {
				o.TypeComp = &oaut.TypeComp{}
			}
			if err := o.TypeComp.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppTypeComp := func(ptr interface{}) { o.TypeComp = *ptr.(**oaut.TypeComp) }
		if err := w.ReadPointer(&o.TypeComp, _s_ppTypeComp, _ptr_ppTypeComp); err != nil {
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

// BindRequest structure represents the Bind operation request
type BindRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// szName: MUST be set to a string.
	Name string `idl:"name:szName" json:"name"`
	// lHashVal: MUST be set to the hash value that corresponds to the value of szName (as
	// specified in section 2.2.51) or 0. Whether 0 or a computed hash value is specified
	// for this argument, the server MUST return the same values in ppTInfo, pDescKind,
	// ppFuncDesc, ppVarDesc, and ppTypeComp.
	HashValue uint32 `idl:"name:lHashVal" json:"hash_value"`
	// wFlags: MUST be set to a value of the INVOKEKIND enumeration, as specified in section
	// 2.2.14, or 0.
	Flags uint16 `idl:"name:wFlags" json:"flags"`
}

func (o *BindRequest) xxx_ToOp(ctx context.Context) *xxx_BindOperation {
	if o == nil {
		return &xxx_BindOperation{}
	}
	return &xxx_BindOperation{
		This:      o.This,
		Name:      o.Name,
		HashValue: o.HashValue,
		Flags:     o.Flags,
	}
}

func (o *BindRequest) xxx_FromOp(ctx context.Context, op *xxx_BindOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Name = op.Name
	o.HashValue = op.HashValue
	o.Flags = op.Flags
}
func (o *BindRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *BindRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BindOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BindResponse structure represents the Bind operation response
type BindResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppTInfo: MUST be set to a reference to the ITypeInfo server instance that corresponds
	// to the element whose name matches the value of szName or NULL, as specified in sections
	// 3.5.4.1.2 and 2.2.50. MUST be set to NULL if szName does not match the name of any
	// element in the binding context (see section 3.5.4.1.1).
	TypeInfo *oaut.TypeInfo `idl:"name:ppTInfo" json:"type_info"`
	// pDescKind: MUST be set to one of the following values of the DESCKIND enumeration
	// (section 2.2.22):
	//
	// * MUST be set to DESCKIND_FUNCDESC if the values of szName and wFlags specify a method
	// or property accessor method in the binding context of the ITypeComp server ( 7894019f-de1e-455e-b2aa-3b899c2e50f6
	// ).
	//
	// * MUST be set to DESCKIND_VARDESC if the values of szName and wFlags specify a data
	// member of an enumeration, module, or ODL dispinterface ( 5583e1b8-454c-4147-9f56-f72416a15bee#gt_544446d5-79ab-4b08-85c4-214f1b64fdf2
	// ) in the binding context of the ITypeComp server, or if the binding server is an
	// ITypeLib server ( 5daecf67-bc6e-4e17-bcf8-797bdba1748b ) and szName specifies the
	// name of an appobject coclass ( 5583e1b8-454c-4147-9f56-f72416a15bee#gt_670b0ee2-d101-41b0-ac77-6ac7dbeee7dc
	// ) in the binding context of the ITypeComp server.
	//
	// * MUST be set to DESCKIND_TYPECOMP if the value of szName specifies an enumeration
	// or module in the binding context of the ITypeComp server.
	//
	// * MUST be set to DESCKIND_IMPLICITAPPOBJ if the binding server is an ITypeLib server,
	// the value of szName specifies an element in the binding context of an appobject coclass,
	// and the appobject coclass is in the binding context of the ITypeComp server.
	//
	// * Otherwise, MUST be set to DESCKIND_NONE.
	DescKind oaut.DescKind `idl:"name:pDescKind" json:"desc_kind"`
	// ppFuncDesc: MUST be set to a FUNCDESC that describes the method or property whose
	// name matches the value of szName, if pDescKind is DESCKIND_FUNCDESC. Otherwise, MUST
	// be set to NULL.
	FuncDesc *oaut.FuncDesc `idl:"name:ppFuncDesc" json:"func_desc"`
	// ppVarDesc: MUST be set to a value that is specified by the value of pDescKind.
	//
	// * DESCKIND_VARDESC: MUST be set to a VARDESC that describes a data member of an enumeration,
	// module, or ODL dispinterface if the name of the data member matches szName.
	//
	// * DESCKIND_IMPLICITAPPOBJ: MUST be set to a VARDESC that describes the appobject
	// coclass whose binding context contains a member whose name matches szName.
	//
	// * DESCKIND_FUNCDESC, DESCKIND_TYPECOMP, DESCKIND_NONE: MUST be set to NULL.
	VarDesc *oaut.VarDesc `idl:"name:ppVarDesc" json:"var_desc"`
	// ppTypeComp: MUST be set to a reference to an ITypeComp server instance that provides
	// access to the binding context of the bound element, if pDescKind is DESCKIND_TYPECOMP
	// or DESCKIND_IMPLICITAPPOBJ. Otherwise, MUST be set to NULL.
	TypeComp *oaut.TypeComp `idl:"name:ppTypeComp" json:"type_comp"`
	// Return: The Bind return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *BindResponse) xxx_ToOp(ctx context.Context) *xxx_BindOperation {
	if o == nil {
		return &xxx_BindOperation{}
	}
	return &xxx_BindOperation{
		That:     o.That,
		TypeInfo: o.TypeInfo,
		DescKind: o.DescKind,
		FuncDesc: o.FuncDesc,
		VarDesc:  o.VarDesc,
		TypeComp: o.TypeComp,
		Return:   o.Return,
	}
}

func (o *BindResponse) xxx_FromOp(ctx context.Context, op *xxx_BindOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TypeInfo = op.TypeInfo
	o.DescKind = op.DescKind
	o.FuncDesc = op.FuncDesc
	o.VarDesc = op.VarDesc
	o.TypeComp = op.TypeComp
	o.Return = op.Return
}
func (o *BindResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *BindResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BindOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BindTypeOperation structure represents the BindType operation
type xxx_BindTypeOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Name      string         `idl:"name:szName" json:"name"`
	HashValue uint32         `idl:"name:lHashVal" json:"hash_value"`
	TypeInfo  *oaut.TypeInfo `idl:"name:ppTInfo" json:"type_info"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_BindTypeOperation) OpNum() int { return 4 }

func (o *xxx_BindTypeOperation) OpName() string { return "/ITypeComp/v0/BindType" }

func (o *xxx_BindTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BindTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// szName {in} (1:{string, alias=LPOLESTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
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

func (o *xxx_BindTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// szName {in} (1:{string, alias=LPOLESTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
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

func (o *xxx_BindTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BindTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_BindTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// BindTypeRequest structure represents the BindType operation request
type BindTypeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// szName: MUST be set to a string.
	Name string `idl:"name:szName" json:"name"`
	// lHashVal: MUST be the hash value associated with the value of szName as specified
	// in section 2.2.51, or 0. The server MUST return the same values in ppTInfo in either
	// case.
	HashValue uint32 `idl:"name:lHashVal" json:"hash_value"`
}

func (o *BindTypeRequest) xxx_ToOp(ctx context.Context) *xxx_BindTypeOperation {
	if o == nil {
		return &xxx_BindTypeOperation{}
	}
	return &xxx_BindTypeOperation{
		This:      o.This,
		Name:      o.Name,
		HashValue: o.HashValue,
	}
}

func (o *BindTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_BindTypeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Name = op.Name
	o.HashValue = op.HashValue
}
func (o *BindTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *BindTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BindTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BindTypeResponse structure represents the BindType operation response
type BindTypeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppTInfo: MUST be set to a reference to an ITypeInfo server instance, or NULL.
	//
	// If the ITypeComp server is associated with an ITypeLib server, ppTInfo MUST specify
	// a type defined in its automation scope whose name matches the value of szName according
	// to the string-matching criteria specified in section 2.2.50, or be set to NULL if
	// no such type exists.
	TypeInfo *oaut.TypeInfo `idl:"name:ppTInfo" json:"type_info"`
	// Return: The BindType return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *BindTypeResponse) xxx_ToOp(ctx context.Context) *xxx_BindTypeOperation {
	if o == nil {
		return &xxx_BindTypeOperation{}
	}
	return &xxx_BindTypeOperation{
		That:     o.That,
		TypeInfo: o.TypeInfo,
		Return:   o.Return,
	}
}

func (o *BindTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_BindTypeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TypeInfo = op.TypeInfo
	o.Return = op.Return
}
func (o *BindTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *BindTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BindTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
