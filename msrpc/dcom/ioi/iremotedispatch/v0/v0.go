package iremotedispatch

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = dcom.GoPackage
	_ = idispatch.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/ioi"
)

var (
	// IRemoteDispatch interface identifier 6619a740-8154-43be-a186-0319578e02db
	RemoteDispatchIID = &dcom.IID{Data1: 0x6619a740, Data2: 0x8154, Data3: 0x43be, Data4: []byte{0xa1, 0x86, 0x03, 0x19, 0x57, 0x8e, 0x02, 0xdb}}
	// Syntax UUID
	RemoteDispatchSyntaxUUID = &uuid.UUID{TimeLow: 0x6619a740, TimeMid: 0x8154, TimeHiAndVersion: 0x43be, ClockSeqHiAndReserved: 0xa1, ClockSeqLow: 0x86, Node: [6]uint8{0x3, 0x19, 0x57, 0x8e, 0x2, 0xdb}}
	// Syntax ID
	RemoteDispatchSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: RemoteDispatchSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IRemoteDispatch interface.
type RemoteDispatchClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// The RemoteDispatchAutoDone method is called by the client to invoke a method on the
	// server.
	//
	// Return Values: An HRESULT that specifies success or failure. All success HRESULT
	// values MUST be treated as success and all failure HRESULT values MUST be treated
	// as failure.
	RemoteDispatchAutoDone(context.Context, *RemoteDispatchAutoDoneRequest, ...dcerpc.CallOption) (*RemoteDispatchAutoDoneResponse, error)

	// The RemoteDispatchNotAutoDone method is called by the client to invoke a method on
	// the server.
	//
	// Return Values: An HRESULT that specifies success or failure. All success HRESULT
	// values MUST be treated as success and all failure HRESULT values MUST be treated
	// as failure.
	RemoteDispatchNotAutoDone(context.Context, *RemoteDispatchNotAutoDoneRequest, ...dcerpc.CallOption) (*RemoteDispatchNotAutoDoneResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) RemoteDispatchClient
}

type xxx_DefaultRemoteDispatchClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultRemoteDispatchClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultRemoteDispatchClient) RemoteDispatchAutoDone(ctx context.Context, in *RemoteDispatchAutoDoneRequest, opts ...dcerpc.CallOption) (*RemoteDispatchAutoDoneResponse, error) {
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
	out := &RemoteDispatchAutoDoneResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteDispatchClient) RemoteDispatchNotAutoDone(ctx context.Context, in *RemoteDispatchNotAutoDoneRequest, opts ...dcerpc.CallOption) (*RemoteDispatchNotAutoDoneResponse, error) {
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
	out := &RemoteDispatchNotAutoDoneResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteDispatchClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultRemoteDispatchClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultRemoteDispatchClient) IPID(ctx context.Context, ipid *dcom.IPID) RemoteDispatchClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultRemoteDispatchClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewRemoteDispatchClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (RemoteDispatchClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(RemoteDispatchSyntaxV0_0))...)
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
	return &xxx_DefaultRemoteDispatchClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_RemoteDispatchAutoDoneOperation structure represents the RemoteDispatchAutoDone operation
type xxx_RemoteDispatchAutoDoneOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	String      *oaut.String   `idl:"name:s" json:"string"`
	ReturnValue *oaut.String   `idl:"name:pRetVal" json:"return_value"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoteDispatchAutoDoneOperation) OpNum() int { return 7 }

func (o *xxx_RemoteDispatchAutoDoneOperation) OpName() string {
	return "/IRemoteDispatch/v0/RemoteDispatchAutoDone"
}

func (o *xxx_RemoteDispatchAutoDoneOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteDispatchAutoDoneOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// s {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.String != nil {
			_ptr_s := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
			if err := w.WritePointer(&o.String, _ptr_s); err != nil {
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

func (o *xxx_RemoteDispatchAutoDoneOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// s {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_s := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.String == nil {
				o.String = &oaut.String{}
			}
			if err := o.String.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_s := func(ptr interface{}) { o.String = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.String, _s_s, _ptr_s); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteDispatchAutoDoneOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteDispatchAutoDoneOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pRetVal {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ReturnValue != nil {
			_ptr_pRetVal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReturnValue != nil {
					if err := o.ReturnValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReturnValue, _ptr_pRetVal); err != nil {
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

func (o *xxx_RemoteDispatchAutoDoneOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pRetVal {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pRetVal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReturnValue == nil {
				o.ReturnValue = &oaut.String{}
			}
			if err := o.ReturnValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pRetVal := func(ptr interface{}) { o.ReturnValue = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ReturnValue, _s_pRetVal, _ptr_pRetVal); err != nil {
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

// RemoteDispatchAutoDoneRequest structure represents the RemoteDispatchAutoDone operation request
type RemoteDispatchAutoDoneRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// s: The s parameter contains binary data representing the input parameters of the
	// method called on the server. The binary data MUST be marshaled as specified in [MS-NRTP]
	// section 3.1.5.1.1. The data is specified as is in the BSTR, such that the length
	// of the BSTR is the size of the data divided by 2 (rounded up if necessary).
	String *oaut.String `idl:"name:s" json:"string"`
}

func (o *RemoteDispatchAutoDoneRequest) xxx_ToOp(ctx context.Context, op *xxx_RemoteDispatchAutoDoneOperation) *xxx_RemoteDispatchAutoDoneOperation {
	if op == nil {
		op = &xxx_RemoteDispatchAutoDoneOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.String = o.String
	return op
}

func (o *RemoteDispatchAutoDoneRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoteDispatchAutoDoneOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.String = op.String
}
func (o *RemoteDispatchAutoDoneRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RemoteDispatchAutoDoneRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteDispatchAutoDoneOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoteDispatchAutoDoneResponse structure represents the RemoteDispatchAutoDone operation response
type RemoteDispatchAutoDoneResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pRetVal: The pRetVal parameter contains the binary data representing the output parameters
	// of the method called on the server. The binary data MUST be marshaled as specified
	// in [MS-NRTP] section 3.1.5.1.1. The data is specified as is in the BSTR, such that
	// the length of the BSTR is the size of the data divided by 2 (rounded up if necessary).
	ReturnValue *oaut.String `idl:"name:pRetVal" json:"return_value"`
	// Return: The RemoteDispatchAutoDone return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemoteDispatchAutoDoneResponse) xxx_ToOp(ctx context.Context, op *xxx_RemoteDispatchAutoDoneOperation) *xxx_RemoteDispatchAutoDoneOperation {
	if op == nil {
		op = &xxx_RemoteDispatchAutoDoneOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *RemoteDispatchAutoDoneResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoteDispatchAutoDoneOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *RemoteDispatchAutoDoneResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RemoteDispatchAutoDoneResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteDispatchAutoDoneOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoteDispatchNotAutoDoneOperation structure represents the RemoteDispatchNotAutoDone operation
type xxx_RemoteDispatchNotAutoDoneOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	String      *oaut.String   `idl:"name:s" json:"string"`
	ReturnValue *oaut.String   `idl:"name:pRetVal" json:"return_value"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoteDispatchNotAutoDoneOperation) OpNum() int { return 8 }

func (o *xxx_RemoteDispatchNotAutoDoneOperation) OpName() string {
	return "/IRemoteDispatch/v0/RemoteDispatchNotAutoDone"
}

func (o *xxx_RemoteDispatchNotAutoDoneOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteDispatchNotAutoDoneOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// s {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.String != nil {
			_ptr_s := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
			if err := w.WritePointer(&o.String, _ptr_s); err != nil {
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

func (o *xxx_RemoteDispatchNotAutoDoneOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// s {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_s := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.String == nil {
				o.String = &oaut.String{}
			}
			if err := o.String.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_s := func(ptr interface{}) { o.String = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.String, _s_s, _ptr_s); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteDispatchNotAutoDoneOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteDispatchNotAutoDoneOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pRetVal {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ReturnValue != nil {
			_ptr_pRetVal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReturnValue != nil {
					if err := o.ReturnValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReturnValue, _ptr_pRetVal); err != nil {
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

func (o *xxx_RemoteDispatchNotAutoDoneOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pRetVal {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pRetVal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReturnValue == nil {
				o.ReturnValue = &oaut.String{}
			}
			if err := o.ReturnValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pRetVal := func(ptr interface{}) { o.ReturnValue = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ReturnValue, _s_pRetVal, _ptr_pRetVal); err != nil {
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

// RemoteDispatchNotAutoDoneRequest structure represents the RemoteDispatchNotAutoDone operation request
type RemoteDispatchNotAutoDoneRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// s: The s parameter contains binary data representing the input parameters of the
	// method called on the server. The binary data MUST be marshaled as specified in [MS-NRTP]
	// section 3.1.5.1.1. The data is specified as is in the BSTR, such that the length
	// of the BSTR is the size of the data divided by 2 (rounded up if necessary).
	String *oaut.String `idl:"name:s" json:"string"`
}

func (o *RemoteDispatchNotAutoDoneRequest) xxx_ToOp(ctx context.Context, op *xxx_RemoteDispatchNotAutoDoneOperation) *xxx_RemoteDispatchNotAutoDoneOperation {
	if op == nil {
		op = &xxx_RemoteDispatchNotAutoDoneOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.String = o.String
	return op
}

func (o *RemoteDispatchNotAutoDoneRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoteDispatchNotAutoDoneOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.String = op.String
}
func (o *RemoteDispatchNotAutoDoneRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RemoteDispatchNotAutoDoneRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteDispatchNotAutoDoneOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoteDispatchNotAutoDoneResponse structure represents the RemoteDispatchNotAutoDone operation response
type RemoteDispatchNotAutoDoneResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pRetVal: The pRetVal parameter contains the binary data representing the output parameters
	// of the method called on the server. The binary data MUST be marshaled as specified
	// in [MS-NRTP] section 3.1.5.1.1. The data is specified as is in the BSTR, such that
	// the length of the BSTR is the size of the data divided by 2 (rounded up if necessary).
	ReturnValue *oaut.String `idl:"name:pRetVal" json:"return_value"`
	// Return: The RemoteDispatchNotAutoDone return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemoteDispatchNotAutoDoneResponse) xxx_ToOp(ctx context.Context, op *xxx_RemoteDispatchNotAutoDoneOperation) *xxx_RemoteDispatchNotAutoDoneOperation {
	if op == nil {
		op = &xxx_RemoteDispatchNotAutoDoneOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *RemoteDispatchNotAutoDoneResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoteDispatchNotAutoDoneOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *RemoteDispatchNotAutoDoneResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RemoteDispatchNotAutoDoneResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteDispatchNotAutoDoneOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
