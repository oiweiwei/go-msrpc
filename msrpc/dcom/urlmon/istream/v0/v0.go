package istream

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
	urlmon "github.com/oiweiwei/go-msrpc/msrpc/dcom/urlmon"
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
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
	_ = dtyp.GoPackage
	_ = urlmon.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/urlmon"
)

var (
	// IStream interface identifier 0000000c-0000-0000-c000-000000000046
	StreamIID = &dcom.IID{Data1: 0x0000000c, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
	// Syntax UUID
	StreamSyntaxUUID = &uuid.UUID{TimeLow: 0xc, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xc0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x46}}
	// Syntax ID
	StreamSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: StreamSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IStream interface.
type StreamClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// Read operation.
	Read(context.Context, *ReadRequest, ...dcerpc.CallOption) (*ReadResponse, error)

	// Write operation.
	Write(context.Context, *WriteRequest, ...dcerpc.CallOption) (*WriteResponse, error)

	// Seek operation.
	Seek(context.Context, *SeekRequest, ...dcerpc.CallOption) (*SeekResponse, error)

	// SetSize operation.
	SetSize(context.Context, *SetSizeRequest, ...dcerpc.CallOption) (*SetSizeResponse, error)

	// CopyTo operation.
	CopyTo(context.Context, *CopyToRequest, ...dcerpc.CallOption) (*CopyToResponse, error)

	// Commit operation.
	Commit(context.Context, *CommitRequest, ...dcerpc.CallOption) (*CommitResponse, error)

	// Revert operation.
	Revert(context.Context, *RevertRequest, ...dcerpc.CallOption) (*RevertResponse, error)

	// LockRegion operation.
	LockRegion(context.Context, *LockRegionRequest, ...dcerpc.CallOption) (*LockRegionResponse, error)

	// UnlockRegion operation.
	UnlockRegion(context.Context, *UnlockRegionRequest, ...dcerpc.CallOption) (*UnlockRegionResponse, error)

	// Stat operation.
	Stat(context.Context, *StatRequest, ...dcerpc.CallOption) (*StatResponse, error)

	// Clone operation.
	Clone(context.Context, *CloneRequest, ...dcerpc.CallOption) (*CloneResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) StreamClient
}

type xxx_DefaultStreamClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultStreamClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultStreamClient) Read(ctx context.Context, in *ReadRequest, opts ...dcerpc.CallOption) (*ReadResponse, error) {
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
	out := &ReadResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultStreamClient) Write(ctx context.Context, in *WriteRequest, opts ...dcerpc.CallOption) (*WriteResponse, error) {
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
	out := &WriteResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultStreamClient) Seek(ctx context.Context, in *SeekRequest, opts ...dcerpc.CallOption) (*SeekResponse, error) {
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
	out := &SeekResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultStreamClient) SetSize(ctx context.Context, in *SetSizeRequest, opts ...dcerpc.CallOption) (*SetSizeResponse, error) {
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
	out := &SetSizeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultStreamClient) CopyTo(ctx context.Context, in *CopyToRequest, opts ...dcerpc.CallOption) (*CopyToResponse, error) {
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
	out := &CopyToResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultStreamClient) Commit(ctx context.Context, in *CommitRequest, opts ...dcerpc.CallOption) (*CommitResponse, error) {
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
	out := &CommitResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultStreamClient) Revert(ctx context.Context, in *RevertRequest, opts ...dcerpc.CallOption) (*RevertResponse, error) {
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
	out := &RevertResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultStreamClient) LockRegion(ctx context.Context, in *LockRegionRequest, opts ...dcerpc.CallOption) (*LockRegionResponse, error) {
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
	out := &LockRegionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultStreamClient) UnlockRegion(ctx context.Context, in *UnlockRegionRequest, opts ...dcerpc.CallOption) (*UnlockRegionResponse, error) {
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
	out := &UnlockRegionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultStreamClient) Stat(ctx context.Context, in *StatRequest, opts ...dcerpc.CallOption) (*StatResponse, error) {
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
	out := &StatResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultStreamClient) Clone(ctx context.Context, in *CloneRequest, opts ...dcerpc.CallOption) (*CloneResponse, error) {
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
	out := &CloneResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultStreamClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultStreamClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultStreamClient) IPID(ctx context.Context, ipid *dcom.IPID) StreamClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultStreamClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewStreamClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (StreamClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(StreamSyntaxV0_0))...)
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
	return &xxx_DefaultStreamClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_ReadOperation structure represents the Read operation
type xxx_ReadOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	Pv         []byte         `idl:"name:pv;size_is:(cb)" json:"pv"`
	Length     uint32         `idl:"name:cb" json:"length"`
	ReadLength uint32         `idl:"name:pcbRead" json:"read_length"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ReadOperation) OpNum() int { return 3 }

func (o *xxx_ReadOperation) OpName() string { return "/IStream/v0/Read" }

func (o *xxx_ReadOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReadOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// cb {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Length); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReadOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// cb {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Length); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReadOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReadOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pv {out} (1:{pointer=ref}*(1))(2:{alias=BYTE}[dim:0,size_is=cb](uint8))
	{
		dimSize1 := uint64(o.Length)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Pv {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Pv[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Pv); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// pcbRead {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.ReadLength); err != nil {
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

func (o *xxx_ReadOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pv {out} (1:{pointer=ref}*(1))(2:{alias=BYTE}[dim:0,size_is=cb](uint8))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Pv", sizeInfo[0])
		}
		o.Pv = make([]byte, sizeInfo[0])
		for i1 := range o.Pv {
			i1 := i1
			if err := w.ReadData(&o.Pv[i1]); err != nil {
				return err
			}
		}
	}
	// pcbRead {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.ReadLength); err != nil {
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

// ReadRequest structure represents the Read operation request
type ReadRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	Length uint32         `idl:"name:cb" json:"length"`
}

func (o *ReadRequest) xxx_ToOp(ctx context.Context, op *xxx_ReadOperation) *xxx_ReadOperation {
	if op == nil {
		op = &xxx_ReadOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Length = o.Length
	return op
}

func (o *ReadRequest) xxx_FromOp(ctx context.Context, op *xxx_ReadOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Length = op.Length
}
func (o *ReadRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ReadRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReadOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReadResponse structure represents the Read operation response
type ReadResponse struct {
	// XXX: cb is an implicit input depedency for output parameters
	Length uint32 `idl:"name:cb" json:"length"`

	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	Pv         []byte         `idl:"name:pv;size_is:(cb)" json:"pv"`
	ReadLength uint32         `idl:"name:pcbRead" json:"read_length"`
	// Return: The Read return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReadResponse) xxx_ToOp(ctx context.Context, op *xxx_ReadOperation) *xxx_ReadOperation {
	if op == nil {
		op = &xxx_ReadOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.Length == uint32(0) {
		op.Length = o.Length
	}

	op.That = o.That
	op.Pv = o.Pv
	op.ReadLength = o.ReadLength
	op.Return = o.Return
	return op
}

func (o *ReadResponse) xxx_FromOp(ctx context.Context, op *xxx_ReadOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.Length = op.Length

	o.That = op.That
	o.Pv = op.Pv
	o.ReadLength = op.ReadLength
	o.Return = op.Return
}
func (o *ReadResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ReadResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReadOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WriteOperation structure represents the Write operation
type xxx_WriteOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	Pv            []byte         `idl:"name:pv;size_is:(cb)" json:"pv"`
	Length        uint32         `idl:"name:cb" json:"length"`
	WrittenLength uint32         `idl:"name:pcbWritten" json:"written_length"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_WriteOperation) OpNum() int { return 4 }

func (o *xxx_WriteOperation) OpName() string { return "/IStream/v0/Write" }

func (o *xxx_WriteOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Pv != nil && o.Length == 0 {
		o.Length = uint32(len(o.Pv))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WriteOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pv {in} (1:{pointer=ref}*(1))(2:{alias=BYTE}[dim:0,size_is=cb](uint8))
	{
		dimSize1 := uint64(o.Length)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Pv {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Pv[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Pv); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// cb {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Length); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WriteOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pv {in} (1:{pointer=ref}*(1))(2:{alias=BYTE}[dim:0,size_is=cb](uint8))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Pv", sizeInfo[0])
		}
		o.Pv = make([]byte, sizeInfo[0])
		for i1 := range o.Pv {
			i1 := i1
			if err := w.ReadData(&o.Pv[i1]); err != nil {
				return err
			}
		}
	}
	// cb {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Length); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WriteOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WriteOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pcbWritten {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.WrittenLength); err != nil {
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

func (o *xxx_WriteOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pcbWritten {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.WrittenLength); err != nil {
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

// WriteRequest structure represents the Write operation request
type WriteRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	Pv     []byte         `idl:"name:pv;size_is:(cb)" json:"pv"`
	Length uint32         `idl:"name:cb" json:"length"`
}

func (o *WriteRequest) xxx_ToOp(ctx context.Context, op *xxx_WriteOperation) *xxx_WriteOperation {
	if op == nil {
		op = &xxx_WriteOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Pv = o.Pv
	op.Length = o.Length
	return op
}

func (o *WriteRequest) xxx_FromOp(ctx context.Context, op *xxx_WriteOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Pv = op.Pv
	o.Length = op.Length
}
func (o *WriteRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WriteRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WriteOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WriteResponse structure represents the Write operation response
type WriteResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	WrittenLength uint32         `idl:"name:pcbWritten" json:"written_length"`
	// Return: The Write return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *WriteResponse) xxx_ToOp(ctx context.Context, op *xxx_WriteOperation) *xxx_WriteOperation {
	if op == nil {
		op = &xxx_WriteOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.WrittenLength = o.WrittenLength
	op.Return = o.Return
	return op
}

func (o *WriteResponse) xxx_FromOp(ctx context.Context, op *xxx_WriteOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.WrittenLength = op.WrittenLength
	o.Return = op.Return
}
func (o *WriteResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WriteResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WriteOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SeekOperation structure represents the Seek operation
type xxx_SeekOperation struct {
	This            *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DlibMove        *dtyp.LargeInteger  `idl:"name:dlibMove" json:"dlib_move"`
	Origin          uint32              `idl:"name:dwOrigin" json:"origin"`
	PlibNewPosition *dtyp.UlargeInteger `idl:"name:plibNewPosition" json:"plib_new_position"`
	Return          int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_SeekOperation) OpNum() int { return 5 }

func (o *xxx_SeekOperation) OpName() string { return "/IStream/v0/Seek" }

func (o *xxx_SeekOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SeekOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// dlibMove {in} (1:{alias=LARGE_INTEGER}(struct))
	{
		if o.DlibMove != nil {
			if err := o.DlibMove.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwOrigin {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Origin); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SeekOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// dlibMove {in} (1:{alias=LARGE_INTEGER}(struct))
	{
		if o.DlibMove == nil {
			o.DlibMove = &dtyp.LargeInteger{}
		}
		if err := o.DlibMove.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwOrigin {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Origin); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SeekOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SeekOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plibNewPosition {out} (1:{pointer=ref}*(1))(2:{alias=ULARGE_INTEGER}(struct))
	{
		if o.PlibNewPosition != nil {
			if err := o.PlibNewPosition.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UlargeInteger{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
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

func (o *xxx_SeekOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plibNewPosition {out} (1:{pointer=ref}*(1))(2:{alias=ULARGE_INTEGER}(struct))
	{
		if o.PlibNewPosition == nil {
			o.PlibNewPosition = &dtyp.UlargeInteger{}
		}
		if err := o.PlibNewPosition.UnmarshalNDR(ctx, w); err != nil {
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

// SeekRequest structure represents the Seek operation request
type SeekRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis     `idl:"name:This" json:"this"`
	DlibMove *dtyp.LargeInteger `idl:"name:dlibMove" json:"dlib_move"`
	Origin   uint32             `idl:"name:dwOrigin" json:"origin"`
}

func (o *SeekRequest) xxx_ToOp(ctx context.Context, op *xxx_SeekOperation) *xxx_SeekOperation {
	if op == nil {
		op = &xxx_SeekOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.DlibMove = o.DlibMove
	op.Origin = o.Origin
	return op
}

func (o *SeekRequest) xxx_FromOp(ctx context.Context, op *xxx_SeekOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DlibMove = op.DlibMove
	o.Origin = op.Origin
}
func (o *SeekRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SeekRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SeekOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SeekResponse structure represents the Seek operation response
type SeekResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That            *dcom.ORPCThat      `idl:"name:That" json:"that"`
	PlibNewPosition *dtyp.UlargeInteger `idl:"name:plibNewPosition" json:"plib_new_position"`
	// Return: The Seek return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SeekResponse) xxx_ToOp(ctx context.Context, op *xxx_SeekOperation) *xxx_SeekOperation {
	if op == nil {
		op = &xxx_SeekOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PlibNewPosition = o.PlibNewPosition
	op.Return = o.Return
	return op
}

func (o *SeekResponse) xxx_FromOp(ctx context.Context, op *xxx_SeekOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PlibNewPosition = op.PlibNewPosition
	o.Return = op.Return
}
func (o *SeekResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SeekResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SeekOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSizeOperation structure represents the SetSize operation
type xxx_SetSizeOperation struct {
	This       *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat      `idl:"name:That" json:"that"`
	LibNewSize *dtyp.UlargeInteger `idl:"name:libNewSize" json:"lib_new_size"`
	Return     int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSizeOperation) OpNum() int { return 6 }

func (o *xxx_SetSizeOperation) OpName() string { return "/IStream/v0/SetSize" }

func (o *xxx_SetSizeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSizeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// libNewSize {in} (1:{alias=ULARGE_INTEGER}(struct))
	{
		if o.LibNewSize != nil {
			if err := o.LibNewSize.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UlargeInteger{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_SetSizeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// libNewSize {in} (1:{alias=ULARGE_INTEGER}(struct))
	{
		if o.LibNewSize == nil {
			o.LibNewSize = &dtyp.UlargeInteger{}
		}
		if err := o.LibNewSize.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSizeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSizeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSizeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetSizeRequest structure represents the SetSize operation request
type SetSizeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis      `idl:"name:This" json:"this"`
	LibNewSize *dtyp.UlargeInteger `idl:"name:libNewSize" json:"lib_new_size"`
}

func (o *SetSizeRequest) xxx_ToOp(ctx context.Context, op *xxx_SetSizeOperation) *xxx_SetSizeOperation {
	if op == nil {
		op = &xxx_SetSizeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.LibNewSize = o.LibNewSize
	return op
}

func (o *SetSizeRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSizeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.LibNewSize = op.LibNewSize
}
func (o *SetSizeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetSizeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSizeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSizeResponse structure represents the SetSize operation response
type SetSizeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetSize return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSizeResponse) xxx_ToOp(ctx context.Context, op *xxx_SetSizeOperation) *xxx_SetSizeOperation {
	if op == nil {
		op = &xxx_SetSizeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetSizeResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSizeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSizeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetSizeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSizeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CopyToOperation structure represents the CopyTo operation
type xxx_CopyToOperation struct {
	This          *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat      `idl:"name:That" json:"that"`
	Pstm          *urlmon.Stream      `idl:"name:pstm" json:"pstm"`
	Length        *dtyp.UlargeInteger `idl:"name:cb" json:"length"`
	ReadLength    *dtyp.UlargeInteger `idl:"name:pcbRead" json:"read_length"`
	WrittenLength *dtyp.UlargeInteger `idl:"name:pcbWritten" json:"written_length"`
	Return        int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_CopyToOperation) OpNum() int { return 7 }

func (o *xxx_CopyToOperation) OpName() string { return "/IStream/v0/CopyTo" }

func (o *xxx_CopyToOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CopyToOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pstm {in} (1:{pointer=ref}*(1))(2:{alias=IStream}(interface))
	{
		if o.Pstm != nil {
			_ptr_pstm := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Pstm != nil {
					if err := o.Pstm.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Stream{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Pstm, _ptr_pstm); err != nil {
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
	// cb {in} (1:{alias=ULARGE_INTEGER}(struct))
	{
		if o.Length != nil {
			if err := o.Length.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UlargeInteger{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CopyToOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pstm {in} (1:{pointer=ref}*(1))(2:{alias=IStream}(interface))
	{
		_ptr_pstm := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Pstm == nil {
				o.Pstm = &urlmon.Stream{}
			}
			if err := o.Pstm.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pstm := func(ptr interface{}) { o.Pstm = *ptr.(**urlmon.Stream) }
		if err := w.ReadPointer(&o.Pstm, _s_pstm, _ptr_pstm); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cb {in} (1:{alias=ULARGE_INTEGER}(struct))
	{
		if o.Length == nil {
			o.Length = &dtyp.UlargeInteger{}
		}
		if err := o.Length.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CopyToOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CopyToOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pcbRead {out} (1:{pointer=ref}*(1))(2:{alias=ULARGE_INTEGER}(struct))
	{
		if o.ReadLength != nil {
			if err := o.ReadLength.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UlargeInteger{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pcbWritten {out} (1:{pointer=ref}*(1))(2:{alias=ULARGE_INTEGER}(struct))
	{
		if o.WrittenLength != nil {
			if err := o.WrittenLength.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UlargeInteger{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
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

func (o *xxx_CopyToOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pcbRead {out} (1:{pointer=ref}*(1))(2:{alias=ULARGE_INTEGER}(struct))
	{
		if o.ReadLength == nil {
			o.ReadLength = &dtyp.UlargeInteger{}
		}
		if err := o.ReadLength.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pcbWritten {out} (1:{pointer=ref}*(1))(2:{alias=ULARGE_INTEGER}(struct))
	{
		if o.WrittenLength == nil {
			o.WrittenLength = &dtyp.UlargeInteger{}
		}
		if err := o.WrittenLength.UnmarshalNDR(ctx, w); err != nil {
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

// CopyToRequest structure represents the CopyTo operation request
type CopyToRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This   *dcom.ORPCThis      `idl:"name:This" json:"this"`
	Pstm   *urlmon.Stream      `idl:"name:pstm" json:"pstm"`
	Length *dtyp.UlargeInteger `idl:"name:cb" json:"length"`
}

func (o *CopyToRequest) xxx_ToOp(ctx context.Context, op *xxx_CopyToOperation) *xxx_CopyToOperation {
	if op == nil {
		op = &xxx_CopyToOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Pstm = o.Pstm
	op.Length = o.Length
	return op
}

func (o *CopyToRequest) xxx_FromOp(ctx context.Context, op *xxx_CopyToOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Pstm = op.Pstm
	o.Length = op.Length
}
func (o *CopyToRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CopyToRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CopyToOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CopyToResponse structure represents the CopyTo operation response
type CopyToResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That          *dcom.ORPCThat      `idl:"name:That" json:"that"`
	ReadLength    *dtyp.UlargeInteger `idl:"name:pcbRead" json:"read_length"`
	WrittenLength *dtyp.UlargeInteger `idl:"name:pcbWritten" json:"written_length"`
	// Return: The CopyTo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CopyToResponse) xxx_ToOp(ctx context.Context, op *xxx_CopyToOperation) *xxx_CopyToOperation {
	if op == nil {
		op = &xxx_CopyToOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReadLength = o.ReadLength
	op.WrittenLength = o.WrittenLength
	op.Return = o.Return
	return op
}

func (o *CopyToResponse) xxx_FromOp(ctx context.Context, op *xxx_CopyToOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReadLength = op.ReadLength
	o.WrittenLength = op.WrittenLength
	o.Return = op.Return
}
func (o *CopyToResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CopyToResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CopyToOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CommitOperation structure represents the Commit operation
type xxx_CommitOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	GrfCommitFlags uint32         `idl:"name:grfCommitFlags" json:"grf_commit_flags"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CommitOperation) OpNum() int { return 8 }

func (o *xxx_CommitOperation) OpName() string { return "/IStream/v0/Commit" }

func (o *xxx_CommitOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// grfCommitFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.GrfCommitFlags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// grfCommitFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.GrfCommitFlags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CommitOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CommitRequest structure represents the Commit operation request
type CommitRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	GrfCommitFlags uint32         `idl:"name:grfCommitFlags" json:"grf_commit_flags"`
}

func (o *CommitRequest) xxx_ToOp(ctx context.Context, op *xxx_CommitOperation) *xxx_CommitOperation {
	if op == nil {
		op = &xxx_CommitOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.GrfCommitFlags = o.GrfCommitFlags
	return op
}

func (o *CommitRequest) xxx_FromOp(ctx context.Context, op *xxx_CommitOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.GrfCommitFlags = op.GrfCommitFlags
}
func (o *CommitRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CommitRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CommitOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CommitResponse structure represents the Commit operation response
type CommitResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Commit return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CommitResponse) xxx_ToOp(ctx context.Context, op *xxx_CommitOperation) *xxx_CommitOperation {
	if op == nil {
		op = &xxx_CommitOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CommitResponse) xxx_FromOp(ctx context.Context, op *xxx_CommitOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CommitResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CommitResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CommitOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RevertOperation structure represents the Revert operation
type xxx_RevertOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RevertOperation) OpNum() int { return 9 }

func (o *xxx_RevertOperation) OpName() string { return "/IStream/v0/Revert" }

func (o *xxx_RevertOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RevertOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RevertOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_RevertOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RevertOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RevertOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RevertRequest structure represents the Revert operation request
type RevertRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *RevertRequest) xxx_ToOp(ctx context.Context, op *xxx_RevertOperation) *xxx_RevertOperation {
	if op == nil {
		op = &xxx_RevertOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *RevertRequest) xxx_FromOp(ctx context.Context, op *xxx_RevertOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *RevertRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RevertRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RevertOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RevertResponse structure represents the Revert operation response
type RevertResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Revert return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RevertResponse) xxx_ToOp(ctx context.Context, op *xxx_RevertOperation) *xxx_RevertOperation {
	if op == nil {
		op = &xxx_RevertOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RevertResponse) xxx_FromOp(ctx context.Context, op *xxx_RevertOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RevertResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RevertResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RevertOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_LockRegionOperation structure represents the LockRegion operation
type xxx_LockRegionOperation struct {
	This      *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat      `idl:"name:That" json:"that"`
	LibOffset *dtyp.UlargeInteger `idl:"name:libOffset" json:"lib_offset"`
	Length    *dtyp.UlargeInteger `idl:"name:cb" json:"length"`
	LockType  uint32              `idl:"name:dwLockType" json:"lock_type"`
	Return    int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_LockRegionOperation) OpNum() int { return 10 }

func (o *xxx_LockRegionOperation) OpName() string { return "/IStream/v0/LockRegion" }

func (o *xxx_LockRegionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LockRegionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// libOffset {in} (1:{alias=ULARGE_INTEGER}(struct))
	{
		if o.LibOffset != nil {
			if err := o.LibOffset.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UlargeInteger{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// cb {in} (1:{alias=ULARGE_INTEGER}(struct))
	{
		if o.Length != nil {
			if err := o.Length.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UlargeInteger{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwLockType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.LockType); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LockRegionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// libOffset {in} (1:{alias=ULARGE_INTEGER}(struct))
	{
		if o.LibOffset == nil {
			o.LibOffset = &dtyp.UlargeInteger{}
		}
		if err := o.LibOffset.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// cb {in} (1:{alias=ULARGE_INTEGER}(struct))
	{
		if o.Length == nil {
			o.Length = &dtyp.UlargeInteger{}
		}
		if err := o.Length.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwLockType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.LockType); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LockRegionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LockRegionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_LockRegionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// LockRegionRequest structure represents the LockRegion operation request
type LockRegionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This      *dcom.ORPCThis      `idl:"name:This" json:"this"`
	LibOffset *dtyp.UlargeInteger `idl:"name:libOffset" json:"lib_offset"`
	Length    *dtyp.UlargeInteger `idl:"name:cb" json:"length"`
	LockType  uint32              `idl:"name:dwLockType" json:"lock_type"`
}

func (o *LockRegionRequest) xxx_ToOp(ctx context.Context, op *xxx_LockRegionOperation) *xxx_LockRegionOperation {
	if op == nil {
		op = &xxx_LockRegionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.LibOffset = o.LibOffset
	op.Length = o.Length
	op.LockType = o.LockType
	return op
}

func (o *LockRegionRequest) xxx_FromOp(ctx context.Context, op *xxx_LockRegionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.LibOffset = op.LibOffset
	o.Length = op.Length
	o.LockType = op.LockType
}
func (o *LockRegionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *LockRegionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LockRegionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LockRegionResponse structure represents the LockRegion operation response
type LockRegionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The LockRegion return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *LockRegionResponse) xxx_ToOp(ctx context.Context, op *xxx_LockRegionOperation) *xxx_LockRegionOperation {
	if op == nil {
		op = &xxx_LockRegionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *LockRegionResponse) xxx_FromOp(ctx context.Context, op *xxx_LockRegionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *LockRegionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *LockRegionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LockRegionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_UnlockRegionOperation structure represents the UnlockRegion operation
type xxx_UnlockRegionOperation struct {
	This      *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat      `idl:"name:That" json:"that"`
	LibOffset *dtyp.UlargeInteger `idl:"name:libOffset" json:"lib_offset"`
	Length    *dtyp.UlargeInteger `idl:"name:cb" json:"length"`
	LockType  uint32              `idl:"name:dwLockType" json:"lock_type"`
	Return    int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_UnlockRegionOperation) OpNum() int { return 11 }

func (o *xxx_UnlockRegionOperation) OpName() string { return "/IStream/v0/UnlockRegion" }

func (o *xxx_UnlockRegionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnlockRegionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// libOffset {in} (1:{alias=ULARGE_INTEGER}(struct))
	{
		if o.LibOffset != nil {
			if err := o.LibOffset.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UlargeInteger{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// cb {in} (1:{alias=ULARGE_INTEGER}(struct))
	{
		if o.Length != nil {
			if err := o.Length.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UlargeInteger{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwLockType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.LockType); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnlockRegionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// libOffset {in} (1:{alias=ULARGE_INTEGER}(struct))
	{
		if o.LibOffset == nil {
			o.LibOffset = &dtyp.UlargeInteger{}
		}
		if err := o.LibOffset.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// cb {in} (1:{alias=ULARGE_INTEGER}(struct))
	{
		if o.Length == nil {
			o.Length = &dtyp.UlargeInteger{}
		}
		if err := o.Length.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwLockType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.LockType); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnlockRegionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnlockRegionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_UnlockRegionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// UnlockRegionRequest structure represents the UnlockRegion operation request
type UnlockRegionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This      *dcom.ORPCThis      `idl:"name:This" json:"this"`
	LibOffset *dtyp.UlargeInteger `idl:"name:libOffset" json:"lib_offset"`
	Length    *dtyp.UlargeInteger `idl:"name:cb" json:"length"`
	LockType  uint32              `idl:"name:dwLockType" json:"lock_type"`
}

func (o *UnlockRegionRequest) xxx_ToOp(ctx context.Context, op *xxx_UnlockRegionOperation) *xxx_UnlockRegionOperation {
	if op == nil {
		op = &xxx_UnlockRegionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.LibOffset = o.LibOffset
	op.Length = o.Length
	op.LockType = o.LockType
	return op
}

func (o *UnlockRegionRequest) xxx_FromOp(ctx context.Context, op *xxx_UnlockRegionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.LibOffset = op.LibOffset
	o.Length = op.Length
	o.LockType = op.LockType
}
func (o *UnlockRegionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *UnlockRegionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UnlockRegionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// UnlockRegionResponse structure represents the UnlockRegion operation response
type UnlockRegionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The UnlockRegion return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *UnlockRegionResponse) xxx_ToOp(ctx context.Context, op *xxx_UnlockRegionOperation) *xxx_UnlockRegionOperation {
	if op == nil {
		op = &xxx_UnlockRegionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *UnlockRegionResponse) xxx_FromOp(ctx context.Context, op *xxx_UnlockRegionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *UnlockRegionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *UnlockRegionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UnlockRegionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_StatOperation structure represents the Stat operation
type xxx_StatOperation struct {
	This        *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Pstatstg    *urlmon.Statstg `idl:"name:pstatstg" json:"pstatstg"`
	GrfStatFlag uint32          `idl:"name:grfStatFlag" json:"grf_stat_flag"`
	Return      int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_StatOperation) OpNum() int { return 12 }

func (o *xxx_StatOperation) OpName() string { return "/IStream/v0/Stat" }

func (o *xxx_StatOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StatOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// grfStatFlag {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.GrfStatFlag); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StatOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// grfStatFlag {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.GrfStatFlag); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StatOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StatOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pstatstg {out} (1:{pointer=ref}*(1))(2:{alias=STATSTG}(struct))
	{
		if o.Pstatstg != nil {
			if err := o.Pstatstg.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&urlmon.Statstg{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_StatOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pstatstg {out} (1:{pointer=ref}*(1))(2:{alias=STATSTG}(struct))
	{
		if o.Pstatstg == nil {
			o.Pstatstg = &urlmon.Statstg{}
		}
		if err := o.Pstatstg.UnmarshalNDR(ctx, w); err != nil {
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

// StatRequest structure represents the Stat operation request
type StatRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	GrfStatFlag uint32         `idl:"name:grfStatFlag" json:"grf_stat_flag"`
}

func (o *StatRequest) xxx_ToOp(ctx context.Context, op *xxx_StatOperation) *xxx_StatOperation {
	if op == nil {
		op = &xxx_StatOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.GrfStatFlag = o.GrfStatFlag
	return op
}

func (o *StatRequest) xxx_FromOp(ctx context.Context, op *xxx_StatOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.GrfStatFlag = op.GrfStatFlag
}
func (o *StatRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *StatRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StatOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// StatResponse structure represents the Stat operation response
type StatResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Pstatstg *urlmon.Statstg `idl:"name:pstatstg" json:"pstatstg"`
	// Return: The Stat return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *StatResponse) xxx_ToOp(ctx context.Context, op *xxx_StatOperation) *xxx_StatOperation {
	if op == nil {
		op = &xxx_StatOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Pstatstg = o.Pstatstg
	op.Return = o.Return
	return op
}

func (o *StatResponse) xxx_FromOp(ctx context.Context, op *xxx_StatOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Pstatstg = op.Pstatstg
	o.Return = op.Return
}
func (o *StatResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *StatResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StatOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloneOperation structure represents the Clone operation
type xxx_CloneOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Ppstm  *urlmon.Stream `idl:"name:ppstm" json:"ppstm"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CloneOperation) OpNum() int { return 13 }

func (o *xxx_CloneOperation) OpName() string { return "/IStream/v0/Clone" }

func (o *xxx_CloneOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloneOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CloneOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_CloneOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloneOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppstm {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IStream}(interface))
	{
		if o.Ppstm != nil {
			_ptr_ppstm := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Ppstm != nil {
					if err := o.Ppstm.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&urlmon.Stream{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Ppstm, _ptr_ppstm); err != nil {
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

func (o *xxx_CloneOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppstm {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IStream}(interface))
	{
		_ptr_ppstm := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Ppstm == nil {
				o.Ppstm = &urlmon.Stream{}
			}
			if err := o.Ppstm.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppstm := func(ptr interface{}) { o.Ppstm = *ptr.(**urlmon.Stream) }
		if err := w.ReadPointer(&o.Ppstm, _s_ppstm, _ptr_ppstm); err != nil {
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

// CloneRequest structure represents the Clone operation request
type CloneRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *CloneRequest) xxx_ToOp(ctx context.Context, op *xxx_CloneOperation) *xxx_CloneOperation {
	if op == nil {
		op = &xxx_CloneOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *CloneRequest) xxx_FromOp(ctx context.Context, op *xxx_CloneOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *CloneRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CloneRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloneOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloneResponse structure represents the Clone operation response
type CloneResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat `idl:"name:That" json:"that"`
	Ppstm *urlmon.Stream `idl:"name:ppstm" json:"ppstm"`
	// Return: The Clone return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CloneResponse) xxx_ToOp(ctx context.Context, op *xxx_CloneOperation) *xxx_CloneOperation {
	if op == nil {
		op = &xxx_CloneOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Ppstm = o.Ppstm
	op.Return = o.Return
	return op
}

func (o *CloneResponse) xxx_FromOp(ctx context.Context, op *xxx_CloneOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Ppstm = op.Ppstm
	o.Return = op.Return
}
func (o *CloneResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CloneResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloneOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
