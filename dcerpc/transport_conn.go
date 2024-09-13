package dcerpc

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"
)

var (
	// Too long packet error.
	ErrPacketTooLong = errors.New("packet is too long")
	// Too small buffer error.
	ErrBufferTooSmall = errors.New("buffer is too small")
)

// Call interface provides the exclusive access to the transport.
type Call interface {
	// The call identifier.
	ID() uint32
	// The function must be called to indicate that transport
	// buffer can be released.
	Ready(context.Context)
	// ReadBuffer function copies the data into the buffer.
	ReadBuffer(context.Context, []byte) (Header, error)
	// WriteBuffer function writes the data from the buffer.
	WriteBuffer(context.Context, Header, []byte) error
}

// call implements the Call interface.
type call struct {
	// The call context.
	ctx context.Context
	// The call identifier.
	id uint32
	// The transmit/receive buffer for the call.
	// (transport's transmit/receive buffers).
	xmit, recv []byte
	// Input/output channels for requests/indications.
	inQ  chan error
	outQ chan Header
	// The flag that indicates whether to perform copy
	// from xmit/recv buffers to connection buffer.
	noCopy bool
	done   chan struct{}
}

// ID returns the call identifier.
func (c *call) ID() uint32 {
	return c.id
}

// ReadBuffer function waits for the transport to retrieve the full fragment,
// then copy this data to `p`. Client must indicate that transport can proceed
// with the next fragment by calling `Ready`. (this is done to acquire security
// lock before next request will be able to use it).
func (c *call) ReadBuffer(ctx context.Context, p []byte) (Header, error) {
	var hdr Header
	var ok bool

	// wait.
	select {
	case hdr, ok = <-c.outQ:
		if !ok {
			return Header{}, io.ErrUnexpectedEOF
		}
	case <-ctx.Done():
		return Header{}, ctx.Err()
	}
	if !c.noCopy {
		copy(p, c.recv)
	}
	// done. (separate call).
	// c.inQ <- nil
	return hdr, nil
}

// Ready function is used to lock the call until particular action from
// the reader will be performed.
func (c *call) Ready(ctx context.Context) {
	// done.
	select {
	case c.inQ <- nil:
	case <-ctx.Done():
	}
}

// WriteBuffer function copies the buffer `p` to xmit buffer and notifies
// the transport and waits for transport done with the buffer transmission.
func (c *call) WriteBuffer(ctx context.Context, hdr Header, p []byte) error {
	// copy the buffer to write.
	if !c.noCopy {
		copy(c.xmit, p)
	}
	// ready.
	select {
	case c.outQ <- hdr:
	case <-ctx.Done():
		return ctx.Err()
	}
	// wait for response.
	select {
	case err := <-c.inQ:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// WriteBuffer function writes the data `p` to the wire.
func (c *transport) WriteBuffer(ctx context.Context, hdr Header, p []byte) error {

	if int(hdr.FragLength) > c.settings.MaxXmitFrag {
		return ErrPacketTooLong
	}

	p = p[:hdr.FragLength]

	return doWithTimeout(ctx, c.settings.Timeout, func() error {
		for n := 0; n < int(hdr.FragLength); {
			actual, err := c.cc.Write(p[n:])
			if err != nil {
				return err
			}
			n += actual
		}
		return nil
	})
}

// ReadBuffer function reads the bytes from the wire into the buffer `p`.
func (c *transport) ReadBuffer(ctx context.Context, p []byte) (Header, error) {

	var hdr Header

	ctx, cancel := context.WithTimeout(ctx, c.settings.Timeout)
	defer cancel()

	if len(p) < c.settings.MaxXmitFrag {
		return hdr, ErrBufferTooSmall
	}

	p = p[:c.settings.MaxRecvFrag]

	// read header.
	if err := doWithTimeout(ctx, c.settings.Timeout, func() error {
		for n := 0; n < HeaderSize; {
			actual, err := c.cc.Read(p[n:HeaderSize])
			if err != nil {
				return err
			}
			n += actual
		}
		return nil
	}); err != nil {
		return hdr, err
	}
	// unmarshal header.
	if err := (&hdr).ReadFrom(ctx, c.Codec(p[:HeaderSize], c.settings.DataRepresentation)); err != nil {
		return hdr, err
	}
	// check fragment length.
	if int(hdr.FragLength) > len(p) {
		return hdr, ErrPacketTooLong
	}
	// read remaining fragment.
	if err := doWithTimeout(ctx, c.settings.Timeout, func() error {
		for n := HeaderSize; n < int(hdr.FragLength); {
			actual, err := c.cc.Read(p[n:hdr.FragLength])
			if err != nil {
				return err
			}
			n += actual
		}
		return nil
	}); err != nil {
		return hdr, err
	}

	return hdr, nil
}

// recvLoop.
func (t *transport) recvLoop(ctx context.Context) error {

	t.logger.Debug().Msg("started receiver routine")

	for {
		select {
		case call := <-t.rxQ:
			t.logger.Debug().Uint32("call_id", call.ID()).Msg("serving response")
			if err := t.recv(ctx, call); err != nil {
				t.logger.Error().Uint32("call_id", call.ID()).Err(err).Msg("serving response error")
			}
		case <-ctx.Done():
			return nil
		}
	}
}

// recv.
func (t *transport) recv(ctx context.Context, call *call) error {

	if t.settings.Multiplexing {
		defer close(call.done)
	}

	// close output query for preventing the deadlock.
	defer close(call.outQ)

	if err := t.HasErr(); err != nil {
		return err
	}

	var deadline *time.Timer
	defer clearTimer(&deadline)

	for {
		// read packet from buffer.
		hdr, err := t.ReadBuffer(ctx, t.rx)
		if err != nil {
			t.logger.Error().Err(err).Msg("receiver: read buffer error")
			// critical error.
			return t.WithErr(err)
		}

		if call.id != hdr.CallID {
			t.logger.Error().Msgf("receiver: call_id %d was not found", hdr.CallID)
			// discard packet for unknown call id.
			continue
		}

		// indicate ready to copy the buffer.
		newTimer(&deadline, t.settings.Deadline)
		select {
		case call.outQ <- hdr:
		case <-deadline.C:
			return t.WithErr(fmt.Errorf("caller-receiver timer expired"))
		case <-ctx.Done():
			return nil
		}

		// wait for buffer copy.
		newTimer(&deadline, t.settings.Deadline)
		select {
		case <-call.inQ:
		case <-deadline.C:
			return t.WithErr(fmt.Errorf("caller-ready timer expired"))
		case <-ctx.Done():
			return nil
		}

		// remove caller from the wait queue.
		if hdr.PacketFlags.IsSet(PacketFlagLastFrag) {
			break
		}
	}
	return nil
}

// sendLoop.
func (t *transport) sendLoop(ctx context.Context) error {

	t.logger.Debug().Msg("started sender routine")

	for {

		select {
		case call := <-t.txQ:
			t.logger.Debug().Uint32("call_id", call.ID()).Msg("serving call")
			if err := t.send(ctx, call); err != nil {

				t.logger.Error().Uint32("call_id", call.ID()).Err(err).Msg("serving call error")
			}
		case <-ctx.Done():
			return nil
		}
	}
}

// send.
func (t *transport) send(ctx context.Context, call *call) error {

	var deadline *time.Timer
	defer clearTimer(&deadline)

	if err := t.HasErr(); err != nil {
		return err
	}

	for {
		// wait for write done.
		var hdr Header
		select {
		case hdr = <-call.outQ:
		case <-ctx.Done():
			return nil
		}
		// write buffer.
		err := t.WriteBuffer(ctx, hdr, t.tx)

		// report status back.
		newTimer(&deadline, t.settings.Deadline)
		select {
		case call.inQ <- err:
		case <-deadline.C:
			err = fmt.Errorf("client-ack timer expired")
		case <-ctx.Done():
			return nil
		}

		if err := t.WithErr(err); err != nil {
			return err
		}

		// finish send.
		if hdr.PacketFlags.IsSet(PacketFlagLastFrag) {
			t.logger.Debug().Uint32("call_id", hdr.CallID).Msg("send is done")
			if hdr.PacketType == PacketTypeAuth3 {
				return nil
			}
			break
		}
	}

	// publish to receiver queue.
	select {
	case t.rxQ <- call:
	case <-ctx.Done():
	}

	// wait for receive completes.
	if t.settings.Multiplexing {
		select {
		case <-call.done:
		case <-ctx.Done():
		}
	}

	return nil
}

// newTimer.
func newTimer(t **time.Timer, dur time.Duration) {
	if *t == nil {
		*t = time.NewTimer(dur)
		return
	}
	if !(*t).Stop() {
		<-((*t).C)
	}
	(*t).Reset(dur)
}

// clearTimer.
func clearTimer(t **time.Timer) {
	if *t != nil && !(*t).Stop() {
		<-((*t).C)
	}
}

// doWithTimeout.
func doWithTimeout(ctx context.Context, timout time.Duration, f func() error) error {

	done := make(chan error, 1)
	go func() {
		done <- f()
	}()

	timer := time.NewTimer(timout)

	var err error

	select {
	case err = <-done:
	case <-timer.C:
		err = context.DeadlineExceeded
	case <-ctx.Done():
		err = ctx.Err()
	}

	if !timer.Stop() {
		<-timer.C
	}

	return err
}
