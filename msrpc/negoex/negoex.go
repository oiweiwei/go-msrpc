// The negoex package implements the NEGOEX client protocol.
//
// # Introduction
//
// The SPNEGO Extended Negotiation (NEGOEX) Security Mechanism enhances the capabilities
// of SPNEGO by providing a security mechanism that can be negotiated by the SPNEGO
// protocol. When the NEGOEX security mechanism is selected by SPNEGO, NEGOEX provides
// a method that allows the selection of a common authentication protocol based on metadata
// such as trust configurations.
//
// # Overview
//
// The SPNEGO Extended Negotiation Security Mechanism (NEGOEX) extends Simple and Protected
// GSS-API Negotiation Mechanism (SPNEGO) described in [RFC4178]. SPNEGO provides a
// negotiation mechanism for Generic Security Services (GSS) API (GSS-API), as described
// in [RFC2743]. NEGOEX is based on the NEGOEX version 4 draft [IETFDRAFT-NEGOEX-04]
// that enhances the capabilities of SPNEGO and provides a security mechanism that can
// be negotiated by the SPNEGO protocol. NEGOEX defines a few new GSS-API extensions
// that a security mechanism MUST support to be negotiated by NEGOEX. This document
// defines these GSS-API extensions. Unlike SPNEGO, NEGOEX defines its own way for signing
// the protocol messages to protect the protocol negotiation.
//
// The NEGOEX protocol is designed to address the drawbacks of the SPNEGO negotiation
// model. When negotiated by SPNEGO, NEGOEX uses the concepts developed in the GSS-API
// specification. The negotiation data is encapsulated in context-level tokens. Therefore,
// callers of the GSS-API do not need to be aware of the existence of the negotiation
// tokens but only of the SPNEGO pseudo-security mechanism. When selected, NEGOEX provides
// a method that allows selection of a common authentication protocol. It preserves
// the optimistic token semantics of SPNEGO and applies that recursively. Accordingly,
// a context establishment mechanism token can be included in the initial NEGOEX message,
// such that NEGOEX does not require an extra round trip when the initiator’s or client’s
// optimistic token is accepted by the target (or server acceptor).
//
// Standard GSS has a strict interpretation of client (initiator) and server (acceptor).
// SPNEGO Extension (SPNG) has extended [RFC4178] to allow the server to initiate SPNG
// message flow. The message flow can begin from either the client or the server as
// the initiator, whereas the receiver is the acceptor. See [MS-SPNG] for client/server
// roles and variations.
package negoex

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
	GoPackage = "negoex"
)

// MessageSignature represents the MESSAGE_SIGNATURE RPC constant
var MessageSignature = 6004521257897772366

// ChecksumSchemeRfc3961 represents the CHECKSUM_SCHEME_RFC3961 RPC constant
var ChecksumSchemeRfc3961 = 1

// AlertTypePulse represents the ALERT_TYPE_PULSE RPC constant
var AlertTypePulse = 1

// AlertVerifyNoKey represents the ALERT_VERIFY_NO_KEY RPC constant
var AlertVerifyNoKey = 1

// ByteVector structure represents BYTE_VECTOR RPC structure.
//
// The BYTE_VECTOR structure encapsulates a variable-length array of octets (or bytes)
// that are stored consecutively. The BYTE_VECTOR structure is used in the following
// structures:
//
// * *CHECKSUM* structure (section 2.2.5.1.3 ( 0655f940-7dc3-4fd1-b249-2d0cb246e5d3
// ) )
//
// * *EXTENSION* structure (section 2.2.5.1.4 ( f0d91d22-fbc6-4ced-870a-be635f72e23b
// ) )
//
// * *EXCHANGE_MESSAGE* structure (section 2.2.6.4 ( aeeaeae6-1782-4d72-af14-a320ad51f0fa
// ) )
//
// struct
//
// {
//
// ULONG ByteArrayOffset;
//
// ULONG ByteArrayLength;
//
// } BYTE_VECTOR;
type ByteVector struct {
	// ByteArrayOffset: A ULONG type array. Each element contains 1 byte.
	ByteArrayOffset uint32 `idl:"name:ByteArrayOffset" json:"byte_array_offset"`
	// ByteArrayLength: A ULONG type that contains the length of the ByteArrayOffset field.
	ByteArrayLength uint32 `idl:"name:ByteArrayLength" json:"byte_array_length"`
}

func (o *ByteVector) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ByteVector) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ByteArrayOffset); err != nil {
		return err
	}
	if err := w.WriteData(o.ByteArrayLength); err != nil {
		return err
	}
	return nil
}
func (o *ByteVector) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ByteArrayOffset); err != nil {
		return err
	}
	if err := w.ReadData(&o.ByteArrayLength); err != nil {
		return err
	}
	return nil
}

// AuthSchemeVector structure represents AUTH_SCHEME_VECTOR RPC structure.
//
// The AUTH_SCHEME_VECTOR structure encapsulates a variable-length array of AUTH_SCHEME
// GUIDs that are stored consecutively.
type AuthSchemeVector struct {
	// AuthSchemeArrayOffset: A ULONG type array of ordered AUTH_SCHEME GUID values, specified
	// in section 2.2.2, that represents the security mechanism's ID in decreasing order
	// of preference.
	AuthSchemeArrayOffset uint32 `idl:"name:AuthSchemeArrayOffset" json:"auth_scheme_array_offset"`
	// AuthSchemeCount: A USHORT that contains the count of AUTH_SCHEME values.
	AuthSchemeCount uint16 `idl:"name:AuthSchemeCount" json:"auth_scheme_count"`
}

func (o *AuthSchemeVector) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthSchemeVector) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.AuthSchemeArrayOffset); err != nil {
		return err
	}
	if err := w.WriteData(o.AuthSchemeCount); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *AuthSchemeVector) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuthSchemeArrayOffset); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuthSchemeCount); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// ExtensionVector structure represents EXTENSION_VECTOR RPC structure.
//
// The EXTENSION_VECTOR structure encapsulates a variable-length array of EXTENSION
// structures (section 2.2.5.1.4) that are stored consecutively. The EXTENSION_VECTOR
// structure is used in the Extensions field in the NEGO_MESSAGE structure, as specified
// in section 2.2.6.3.
type ExtensionVector struct {
	// ExtensionArrayOffset: A ULONG type array. Each element contains an EXTENSION structure,
	// as specified in section 2.2.5.1.4.
	ExtensionArrayOffset uint32 `idl:"name:ExtensionArrayOffset" json:"extension_array_offset"`
	// ExtensionCount: A USHORT that contains the count of elements in the ExtensionArrayOffset
	// field.
	ExtensionCount uint16 `idl:"name:ExtensionCount" json:"extension_count"`
}

func (o *ExtensionVector) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ExtensionVector) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ExtensionArrayOffset); err != nil {
		return err
	}
	if err := w.WriteData(o.ExtensionCount); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *ExtensionVector) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ExtensionArrayOffset); err != nil {
		return err
	}
	if err := w.ReadData(&o.ExtensionCount); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// Extension structure represents EXTENSION RPC structure.
//
// An EXTENSION structure is used in the EXTENSION_VECTOR structure (section 2.2.5.2.4)
// as the Extensions field in the NEGO_MESSAGE structure, as specified in section 2.2.6.3.
type Extension struct {
	// ExtensionType: A ULONG that indicates how the extension data should be interpreted.
	// All negative extension types (the highest bit is set to 1) are critical. If the receiver
	// does not understand a critical extension, the authentication attempt MUST be rejected.
	ExtensionType uint32 `idl:"name:ExtensionType" json:"extension_type"`
	// ExtensionValue: A BYTE_VECTOR structure that contains the extension data.
	ExtensionValue *ByteVector `idl:"name:ExtensionValue" json:"extension_value"`
}

func (o *Extension) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Extension) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ExtensionType); err != nil {
		return err
	}
	if o.ExtensionValue != nil {
		if err := o.ExtensionValue.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ByteVector{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Extension) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ExtensionType); err != nil {
		return err
	}
	if o.ExtensionValue == nil {
		o.ExtensionValue = &ByteVector{}
	}
	if err := o.ExtensionValue.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Checksum structure represents CHECKSUM RPC structure.
//
// The CHECKSUM structure is used in the VERIFY_MESSAGE structure (section 2.2.6.5)
// and is defined as follows.
type Checksum struct {
	// cbHeaderLength: A ULONG that contains the length of the structure definition in octets;
	// this field has a value of 20.
	HeaderLength uint32 `idl:"name:cbHeaderLength" json:"header_length"`
	// ChecksumScheme: A ULONG that describes how checksum is computed and verified. Only
	// the CHECKSUM_SCHEME_RFC3961 is defined, as specified in section 2.2.3. When the value
	// of the ChecksumScheme field is 1 (CHECKSUM_SCHEME_RFC3961), the ChecksumValue field
	// contains a sequence of octets computed according to [RFC3961] and the ChecksumType
	// field contains the checksum type value defined according to [RFC3961].
	ChecksumScheme uint32 `idl:"name:ChecksumScheme" json:"checksum_scheme"`
	// ChecksumType: A ULONG that contains the checksum type of value defined according
	// to [RFC3961].
	ChecksumType uint32 `idl:"name:ChecksumType" json:"checksum_type"`
	// ChecksumValue: A BYTE_VECTOR structure that contains a sequence of octets computed
	// according to [RFC3961].
	ChecksumValue *ByteVector `idl:"name:ChecksumValue" json:"checksum_value"`
}

func (o *Checksum) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Checksum) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.HeaderLength); err != nil {
		return err
	}
	if err := w.WriteData(o.ChecksumScheme); err != nil {
		return err
	}
	if err := w.WriteData(o.ChecksumType); err != nil {
		return err
	}
	if o.ChecksumValue != nil {
		if err := o.ChecksumValue.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ByteVector{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Checksum) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.HeaderLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.ChecksumScheme); err != nil {
		return err
	}
	if err := w.ReadData(&o.ChecksumType); err != nil {
		return err
	}
	if o.ChecksumValue == nil {
		o.ChecksumValue = &ByteVector{}
	}
	if err := o.ChecksumValue.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// AuthScheme structure represents AUTH_SCHEME RPC structure.
type AuthScheme dtyp.GUID

func (o *AuthScheme) GUID() *dtyp.GUID { return (*dtyp.GUID)(o) }

func (o *AuthScheme) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthScheme) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Data1); err != nil {
		return err
	}
	if err := w.WriteData(o.Data2); err != nil {
		return err
	}
	if err := w.WriteData(o.Data3); err != nil {
		return err
	}
	for i1 := range o.Data4 {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.Data4[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data4); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *AuthScheme) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Data1); err != nil {
		return err
	}
	if err := w.ReadData(&o.Data2); err != nil {
		return err
	}
	if err := w.ReadData(&o.Data3); err != nil {
		return err
	}
	o.Data4 = make([]byte, 8)
	for i1 := range o.Data4 {
		i1 := i1
		if err := w.ReadData(&o.Data4[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// ConversationID structure represents CONVERSATION_ID RPC structure.
type ConversationID dtyp.GUID

func (o *ConversationID) GUID() *dtyp.GUID { return (*dtyp.GUID)(o) }

func (o *ConversationID) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ConversationID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Data1); err != nil {
		return err
	}
	if err := w.WriteData(o.Data2); err != nil {
		return err
	}
	if err := w.WriteData(o.Data3); err != nil {
		return err
	}
	for i1 := range o.Data4 {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.Data4[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data4); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *ConversationID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Data1); err != nil {
		return err
	}
	if err := w.ReadData(&o.Data2); err != nil {
		return err
	}
	if err := w.ReadData(&o.Data3); err != nil {
		return err
	}
	o.Data4 = make([]byte, 8)
	for i1 := range o.Data4 {
		i1 := i1
		if err := w.ReadData(&o.Data4[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// MessageType type represents MESSAGE_TYPE RPC enumeration.
//
// The MESSAGE_TYPE enumeration defines the types of messages sent in the MESSAGE_HEADER
// structure MessageType field in every message, as specified in section 2.2.6.2. MESSAGE_TYPE_INITIATOR_NEGO
// type has the value 0, and MESSAGE_TYPE_ALERT type has the value 7. A value is a 4-octet
// number encoded in little-endian.
type MessageType uint16

var (
	// MESSAGE_TYPE_INITIATOR_NEGO: Used in NEGO_MESSAGE, as specified in section 2.2.6.3,
	// to begin negotiation of security mechanisms.
	MessageTypeInitiatorNego MessageType = 0
	// MESSAGE_TYPE_ACCEPTOR_NEGO: Used in NEGO_MESSAGE, as specified in section 2.2.6.3,
	// for the acceptor's output token.
	MessageTypeAcceptorNego MessageType = 1
	// MESSAGE_TYPE_INITIATOR_META_DATA: Used in EXCHANGE_MESSAGE, as specified in section
	// 2.2.6.4, to return a metadata token to NEGOEX for a security mechanism by the initiator.
	MessageTypeInitiatorMetadata MessageType = 2
	// MESSAGE_TYPE_ACCEPTOR_META_DATA: Used in EXCHANGE_MESSAGE, as specified in section
	// 2.2.6.4, to return a metadata token to NEGOEX for a security mechanism by the acceptor.
	MessageTypeAcceptorMetadata MessageType = 3
	// MESSAGE_TYPE_CHALLENGE: Used in EXCHANGE_MESSAGE, as specified in section 2.2.6.4,
	// to encapsulate context tokens of the negotiated security mechanism by the acceptor.
	MessageTypeChallenge MessageType = 4
	// MESSAGE_TYPE_AP_REQUEST: Used in EXCHANGE_MESSAGE, as specified in section 2.2.6.4,
	// to encapsulate context tokens of the negotiated security mechanism by the initiator.
	MessageTypeApRequest MessageType = 5
	// MESSAGE_TYPE_VERIFY: Used in VERIFY_MESSAGE, as specified in section 2.2.6.5, when
	// there is a shared key established that is used to sign all the NEGOEX messages in
	// the negotiation context.
	MessageTypeVerify MessageType = 6
	// MESSAGE_TYPE_ALERT: Used in ALERT_MESSAGE, as specified in section 2.2.6.6, to indicate
	// that the message needs to be resent. Contains the security mechanism, error codes,
	// and various alert types.
	MessageTypeAlert MessageType = 7
)

func (o MessageType) String() string {
	switch o {
	case MessageTypeInitiatorNego:
		return "MessageTypeInitiatorNego"
	case MessageTypeAcceptorNego:
		return "MessageTypeAcceptorNego"
	case MessageTypeInitiatorMetadata:
		return "MessageTypeInitiatorMetadata"
	case MessageTypeAcceptorMetadata:
		return "MessageTypeAcceptorMetadata"
	case MessageTypeChallenge:
		return "MessageTypeChallenge"
	case MessageTypeApRequest:
		return "MessageTypeApRequest"
	case MessageTypeVerify:
		return "MessageTypeVerify"
	case MessageTypeAlert:
		return "MessageTypeAlert"
	}
	return "Invalid"
}

// MessageHeader structure represents MESSAGE_HEADER RPC structure.
//
// The MESSAGE_HEADER structure is a member of other message structures and is used
// to provide metadata about each message. The fields are common for all the NEGOEX
// messages in a conversation exchange except for the MESSAGE_TYPE field, which varies
// according to the message.
type MessageHeader struct {
	// Signature: A ULONG64 type that contains the MESSAGE_SIGNATURE constant in hexadecimal
	// format that indicates "NEGOEXTS", as specified in section 2.2.3. The message signature
	// should remain the same throughout the negotiation process.
	Signature uint64 `idl:"name:Signature" json:"signature"`
	// MessageType: A value of the MESSAGE_TYPE enumeration, as specified in section 2.2.6.1,
	// that contains the type of message.
	MessageType MessageType `idl:"name:MessageType" json:"message_type"`
	// SequenceNum: A ULONG type that contains the message sequence number of the specific
	// conversation, starting with 0 and incremented sequentially.
	SequenceNum uint32 `idl:"name:SequenceNum" json:"sequence_num"`
	// cbHeaderLength: A ULONG type that contains the header length of the message, which
	// includes the message-specific header and excludes the payload.
	HeaderLength uint32 `idl:"name:cbHeaderLength" json:"header_length"`
	// cbMessageLength: A ULONG type that contains the length of the message.
	MessageLength uint32 `idl:"name:cbMessageLength" json:"message_length"`
	// ConversationId: A CONVERSATION_ID GUID, as specified in section 2.2.2, that the initiator
	// and the acceptor use as a context handle to identify an exchange conversation. The
	// CONVERSATION_ID is referred to as ConversationID (section 3.1.5.2). The ConversationID
	// MUST remain the same throughout the entire exchange.
	ConversationID *ConversationID `idl:"name:ConversationId" json:"conversation_id"`
}

func (o *MessageHeader) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MessageHeader) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.Signature); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.MessageType)); err != nil {
		return err
	}
	if err := w.WriteData(o.SequenceNum); err != nil {
		return err
	}
	if err := w.WriteData(o.HeaderLength); err != nil {
		return err
	}
	if err := w.WriteData(o.MessageLength); err != nil {
		return err
	}
	if o.ConversationID != nil {
		if err := o.ConversationID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ConversationID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *MessageHeader) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.Signature); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.MessageType)); err != nil {
		return err
	}
	if err := w.ReadData(&o.SequenceNum); err != nil {
		return err
	}
	if err := w.ReadData(&o.HeaderLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.MessageLength); err != nil {
		return err
	}
	if o.ConversationID == nil {
		o.ConversationID = &ConversationID{}
	}
	if err := o.ConversationID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// NegoMessage structure represents NEGO_MESSAGE RPC structure.
//
// The NEGO_MESSAGE structure message is used to begin and exchange negotiation of security
// mechanisms. This message is sent from the initiator to the acceptor with the message
// type set to MESSAGE_TYPE_INITIATOR_NEGO to begin the negotiation. The initiator uses
// this message to specify the set of supported security mechanisms. The acceptor then
// responds with a NEGO_MESSAGE message, with the message type set to MESSAGE_TYPE_ACCEPTOR_NEGO
// and with its own list of supported security mechanisms. This message contains signatures
// for protecting the NEGOEX negotiation and might also contain the optimistic mechanism
// token.
//
// The NEGO_MESSAGE structure has the following definition.
type NegoMessage struct {
	// Header: A MESSAGE_HEADER structure, as specified in section 2.2.6.2. Its MessageType
	// field can carry from the MESSAGE_TYPE enumeration (section 2.2.6.1) either the value
	// MESSAGE_TYPE_INITIATOR_NEGO for the initiator or the value MESSAGE_TYPE_ACCEPTOR_NEGO
	// for the acceptor.
	Header *MessageHeader `idl:"name:Header" json:"header"`
	// Random: A UCHAR integer array. The Random field is filled using a secure random number
	// generator, as specified in section 2.2.4.
	Random []byte `idl:"name:Random" json:"random"`
	// ProtocolVersion: A ULONG64 type that indicates the numbered version of this protocol.
	// This field contains 0.
	ProtocolVersion uint64 `idl:"name:ProtocolVersion" json:"protocol_version"`
	// AuthSchemes: An AUTH_SCHEME_VECTOR structure, as specified in section 2.2.5.2.2,
	// that contains an ordered list of available, supported security mechanism IDs in decreasing
	// order of preference.
	AuthSchemes *AuthSchemeVector `idl:"name:AuthSchemes" json:"auth_schemes"`
	// Extensions: All negative extension types are critical (the highest bit is set to
	// 1). If the receiver does not understand a critical extension, the authentication
	// attempt MUST be rejected.
	Extensions *ExtensionVector `idl:"name:Extensions" json:"extensions"`
}

func (o *NegoMessage) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *NegoMessage) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.Header != nil {
		if err := o.Header.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&MessageHeader{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	for i1 := range o.Random {
		i1 := i1
		if uint64(i1) >= 32 {
			break
		}
		if err := w.WriteData(o.Random[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Random); uint64(i1) < 32; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ProtocolVersion); err != nil {
		return err
	}
	if o.AuthSchemes != nil {
		if err := o.AuthSchemes.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&AuthSchemeVector{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Extensions != nil {
		if err := o.Extensions.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ExtensionVector{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *NegoMessage) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.Header == nil {
		o.Header = &MessageHeader{}
	}
	if err := o.Header.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	o.Random = make([]byte, 32)
	for i1 := range o.Random {
		i1 := i1
		if err := w.ReadData(&o.Random[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.ProtocolVersion); err != nil {
		return err
	}
	if o.AuthSchemes == nil {
		o.AuthSchemes = &AuthSchemeVector{}
	}
	if err := o.AuthSchemes.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Extensions == nil {
		o.Extensions = &ExtensionVector{}
	}
	if err := o.Extensions.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// ExchangeMessage structure represents EXCHANGE_MESSAGE RPC structure.
//
// The EXCHANGE_MESSAGE structure message is used to encapsulate context tokens of the
// negotiated security mechanism for either the initiator or the acceptor.
type ExchangeMessage struct {
	// Header: A MESSAGE_HEADER structure, as specified in section 2.2.6.2. The Header’s
	// MessageType field is set from the values of the MESSAGE_TYPE enumeration, as specified
	// in section 2.2.6.1. The MessageType field MUST be set to MESSAGE_TYPE_AP_REQUEST
	// type for the initiator or MESSAGE_TYPE_CHALLENGE type for the acceptor when context
	// tokens are being exchanged. The MessageType field MUST be set to MESSAGE_TYPE_INITIATOR_META_DATA
	// type for the initiator or MESSAGE_TYPE_ACCEPTOR_META_DATA type for the acceptor when
	// metadata tokens are being exchanged.
	Header *MessageHeader `idl:"name:Header" json:"header"`
	// AuthScheme: An AUTH_SCHEME GUID that contains the security mechanism's ID, as specified
	// in section 2.2.2.
	AuthScheme *AuthScheme `idl:"name:AuthScheme" json:"auth_scheme"`
	// Exchange: A BYTE_VECTOR structure, specified in section 2.2.5.2.3, that contains
	// the opaque handshake message for the client authentication scheme.
	Exchange *ByteVector `idl:"name:Exchange" json:"exchange"`
}

func (o *ExchangeMessage) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ExchangeMessage) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.Header != nil {
		if err := o.Header.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&MessageHeader{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.AuthScheme != nil {
		if err := o.AuthScheme.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&AuthScheme{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Exchange != nil {
		if err := o.Exchange.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ByteVector{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *ExchangeMessage) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.Header == nil {
		o.Header = &MessageHeader{}
	}
	if err := o.Header.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.AuthScheme == nil {
		o.AuthScheme = &AuthScheme{}
	}
	if err := o.AuthScheme.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Exchange == nil {
		o.Exchange = &ByteVector{}
	}
	if err := o.Exchange.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// VerifyMessage structure represents VERIFY_MESSAGE RPC structure.
//
// A VERIFY_MESSAGE structure message is produced using the required checksum mechanism
// per [RFC3961] and is included in the output token.
type VerifyMessage struct {
	// Header: A MESSAGE_HEADER structure, as specified in section 2.2.6.2. The header’s
	// message type MUST be set to the MESSAGE_TYPE_VERIFY value from the MESSAGE_TYPE enumeration,
	// as specified in section 2.2.6.1.
	Header *MessageHeader `idl:"name:Header" json:"header"`
	// AuthScheme: An AUTH_SCHEME GUID, as specified in section 2.2.2, that identifies the
	// security mechanism ID from which the protocol key was obtained.
	AuthScheme *AuthScheme `idl:"name:AuthScheme" json:"auth_scheme"`
	// Checksum: A CHECKSUM structure, specified in section 2.2.5.1.3, that contains the
	// checksum of all the previously exchanged messages in the order they were sent in
	// the conversation. The checksum is calculated based on [RFC3961].
	Checksum *Checksum `idl:"name:Checksum" json:"checksum"`
}

func (o *VerifyMessage) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *VerifyMessage) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.Header != nil {
		if err := o.Header.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&MessageHeader{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.AuthScheme != nil {
		if err := o.AuthScheme.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&AuthScheme{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Checksum != nil {
		if err := o.Checksum.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Checksum{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *VerifyMessage) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.Header == nil {
		o.Header = &MessageHeader{}
	}
	if err := o.Header.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.AuthScheme == nil {
		o.AuthScheme = &AuthScheme{}
	}
	if err := o.AuthScheme.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Checksum == nil {
		o.Checksum = &Checksum{}
	}
	if err := o.Checksum.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// Alert structure represents ALERT RPC structure.
//
// The ALERT structure is used in the ALERT_VECTOR structure, which is used in the Alerts
// field of the ALERT_MESSAGE structure message, as specified in section 2.2.6.6.
type Alert struct {
	// AlertType: A ULONG that indicates the type of the alert.
	AlertType uint32 `idl:"name:AlertType" json:"alert_type"`
	// AlertValue: A BYTE_VECTOR structure, as specified in section 2.2.5.2.3, that contains
	// an array of alert values.
	AlertValue *ByteVector `idl:"name:AlertValue" json:"alert_value"`
}

func (o *Alert) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Alert) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.AlertType); err != nil {
		return err
	}
	if o.AlertValue != nil {
		if err := o.AlertValue.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ByteVector{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Alert) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.AlertType); err != nil {
		return err
	}
	if o.AlertValue == nil {
		o.AlertValue = &ByteVector{}
	}
	if err := o.AlertValue.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// AlertPulse structure represents ALERT_PULSE RPC structure.
//
// The ALERT_PULSE structure is used in the ALERT_MESSAGE structure message, as specified
// in section 2.2.6.6.
type AlertPulse struct {
	// cbHeaderLength: A ULONG that contains the header length of this message.
	HeaderLength uint32 `idl:"name:cbHeaderLength" json:"header_length"`
	// Reason: A ULONG that contains the reason code for the heartbeat message ALERT_VERIFY_NO_KEY,
	// as specified in section 2.2.3.
	Reason uint32 `idl:"name:Reason" json:"reason"`
}

func (o *AlertPulse) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AlertPulse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.HeaderLength); err != nil {
		return err
	}
	if err := w.WriteData(o.Reason); err != nil {
		return err
	}
	return nil
}
func (o *AlertPulse) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.HeaderLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.Reason); err != nil {
		return err
	}
	return nil
}

// AlertVector structure represents ALERT_VECTOR RPC structure.
//
// The ALERT_VECTOR structure contains the alert types and the count of alerts used
// in the ALERT_MESSAGE structure, as specified in section 2.2.6.6.
type AlertVector struct {
	// AlertArrayOffset: A ULONG that contains array of ALERT structures, as specified in
	// section 2.2.5.1.1.
	AlertArrayOffset uint32 `idl:"name:AlertArrayOffset" json:"alert_array_offset"`
	// AlertCount: A USHORT that contains the number of alerts in the AlertArrayOffset field.
	AlertCount uint16 `idl:"name:AlertCount" json:"alert_count"`
}

func (o *AlertVector) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AlertVector) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.AlertArrayOffset); err != nil {
		return err
	}
	if err := w.WriteData(o.AlertCount); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *AlertVector) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.AlertArrayOffset); err != nil {
		return err
	}
	if err := w.ReadData(&o.AlertCount); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// AlertMessage structure represents ALERT_MESSAGE RPC structure.
//
// The ALERT_MESSAGE structure message is sent by the initiator or the acceptor requesting
// that the peer resend the message. The ALERT_MESSAGE might not always be sent.
type AlertMessage struct {
	// Header: A MESSAGE_HEADER structure, as specified in section 2.2.6.2. The header’s
	// message type MUST be set to the MESSAGE_TYPE_ALERT value from the MESSAGE_TYPE enumeration,
	// as specified in section 2.2.6.1.
	Header *MessageHeader `idl:"name:Header" json:"header"`
	// AuthScheme: An AUTH_SCHEME GUID, as specified in section 2.2.2, that indicates the
	// security mechanism ID to which the alert message is targeted.
	AuthScheme *AuthScheme `idl:"name:AuthScheme" json:"auth_scheme"`
	// ErrorCode: A ULONG type indicating an NTSTATUS code, as specified in [MS-ERREF] section
	// 2.3.
	ErrorCode uint32 `idl:"name:ErrorCode" json:"error_code"`
	// Alerts: An ALERT_VECTOR structure, as specified in section 2.2.5.2.1, that contains
	// ALERT structures, as specified in section 2.2.5.1.1.
	Alerts *AlertVector `idl:"name:Alerts" json:"alerts"`
}

func (o *AlertMessage) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AlertMessage) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.Header != nil {
		if err := o.Header.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&MessageHeader{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.AuthScheme != nil {
		if err := o.AuthScheme.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&AuthScheme{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ErrorCode); err != nil {
		return err
	}
	if o.Alerts != nil {
		if err := o.Alerts.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&AlertVector{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *AlertMessage) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.Header == nil {
		o.Header = &MessageHeader{}
	}
	if err := o.Header.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.AuthScheme == nil {
		o.AuthScheme = &AuthScheme{}
	}
	if err := o.AuthScheme.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.ErrorCode); err != nil {
		return err
	}
	if o.Alerts == nil {
		o.Alerts = &AlertVector{}
	}
	if err := o.Alerts.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}
