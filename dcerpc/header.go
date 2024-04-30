package dcerpc

// header.go contains the connection-oriented common
// header definitions.

import (
	"context"

	"github.com/oiweiwei/go-msrpc/midl/uuid"
	"github.com/oiweiwei/go-msrpc/ndr"
	"github.com/rs/zerolog"
)

// The common header size.
const HeaderSize = 16

type PacketType uint8

const (
	PacketTypeRequest              PacketType = 0
	PacketTypeResponse             PacketType = 2
	PacketTypeFault                PacketType = 3
	PacketTypeBind                 PacketType = 11
	PacketTypeBindAck              PacketType = 12
	PacketTypeBindNak              PacketType = 13
	PacketTypeAlterContext         PacketType = 14
	PacketTypeAlterContextResponse PacketType = 15
	PacketTypeAuth3                PacketType = 16
	PacketTypeShutdown             PacketType = 17
	PacketTypeCancel               PacketType = 18
	PacketTypeOrphaned             PacketType = 19
)

func (f PacketType) String() string {
	switch f {
	case PacketTypeRequest:
		return "request"
	case PacketTypeResponse:
		return "response"
	case PacketTypeFault:
		return "fault"
	case PacketTypeBind:
		return "bind"
	case PacketTypeBindAck:
		return "bind_ack"
	case PacketTypeBindNak:
		return "bind_nak"
	case PacketTypeAlterContext:
		return "alter_context"
	case PacketTypeAlterContextResponse:
		return "alter_context_response"
	case PacketTypeAuth3:
		return "auth3"
	case PacketTypeShutdown:
		return "shutdown"
	case PacketTypeCancel:
		return "cancel"
	case PacketTypeOrphaned:
		return "orphaned"
	}

	return "unknown"
}

type PacketFlag uint8

func (f PacketFlag) IsSet(ff PacketFlag) bool {
	return f&ff != 0
}

func (f PacketFlag) String() string {

	ret, sep := "", ""
	if f.IsSet(PacketFlagFirstFrag) {
		ret, sep = ret+sep+"first", "|"
	}

	if f.IsSet(PacketFlagLastFrag) {
		ret, sep = ret+sep+"last", "|"
	}

	if f.IsSet(PacketFlagSupportHeaderSign) {
		ret, sep = ret+sep+"support_header_sign_or_cancel", "|"
	}

	if f.IsSet(PacketFlagConcMPX) {
		ret, sep = ret+sep+"multiplexing", "|"
	}

	if f.IsSet(PacketFlagDidNotExecute) {
		ret, sep = ret+sep+"did_not_execute", "|"
	}

	if f.IsSet(PacketFlagMaybe) {
		ret, sep = ret+sep+"maybe", "|"
	}

	if f.IsSet(PacketFlagObjectUUID) {
		ret, sep = ret+sep+"object", "|"
	}

	return ret
}

const (
	// PFC_FIRST_FRAG: First fragment.
	PacketFlagFirstFrag PacketFlag = 1 << 0
	// PFC_LAST_FRAG: Last fragment.
	PacketFlagLastFrag PacketFlag = 1 << 1
	// PFC_PENDING_CANCEL: Cancel was pending at sender.
	PacketFlagPendingCancel PacketFlag = 1 << 2
	// PFC_SUPPORT_HEADER_SIGN: Supports the header/trailer protection.
	// (same as PFC_PENDING_CANCEL, only during Bind/BindAck).
	PacketFlagSupportHeaderSign PacketFlag = 1 << 2
	// PFC_RESERVED_1: Reserved1.
	PacketFlagReserved1 PacketFlag = 1 << 3
	// PFC_CONC_MPX: Supports concurrent multiplexing of a single connection.
	PacketFlagConcMPX PacketFlag = 1 << 4
	// PFC_DID_NOT_EXECUTE: Only meaningful on 'fault' packet; If true,
	// guaranteed call did not execute.
	PacketFlagDidNotExecute PacketFlag = 1 << 5
	// PFC_MAYBE: 'maybe' call semantics requested
	PacketFlagMaybe PacketFlag = 1 << 6
	// PFC_OBJECT_UUID: if true, a non-nil object UUID was specified in the
	// handle, and is present in the optional object field. If false, the
	// object field is omitted.
	PacketFlagObjectUUID PacketFlag = 1 << 7
)

// FlagObjectUUID function computes the PFC_OBJECT_UUID flag based on
// the value `obj`.
func FlagObjectUUID(obj *uuid.UUID) PacketFlag {
	if obj != nil {
		return PacketFlagObjectUUID
	}
	return 0
}

type Header struct {
	RPCVersion      uint8
	RPCVersionMinor uint8
	PacketType      PacketType
	PacketFlags     PacketFlag
	PacketDRep      ndr.DataRepresentation
	FragLength      uint16
	AuthLength      uint16
	CallID          uint32
}

func (h Header) MarshalZerologObject(e *zerolog.Event) {
	e.Stringer("packet_type", h.PacketType)
	e.Stringer("flags", h.PacketFlags)
	e.Uint16("frag_length", h.FragLength)
	e.Uint16("auth_length", h.AuthLength)
	e.Uint32("call_id", h.CallID)
}

func (h *Header) WriteTo(ctx context.Context, w ndr.Writer) error {
	w.WriteData(h.RPCVersion)
	w.WriteData(h.RPCVersionMinor)
	w.WriteData((uint8)(h.PacketType))
	w.WriteData((uint8)(h.PacketFlags))
	w.WriteData(h.PacketDRep)
	w.WriteData(h.FragLength)
	w.WriteData(h.AuthLength)
	w.WriteData(h.CallID)
	return w.Err()
}

func (h *Header) ReadFrom(ctx context.Context, r ndr.Reader) error {
	r.ReadData(&h.RPCVersion)
	r.ReadData(&h.RPCVersionMinor)
	r.ReadData((*uint8)(&h.PacketType))
	r.ReadData((*uint8)(&h.PacketFlags))
	r.ReadData(&h.PacketDRep)
	r.ReadData(&h.FragLength)
	r.ReadData(&h.AuthLength)
	r.ReadData(&h.CallID)
	return r.Err()
}
