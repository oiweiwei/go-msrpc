package iapphostchangehandler

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
	GoPackage = "dcom/iisa"
)

var (
	// IAppHostChangeHandler interface identifier 09829352-87c2-418d-8d79-4133969a489d
	AppHostChangeHandlerIID = &dcom.IID{Data1: 0x09829352, Data2: 0x87c2, Data3: 0x418d, Data4: []byte{0x8d, 0x79, 0x41, 0x33, 0x96, 0x9a, 0x48, 0x9d}}
	// Syntax UUID
	AppHostChangeHandlerSyntaxUUID = &uuid.UUID{TimeLow: 0x9829352, TimeMid: 0x87c2, TimeHiAndVersion: 0x418d, ClockSeqHiAndReserved: 0x8d, ClockSeqLow: 0x79, Node: [6]uint8{0x41, 0x33, 0x96, 0x9a, 0x48, 0x9d}}
	// Syntax ID
	AppHostChangeHandlerSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: AppHostChangeHandlerSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IAppHostChangeHandler interface.
type AppHostChangeHandlerClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The OnSectionChanges method is called by the server by using an RPC_REQUEST packet.
	// This method is called when a change in the path of the administration system hierarchy
	// is detected. The callee (the client-implemented object) can react to this notification
	// as it determines. It can return any error and the server MUST ignore it.
	//
	// Return Values: The return value MUST be ignored by the server.
	OnSectionChanges(context.Context, *OnSectionChangesRequest, ...dcerpc.CallOption) (*OnSectionChangesResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) AppHostChangeHandlerClient
}

type xxx_DefaultAppHostChangeHandlerClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultAppHostChangeHandlerClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultAppHostChangeHandlerClient) OnSectionChanges(ctx context.Context, in *OnSectionChangesRequest, opts ...dcerpc.CallOption) (*OnSectionChangesResponse, error) {
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
	out := &OnSectionChangesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostChangeHandlerClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultAppHostChangeHandlerClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultAppHostChangeHandlerClient) IPID(ctx context.Context, ipid *dcom.IPID) AppHostChangeHandlerClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultAppHostChangeHandlerClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewAppHostChangeHandlerClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (AppHostChangeHandlerClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(AppHostChangeHandlerSyntaxV0_0))...)
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
	return &xxx_DefaultAppHostChangeHandlerClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_OnSectionChangesOperation structure represents the OnSectionChanges operation
type xxx_OnSectionChangesOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	SectionName *oaut.String   `idl:"name:bstrSectionName" json:"section_name"`
	ConfigPath  *oaut.String   `idl:"name:bstrConfigPath" json:"config_path"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_OnSectionChangesOperation) OpNum() int { return 3 }

func (o *xxx_OnSectionChangesOperation) OpName() string {
	return "/IAppHostChangeHandler/v0/OnSectionChanges"
}

func (o *xxx_OnSectionChangesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OnSectionChangesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrSectionName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SectionName != nil {
			_ptr_bstrSectionName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SectionName != nil {
					if err := o.SectionName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SectionName, _ptr_bstrSectionName); err != nil {
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
	// bstrConfigPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ConfigPath != nil {
			_ptr_bstrConfigPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConfigPath != nil {
					if err := o.ConfigPath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConfigPath, _ptr_bstrConfigPath); err != nil {
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

func (o *xxx_OnSectionChangesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrSectionName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrSectionName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SectionName == nil {
				o.SectionName = &oaut.String{}
			}
			if err := o.SectionName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrSectionName := func(ptr interface{}) { o.SectionName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SectionName, _s_bstrSectionName, _ptr_bstrSectionName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrConfigPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrConfigPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConfigPath == nil {
				o.ConfigPath = &oaut.String{}
			}
			if err := o.ConfigPath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrConfigPath := func(ptr interface{}) { o.ConfigPath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ConfigPath, _s_bstrConfigPath, _ptr_bstrConfigPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OnSectionChangesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OnSectionChangesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_OnSectionChangesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// OnSectionChangesRequest structure represents the OnSectionChanges operation request
type OnSectionChangesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrSectionName: The name of the IAppHostElement on the server that changed. A server
	// is free to not implement this parameter and always passes NULL.
	SectionName *oaut.String `idl:"name:bstrSectionName" json:"section_name"`
	// bstrConfigPath: The path in the hierarchy where the change was detected by the administration
	// system.
	ConfigPath *oaut.String `idl:"name:bstrConfigPath" json:"config_path"`
}

func (o *OnSectionChangesRequest) xxx_ToOp(ctx context.Context, op *xxx_OnSectionChangesOperation) *xxx_OnSectionChangesOperation {
	if op == nil {
		op = &xxx_OnSectionChangesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.SectionName = o.SectionName
	op.ConfigPath = o.ConfigPath
	return op
}

func (o *OnSectionChangesRequest) xxx_FromOp(ctx context.Context, op *xxx_OnSectionChangesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SectionName = op.SectionName
	o.ConfigPath = op.ConfigPath
}
func (o *OnSectionChangesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OnSectionChangesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OnSectionChangesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OnSectionChangesResponse structure represents the OnSectionChanges operation response
type OnSectionChangesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The OnSectionChanges return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OnSectionChangesResponse) xxx_ToOp(ctx context.Context, op *xxx_OnSectionChangesOperation) *xxx_OnSectionChangesOperation {
	if op == nil {
		op = &xxx_OnSectionChangesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *OnSectionChangesResponse) xxx_FromOp(ctx context.Context, op *xxx_OnSectionChangesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *OnSectionChangesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OnSectionChangesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OnSectionChangesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
