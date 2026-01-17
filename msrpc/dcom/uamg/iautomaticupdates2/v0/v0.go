package iautomaticupdates2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	uamg "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg"
	iautomaticupdates "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iautomaticupdates/v0"
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
	_ = iautomaticupdates.GoPackage
	_ = uamg.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/uamg"
)

var (
	// IAutomaticUpdates2 interface identifier 4a2f5c31-cfd9-410e-b7fb-29a653973a0f
	AutomaticUpdates2IID = &dcom.IID{Data1: 0x4a2f5c31, Data2: 0xcfd9, Data3: 0x410e, Data4: []byte{0xb7, 0xfb, 0x29, 0xa6, 0x53, 0x97, 0x3a, 0x0f}}
	// Syntax UUID
	AutomaticUpdates2SyntaxUUID = &uuid.UUID{TimeLow: 0x4a2f5c31, TimeMid: 0xcfd9, TimeHiAndVersion: 0x410e, ClockSeqHiAndReserved: 0xb7, ClockSeqLow: 0xfb, Node: [6]uint8{0x29, 0xa6, 0x53, 0x97, 0x3a, 0xf}}
	// Syntax ID
	AutomaticUpdates2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: AutomaticUpdates2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IAutomaticUpdates2 interface.
type AutomaticUpdates2Client interface {

	// IAutomaticUpdates retrieval method.
	AutomaticUpdates() iautomaticupdates.AutomaticUpdatesClient

	// The IAutomaticUpdates2::Results (opnum 15) method retrieves an IAutomaticUpdatesResults
	// instance.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the Results ADM element.
	GetResults(context.Context, *GetResultsRequest, ...dcerpc.CallOption) (*GetResultsResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) AutomaticUpdates2Client
}

type xxx_DefaultAutomaticUpdates2Client struct {
	iautomaticupdates.AutomaticUpdatesClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultAutomaticUpdates2Client) AutomaticUpdates() iautomaticupdates.AutomaticUpdatesClient {
	return o.AutomaticUpdatesClient
}

func (o *xxx_DefaultAutomaticUpdates2Client) GetResults(ctx context.Context, in *GetResultsRequest, opts ...dcerpc.CallOption) (*GetResultsResponse, error) {
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
	out := &GetResultsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAutomaticUpdates2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultAutomaticUpdates2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultAutomaticUpdates2Client) IPID(ctx context.Context, ipid *dcom.IPID) AutomaticUpdates2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultAutomaticUpdates2Client{
		AutomaticUpdatesClient: o.AutomaticUpdatesClient.IPID(ctx, ipid),
		cc:                     o.cc,
		ipid:                   ipid,
	}
}

func NewAutomaticUpdates2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (AutomaticUpdates2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(AutomaticUpdates2SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iautomaticupdates.NewAutomaticUpdatesClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultAutomaticUpdates2Client{
		AutomaticUpdatesClient: base,
		cc:                     cc,
		ipid:                   ipid,
	}, nil
}

// xxx_GetResultsOperation structure represents the Results operation
type xxx_GetResultsOperation struct {
	This        *dcom.ORPCThis                `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat                `idl:"name:That" json:"that"`
	ReturnValue *uamg.AutomaticUpdatesResults `idl:"name:retval" json:"return_value"`
	Return      int32                         `idl:"name:Return" json:"return"`
}

func (o *xxx_GetResultsOperation) OpNum() int { return 14 }

func (o *xxx_GetResultsOperation) OpName() string { return "/IAutomaticUpdates2/v0/Results" }

func (o *xxx_GetResultsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResultsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetResultsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetResultsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResultsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// retval {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAutomaticUpdatesResults}(interface))
	{
		if o.ReturnValue != nil {
			_ptr_retval := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReturnValue != nil {
					if err := o.ReturnValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&uamg.AutomaticUpdatesResults{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReturnValue, _ptr_retval); err != nil {
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

func (o *xxx_GetResultsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// retval {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAutomaticUpdatesResults}(interface))
	{
		_ptr_retval := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReturnValue == nil {
				o.ReturnValue = &uamg.AutomaticUpdatesResults{}
			}
			if err := o.ReturnValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_retval := func(ptr interface{}) { o.ReturnValue = *ptr.(**uamg.AutomaticUpdatesResults) }
		if err := w.ReadPointer(&o.ReturnValue, _s_retval, _ptr_retval); err != nil {
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

// GetResultsRequest structure represents the Results operation request
type GetResultsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetResultsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetResultsOperation) *xxx_GetResultsOperation {
	if op == nil {
		op = &xxx_GetResultsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetResultsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetResultsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetResultsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetResultsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResultsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetResultsResponse structure represents the Results operation response
type GetResultsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// retval: MUST be set to an instance of the IAutomaticUpdatesResults interface.
	ReturnValue *uamg.AutomaticUpdatesResults `idl:"name:retval" json:"return_value"`
	// Return: The Results return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetResultsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetResultsOperation) *xxx_GetResultsOperation {
	if op == nil {
		op = &xxx_GetResultsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *GetResultsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetResultsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *GetResultsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetResultsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResultsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
