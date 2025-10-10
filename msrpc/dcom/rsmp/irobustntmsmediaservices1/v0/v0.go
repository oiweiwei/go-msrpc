package irobustntmsmediaservices1

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	intmsmediaservices1 "github.com/oiweiwei/go-msrpc/msrpc/dcom/rsmp/intmsmediaservices1/v0"
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
	_ = intmsmediaservices1.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/rsmp"
)

var (
	// IRobustNtmsMediaServices1 interface identifier 7d07f313-a53f-459a-bb12-012c15b1846e
	RobustNTMSMediaServices1IID = &dcom.IID{Data1: 0x7d07f313, Data2: 0xa53f, Data3: 0x459a, Data4: []byte{0xbb, 0x12, 0x01, 0x2c, 0x15, 0xb1, 0x84, 0x6e}}
	// Syntax UUID
	RobustNTMSMediaServices1SyntaxUUID = &uuid.UUID{TimeLow: 0x7d07f313, TimeMid: 0xa53f, TimeHiAndVersion: 0x459a, ClockSeqHiAndReserved: 0xbb, ClockSeqLow: 0x12, Node: [6]uint8{0x1, 0x2c, 0x15, 0xb1, 0x84, 0x6e}}
	// Syntax ID
	RobustNTMSMediaServices1SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: RobustNTMSMediaServices1SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IRobustNtmsMediaServices1 interface.
type RobustNTMSMediaServices1Client interface {

	// INtmsMediaServices1 retrieval method.
	NTMSMediaServices1() intmsmediaservices1.NTMSMediaServices1Client

	GetNTMSMediaPoolNameAr(context.Context, *GetNTMSMediaPoolNameArRequest, ...dcerpc.CallOption) (*GetNTMSMediaPoolNameArResponse, error)

	GetNTMSMediaPoolNameWr(context.Context, *GetNTMSMediaPoolNameWrRequest, ...dcerpc.CallOption) (*GetNTMSMediaPoolNameWrResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) RobustNTMSMediaServices1Client
}

type xxx_DefaultRobustNTMSMediaServices1Client struct {
	intmsmediaservices1.NTMSMediaServices1Client
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultRobustNTMSMediaServices1Client) NTMSMediaServices1() intmsmediaservices1.NTMSMediaServices1Client {
	return o.NTMSMediaServices1Client
}

func (o *xxx_DefaultRobustNTMSMediaServices1Client) GetNTMSMediaPoolNameAr(ctx context.Context, in *GetNTMSMediaPoolNameArRequest, opts ...dcerpc.CallOption) (*GetNTMSMediaPoolNameArResponse, error) {
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
	out := &GetNTMSMediaPoolNameArResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRobustNTMSMediaServices1Client) GetNTMSMediaPoolNameWr(ctx context.Context, in *GetNTMSMediaPoolNameWrRequest, opts ...dcerpc.CallOption) (*GetNTMSMediaPoolNameWrResponse, error) {
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
	out := &GetNTMSMediaPoolNameWrResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRobustNTMSMediaServices1Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultRobustNTMSMediaServices1Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultRobustNTMSMediaServices1Client) IPID(ctx context.Context, ipid *dcom.IPID) RobustNTMSMediaServices1Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultRobustNTMSMediaServices1Client{
		NTMSMediaServices1Client: o.NTMSMediaServices1Client.IPID(ctx, ipid),
		cc:                       o.cc,
		ipid:                     ipid,
	}
}

func NewRobustNTMSMediaServices1Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (RobustNTMSMediaServices1Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(RobustNTMSMediaServices1SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := intmsmediaservices1.NewNTMSMediaServices1Client(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultRobustNTMSMediaServices1Client{
		NTMSMediaServices1Client: base,
		cc:                       cc,
		ipid:                     ipid,
	}, nil
}

// xxx_GetNTMSMediaPoolNameArOperation structure represents the GetNtmsMediaPoolNameAR operation
type xxx_GetNTMSMediaPoolNameArOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	PoolID         *dtyp.GUID     `idl:"name:lpPoolId" json:"pool_id"`
	BufferName     []byte         `idl:"name:lpBufName;size_is:(lpdwNameSizeBuf);length_is:(lpdwNameSize)" json:"buffer_name"`
	NameSizeBuffer uint32         `idl:"name:lpdwNameSizeBuf" json:"name_size_buffer"`
	NameSize       uint32         `idl:"name:lpdwNameSize" json:"name_size"`
	OutputSize     uint32         `idl:"name:lpdwOutputSize" json:"output_size"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNTMSMediaPoolNameArOperation) OpNum() int { return 18 }

func (o *xxx_GetNTMSMediaPoolNameArOperation) OpName() string {
	return "/IRobustNtmsMediaServices1/v0/GetNtmsMediaPoolNameAR"
}

func (o *xxx_GetNTMSMediaPoolNameArOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSMediaPoolNameArOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpPoolId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.PoolID != nil {
			if err := o.PoolID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpdwNameSizeBuf {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NameSizeBuffer); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSMediaPoolNameArOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpPoolId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.PoolID == nil {
			o.PoolID = &dtyp.GUID{}
		}
		if err := o.PoolID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpdwNameSizeBuf {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NameSizeBuffer); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSMediaPoolNameArOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.BufferName != nil && o.NameSize == 0 {
		o.NameSize = uint32(len(o.BufferName))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSMediaPoolNameArOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpBufName {out} (1:{pointer=ref}*(1)[dim:0,size_is=lpdwNameSizeBuf,length_is=lpdwNameSize](uchar))
	{
		dimSize1 := uint64(o.NameSizeBuffer)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.NameSize)
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
		for i1 := range o.BufferName {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.BufferName[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.BufferName); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// lpdwNameSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NameSize); err != nil {
			return err
		}
	}
	// lpdwOutputSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OutputSize); err != nil {
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

func (o *xxx_GetNTMSMediaPoolNameArOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpBufName {out} (1:{pointer=ref}*(1)[dim:0,size_is=lpdwNameSizeBuf,length_is=lpdwNameSize](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.BufferName", sizeInfo[0])
		}
		o.BufferName = make([]byte, sizeInfo[0])
		for i1 := range o.BufferName {
			i1 := i1
			if err := w.ReadData(&o.BufferName[i1]); err != nil {
				return err
			}
		}
	}
	// lpdwNameSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NameSize); err != nil {
			return err
		}
	}
	// lpdwOutputSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OutputSize); err != nil {
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

// GetNTMSMediaPoolNameArRequest structure represents the GetNtmsMediaPoolNameAR operation request
type GetNTMSMediaPoolNameArRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	PoolID         *dtyp.GUID     `idl:"name:lpPoolId" json:"pool_id"`
	NameSizeBuffer uint32         `idl:"name:lpdwNameSizeBuf" json:"name_size_buffer"`
}

func (o *GetNTMSMediaPoolNameArRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSMediaPoolNameArOperation) *xxx_GetNTMSMediaPoolNameArOperation {
	if op == nil {
		op = &xxx_GetNTMSMediaPoolNameArOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PoolID = o.PoolID
	op.NameSizeBuffer = o.NameSizeBuffer
	return op
}

func (o *GetNTMSMediaPoolNameArRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSMediaPoolNameArOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PoolID = op.PoolID
	o.NameSizeBuffer = op.NameSizeBuffer
}
func (o *GetNTMSMediaPoolNameArRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNTMSMediaPoolNameArRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSMediaPoolNameArOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNTMSMediaPoolNameArResponse structure represents the GetNtmsMediaPoolNameAR operation response
type GetNTMSMediaPoolNameArResponse struct {
	// XXX: lpdwNameSizeBuf is an implicit input depedency for output parameters
	NameSizeBuffer uint32 `idl:"name:lpdwNameSizeBuf" json:"name_size_buffer"`

	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	BufferName []byte         `idl:"name:lpBufName;size_is:(lpdwNameSizeBuf);length_is:(lpdwNameSize)" json:"buffer_name"`
	NameSize   uint32         `idl:"name:lpdwNameSize" json:"name_size"`
	OutputSize uint32         `idl:"name:lpdwOutputSize" json:"output_size"`
	// Return: The GetNtmsMediaPoolNameAR return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNTMSMediaPoolNameArResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSMediaPoolNameArOperation) *xxx_GetNTMSMediaPoolNameArOperation {
	if op == nil {
		op = &xxx_GetNTMSMediaPoolNameArOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.NameSizeBuffer == uint32(0) {
		op.NameSizeBuffer = o.NameSizeBuffer
	}

	op.That = o.That
	op.BufferName = o.BufferName
	op.NameSize = o.NameSize
	op.OutputSize = o.OutputSize
	op.Return = o.Return
	return op
}

func (o *GetNTMSMediaPoolNameArResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSMediaPoolNameArOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.NameSizeBuffer = op.NameSizeBuffer

	o.That = op.That
	o.BufferName = op.BufferName
	o.NameSize = op.NameSize
	o.OutputSize = op.OutputSize
	o.Return = op.Return
}
func (o *GetNTMSMediaPoolNameArResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNTMSMediaPoolNameArResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSMediaPoolNameArOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNTMSMediaPoolNameWrOperation structure represents the GetNtmsMediaPoolNameWR operation
type xxx_GetNTMSMediaPoolNameWrOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	PoolID         *dtyp.GUID     `idl:"name:lpPoolId" json:"pool_id"`
	BufferName     string         `idl:"name:lpBufName;size_is:(lpdwNameSizeBuf);length_is:(lpdwNameSize)" json:"buffer_name"`
	NameSizeBuffer uint32         `idl:"name:lpdwNameSizeBuf" json:"name_size_buffer"`
	NameSize       uint32         `idl:"name:lpdwNameSize" json:"name_size"`
	OutputSize     uint32         `idl:"name:lpdwOutputSize" json:"output_size"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNTMSMediaPoolNameWrOperation) OpNum() int { return 19 }

func (o *xxx_GetNTMSMediaPoolNameWrOperation) OpName() string {
	return "/IRobustNtmsMediaServices1/v0/GetNtmsMediaPoolNameWR"
}

func (o *xxx_GetNTMSMediaPoolNameWrOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSMediaPoolNameWrOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpPoolId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.PoolID != nil {
			if err := o.PoolID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpdwNameSizeBuf {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NameSizeBuffer); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSMediaPoolNameWrOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpPoolId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.PoolID == nil {
			o.PoolID = &dtyp.GUID{}
		}
		if err := o.PoolID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpdwNameSizeBuf {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NameSizeBuffer); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSMediaPoolNameWrOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.BufferName != "" && o.NameSize == 0 {
		o.NameSize = uint32(ndr.UTF16Len(o.BufferName))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSMediaPoolNameWrOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpBufName {out} (1:{pointer=ref}*(1)[dim:0,size_is=lpdwNameSizeBuf,length_is=lpdwNameSize,string](wchar))
	{
		dimSize1 := uint64(o.NameSizeBuffer)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.NameSize)
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
		_BufferName_buf := utf16.Encode([]rune(o.BufferName))
		if uint64(len(_BufferName_buf)) > sizeInfo[0] {
			_BufferName_buf = _BufferName_buf[:sizeInfo[0]]
		}
		for i1 := range _BufferName_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_BufferName_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_BufferName_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	// lpdwNameSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NameSize); err != nil {
			return err
		}
	}
	// lpdwOutputSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OutputSize); err != nil {
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

func (o *xxx_GetNTMSMediaPoolNameWrOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpBufName {out} (1:{pointer=ref}*(1)[dim:0,size_is=lpdwNameSizeBuf,length_is=lpdwNameSize,string](wchar))
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
		var _BufferName_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _BufferName_buf", sizeInfo[0])
		}
		_BufferName_buf = make([]uint16, sizeInfo[0])
		for i1 := range _BufferName_buf {
			i1 := i1
			if err := w.ReadData(&_BufferName_buf[i1]); err != nil {
				return err
			}
		}
		o.BufferName = strings.TrimRight(string(utf16.Decode(_BufferName_buf)), ndr.ZeroString)
	}
	// lpdwNameSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NameSize); err != nil {
			return err
		}
	}
	// lpdwOutputSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OutputSize); err != nil {
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

// GetNTMSMediaPoolNameWrRequest structure represents the GetNtmsMediaPoolNameWR operation request
type GetNTMSMediaPoolNameWrRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	PoolID         *dtyp.GUID     `idl:"name:lpPoolId" json:"pool_id"`
	NameSizeBuffer uint32         `idl:"name:lpdwNameSizeBuf" json:"name_size_buffer"`
}

func (o *GetNTMSMediaPoolNameWrRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSMediaPoolNameWrOperation) *xxx_GetNTMSMediaPoolNameWrOperation {
	if op == nil {
		op = &xxx_GetNTMSMediaPoolNameWrOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PoolID = o.PoolID
	op.NameSizeBuffer = o.NameSizeBuffer
	return op
}

func (o *GetNTMSMediaPoolNameWrRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSMediaPoolNameWrOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PoolID = op.PoolID
	o.NameSizeBuffer = op.NameSizeBuffer
}
func (o *GetNTMSMediaPoolNameWrRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNTMSMediaPoolNameWrRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSMediaPoolNameWrOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNTMSMediaPoolNameWrResponse structure represents the GetNtmsMediaPoolNameWR operation response
type GetNTMSMediaPoolNameWrResponse struct {
	// XXX: lpdwNameSizeBuf is an implicit input depedency for output parameters
	NameSizeBuffer uint32 `idl:"name:lpdwNameSizeBuf" json:"name_size_buffer"`

	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	BufferName string         `idl:"name:lpBufName;size_is:(lpdwNameSizeBuf);length_is:(lpdwNameSize)" json:"buffer_name"`
	NameSize   uint32         `idl:"name:lpdwNameSize" json:"name_size"`
	OutputSize uint32         `idl:"name:lpdwOutputSize" json:"output_size"`
	// Return: The GetNtmsMediaPoolNameWR return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNTMSMediaPoolNameWrResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSMediaPoolNameWrOperation) *xxx_GetNTMSMediaPoolNameWrOperation {
	if op == nil {
		op = &xxx_GetNTMSMediaPoolNameWrOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.NameSizeBuffer == uint32(0) {
		op.NameSizeBuffer = o.NameSizeBuffer
	}

	op.That = o.That
	op.BufferName = o.BufferName
	op.NameSize = o.NameSize
	op.OutputSize = o.OutputSize
	op.Return = o.Return
	return op
}

func (o *GetNTMSMediaPoolNameWrResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSMediaPoolNameWrOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.NameSizeBuffer = op.NameSizeBuffer

	o.That = op.That
	o.BufferName = op.BufferName
	o.NameSize = op.NameSize
	o.OutputSize = op.OutputSize
	o.Return = op.Return
}
func (o *GetNTMSMediaPoolNameWrResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNTMSMediaPoolNameWrResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSMediaPoolNameWrOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
