package frsrpc

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
)

var (
	// import guard
	GoPackage = "frs1"
)

var (
	// Syntax UUID
	FrsrpcSyntaxUUID = &uuid.UUID{TimeLow: 0xf5cc59b4, TimeMid: 0x4264, TimeHiAndVersion: 0x101a, ClockSeqHiAndReserved: 0x8c, ClockSeqLow: 0x59, Node: [6]uint8{0x8, 0x0, 0x2b, 0x2f, 0x84, 0x26}}
	// Syntax ID
	FrsrpcSyntaxV1_1 = &dcerpc.SyntaxID{IfUUID: FrsrpcSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 1}
)

// frsrpc interface.
type FrsrpcClient interface {
	FrsRPCSendCommPacket(context.Context, *FrsRPCSendCommPacketRequest, ...dcerpc.CallOption) (*FrsRPCSendCommPacketResponse, error)

	FrsRPCVerifyPromotionParent(context.Context, *FrsRPCVerifyPromotionParentRequest, ...dcerpc.CallOption) (*FrsRPCVerifyPromotionParentResponse, error)

	FrsRPCStartPromotionParent(context.Context, *FrsRPCStartPromotionParentRequest, ...dcerpc.CallOption) (*FrsRPCStartPromotionParentResponse, error)

	FrsNop(context.Context, *FrsNopRequest, ...dcerpc.CallOption) (*FrsNopResponse, error)

	// Opnum4NotUsedOnWire operation.
	// Opnum4NotUsedOnWire

	// Opnum5NotUsedOnWire operation.
	// Opnum5NotUsedOnWire

	// Opnum6NotUsedOnWire operation.
	// Opnum6NotUsedOnWire

	// Opnum7NotUsedOnWire operation.
	// Opnum7NotUsedOnWire

	// Opnum8NotUsedOnWire operation.
	// Opnum8NotUsedOnWire

	// Opnum9NotUsedOnWire operation.
	// Opnum9NotUsedOnWire

	// Opnum10NotUsedOnWire operation.
	// Opnum10NotUsedOnWire

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// CommPacket structure represents COMM_PACKET RPC structure.
type CommPacket struct {
	Major        uint32 `idl:"name:Major" json:"major"`
	Minor        uint32 `idl:"name:Minor" json:"minor"`
	CSID         uint32 `idl:"name:CsId" json:"cs_id"`
	MemLength    uint32 `idl:"name:MemLen" json:"mem_length"`
	PacketLength uint32 `idl:"name:PktLen" json:"packet_length"`
	UpkLength    uint32 `idl:"name:UpkLen" json:"upk_length"`
	Packet       []byte `idl:"name:Pkt;size_is:(PktLen)" json:"packet"`
	DataName     []byte `idl:"name:DataName" json:"data_name"`
	DataHandle   []byte `idl:"name:DataHandle" json:"data_handle"`
}

func (o *CommPacket) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Packet != nil && o.PacketLength == 0 {
		o.PacketLength = uint32(len(o.Packet))
	}
	if o.PacketLength > uint32(262144) {
		return fmt.Errorf("PacketLength is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *CommPacket) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Major); err != nil {
		return err
	}
	if err := w.WriteData(o.Minor); err != nil {
		return err
	}
	if err := w.WriteData(o.CSID); err != nil {
		return err
	}
	if err := w.WriteData(o.MemLength); err != nil {
		return err
	}
	if err := w.WriteData(o.PacketLength); err != nil {
		return err
	}
	if err := w.WriteData(o.UpkLength); err != nil {
		return err
	}
	if o.Packet != nil || o.PacketLength > 0 {
		_ptr_Pkt := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.PacketLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Packet {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Packet[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Packet); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Packet, _ptr_Pkt); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *CommPacket) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Major); err != nil {
		return err
	}
	if err := w.ReadData(&o.Minor); err != nil {
		return err
	}
	if err := w.ReadData(&o.CSID); err != nil {
		return err
	}
	if err := w.ReadData(&o.MemLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.PacketLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.UpkLength); err != nil {
		return err
	}
	_ptr_Pkt := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.PacketLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.PacketLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Packet", sizeInfo[0])
		}
		o.Packet = make([]byte, sizeInfo[0])
		for i1 := range o.Packet {
			i1 := i1
			if err := w.ReadData(&o.Packet[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Pkt := func(ptr interface{}) { o.Packet = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Packet, _s_Pkt, _ptr_Pkt); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultFrsrpcClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultFrsrpcClient) FrsRPCSendCommPacket(ctx context.Context, in *FrsRPCSendCommPacketRequest, opts ...dcerpc.CallOption) (*FrsRPCSendCommPacketResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FrsRPCSendCommPacketResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFrsrpcClient) FrsRPCVerifyPromotionParent(ctx context.Context, in *FrsRPCVerifyPromotionParentRequest, opts ...dcerpc.CallOption) (*FrsRPCVerifyPromotionParentResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FrsRPCVerifyPromotionParentResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFrsrpcClient) FrsRPCStartPromotionParent(ctx context.Context, in *FrsRPCStartPromotionParentRequest, opts ...dcerpc.CallOption) (*FrsRPCStartPromotionParentResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FrsRPCStartPromotionParentResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFrsrpcClient) FrsNop(ctx context.Context, in *FrsNopRequest, opts ...dcerpc.CallOption) (*FrsNopResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FrsNopResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFrsrpcClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultFrsrpcClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewFrsrpcClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (FrsrpcClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(FrsrpcSyntaxV1_1))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultFrsrpcClient{cc: cc}, nil
}

// xxx_FrsRPCSendCommPacketOperation structure represents the FrsRpcSendCommPkt operation
type xxx_FrsRPCSendCommPacketOperation struct {
	CommPacket *CommPacket `idl:"name:CommPkt" json:"comm_packet"`
	Return     uint32      `idl:"name:Return" json:"return"`
}

func (o *xxx_FrsRPCSendCommPacketOperation) OpNum() int { return 0 }

func (o *xxx_FrsRPCSendCommPacketOperation) OpName() string { return "/frsrpc/v1.1/FrsRpcSendCommPkt" }

func (o *xxx_FrsRPCSendCommPacketOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FrsRPCSendCommPacketOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// CommPkt {in} (1:{alias=PCOMM_PACKET}*(1))(2:{alias=COMM_PACKET}(struct))
	{
		if o.CommPacket != nil {
			if err := o.CommPacket.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&CommPacket{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FrsRPCSendCommPacketOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// CommPkt {in} (1:{alias=PCOMM_PACKET,pointer=ref}*(1))(2:{alias=COMM_PACKET}(struct))
	{
		if o.CommPacket == nil {
			o.CommPacket = &CommPacket{}
		}
		if err := o.CommPacket.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FrsRPCSendCommPacketOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FrsRPCSendCommPacketOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_FrsRPCSendCommPacketOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// FrsRPCSendCommPacketRequest structure represents the FrsRpcSendCommPkt operation request
type FrsRPCSendCommPacketRequest struct {
	CommPacket *CommPacket `idl:"name:CommPkt" json:"comm_packet"`
}

func (o *FrsRPCSendCommPacketRequest) xxx_ToOp(ctx context.Context, op *xxx_FrsRPCSendCommPacketOperation) *xxx_FrsRPCSendCommPacketOperation {
	if op == nil {
		op = &xxx_FrsRPCSendCommPacketOperation{}
	}
	if o == nil {
		return op
	}
	op.CommPacket = o.CommPacket
	return op
}

func (o *FrsRPCSendCommPacketRequest) xxx_FromOp(ctx context.Context, op *xxx_FrsRPCSendCommPacketOperation) {
	if o == nil {
		return
	}
	o.CommPacket = op.CommPacket
}
func (o *FrsRPCSendCommPacketRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *FrsRPCSendCommPacketRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FrsRPCSendCommPacketOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FrsRPCSendCommPacketResponse structure represents the FrsRpcSendCommPkt operation response
type FrsRPCSendCommPacketResponse struct {
	// Return: The FrsRpcSendCommPkt return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *FrsRPCSendCommPacketResponse) xxx_ToOp(ctx context.Context, op *xxx_FrsRPCSendCommPacketOperation) *xxx_FrsRPCSendCommPacketOperation {
	if op == nil {
		op = &xxx_FrsRPCSendCommPacketOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *FrsRPCSendCommPacketResponse) xxx_FromOp(ctx context.Context, op *xxx_FrsRPCSendCommPacketOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *FrsRPCSendCommPacketResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *FrsRPCSendCommPacketResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FrsRPCSendCommPacketOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FrsRPCVerifyPromotionParentOperation structure represents the FrsRpcVerifyPromotionParent operation
type xxx_FrsRPCVerifyPromotionParentOperation struct {
	ParentAccount    string `idl:"name:ParentAccount;string;pointer:unique" json:"parent_account"`
	ParentPassword   string `idl:"name:ParentPassword;string;pointer:unique" json:"parent_password"`
	SetNameReplica   string `idl:"name:ReplicaSetName;string;pointer:unique" json:"set_name_replica"`
	SetTypeReplica   string `idl:"name:ReplicaSetType;string;pointer:unique" json:"set_type_replica"`
	PartnerAuthLevel uint32 `idl:"name:PartnerAuthLevel" json:"partner_auth_level"`
	GUIDSize         uint32 `idl:"name:GuidSize" json:"guid_size"`
	Return           uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_FrsRPCVerifyPromotionParentOperation) OpNum() int { return 1 }

func (o *xxx_FrsRPCVerifyPromotionParentOperation) OpName() string {
	return "/frsrpc/v1.1/FrsRpcVerifyPromotionParent"
}

func (o *xxx_FrsRPCVerifyPromotionParentOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FrsRPCVerifyPromotionParentOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ParentAccount {in} (1:{string, pointer=unique, alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.ParentAccount != "" {
			_ptr_ParentAccount := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ParentAccount); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ParentAccount, _ptr_ParentAccount); err != nil {
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
	// ParentPassword {in} (1:{string, pointer=unique, alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.ParentPassword != "" {
			_ptr_ParentPassword := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ParentPassword); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ParentPassword, _ptr_ParentPassword); err != nil {
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
	// ReplicaSetType {in} (1:{string, pointer=unique, alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.SetTypeReplica != "" {
			_ptr_ReplicaSetType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.SetTypeReplica); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.SetTypeReplica, _ptr_ReplicaSetType); err != nil {
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
	// PartnerAuthLevel {in} (1:(uint32))
	{
		if err := w.WriteData(o.PartnerAuthLevel); err != nil {
			return err
		}
	}
	// GuidSize {in} (1:(uint32))
	{
		if err := w.WriteData(o.GUIDSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FrsRPCVerifyPromotionParentOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ParentAccount {in} (1:{string, pointer=unique, alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_ParentAccount := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ParentAccount); err != nil {
				return err
			}
			return nil
		})
		_s_ParentAccount := func(ptr interface{}) { o.ParentAccount = *ptr.(*string) }
		if err := w.ReadPointer(&o.ParentAccount, _s_ParentAccount, _ptr_ParentAccount); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ParentPassword {in} (1:{string, pointer=unique, alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_ParentPassword := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ParentPassword); err != nil {
				return err
			}
			return nil
		})
		_s_ParentPassword := func(ptr interface{}) { o.ParentPassword = *ptr.(*string) }
		if err := w.ReadPointer(&o.ParentPassword, _s_ParentPassword, _ptr_ParentPassword); err != nil {
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
	// ReplicaSetType {in} (1:{string, pointer=unique, alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_ReplicaSetType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.SetTypeReplica); err != nil {
				return err
			}
			return nil
		})
		_s_ReplicaSetType := func(ptr interface{}) { o.SetTypeReplica = *ptr.(*string) }
		if err := w.ReadPointer(&o.SetTypeReplica, _s_ReplicaSetType, _ptr_ReplicaSetType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// PartnerAuthLevel {in} (1:(uint32))
	{
		if err := w.ReadData(&o.PartnerAuthLevel); err != nil {
			return err
		}
	}
	// GuidSize {in} (1:(uint32))
	{
		if err := w.ReadData(&o.GUIDSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FrsRPCVerifyPromotionParentOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FrsRPCVerifyPromotionParentOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_FrsRPCVerifyPromotionParentOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// FrsRPCVerifyPromotionParentRequest structure represents the FrsRpcVerifyPromotionParent operation request
type FrsRPCVerifyPromotionParentRequest struct {
	ParentAccount    string `idl:"name:ParentAccount;string;pointer:unique" json:"parent_account"`
	ParentPassword   string `idl:"name:ParentPassword;string;pointer:unique" json:"parent_password"`
	SetNameReplica   string `idl:"name:ReplicaSetName;string;pointer:unique" json:"set_name_replica"`
	SetTypeReplica   string `idl:"name:ReplicaSetType;string;pointer:unique" json:"set_type_replica"`
	PartnerAuthLevel uint32 `idl:"name:PartnerAuthLevel" json:"partner_auth_level"`
	GUIDSize         uint32 `idl:"name:GuidSize" json:"guid_size"`
}

func (o *FrsRPCVerifyPromotionParentRequest) xxx_ToOp(ctx context.Context, op *xxx_FrsRPCVerifyPromotionParentOperation) *xxx_FrsRPCVerifyPromotionParentOperation {
	if op == nil {
		op = &xxx_FrsRPCVerifyPromotionParentOperation{}
	}
	if o == nil {
		return op
	}
	op.ParentAccount = o.ParentAccount
	op.ParentPassword = o.ParentPassword
	op.SetNameReplica = o.SetNameReplica
	op.SetTypeReplica = o.SetTypeReplica
	op.PartnerAuthLevel = o.PartnerAuthLevel
	op.GUIDSize = o.GUIDSize
	return op
}

func (o *FrsRPCVerifyPromotionParentRequest) xxx_FromOp(ctx context.Context, op *xxx_FrsRPCVerifyPromotionParentOperation) {
	if o == nil {
		return
	}
	o.ParentAccount = op.ParentAccount
	o.ParentPassword = op.ParentPassword
	o.SetNameReplica = op.SetNameReplica
	o.SetTypeReplica = op.SetTypeReplica
	o.PartnerAuthLevel = op.PartnerAuthLevel
	o.GUIDSize = op.GUIDSize
}
func (o *FrsRPCVerifyPromotionParentRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *FrsRPCVerifyPromotionParentRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FrsRPCVerifyPromotionParentOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FrsRPCVerifyPromotionParentResponse structure represents the FrsRpcVerifyPromotionParent operation response
type FrsRPCVerifyPromotionParentResponse struct {
	// Return: The FrsRpcVerifyPromotionParent return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *FrsRPCVerifyPromotionParentResponse) xxx_ToOp(ctx context.Context, op *xxx_FrsRPCVerifyPromotionParentOperation) *xxx_FrsRPCVerifyPromotionParentOperation {
	if op == nil {
		op = &xxx_FrsRPCVerifyPromotionParentOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *FrsRPCVerifyPromotionParentResponse) xxx_FromOp(ctx context.Context, op *xxx_FrsRPCVerifyPromotionParentOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *FrsRPCVerifyPromotionParentResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *FrsRPCVerifyPromotionParentResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FrsRPCVerifyPromotionParentOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FrsRPCStartPromotionParentOperation structure represents the FrsRpcStartPromotionParent operation
type xxx_FrsRPCStartPromotionParentOperation struct {
	ParentAccount    string `idl:"name:ParentAccount;string;pointer:unique" json:"parent_account"`
	ParentPassword   string `idl:"name:ParentPassword;string;pointer:unique" json:"parent_password"`
	SetNameReplica   string `idl:"name:ReplicaSetName;string;pointer:unique" json:"set_name_replica"`
	SetTypeReplica   string `idl:"name:ReplicaSetType;string;pointer:unique" json:"set_type_replica"`
	CxtionName       string `idl:"name:CxtionName;string;pointer:unique" json:"cxtion_name"`
	PartnerName      string `idl:"name:PartnerName;string;pointer:unique" json:"partner_name"`
	PartnerPrincName string `idl:"name:PartnerPrincName;string;pointer:unique" json:"partner_princ_name"`
	PartnerAuthLevel uint32 `idl:"name:PartnerAuthLevel" json:"partner_auth_level"`
	GUIDSize         uint32 `idl:"name:GuidSize" json:"guid_size"`
	CxtionGUID       []byte `idl:"name:CxtionGuid;size_is:(GuidSize);pointer:unique" json:"cxtion_guid"`
	PartnerGUID      []byte `idl:"name:PartnerGuid;size_is:(GuidSize);pointer:unique" json:"partner_guid"`
	ParentGUID       []byte `idl:"name:ParentGuid;size_is:(GuidSize);pointer:unique" json:"parent_guid"`
	Return           uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_FrsRPCStartPromotionParentOperation) OpNum() int { return 2 }

func (o *xxx_FrsRPCStartPromotionParentOperation) OpName() string {
	return "/frsrpc/v1.1/FrsRpcStartPromotionParent"
}

func (o *xxx_FrsRPCStartPromotionParentOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.CxtionGUID != nil && o.GUIDSize == 0 {
		o.GUIDSize = uint32(len(o.CxtionGUID))
	}
	if o.PartnerGUID != nil && o.GUIDSize == 0 {
		o.GUIDSize = uint32(len(o.PartnerGUID))
	}
	if o.ParentGUID != nil && o.GUIDSize == 0 {
		o.GUIDSize = uint32(len(o.ParentGUID))
	}
	if o.GUIDSize < uint32(15) || o.GUIDSize > uint32(15) {
		return fmt.Errorf("GUIDSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FrsRPCStartPromotionParentOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ParentAccount {in} (1:{string, pointer=unique, alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.ParentAccount != "" {
			_ptr_ParentAccount := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ParentAccount); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ParentAccount, _ptr_ParentAccount); err != nil {
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
	// ParentPassword {in} (1:{string, pointer=unique, alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.ParentPassword != "" {
			_ptr_ParentPassword := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ParentPassword); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ParentPassword, _ptr_ParentPassword); err != nil {
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
	// ReplicaSetType {in} (1:{string, pointer=unique, alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.SetTypeReplica != "" {
			_ptr_ReplicaSetType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.SetTypeReplica); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.SetTypeReplica, _ptr_ReplicaSetType); err != nil {
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
	// CxtionName {in} (1:{string, pointer=unique, alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.CxtionName != "" {
			_ptr_CxtionName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.CxtionName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.CxtionName, _ptr_CxtionName); err != nil {
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
	// PartnerName {in} (1:{string, pointer=unique, alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.PartnerName != "" {
			_ptr_PartnerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.PartnerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.PartnerName, _ptr_PartnerName); err != nil {
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
	// PartnerPrincName {in} (1:{string, pointer=unique, alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.PartnerPrincName != "" {
			_ptr_PartnerPrincName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.PartnerPrincName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.PartnerPrincName, _ptr_PartnerPrincName); err != nil {
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
	// PartnerAuthLevel {in} (1:(uint32))
	{
		if err := w.WriteData(o.PartnerAuthLevel); err != nil {
			return err
		}
	}
	// GuidSize {in} (1:{range=(15,15)}(uint32))
	{
		if err := w.WriteData(o.GUIDSize); err != nil {
			return err
		}
	}
	// CxtionGuid {in} (1:{pointer=unique}*(1)[dim:0,size_is=GuidSize](uchar))
	{
		if o.CxtionGUID != nil || o.GUIDSize > 0 {
			_ptr_CxtionGuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.GUIDSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.CxtionGUID {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.CxtionGUID[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.CxtionGUID); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
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
	// PartnerGuid {in} (1:{pointer=unique}*(1)[dim:0,size_is=GuidSize](uchar))
	{
		if o.PartnerGUID != nil || o.GUIDSize > 0 {
			_ptr_PartnerGuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.GUIDSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.PartnerGUID {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.PartnerGUID[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.PartnerGUID); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PartnerGUID, _ptr_PartnerGuid); err != nil {
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
	// ParentGuid {in, out} (1:{pointer=unique}*(1)[dim:0,size_is=GuidSize](uchar))
	{
		if o.ParentGUID != nil || o.GUIDSize > 0 {
			_ptr_ParentGuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.GUIDSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ParentGUID {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.ParentGUID[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.ParentGUID); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ParentGUID, _ptr_ParentGuid); err != nil {
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

func (o *xxx_FrsRPCStartPromotionParentOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ParentAccount {in} (1:{string, pointer=unique, alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_ParentAccount := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ParentAccount); err != nil {
				return err
			}
			return nil
		})
		_s_ParentAccount := func(ptr interface{}) { o.ParentAccount = *ptr.(*string) }
		if err := w.ReadPointer(&o.ParentAccount, _s_ParentAccount, _ptr_ParentAccount); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ParentPassword {in} (1:{string, pointer=unique, alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_ParentPassword := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ParentPassword); err != nil {
				return err
			}
			return nil
		})
		_s_ParentPassword := func(ptr interface{}) { o.ParentPassword = *ptr.(*string) }
		if err := w.ReadPointer(&o.ParentPassword, _s_ParentPassword, _ptr_ParentPassword); err != nil {
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
	// ReplicaSetType {in} (1:{string, pointer=unique, alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_ReplicaSetType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.SetTypeReplica); err != nil {
				return err
			}
			return nil
		})
		_s_ReplicaSetType := func(ptr interface{}) { o.SetTypeReplica = *ptr.(*string) }
		if err := w.ReadPointer(&o.SetTypeReplica, _s_ReplicaSetType, _ptr_ReplicaSetType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// CxtionName {in} (1:{string, pointer=unique, alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_CxtionName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.CxtionName); err != nil {
				return err
			}
			return nil
		})
		_s_CxtionName := func(ptr interface{}) { o.CxtionName = *ptr.(*string) }
		if err := w.ReadPointer(&o.CxtionName, _s_CxtionName, _ptr_CxtionName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// PartnerName {in} (1:{string, pointer=unique, alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_PartnerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.PartnerName); err != nil {
				return err
			}
			return nil
		})
		_s_PartnerName := func(ptr interface{}) { o.PartnerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.PartnerName, _s_PartnerName, _ptr_PartnerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// PartnerPrincName {in} (1:{string, pointer=unique, alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_PartnerPrincName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.PartnerPrincName); err != nil {
				return err
			}
			return nil
		})
		_s_PartnerPrincName := func(ptr interface{}) { o.PartnerPrincName = *ptr.(*string) }
		if err := w.ReadPointer(&o.PartnerPrincName, _s_PartnerPrincName, _ptr_PartnerPrincName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// PartnerAuthLevel {in} (1:(uint32))
	{
		if err := w.ReadData(&o.PartnerAuthLevel); err != nil {
			return err
		}
	}
	// GuidSize {in} (1:{range=(15,15)}(uint32))
	{
		if err := w.ReadData(&o.GUIDSize); err != nil {
			return err
		}
	}
	// CxtionGuid {in} (1:{pointer=unique}*(1)[dim:0,size_is=GuidSize](uchar))
	{
		_ptr_CxtionGuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.CxtionGUID", sizeInfo[0])
			}
			o.CxtionGUID = make([]byte, sizeInfo[0])
			for i1 := range o.CxtionGUID {
				i1 := i1
				if err := w.ReadData(&o.CxtionGUID[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_CxtionGuid := func(ptr interface{}) { o.CxtionGUID = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.CxtionGUID, _s_CxtionGuid, _ptr_CxtionGuid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// PartnerGuid {in} (1:{pointer=unique}*(1)[dim:0,size_is=GuidSize](uchar))
	{
		_ptr_PartnerGuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.PartnerGUID", sizeInfo[0])
			}
			o.PartnerGUID = make([]byte, sizeInfo[0])
			for i1 := range o.PartnerGUID {
				i1 := i1
				if err := w.ReadData(&o.PartnerGUID[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_PartnerGuid := func(ptr interface{}) { o.PartnerGUID = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.PartnerGUID, _s_PartnerGuid, _ptr_PartnerGuid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ParentGuid {in, out} (1:{pointer=unique}*(1)[dim:0,size_is=GuidSize](uchar))
	{
		_ptr_ParentGuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ParentGUID", sizeInfo[0])
			}
			o.ParentGUID = make([]byte, sizeInfo[0])
			for i1 := range o.ParentGUID {
				i1 := i1
				if err := w.ReadData(&o.ParentGUID[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ParentGuid := func(ptr interface{}) { o.ParentGUID = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.ParentGUID, _s_ParentGuid, _ptr_ParentGuid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FrsRPCStartPromotionParentOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FrsRPCStartPromotionParentOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ParentGuid {in, out} (1:{pointer=unique}*(1)[dim:0,size_is=GuidSize](uchar))
	{
		if o.ParentGUID != nil || o.GUIDSize > 0 {
			_ptr_ParentGuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.GUIDSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ParentGUID {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.ParentGUID[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.ParentGUID); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ParentGUID, _ptr_ParentGuid); err != nil {
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

func (o *xxx_FrsRPCStartPromotionParentOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ParentGuid {in, out} (1:{pointer=unique}*(1)[dim:0,size_is=GuidSize](uchar))
	{
		_ptr_ParentGuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ParentGUID", sizeInfo[0])
			}
			o.ParentGUID = make([]byte, sizeInfo[0])
			for i1 := range o.ParentGUID {
				i1 := i1
				if err := w.ReadData(&o.ParentGUID[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ParentGuid := func(ptr interface{}) { o.ParentGUID = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.ParentGUID, _s_ParentGuid, _ptr_ParentGuid); err != nil {
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

// FrsRPCStartPromotionParentRequest structure represents the FrsRpcStartPromotionParent operation request
type FrsRPCStartPromotionParentRequest struct {
	ParentAccount    string `idl:"name:ParentAccount;string;pointer:unique" json:"parent_account"`
	ParentPassword   string `idl:"name:ParentPassword;string;pointer:unique" json:"parent_password"`
	SetNameReplica   string `idl:"name:ReplicaSetName;string;pointer:unique" json:"set_name_replica"`
	SetTypeReplica   string `idl:"name:ReplicaSetType;string;pointer:unique" json:"set_type_replica"`
	CxtionName       string `idl:"name:CxtionName;string;pointer:unique" json:"cxtion_name"`
	PartnerName      string `idl:"name:PartnerName;string;pointer:unique" json:"partner_name"`
	PartnerPrincName string `idl:"name:PartnerPrincName;string;pointer:unique" json:"partner_princ_name"`
	PartnerAuthLevel uint32 `idl:"name:PartnerAuthLevel" json:"partner_auth_level"`
	GUIDSize         uint32 `idl:"name:GuidSize" json:"guid_size"`
	CxtionGUID       []byte `idl:"name:CxtionGuid;size_is:(GuidSize);pointer:unique" json:"cxtion_guid"`
	PartnerGUID      []byte `idl:"name:PartnerGuid;size_is:(GuidSize);pointer:unique" json:"partner_guid"`
	ParentGUID       []byte `idl:"name:ParentGuid;size_is:(GuidSize);pointer:unique" json:"parent_guid"`
}

func (o *FrsRPCStartPromotionParentRequest) xxx_ToOp(ctx context.Context, op *xxx_FrsRPCStartPromotionParentOperation) *xxx_FrsRPCStartPromotionParentOperation {
	if op == nil {
		op = &xxx_FrsRPCStartPromotionParentOperation{}
	}
	if o == nil {
		return op
	}
	op.ParentAccount = o.ParentAccount
	op.ParentPassword = o.ParentPassword
	op.SetNameReplica = o.SetNameReplica
	op.SetTypeReplica = o.SetTypeReplica
	op.CxtionName = o.CxtionName
	op.PartnerName = o.PartnerName
	op.PartnerPrincName = o.PartnerPrincName
	op.PartnerAuthLevel = o.PartnerAuthLevel
	op.GUIDSize = o.GUIDSize
	op.CxtionGUID = o.CxtionGUID
	op.PartnerGUID = o.PartnerGUID
	op.ParentGUID = o.ParentGUID
	return op
}

func (o *FrsRPCStartPromotionParentRequest) xxx_FromOp(ctx context.Context, op *xxx_FrsRPCStartPromotionParentOperation) {
	if o == nil {
		return
	}
	o.ParentAccount = op.ParentAccount
	o.ParentPassword = op.ParentPassword
	o.SetNameReplica = op.SetNameReplica
	o.SetTypeReplica = op.SetTypeReplica
	o.CxtionName = op.CxtionName
	o.PartnerName = op.PartnerName
	o.PartnerPrincName = op.PartnerPrincName
	o.PartnerAuthLevel = op.PartnerAuthLevel
	o.GUIDSize = op.GUIDSize
	o.CxtionGUID = op.CxtionGUID
	o.PartnerGUID = op.PartnerGUID
	o.ParentGUID = op.ParentGUID
}
func (o *FrsRPCStartPromotionParentRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *FrsRPCStartPromotionParentRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FrsRPCStartPromotionParentOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FrsRPCStartPromotionParentResponse structure represents the FrsRpcStartPromotionParent operation response
type FrsRPCStartPromotionParentResponse struct {
	// XXX: GuidSize is an implicit input depedency for output parameters
	GUIDSize uint32 `idl:"name:GuidSize" json:"guid_size"`

	ParentGUID []byte `idl:"name:ParentGuid;size_is:(GuidSize);pointer:unique" json:"parent_guid"`
	// Return: The FrsRpcStartPromotionParent return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *FrsRPCStartPromotionParentResponse) xxx_ToOp(ctx context.Context, op *xxx_FrsRPCStartPromotionParentOperation) *xxx_FrsRPCStartPromotionParentOperation {
	if op == nil {
		op = &xxx_FrsRPCStartPromotionParentOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.GUIDSize == uint32(0) {
		op.GUIDSize = o.GUIDSize
	}

	op.ParentGUID = o.ParentGUID
	op.Return = o.Return
	return op
}

func (o *FrsRPCStartPromotionParentResponse) xxx_FromOp(ctx context.Context, op *xxx_FrsRPCStartPromotionParentOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.GUIDSize = op.GUIDSize

	o.ParentGUID = op.ParentGUID
	o.Return = op.Return
}
func (o *FrsRPCStartPromotionParentResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *FrsRPCStartPromotionParentResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FrsRPCStartPromotionParentOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FrsNopOperation structure represents the FrsNOP operation
type xxx_FrsNopOperation struct {
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_FrsNopOperation) OpNum() int { return 3 }

func (o *xxx_FrsNopOperation) OpName() string { return "/frsrpc/v1.1/FrsNOP" }

func (o *xxx_FrsNopOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FrsNopOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_FrsNopOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_FrsNopOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FrsNopOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_FrsNopOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// FrsNopRequest structure represents the FrsNOP operation request
type FrsNopRequest struct {
}

func (o *FrsNopRequest) xxx_ToOp(ctx context.Context, op *xxx_FrsNopOperation) *xxx_FrsNopOperation {
	if op == nil {
		op = &xxx_FrsNopOperation{}
	}
	if o == nil {
		return op
	}
	return op
}

func (o *FrsNopRequest) xxx_FromOp(ctx context.Context, op *xxx_FrsNopOperation) {
	if o == nil {
		return
	}
}
func (o *FrsNopRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *FrsNopRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FrsNopOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FrsNopResponse structure represents the FrsNOP operation response
type FrsNopResponse struct {
	// Return: The FrsNOP return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *FrsNopResponse) xxx_ToOp(ctx context.Context, op *xxx_FrsNopOperation) *xxx_FrsNopOperation {
	if op == nil {
		op = &xxx_FrsNopOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *FrsNopResponse) xxx_FromOp(ctx context.Context, op *xxx_FrsNopOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *FrsNopResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *FrsNopResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FrsNopOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
