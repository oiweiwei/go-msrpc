package ivdsswprovider

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
	vds "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds"
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
	_ = vds.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/vds"
)

var (
	// IVdsSwProvider interface identifier 9aa58360-ce33-4f92-b658-ed24b14425b8
	SwProviderIID = &dcom.IID{Data1: 0x9aa58360, Data2: 0xce33, Data3: 0x4f92, Data4: []byte{0xb6, 0x58, 0xed, 0x24, 0xb1, 0x44, 0x25, 0xb8}}
	// Syntax UUID
	SwProviderSyntaxUUID = &uuid.UUID{TimeLow: 0x9aa58360, TimeMid: 0xce33, TimeHiAndVersion: 0x4f92, ClockSeqHiAndReserved: 0xb6, ClockSeqLow: 0x58, Node: [6]uint8{0xed, 0x24, 0xb1, 0x44, 0x25, 0xb8}}
	// Syntax ID
	SwProviderSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: SwProviderSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsSwProvider interface.
type SwProviderClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The QueryPacks method retrieves the provider disk packs.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryPacks(context.Context, *QueryPacksRequest, ...dcerpc.CallOption) (*QueryPacksResponse, error)

	// The CreatePack method creates a disk pack.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	CreatePack(context.Context, *CreatePackRequest, ...dcerpc.CallOption) (*CreatePackResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) SwProviderClient
}

type xxx_DefaultSwProviderClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultSwProviderClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultSwProviderClient) QueryPacks(ctx context.Context, in *QueryPacksRequest, opts ...dcerpc.CallOption) (*QueryPacksResponse, error) {
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
	out := &QueryPacksResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSwProviderClient) CreatePack(ctx context.Context, in *CreatePackRequest, opts ...dcerpc.CallOption) (*CreatePackResponse, error) {
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
	out := &CreatePackResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSwProviderClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultSwProviderClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultSwProviderClient) IPID(ctx context.Context, ipid *dcom.IPID) SwProviderClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultSwProviderClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewSwProviderClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (SwProviderClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(SwProviderSyntaxV0_0))...)
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
	return &xxx_DefaultSwProviderClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_QueryPacksOperation structure represents the QueryPacks operation
type xxx_QueryPacksOperation struct {
	This   *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Enum   *vds.EnumObject `idl:"name:ppEnum" json:"enum"`
	Return int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryPacksOperation) OpNum() int { return 3 }

func (o *xxx_QueryPacksOperation) OpName() string { return "/IVdsSwProvider/v0/QueryPacks" }

func (o *xxx_QueryPacksOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryPacksOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_QueryPacksOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_QueryPacksOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryPacksOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumVdsObject}(interface))
	{
		if o.Enum != nil {
			_ptr_ppEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Enum != nil {
					if err := o.Enum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&vds.EnumObject{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Enum, _ptr_ppEnum); err != nil {
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

func (o *xxx_QueryPacksOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumVdsObject}(interface))
	{
		_ptr_ppEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Enum == nil {
				o.Enum = &vds.EnumObject{}
			}
			if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppEnum := func(ptr interface{}) { o.Enum = *ptr.(**vds.EnumObject) }
		if err := w.ReadPointer(&o.Enum, _s_ppEnum, _ptr_ppEnum); err != nil {
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

// QueryPacksRequest structure represents the QueryPacks operation request
type QueryPacksRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *QueryPacksRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryPacksOperation) *xxx_QueryPacksOperation {
	if op == nil {
		op = &xxx_QueryPacksOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *QueryPacksRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryPacksOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *QueryPacksRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryPacksRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryPacksOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryPacksResponse structure represents the QueryPacks operation response
type QueryPacksResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppEnum: A pointer to an IEnumVdsObject interface that, if the operation is successfully
	// completed, receives the IEnumVdsObject interface of the object containing an enumeration
	// of pack objects in the provider. Callers MUST release the interface when they are
	// done with it.
	Enum *vds.EnumObject `idl:"name:ppEnum" json:"enum"`
	// Return: The QueryPacks return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryPacksResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryPacksOperation) *xxx_QueryPacksOperation {
	if op == nil {
		op = &xxx_QueryPacksOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Enum = op.Enum
	o.Return = op.Return
	return op
}

func (o *QueryPacksResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryPacksOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enum = op.Enum
	o.Return = op.Return
}
func (o *QueryPacksResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryPacksResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryPacksOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreatePackOperation structure represents the CreatePack operation
type xxx_CreatePackOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Pack   *vds.Pack      `idl:"name:ppPack" json:"pack"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CreatePackOperation) OpNum() int { return 4 }

func (o *xxx_CreatePackOperation) OpName() string { return "/IVdsSwProvider/v0/CreatePack" }

func (o *xxx_CreatePackOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePackOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreatePackOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_CreatePackOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePackOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppPack {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsPack}(interface))
	{
		if o.Pack != nil {
			_ptr_ppPack := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Pack != nil {
					if err := o.Pack.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&vds.Pack{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Pack, _ptr_ppPack); err != nil {
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

func (o *xxx_CreatePackOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppPack {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsPack}(interface))
	{
		_ptr_ppPack := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Pack == nil {
				o.Pack = &vds.Pack{}
			}
			if err := o.Pack.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppPack := func(ptr interface{}) { o.Pack = *ptr.(**vds.Pack) }
		if err := w.ReadPointer(&o.Pack, _s_ppPack, _ptr_ppPack); err != nil {
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

// CreatePackRequest structure represents the CreatePack operation request
type CreatePackRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *CreatePackRequest) xxx_ToOp(ctx context.Context, op *xxx_CreatePackOperation) *xxx_CreatePackOperation {
	if op == nil {
		op = &xxx_CreatePackOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *CreatePackRequest) xxx_FromOp(ctx context.Context, op *xxx_CreatePackOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *CreatePackRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreatePackRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreatePackOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreatePackResponse structure represents the CreatePack operation response
type CreatePackResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppPack: A pointer to an IVdsPack interface that, if the operation is successfully
	// completed, receives the IVdsPack interface of the newly created disk pack. Callers
	// MUST release the interface when they are done with it.
	Pack *vds.Pack `idl:"name:ppPack" json:"pack"`
	// Return: The CreatePack return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreatePackResponse) xxx_ToOp(ctx context.Context, op *xxx_CreatePackOperation) *xxx_CreatePackOperation {
	if op == nil {
		op = &xxx_CreatePackOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Pack = op.Pack
	o.Return = op.Return
	return op
}

func (o *CreatePackResponse) xxx_FromOp(ctx context.Context, op *xxx_CreatePackOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Pack = op.Pack
	o.Return = op.Return
}
func (o *CreatePackResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreatePackResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreatePackOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
