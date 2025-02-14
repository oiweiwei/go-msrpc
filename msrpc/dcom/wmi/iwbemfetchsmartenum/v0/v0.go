package iwbemfetchsmartenum

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
	wmi "github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi"
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
	_ = wmi.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wmi"
)

var (
	// IWbemFetchSmartEnum interface identifier 1c1c45ee-4395-11d2-b60b-00104b703efd
	FetchSmartEnumIID = &dcom.IID{Data1: 0x1c1c45ee, Data2: 0x4395, Data3: 0x11d2, Data4: []byte{0xb6, 0x0b, 0x00, 0x10, 0x4b, 0x70, 0x3e, 0xfd}}
	// Syntax UUID
	FetchSmartEnumSyntaxUUID = &uuid.UUID{TimeLow: 0x1c1c45ee, TimeMid: 0x4395, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0xb6, ClockSeqLow: 0xb, Node: [6]uint8{0x0, 0x10, 0x4b, 0x70, 0x3e, 0xfd}}
	// Syntax ID
	FetchSmartEnumSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: FetchSmartEnumSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IWbemFetchSmartEnum interface.
type FetchSmartEnumClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The IWbemFetchSmartEnum::GetSmartEnum method retrieves an IWbemWCOSmartEnum (section
	// 3.1.4.7) interface, which is a network-optimized enumerator interface.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	GetSmartEnum(context.Context, *GetSmartEnumRequest, ...dcerpc.CallOption) (*GetSmartEnumResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) FetchSmartEnumClient
}

type xxx_DefaultFetchSmartEnumClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultFetchSmartEnumClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultFetchSmartEnumClient) GetSmartEnum(ctx context.Context, in *GetSmartEnumRequest, opts ...dcerpc.CallOption) (*GetSmartEnumResponse, error) {
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
	out := &GetSmartEnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFetchSmartEnumClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultFetchSmartEnumClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultFetchSmartEnumClient) IPID(ctx context.Context, ipid *dcom.IPID) FetchSmartEnumClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultFetchSmartEnumClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewFetchSmartEnumClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (FetchSmartEnumClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(FetchSmartEnumSyntaxV0_0))...)
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
	return &xxx_DefaultFetchSmartEnumClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetSmartEnumOperation structure represents the GetSmartEnum operation
type xxx_GetSmartEnumOperation struct {
	This      *dcom.ORPCThis    `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat    `idl:"name:That" json:"that"`
	SmartEnum *wmi.WCOSmartEnum `idl:"name:ppSmartEnum" json:"smart_enum"`
	Return    int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSmartEnumOperation) OpNum() int { return 3 }

func (o *xxx_GetSmartEnumOperation) OpName() string { return "/IWbemFetchSmartEnum/v0/GetSmartEnum" }

func (o *xxx_GetSmartEnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSmartEnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSmartEnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSmartEnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSmartEnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppSmartEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IWbemWCOSmartEnum}(interface))
	{
		if o.SmartEnum != nil {
			_ptr_ppSmartEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SmartEnum != nil {
					if err := o.SmartEnum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.WCOSmartEnum{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SmartEnum, _ptr_ppSmartEnum); err != nil {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSmartEnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppSmartEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IWbemWCOSmartEnum}(interface))
	{
		_ptr_ppSmartEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SmartEnum == nil {
				o.SmartEnum = &wmi.WCOSmartEnum{}
			}
			if err := o.SmartEnum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppSmartEnum := func(ptr interface{}) { o.SmartEnum = *ptr.(**wmi.WCOSmartEnum) }
		if err := w.ReadPointer(&o.SmartEnum, _s_ppSmartEnum, _ptr_ppSmartEnum); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetSmartEnumRequest structure represents the GetSmartEnum operation request
type GetSmartEnumRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSmartEnumRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSmartEnumOperation) *xxx_GetSmartEnumOperation {
	if op == nil {
		op = &xxx_GetSmartEnumOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetSmartEnumRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSmartEnumOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSmartEnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSmartEnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSmartEnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSmartEnumResponse structure represents the GetSmartEnum operation response
type GetSmartEnumResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppSmartEnum: MUST be a pointer to a network-optimized enumerator interface. This
	// parameter MUST NOT be NULL. Upon return by the server, this parameter can be NULL
	// if there is a failure or if there are no results.
	SmartEnum *wmi.WCOSmartEnum `idl:"name:ppSmartEnum" json:"smart_enum"`
	// Return: The GetSmartEnum return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSmartEnumResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSmartEnumOperation) *xxx_GetSmartEnumOperation {
	if op == nil {
		op = &xxx_GetSmartEnumOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.SmartEnum = o.SmartEnum
	op.Return = o.Return
	return op
}

func (o *GetSmartEnumResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSmartEnumOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SmartEnum = op.SmartEnum
	o.Return = op.Return
}
func (o *GetSmartEnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSmartEnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSmartEnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
