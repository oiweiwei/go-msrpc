package iwrmprotocol

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
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
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
	_ = idispatch.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wsrm"
)

var (
	// IWRMProtocol interface identifier f31931a9-832d-481c-9503-887a0e6a79f0
	ProtocolIID = &dcom.IID{Data1: 0xf31931a9, Data2: 0x832d, Data3: 0x481c, Data4: []byte{0x95, 0x03, 0x88, 0x7a, 0x0e, 0x6a, 0x79, 0xf0}}
	// Syntax UUID
	ProtocolSyntaxUUID = &uuid.UUID{TimeLow: 0xf31931a9, TimeMid: 0x832d, TimeHiAndVersion: 0x481c, ClockSeqHiAndReserved: 0x95, ClockSeqLow: 0x3, Node: [6]uint8{0x88, 0x7a, 0xe, 0x6a, 0x79, 0xf0}}
	// Syntax ID
	ProtocolSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ProtocolSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IWRMProtocol interface.
type ProtocolClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// The GetSupportedClient method retrieves the level of support for clients on the WSRM
	// server.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------+-----------------------+
	//	|      RETURN       |                       |
	//	|    VALUE/CODE     |      DESCRIPTION      |
	//	|                   |                       |
	//	+-------------------+-----------------------+
	//	+-------------------+-----------------------+
	//	| 0x00000000 S_OK   | Operation successful. |
	//	+-------------------+-----------------------+
	//
	// Additional IWRMProtocol interface methods are specified in section 3.2.4.8.
	GetSupportedClient(context.Context, *GetSupportedClientRequest, ...dcerpc.CallOption) (*GetSupportedClientResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ProtocolClient
}

type xxx_DefaultProtocolClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultProtocolClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultProtocolClient) GetSupportedClient(ctx context.Context, in *GetSupportedClientRequest, opts ...dcerpc.CallOption) (*GetSupportedClientResponse, error) {
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
	out := &GetSupportedClientResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultProtocolClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultProtocolClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultProtocolClient) IPID(ctx context.Context, ipid *dcom.IPID) ProtocolClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultProtocolClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewProtocolClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ProtocolClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ProtocolSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := idispatch.NewDispatchClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultProtocolClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetSupportedClientOperation structure represents the GetSupportedClient operation
type xxx_GetSupportedClientOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	SupportedClients *oaut.String   `idl:"name:pbstrSupportedClients" json:"supported_clients"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSupportedClientOperation) OpNum() int { return 7 }

func (o *xxx_GetSupportedClientOperation) OpName() string {
	return "/IWRMProtocol/v0/GetSupportedClient"
}

func (o *xxx_GetSupportedClientOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSupportedClientOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSupportedClientOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSupportedClientOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSupportedClientOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrSupportedClients {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SupportedClients != nil {
			_ptr_pbstrSupportedClients := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SupportedClients != nil {
					if err := o.SupportedClients.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SupportedClients, _ptr_pbstrSupportedClients); err != nil {
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

func (o *xxx_GetSupportedClientOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrSupportedClients {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrSupportedClients := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SupportedClients == nil {
				o.SupportedClients = &oaut.String{}
			}
			if err := o.SupportedClients.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrSupportedClients := func(ptr interface{}) { o.SupportedClients = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SupportedClients, _s_pbstrSupportedClients, _ptr_pbstrSupportedClients); err != nil {
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

// GetSupportedClientRequest structure represents the GetSupportedClient operation request
type GetSupportedClientRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSupportedClientRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSupportedClientOperation) *xxx_GetSupportedClientOperation {
	if op == nil {
		op = &xxx_GetSupportedClientOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetSupportedClientRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSupportedClientOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSupportedClientRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSupportedClientRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSupportedClientOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSupportedClientResponse structure represents the GetSupportedClient operation response
type GetSupportedClientResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbstrSupportedClients: A pointer to a string that returns the level of support for
	// clients, in the format specified in SupportedClients (section 2.2.5.29).
	SupportedClients *oaut.String `idl:"name:pbstrSupportedClients" json:"supported_clients"`
	// Return: The GetSupportedClient return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSupportedClientResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSupportedClientOperation) *xxx_GetSupportedClientOperation {
	if op == nil {
		op = &xxx_GetSupportedClientOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.SupportedClients = o.SupportedClients
	op.Return = o.Return
	return op
}

func (o *GetSupportedClientResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSupportedClientOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SupportedClients = op.SupportedClients
	o.Return = op.Return
}
func (o *GetSupportedClientResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSupportedClientResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSupportedClientOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
