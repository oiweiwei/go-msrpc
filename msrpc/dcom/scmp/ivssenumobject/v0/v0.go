package ivssenumobject

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
	scmp "github.com/oiweiwei/go-msrpc/msrpc/dcom/scmp"
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
	_ = scmp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/scmp"
)

var (
	// IVssEnumObject interface identifier ae1c7110-2f60-11d3-8a39-00c04f72d8e3
	VSSEnumObjectIID = &dcom.IID{Data1: 0xae1c7110, Data2: 0x2f60, Data3: 0x11d3, Data4: []byte{0x8a, 0x39, 0x00, 0xc0, 0x4f, 0x72, 0xd8, 0xe3}}
	// Syntax UUID
	VSSEnumObjectSyntaxUUID = &uuid.UUID{TimeLow: 0xae1c7110, TimeMid: 0x2f60, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x8a, ClockSeqLow: 0x39, Node: [6]uint8{0x0, 0xc0, 0x4f, 0x72, 0xd8, 0xe3}}
	// Syntax ID
	VSSEnumObjectSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: VSSEnumObjectSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVssEnumObject interface.
type VSSEnumObjectClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// Next operation.
	Next(context.Context, *NextRequest, ...dcerpc.CallOption) (*NextResponse, error)

	// Skip operation.
	Skip(context.Context, *SkipRequest, ...dcerpc.CallOption) (*SkipResponse, error)

	// Reset operation.
	Reset(context.Context, *ResetRequest, ...dcerpc.CallOption) (*ResetResponse, error)

	// Clone operation.
	Clone(context.Context, *CloneRequest, ...dcerpc.CallOption) (*CloneResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) VSSEnumObjectClient
}

type xxx_DefaultVSSEnumObjectClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultVSSEnumObjectClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultVSSEnumObjectClient) Next(ctx context.Context, in *NextRequest, opts ...dcerpc.CallOption) (*NextResponse, error) {
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
	out := &NextResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVSSEnumObjectClient) Skip(ctx context.Context, in *SkipRequest, opts ...dcerpc.CallOption) (*SkipResponse, error) {
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
	out := &SkipResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVSSEnumObjectClient) Reset(ctx context.Context, in *ResetRequest, opts ...dcerpc.CallOption) (*ResetResponse, error) {
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
	out := &ResetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVSSEnumObjectClient) Clone(ctx context.Context, in *CloneRequest, opts ...dcerpc.CallOption) (*CloneResponse, error) {
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

func (o *xxx_DefaultVSSEnumObjectClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultVSSEnumObjectClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultVSSEnumObjectClient) IPID(ctx context.Context, ipid *dcom.IPID) VSSEnumObjectClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultVSSEnumObjectClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewVSSEnumObjectClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (VSSEnumObjectClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(VSSEnumObjectSyntaxV0_0))...)
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
	return &xxx_DefaultVSSEnumObjectClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_NextOperation structure represents the Next operation
type xxx_NextOperation struct {
	This    *dcom.ORPCThis            `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat            `idl:"name:That" json:"that"`
	Count   uint32                    `idl:"name:celt" json:"count"`
	Entries []*scmp.VSSObjectProperty `idl:"name:rgelt;size_is:(celt);length_is:(pceltFetched)" json:"entries"`
	Fetched uint32                    `idl:"name:pceltFetched" json:"fetched"`
	Return  int32                     `idl:"name:Return" json:"return"`
}

func (o *xxx_NextOperation) OpNum() int { return 3 }

func (o *xxx_NextOperation) OpName() string { return "/IVssEnumObject/v0/Next" }

func (o *xxx_NextOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NextOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// celt {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NextOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// celt {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NextOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Entries != nil && o.Fetched == 0 {
		o.Fetched = uint32(len(o.Entries))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NextOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// rgelt {out} (1:{pointer=ref}*(1))(2:{alias=VSS_OBJECT_PROP}[dim:0,size_is=celt,length_is=pceltFetched](struct))
	{
		dimSize1 := uint64(o.Count)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.Fetched)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.Entries {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Entries[i1] != nil {
				if err := o.Entries[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&scmp.VSSObjectProperty{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Entries); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&scmp.VSSObjectProperty{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pceltFetched {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Fetched); err != nil {
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

func (o *xxx_NextOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// rgelt {out} (1:{pointer=ref}*(1))(2:{alias=VSS_OBJECT_PROP}[dim:0,size_is=celt,length_is=pceltFetched](struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Entries", sizeInfo[0])
		}
		o.Entries = make([]*scmp.VSSObjectProperty, sizeInfo[0])
		for i1 := range o.Entries {
			i1 := i1
			if o.Entries[i1] == nil {
				o.Entries[i1] = &scmp.VSSObjectProperty{}
			}
			if err := o.Entries[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pceltFetched {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Fetched); err != nil {
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

// NextRequest structure represents the Next operation request
type NextRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis `idl:"name:This" json:"this"`
	Count uint32         `idl:"name:celt" json:"count"`
}

func (o *NextRequest) xxx_ToOp(ctx context.Context, op *xxx_NextOperation) *xxx_NextOperation {
	if op == nil {
		op = &xxx_NextOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Count = o.Count
	return op
}

func (o *NextRequest) xxx_FromOp(ctx context.Context, op *xxx_NextOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Count = op.Count
}
func (o *NextRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *NextRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NextOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// NextResponse structure represents the Next operation response
type NextResponse struct {
	// XXX: celt is an implicit input depedency for output parameters
	Count uint32 `idl:"name:celt" json:"count"`

	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat            `idl:"name:That" json:"that"`
	Entries []*scmp.VSSObjectProperty `idl:"name:rgelt;size_is:(celt);length_is:(pceltFetched)" json:"entries"`
	Fetched uint32                    `idl:"name:pceltFetched" json:"fetched"`
	// Return: The Next return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *NextResponse) xxx_ToOp(ctx context.Context, op *xxx_NextOperation) *xxx_NextOperation {
	if op == nil {
		op = &xxx_NextOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.Count == uint32(0) {
		op.Count = o.Count
	}

	op.That = o.That
	op.Entries = o.Entries
	op.Fetched = o.Fetched
	op.Return = o.Return
	return op
}

func (o *NextResponse) xxx_FromOp(ctx context.Context, op *xxx_NextOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.Count = op.Count

	o.That = op.That
	o.Entries = op.Entries
	o.Fetched = op.Fetched
	o.Return = op.Return
}
func (o *NextResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *NextResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NextOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SkipOperation structure represents the Skip operation
type xxx_SkipOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Count  uint32         `idl:"name:celt" json:"count"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SkipOperation) OpNum() int { return 4 }

func (o *xxx_SkipOperation) OpName() string { return "/IVssEnumObject/v0/Skip" }

func (o *xxx_SkipOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SkipOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// celt {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SkipOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// celt {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SkipOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SkipOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SkipOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SkipRequest structure represents the Skip operation request
type SkipRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis `idl:"name:This" json:"this"`
	Count uint32         `idl:"name:celt" json:"count"`
}

func (o *SkipRequest) xxx_ToOp(ctx context.Context, op *xxx_SkipOperation) *xxx_SkipOperation {
	if op == nil {
		op = &xxx_SkipOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Count = o.Count
	return op
}

func (o *SkipRequest) xxx_FromOp(ctx context.Context, op *xxx_SkipOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Count = op.Count
}
func (o *SkipRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SkipRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SkipOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SkipResponse structure represents the Skip operation response
type SkipResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Skip return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SkipResponse) xxx_ToOp(ctx context.Context, op *xxx_SkipOperation) *xxx_SkipOperation {
	if op == nil {
		op = &xxx_SkipOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SkipResponse) xxx_FromOp(ctx context.Context, op *xxx_SkipOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SkipResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SkipResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SkipOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ResetOperation structure represents the Reset operation
type xxx_ResetOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ResetOperation) OpNum() int { return 5 }

func (o *xxx_ResetOperation) OpName() string { return "/IVssEnumObject/v0/Reset" }

func (o *xxx_ResetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ResetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_ResetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ResetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ResetRequest structure represents the Reset operation request
type ResetRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *ResetRequest) xxx_ToOp(ctx context.Context, op *xxx_ResetOperation) *xxx_ResetOperation {
	if op == nil {
		op = &xxx_ResetOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *ResetRequest) xxx_FromOp(ctx context.Context, op *xxx_ResetOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *ResetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ResetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ResetResponse structure represents the Reset operation response
type ResetResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Reset return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ResetResponse) xxx_ToOp(ctx context.Context, op *xxx_ResetOperation) *xxx_ResetOperation {
	if op == nil {
		op = &xxx_ResetOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ResetResponse) xxx_FromOp(ctx context.Context, op *xxx_ResetOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ResetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ResetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloneOperation structure represents the Clone operation
type xxx_CloneOperation struct {
	This   *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat      `idl:"name:That" json:"that"`
	Enum   *scmp.VSSEnumObject `idl:"name:ppenum" json:"enum"`
	Return int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_CloneOperation) OpNum() int { return 6 }

func (o *xxx_CloneOperation) OpName() string { return "/IVssEnumObject/v0/Clone" }

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
	// ppenum {in, out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVssEnumObject}(interface))
	{
		if o.Enum != nil {
			_ptr_ppenum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Enum != nil {
					if err := o.Enum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&scmp.VSSEnumObject{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Enum, _ptr_ppenum); err != nil {
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
	// ppenum {in, out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVssEnumObject}(interface))
	{
		_ptr_ppenum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Enum == nil {
				o.Enum = &scmp.VSSEnumObject{}
			}
			if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppenum := func(ptr interface{}) { o.Enum = *ptr.(**scmp.VSSEnumObject) }
		if err := w.ReadPointer(&o.Enum, _s_ppenum, _ptr_ppenum); err != nil {
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
	// ppenum {in, out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVssEnumObject}(interface))
	{
		if o.Enum != nil {
			_ptr_ppenum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Enum != nil {
					if err := o.Enum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&scmp.VSSEnumObject{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Enum, _ptr_ppenum); err != nil {
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
	// ppenum {in, out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVssEnumObject}(interface))
	{
		_ptr_ppenum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Enum == nil {
				o.Enum = &scmp.VSSEnumObject{}
			}
			if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppenum := func(ptr interface{}) { o.Enum = *ptr.(**scmp.VSSEnumObject) }
		if err := w.ReadPointer(&o.Enum, _s_ppenum, _ptr_ppenum); err != nil {
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
	This *dcom.ORPCThis      `idl:"name:This" json:"this"`
	Enum *scmp.VSSEnumObject `idl:"name:ppenum" json:"enum"`
}

func (o *CloneRequest) xxx_ToOp(ctx context.Context, op *xxx_CloneOperation) *xxx_CloneOperation {
	if op == nil {
		op = &xxx_CloneOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Enum = o.Enum
	return op
}

func (o *CloneRequest) xxx_FromOp(ctx context.Context, op *xxx_CloneOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Enum = op.Enum
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
	That *dcom.ORPCThat      `idl:"name:That" json:"that"`
	Enum *scmp.VSSEnumObject `idl:"name:ppenum" json:"enum"`
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
	op.Enum = o.Enum
	op.Return = o.Return
	return op
}

func (o *CloneResponse) xxx_FromOp(ctx context.Context, op *xxx_CloneOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enum = op.Enum
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
