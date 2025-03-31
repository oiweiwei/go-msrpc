package dcerpc

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/rs/zerolog"

	"github.com/oiweiwei/go-msrpc/smb2"
	"github.com/oiweiwei/go-msrpc/ssp/gssapi"
)

var (
	// The transport was not binded.
	ErrNotBinded = errors.New("not binded")
	// The transport is closed.
	ErrClosed = errors.New("transport is closed")
)

// transport represents the single TCP/IP connection or instance
// of the named pipe.
type transport struct {
	// The mutex for close connection.
	mu sync.RWMutex
	// The exclusive access to the caller channel.
	// (needed for alter_context/bind operation).
	callMu sync.RWMutex
	// The transport identifier (smb connection port, or tcp/ip port).
	id int
	// binded flag indicates whether the connection has already
	// been binded. (If this flag is `true`, only alter context
	// queries are allowed).
	binded bool
	// The raw connection.
	cc *BufferedConn
	// The next call identifier.
	cid atomic.Uint32
	// The connection settings.
	settings *Transport
	// The transmit, receive buffers.
	tx, rx []byte
	// The channel for the callers.
	txQ, rxQ chan *call
	// logger.
	logger zerolog.Logger
	// The transport error.
	err error
	// The close function and wait group.
	close     func()
	closeWait *sync.WaitGroup
	// The transport connection.
	conn *conn
}

func (t *transport) IsBinded() bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.binded
}

func (t *transport) Binded() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.binded = true
}

func (t *transport) HasErr() error {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.err
}

func (t *transport) WithErr(err error) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.err != nil {
		return t.err
	}
	t.err = err
	return t.err
}

// do not copy data to the buffer.
type noCopy struct{}

func (t *transport) CallLock() {
	t.callMu.RLock()
}

func (t *transport) CallUnlock() {
	t.callMu.RUnlock()
}

// MakeCall function acquires the read/write access for the transport.
func (t *transport) MakeCall(ctx context.Context, opts ...any) (Call, error) {

	if err := t.HasErr(); err != nil {
		return nil, err
	}

	call := &call{
		ctx:  ctx,
		id:   t.CallID(),
		xmit: t.tx,
		recv: t.rx,
		inQ:  make(chan error),
		outQ: make(chan Header),
		done: make(chan struct{}),
	}

	for _, opt := range opts {
		switch opt.(type) {
		case noCopy:
			call.noCopy = true
		}
	}

	if t.IsBinded() {
		t.txQ <- call
	}

	return call, nil
}

func (c *transport) CallID() uint32 {
	return c.cid.Add(1)
}

// ExportSMBSecurity exports SMB keying material for security context.
// refs:
//   - https://www.snia.org/sites/default/files/SDC/2016/presentations/smb/Stefan_Metzmacher_Improving_DCERPC_Security.pdf
//   - https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-smb2/901ae284-31d3-4ea1-ae8a-766fc8bfe00e
func (c *transport) ExportSMBSecurity(o *Security) {

	if o == nil {
		return
	}

	if pipe, ok := c.cc.RawConn.(*smb2.NamedPipe); ok {
		o.SetAttribute(gssapi.AttributeSMBApplicationKey, pipe.ApplicationKey())
		o.SetAttribute(gssapi.AttributeSMBSessionKey, pipe.SessionKey())
		if o.Level > AuthLevelConnect {
			o.SetAttribute(gssapi.AttributeSMBEffectiveSessionKey, []byte("SystemLibraryDTC"))
		} else if pipe.Dialect() >= smb2.SMB300 {
			o.SetAttribute(gssapi.AttributeSMBEffectiveSessionKey, pipe.ApplicationKey())
		} else {
			o.SetAttribute(gssapi.AttributeSMBEffectiveSessionKey, pipe.SessionKey())
		}
	}

}

// AlterContext function establishes new presentation or security (or both) context(s).
func (c *transport) AlterContext(ctx context.Context, opts ...Option) (Conn, error) {

	if err := c.HasErr(); err != nil {
		return nil, fmt.Errorf("alter context: %w", err)
	}

	o, err := ParseOptions(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("alter context: parse options: %w", err)
	}

	c.ExportSMBSecurity(o.Security)

	c.callMu.Lock()
	defer c.callMu.Unlock()

	call, err := c.MakeCall(ctx, noCopy{})
	if err != nil {
		return nil, fmt.Errorf("alter context: allocate call: %w", err)
	}

	contexts := PresentationsToContextList(o.Presentations, o.TransferSyntaxes)

	pkt := &Packet{
		Header: Header{
			PacketFlags: PacketFlagFirstFrag | PacketFlagLastFrag | o.Security.RequestHeaderSign,
		},
		PDU:             &AlterContext{ContextList: contexts},
		SecurityTrailer: o.Security.SecurityTrailer(),
	}

	if o.Security == nil || o.Security.Established() {
		pkt.SecurityTrailer = SecurityTrailer{}
	}

	// set auth data.
	if pkt.AuthData, err = o.Security.Init(ctx, nil); err != nil {
		return nil, fmt.Errorf("alter context: init security: %w", err)
	}
	// write bind pdu.
	if err = c.WritePacket(ctx, call, pkt); err != nil {
		return nil, fmt.Errorf("alter context: write packet: %w", err)
	}
	// read bind response (bind-ack, bind-nak).
	if pkt, err = c.ReadPacket(ctx, call, pkt); err != nil {
		return nil, fmt.Errorf("alter context: read packet: %w", err)
	}

	pdu, ok := pkt.PDU.(*AlterContextResponse)
	if !ok {
		return nil, fmt.Errorf("alter context: unexpected response: %s", pkt.Header.PacketType)
	}

	o.Security.SignHeader = pkt.Header.PacketFlags.IsSet(PacketFlagSupportHeaderSign)

	PresentationFromContextList(o.Presentations, pdu.ResultList)

	// alter context until the security context is established.
	if err := c.CompleteContext(ctx, contexts, o.Security, pkt.AuthData); err != nil {
		return nil, err
	}

	if o.IsNewSecurity && o.Security.Level >= AuthLevelConnect {
		// increment security context count for multiplexing.
		c.settings.SecurityContextCount++
	}

	return c.makeConn(o), nil
}

// makeVerify function constructs the verification trailer for the connection.
func (c *transport) makeVerify(i int, o *option) *VerificationTrailer {

	out := make([]*VerificationCommand, 0, len(o.Verify))

	for _, verify := range o.Verify {
		switch verify.Command.(type) {
		case *VerifyPresentation:
			out = append(out, &VerificationCommand{
				Command: &VerifyPresentation{
					InterfaceID:    o.Presentations[i].AbstractSyntax,
					TransferSyntax: o.Presentations[i].TransferSyntax,
				},
				Required: verify.Required,
			})
		case VerifyBitMask:
			bitMask := VerifyBitMask(0)
			if o.Security.SignHeader {
				bitMask |= VerifyBitMaskSupportHeaderSign
			}
			out = append(out, &VerificationCommand{
				Command:  bitMask,
				Required: verify.Required,
			})
		case VerifyHeader2:
			out = append(out, &VerificationCommand{
				Command:  &VerifyHeader2{},
				Required: verify.Required,
			})
		default:
			out = append(out, verify)
		}
	}

	return &VerificationTrailer{Commands: out}
}

// makeConn function constructs the client connection for each
// presentation context and returns the first presentation context
// connection as a client connection. The rest client connections
// can be accessed via SubConns method.
func (c *transport) makeConn(o *option) *clientConn {

	conns, mu := make([]*clientConn, len(o.Presentations)), new(sync.RWMutex)

	for i := range conns {
		conns[i] = &clientConn{
			mu:           mu,
			transport:    c,
			buffer:       make([]byte, c.settings.FragmentSize()),
			security:     o.Security,
			verify:       c.makeVerify(i, o),
			presentation: o.Presentations[i],
			subs:         conns,
			logger:       o.Logger,
		}
	}

	return conns[0]
}

// alter context until the security context is established.
func (c *transport) CompleteContext(ctx context.Context, p []*Context, o *Security, authData []byte) error {

	var err error

	pkt := &Packet{AuthData: authData}

	for !o.Established() {
		// alter context until the security context is established.
		pkt = &Packet{
			Header: Header{
				PacketFlags: PacketFlagFirstFrag | PacketFlagLastFrag | o.RequestHeaderSign,
			},
			PDU:             &AlterContext{ContextList: p},
			SecurityTrailer: o.SecurityTrailer(),
			AuthData:        pkt.AuthData,
		}
		// intitialize the context.
		if pkt.AuthData, err = o.Init(ctx, pkt.AuthData); err != nil {
			return fmt.Errorf("alter context: init security context: %w", err)
		}
		// context has been successfully established.
		if len(pkt.AuthData) == 0 && o.Established() {
			break
		}

		call, err := c.MakeCall(ctx, noCopy{})
		if err != nil {
			return fmt.Errorf("alter context: allocate channel: %w", err)
		}

		if o.Type.Legs() == LegsOdd {
			// replace type with auth3.
			pkt.PDU = &Auth3{}
			// write auth3 pdu.
			if err = c.WritePacket(ctx, call, pkt); err != nil {
				return fmt.Errorf("alter context: auth3: write packet: %w", err)
			}
			// no response is assumed.
			break
		}
		// write alter_context request.
		if err = c.WritePacket(ctx, call, pkt); err != nil {
			return fmt.Errorf("alter context: write packet: %w", err)
		}
		// read alter_context response.
		if pkt, err = c.ReadPacket(ctx, call, pkt); err != nil {
			return fmt.Errorf("alter context: read packet: %w", err)
		}
		// check response.
		if _, ok := pkt.PDU.(*AlterContextResponse); !ok {
			return fmt.Errorf("alter context: unexpected response: %s", pkt.Header.PacketType)
		}

		o.SignHeader = pkt.Header.PacketFlags.IsSet(PacketFlagSupportHeaderSign)
	}

	return nil
}

func (c *transport) Bind(ctx context.Context, opts ...Option) (Conn, error) {

	if err := c.HasErr(); err != nil {
		return nil, fmt.Errorf("bind: %w", err)
	}

	if conn, ok := HasNoBind(opts); ok {
		return conn, nil
	}

	if c.IsBinded() {
		return c.AlterContext(ctx, opts...)
	}

	o, err := ParseOptions(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("bind: parse options: %w", err)
	}

	c.ExportSMBSecurity(o.Security)

	c.logger = o.Logger

	c.callMu.Lock()
	defer c.callMu.Unlock()

	call, err := c.MakeCall(ctx, noCopy{})
	if err != nil {
		return nil, fmt.Errorf("bind: allocate channel: %w", err)
	}

	// set/override the settings group id if association is non-zero.
	c.settings.GroupID = o.Group.SetID(c.settings.GroupID)

	contexts := PresentationsToContextList(o.Presentations, o.TransferSyntaxes)
	if len(contexts) > 0 {
		contexts = append(contexts, &Context{
			AbstractSyntax:   contexts[len(contexts)-1].AbstractSyntax,
			TransferSyntaxes: []*SyntaxID{BindFeatureSyntaxV1_0},
		})
	}

	pkt := &Packet{
		Header: Header{
			PacketFlags: PacketFlagFirstFrag | PacketFlagLastFrag | PacketFlagConcMPX | o.Security.RequestHeaderSign,
		},
		PDU: &Bind{
			MaxXmitFrag:  uint16(c.settings.MaxXmitFrag),
			MaxRecvFrag:  uint16(c.settings.MaxRecvFrag),
			AssocGroupID: uint32(c.settings.GroupID),
			ContextList:  contexts,
		},
		SecurityTrailer: o.Security.SecurityTrailer(),
	}
	// set auth data.
	if pkt.AuthData, err = o.Security.Init(ctx, nil); err != nil {
		return nil, fmt.Errorf("bind: %w", err)
	}
	// write bind pdu.
	if err = c.WritePacket(ctx, call, pkt); err != nil {
		return nil, fmt.Errorf("bind: write packet: %w", err)
	}
	// read bind response (bind-ack, bind-nak).
	if pkt, err = c.ReadPacket(ctx, call, pkt); err != nil {
		return nil, fmt.Errorf("bind: read packet: %w", err)
	}

	switch pdu := (interface{})(pkt.PDU).(type) {

	case *BindAck:

		if c.settings.MaxRecvFrag != int(pdu.MaxRecvFrag) {
			// reset buffered connector.
			c.cc = c.cc.Resized(int(pdu.MaxRecvFrag))
		}

		sz := c.settings.FragmentSize()

		// save retrieved parameters.
		c.settings.MaxRecvFrag = int(pdu.MaxRecvFrag)
		c.settings.MaxXmitFrag = int(pdu.MaxXmitFrag)
		c.settings.GroupID = int(pdu.AssocGroupID)
		c.settings.SecondaryAddr = pdu.PortSpec

		// set the group id for the association.
		o.Group.SetID(int(pdu.AssocGroupID))

		if sz != c.settings.FragmentSize() {
			c.tx, c.rx = make([]byte, c.settings.FragmentSize()), make([]byte, c.settings.FragmentSize())
		}

		// save negotiated header sign parameter.
		o.Security.SignHeader = pkt.Header.PacketFlags.IsSet(PacketFlagSupportHeaderSign)
		c.settings.Multiplexing = pkt.Header.PacketFlags.IsSet(PacketFlagConcMPX)

		feature := PresentationFromContextList(o.Presentations, pdu.ResultList)
		c.settings.KeepConnOpenOnOrphaned = feature.KeepConnOpenOnOrphaned()
		c.settings.SecurityContextMultiplexing = feature.SecurityContextMultiplexing()

		c.logger.Debug().EmbedObject(c.settings).Msg("negotiated_features")

	case *BindNak:
		return nil, c.asyncClose(ctx, fmt.Errorf("bind: %w", pdu))
	default:
		return nil, c.asyncClose(ctx, fmt.Errorf("bind: unexpected response: %s", pkt.Header.PacketType))
	}

	if err := c.CompleteContext(ctx, contexts, o.Security, pkt.AuthData); err != nil {
		return nil, c.asyncClose(ctx, fmt.Errorf("bind: %w", err))
	}

	c.Binded()

	if o.IsNewSecurity && o.Security.Level >= AuthLevelConnect {
		// increment security context count for multiplexing.
		c.settings.SecurityContextCount++
	}

	ctx, c.close = context.WithCancel(ctx)
	c.closeWait = new(sync.WaitGroup)

	// run receiver.
	c.closeWait.Add(1)
	go func() {
		defer c.closeWait.Done()
		if err := c.recvLoop(ctx); err != nil {
			c.WithErr(err)
			c.logger.Info().Err(err).Msg("receiver terminated")
		}
	}()

	// run sender.
	c.closeWait.Add(1)
	go func() {
		defer c.closeWait.Done()
		if err := c.sendLoop(ctx); err != nil {
			c.WithErr(err)
			c.logger.Info().Err(err).Msg("sender terminated")
		}
	}()

	return c.makeConn(o), nil
}

// WritePacket function encodes and writes packet to the connection.
func (t *transport) WritePacket(ctx context.Context, call Call, pkt *Packet) error {

	if err := t.writePacket(ctx, call, pkt); err != nil {
		return t.asyncClose(ctx, err)
	}

	return nil
}

// writePacket.
func (t *transport) writePacket(ctx context.Context, call Call, pkt *Packet) error {

	pkt.Header.CallID = call.ID()

	// encode the packet to the buffer.
	if err := t.EncodePacket(ctx, pkt, t.tx); err != nil {
		return fmt.Errorf("encode packet: %w", err)
	}

	if t.IsBinded() {
		// use call to access the buffer.
		if err := call.WriteBuffer(ctx, pkt.Header, t.tx); err != nil {
			return fmt.Errorf("write buffer: %w", err)
		}
	} else {
		// write data to the buffer directly.
		if err := t.WriteBuffer(ctx, pkt.Header, t.tx); err != nil {
			return fmt.Errorf("write buffer: %w", err)
		}
	}

	t.logger.Debug().EmbedObject(pkt.Header).EmbedObject(pkt.PDU).Msg("write_packet")

	return nil
}

// ReadPacket function reads and decodes the packet from the connection.
func (t *transport) ReadPacket(ctx context.Context, call Call, pkt *Packet) (*Packet, error) {

	var err error

	pkt, err = t.readPacket(ctx, call, pkt)
	if err != nil {
		return nil, t.asyncClose(ctx, err)
	}

	return pkt, nil
}

// readPacket.
func (t *transport) readPacket(ctx context.Context, call Call, pkt *Packet) (*Packet, error) {

	var err error

	if t.IsBinded() {
		// use call to access the readers channel.
		if _, err := call.ReadBuffer(ctx, t.rx); err != nil {
			return nil, fmt.Errorf("read buffer: %w", err)
		}
		// unlock the buffer for the next read.
		defer call.Ready(ctx)
	} else {
		// read the data directly from the connection.
		if _, err := t.ReadBuffer(ctx, t.rx); err != nil {
			return nil, fmt.Errorf("read buffer: %w", err)
		}
	}

	// decode the retrieved packet.
	if pkt, err = t.DecodePacket(ctx, pkt, t.rx); err != nil {
		return nil, fmt.Errorf("decode packet: %w", err)
	}

	t.logger.Debug().EmbedObject(pkt.Header).EmbedObject(pkt.PDU).Msg("read packet done")

	return pkt, nil
}

func (t *transport) shutdown(ctx context.Context) error {

	if !t.IsBinded() {
		return nil
	}

	t.logger.Debug().Msg("closing sender/receiver loops")
	// close sender/receiver.
	t.close()
	// wait for sender/receiver done.
	t.closeWait.Wait()
	t.logger.Debug().Msg("closing sender/receiver loops done")

	t.logger.Debug().Msg("closing connection")

	// close underlying connection.
	if err := t.cc.Close(); err != nil {
		t.logger.Error().Err(err).Msg("close connection error")
		return err
	}

	return nil
}

func (t *transport) asyncClose(ctx context.Context, err error) error {
	if err != nil {
		go func() { t.logger.Err(t.Close(ctx)).Msg("transport is closing") }()
	}
	return err
}

// Close function closes the transport's underlying connection and removes
// the transport from the conn's active transports.
func (t *transport) Close(ctx context.Context) error {

	if err := t.HasErr(); err != nil && err == ErrClosed {
		return err
	}

	// set closed whatever it takes.
	defer t.WithErr(ErrClosed)

	// remove the transport from the list of active transports.
	if err := t.conn.closeTransport(ctx, t); err != nil {
		t.logger.Error().Err(err).Msg("close transport error")
	}

	if err := t.shutdown(ctx); err != nil {
		t.logger.Error().Err(err).Msg("close transport error")
	}

	t.logger.Debug().Msg("closed")

	// done.
	return nil
}
