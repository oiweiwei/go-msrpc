package intmsobjectmanagement3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	intmsobjectmanagement2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/rsmp/intmsobjectmanagement2/v0"
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
	_ = intmsobjectmanagement2.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/rsmp"
)

var (
	// INtmsObjectManagement3 interface identifier 3bbed8d9-2c9a-4b21-8936-acb2f995be6c
	NTMSObjectManagement3IID = &dcom.IID{Data1: 0x3bbed8d9, Data2: 0x2c9a, Data3: 0x4b21, Data4: []byte{0x89, 0x36, 0xac, 0xb2, 0xf9, 0x95, 0xbe, 0x6c}}
	// Syntax UUID
	NTMSObjectManagement3SyntaxUUID = &uuid.UUID{TimeLow: 0x3bbed8d9, TimeMid: 0x2c9a, TimeHiAndVersion: 0x4b21, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x36, Node: [6]uint8{0xac, 0xb2, 0xf9, 0x95, 0xbe, 0x6c}}
	// Syntax ID
	NTMSObjectManagement3SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: NTMSObjectManagement3SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// INtmsObjectManagement3 interface.
type NTMSObjectManagement3Client interface {

	// INtmsObjectManagement2 retrieval method.
	NTMSObjectManagement2() intmsobjectmanagement2.NTMSObjectManagement2Client

	GetNTMSObjectAttributeAr(context.Context, *GetNTMSObjectAttributeArRequest, ...dcerpc.CallOption) (*GetNTMSObjectAttributeArResponse, error)

	GetNTMSObjectAttributeWr(context.Context, *GetNTMSObjectAttributeWrRequest, ...dcerpc.CallOption) (*GetNTMSObjectAttributeWrResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) NTMSObjectManagement3Client
}

type xxx_DefaultNTMSObjectManagement3Client struct {
	intmsobjectmanagement2.NTMSObjectManagement2Client
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultNTMSObjectManagement3Client) NTMSObjectManagement2() intmsobjectmanagement2.NTMSObjectManagement2Client {
	return o.NTMSObjectManagement2Client
}

func (o *xxx_DefaultNTMSObjectManagement3Client) GetNTMSObjectAttributeAr(ctx context.Context, in *GetNTMSObjectAttributeArRequest, opts ...dcerpc.CallOption) (*GetNTMSObjectAttributeArResponse, error) {
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
	out := &GetNTMSObjectAttributeArResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNTMSObjectManagement3Client) GetNTMSObjectAttributeWr(ctx context.Context, in *GetNTMSObjectAttributeWrRequest, opts ...dcerpc.CallOption) (*GetNTMSObjectAttributeWrResponse, error) {
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
	out := &GetNTMSObjectAttributeWrResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNTMSObjectManagement3Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultNTMSObjectManagement3Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultNTMSObjectManagement3Client) IPID(ctx context.Context, ipid *dcom.IPID) NTMSObjectManagement3Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultNTMSObjectManagement3Client{
		NTMSObjectManagement2Client: o.NTMSObjectManagement2Client.IPID(ctx, ipid),
		cc:                          o.cc,
		ipid:                        ipid,
	}
}

func NewNTMSObjectManagement3Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (NTMSObjectManagement3Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(NTMSObjectManagement3SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := intmsobjectmanagement2.NewNTMSObjectManagement2Client(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultNTMSObjectManagement3Client{
		NTMSObjectManagement2Client: base,
		cc:                          cc,
		ipid:                        ipid,
	}, nil
}

// xxx_GetNTMSObjectAttributeArOperation structure represents the GetNtmsObjectAttributeAR operation
type xxx_GetNTMSObjectAttributeArOperation struct {
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	ObjectID            *dtyp.GUID     `idl:"name:lpObjectId" json:"object_id"`
	Type                uint32         `idl:"name:dwType" json:"type"`
	AttributeName       uint8          `idl:"name:lpAttributeName" json:"attribute_name"`
	AttributeData       []byte         `idl:"name:lpAttributeData;size_is:(lpdwAttributeBufferSize);length_is:(lpAttributeSize)" json:"attribute_data"`
	AttributeBufferSize uint32         `idl:"name:lpdwAttributeBufferSize" json:"attribute_buffer_size"`
	AttributeSize       uint32         `idl:"name:lpAttributeSize" json:"attribute_size"`
	ActualAttributeSize uint32         `idl:"name:lpActualAttributeSize" json:"actual_attribute_size"`
	Return              int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNTMSObjectAttributeArOperation) OpNum() int { return 14 }

func (o *xxx_GetNTMSObjectAttributeArOperation) OpName() string {
	return "/INtmsObjectManagement3/v0/GetNtmsObjectAttributeAR"
}

func (o *xxx_GetNTMSObjectAttributeArOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSObjectAttributeArOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpObjectId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID != nil {
			if err := o.ObjectID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// lpAttributeName {in} (1:{pointer=ref}*(1)(char))
	{
		if err := w.WriteData(o.AttributeName); err != nil {
			return err
		}
	}
	// lpdwAttributeBufferSize {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.AttributeBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSObjectAttributeArOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpObjectId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID == nil {
			o.ObjectID = &dtyp.GUID{}
		}
		if err := o.ObjectID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// lpAttributeName {in} (1:{pointer=ref}*(1)(char))
	{
		if err := w.ReadData(&o.AttributeName); err != nil {
			return err
		}
	}
	// lpdwAttributeBufferSize {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.AttributeBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSObjectAttributeArOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.AttributeData != nil && o.AttributeSize == 0 {
		o.AttributeSize = uint32(len(o.AttributeData))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSObjectAttributeArOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpAttributeData {out} (1:{pointer=ref}*(1)[dim:0,size_is=lpdwAttributeBufferSize,length_is=lpAttributeSize](uint8))
	{
		dimSize1 := uint64(o.AttributeBufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.AttributeSize)
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
		for i1 := range o.AttributeData {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.AttributeData[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.AttributeData); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// lpAttributeSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.AttributeSize); err != nil {
			return err
		}
	}
	// lpActualAttributeSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ActualAttributeSize); err != nil {
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

func (o *xxx_GetNTMSObjectAttributeArOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpAttributeData {out} (1:{pointer=ref}*(1)[dim:0,size_is=lpdwAttributeBufferSize,length_is=lpAttributeSize](uint8))
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
			return fmt.Errorf("buffer overflow for size %d of array o.AttributeData", sizeInfo[0])
		}
		o.AttributeData = make([]byte, sizeInfo[0])
		for i1 := range o.AttributeData {
			i1 := i1
			if err := w.ReadData(&o.AttributeData[i1]); err != nil {
				return err
			}
		}
	}
	// lpAttributeSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.AttributeSize); err != nil {
			return err
		}
	}
	// lpActualAttributeSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ActualAttributeSize); err != nil {
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

// GetNTMSObjectAttributeArRequest structure represents the GetNtmsObjectAttributeAR operation request
type GetNTMSObjectAttributeArRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	ObjectID            *dtyp.GUID     `idl:"name:lpObjectId" json:"object_id"`
	Type                uint32         `idl:"name:dwType" json:"type"`
	AttributeName       uint8          `idl:"name:lpAttributeName" json:"attribute_name"`
	AttributeBufferSize uint32         `idl:"name:lpdwAttributeBufferSize" json:"attribute_buffer_size"`
}

func (o *GetNTMSObjectAttributeArRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSObjectAttributeArOperation) *xxx_GetNTMSObjectAttributeArOperation {
	if op == nil {
		op = &xxx_GetNTMSObjectAttributeArOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ObjectID = o.ObjectID
	op.Type = o.Type
	op.AttributeName = o.AttributeName
	op.AttributeBufferSize = o.AttributeBufferSize
	return op
}

func (o *GetNTMSObjectAttributeArRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSObjectAttributeArOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectID = op.ObjectID
	o.Type = op.Type
	o.AttributeName = op.AttributeName
	o.AttributeBufferSize = op.AttributeBufferSize
}
func (o *GetNTMSObjectAttributeArRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNTMSObjectAttributeArRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSObjectAttributeArOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNTMSObjectAttributeArResponse structure represents the GetNtmsObjectAttributeAR operation response
type GetNTMSObjectAttributeArResponse struct {
	// XXX: lpdwAttributeBufferSize is an implicit input depedency for output parameters
	AttributeBufferSize uint32 `idl:"name:lpdwAttributeBufferSize" json:"attribute_buffer_size"`

	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	AttributeData       []byte         `idl:"name:lpAttributeData;size_is:(lpdwAttributeBufferSize);length_is:(lpAttributeSize)" json:"attribute_data"`
	AttributeSize       uint32         `idl:"name:lpAttributeSize" json:"attribute_size"`
	ActualAttributeSize uint32         `idl:"name:lpActualAttributeSize" json:"actual_attribute_size"`
	// Return: The GetNtmsObjectAttributeAR return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNTMSObjectAttributeArResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSObjectAttributeArOperation) *xxx_GetNTMSObjectAttributeArOperation {
	if op == nil {
		op = &xxx_GetNTMSObjectAttributeArOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.AttributeBufferSize == uint32(0) {
		op.AttributeBufferSize = o.AttributeBufferSize
	}

	op.That = o.That
	op.AttributeData = o.AttributeData
	op.AttributeSize = o.AttributeSize
	op.ActualAttributeSize = o.ActualAttributeSize
	op.Return = o.Return
	return op
}

func (o *GetNTMSObjectAttributeArResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSObjectAttributeArOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.AttributeBufferSize = op.AttributeBufferSize

	o.That = op.That
	o.AttributeData = op.AttributeData
	o.AttributeSize = op.AttributeSize
	o.ActualAttributeSize = op.ActualAttributeSize
	o.Return = op.Return
}
func (o *GetNTMSObjectAttributeArResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNTMSObjectAttributeArResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSObjectAttributeArOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNTMSObjectAttributeWrOperation structure represents the GetNtmsObjectAttributeWR operation
type xxx_GetNTMSObjectAttributeWrOperation struct {
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	ObjectID            *dtyp.GUID     `idl:"name:lpObjectId" json:"object_id"`
	Type                uint32         `idl:"name:dwType" json:"type"`
	AttributeName       string         `idl:"name:lpAttributeName;string" json:"attribute_name"`
	AttributeData       []byte         `idl:"name:lpAttributeData;size_is:(lpdwAttributeBufferSize);length_is:(lpAttributeSize)" json:"attribute_data"`
	AttributeBufferSize uint32         `idl:"name:lpdwAttributeBufferSize" json:"attribute_buffer_size"`
	AttributeSize       uint32         `idl:"name:lpAttributeSize" json:"attribute_size"`
	ActualAttributeSize uint32         `idl:"name:lpActualAttributeSize" json:"actual_attribute_size"`
	Return              int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNTMSObjectAttributeWrOperation) OpNum() int { return 15 }

func (o *xxx_GetNTMSObjectAttributeWrOperation) OpName() string {
	return "/INtmsObjectManagement3/v0/GetNtmsObjectAttributeWR"
}

func (o *xxx_GetNTMSObjectAttributeWrOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSObjectAttributeWrOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpObjectId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID != nil {
			if err := o.ObjectID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// lpAttributeName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.AttributeName); err != nil {
			return err
		}
	}
	// lpdwAttributeBufferSize {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.AttributeBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSObjectAttributeWrOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpObjectId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID == nil {
			o.ObjectID = &dtyp.GUID{}
		}
		if err := o.ObjectID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// lpAttributeName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.AttributeName); err != nil {
			return err
		}
	}
	// lpdwAttributeBufferSize {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.AttributeBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSObjectAttributeWrOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.AttributeData != nil && o.AttributeSize == 0 {
		o.AttributeSize = uint32(len(o.AttributeData))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSObjectAttributeWrOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpAttributeData {out} (1:{pointer=ref}*(1)[dim:0,size_is=lpdwAttributeBufferSize,length_is=lpAttributeSize](uint8))
	{
		dimSize1 := uint64(o.AttributeBufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.AttributeSize)
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
		for i1 := range o.AttributeData {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.AttributeData[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.AttributeData); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// lpAttributeSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.AttributeSize); err != nil {
			return err
		}
	}
	// lpActualAttributeSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ActualAttributeSize); err != nil {
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

func (o *xxx_GetNTMSObjectAttributeWrOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpAttributeData {out} (1:{pointer=ref}*(1)[dim:0,size_is=lpdwAttributeBufferSize,length_is=lpAttributeSize](uint8))
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
			return fmt.Errorf("buffer overflow for size %d of array o.AttributeData", sizeInfo[0])
		}
		o.AttributeData = make([]byte, sizeInfo[0])
		for i1 := range o.AttributeData {
			i1 := i1
			if err := w.ReadData(&o.AttributeData[i1]); err != nil {
				return err
			}
		}
	}
	// lpAttributeSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.AttributeSize); err != nil {
			return err
		}
	}
	// lpActualAttributeSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ActualAttributeSize); err != nil {
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

// GetNTMSObjectAttributeWrRequest structure represents the GetNtmsObjectAttributeWR operation request
type GetNTMSObjectAttributeWrRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	ObjectID            *dtyp.GUID     `idl:"name:lpObjectId" json:"object_id"`
	Type                uint32         `idl:"name:dwType" json:"type"`
	AttributeName       string         `idl:"name:lpAttributeName;string" json:"attribute_name"`
	AttributeBufferSize uint32         `idl:"name:lpdwAttributeBufferSize" json:"attribute_buffer_size"`
}

func (o *GetNTMSObjectAttributeWrRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSObjectAttributeWrOperation) *xxx_GetNTMSObjectAttributeWrOperation {
	if op == nil {
		op = &xxx_GetNTMSObjectAttributeWrOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ObjectID = o.ObjectID
	op.Type = o.Type
	op.AttributeName = o.AttributeName
	op.AttributeBufferSize = o.AttributeBufferSize
	return op
}

func (o *GetNTMSObjectAttributeWrRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSObjectAttributeWrOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectID = op.ObjectID
	o.Type = op.Type
	o.AttributeName = op.AttributeName
	o.AttributeBufferSize = op.AttributeBufferSize
}
func (o *GetNTMSObjectAttributeWrRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNTMSObjectAttributeWrRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSObjectAttributeWrOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNTMSObjectAttributeWrResponse structure represents the GetNtmsObjectAttributeWR operation response
type GetNTMSObjectAttributeWrResponse struct {
	// XXX: lpdwAttributeBufferSize is an implicit input depedency for output parameters
	AttributeBufferSize uint32 `idl:"name:lpdwAttributeBufferSize" json:"attribute_buffer_size"`

	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	AttributeData       []byte         `idl:"name:lpAttributeData;size_is:(lpdwAttributeBufferSize);length_is:(lpAttributeSize)" json:"attribute_data"`
	AttributeSize       uint32         `idl:"name:lpAttributeSize" json:"attribute_size"`
	ActualAttributeSize uint32         `idl:"name:lpActualAttributeSize" json:"actual_attribute_size"`
	// Return: The GetNtmsObjectAttributeWR return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNTMSObjectAttributeWrResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSObjectAttributeWrOperation) *xxx_GetNTMSObjectAttributeWrOperation {
	if op == nil {
		op = &xxx_GetNTMSObjectAttributeWrOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.AttributeBufferSize == uint32(0) {
		op.AttributeBufferSize = o.AttributeBufferSize
	}

	op.That = o.That
	op.AttributeData = o.AttributeData
	op.AttributeSize = o.AttributeSize
	op.ActualAttributeSize = o.ActualAttributeSize
	op.Return = o.Return
	return op
}

func (o *GetNTMSObjectAttributeWrResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSObjectAttributeWrOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.AttributeBufferSize = op.AttributeBufferSize

	o.That = op.That
	o.AttributeData = op.AttributeData
	o.AttributeSize = op.AttributeSize
	o.ActualAttributeSize = op.ActualAttributeSize
	o.Return = op.Return
}
func (o *GetNTMSObjectAttributeWrResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNTMSObjectAttributeWrResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSObjectAttributeWrOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
