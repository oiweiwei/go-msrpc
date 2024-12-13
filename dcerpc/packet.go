package dcerpc

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/oiweiwei/go-msrpc/dcerpc/errors"
	"github.com/oiweiwei/go-msrpc/midl/uuid"
	"github.com/oiweiwei/go-msrpc/ndr"
)

// SecurityTrailerSize is the size of the security trailer.
const SecurityTrailerSize = 8

// MaxPad specifies the maximum possible padding.
const MaxPad = 16

// Packet represents the raw DCE/RPC packet structure.
type Packet struct {
	// The packet header.
	Header Header
	// The PDU header.
	PDU PDU
	// Verification trailer.
	VerificationTrailer VerificationTrailer
	// Security trailer.
	SecurityTrailer SecurityTrailer
	// Body.
	Body *Body
	// Auth token.
	AuthData []byte
	// Raw data.
	raw []byte
	// stub start and stub end.
	start, end int
}

func (p *Packet) IsLastFrag() bool {
	return p.Header.PacketFlags&PacketFlagLastFrag != 0
}

// Bytes function returns the PDU bytes.
func (p *Packet) Bytes() []byte {
	return p.raw
}

func (p *Packet) Len() int {
	return len(p.raw)
}

// HeaderBytes function returns the header bytes.
func (p *Packet) HeaderBytes() []byte {
	return p.raw[:p.start]
}

// StubDataBytes function returns the stub data bytes.
// (will include padding only when request/response is not a case)
func (p *Packet) StubDataBytes() []byte {
	return p.raw[p.start:p.end]
}

// SecurityTrailerBytes function returns the slice pointing to the
// security trailer bytes.
func (p *Packet) SecurityTrailerBytes() []byte {
	if p.Header.AuthLength == 0 {
		return nil
	}
	return p.raw[p.end : p.end+SecurityTrailerSize]
}

// AuthDataBytes function returns the slice pointing to the auth_data
// encoded value.
func (p *Packet) AuthDataBytes() []byte {
	if p.Header.AuthLength == 0 {
		return nil
	}
	return p.raw[p.Header.FragLength-p.Header.AuthLength:]
}

// Body structure represents the read-writer for the chunked
// marshalling/unmarshalling operations.
type Body struct {
	mu sync.Mutex
	// NDR.
	ndr ndr.NDR
	// Operation.
	op Operation
	// Wait chunk.
	chnk *ndr.WaitChunk
	// done.
	done bool
}

func (body *Body) SetDone() {
	body.mu.Lock()
	body.done = true
	body.chnk.Done()
	body.mu.Unlock()
}

func (body *Body) IsDone() bool {
	if body != nil {
		body.mu.Lock()
		done := body.done
		body.mu.Unlock()
		return done
	}
	return true
}

// NewBody function returns the new stub reader (when marshal is `false`), or
// writer, (when marshal is `true`).
func NewBody(ctx context.Context, op Operation, p *Presentation, marshal bool) *Body {

	chnk := ndr.NewWaitChunk()
	body := &Body{chnk: chnk, ndr: p.TransferEncoding()(nil, chnk)}

	// start marshaling/unmarshaling routine.
	go func() {
		// done will indicate the end of marshaling.
		defer body.SetDone()
		if marshal {
			body.ndr.Write(nil) // do nil write.
			op.MarshalNDRRequest(ctx, body.ndr)
		} else {
			body.ndr.Read(nil) // do nil read.
			op.UnmarshalNDRResponse(ctx, body.ndr)
		}
	}()

	return body
}

// DecodeFrom function performs the blocking read from the buffer `b`
// to unmarshaller.
func (body *Body) DecodeFrom(b []byte, frmt ndr.DataRepresentation, maxLen int) (int, error) {

	if body.IsDone() {
		return 0, io.EOF
	}

	n := body.chnk.Wait(b, frmt, maxLen)
	if body.ndr.Err() != nil {
		return n, body.ndr.Err()
	}

	if body.IsDone() {
		return n, io.EOF
	}

	return n, nil
}

// Close function terminates the marshaling/unmarshaling routines
// that haven't yet completed.
func (body *Body) Close() {
	body.chnk.Close()
}

// WriteTo function performs the blocking write from the marshaler to
// buffer `b`.
func (body *Body) EncodeTo(b []byte, frmt ndr.DataRepresentation, maxLen int) (int, error) {

	if body.IsDone() {
		return 0, io.EOF
	}

	n := body.chnk.Wait(b, frmt, maxLen)
	if body.ndr.Err() != nil {
		return n, body.ndr.Err()
	}

	if body.IsDone() {
		return n, io.EOF
	}

	return n, nil
}

// Codec function returns the NDR codec.
func (c *transport) Codec(raw []byte, drep ndr.DataRepresentation) ndr.NDR {
	return ndr.NDR20(raw, drep)
}

// IsSecurePacket function returns `true` if security trailer is/must_be appended
// to the packet.
func (c *transport) IsSecurePacket(ctx context.Context, pkt *Packet) bool {
	return pkt.Header.AuthLength != 0 || (pkt.SecurityTrailer.AuthLevel >= AuthLevelConnect && c.settings.IsSecurityMultiplexed())
}

// DecodeFragmentFromBuffer function decodes the fragment in buffer `raw`, unwraps the contents
// of the stub and returns the packet.
func (c *transport) DecodePacket(ctx context.Context, pkt *Packet, raw []byte) (*Packet, error) {

	raw = raw[:c.settings.MaxRecvFrag]

	r := c.Codec(raw, c.settings.DataRepresentation)

	// unmarshal header.
	if err := pkt.Header.ReadFrom(ctx, r); err != nil {
		return nil, fmt.Errorf("read_header: %v", err)
	}
	// set proper fragment length.
	pkt.raw, pkt.end = raw[:pkt.Header.FragLength], int(pkt.Header.FragLength)
	// adjust auth length.
	if pkt.Header.AuthLength != 0 {
		// adjust stub end to include security trailer.
		pkt.end -= SecurityTrailerSize + int(pkt.Header.AuthLength)
	}

	// unmarshal pdu.
	switch pkt.Header.PacketType {
	case PacketTypeRequest:
		req := new(Request)
		// expect object uuid.
		if pkt.Header.PacketFlags&PacketFlagObjectUUID != 0 {
			// capture object uuid in packet level.
			req.ObjectUUID = new(uuid.UUID)
		}
		pkt.PDU = req
	case PacketTypeResponse:
		pkt.PDU = new(Response)
	case PacketTypeFault:
		pkt.PDU = new(Fault)
	case PacketTypeBind:
		pkt.PDU = new(Bind)
	case PacketTypeBindAck:
		pkt.PDU = new(BindAck)
	case PacketTypeBindNak:
		pkt.PDU = new(BindNak)
	case PacketTypeAlterContext:
		pkt.PDU = new(AlterContext)
	case PacketTypeAlterContextResponse:
		pkt.PDU = new(AlterContextResponse)
	case PacketTypeAuth3:
		pkt.PDU = new(Auth3)
	case PacketTypeShutdown:
		pkt.PDU = new(Shutdown)
	case PacketTypeCancel:
		pkt.PDU = new(Cancel)
	case PacketTypeOrphaned:
		pkt.PDU = new(Orphaned)
	default:
		return nil, fmt.Errorf("read_pdu_header: unknown header type 0x%02x", pkt.Header.PacketType)
	}
	// unmarshal pdu header.
	if err := pkt.PDU.ReadFrom(ctx, r); err != nil {
		return nil, fmt.Errorf("read_pdu_header: %v", err)
	}

	maxLen := 0
	// process fault/shutdown message.
	switch pdu := pkt.PDU.(type) {
	case *Request:
		maxLen = int(pdu.AllocHint)
	case *Response:
		maxLen = int(pdu.AllocHint)
	case *Fault:
		if pdu.Status != 0 {
			return nil, errors.New(ctx, pdu.Status)
		}
		maxLen = int(pdu.AllocHint)
	case *BindNak:
		return pkt, nil
	case *Shutdown:
		return nil, ErrShutdown
	}
	// set proper beginning of the stub data.
	pkt.start = r.Offset()

	// read auth data.
	if c.IsSecurePacket(ctx, pkt) {
		r := c.Codec(pkt.raw[pkt.end:], pkt.Header.PacketDRep)
		// read security trailer.
		if err := pkt.SecurityTrailer.ReadFrom(ctx, r); err != nil {
			return nil, fmt.Errorf("read_security_trailer: %v", err)
		}
		// read auth data.
		pkt.AuthData = r.Bytes()[:pkt.Header.AuthLength]
	}
	// trim auth padding from decrypted stub.
	if pad := int(pkt.SecurityTrailer.AuthPadLength); pkt.end-pad >= pkt.start {
		pkt.end -= pad
	}
	// decode the stub data.
	n, err := pkt.Body.DecodeFrom(pkt.raw[pkt.start:pkt.end], pkt.Header.PacketDRep, maxLen)
	if err != nil {
		if err != io.EOF {
			return nil, fmt.Errorf("decode_stub_data: %v", err)
		}
	}
	// check the trailer.
	if expN := pkt.end - pkt.start; expN != n {
		// find the verification trailer signature.
		idx := bytes.Index(pkt.raw[pkt.start+n:pkt.end], VerificationTrailerSignature[:])
		if idx == -1 {
			return nil, fmt.Errorf("decode_stub_data: unread portion of %d bytes", expN-n)
		}
		// always little-endian.
		r := c.Codec(pkt.raw[pkt.start+n+idx:pkt.end], ndr.DefaultDataRepresentation)
		if err := pkt.VerificationTrailer.ReadFrom(ctx, r); err != nil {
			return nil, fmt.Errorf("decode_verification_trailer: %v", err)
		}
		// adjust packet end to the stub data end.
		pkt.end = pkt.start + n
	}

	return pkt, nil
}

// zeroPad for writing header and auth padding data.
var zeroPad = [16]byte{}

// EncodeFragmentToBuffer function encodes the packet to the buffer `raw`.
func (c *transport) EncodePacket(ctx context.Context, pkt *Packet, raw []byte) error {

	raw = raw[:c.settings.MaxXmitFrag]

	// set initial buffer settings.
	pkt.raw, pkt.end = raw, len(raw)

	// set packet rpc version.
	pkt.Header.RPCVersion, pkt.Header.RPCVersionMinor = 5, 0
	// set packet drep.
	pkt.Header.PacketDRep = c.settings.DataRepresentation
	// set packet type.
	pkt.Header.PacketType = PDUToPacketType(pkt.PDU)
	// set auth_data length.
	pkt.Header.AuthLength = uint16(len(pkt.AuthData))
	// populate verification trailer header2.
	if pdu, ok := pkt.PDU.(*Request); ok {
		for i := range pkt.VerificationTrailer.Commands {
			cmd, ok := pkt.VerificationTrailer.Commands[i].Command.(*VerifyHeader2)
			if !ok {
				continue
			}

			cpy := *cmd

			cpy.PacketDRep = pkt.Header.PacketDRep
			cpy.PacketType = pkt.Header.PacketType
			cpy.CallID = pkt.Header.CallID
			cpy.ContextID = pdu.ContextID
			cpy.OpNum = pdu.OpNum

			pkt.VerificationTrailer.Commands[i] = &VerificationCommand{
				Command:  &cpy,
				Required: pkt.VerificationTrailer.Commands[i].Required,
			}
		}
	}

	// adjust stub buffer to include security trailer.
	if pkt.Header.AuthLength > 0 {
		pkt.end -= MaxPad + SecurityTrailerSize + int(pkt.Header.AuthLength)
	}

	// XXX: verification is computed for every fragment, however it's being
	// written only for last fragment.
	verifyLen := pkt.VerificationTrailer.Size()
	if verifyLen > 0 {
		pkt.end -= VerificationTrailerMaxPad + verifyLen /* VerificationMaxPad */
	}

	w := c.Codec(raw, pkt.Header.PacketDRep)

	// skip over the header. (flags and length will be determined afterwards).
	if _, err := w.Write(zeroPad[:HeaderSize]); err != nil {
		return fmt.Errorf("write_zero_header: %v", err)
	}
	// marshal pdu header.
	if err := pkt.PDU.WriteTo(ctx, w); err != nil {
		return fmt.Errorf("write_pdu_header: %v", err)
	}
	// end of header / start of stub data.
	pkt.start = w.Offset()
	// encode stub data.
	n, err := pkt.Body.EncodeTo(raw[pkt.start:pkt.end], pkt.Header.PacketDRep, 0)
	if err != nil {
		if err != io.EOF {
			return fmt.Errorf("encode_stub_data: %v", err)
		}
		pkt.Header.PacketFlags |= PacketFlagLastFrag
	}

	// adjust stub data end.
	pkt.raw, pkt.end = raw[:pkt.start+n], pkt.start+n

	// verification_trailer.
	if pkt.IsLastFrag() && verifyLen > 0 {
		// XXX: always little-endian.
		w := c.Codec(raw[pkt.end:], ndr.DefaultDataRepresentation)
		// write pad.
		pad := (4 - (pkt.end % 4)) % 4
		if _, err := w.Write(zeroPad[:pad]); err != nil {
			return fmt.Errorf("write_verification_trailer_pad: %v", err)
		}
		// write verification_trailer.
		w = w.WithBytes(raw[pkt.end+pad:])
		if err := pkt.VerificationTrailer.WriteTo(ctx, w); err != nil {
			return fmt.Errorf("write_verification_trailer: %v", err)
		}
		verifyLen = w.Offset()
		// adjust stub end.
		pkt.raw, pkt.end = raw[:pkt.end+pad+verifyLen], pkt.end+pad+verifyLen
	}

	// security_trailer.
	if c.IsSecurePacket(ctx, pkt) {
		w := c.Codec(raw[pkt.end:], pkt.Header.PacketDRep)
		// pad for security trailer. (pkt.end is the current offset).
		pad := (4 - (pkt.end % 4)) % 4
		// pad to avoid cryptographic residue.
		pad += (16 - ((n + pad) % 16)) % 16
		if _, err := w.Write(zeroPad[:pad]); err != nil {
			return fmt.Errorf("write_auth_pad: %v", err)
		}
		// reset writer to align the trailer.
		w = c.Codec(raw[pkt.end+pad:], pkt.Header.PacketDRep)
		// set auth_pad_length.
		pkt.SecurityTrailer.AuthPadLength = uint8(pad)
		// write security trailer.
		if err := pkt.SecurityTrailer.WriteTo(ctx, w); err != nil {
			return fmt.Errorf("write_security_trailer: %v", err)
		}
		// for request/response auth_data will be calculated
		// based on the marshalled packet.
		if _, err := w.Write(pkt.AuthData); err != nil {
			return fmt.Errorf("write_auth_data: %v", err)
		}
		authLen := w.Offset()
		// set proper raw data.
		pkt.raw, pkt.end = raw[:pkt.end+pad+authLen], pkt.end+pad
	}

	// set proper fragment length.
	pkt.Header.FragLength = uint16(len(pkt.raw))
	// set proper header data.
	if err := pkt.Header.WriteTo(ctx, c.Codec(raw, pkt.Header.PacketDRep)); err != nil {
		return fmt.Errorf("write_header: %v", err)
	}

	return nil
}
