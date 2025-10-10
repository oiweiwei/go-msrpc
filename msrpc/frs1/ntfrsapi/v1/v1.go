package ntfrsapi

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
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
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "frs1"
)

var (
	// Syntax UUID
	NTFrsAPISyntaxUUID = &uuid.UUID{TimeLow: 0xd049b186, TimeMid: 0x814f, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x9a, ClockSeqLow: 0x3c, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xc9, 0xb2, 0x32}}
	// Syntax ID
	NTFrsAPISyntaxV1_1 = &dcerpc.SyntaxID{IfUUID: NTFrsAPISyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 1}
)

// NtFrsApi interface.
type NTFrsAPIClient interface {

	// Opnum0NotUsedOnWire operation.
	// Opnum0NotUsedOnWire

	// Opnum1NotUsedOnWire operation.
	// Opnum1NotUsedOnWire

	// Opnum2NotUsedOnWire operation.
	// Opnum2NotUsedOnWire

	// Opnum3NotUsedOnWire operation.
	// Opnum3NotUsedOnWire

	SetDSPollingIntervalW(context.Context, *SetDSPollingIntervalWRequest, ...dcerpc.CallOption) (*SetDSPollingIntervalWResponse, error)

	GetDSPollingIntervalW(context.Context, *GetDSPollingIntervalWRequest, ...dcerpc.CallOption) (*GetDSPollingIntervalWResponse, error)

	// Opnum6NotUsedOnWire operation.
	// Opnum6NotUsedOnWire

	InfoW(context.Context, *InfoWRequest, ...dcerpc.CallOption) (*InfoWResponse, error)

	IsPathReplicated(context.Context, *IsPathReplicatedRequest, ...dcerpc.CallOption) (*IsPathReplicatedResponse, error)

	WriterCommand(context.Context, *WriterCommandRequest, ...dcerpc.CallOption) (*WriterCommandResponse, error)

	ForceReplication(context.Context, *ForceReplicationRequest, ...dcerpc.CallOption) (*ForceReplicationResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultNTFrsAPIClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultNTFrsAPIClient) SetDSPollingIntervalW(ctx context.Context, in *SetDSPollingIntervalWRequest, opts ...dcerpc.CallOption) (*SetDSPollingIntervalWResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetDSPollingIntervalWResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNTFrsAPIClient) GetDSPollingIntervalW(ctx context.Context, in *GetDSPollingIntervalWRequest, opts ...dcerpc.CallOption) (*GetDSPollingIntervalWResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetDSPollingIntervalWResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNTFrsAPIClient) InfoW(ctx context.Context, in *InfoWRequest, opts ...dcerpc.CallOption) (*InfoWResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &InfoWResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNTFrsAPIClient) IsPathReplicated(ctx context.Context, in *IsPathReplicatedRequest, opts ...dcerpc.CallOption) (*IsPathReplicatedResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &IsPathReplicatedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNTFrsAPIClient) WriterCommand(ctx context.Context, in *WriterCommandRequest, opts ...dcerpc.CallOption) (*WriterCommandResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WriterCommandResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNTFrsAPIClient) ForceReplication(ctx context.Context, in *ForceReplicationRequest, opts ...dcerpc.CallOption) (*ForceReplicationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ForceReplicationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNTFrsAPIClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultNTFrsAPIClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewNTFrsAPIClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (NTFrsAPIClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(NTFrsAPISyntaxV1_1))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultNTFrsAPIClient{cc: cc}, nil
}

// xxx_SetDSPollingIntervalWOperation structure represents the NtFrsApi_Rpc_Set_DsPollingIntervalW operation
type xxx_SetDSPollingIntervalWOperation struct {
	UseShortInterval uint32 `idl:"name:UseShortInterval" json:"use_short_interval"`
	LongInterval     uint32 `idl:"name:LongInterval" json:"long_interval"`
	ShortInterval    uint32 `idl:"name:ShortInterval" json:"short_interval"`
	Return           uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_SetDSPollingIntervalWOperation) OpNum() int { return 4 }

func (o *xxx_SetDSPollingIntervalWOperation) OpName() string {
	return "/NtFrsApi/v1.1/NtFrsApi_Rpc_Set_DsPollingIntervalW"
}

func (o *xxx_SetDSPollingIntervalWOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDSPollingIntervalWOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// UseShortInterval {in} (1:(uint32))
	{
		if err := w.WriteData(o.UseShortInterval); err != nil {
			return err
		}
	}
	// LongInterval {in} (1:(uint32))
	{
		if err := w.WriteData(o.LongInterval); err != nil {
			return err
		}
	}
	// ShortInterval {in} (1:(uint32))
	{
		if err := w.WriteData(o.ShortInterval); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDSPollingIntervalWOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// UseShortInterval {in} (1:(uint32))
	{
		if err := w.ReadData(&o.UseShortInterval); err != nil {
			return err
		}
	}
	// LongInterval {in} (1:(uint32))
	{
		if err := w.ReadData(&o.LongInterval); err != nil {
			return err
		}
	}
	// ShortInterval {in} (1:(uint32))
	{
		if err := w.ReadData(&o.ShortInterval); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDSPollingIntervalWOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDSPollingIntervalWOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDSPollingIntervalWOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetDSPollingIntervalWRequest structure represents the NtFrsApi_Rpc_Set_DsPollingIntervalW operation request
type SetDSPollingIntervalWRequest struct {
	UseShortInterval uint32 `idl:"name:UseShortInterval" json:"use_short_interval"`
	LongInterval     uint32 `idl:"name:LongInterval" json:"long_interval"`
	ShortInterval    uint32 `idl:"name:ShortInterval" json:"short_interval"`
}

func (o *SetDSPollingIntervalWRequest) xxx_ToOp(ctx context.Context, op *xxx_SetDSPollingIntervalWOperation) *xxx_SetDSPollingIntervalWOperation {
	if op == nil {
		op = &xxx_SetDSPollingIntervalWOperation{}
	}
	if o == nil {
		return op
	}
	op.UseShortInterval = o.UseShortInterval
	op.LongInterval = o.LongInterval
	op.ShortInterval = o.ShortInterval
	return op
}

func (o *SetDSPollingIntervalWRequest) xxx_FromOp(ctx context.Context, op *xxx_SetDSPollingIntervalWOperation) {
	if o == nil {
		return
	}
	o.UseShortInterval = op.UseShortInterval
	o.LongInterval = op.LongInterval
	o.ShortInterval = op.ShortInterval
}
func (o *SetDSPollingIntervalWRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetDSPollingIntervalWRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDSPollingIntervalWOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetDSPollingIntervalWResponse structure represents the NtFrsApi_Rpc_Set_DsPollingIntervalW operation response
type SetDSPollingIntervalWResponse struct {
	// Return: The NtFrsApi_Rpc_Set_DsPollingIntervalW return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetDSPollingIntervalWResponse) xxx_ToOp(ctx context.Context, op *xxx_SetDSPollingIntervalWOperation) *xxx_SetDSPollingIntervalWOperation {
	if op == nil {
		op = &xxx_SetDSPollingIntervalWOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *SetDSPollingIntervalWResponse) xxx_FromOp(ctx context.Context, op *xxx_SetDSPollingIntervalWOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetDSPollingIntervalWResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetDSPollingIntervalWResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDSPollingIntervalWOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDSPollingIntervalWOperation structure represents the NtFrsApi_Rpc_Get_DsPollingIntervalW operation
type xxx_GetDSPollingIntervalWOperation struct {
	Interval      uint32 `idl:"name:Interval" json:"interval"`
	LongInterval  uint32 `idl:"name:LongInterval" json:"long_interval"`
	ShortInterval uint32 `idl:"name:ShortInterval" json:"short_interval"`
	Return        uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDSPollingIntervalWOperation) OpNum() int { return 5 }

func (o *xxx_GetDSPollingIntervalWOperation) OpName() string {
	return "/NtFrsApi/v1.1/NtFrsApi_Rpc_Get_DsPollingIntervalW"
}

func (o *xxx_GetDSPollingIntervalWOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDSPollingIntervalWOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_GetDSPollingIntervalWOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_GetDSPollingIntervalWOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDSPollingIntervalWOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Interval {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Interval); err != nil {
			return err
		}
	}
	// LongInterval {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.LongInterval); err != nil {
			return err
		}
	}
	// ShortInterval {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.ShortInterval); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDSPollingIntervalWOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Interval {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Interval); err != nil {
			return err
		}
	}
	// LongInterval {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.LongInterval); err != nil {
			return err
		}
	}
	// ShortInterval {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.ShortInterval); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetDSPollingIntervalWRequest structure represents the NtFrsApi_Rpc_Get_DsPollingIntervalW operation request
type GetDSPollingIntervalWRequest struct {
}

func (o *GetDSPollingIntervalWRequest) xxx_ToOp(ctx context.Context, op *xxx_GetDSPollingIntervalWOperation) *xxx_GetDSPollingIntervalWOperation {
	if op == nil {
		op = &xxx_GetDSPollingIntervalWOperation{}
	}
	if o == nil {
		return op
	}
	return op
}

func (o *GetDSPollingIntervalWRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDSPollingIntervalWOperation) {
	if o == nil {
		return
	}
}
func (o *GetDSPollingIntervalWRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetDSPollingIntervalWRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDSPollingIntervalWOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDSPollingIntervalWResponse structure represents the NtFrsApi_Rpc_Get_DsPollingIntervalW operation response
type GetDSPollingIntervalWResponse struct {
	Interval      uint32 `idl:"name:Interval" json:"interval"`
	LongInterval  uint32 `idl:"name:LongInterval" json:"long_interval"`
	ShortInterval uint32 `idl:"name:ShortInterval" json:"short_interval"`
	// Return: The NtFrsApi_Rpc_Get_DsPollingIntervalW return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetDSPollingIntervalWResponse) xxx_ToOp(ctx context.Context, op *xxx_GetDSPollingIntervalWOperation) *xxx_GetDSPollingIntervalWOperation {
	if op == nil {
		op = &xxx_GetDSPollingIntervalWOperation{}
	}
	if o == nil {
		return op
	}
	op.Interval = o.Interval
	op.LongInterval = o.LongInterval
	op.ShortInterval = o.ShortInterval
	op.Return = o.Return
	return op
}

func (o *GetDSPollingIntervalWResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDSPollingIntervalWOperation) {
	if o == nil {
		return
	}
	o.Interval = op.Interval
	o.LongInterval = op.LongInterval
	o.ShortInterval = op.ShortInterval
	o.Return = op.Return
}
func (o *GetDSPollingIntervalWResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetDSPollingIntervalWResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDSPollingIntervalWOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_InfoWOperation structure represents the NtFrsApi_Rpc_InfoW operation
type xxx_InfoWOperation struct {
	BlobSize uint32 `idl:"name:BlobSize" json:"blob_size"`
	Blob     []byte `idl:"name:Blob;size_is:(BlobSize);pointer:unique" json:"blob"`
	Return   uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_InfoWOperation) OpNum() int { return 7 }

func (o *xxx_InfoWOperation) OpName() string { return "/NtFrsApi/v1.1/NtFrsApi_Rpc_InfoW" }

func (o *xxx_InfoWOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Blob != nil && o.BlobSize == 0 {
		o.BlobSize = uint32(len(o.Blob))
	}
	if o.BlobSize > uint32(65536) {
		return fmt.Errorf("BlobSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InfoWOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// BlobSize {in} (1:{range=(0,65536)}(uint32))
	{
		if err := w.WriteData(o.BlobSize); err != nil {
			return err
		}
	}
	// Blob {in, out} (1:{pointer=unique}*(1)[dim:0,size_is=BlobSize](uchar))
	{
		if o.Blob != nil || o.BlobSize > 0 {
			_ptr_Blob := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.BlobSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Blob {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.Blob[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.Blob); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Blob, _ptr_Blob); err != nil {
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

func (o *xxx_InfoWOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// BlobSize {in} (1:{range=(0,65536)}(uint32))
	{
		if err := w.ReadData(&o.BlobSize); err != nil {
			return err
		}
	}
	// Blob {in, out} (1:{pointer=unique}*(1)[dim:0,size_is=BlobSize](uchar))
	{
		_ptr_Blob := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Blob", sizeInfo[0])
			}
			o.Blob = make([]byte, sizeInfo[0])
			for i1 := range o.Blob {
				i1 := i1
				if err := w.ReadData(&o.Blob[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_Blob := func(ptr interface{}) { o.Blob = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Blob, _s_Blob, _ptr_Blob); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InfoWOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InfoWOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Blob {in, out} (1:{pointer=unique}*(1)[dim:0,size_is=BlobSize](uchar))
	{
		if o.Blob != nil || o.BlobSize > 0 {
			_ptr_Blob := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.BlobSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Blob {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.Blob[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.Blob); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Blob, _ptr_Blob); err != nil {
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
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InfoWOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Blob {in, out} (1:{pointer=unique}*(1)[dim:0,size_is=BlobSize](uchar))
	{
		_ptr_Blob := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Blob", sizeInfo[0])
			}
			o.Blob = make([]byte, sizeInfo[0])
			for i1 := range o.Blob {
				i1 := i1
				if err := w.ReadData(&o.Blob[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_Blob := func(ptr interface{}) { o.Blob = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Blob, _s_Blob, _ptr_Blob); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// InfoWRequest structure represents the NtFrsApi_Rpc_InfoW operation request
type InfoWRequest struct {
	BlobSize uint32 `idl:"name:BlobSize" json:"blob_size"`
	Blob     []byte `idl:"name:Blob;size_is:(BlobSize);pointer:unique" json:"blob"`
}

func (o *InfoWRequest) xxx_ToOp(ctx context.Context, op *xxx_InfoWOperation) *xxx_InfoWOperation {
	if op == nil {
		op = &xxx_InfoWOperation{}
	}
	if o == nil {
		return op
	}
	op.BlobSize = o.BlobSize
	op.Blob = o.Blob
	return op
}

func (o *InfoWRequest) xxx_FromOp(ctx context.Context, op *xxx_InfoWOperation) {
	if o == nil {
		return
	}
	o.BlobSize = op.BlobSize
	o.Blob = op.Blob
}
func (o *InfoWRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *InfoWRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InfoWOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// InfoWResponse structure represents the NtFrsApi_Rpc_InfoW operation response
type InfoWResponse struct {
	// XXX: BlobSize is an implicit input depedency for output parameters
	BlobSize uint32 `idl:"name:BlobSize" json:"blob_size"`

	Blob []byte `idl:"name:Blob;size_is:(BlobSize);pointer:unique" json:"blob"`
	// Return: The NtFrsApi_Rpc_InfoW return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *InfoWResponse) xxx_ToOp(ctx context.Context, op *xxx_InfoWOperation) *xxx_InfoWOperation {
	if op == nil {
		op = &xxx_InfoWOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.BlobSize == uint32(0) {
		op.BlobSize = o.BlobSize
	}

	op.Blob = o.Blob
	op.Return = o.Return
	return op
}

func (o *InfoWResponse) xxx_FromOp(ctx context.Context, op *xxx_InfoWOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.BlobSize = op.BlobSize

	o.Blob = op.Blob
	o.Return = op.Return
}
func (o *InfoWResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *InfoWResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InfoWOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IsPathReplicatedOperation structure represents the NtFrsApi_Rpc_IsPathReplicated operation
type xxx_IsPathReplicatedOperation struct {
	Path                     string     `idl:"name:Path;string;pointer:unique" json:"path"`
	SetTypeOfInterestReplica uint32     `idl:"name:ReplicaSetTypeOfInterest" json:"set_type_of_interest_replica"`
	Replicated               uint32     `idl:"name:Replicated" json:"replicated"`
	Primary                  uint32     `idl:"name:Primary" json:"primary"`
	Root                     uint32     `idl:"name:Root" json:"root"`
	SetGUIDReplica           *dtyp.GUID `idl:"name:ReplicaSetGuid" json:"set_guid_replica"`
	Return                   uint32     `idl:"name:Return" json:"return"`
}

func (o *xxx_IsPathReplicatedOperation) OpNum() int { return 8 }

func (o *xxx_IsPathReplicatedOperation) OpName() string {
	return "/NtFrsApi/v1.1/NtFrsApi_Rpc_IsPathReplicated"
}

func (o *xxx_IsPathReplicatedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsPathReplicatedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Path {in} (1:{string, pointer=unique, alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_Path := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_Path); err != nil {
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
	// ReplicaSetTypeOfInterest {in} (1:(uint32))
	{
		if err := w.WriteData(o.SetTypeOfInterestReplica); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsPathReplicatedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Path {in} (1:{string, pointer=unique, alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_Path := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_Path := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_Path, _ptr_Path); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ReplicaSetTypeOfInterest {in} (1:(uint32))
	{
		if err := w.ReadData(&o.SetTypeOfInterestReplica); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsPathReplicatedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsPathReplicatedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Replicated {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Replicated); err != nil {
			return err
		}
	}
	// Primary {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Primary); err != nil {
			return err
		}
	}
	// Root {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Root); err != nil {
			return err
		}
	}
	// ReplicaSetGuid {out} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.SetGUIDReplica != nil {
			if err := o.SetGUIDReplica.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsPathReplicatedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Replicated {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Replicated); err != nil {
			return err
		}
	}
	// Primary {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Primary); err != nil {
			return err
		}
	}
	// Root {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Root); err != nil {
			return err
		}
	}
	// ReplicaSetGuid {out} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.SetGUIDReplica == nil {
			o.SetGUIDReplica = &dtyp.GUID{}
		}
		if err := o.SetGUIDReplica.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// IsPathReplicatedRequest structure represents the NtFrsApi_Rpc_IsPathReplicated operation request
type IsPathReplicatedRequest struct {
	Path                     string `idl:"name:Path;string;pointer:unique" json:"path"`
	SetTypeOfInterestReplica uint32 `idl:"name:ReplicaSetTypeOfInterest" json:"set_type_of_interest_replica"`
}

func (o *IsPathReplicatedRequest) xxx_ToOp(ctx context.Context, op *xxx_IsPathReplicatedOperation) *xxx_IsPathReplicatedOperation {
	if op == nil {
		op = &xxx_IsPathReplicatedOperation{}
	}
	if o == nil {
		return op
	}
	op.Path = o.Path
	op.SetTypeOfInterestReplica = o.SetTypeOfInterestReplica
	return op
}

func (o *IsPathReplicatedRequest) xxx_FromOp(ctx context.Context, op *xxx_IsPathReplicatedOperation) {
	if o == nil {
		return
	}
	o.Path = op.Path
	o.SetTypeOfInterestReplica = op.SetTypeOfInterestReplica
}
func (o *IsPathReplicatedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *IsPathReplicatedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsPathReplicatedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsPathReplicatedResponse structure represents the NtFrsApi_Rpc_IsPathReplicated operation response
type IsPathReplicatedResponse struct {
	Replicated     uint32     `idl:"name:Replicated" json:"replicated"`
	Primary        uint32     `idl:"name:Primary" json:"primary"`
	Root           uint32     `idl:"name:Root" json:"root"`
	SetGUIDReplica *dtyp.GUID `idl:"name:ReplicaSetGuid" json:"set_guid_replica"`
	// Return: The NtFrsApi_Rpc_IsPathReplicated return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *IsPathReplicatedResponse) xxx_ToOp(ctx context.Context, op *xxx_IsPathReplicatedOperation) *xxx_IsPathReplicatedOperation {
	if op == nil {
		op = &xxx_IsPathReplicatedOperation{}
	}
	if o == nil {
		return op
	}
	op.Replicated = o.Replicated
	op.Primary = o.Primary
	op.Root = o.Root
	op.SetGUIDReplica = o.SetGUIDReplica
	op.Return = o.Return
	return op
}

func (o *IsPathReplicatedResponse) xxx_FromOp(ctx context.Context, op *xxx_IsPathReplicatedOperation) {
	if o == nil {
		return
	}
	o.Replicated = op.Replicated
	o.Primary = op.Primary
	o.Root = op.Root
	o.SetGUIDReplica = op.SetGUIDReplica
	o.Return = op.Return
}
func (o *IsPathReplicatedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *IsPathReplicatedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsPathReplicatedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WriterCommandOperation structure represents the NtFrsApi_Rpc_WriterCommand operation
type xxx_WriterCommandOperation struct {
	Command uint32 `idl:"name:Command" json:"command"`
	Return  uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_WriterCommandOperation) OpNum() int { return 9 }

func (o *xxx_WriterCommandOperation) OpName() string {
	return "/NtFrsApi/v1.1/NtFrsApi_Rpc_WriterCommand"
}

func (o *xxx_WriterCommandOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WriterCommandOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Command {in} (1:(uint32))
	{
		if err := w.WriteData(o.Command); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WriterCommandOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Command {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Command); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WriterCommandOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WriterCommandOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WriterCommandOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WriterCommandRequest structure represents the NtFrsApi_Rpc_WriterCommand operation request
type WriterCommandRequest struct {
	Command uint32 `idl:"name:Command" json:"command"`
}

func (o *WriterCommandRequest) xxx_ToOp(ctx context.Context, op *xxx_WriterCommandOperation) *xxx_WriterCommandOperation {
	if op == nil {
		op = &xxx_WriterCommandOperation{}
	}
	if o == nil {
		return op
	}
	op.Command = o.Command
	return op
}

func (o *WriterCommandRequest) xxx_FromOp(ctx context.Context, op *xxx_WriterCommandOperation) {
	if o == nil {
		return
	}
	o.Command = op.Command
}
func (o *WriterCommandRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WriterCommandRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WriterCommandOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WriterCommandResponse structure represents the NtFrsApi_Rpc_WriterCommand operation response
type WriterCommandResponse struct {
	// Return: The NtFrsApi_Rpc_WriterCommand return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *WriterCommandResponse) xxx_ToOp(ctx context.Context, op *xxx_WriterCommandOperation) *xxx_WriterCommandOperation {
	if op == nil {
		op = &xxx_WriterCommandOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *WriterCommandResponse) xxx_FromOp(ctx context.Context, op *xxx_WriterCommandOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *WriterCommandResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WriterCommandResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WriterCommandOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ForceReplicationOperation structure represents the NtFrsApi_Rpc_ForceReplication operation
type xxx_ForceReplicationOperation struct {
	SetGUIDReplica *dtyp.GUID `idl:"name:ReplicaSetGuid;pointer:unique" json:"set_guid_replica"`
	CxtionGUID     *dtyp.GUID `idl:"name:CxtionGuid;pointer:unique" json:"cxtion_guid"`
	SetNameReplica string     `idl:"name:ReplicaSetName;string;pointer:unique" json:"set_name_replica"`
	PartnerDNSName string     `idl:"name:PartnerDnsName;string;pointer:unique" json:"partner_dns_name"`
	Return         uint32     `idl:"name:Return" json:"return"`
}

func (o *xxx_ForceReplicationOperation) OpNum() int { return 10 }

func (o *xxx_ForceReplicationOperation) OpName() string {
	return "/NtFrsApi/v1.1/NtFrsApi_Rpc_ForceReplication"
}

func (o *xxx_ForceReplicationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ForceReplicationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ReplicaSetGuid {in} (1:{pointer=unique}*(1))(2:{alias=GUID}(struct))
	{
		if o.SetGUIDReplica != nil {
			_ptr_ReplicaSetGuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SetGUIDReplica != nil {
					if err := o.SetGUIDReplica.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SetGUIDReplica, _ptr_ReplicaSetGuid); err != nil {
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
	// CxtionGuid {in} (1:{pointer=unique}*(1))(2:{alias=GUID}(struct))
	{
		if o.CxtionGUID != nil {
			_ptr_CxtionGuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CxtionGUID != nil {
					if err := o.CxtionGUID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CxtionGUID, _ptr_CxtionGuid); err != nil {
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
	// ReplicaSetName {in} (1:{string, pointer=unique, alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.SetNameReplica != "" {
			_ptr_ReplicaSetName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.SetNameReplica); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.SetNameReplica, _ptr_ReplicaSetName); err != nil {
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
	// PartnerDnsName {in} (1:{string, pointer=unique, alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.PartnerDNSName != "" {
			_ptr_PartnerDnsName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.PartnerDNSName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.PartnerDNSName, _ptr_PartnerDnsName); err != nil {
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

func (o *xxx_ForceReplicationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ReplicaSetGuid {in} (1:{pointer=unique}*(1))(2:{alias=GUID}(struct))
	{
		_ptr_ReplicaSetGuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SetGUIDReplica == nil {
				o.SetGUIDReplica = &dtyp.GUID{}
			}
			if err := o.SetGUIDReplica.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReplicaSetGuid := func(ptr interface{}) { o.SetGUIDReplica = *ptr.(**dtyp.GUID) }
		if err := w.ReadPointer(&o.SetGUIDReplica, _s_ReplicaSetGuid, _ptr_ReplicaSetGuid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// CxtionGuid {in} (1:{pointer=unique}*(1))(2:{alias=GUID}(struct))
	{
		_ptr_CxtionGuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CxtionGUID == nil {
				o.CxtionGUID = &dtyp.GUID{}
			}
			if err := o.CxtionGUID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_CxtionGuid := func(ptr interface{}) { o.CxtionGUID = *ptr.(**dtyp.GUID) }
		if err := w.ReadPointer(&o.CxtionGUID, _s_CxtionGuid, _ptr_CxtionGuid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ReplicaSetName {in} (1:{string, pointer=unique, alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_ReplicaSetName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.SetNameReplica); err != nil {
				return err
			}
			return nil
		})
		_s_ReplicaSetName := func(ptr interface{}) { o.SetNameReplica = *ptr.(*string) }
		if err := w.ReadPointer(&o.SetNameReplica, _s_ReplicaSetName, _ptr_ReplicaSetName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// PartnerDnsName {in} (1:{string, pointer=unique, alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_PartnerDnsName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.PartnerDNSName); err != nil {
				return err
			}
			return nil
		})
		_s_PartnerDnsName := func(ptr interface{}) { o.PartnerDNSName = *ptr.(*string) }
		if err := w.ReadPointer(&o.PartnerDNSName, _s_PartnerDnsName, _ptr_PartnerDnsName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ForceReplicationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ForceReplicationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ForceReplicationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ForceReplicationRequest structure represents the NtFrsApi_Rpc_ForceReplication operation request
type ForceReplicationRequest struct {
	SetGUIDReplica *dtyp.GUID `idl:"name:ReplicaSetGuid;pointer:unique" json:"set_guid_replica"`
	CxtionGUID     *dtyp.GUID `idl:"name:CxtionGuid;pointer:unique" json:"cxtion_guid"`
	SetNameReplica string     `idl:"name:ReplicaSetName;string;pointer:unique" json:"set_name_replica"`
	PartnerDNSName string     `idl:"name:PartnerDnsName;string;pointer:unique" json:"partner_dns_name"`
}

func (o *ForceReplicationRequest) xxx_ToOp(ctx context.Context, op *xxx_ForceReplicationOperation) *xxx_ForceReplicationOperation {
	if op == nil {
		op = &xxx_ForceReplicationOperation{}
	}
	if o == nil {
		return op
	}
	op.SetGUIDReplica = o.SetGUIDReplica
	op.CxtionGUID = o.CxtionGUID
	op.SetNameReplica = o.SetNameReplica
	op.PartnerDNSName = o.PartnerDNSName
	return op
}

func (o *ForceReplicationRequest) xxx_FromOp(ctx context.Context, op *xxx_ForceReplicationOperation) {
	if o == nil {
		return
	}
	o.SetGUIDReplica = op.SetGUIDReplica
	o.CxtionGUID = op.CxtionGUID
	o.SetNameReplica = op.SetNameReplica
	o.PartnerDNSName = op.PartnerDNSName
}
func (o *ForceReplicationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ForceReplicationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ForceReplicationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ForceReplicationResponse structure represents the NtFrsApi_Rpc_ForceReplication operation response
type ForceReplicationResponse struct {
	// Return: The NtFrsApi_Rpc_ForceReplication return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ForceReplicationResponse) xxx_ToOp(ctx context.Context, op *xxx_ForceReplicationOperation) *xxx_ForceReplicationOperation {
	if op == nil {
		op = &xxx_ForceReplicationOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *ForceReplicationResponse) xxx_FromOp(ctx context.Context, op *xxx_ForceReplicationOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ForceReplicationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ForceReplicationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ForceReplicationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
