package ivdsvolume2

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
	// IVdsVolume2 interface identifier 72ae6713-dcbb-4a03-b36b-371f6ac6b53d
	Volume2IID = &dcom.IID{Data1: 0x72ae6713, Data2: 0xdcbb, Data3: 0x4a03, Data4: []byte{0xb3, 0x6b, 0x37, 0x1f, 0x6a, 0xc6, 0xb5, 0x3d}}
	// Syntax UUID
	Volume2SyntaxUUID = &uuid.UUID{TimeLow: 0x72ae6713, TimeMid: 0xdcbb, TimeHiAndVersion: 0x4a03, ClockSeqHiAndReserved: 0xb3, ClockSeqLow: 0x6b, Node: [6]uint8{0x37, 0x1f, 0x6a, 0xc6, 0xb5, 0x3d}}
	// Syntax ID
	Volume2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: Volume2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsVolume2 interface.
type Volume2Client interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// GetProperties2 operation.
	GetProperties2(context.Context, *GetProperties2Request, ...dcerpc.CallOption) (*GetProperties2Response, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) Volume2Client
}

type xxx_DefaultVolume2Client struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultVolume2Client) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultVolume2Client) GetProperties2(ctx context.Context, in *GetProperties2Request, opts ...dcerpc.CallOption) (*GetProperties2Response, error) {
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
	out := &GetProperties2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolume2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultVolume2Client) IPID(ctx context.Context, ipid *dcom.IPID) Volume2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultVolume2Client{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}
func NewVolume2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Volume2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(Volume2SyntaxV0_0))...)
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
	return &xxx_DefaultVolume2Client{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetProperties2Operation structure represents the GetProperties2 operation
type xxx_GetProperties2Operation struct {
	This             *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat       `idl:"name:That" json:"that"`
	VolumeProperties *vds.VolumeProperty2 `idl:"name:pVolumeProperties" json:"volume_properties"`
	Return           int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_GetProperties2Operation) OpNum() int { return 3 }

func (o *xxx_GetProperties2Operation) OpName() string { return "/IVdsVolume2/v0/GetProperties2" }

func (o *xxx_GetProperties2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetProperties2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetProperties2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetProperties2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetProperties2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pVolumeProperties {out} (1:{pointer=ref}*(1))(2:{alias=VDS_VOLUME_PROP2}(struct))
	{
		if o.VolumeProperties != nil {
			if err := o.VolumeProperties.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&vds.VolumeProperty2{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetProperties2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pVolumeProperties {out} (1:{pointer=ref}*(1))(2:{alias=VDS_VOLUME_PROP2}(struct))
	{
		if o.VolumeProperties == nil {
			o.VolumeProperties = &vds.VolumeProperty2{}
		}
		if err := o.VolumeProperties.UnmarshalNDR(ctx, w); err != nil {
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

// GetProperties2Request structure represents the GetProperties2 operation request
type GetProperties2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetProperties2Request) xxx_ToOp(ctx context.Context) *xxx_GetProperties2Operation {
	if o == nil {
		return &xxx_GetProperties2Operation{}
	}
	return &xxx_GetProperties2Operation{
		This: o.This,
	}
}

func (o *GetProperties2Request) xxx_FromOp(ctx context.Context, op *xxx_GetProperties2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetProperties2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetProperties2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetProperties2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetProperties2Response structure represents the GetProperties2 operation response
type GetProperties2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That             *dcom.ORPCThat       `idl:"name:That" json:"that"`
	VolumeProperties *vds.VolumeProperty2 `idl:"name:pVolumeProperties" json:"volume_properties"`
	// Return: The GetProperties2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetProperties2Response) xxx_ToOp(ctx context.Context) *xxx_GetProperties2Operation {
	if o == nil {
		return &xxx_GetProperties2Operation{}
	}
	return &xxx_GetProperties2Operation{
		That:             o.That,
		VolumeProperties: o.VolumeProperties,
		Return:           o.Return,
	}
}

func (o *GetProperties2Response) xxx_FromOp(ctx context.Context, op *xxx_GetProperties2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.VolumeProperties = op.VolumeProperties
	o.Return = op.Return
}
func (o *GetProperties2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetProperties2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetProperties2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
