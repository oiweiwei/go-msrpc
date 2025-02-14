package iwbemloginclientid

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
)

var (
	// import guard
	GoPackage = "dcom/wmi"
)

var (
	// IWbemLoginClientID interface identifier d4781cd6-e5d3-44df-ad94-930efe48a887
	LoginClientIDIID = &dcom.IID{Data1: 0xd4781cd6, Data2: 0xe5d3, Data3: 0x44df, Data4: []byte{0xad, 0x94, 0x93, 0x0e, 0xfe, 0x48, 0xa8, 0x87}}
	// Syntax UUID
	LoginClientIDSyntaxUUID = &uuid.UUID{TimeLow: 0xd4781cd6, TimeMid: 0xe5d3, TimeHiAndVersion: 0x44df, ClockSeqHiAndReserved: 0xad, ClockSeqLow: 0x94, Node: [6]uint8{0x93, 0xe, 0xfe, 0x48, 0xa8, 0x87}}
	// Syntax ID
	LoginClientIDSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: LoginClientIDSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IWbemLoginClientID interface.
type LoginClientIDClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The IWbemLoginClientID::SetClientInfo method MUST pass the client NETBIOS name and
	// a unique client-generated number to the server.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	//
	// In case of failure, the server MUST return an HRESULT whose S (severity) bit is set
	// as specified in [MS-ERREF] section 2.1. The actual HRESULT value is implementation
	// dependent.
	SetClientInfo(context.Context, *SetClientInfoRequest, ...dcerpc.CallOption) (*SetClientInfoResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) LoginClientIDClient
}

type xxx_DefaultLoginClientIDClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultLoginClientIDClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultLoginClientIDClient) SetClientInfo(ctx context.Context, in *SetClientInfoRequest, opts ...dcerpc.CallOption) (*SetClientInfoResponse, error) {
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
	out := &SetClientInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLoginClientIDClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultLoginClientIDClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultLoginClientIDClient) IPID(ctx context.Context, ipid *dcom.IPID) LoginClientIDClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultLoginClientIDClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewLoginClientIDClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (LoginClientIDClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(LoginClientIDSyntaxV0_0))...)
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
	return &xxx_DefaultLoginClientIDClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_SetClientInfoOperation structure represents the SetClientInfo operation
type xxx_SetClientInfoOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	ClientMachine string         `idl:"name:wszClientMachine;string;pointer:unique" json:"client_machine"`
	ClientProcID  int32          `idl:"name:lClientProcId" json:"client_proc_id"`
	_             int32          `idl:"name:lReserved"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetClientInfoOperation) OpNum() int { return 3 }

func (o *xxx_SetClientInfoOperation) OpName() string { return "/IWbemLoginClientID/v0/SetClientInfo" }

func (o *xxx_SetClientInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClientInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// wszClientMachine {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ClientMachine != "" {
			_ptr_wszClientMachine := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ClientMachine); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ClientMachine, _ptr_wszClientMachine); err != nil {
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
	// lClientProcId {in} (1:(int32))
	{
		if err := w.WriteData(o.ClientProcID); err != nil {
			return err
		}
	}
	// lReserved {in} (1:(int32))
	{
		// reserved lReserved
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClientInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// wszClientMachine {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_wszClientMachine := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ClientMachine); err != nil {
				return err
			}
			return nil
		})
		_s_wszClientMachine := func(ptr interface{}) { o.ClientMachine = *ptr.(*string) }
		if err := w.ReadPointer(&o.ClientMachine, _s_wszClientMachine, _ptr_wszClientMachine); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lClientProcId {in} (1:(int32))
	{
		if err := w.ReadData(&o.ClientProcID); err != nil {
			return err
		}
	}
	// lReserved {in} (1:(int32))
	{
		// reserved lReserved
		var _lReserved int32
		if err := w.ReadData(&_lReserved); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClientInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClientInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClientInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetClientInfoRequest structure represents the SetClientInfo operation request
type SetClientInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// wszClientMachine: MUST specify the client NETBIOS name. This parameter MUST NOT be
	// NULL.
	ClientMachine string `idl:"name:wszClientMachine;string;pointer:unique" json:"client_machine"`
	// lClientProcId: Specifies a client-generated number. The server MAY use this for logging
	// purposes.<56>
	ClientProcID int32 `idl:"name:lClientProcId" json:"client_proc_id"`
}

func (o *SetClientInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_SetClientInfoOperation) *xxx_SetClientInfoOperation {
	if op == nil {
		op = &xxx_SetClientInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ClientMachine = o.ClientMachine
	op.ClientProcID = o.ClientProcID
	return op
}

func (o *SetClientInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_SetClientInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ClientMachine = op.ClientMachine
	o.ClientProcID = op.ClientProcID
}
func (o *SetClientInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetClientInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetClientInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetClientInfoResponse structure represents the SetClientInfo operation response
type SetClientInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetClientInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetClientInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_SetClientInfoOperation) *xxx_SetClientInfoOperation {
	if op == nil {
		op = &xxx_SetClientInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetClientInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_SetClientInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetClientInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetClientInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetClientInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
