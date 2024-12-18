package ivdsdisk2

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
	GoPackage = "dcom/vds"
)

var (
	// IVdsDisk2 interface identifier 40f73c8b-687d-4a13-8d96-3d7f2e683936
	Disk2IID = &dcom.IID{Data1: 0x40f73c8b, Data2: 0x687d, Data3: 0x4a13, Data4: []byte{0x8d, 0x96, 0x3d, 0x7f, 0x2e, 0x68, 0x39, 0x36}}
	// Syntax UUID
	Disk2SyntaxUUID = &uuid.UUID{TimeLow: 0x40f73c8b, TimeMid: 0x687d, TimeHiAndVersion: 0x4a13, ClockSeqHiAndReserved: 0x8d, ClockSeqLow: 0x96, Node: [6]uint8{0x3d, 0x7f, 0x2e, 0x68, 0x39, 0x36}}
	// Syntax ID
	Disk2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: Disk2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsDisk2 interface.
type Disk2Client interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The SetSANMode method sets the SAN mode of a disk to either offline or online. A
	// disk that is offline exposes no volume devices for partitions or volumes with extents
	// on that disk. A disk can be REAONLY or READWRITE independent of the offline or online
	// setting.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	//
	// ERROR_SUCCESS (0x00000000)
	SetSANMode(context.Context, *SetSANModeRequest, ...dcerpc.CallOption) (*SetSANModeResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) Disk2Client
}

type xxx_DefaultDisk2Client struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultDisk2Client) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultDisk2Client) SetSANMode(ctx context.Context, in *SetSANModeRequest, opts ...dcerpc.CallOption) (*SetSANModeResponse, error) {
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
	out := &SetSANModeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDisk2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultDisk2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultDisk2Client) IPID(ctx context.Context, ipid *dcom.IPID) Disk2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultDisk2Client{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewDisk2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Disk2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(Disk2SyntaxV0_0))...)
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
	return &xxx_DefaultDisk2Client{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_SetSANModeOperation structure represents the SetSANMode operation
type xxx_SetSANModeOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Enable int32          `idl:"name:bEnable" json:"enable"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSANModeOperation) OpNum() int { return 3 }

func (o *xxx_SetSANModeOperation) OpName() string { return "/IVdsDisk2/v0/SetSANMode" }

func (o *xxx_SetSANModeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSANModeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bEnable {in} (1:(int32))
	{
		if err := w.WriteData(o.Enable); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSANModeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bEnable {in} (1:(int32))
	{
		if err := w.ReadData(&o.Enable); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSANModeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSANModeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSANModeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetSANModeRequest structure represents the SetSANMode operation request
type SetSANModeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bEnable: A Boolean value that indicates whether to set the disk to either online
	// or offline.
	Enable int32 `idl:"name:bEnable" json:"enable"`
}

func (o *SetSANModeRequest) xxx_ToOp(ctx context.Context) *xxx_SetSANModeOperation {
	if o == nil {
		return &xxx_SetSANModeOperation{}
	}
	return &xxx_SetSANModeOperation{
		This:   o.This,
		Enable: o.Enable,
	}
}

func (o *SetSANModeRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSANModeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Enable = op.Enable
}
func (o *SetSANModeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetSANModeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSANModeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSANModeResponse structure represents the SetSANMode operation response
type SetSANModeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetSANMode return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSANModeResponse) xxx_ToOp(ctx context.Context) *xxx_SetSANModeOperation {
	if o == nil {
		return &xxx_SetSANModeOperation{}
	}
	return &xxx_SetSANModeOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetSANModeResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSANModeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSANModeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetSANModeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSANModeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
