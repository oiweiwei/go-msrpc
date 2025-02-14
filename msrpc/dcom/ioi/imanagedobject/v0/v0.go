package imanagedobject

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
	GoPackage = "dcom/ioi"
)

var (
	// IManagedObject interface identifier c3fcc19e-a970-11d2-8b5a-00a0c9b7c9c4
	ManagedObjectIID = &dcom.IID{Data1: 0xc3fcc19e, Data2: 0xa970, Data3: 0x11d2, Data4: []byte{0x8b, 0x5a, 0x00, 0xa0, 0xc9, 0xb7, 0xc9, 0xc4}}
	// Syntax UUID
	ManagedObjectSyntaxUUID = &uuid.UUID{TimeLow: 0xc3fcc19e, TimeMid: 0xa970, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0x8b, ClockSeqLow: 0x5a, Node: [6]uint8{0x0, 0xa0, 0xc9, 0xb7, 0xc9, 0xc4}}
	// Syntax ID
	ManagedObjectSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ManagedObjectSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IManagedObject interface.
type ManagedObjectClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The GetSerializedBuffer method converts the given managed object to a binary-formatted
	// string representation that can be used to create a managed object.
	//
	// Return Values: The method MUST return a positive value or 0 to indicate successful
	// completion or a negative value to indicate failure.
	//
	//	+--------------------------+-------------+
	//	|          RETURN          |             |
	//	|        VALUE/CODE        | DESCRIPTION |
	//	|                          |             |
	//	+--------------------------+-------------+
	//	+--------------------------+-------------+
	//	| 0x00000000 ERROR_SUCCESS | Success.    |
	//	+--------------------------+-------------+
	//
	// Exceptions Thrown: No exceptions are thrown from this method beyond those thrown
	// by the underlying RPC protocol.
	GetSerializedBuffer(context.Context, *GetSerializedBufferRequest, ...dcerpc.CallOption) (*GetSerializedBufferResponse, error)

	// The IManagedObject::GetObjectIdentity method is used by a CLR instance to determine
	// whether a COM object entering the system is really a managed object that originated
	// in this CLR instance and within the current process division.
	//
	// Return Values: The method MUST return a positive value or 0 to indicate successful
	// completion or a negative value to indicate failure.
	//
	//	+--------------------------+-------------+
	//	|          RETURN          |             |
	//	|        VALUE/CODE        | DESCRIPTION |
	//	|                          |             |
	//	+--------------------------+-------------+
	//	+--------------------------+-------------+
	//	| 0x00000000 ERROR_SUCCESS | Success     |
	//	+--------------------------+-------------+
	//
	// Exceptions Thrown: No exceptions are thrown from this method beyond those thrown
	// by the underlying RPC protocol.
	GetObjectIdentity(context.Context, *GetObjectIdentityRequest, ...dcerpc.CallOption) (*GetObjectIdentityResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ManagedObjectClient
}

type xxx_DefaultManagedObjectClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultManagedObjectClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultManagedObjectClient) GetSerializedBuffer(ctx context.Context, in *GetSerializedBufferRequest, opts ...dcerpc.CallOption) (*GetSerializedBufferResponse, error) {
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
	out := &GetSerializedBufferResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultManagedObjectClient) GetObjectIdentity(ctx context.Context, in *GetObjectIdentityRequest, opts ...dcerpc.CallOption) (*GetObjectIdentityResponse, error) {
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
	out := &GetObjectIdentityResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultManagedObjectClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultManagedObjectClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultManagedObjectClient) IPID(ctx context.Context, ipid *dcom.IPID) ManagedObjectClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultManagedObjectClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewManagedObjectClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ManagedObjectClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ManagedObjectSyntaxV0_0))...)
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
	return &xxx_DefaultManagedObjectClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetSerializedBufferOperation structure represents the GetSerializedBuffer operation
type xxx_GetSerializedBufferOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	String *oaut.String   `idl:"name:pBSTR" json:"string"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSerializedBufferOperation) OpNum() int { return 3 }

func (o *xxx_GetSerializedBufferOperation) OpName() string {
	return "/IManagedObject/v0/GetSerializedBuffer"
}

func (o *xxx_GetSerializedBufferOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSerializedBufferOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSerializedBufferOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSerializedBufferOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSerializedBufferOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pBSTR {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.String != nil {
			_ptr_pBSTR := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.String != nil {
					if err := o.String.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.String, _ptr_pBSTR); err != nil {
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

func (o *xxx_GetSerializedBufferOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pBSTR {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pBSTR := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.String == nil {
				o.String = &oaut.String{}
			}
			if err := o.String.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pBSTR := func(ptr interface{}) { o.String = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.String, _s_pBSTR, _ptr_pBSTR); err != nil {
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

// GetSerializedBufferRequest structure represents the GetSerializedBuffer operation request
type GetSerializedBufferRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSerializedBufferRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSerializedBufferOperation) *xxx_GetSerializedBufferOperation {
	if op == nil {
		op = &xxx_GetSerializedBufferOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *GetSerializedBufferRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSerializedBufferOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSerializedBufferRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSerializedBufferRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSerializedBufferOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSerializedBufferResponse structure represents the GetSerializedBuffer operation response
type GetSerializedBufferResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	String *oaut.String   `idl:"name:pBSTR" json:"string"`
	// Return: The GetSerializedBuffer return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSerializedBufferResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSerializedBufferOperation) *xxx_GetSerializedBufferOperation {
	if op == nil {
		op = &xxx_GetSerializedBufferOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.String = op.String
	o.Return = op.Return
	return op
}

func (o *GetSerializedBufferResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSerializedBufferOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.String = op.String
	o.Return = op.Return
}
func (o *GetSerializedBufferResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSerializedBufferResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSerializedBufferOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetObjectIdentityOperation structure represents the GetObjectIdentity operation
type xxx_GetObjectIdentityOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	GUID        *oaut.String   `idl:"name:pBSTRGUID" json:"guid"`
	AppDomainID int32          `idl:"name:AppDomainID" json:"app_domain_id"`
	CCW         int64          `idl:"name:pCCW" json:"ccw"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetObjectIdentityOperation) OpNum() int { return 4 }

func (o *xxx_GetObjectIdentityOperation) OpName() string {
	return "/IManagedObject/v0/GetObjectIdentity"
}

func (o *xxx_GetObjectIdentityOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectIdentityOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetObjectIdentityOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetObjectIdentityOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectIdentityOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pBSTRGUID {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.GUID != nil {
			_ptr_pBSTRGUID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.GUID != nil {
					if err := o.GUID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.GUID, _ptr_pBSTRGUID); err != nil {
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
	// AppDomainID {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.AppDomainID); err != nil {
			return err
		}
	}
	// pCCW {out} (1:{alias=CCW_PTR}*(1)(int64))
	{
		if err := w.WriteData(o.CCW); err != nil {
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

func (o *xxx_GetObjectIdentityOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pBSTRGUID {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pBSTRGUID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.GUID == nil {
				o.GUID = &oaut.String{}
			}
			if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pBSTRGUID := func(ptr interface{}) { o.GUID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.GUID, _s_pBSTRGUID, _ptr_pBSTRGUID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// AppDomainID {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.AppDomainID); err != nil {
			return err
		}
	}
	// pCCW {out} (1:{alias=CCW_PTR,pointer=ref}*(1)(int64))
	{
		if err := w.ReadData(&o.CCW); err != nil {
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

// GetObjectIdentityRequest structure represents the GetObjectIdentity operation request
type GetObjectIdentityRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetObjectIdentityRequest) xxx_ToOp(ctx context.Context, op *xxx_GetObjectIdentityOperation) *xxx_GetObjectIdentityOperation {
	if op == nil {
		op = &xxx_GetObjectIdentityOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *GetObjectIdentityRequest) xxx_FromOp(ctx context.Context, op *xxx_GetObjectIdentityOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetObjectIdentityRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetObjectIdentityRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetObjectIdentityOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetObjectIdentityResponse structure represents the GetObjectIdentity operation response
type GetObjectIdentityResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pBSTRGUID: The pBSTRGUID parameter is a GUID ([MS-DTYP] section 2.3.4.3). The pBSTRGUID
	// parameter MUST indicate the CLR instance in which this object was created.
	GUID *oaut.String `idl:"name:pBSTRGUID" json:"guid"`
	// AppDomainID: Optional parameter that contains implementation-specific, opaque, process-unique
	// identifiers. If present, the AppDomainID parameter MUST denote the process subdivision
	// in which this object resides.
	AppDomainID int32 `idl:"name:AppDomainID" json:"app_domain_id"`
	// pCCW: Optional field. Implementation-specific, opaque value that helps identify the
	// managed object. If present, this field MUST map back to the implementation's internal
	// representation of a managed object.
	CCW int64 `idl:"name:pCCW" json:"ccw"`
	// Return: The GetObjectIdentity return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetObjectIdentityResponse) xxx_ToOp(ctx context.Context, op *xxx_GetObjectIdentityOperation) *xxx_GetObjectIdentityOperation {
	if op == nil {
		op = &xxx_GetObjectIdentityOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.GUID = op.GUID
	o.AppDomainID = op.AppDomainID
	o.CCW = op.CCW
	o.Return = op.Return
	return op
}

func (o *GetObjectIdentityResponse) xxx_FromOp(ctx context.Context, op *xxx_GetObjectIdentityOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.GUID = op.GUID
	o.AppDomainID = op.AppDomainID
	o.CCW = op.CCW
	o.Return = op.Return
}
func (o *GetObjectIdentityResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetObjectIdentityResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetObjectIdentityOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
