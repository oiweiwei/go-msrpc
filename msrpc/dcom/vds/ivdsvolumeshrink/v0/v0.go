package ivdsvolumeshrink

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
	// IVdsVolumeShrink interface identifier d68168c9-82a2-4f85-b6e9-74707c49a58f
	VolumeShrinkIID = &dcom.IID{Data1: 0xd68168c9, Data2: 0x82a2, Data3: 0x4f85, Data4: []byte{0xb6, 0xe9, 0x74, 0x70, 0x7c, 0x49, 0xa5, 0x8f}}
	// Syntax UUID
	VolumeShrinkSyntaxUUID = &uuid.UUID{TimeLow: 0xd68168c9, TimeMid: 0x82a2, TimeHiAndVersion: 0x4f85, ClockSeqHiAndReserved: 0xb6, ClockSeqLow: 0xe9, Node: [6]uint8{0x74, 0x70, 0x7c, 0x49, 0xa5, 0x8f}}
	// Syntax ID
	VolumeShrinkSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: VolumeShrinkSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsVolumeShrink interface.
type VolumeShrinkClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The QueryMaxReclaimableBytes method retrieves the maximum number of bytes that can
	// be reclaimed from the current volume.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryMaxReclaimableBytes(context.Context, *QueryMaxReclaimableBytesRequest, ...dcerpc.CallOption) (*QueryMaxReclaimableBytesResponse, error)

	// Shrink operation.
	Shrink(context.Context, *ShrinkRequest, ...dcerpc.CallOption) (*ShrinkResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) VolumeShrinkClient
}

type xxx_DefaultVolumeShrinkClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultVolumeShrinkClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultVolumeShrinkClient) QueryMaxReclaimableBytes(ctx context.Context, in *QueryMaxReclaimableBytesRequest, opts ...dcerpc.CallOption) (*QueryMaxReclaimableBytesResponse, error) {
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
	out := &QueryMaxReclaimableBytesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeShrinkClient) Shrink(ctx context.Context, in *ShrinkRequest, opts ...dcerpc.CallOption) (*ShrinkResponse, error) {
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
	out := &ShrinkResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeShrinkClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultVolumeShrinkClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultVolumeShrinkClient) IPID(ctx context.Context, ipid *dcom.IPID) VolumeShrinkClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultVolumeShrinkClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewVolumeShrinkClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (VolumeShrinkClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(VolumeShrinkSyntaxV0_0))...)
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
	return &xxx_DefaultVolumeShrinkClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_QueryMaxReclaimableBytesOperation structure represents the QueryMaxReclaimableBytes operation
type xxx_QueryMaxReclaimableBytesOperation struct {
	This                            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                            *dcom.ORPCThat `idl:"name:That" json:"that"`
	PullMaxNumberOfReclaimableBytes uint64         `idl:"name:pullMaxNumberOfReclaimableBytes" json:"pull_max_number_of_reclaimable_bytes"`
	Return                          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryMaxReclaimableBytesOperation) OpNum() int { return 3 }

func (o *xxx_QueryMaxReclaimableBytesOperation) OpName() string {
	return "/IVdsVolumeShrink/v0/QueryMaxReclaimableBytes"
}

func (o *xxx_QueryMaxReclaimableBytesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryMaxReclaimableBytesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_QueryMaxReclaimableBytesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_QueryMaxReclaimableBytesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryMaxReclaimableBytesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pullMaxNumberOfReclaimableBytes {out} (1:{pointer=ref}*(1))(2:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.PullMaxNumberOfReclaimableBytes); err != nil {
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

func (o *xxx_QueryMaxReclaimableBytesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pullMaxNumberOfReclaimableBytes {out} (1:{pointer=ref}*(1))(2:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.PullMaxNumberOfReclaimableBytes); err != nil {
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

// QueryMaxReclaimableBytesRequest structure represents the QueryMaxReclaimableBytes operation request
type QueryMaxReclaimableBytesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *QueryMaxReclaimableBytesRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryMaxReclaimableBytesOperation) *xxx_QueryMaxReclaimableBytesOperation {
	if op == nil {
		op = &xxx_QueryMaxReclaimableBytesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *QueryMaxReclaimableBytesRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryMaxReclaimableBytesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *QueryMaxReclaimableBytesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryMaxReclaimableBytesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryMaxReclaimableBytesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryMaxReclaimableBytesResponse structure represents the QueryMaxReclaimableBytes operation response
type QueryMaxReclaimableBytesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pullMaxNumberOfReclaimableBytes: A pointer to a variable that, if the operation is
	// successfully completed, receives the maximum number of bytes that can be reclaimed
	// from the current volume. This number is always a multiple of the file system cluster
	// size, which is in turn a multiple of the disk sector size.
	PullMaxNumberOfReclaimableBytes uint64 `idl:"name:pullMaxNumberOfReclaimableBytes" json:"pull_max_number_of_reclaimable_bytes"`
	// Return: The QueryMaxReclaimableBytes return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryMaxReclaimableBytesResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryMaxReclaimableBytesOperation) *xxx_QueryMaxReclaimableBytesOperation {
	if op == nil {
		op = &xxx_QueryMaxReclaimableBytesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PullMaxNumberOfReclaimableBytes = o.PullMaxNumberOfReclaimableBytes
	op.Return = o.Return
	return op
}

func (o *QueryMaxReclaimableBytesResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryMaxReclaimableBytesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PullMaxNumberOfReclaimableBytes = op.PullMaxNumberOfReclaimableBytes
	o.Return = op.Return
}
func (o *QueryMaxReclaimableBytesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryMaxReclaimableBytesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryMaxReclaimableBytesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ShrinkOperation structure represents the Shrink operation
type xxx_ShrinkOperation struct {
	This                            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                            *dcom.ORPCThat `idl:"name:That" json:"that"`
	DesiredNumberOfReclaimableBytes uint64         `idl:"name:ullDesiredNumberOfReclaimableBytes" json:"desired_number_of_reclaimable_bytes"`
	MinNumberOfReclaimableBytes     uint64         `idl:"name:ullMinNumberOfReclaimableBytes" json:"min_number_of_reclaimable_bytes"`
	Async                           *vds.Async     `idl:"name:ppAsync" json:"async"`
	Return                          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ShrinkOperation) OpNum() int { return 4 }

func (o *xxx_ShrinkOperation) OpName() string { return "/IVdsVolumeShrink/v0/Shrink" }

func (o *xxx_ShrinkOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ShrinkOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ullDesiredNumberOfReclaimableBytes {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.DesiredNumberOfReclaimableBytes); err != nil {
			return err
		}
	}
	// ullMinNumberOfReclaimableBytes {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.MinNumberOfReclaimableBytes); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ShrinkOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ullDesiredNumberOfReclaimableBytes {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.DesiredNumberOfReclaimableBytes); err != nil {
			return err
		}
	}
	// ullMinNumberOfReclaimableBytes {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.MinNumberOfReclaimableBytes); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ShrinkOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ShrinkOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppAsync {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsAsync}(interface))
	{
		if o.Async != nil {
			_ptr_ppAsync := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Async != nil {
					if err := o.Async.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&vds.Async{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Async, _ptr_ppAsync); err != nil {
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

func (o *xxx_ShrinkOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppAsync {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsAsync}(interface))
	{
		_ptr_ppAsync := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Async == nil {
				o.Async = &vds.Async{}
			}
			if err := o.Async.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppAsync := func(ptr interface{}) { o.Async = *ptr.(**vds.Async) }
		if err := w.ReadPointer(&o.Async, _s_ppAsync, _ptr_ppAsync); err != nil {
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

// ShrinkRequest structure represents the Shrink operation request
type ShrinkRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                            *dcom.ORPCThis `idl:"name:This" json:"this"`
	DesiredNumberOfReclaimableBytes uint64         `idl:"name:ullDesiredNumberOfReclaimableBytes" json:"desired_number_of_reclaimable_bytes"`
	MinNumberOfReclaimableBytes     uint64         `idl:"name:ullMinNumberOfReclaimableBytes" json:"min_number_of_reclaimable_bytes"`
}

func (o *ShrinkRequest) xxx_ToOp(ctx context.Context, op *xxx_ShrinkOperation) *xxx_ShrinkOperation {
	if op == nil {
		op = &xxx_ShrinkOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.DesiredNumberOfReclaimableBytes = o.DesiredNumberOfReclaimableBytes
	op.MinNumberOfReclaimableBytes = o.MinNumberOfReclaimableBytes
	return op
}

func (o *ShrinkRequest) xxx_FromOp(ctx context.Context, op *xxx_ShrinkOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DesiredNumberOfReclaimableBytes = op.DesiredNumberOfReclaimableBytes
	o.MinNumberOfReclaimableBytes = op.MinNumberOfReclaimableBytes
}
func (o *ShrinkRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ShrinkRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ShrinkOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ShrinkResponse structure represents the Shrink operation response
type ShrinkResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat `idl:"name:That" json:"that"`
	Async *vds.Async     `idl:"name:ppAsync" json:"async"`
	// Return: The Shrink return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ShrinkResponse) xxx_ToOp(ctx context.Context, op *xxx_ShrinkOperation) *xxx_ShrinkOperation {
	if op == nil {
		op = &xxx_ShrinkOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Async = o.Async
	op.Return = o.Return
	return op
}

func (o *ShrinkResponse) xxx_FromOp(ctx context.Context, op *xxx_ShrinkOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
}
func (o *ShrinkResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ShrinkResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ShrinkOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
