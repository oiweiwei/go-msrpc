package dcerpc

// pdu.go contains the connection-oriented PDU definitions.

import (
	"context"

	"github.com/oiweiwei/go-msrpc/midl/uuid"
	"github.com/oiweiwei/go-msrpc/ndr"
	"github.com/rs/zerolog"
)

const (
	RequestSize    = HeaderSize + 8
	ObjectUUIDSize = 16
)

type PDU interface {
	WriteTo(context.Context, ndr.Writer) error
	ReadFrom(context.Context, ndr.Reader) error
	MarshalZerologObject(*zerolog.Event)
}

var (
	_ PDU = (*Auth3)(nil)
	_ PDU = (*AlterContext)(nil)
	_ PDU = (*AlterContextResponse)(nil)
	_ PDU = (*Bind)(nil)
	_ PDU = (*BindAck)(nil)
	_ PDU = (*BindNak)(nil)
	_ PDU = (*Cancel)(nil)
	_ PDU = (*Fault)(nil)
	_ PDU = (*Orphaned)(nil)
	_ PDU = (*Request)(nil)
	_ PDU = (*Response)(nil)
	_ PDU = (*Shutdown)(nil)
)

func PDUToPacketType(pdu PDU) PacketType {
	switch pdu.(type) {
	case *Auth3:
		return PacketTypeAuth3
	case *AlterContext:
		return PacketTypeAlterContext
	case *AlterContextResponse:
		return PacketTypeAlterContextResponse
	case *Bind:
		return PacketTypeBind
	case *BindAck:
		return PacketTypeBindAck
	case *BindNak:
		return PacketTypeBindNak
	case *Cancel:
		return PacketTypeCancel
	case *Fault:
		return PacketTypeFault
	case *Orphaned:
		return PacketTypeOrphaned
	case *Request:
		return PacketTypeRequest
	case *Response:
		return PacketTypeResponse
	case *Shutdown:
		return PacketTypeShutdown
	}
	return 0
}

type Auth3 struct {
	Pad [4]byte
}

func (pdu *Auth3) Size() int {
	return 4
}

func (pdu *Auth3) MarshalZerologObject(e *zerolog.Event) {}

// marshal function ...
func (pdu *Auth3) WriteTo(ctx context.Context, w ndr.Writer) error {
	for i1 := range pdu.Pad {
		w.WriteData(pdu.Pad[i1])
	}
	return w.Err()
}

// ReadFrom function ...
func (pdu *Auth3) ReadFrom(ctx context.Context, r ndr.Reader) error {
	for i1 := range pdu.Pad {
		r.ReadData(&pdu.Pad[i1])
	}
	return r.Err()
}

// The AlterContextPDU is used to request additional presentation negotiation
// for another interface and/or version, or to negotiate a new security context,
// or both. The format is identical to the BindPDU, except that the value of
// the COHeader.PacketType field is set to PacketTypeAlterContext.
// The MaxXmitFrag, MaxRecvFrag, AssocGroupID and SecAddressLength/PortSpec fields
// are ignored.
type AlterContext struct {
	MaxXmitFrag  uint16
	MaxRecvFrag  uint16
	AssocGroupID uint32
	ContextList  []*Context
}

func (pdu *AlterContext) Size() int {
	size := 8 + 4
	for _, ctx := range pdu.ContextList {
		size += ctx.Size()
	}
	return size
}

func (pdu *AlterContext) MarshalZerologObject(e *zerolog.Event) {}

func (pdu *AlterContext) WriteTo(ctx context.Context, w ndr.Writer) error {
	w.WriteData(pdu.MaxXmitFrag)
	w.WriteData(pdu.MaxRecvFrag)
	w.WriteData(pdu.AssocGroupID)
	_context_list := &ContextList{NContextElem: uint8(len(pdu.ContextList)), ContextElem: pdu.ContextList}
	_context_list.WriteTo(ctx, w)
	return w.Err()
}

func (pdu *AlterContext) ReadFrom(ctx context.Context, r ndr.Reader) error {
	r.ReadData(&pdu.MaxXmitFrag)
	r.ReadData(&pdu.MaxRecvFrag)
	r.ReadData(&pdu.AssocGroupID)
	_context_list := &ContextList{}
	_context_list.ReadFrom(ctx, r)
	pdu.ContextList = _context_list.ContextElem
	return r.Err()
}

// The AlterContextResponsePDU is used to indicate the server's response to an
// AlterContext request. The PDU format is identical to BindAckPDU, except that
// the value of the COHeader.PacketType field is set to PacketTypeAlterContextResponse.
// The MaxXmitFrag, MaxRecvFrag, AssocGroupID and SecAddressLength/PortSpec fields
// are ignored.
type AlterContextResponse struct {
	MaxXmitFrag  uint16
	MaxRecvFrag  uint16
	AssocGroupID uint32
	PortSpec     string
	ResultList   []*Result
}

func (pdu *AlterContextResponse) MarshalZerologObject(e *zerolog.Event) {}

func (pdu *AlterContextResponse) WriteTo(ctx context.Context, w ndr.Writer) error {
	w.WriteData(pdu.MaxXmitFrag)
	w.WriteData(pdu.MaxRecvFrag)
	w.WriteData(pdu.AssocGroupID)
	_port_any := &PortAny{PortSpec: pdu.PortSpec}
	_port_any.WriteTo(ctx, w)
	w.WriteAlign(4)
	_results := &ResultList{NResults: uint8(len(pdu.ResultList)), Results: pdu.ResultList}
	_results.WriteTo(ctx, w)
	return w.Err()
}

func (pdu *AlterContextResponse) ReadFrom(ctx context.Context, r ndr.Reader) error {
	r.ReadData(&pdu.MaxXmitFrag)
	r.ReadData(&pdu.MaxRecvFrag)
	r.ReadData(&pdu.AssocGroupID)
	_port_any := &PortAny{}
	_port_any.ReadFrom(ctx, r)
	pdu.PortSpec = _port_any.PortSpec
	r.ReadAlign(4)
	_results := &ResultList{}
	_results.ReadFrom(ctx, r)
	pdu.ResultList = _results.Results
	return r.Err()
}

// The BindPDU is used to initiate the presentation negotiation for the body
// data, and optionally, authentication. The presentation negotiation follows
// the model of the OSI presentation layer.
//
// The PDU contains a priority-ordered list of supported presentation syntaxes,
// both abstract and transfer, and context identifiers (local handles). (This
// differs from OSI, which does not specify any order for the list.) The abstract
// and transfer syntaxes are represented as a record of interface UUID and
// interface version. (These may map one-to-one into OSI object identifiers by
// providing suitable prefixes and changing the encoding.) Each supported data
// representation, such as NDR, will be assigned an interface UUID, and will
// use that UUID as part of its transfer syntax value. Each stub computes its
// abstract syntax value given its interface UUID and interface version.
//
// If COHeader.PFCFlags does not have PFCFlagCOLastFrag set and RPCVersMinor is 1,
// then the PDU has fragmented AuthVerifier data. The server will assemble the
// data concatenating sequentially each AuthVerifier field until a PDU is sent
// with PFCFlagCOLastFrag flag set. This completed buffer is then used as
// AuthVerifier data.
//
// The fields MaxXmitFrag and MaxRecvFrag are used for fragment size
// negotiation.
//
// The AssocGroupID field contains either an association group identifier that
// was created during a previous bind negotiation or 0 (zero) to indicate a
// request for a new group.
//
// This PDU shall not exceed the MustRecvFragSize, since no size negotiation
// has yet occurred. If the ContextList is too long, the leading subset should
// be transmitted, and additional presentation context negotiation can occur
// in subsequent AlterContext PDUs, as needed, after a successful BindAck.
type Bind struct {
	MaxXmitFrag  uint16
	MaxRecvFrag  uint16
	AssocGroupID uint32
	ContextList  []*Context
}

func (pdu *Bind) Size() int {
	size := 8 + 4
	for _, ctx := range pdu.ContextList {
		size += ctx.Size()
	}
	return size
}

func (pdu *Bind) MarshalZerologObject(e *zerolog.Event) {}

func (pdu *Bind) WriteTo(ctx context.Context, w ndr.Writer) error {
	w.WriteData(pdu.MaxXmitFrag)
	w.WriteData(pdu.MaxRecvFrag)
	w.WriteData(pdu.AssocGroupID)
	_context_list := &ContextList{NContextElem: uint8(len(pdu.ContextList)), ContextElem: pdu.ContextList}
	_context_list.WriteTo(ctx, w)
	return w.Err()
}

func (pdu *Bind) ReadFrom(ctx context.Context, r ndr.Reader) error {
	r.ReadData(&pdu.MaxXmitFrag)
	r.ReadData(&pdu.MaxRecvFrag)
	r.ReadData(&pdu.AssocGroupID)
	_context_list := &ContextList{}
	_context_list.ReadFrom(ctx, r)
	pdu.ContextList = _context_list.ContextElem
	return r.Err()
}

// The BindAckPDU is returned by the server when it accepts a bind request
// initiated by the client's bindPDU. It contains the results of presentation
// context and fragment size negotiations.
//
// It may also contain a new association group identifier if one was requested
// by the client.
//
// The MaxXmitFrag and MaxRecvFrag fields contain the maximum transmit and
// receive fragment sizes as determined by the server in response to the
// client's desired sizes.
//
// The ResultList contains the results of the presentation context negotiation
// initiated by the client. It is possible for a BindAckPDU not to contain any
// mutually supported syntaxes.
//
// If the client requested a new association group, AssocGroupID contains the
// identifier of the new association group created by the server. Otherwise,
// it contains the identifier of the previously created association group
// requested by the client.
type BindAck struct {
	MaxXmitFrag  uint16
	MaxRecvFrag  uint16
	AssocGroupID uint32
	PortSpec     string
	ResultList   []*Result
}

func (pdu *BindAck) MarshalZerologObject(e *zerolog.Event) {}

func (pdu *BindAck) WriteTo(ctx context.Context, w ndr.Writer) error {
	w.WriteData(pdu.MaxXmitFrag)
	w.WriteData(pdu.MaxRecvFrag)
	w.WriteData(pdu.AssocGroupID)
	(&PortAny{PortSpec: pdu.PortSpec}).WriteTo(ctx, w)
	w.WriteAlign(4)
	_results := &ResultList{NResults: uint8(len(pdu.ResultList)), Results: pdu.ResultList}
	_results.WriteTo(ctx, w)
	return w.Err()
}

func (pdu *BindAck) ReadFrom(ctx context.Context, r ndr.Reader) error {
	r.ReadData(&pdu.MaxXmitFrag)
	r.ReadData(&pdu.MaxRecvFrag)
	r.ReadData(&pdu.AssocGroupID)
	_port_any := &PortAny{}
	_port_any.ReadFrom(ctx, r)
	pdu.PortSpec = _port_any.PortSpec
	r.ReadAlign(4)
	_results := &ResultList{}
	_results.ReadFrom(ctx, r)
	pdu.ResultList = _results.Results
	return r.Err()
}

// The BindNakPDU is returned by the server when it rejects an association
// request initiated by the client's BindPDU. The ProviderRejectReason field
// holds the rejection reason code. When the reject reason is
// ProtocolVersionNotSupported, the versions field contains a list of runtime
// protocol versions supported by the server.
//
// The BindNakPDU never contains an authentication verifier.
type BindNak struct {
	ProviderRejectReason ProviderReason
	// VersionList represents the list of supported protocols returned when the
	// protocol negotiation fails.
	VersionList []*Version
}

func (pdu *BindNak) MarshalZerologObject(e *zerolog.Event) {}

func (pdu *BindNak) Error() string {

	switch pdu.ProviderRejectReason {
	case AbstractSyntaxNotSupported:
		return "bind: abstract syntax is not supported"
	case ProposedTransferSyntaxesNotSupported:
		return "bind: proposed transfer syntaxes not supported"
	case LocalLimitExceeded:
		return "bind: local limit exceeded"
	case AuthTypeNotRecognized:
		return "bind: authentication type was not recognized"
	case InvalidChecksum:
		return "bind: invalid checksum"
	case ProtocolVersionNotSupported:
		return "bind: protocol version not supported"
	default:
		return "bind: unknown error"
	}
}

func (pdu *BindNak) WriteTo(ctx context.Context, w ndr.Writer) error {
	w.WriteData((uint16)(pdu.ProviderRejectReason))
	_versions := &VersionList{NProtocols: uint8(len(pdu.VersionList)), Versions: pdu.VersionList}
	_versions.WriteTo(ctx, w)
	return w.Err()
}

func (pdu *BindNak) ReadFrom(ctx context.Context, r ndr.Reader) error {
	r.ReadData((*uint16)(&pdu.ProviderRejectReason))
	_versions := &VersionList{}
	_versions.ReadFrom(ctx, r)
	pdu.VersionList = _versions.Versions
	return r.Err()
}

// The CoCancelPDU is used to forward a cancel.
type Cancel struct{}

func (pdu *Cancel) MarshalZerologObject(e *zerolog.Event) {}

func (pdu *Cancel) WriteTo(ctx context.Context, w ndr.Writer) error { return w.Err() }

func (pdu *Cancel) ReadFrom(ctx context.Context, r ndr.Reader) error { return r.Err() }

// The FaultPDU is used to indicate either an RPC run-time, RPC stub, or
// RPC-specific exception to the client. The ContextID field holds a context
// identifier that identifies the data representation.
//
// The AllocHint field is optionally used by the client to provide a hint to
// the receiver of the amount of buffer space to allocate contiguously for
// fragmented requests. This is only a potential optimisation. The server must
// work correctly regardless of the value passed. The value 0 (zero) is reserved
// to indicate that the transmitter is not supplying any information.
//
// The Status field indicates run-time status. The value may either be an architected
// non-zero value, indicating a run-time error, such as an interface version
// mismatch, or 0 (zero), indicating a stub defined exception that is specified
// with the stub data. If a non-zero value is present, no stub data is allowed.
//
// Certain status values imply that the call did not execute. To keep such
// status values consistent with the flag, an implementation should model all
// fault messages as being initialised with the COHeader.PFCFlagCODidNotExecute flag
// set, then cleared when the run-time system (or stub, if the implementation
// allows) passes control to the server stub routine.
type Fault struct {
	AllocHint   uint32
	ContextID   uint16
	CancelCount uint8
	Flags       uint8
	Status      uint32
	Pad         [4]byte
}

func (pdu *Fault) MarshalZerologObject(e *zerolog.Event) {
	e.Uint32("alloc_hint", pdu.AllocHint)
	e.Uint16("context_id", pdu.ContextID)
	e.Uint8("cancel_count", pdu.CancelCount)
	e.Uint8("flags", pdu.Flags)
	e.Uint32("status", pdu.Status)
}

func (pdu *Fault) WriteTo(ctx context.Context, w ndr.Writer) error {
	w.WriteData(pdu.AllocHint)
	w.WriteData(pdu.ContextID)
	w.WriteData(pdu.CancelCount)
	w.WriteData(pdu.Flags)
	w.WriteData(pdu.Status)
	for i1 := range pdu.Pad {
		w.WriteData(pdu.Pad[i1])
	}
	return w.Err()
}

func (pdu *Fault) ReadFrom(ctx context.Context, r ndr.Reader) error {
	r.ReadData(&pdu.AllocHint)
	r.ReadData(&pdu.ContextID)
	r.ReadData(&pdu.CancelCount)
	r.ReadData(&pdu.Flags)
	r.ReadData(&pdu.Status)
	for i1 := range pdu.Pad {
		r.ReadData(&pdu.Pad[i1])
	}
	return r.Err()
}

// The OrphanedPDU is used by a client to notify a server that it is aborting
// a request in progress that has not been entirely transmitted yet, or that
// it is aborting a (possibly lengthy) response in progress.
type Orphaned struct{}

func (pdu *Orphaned) MarshalZerologObject(e *zerolog.Event) {}

func (pdu *Orphaned) WriteTo(ctx context.Context, w ndr.Writer) error { return w.Err() }

func (pdu *Orphaned) ReadFrom(ctx context.Context, r ndr.Reader) error { return r.Err() }

// The RequestPDU is used for an initial call request. The ContextID field
// holds a presentation context identifier that identifies the data representation.
// The OpNum field identifies the operation being invoked within the interface.
//
// The PDU may also contain an Object UUID. In this case the PFCFlagCOObjectUUID
// flag is set in COHeader.PFCFlags, and the PDU includes the object field. If
// the PFCFlagCOObjectUUID flag is not set, the PDU does not include the object
// field.
//
// The AllocHint field is optionally used by the client to provide a hint to the
// receiver of the amount of buffer space to allocate contiguously for
// fragmented requests. This is only a potential optimisation. The server must
// work correctly regardless of the value passed. The value 0 (zero) is reserved to
// indicate that the transmitter is not supplying any information.
type Request struct {
	AllocHint  uint32
	ContextID  uint16
	OpNum      uint16
	ObjectUUID *uuid.UUID
}

func (pdu *Request) MarshalZerologObject(e *zerolog.Event) {
	e.Uint32("alloc_hint", pdu.AllocHint).Uint16("context_id", pdu.ContextID).Uint16("op_num", pdu.OpNum)
	if pdu.ObjectUUID != nil {
		e.Stringer("object", pdu.ObjectUUID)
	}
}

func (pdu *Request) WriteTo(ctx context.Context, w ndr.Writer) error {
	w.WriteData(pdu.AllocHint)
	w.WriteData(pdu.ContextID)
	w.WriteData(pdu.OpNum)
	if pdu.ObjectUUID != nil {
		w.SetErr(pdu.ObjectUUID.Write(w))
	}
	return w.Err()
}

func (pdu *Request) ReadFrom(ctx context.Context, r ndr.Reader) error {
	r.ReadData(&pdu.AllocHint)
	r.ReadData(&pdu.ContextID)
	r.ReadData(&pdu.OpNum)
	if pdu.ObjectUUID != nil {
		r.SetErr(pdu.ObjectUUID.Read(r))
	}
	return r.Err()
}

// The ResponsePDU is used to respond to an active call. The ContextID field
// holds a context identifier that identifies the data representation. The
// CancelCount field holds a count of cancels received.
//
// The AllocHint field is optionally used by the client to provide a hint to the
// receiver of the amount of buffer space to allocate contiguously for
// fragmented requests. This is only a potential optimisation. The server must
// work correctly regardless of the value passed. The value 0 (zero) is reserved to
// indicate that the transmitter is not supplying any information.
type Response struct {
	AllocHint   uint32
	ContextID   uint16
	CancelCount uint8
	_           uint8 // pad.
}

func (pdu *Response) MarshalZerologObject(e *zerolog.Event) {
	e.Uint32("alloc_hint", pdu.AllocHint).Uint16("context_id", pdu.ContextID).Uint8("cancel_count", pdu.CancelCount)
}

func (pdu *Response) WriteTo(ctx context.Context, w ndr.Writer) error {
	w.WriteData(pdu.AllocHint)
	w.WriteData(pdu.ContextID)
	w.WriteData((uint8)(0)) // pad.
	return w.Err()
}

func (pdu *Response) ReadFrom(ctx context.Context, r ndr.Reader) error {
	r.ReadData(&pdu.AllocHint)
	r.ReadData(&pdu.ContextID)
	r.ReadData(&pdu.CancelCount)
	_pad := uint8(0)
	r.ReadData(&_pad)
	return r.Err()
}

// The ShutdownPDU is sent by the server to request that a client terminate
// the connection, freeing the related resources.
//
// The ShutdownPDU never contains an authentication verifier even if
// authentication services are in use.
type Shutdown struct{}

func (pdu *Shutdown) MarshalZerologObject(e *zerolog.Event) {}

func (pdu *Shutdown) WriteTo(ctx context.Context, w ndr.Writer) error { return w.Err() }

func (pdu *Shutdown) ReadFrom(ctx context.Context, r ndr.Reader) error { return r.Err() }
