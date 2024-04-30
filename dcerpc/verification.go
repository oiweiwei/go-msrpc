package dcerpc

import (
	"bytes"
	"context"
	"fmt"

	"github.com/oiweiwei/go-msrpc/ndr"
)

var VerificationTrailerSignature = [...]byte{0x8A, 0xE3, 0x13, 0x71, 0x02, 0xF4, 0x36, 0x71}

// The bits field is a bitmask. A server MUST ignore bits it does not understand.
// Currently, there is only one bit defined: CLIENT_SUPPORT_HEADER_SIGNING
// (bitmask of 0x00000001). If this bit is set, the PFC_SUPPORT_HEADER_SIGN bit,
// MUST be present in the PDU header for the bind PDU on this connection.
type VerifyBitMask uint32

var (
	// If this bit is set, the PFC_SUPPORT_HEADER_SIGN bit, MUST be present
	// in the PDU header for the bind PDU on this connection.
	VerifyBitMaskSupportHeaderSign VerifyBitMask = 0x00000001
)

func (o VerifyBitMask) WriteTo(ctx context.Context, w ndr.Writer) error {
	return w.WriteData((uint32)(o))
}

func (o *VerifyBitMask) ReadFrom(ctx context.Context, r ndr.Reader) error {
	return r.ReadData((*uint32)(o))
}

type VerifyHeader2 struct {
	// MUST be the same as the PTYPE field in the request PDU header.
	PacketType PacketType
	// MUST be set to 0 when sent and MUST be ignored on receipt.
	_ uint8
	// MUST be set to 0 when sent and MUST be ignored on receipt.
	_ uint16
	// MUST be the same as the drep field in the request PDU header.
	PacketDRep ndr.DataRepresentation
	// MUST be the same as the call_id field in the request PDU header.
	CallID uint32
	// MUST be the same as the p_cont_id field in the request PDU header.
	ContextID uint16
	// MUST be the same as the opnum field in the request PDU header.
	OpNum uint16
}

func (o *VerifyHeader2) WriteTo(ctx context.Context, w ndr.Writer) error {
	w.WriteData((uint8)(o.PacketType))
	w.WriteData(uint8(0))
	w.WriteData(uint16(0))
	w.WriteData(o.PacketDRep)
	w.WriteData(o.CallID)
	w.WriteData(o.ContextID)
	w.WriteData(o.OpNum)
	return w.Err()
}

func (o *VerifyHeader2) ReadFrom(ctx context.Context, r ndr.Reader) error {
	r.ReadData((*uint8)(&o.PacketType))
	_pad := uint8(0)
	r.ReadData(&_pad)
	_pad2 := uint16(0)
	r.ReadData(&_pad2)
	r.ReadData(&o.PacketDRep)
	r.ReadData(&o.CallID)
	r.ReadData(&o.ContextID)
	r.ReadData(&o.OpNum)
	return r.Err()
}

type VerifyPresentation struct {
	// The interface identifier for the presentation context of the
	// request PDU in which this verification trailer appears. This value
	// MUST match the chosen abstract_syntax field from the bind or
	// alter_context PDU where the presentation context was negotiated.
	InterfaceID *SyntaxID
	// The transfer syntax identifier for the presentation context of
	// the request PDU in which this verification trailer appears. This
	// value MUST match the chosen transfer_syntax from the bind or
	// alter_context PDU where the presentation context was negotiated.
	TransferSyntax *SyntaxID
}

func (o *VerifyPresentation) WriteTo(ctx context.Context, w ndr.Writer) error {
	if o.InterfaceID != nil {
		o.InterfaceID.WriteTo(ctx, w)
	} else {
		(&SyntaxID{}).WriteTo(ctx, w)
	}
	if o.TransferSyntax != nil {
		o.TransferSyntax.WriteTo(ctx, w)
	} else {
		(&SyntaxID{}).WriteTo(ctx, w)
	}
	return w.Err()
}

func (o *VerifyPresentation) ReadFrom(ctx context.Context, r ndr.Reader) error {
	o.InterfaceID = &SyntaxID{}
	o.InterfaceID.ReadFrom(ctx, r)
	o.TransferSyntax = &SyntaxID{}
	o.TransferSyntax.ReadFrom(ctx, r)
	return r.Err()
}

type VerificationTrailer struct {
	Commands []*VerificationCommand
}

func (o *VerificationTrailer) VerificationTrailer() VerificationTrailer {

	if o == nil {
		return VerificationTrailer{}
	}

	cmd := make([]*VerificationCommand, len(o.Commands))
	copy(cmd, o.Commands)
	return VerificationTrailer{Commands: cmd}
}

func (o *VerificationTrailer) Size() int {

	if len(o.Commands) == 0 {
		return 0
	}

	sz := 8 /* header */
	for i := range o.Commands {
		sz += 4 /* command header */
		switch cmd := o.Commands[i].Command.(type) {
		case VerifyBitMask:
			sz += 0x0004
		case *VerifyHeader2:
			sz += 0x0010
		case *VerifyPresentation:
			sz += 0x0028
		case []byte:
			sz += len(cmd)
		}
	}

	return sz
}

func (o *VerificationTrailer) WriteTo(ctx context.Context, w ndr.Writer) error {

	w.Write(VerificationTrailerSignature[:])

	for i := range o.Commands {
		command := CommandType(0)
		if o.Commands[i].Required {
			command |= CommandTypeRequired
		}
		if i == len(o.Commands)-1 {
			command |= CommandTypeEnd
		}
		switch cmd := o.Commands[i].Command.(type) {
		case VerifyBitMask:
			w.WriteData((uint16)(command | CommandTypeBitMask))
			w.WriteData(uint16(0x0004))
			cmd.WriteTo(ctx, w)
		case *VerifyHeader2:
			w.WriteData((uint16)(command | CommandTypeHeader2))
			w.WriteData(uint16(0x0010))
			cmd.WriteTo(ctx, w)
		case *VerifyPresentation:
			w.WriteData((uint16)(command | CommandTypePresentation))
			w.WriteData(uint16(0x0028))
			cmd.WriteTo(ctx, w)
		case []byte:
			w.WriteData((uint16)(command | o.Commands[i].CommandType))
			w.WriteData((uint16)(len(cmd)))
			w.Write(cmd)
		}
	}

	return w.Err()
}

func (o *VerificationTrailer) ReadFrom(ctx context.Context, r ndr.Reader) error {

	sgn := make([]byte, 8)
	r.Read(sgn)

	if !bytes.Equal(sgn, VerificationTrailerSignature[:]) {
		return r.SetErr(fmt.Errorf("invalid signature"))
	}

	for {

		command := CommandType(0)
		r.ReadData((*uint16)(&command))

		cmd := &VerificationCommand{Required: command.IsRequired(), CommandType: command & CommandMask}

		length := uint16(0)
		r.ReadData(&length)

		b := make([]byte, length)
		r.Read(b)

		switch r := r.WithBytes(b); {
		case command.Is(CommandTypeBitMask):
			v := VerifyBitMask(0)
			(&v).ReadFrom(ctx, r)
			cmd.Command = v
		case command.Is(CommandTypeHeader2):
			v := &VerifyHeader2{}
			v.ReadFrom(ctx, r)
			cmd.Command = v
		case command.Is(CommandTypePresentation):
			v := &VerifyPresentation{}
			v.ReadFrom(ctx, r)
			cmd.Command = v
		default:
			// write bytes.
			cmd.Command = b
		}

		o.Commands = append(o.Commands, cmd)

		if command.IsEnd() || r.Err() != nil {
			break
		}
	}

	return r.Err()
}

type VerificationCommand struct {
	CommandType CommandType
	Required    bool
	Command     any
}

func (o VerificationCommand) Type() CommandType {

	switch o.Command.(type) {
	case VerifyBitMask:
		return CommandTypeBitMask
	case *VerifyHeader2:
		return CommandTypeHeader2
	case *VerifyPresentation:
		return CommandTypePresentation
	}

	return o.CommandType
}

type CommandType uint16

var (
	// SEC_VT_COMMAND_BITMASK_1: This is an rpc_sec_vt_bitmask command.
	CommandTypeBitMask CommandType = 0x0001
	// SEC_VT_COMMAND_PCONTEXT: This is an rpc_sec_vt_pcontext command.
	CommandTypePresentation CommandType = 0x0002
	// SEC_VT_COMMAND_HEADER2: This is an rpc_sec_vt_header2 command.
	CommandTypeHeader2 CommandType = 0x0003
	// SEC_VT_COMMAND_END: This flag MUST be present in the last command
	// in the verification trailer body.
	CommandTypeEnd CommandType = 0x4000
	// SEC_VT_MUST_PROCESS_COMMAND: Indicates that the server MUST process
	// this command. If the server does not support the command, it MUST
	// reject the request.
	CommandTypeRequired CommandType = 0x8000
	// Mask.
	CommandMask CommandType = 0x3FFF
)

func (o CommandType) IsEnd() bool {
	return o&CommandTypeEnd != 0
}

func (o CommandType) IsRequired() bool {
	return o&CommandTypeRequired != 0
}

func (o CommandType) Is(is CommandType) bool {
	return (o & CommandMask) == is
}

const VerificationTrailerMaxPad = 4
