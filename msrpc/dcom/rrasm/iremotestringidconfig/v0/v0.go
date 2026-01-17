package iremotestringidconfig

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/rrasm"
)

var (
	// IRemoteStringIdConfig interface identifier 67e08fc2-2984-4b62-b92e-fc1aae64bbbb
	RemoteStringIDConfigIID = &dcom.IID{Data1: 0x67e08fc2, Data2: 0x2984, Data3: 0x4b62, Data4: []byte{0xb9, 0x2e, 0xfc, 0x1a, 0xae, 0x64, 0xbb, 0xbb}}
	// Syntax UUID
	RemoteStringIDConfigSyntaxUUID = &uuid.UUID{TimeLow: 0x67e08fc2, TimeMid: 0x2984, TimeHiAndVersion: 0x4b62, ClockSeqHiAndReserved: 0xb9, ClockSeqLow: 0x2e, Node: [6]uint8{0xfc, 0x1a, 0xae, 0x64, 0xbb, 0xbb}}
	// Syntax ID
	RemoteStringIDConfigSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: RemoteStringIDConfigSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IRemoteStringIdConfig interface.
type RemoteStringIDConfigClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// GetStringFromId operation.
	GetStringFromID(context.Context, *GetStringFromIDRequest, ...dcerpc.CallOption) (*GetStringFromIDResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) RemoteStringIDConfigClient
}

type xxx_DefaultRemoteStringIDConfigClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultRemoteStringIDConfigClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultRemoteStringIDConfigClient) GetStringFromID(ctx context.Context, in *GetStringFromIDRequest, opts ...dcerpc.CallOption) (*GetStringFromIDResponse, error) {
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
	out := &GetStringFromIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteStringIDConfigClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultRemoteStringIDConfigClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultRemoteStringIDConfigClient) IPID(ctx context.Context, ipid *dcom.IPID) RemoteStringIDConfigClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultRemoteStringIDConfigClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewRemoteStringIDConfigClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (RemoteStringIDConfigClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(RemoteStringIDConfigSyntaxV0_0))...)
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
	return &xxx_DefaultRemoteStringIDConfigClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetStringFromIDOperation structure represents the GetStringFromId operation
type xxx_GetStringFromIDOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	StringID uint32         `idl:"name:stringId" json:"string_id"`
	Name     *oaut.String   `idl:"name:pBstrName" json:"name"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetStringFromIDOperation) OpNum() int { return 3 }

func (o *xxx_GetStringFromIDOperation) OpName() string {
	return "/IRemoteStringIdConfig/v0/GetStringFromId"
}

func (o *xxx_GetStringFromIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetStringFromIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// stringId {in} (1:{alias=UINT}(uint32))
	{
		if err := w.WriteData(o.StringID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetStringFromIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// stringId {in} (1:{alias=UINT}(uint32))
	{
		if err := w.ReadData(&o.StringID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetStringFromIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetStringFromIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pBstrName {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Name != nil {
			_ptr_pBstrName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Name != nil {
					if err := o.Name.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Name, _ptr_pBstrName); err != nil {
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

func (o *xxx_GetStringFromIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pBstrName {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pBstrName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Name == nil {
				o.Name = &oaut.String{}
			}
			if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pBstrName := func(ptr interface{}) { o.Name = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Name, _s_pBstrName, _ptr_pBstrName); err != nil {
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

// GetStringFromIDRequest structure represents the GetStringFromId operation request
type GetStringFromIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	StringID uint32         `idl:"name:stringId" json:"string_id"`
}

func (o *GetStringFromIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetStringFromIDOperation) *xxx_GetStringFromIDOperation {
	if op == nil {
		op = &xxx_GetStringFromIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.StringID = o.StringID
	return op
}

func (o *GetStringFromIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetStringFromIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.StringID = op.StringID
}
func (o *GetStringFromIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetStringFromIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetStringFromIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetStringFromIDResponse structure represents the GetStringFromId operation response
type GetStringFromIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	Name *oaut.String   `idl:"name:pBstrName" json:"name"`
	// Return: The GetStringFromId return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetStringFromIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetStringFromIDOperation) *xxx_GetStringFromIDOperation {
	if op == nil {
		op = &xxx_GetStringFromIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Name = o.Name
	op.Return = o.Return
	return op
}

func (o *GetStringFromIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetStringFromIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Name = op.Name
	o.Return = op.Return
}
func (o *GetStringFromIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetStringFromIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetStringFromIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
