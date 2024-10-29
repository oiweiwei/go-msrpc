package ndr

import (
	"encoding/binary"
	"io"
	go_math "math"
	"sync"

	"github.com/oiweiwei/go-msrpc/ndr/math"
)

// ChunkedBuffer interface is aimed to provide an interface
// for reading the data input chunk by chunk thus enabling
// capabilities for reusing the same buffer within single
// network exchange involving multiple chunks.
type ChunkedBuffer interface {
	EOF() bool
	// Read/Write interface.
	io.ReadWriter
	// Len function must return the total length of the buffer.
	Len() int
	// ReadRepresentation function must read and configure the
	// representation for the current chunk.
	ReadRepresentation(*DataRepresentation) error
	// WriteRepresentation function must write and configure
	// the representation for the current chunk.
	WriteRepresentation(DataRepresentation) error
	// Order function must return the byte order for the chunk.
	Order() binary.ByteOrder
	// Float function must return the floating-point format for
	// the chunk.
	Float() math.FloatFormat
	// Bytes function must return the bytes in the current chunk.
	Bytes() []byte
}

// AlignBuffer provides the alignment capabilities over
// the chunked buffer. Since chunked buffer can contain
// multiple chunks in the implementation, we need to keep
// track of current position in the stream.
type AlignBuffer interface {
	// The chunked buffer interface.
	ChunkedBuffer
	// EOF function must return true if underlying chunk buffer
	// length is zero. (note that it must be used only with chunks
	// that contain of a single buffer).
	EOF() bool
	// Pos must contain the current position of the buffer over
	// multiple chunks. (can be used to set proper pointer values).
	Pos() int
	// SkipMod function must advance the read buffer up to the
	// modulo of integer parameter.
	SkipMod(int) error
	// FillMod function must fill the write buffer up to the
	// modulo of integer parameter.
	FillMod(int) error
}

// buffer implements the AlignBuffer.
type buffer struct {
	// the chunked buffer.
	chk ChunkedBuffer
	// the current read/write position.
	pos int
	// the temporary buffer for padding.
	tmp [16]byte
}

// NewAlignBuffer returns the new instance of aligned buffer.
func NewAlignBuffer(chk ChunkedBuffer) AlignBuffer {
	if chk == nil {
		chk = NewChunk(nil, DefaultDataRepresentation) /* expected direction is write */
	}
	return &buffer{chk: chk}
}

// Pos function returns the current position in the buffer.
func (b *buffer) Pos() int {
	return b.pos
}

// SkipMod function skips the reader position to align to
// the modulo `mod`.
func (b *buffer) SkipMod(mod int) error {
	to := ((b.pos + mod - 1) &^ (mod - 1)) - b.pos
	if to == 0 {
		return nil
	}
	n, err := b.chk.Read(b.tmp[:to])
	if err != nil {
		return err
	}
	if b.pos += n; n != to {
		return io.ErrUnexpectedEOF
	}
	return nil
}

// FillMod function fills the data with zero pad modulo `mod`.
func (b *buffer) FillMod(mod int) error {
	to := ((b.pos + mod - 1) &^ (mod - 1)) - b.pos
	if to == 0 {
		return nil
	}
	n, err := b.chk.Write(b.tmp[:to])
	if err != nil {
		return err
	}
	if b.pos += n; n != to {
		return io.ErrShortWrite
	}
	return nil
}

// Write function implements the io.Writer interface.
func (b *buffer) Write(p []byte) (int, error) {
	n, err := b.chk.Write(p)
	if err != nil {
		return n, err
	}
	if b.pos += n; n != len(p) {
		return n, io.ErrShortWrite
	}
	return n, nil
}

// Read function implements the io.Reader interface.
func (b *buffer) Read(p []byte) (int, error) {
	n, err := b.chk.Read(p)
	if err != nil {
		return n, err
	}
	if b.pos += n; n != len(p) {
		return n, io.ErrUnexpectedEOF
	}
	return n, nil
}

// Order function returns the byte order for the buffer chunk.
func (b *buffer) Order() binary.ByteOrder { return b.chk.Order() }

// Float function return the floating-point format for the buffer chunk.
func (b *buffer) Float() math.FloatFormat { return b.chk.Float() }

// Bytes function returns the bytes written to the current
// chunk.
func (b *buffer) Bytes() []byte { return b.chk.Bytes() }

// Len function returns the total number of bytes written
// by this buffer.
func (b *buffer) Len() int { return b.chk.Len() }

// EOF function returns the indicator of the end of the input.
func (b *buffer) EOF() bool { return b.chk.EOF() }

// ReadRepresentation function reads the data representation from the
// underlying chunk.
func (b *buffer) ReadRepresentation(drep *DataRepresentation) error {
	b.pos += 4
	return b.chk.ReadRepresentation(drep)
}

// WriteRepresentation function writes the data representation to the
// underlying chunk.
func (b *buffer) WriteRepresentation(drep DataRepresentation) error {
	b.pos += 4
	return b.chk.WriteRepresentation(drep)
}

// chunk implements the ChunkedBuffer interface.
type chunk struct {
	b    []byte
	pos  int
	drep DataRepresentation
	read bool
}

// NewChunk function returns the new chunked buffer.
func NewChunk(b []byte, drep DataRepresentation) ChunkedBuffer {
	return &chunk{b: b, pos: 0, drep: drep}
}

// Write function implements the io.Writer interface.
func (c *chunk) Write(p []byte) (int, error) {
	c.read = false
	n := copy(c.b[c.pos:], p)
	if n < len(p) {
		c.b, n = append(c.b, p[n:]...), len(p)
	}
	c.pos += n
	return n, nil
}

// Read function implements the io.Reader interface.
func (c *chunk) Read(p []byte) (int, error) {
	c.read = true
	n := copy(p, c.b[c.pos:])
	c.pos += n
	return n, nil
}

// Order function returns the byte order for the buffer chunk.
func (c *chunk) Order() binary.ByteOrder { return c.drep.ByteOrder() }

// Float function return the floating-point format for the buffer chunk.
func (c *chunk) Float() math.FloatFormat { return c.drep.FloatFormat() }

// Bytes function returns the bytes written/not-yet-read for the
// chunk.
func (c *chunk) Bytes() []byte {
	if c.read {
		return c.b[c.pos:]
	}
	return c.b
}

func (c *chunk) ReadRepresentation(drep *DataRepresentation) error {
	p := make([]byte, 4)
	if n, err := c.Read(p); err != nil || n < 4 {
		return io.ErrUnexpectedEOF
	}
	c.drep = (DataRepresentation)(binary.LittleEndian.Uint32(p))
	*drep = c.drep
	return nil
}

func (c *chunk) WriteRepresentation(drep DataRepresentation) error {
	p := make([]byte, 4)
	binary.LittleEndian.PutUint32(p, uint32(c.drep))
	if n, err := c.Write(p); err != nil || n < 4 {
		return io.ErrShortWrite
	}
	c.drep = drep
	return nil
}

// Len function returns the current chunk remaining data.
func (c *chunk) Len() int { return len(c.Bytes()) }

// EOF function indicates whether the chunk is completed.
func (c *chunk) EOF() bool { return c.Len() == 0 }

// WaitChunk structure represents the buffer that blocks until
// provided via `Wait` bytes will be exhausted either by read or
// by write or by the completion of the operation (Done is called).
type WaitChunk struct {
	// The current buffer/data-representation.
	bAndFormat *bAndFormat
	// The channel to retrieve the buffer.
	in chan *bAndFormat
	// The wait chan to block until the operation
	// is complete (or buffer is exhausted).
	out chan struct{}
	// The sync-once to retrieve the first chunk.
	once sync.Once
}

func NewWaitChunk() *WaitChunk {
	chnk := &WaitChunk{in: make(chan *bAndFormat), out: make(chan struct{})}
	return chnk
}

type bAndFormat struct {
	b      []byte
	format DataRepresentation
	maxLen int
}

func (c *WaitChunk) Wait(b []byte, frmt DataRepresentation, maxLen int) int {
	c.in <- &bAndFormat{b, frmt, maxLen}
	<-c.out
	return len(b) - len(c.bAndFormat.b)
}

// Write function implements the io.Writer interface.
func (c *WaitChunk) Write(p []byte) (int, error) {

	var ok bool

	if c.bAndFormat == nil {
		if c.bAndFormat, ok = <-c.in; !ok {
			return 0, nil
		}
	}

	n := copy(c.bAndFormat.b, p)
	c.bAndFormat.b = c.bAndFormat.b[n:]

	for n < len(p) {
		c.out <- struct{}{}
		if c.bAndFormat, ok = <-c.in; !ok {
			return n, nil
		}
		actual := copy(c.bAndFormat.b, p[n:])
		c.bAndFormat.b, n = c.bAndFormat.b[actual:], n+actual
	}

	c.bAndFormat.maxLen -= n
	return n, nil
}

func (c *WaitChunk) Close() {
	close(c.in)
}

func (c *WaitChunk) Done() {
	close(c.out)
}

// Read function implements the io.Reader interface.
func (c *WaitChunk) Read(p []byte) (int, error) {

	var ok bool

	if c.bAndFormat == nil {
		if c.bAndFormat, ok = <-c.in; !ok {
			return 0, nil
		}
	}

	n := copy(p, c.bAndFormat.b)

	c.bAndFormat.b = c.bAndFormat.b[n:]

	for n < len(p) {
		c.out <- struct{}{}
		if c.bAndFormat, ok = <-c.in; !ok {
			return n, nil
		}
		actual := copy(p[n:], c.bAndFormat.b)
		c.bAndFormat.b, n = c.bAndFormat.b[actual:], n+actual
	}

	c.bAndFormat.maxLen -= n
	return n, nil
}

// Order function returns the byte order for the buffer chunk.
func (c *WaitChunk) Order() binary.ByteOrder {

	var ok bool

	if c.bAndFormat == nil {
		if c.bAndFormat, ok = <-c.in; !ok {
			return binary.LittleEndian
		}
	}

	return c.bAndFormat.format.ByteOrder()
}

// Float function return the floating-point format for the buffer chunk.
func (c *WaitChunk) Float() math.FloatFormat {

	var ok bool

	if c.bAndFormat == nil {
		if c.bAndFormat, ok = <-c.in; !ok {
			return math.IEEE
		}
	}

	return c.bAndFormat.format.FloatFormat()
}

// Bytes function returns the bytes written/not-yet-read for the
// chunk.
func (c *WaitChunk) Bytes() []byte { return nil /* panic("wait_chunk: bytes not supported") */ }

func (c *WaitChunk) ReadRepresentation(drep *DataRepresentation) error {
	panic("wait_chunk: read_representation not supported")
}

func (c *WaitChunk) WriteRepresentation(drep DataRepresentation) error {
	panic("wait_chunk: write_representation not supported")
}

// Len function returns the current WaitChunk remaining data. (used for sanity check).
func (c *WaitChunk) Len() int {

	var ok bool

	if c.bAndFormat == nil {
		if c.bAndFormat, ok = <-c.in; !ok {
			return 0
		}
	}

	if c.bAndFormat.maxLen > 0 {
		return c.bAndFormat.maxLen
	}

	return go_math.MaxInt
}

// EOF function indicates whether the WaitChunk is completed.
func (c *WaitChunk) EOF() bool { panic("wait_chunk: eof not supported") }
