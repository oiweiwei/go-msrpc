package wdsrpc

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
)

var (
	// import guard
	GoPackage = "wdsc"
)

var (
	// Syntax UUID
	WdsRPCSyntaxUUID = &uuid.UUID{TimeLow: 0x1a927394, TimeMid: 0x352e, TimeHiAndVersion: 0x4553, ClockSeqHiAndReserved: 0xae, ClockSeqLow: 0x3f, Node: [6]uint8{0x7c, 0xf4, 0xaa, 0xfc, 0xa6, 0x20}}
	// Syntax ID
	WdsRPCSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: WdsRPCSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// WdsRpc interface.
type WdsRPCClient interface {

	// The WdsRpcMessage (opnum 0) method sends the request packet to the server and returns
	// the corresponding reply packet.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success or a
	// non-zero Win32 error code value if an error occurred.
	Message(context.Context, *MessageRequest, ...dcerpc.CallOption) (*MessageResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultWdsRPCClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultWdsRPCClient) Message(ctx context.Context, in *MessageRequest, opts ...dcerpc.CallOption) (*MessageResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MessageResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWdsRPCClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultWdsRPCClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewWdsRPCClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (WdsRPCClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(WdsRPCSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultWdsRPCClient{cc: cc}, nil
}

// xxx_MessageOperation structure represents the WdsRpcMessage operation
type xxx_MessageOperation struct {
	RequestPacketSize uint32 `idl:"name:uRequestPacketSize" json:"request_packet_size"`
	RequestPacket     []byte `idl:"name:bRequestPacket;size_is:(uRequestPacketSize)" json:"request_packet"`
	ReplyPacketSize   uint32 `idl:"name:puReplyPacketSize" json:"reply_packet_size"`
	ReplyPacket       []byte `idl:"name:pbReplyPacket;size_is:(, puReplyPacketSize)" json:"reply_packet"`
	Return            uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_MessageOperation) OpNum() int { return 0 }

func (o *xxx_MessageOperation) OpName() string { return "/WdsRpc/v1/WdsRpcMessage" }

func (o *xxx_MessageOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.RequestPacket != nil && o.RequestPacketSize == 0 {
		o.RequestPacketSize = uint32(len(o.RequestPacket))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// uRequestPacketSize {in} (1:(uint32))
	{
		if err := w.WriteData(o.RequestPacketSize); err != nil {
			return err
		}
	}
	// bRequestPacket {in} (1:[dim:0,size_is=uRequestPacketSize](uint8))
	{
		dimSize1 := uint64(o.RequestPacketSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.RequestPacket {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.RequestPacket[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.RequestPacket); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_MessageOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// uRequestPacketSize {in} (1:(uint32))
	{
		if err := w.ReadData(&o.RequestPacketSize); err != nil {
			return err
		}
	}
	// bRequestPacket {in} (1:[dim:0,size_is=uRequestPacketSize](uint8))
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
			return fmt.Errorf("buffer overflow for size %d of array o.RequestPacket", sizeInfo[0])
		}
		o.RequestPacket = make([]byte, sizeInfo[0])
		for i1 := range o.RequestPacket {
			i1 := i1
			if err := w.ReadData(&o.RequestPacket[i1]); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_MessageOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ReplyPacket != nil && o.ReplyPacketSize == 0 {
		o.ReplyPacketSize = uint32(len(o.ReplyPacket))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// puReplyPacketSize {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.ReplyPacketSize); err != nil {
			return err
		}
	}
	// pbReplyPacket {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=puReplyPacketSize](uint8))
	{
		if o.ReplyPacket != nil || o.ReplyPacketSize > 0 {
			_ptr_pbReplyPacket := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ReplyPacketSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ReplyPacket {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.ReplyPacket[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.ReplyPacket); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReplyPacket, _ptr_pbReplyPacket); err != nil {
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

func (o *xxx_MessageOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// puReplyPacketSize {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.ReplyPacketSize); err != nil {
			return err
		}
	}
	// pbReplyPacket {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=puReplyPacketSize](uint8))
	{
		_ptr_pbReplyPacket := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ReplyPacket", sizeInfo[0])
			}
			o.ReplyPacket = make([]byte, sizeInfo[0])
			for i1 := range o.ReplyPacket {
				i1 := i1
				if err := w.ReadData(&o.ReplyPacket[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pbReplyPacket := func(ptr interface{}) { o.ReplyPacket = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.ReplyPacket, _s_pbReplyPacket, _ptr_pbReplyPacket); err != nil {
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

// MessageRequest structure represents the WdsRpcMessage operation request
type MessageRequest struct {
	// uRequestPacketSize: The client MUST pass the total size of request packet in bytes.
	RequestPacketSize uint32 `idl:"name:uRequestPacketSize" json:"request_packet_size"`
	// bRequestPacket: A pointer to the request packet. The packet MUST be constructed as
	// specified in section 2.2.1.
	RequestPacket []byte `idl:"name:bRequestPacket;size_is:(uRequestPacketSize)" json:"request_packet"`
}

func (o *MessageRequest) xxx_ToOp(ctx context.Context, op *xxx_MessageOperation) *xxx_MessageOperation {
	if op == nil {
		op = &xxx_MessageOperation{}
	}
	if o == nil {
		return op
	}
	op.RequestPacketSize = o.RequestPacketSize
	op.RequestPacket = o.RequestPacket
	return op
}

func (o *MessageRequest) xxx_FromOp(ctx context.Context, op *xxx_MessageOperation) {
	if o == nil {
		return
	}
	o.RequestPacketSize = op.RequestPacketSize
	o.RequestPacket = op.RequestPacket
}
func (o *MessageRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MessageRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MessageOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MessageResponse structure represents the WdsRpcMessage operation response
type MessageResponse struct {
	// puReplyPacketSize: The WDS Server MUST set this to the total size of the reply packet
	// in bytes.
	ReplyPacketSize uint32 `idl:"name:puReplyPacketSize" json:"reply_packet_size"`
	// pbReplyPacket: The WDS Server MUST set this to the reply packet. The reply packet
	// MUST be constructed as specified in section 2.2.1.
	ReplyPacket []byte `idl:"name:pbReplyPacket;size_is:(, puReplyPacketSize)" json:"reply_packet"`
	// Return: The WdsRpcMessage return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MessageResponse) xxx_ToOp(ctx context.Context, op *xxx_MessageOperation) *xxx_MessageOperation {
	if op == nil {
		op = &xxx_MessageOperation{}
	}
	if o == nil {
		return op
	}
	op.ReplyPacketSize = o.ReplyPacketSize
	op.ReplyPacket = o.ReplyPacket
	op.Return = o.Return
	return op
}

func (o *MessageResponse) xxx_FromOp(ctx context.Context, op *xxx_MessageOperation) {
	if o == nil {
		return
	}
	o.ReplyPacketSize = op.ReplyPacketSize
	o.ReplyPacket = op.ReplyPacket
	o.Return = op.Return
}
func (o *MessageResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MessageResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MessageOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
