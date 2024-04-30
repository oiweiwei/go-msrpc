package ivdsservicehba

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
	// IVdsServiceHba interface identifier 0ac13689-3134-47c6-a17c-4669216801be
	ServiceHBAIID = &dcom.IID{Data1: 0x0ac13689, Data2: 0x3134, Data3: 0x47c6, Data4: []byte{0xa1, 0x7c, 0x46, 0x69, 0x21, 0x68, 0x01, 0xbe}}
	// Syntax UUID
	ServiceHBASyntaxUUID = &uuid.UUID{TimeLow: 0xac13689, TimeMid: 0x3134, TimeHiAndVersion: 0x47c6, ClockSeqHiAndReserved: 0xa1, ClockSeqLow: 0x7c, Node: [6]uint8{0x46, 0x69, 0x21, 0x68, 0x1, 0xbe}}
	// Syntax ID
	ServiceHBASyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ServiceHBASyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsServiceHba interface.
type ServiceHBAClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The QueryHbaPorts method returns an IEnumVdsObject enumeration object that contains
	// a list of the HBA ports that are known to VDS on the system.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryHBAPorts(context.Context, *QueryHBAPortsRequest, ...dcerpc.CallOption) (*QueryHBAPortsResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ServiceHBAClient
}

type xxx_DefaultServiceHBAClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultServiceHBAClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultServiceHBAClient) QueryHBAPorts(ctx context.Context, in *QueryHBAPortsRequest, opts ...dcerpc.CallOption) (*QueryHBAPortsResponse, error) {
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
	out := &QueryHBAPortsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServiceHBAClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultServiceHBAClient) IPID(ctx context.Context, ipid *dcom.IPID) ServiceHBAClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultServiceHBAClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}
func NewServiceHBAClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ServiceHBAClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ServiceHBASyntaxV0_0))...)
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
	return &xxx_DefaultServiceHBAClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_QueryHBAPortsOperation structure represents the QueryHbaPorts operation
type xxx_QueryHBAPortsOperation struct {
	This   *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Enum   *vds.EnumObject `idl:"name:ppEnum" json:"enum"`
	Return int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryHBAPortsOperation) OpNum() int { return 3 }

func (o *xxx_QueryHBAPortsOperation) OpName() string { return "/IVdsServiceHba/v0/QueryHbaPorts" }

func (o *xxx_QueryHBAPortsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryHBAPortsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_QueryHBAPortsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_QueryHBAPortsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryHBAPortsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_QueryHBAPortsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// QueryHBAPortsRequest structure represents the QueryHbaPorts operation request
type QueryHBAPortsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *QueryHBAPortsRequest) xxx_ToOp(ctx context.Context) *xxx_QueryHBAPortsOperation {
	if o == nil {
		return &xxx_QueryHBAPortsOperation{}
	}
	return &xxx_QueryHBAPortsOperation{
		This: o.This,
	}
}

func (o *QueryHBAPortsRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryHBAPortsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *QueryHBAPortsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryHBAPortsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryHBAPortsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryHBAPortsResponse structure represents the QueryHbaPorts operation response
type QueryHBAPortsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppEnum: A pointer to an IEnumVdsObject interface that, if the operation is successfully
	// completed, receives the IEnumVdsObject interface of the object that contains an enumeration
	// of the HBA port objects on the server. Callers MUST release the interface when they
	// are done with it.
	Enum *vds.EnumObject `idl:"name:ppEnum" json:"enum"`
	// Return: The QueryHbaPorts return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryHBAPortsResponse) xxx_ToOp(ctx context.Context) *xxx_QueryHBAPortsOperation {
	if o == nil {
		return &xxx_QueryHBAPortsOperation{}
	}
	return &xxx_QueryHBAPortsOperation{
		That:   o.That,
		Enum:   o.Enum,
		Return: o.Return,
	}
}

func (o *QueryHBAPortsResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryHBAPortsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enum = op.Enum
	o.Return = op.Return
}
func (o *QueryHBAPortsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryHBAPortsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryHBAPortsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
