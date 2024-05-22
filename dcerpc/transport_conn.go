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
	// wait.
	hdr, ok := <-c.outQ
	if !ok {
		return Header{}, io.ErrUnexpectedEOF
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
	c.inQ <- nil
}

// WriteBuffer function copies the buffer `p` to xmit buffer and notifies
// the transport and waits for transport done with the buffer transmission.
func (c *call) WriteBuffer(ctx context.Context, hdr Header, p []byte) error {
	// copy the buffer to write.
	if !c.noCopy {
		copy(c.xmit, p)
	}
	// ready.
	c.outQ <- hdr
	// wait for response.
	return <-c.inQ
}

// WriteBuffer function writes the data `p` to the wire.
func (c *transport) WriteBuffer(ctx context.Context, hdr Header, p []byte) error {

	ctx, cancel := context.WithTimeout(ctx, c.settings.Timeout)
	defer cancel()

	if int(hdr.FragLength) > c.settings.MaxXmitFrag {
		return ErrPacketTooLong
	}

	p = p[:hdr.FragLength]

	return doWithTimeout(ctx, func() error {
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
	if err := doWithTimeout(ctx, func() error {
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
	if err := doWithTimeout(ctx, func() error {
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
			return t.WithErr(ctx.Err())
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
		}

		// wait for buffer copy.
		newTimer(&deadline, t.settings.Deadline)
		select {
		case <-call.inQ:
		case <-deadline.C:
			return t.WithErr(fmt.Errorf("caller-ready timer expired"))
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
			return t.WithErr(ctx.Err())
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
		hdr := <-call.outQ
		// write buffer.
		err := t.WriteBuffer(ctx, hdr, t.tx)

		// report status back.
		newTimer(&deadline, t.settings.Deadline)
		select {
		case call.inQ <- err:
		case <-deadline.C:
			err = fmt.Errorf("client-ack timer expired")
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
	t.rxQ <- call

	// wait for receive completes.
	if t.settings.Multiplexing {
		<-call.done
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
func doWithTimeout(ctx context.Context, f func() error) error {

	done := make(chan error, 1)
	go func() {
		done <- f()
	}()

loop:
	for {
		select {
		case err := <-done:
			if err != nil {
				return err
			}
			break loop
		case <-ctx.Done():
			if ctx.Err() != nil {
				return ctx.Err()
			}
			return context.DeadlineExceeded
		}
	}
	return nil
}
