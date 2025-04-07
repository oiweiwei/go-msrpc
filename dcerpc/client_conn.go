package dcerpc

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/oiweiwei/go-msrpc/midl/uuid"
	"github.com/rs/zerolog"
)

var (
	ErrConnClosed   = errors.New("connection is closed")
	ErrConnNotExist = errors.New("presentation connection does not exist")
)

// The DCE/RPC client connection.
//
// This connection is returned by group connection `conn` Bind method.
type clientConn struct {
	mu *sync.RWMutex
	// The presentation context.
	presentation *Presentation
	// The security context.
	security *Security
	// The verification trailer.
	verify *VerificationTrailer
	// The communication channel.
	transport *transport
	// The receive/transmit buffers.
	buffer []byte
	// The subconnections (connection negotiated with
	// the same bind instruction).
	subs []*clientConn
	// The flag that indicates whether the connection
	// is closed.
	closed bool
	// Logger.
	logger zerolog.Logger
}

// SubConn interface implements the sub-connection query method
// using presentation context syntax identifier.
type SubConn interface {
	// Conn.
	Conn
	// SubConn function returns the client connection established for the abstract syntax.
	// If the abstract syntax was not found ErrConnNotExist is returned.
	// If the connection is closed ErrConnClosed is returned.
	// If presenetation context negotiation was not successful the sub-connection and
	// presentation context error is returned.
	SubConn(context.Context, *SyntaxID) (Conn, error)
}

// SubConn function returns the client connection established for the abstract syntax.
func (c *clientConn) SubConn(ctx context.Context, abstractSyntax *SyntaxID) (Conn, error) {

	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.isClosed() {
		return nil, ErrConnClosed
	}

	if abstractSyntax != nil {
		for _, sub := range c.subs {
			if *sub.presentation.AbstractSyntax == *abstractSyntax {
				return sub, sub.presentation.Error
			}
		}
	}

	return nil, ErrConnNotExist
}

// Bind function establishes new client connection on the same transport
// as for current client connection.
func (c *clientConn) Bind(ctx context.Context, opts ...Option) (Conn, error) {
	return c.transport.Bind(ctx, opts...)
}

// AlterContext function negotiates the new security context for the client connection.
func (c *clientConn) AlterContext(ctx context.Context, opts ...Option) error {

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.isClosed() {
		return fmt.Errorf("alter connection context: %w", ErrConnClosed)
	}

	for _, sub := range c.subs {
		opts = append(opts, withPresentation(sub.presentation))
	}

	new, err := c.transport.AlterContext(ctx, opts...)
	if err != nil {
		return fmt.Errorf("alter connection context: %w", err)
	}

	for _, sub := range c.subs {
		// replace security context.
		sub.security = new.(*clientConn).security
	}

	return nil
}

func (c *clientConn) Context() context.Context {
	if c.security != nil && c.security.ctx != nil {
		return c.security.ctx
	}

	return context.Background()
}

// Invoke function invokes the operation.
func (c *clientConn) Invoke(ctx context.Context, op Operation, opts ...CallOption) error {

	c.mu.RLock()
	defer c.mu.RUnlock()

	if err := c.invoke(ctx, op, opts...); err != nil {
		return fmt.Errorf("dcerpc: invoke: %s: %w", op.OpName(), err)
	}

	return nil
}

// InvokeObject function invokes the operation with ObjectUUID.
func (c *clientConn) InvokeObject(ctx context.Context, obj *uuid.UUID, op Operation, opts ...CallOption) error {

	c.mu.RLock()
	defer c.mu.RUnlock()

	if err := c.invoke(ctx, op, append(opts, WithObjectUUID(obj))...); err != nil {
		return fmt.Errorf("dcerpc: invoke_object: %s: %s: %w", obj.String(), op.OpName(), err)
	}

	return nil
}

// RegsiterServer: NYI.
func (c *clientConn) RegisterServer(h ServerHandle, opts ...Option) {
	// NYI.
}

// BodyReader function returns the new packet body as a reader for sequential
// request marshaling.
func (c *clientConn) BodyReader(ctx context.Context, op Operation) *Body {
	return NewBody(ctx, op, c.presentation, false)
}

// BodyWriter function returns the new packet body as a writer for sequential
// response unmarshaling.
func (c *clientConn) BodyWriter(ctx context.Context, op Operation) *Body {
	return NewBody(ctx, op, c.presentation, true)
}

// invoke.
func (c *clientConn) invoke(ctx context.Context, op Operation, opts ...CallOption) error {

	if c.isClosed() {
		return ErrConnClosed
	}

	if c.presentation.Error != nil {
		return c.presentation.Error
	}

	obj, _ := HasObjectUUID(opts)

	call, err := c.transport.MakeCall(ctx)
	if err != nil {
		return err
	}

	call.Lock()
	defer call.Unlock()

	c.logger.Debug().Uint32("call_id", call.ID()).Interface("in", op).Msg("operation input")

	pkt := &Packet{
		Header: Header{
			PacketFlags: PacketFlagFirstFrag | FlagObjectUUID(obj),
		},
		PDU: &Request{
			ContextID:  uint16(c.presentation.ID()),
			OpNum:      uint16(op.OpNum()),
			ObjectUUID: obj,
		},
		VerificationTrailer: c.verify.VerificationTrailer(),
		SecurityTrailer:     c.security.SecurityTrailer(),
	}

	bodyWriter := c.BodyWriter(ctx, op)
	defer bodyWriter.Close()

	for pkt.Body = bodyWriter; !pkt.IsLastFrag(); {
		// allocate auth_data.
		pkt.AuthData = make([]byte, c.security.AuthLength(ctx, pkt))
		// encode packet fragment.
		if err = c.WritePacket(ctx, call, pkt); err != nil {
			return fmt.Errorf("request: %w", err)
		}
		// clear the first frag.
		pkt.Header.PacketFlags &= ^PacketFlagFirstFrag
	}

	// clear last frag flag to start the for-loop.
	pkt.Header.PacketFlags &= ^PacketFlagLastFrag

	bodyReader := c.BodyReader(ctx, op)
	defer bodyReader.Close()

	for pkt.Body = bodyReader; !pkt.IsLastFrag(); {
		// decode packet fragment.
		if pkt, err = c.ReadPacket(ctx, call, pkt); err != nil {
			return fmt.Errorf("response: %w", err)
		}
	}

	c.logger.Debug().Uint32("call_id", call.ID()).Interface("out", op).Msg("operation output")

	return nil
}

// Wrap function wraps the raw bytes with security service.
func (c *clientConn) Wrap(ctx context.Context, hdr Header, raw []byte, call Call) error {

	pkt := &Packet{Header: hdr, raw: raw[:hdr.FragLength]}

	if pkt.start = RequestSize; pkt.Header.PacketFlags&PacketFlagObjectUUID != 0 {
		pkt.start += ObjectUUIDSize
	}

	if pkt.end = int(pkt.Header.FragLength); pkt.Header.AuthLength != 0 {
		pkt.end -= int(pkt.Header.AuthLength) + SecurityTrailerSize
	}

	if c.security.CanWrap(ctx, pkt) {
		if err := c.security.Wrap(ctx, pkt); err != nil {
			return err
		}
	}

	return nil
}

// Unwrap function unwraps the encrypted/signed raw bytes.
func (c *clientConn) Unwrap(ctx context.Context, hdr Header, raw []byte, call Call) error {

	// indicate ready when security context will be locked.
	// XXX: set the ready flag once we acquire the security context lock.
	// (see afterLock parameter in Unwrap function).
	// defer call.Ready(ctx)

	pkt := &Packet{Header: hdr, raw: raw[:hdr.FragLength]}

	if pkt.start = RequestSize; pkt.Header.PacketFlags&PacketFlagObjectUUID != 0 {
		pkt.start += ObjectUUIDSize /* uuid size */
	}

	if pkt.end = int(pkt.Header.FragLength); pkt.Header.AuthLength != 0 {
		pkt.end -= int(pkt.Header.AuthLength) + SecurityTrailerSize
	}

	if c.security.CanWrap(ctx, pkt) {
		if err := c.security.Unwrap(ctx, pkt, call.Ready); err != nil {
			return err
		}
	} else {
		call.Ready(ctx)
	}

	return nil

}

// WritePacket function encodes, encrypts/signs and sends the packet to the server.
func (c *clientConn) WritePacket(ctx context.Context, call Call, pkt *Packet) error {
	if err := c.writePacket(ctx, call, pkt); err != nil {
		if terr := c.transport.HasErr(); terr != nil {
			err = terr
		}
		// close transport on error.
		c.transport.Close(ctx)
		return err
	}
	return nil
}

func (c *clientConn) writePacket(ctx context.Context, call Call, pkt *Packet) error {

	pkt.Header.CallID = call.ID()

	c.logger.Debug().EmbedObject(pkt.Header).EmbedObject(pkt.PDU).Msg("writing packet")

	if err := c.transport.EncodePacket(ctx, pkt, c.buffer); err != nil {
		return fmt.Errorf("encode packet: %w", err)
	}

	if err := c.Wrap(ctx, pkt.Header, c.buffer, call); err != nil {
		return fmt.Errorf("wrap packet: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, c.transport.settings.Timeout)
	defer cancel()

	if err := call.WriteBuffer(ctx, pkt.Header, c.buffer); err != nil {
		return fmt.Errorf("write buffer: %w", err)
	}

	c.logger.Debug().EmbedObject(pkt.Header).EmbedObject(pkt.PDU).Msg("writing packet done")

	return nil
}

// ReadPacket function reads, decrypts/verifies signature, and decodes the
// packet retrieved from the server.
func (c *clientConn) ReadPacket(ctx context.Context, call Call, pkt *Packet) (*Packet, error) {

	pkt, err := c.readPacket(ctx, call, pkt)
	if err != nil {
		if terr := c.transport.HasErr(); terr != nil {
			err = terr
		}
		// close transport on error.
		c.transport.Close(ctx)
		return nil, err
	}

	return pkt, nil
}

// readPacket.
func (c *clientConn) readPacket(ctx context.Context, call Call, pkt *Packet) (*Packet, error) {

	ctx, cancel := context.WithTimeout(ctx, c.transport.settings.Timeout)
	defer cancel()

	hdr, err := call.ReadBuffer(ctx, c.buffer)
	if err != nil {
		return nil, fmt.Errorf("read buffer: %w", err)
	}

	c.logger.Debug().EmbedObject(hdr).Msg("reading packet")

	if err := c.Unwrap(ctx, hdr, c.buffer, call); err != nil {
		return nil, fmt.Errorf("unwrap packet: %w", err)
	}

	if pkt, err = c.transport.DecodePacket(ctx, pkt, c.buffer); err != nil {
		return nil, fmt.Errorf("decode packet: %w", err)
	}

	c.logger.Debug().EmbedObject(pkt.Header).EmbedObject(pkt.PDU).Msg("reading packet done")

	return pkt, nil
}

// isClosed.
func (c *clientConn) isClosed() bool {
	return c.closed
}

// Close function closes the client connection and underlying transport.
func (c *clientConn) Close(ctx context.Context) error {

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.isClosed() {
		return fmt.Errorf("connection has already been closed")
	}

	for _, sub := range c.subs {
		// mark all sub-connections as closed.
		sub.closed = true
	}

	// close the transport, this will shut down the socket/named pipe
	// and remove the transport from the list of active transports
	// of the group conn.
	return c.transport.Close(ctx)
}
