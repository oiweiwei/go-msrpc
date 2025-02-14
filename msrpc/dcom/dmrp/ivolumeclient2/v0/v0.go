package ivolumeclient2

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
	GoPackage = "dcom/dmrp"
)

var (
	// IVolumeClient2 interface identifier 4bdafc52-fe6a-11d2-93f8-00105a11164a
	VolumeClient2IID = &dcom.IID{Data1: 0x4bdafc52, Data2: 0xfe6a, Data3: 0x11d2, Data4: []byte{0x93, 0xf8, 0x00, 0x10, 0x5a, 0x11, 0x16, 0x4a}}
	// Syntax UUID
	VolumeClient2SyntaxUUID = &uuid.UUID{TimeLow: 0x4bdafc52, TimeMid: 0xfe6a, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0x93, ClockSeqLow: 0xf8, Node: [6]uint8{0x0, 0x10, 0x5a, 0x11, 0x16, 0x4a}}
	// Syntax ID
	VolumeClient2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: VolumeClient2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVolumeClient2 interface.
type VolumeClient2Client interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// GetMaxAdjustedFreeSpace operation.
	GetMaxAdjustedFreeSpace(context.Context, *GetMaxAdjustedFreeSpaceRequest, ...dcerpc.CallOption) (*GetMaxAdjustedFreeSpaceResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) VolumeClient2Client
}

type xxx_DefaultVolumeClient2Client struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultVolumeClient2Client) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultVolumeClient2Client) GetMaxAdjustedFreeSpace(ctx context.Context, in *GetMaxAdjustedFreeSpaceRequest, opts ...dcerpc.CallOption) (*GetMaxAdjustedFreeSpaceResponse, error) {
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
	out := &GetMaxAdjustedFreeSpaceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClient2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultVolumeClient2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultVolumeClient2Client) IPID(ctx context.Context, ipid *dcom.IPID) VolumeClient2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultVolumeClient2Client{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewVolumeClient2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (VolumeClient2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(VolumeClient2SyntaxV0_0))...)
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
	return &xxx_DefaultVolumeClient2Client{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetMaxAdjustedFreeSpaceOperation structure represents the GetMaxAdjustedFreeSpace operation
type xxx_GetMaxAdjustedFreeSpaceOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	DiskID               int64          `idl:"name:diskId" json:"disk_id"`
	MaxAdjustedFreeSpace int64          `idl:"name:maxAdjustedFreeSpace" json:"max_adjusted_free_space"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMaxAdjustedFreeSpaceOperation) OpNum() int { return 3 }

func (o *xxx_GetMaxAdjustedFreeSpaceOperation) OpName() string {
	return "/IVolumeClient2/v0/GetMaxAdjustedFreeSpace"
}

func (o *xxx_GetMaxAdjustedFreeSpaceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMaxAdjustedFreeSpaceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// diskId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.DiskID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMaxAdjustedFreeSpaceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// diskId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.DiskID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMaxAdjustedFreeSpaceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMaxAdjustedFreeSpaceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// maxAdjustedFreeSpace {out} (1:{pointer=ref}*(1))(2:{alias=LONGLONG}(int64))
	{
		if err := w.WriteData(o.MaxAdjustedFreeSpace); err != nil {
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

func (o *xxx_GetMaxAdjustedFreeSpaceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// maxAdjustedFreeSpace {out} (1:{pointer=ref}*(1))(2:{alias=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.MaxAdjustedFreeSpace); err != nil {
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

// GetMaxAdjustedFreeSpaceRequest structure represents the GetMaxAdjustedFreeSpace operation request
type GetMaxAdjustedFreeSpaceRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	DiskID int64          `idl:"name:diskId" json:"disk_id"`
}

func (o *GetMaxAdjustedFreeSpaceRequest) xxx_ToOp(ctx context.Context, op *xxx_GetMaxAdjustedFreeSpaceOperation) *xxx_GetMaxAdjustedFreeSpaceOperation {
	if op == nil {
		op = &xxx_GetMaxAdjustedFreeSpaceOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.DiskID = op.DiskID
	return op
}

func (o *GetMaxAdjustedFreeSpaceRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMaxAdjustedFreeSpaceOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
}
func (o *GetMaxAdjustedFreeSpaceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetMaxAdjustedFreeSpaceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMaxAdjustedFreeSpaceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMaxAdjustedFreeSpaceResponse structure represents the GetMaxAdjustedFreeSpace operation response
type GetMaxAdjustedFreeSpaceResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	MaxAdjustedFreeSpace int64          `idl:"name:maxAdjustedFreeSpace" json:"max_adjusted_free_space"`
	// Return: The GetMaxAdjustedFreeSpace return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMaxAdjustedFreeSpaceResponse) xxx_ToOp(ctx context.Context, op *xxx_GetMaxAdjustedFreeSpaceOperation) *xxx_GetMaxAdjustedFreeSpaceOperation {
	if op == nil {
		op = &xxx_GetMaxAdjustedFreeSpaceOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.MaxAdjustedFreeSpace = op.MaxAdjustedFreeSpace
	o.Return = op.Return
	return op
}

func (o *GetMaxAdjustedFreeSpaceResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMaxAdjustedFreeSpaceOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MaxAdjustedFreeSpace = op.MaxAdjustedFreeSpace
	o.Return = op.Return
}
func (o *GetMaxAdjustedFreeSpaceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetMaxAdjustedFreeSpaceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMaxAdjustedFreeSpaceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
