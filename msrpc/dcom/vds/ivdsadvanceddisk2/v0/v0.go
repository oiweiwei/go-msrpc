package ivdsadvanceddisk2

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
	// IVdsAdvancedDisk2 interface identifier 9723f420-9355-42de-ab66-e31bb15beeac
	AdvancedDisk2IID = &dcom.IID{Data1: 0x9723f420, Data2: 0x9355, Data3: 0x42de, Data4: []byte{0xab, 0x66, 0xe3, 0x1b, 0xb1, 0x5b, 0xee, 0xac}}
	// Syntax UUID
	AdvancedDisk2SyntaxUUID = &uuid.UUID{TimeLow: 0x9723f420, TimeMid: 0x9355, TimeHiAndVersion: 0x42de, ClockSeqHiAndReserved: 0xab, ClockSeqLow: 0x66, Node: [6]uint8{0xe3, 0x1b, 0xb1, 0x5b, 0xee, 0xac}}
	// Syntax ID
	AdvancedDisk2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: AdvancedDisk2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsAdvancedDisk2 interface.
type AdvancedDisk2Client interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The ChangePartitionType method changes the partition type on the disk at a specified
	// byte offset.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	//
	// ERROR_SUCCESS (0x00000000)
	ChangePartitionType(context.Context, *ChangePartitionTypeRequest, ...dcerpc.CallOption) (*ChangePartitionTypeResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) AdvancedDisk2Client
}

type xxx_DefaultAdvancedDisk2Client struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultAdvancedDisk2Client) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultAdvancedDisk2Client) ChangePartitionType(ctx context.Context, in *ChangePartitionTypeRequest, opts ...dcerpc.CallOption) (*ChangePartitionTypeResponse, error) {
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
	out := &ChangePartitionTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAdvancedDisk2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultAdvancedDisk2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultAdvancedDisk2Client) IPID(ctx context.Context, ipid *dcom.IPID) AdvancedDisk2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultAdvancedDisk2Client{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewAdvancedDisk2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (AdvancedDisk2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(AdvancedDisk2SyntaxV0_0))...)
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
	return &xxx_DefaultAdvancedDisk2Client{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_ChangePartitionTypeOperation structure represents the ChangePartitionType operation
type xxx_ChangePartitionTypeOperation struct {
	This       *dcom.ORPCThis                     `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat                     `idl:"name:That" json:"that"`
	Offset     uint64                             `idl:"name:ullOffset" json:"offset"`
	Force      int32                              `idl:"name:bForce" json:"force"`
	Parameters *vds.ChangePartitionTypeParameters `idl:"name:para" json:"parameters"`
	Return     int32                              `idl:"name:Return" json:"return"`
}

func (o *xxx_ChangePartitionTypeOperation) OpNum() int { return 3 }

func (o *xxx_ChangePartitionTypeOperation) OpName() string {
	return "/IVdsAdvancedDisk2/v0/ChangePartitionType"
}

func (o *xxx_ChangePartitionTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangePartitionTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ullOffset {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.Offset); err != nil {
			return err
		}
	}
	// bForce {in} (1:(int32))
	{
		if err := w.WriteData(o.Force); err != nil {
			return err
		}
	}
	// para {in} (1:{pointer=ref}*(1))(2:{alias=CHANGE_PARTITION_TYPE_PARAMETERS}(struct))
	{
		if o.Parameters != nil {
			if err := o.Parameters.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&vds.ChangePartitionTypeParameters{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ChangePartitionTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ullOffset {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.Offset); err != nil {
			return err
		}
	}
	// bForce {in} (1:(int32))
	{
		if err := w.ReadData(&o.Force); err != nil {
			return err
		}
	}
	// para {in} (1:{pointer=ref}*(1))(2:{alias=CHANGE_PARTITION_TYPE_PARAMETERS}(struct))
	{
		if o.Parameters == nil {
			o.Parameters = &vds.ChangePartitionTypeParameters{}
		}
		if err := o.Parameters.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangePartitionTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangePartitionTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ChangePartitionTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ChangePartitionTypeRequest structure represents the ChangePartitionType operation request
type ChangePartitionTypeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ullOffset: The byte offset of the partition, from the beginning of the disk. This
	// offset MUST be the offset of a start of a partition.
	Offset uint64 `idl:"name:ullOffset" json:"offset"`
	// bForce: A Boolean value that indicates whether a change will be forced even if the
	// volume cannot be locked for exclusive access. When bForce is false, ChangePartitionType
	// MUST lock and dismount the volume before changing the partition type. If bForce is
	// true, the volume MUST be dismounted and the change MUST be made even if the locking
	// of the volume fails.
	Force int32 `idl:"name:bForce" json:"force"`
	// para: A pointer to a CHANGE_PARTITION_TYPE_PARAMETERS structure that contains the
	// partition type that the partition at the location specified by ullOffset is changed
	// to.
	Parameters *vds.ChangePartitionTypeParameters `idl:"name:para" json:"parameters"`
}

func (o *ChangePartitionTypeRequest) xxx_ToOp(ctx context.Context, op *xxx_ChangePartitionTypeOperation) *xxx_ChangePartitionTypeOperation {
	if op == nil {
		op = &xxx_ChangePartitionTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Offset = o.Offset
	op.Force = o.Force
	op.Parameters = o.Parameters
	return op
}

func (o *ChangePartitionTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_ChangePartitionTypeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Offset = op.Offset
	o.Force = op.Force
	o.Parameters = op.Parameters
}
func (o *ChangePartitionTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ChangePartitionTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ChangePartitionTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ChangePartitionTypeResponse structure represents the ChangePartitionType operation response
type ChangePartitionTypeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ChangePartitionType return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ChangePartitionTypeResponse) xxx_ToOp(ctx context.Context, op *xxx_ChangePartitionTypeOperation) *xxx_ChangePartitionTypeOperation {
	if op == nil {
		op = &xxx_ChangePartitionTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ChangePartitionTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_ChangePartitionTypeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ChangePartitionTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ChangePartitionTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ChangePartitionTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
