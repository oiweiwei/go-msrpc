package wmio

// codec.go module contains the definitions and implementation
// for the WMIO encoder/decoder.

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
	"strings"
	"unicode"
	"unicode/utf16"
)

// The NULL reference.
const NULLRef = uint32(0xFFFFFFFF)

// NewCodec function returns the new encoder/decoder for the WMIO
// objects.
func NewCodec(b []byte) *Codec {
	return &Codec{buf: bytes.NewBuffer(b), refs: &Heaps{}}
}

// The WMIO encoder/decoder.
type Codec struct {
	// The debugging context.
	ctx []string
	// The first captured error.
	err error
	// The input/output buffer.
	buf *bytes.Buffer
	// The set of references.
	refs *Heaps
}

// Err.
func (c *Codec) Err(err ...error) error {
	if c.err == nil && len(err) > 0 {
		c.err = err[0]
	}
	return c.err
}

// Errf.
func (c *Codec) Errf(frmt string, args ...interface{}) error {
	if c.err == nil {
		c.err = fmt.Errorf(frmt, args...)
	}
	return c.err
}

// Len.
func (c *Codec) Len() int {
	return c.buf.Len()
}

// clone.
func (c *Codec) clone() *Codec {
	return &Codec{err: c.err, ctx: c.ctx, buf: c.buf, refs: c.refs}
}

// Bytes.
func (c *Codec) Bytes() []byte {
	return c.buf.Bytes()
}

// withBytes function sets the buffer bytes to the value `b`.
func (c *Codec) withBytes(b []byte) *Codec { c.buf = bytes.NewBuffer(b); return c }

// withBytesBuffer function sets the buffer to the value `buf`.
func (c *Codec) withBytesBuffer(buf *bytes.Buffer) *Codec { c.buf = buf; return c }

// Begin function is used to start the debugging context.
func (c *Codec) Begin(s any) {
	if c.Err() == nil {
		c.ctx = append(c.ctx, fmt.Sprintf("%v", s))
	}
}

// Done function is used to pop the debugging context.
func (c *Codec) Done() error {
	if c.Err() != nil {
		return c.Err()
	}
	c.ctx = c.ctx[:len(c.ctx)-1]
	return nil
}

// ReadRef function reads the reference for the data, and if reference
// is not NULL, appends the reference to the list of heap references.
func (c *Codec) ReadRef(data any, debug ...string) error {

	var ref uint32

	if c.ReadData(&ref); ref == NULLRef {
		// it's nil, nothing to do.
		return nil
	}

	if ref&uint32(1<<31) != 0 {
		// it's a dictionary value.
		s, ok := data.(*string)
		if !ok {
			return c.Errf("%s: invalid dictionary reference 0x%08x",
				strings.Join(append(c.ctx, debug...), "."), ref)
		}
		*s = LookupDictionary(ref)
		return nil
	}

	// add references to the heap references to be read later
	// (once the heap will be available).
	c.refs.Append(&HeapRef{ref, data, append(c.ctx, debug...)})
	return nil
}

// WriteRef function writes the reference on the current heap.
// NOTE: for situations when written value is a reference itself,
// WriteDataOnHeap must be used, to allocate enough space for the
// reference itself, and the value it is referred by.
func (c *Codec) WriteRef(data any, debug ...string) error {

	if data == nil || reflect.ValueOf(data).IsZero() {
		// it's null, write NULL-ref.
		return c.WriteData(NULLRef)
	}

	if s, ok := data.(string); ok {
		// it's a string within dictionary.
		if ref := ReverseLookupDictionary(s); ref != 0 {
			return c.WriteData(ref)
		}
	}

	// use heap as a codec.
	heap := c.clone().withBytesBuffer(c.refs.HeapBuffer())

	// write current heap length as a reference.
	c.WriteData((uint32)(heap.Len()))

	// write actual data.
	return c.Err(heap.WriteData(data))
}

// DecodeWithBytes function decodes the `data` within given
// byte chunk `b`.
func (c *Codec) DecodeWithBytes(b []byte, data any) error {
	return c.Err(c.clone().withBytes(b).ReadData(data))
}

// DecodeWithSize32 function reads the sized chunk and decodes
// the data within it.
func (c *Codec) DecodeWithSize32(sz uint32, data any) error {

	if sz == 0 {
		return nil
	}

	if sz > uint32(c.buf.Len()) {
		return c.Errf("decode_with_size32: expected size is greater than actual: %d > %d", sz, c.buf.Len())
	}

	b := make([]byte, sz)
	c.ReadData(b)

	return c.DecodeWithBytes(b, data)
}

// EncodeBytesWithSize32 function writes the data and saves the captured
// length into the value `sz`.
func (c *Codec) EncodeBytesWithSize32(sz *uint32, data any) ([]byte, error) {

	clone := c.clone().withBytes(nil)

	if c.Err(clone.WriteData(data)) != nil {
		return nil, c.Err()
	}

	b := clone.buf.Bytes()

	if sz != nil {
		*sz = uint32(len(b))
	}

	return b, nil
}

// DecodeWithLength32 function reads the uint32 data length prefix
// (length is calculated as len + size_of(len)).
func (c *Codec) DecodeWithLength32(data any) error {

	var len uint32
	c.ReadData(&len)

	return c.DecodeWithSize32(len-4 /* length is included */, data)
}

// EncodeWithLength32 function encodes the data and length prefix.
func (c *Codec) EncodeWithLength32(data any) error {

	clone := c.clone().withBytes(nil)

	if c.Err(clone.WriteData(data)) != nil {
		return c.Err()
	}

	b := clone.buf.Bytes()
	c.WriteData((uint32(len(b)) + 4))

	return c.WriteData(b)
}

// ReadHeap function reads the heap data. (size of heap is calculated
// by inverting the left-most bit).
func (c *Codec) ReadHeap() error {
	var hSz uint32
	c.ReadData(&hSz)
	hSz &= ^(uint32(1 << 31)) // invart left-most bit.
	h := make([]byte, hSz)
	c.ReadData(h)
	c.refs.SetHeap(h)
	return nil
}

// WriteHeap function writes the heap binary data to the buffer.
func (c *Codec) WriteHeap() error {
	h := c.refs.Heap()
	hSz := uint32(len(h)) | uint32(1<<31)
	c.WriteData(hSz)
	return c.WriteData(h)
}

// DecodeHeap function decodes the references on the current heap.
// If decoded data has more references, they will be decoded as well.
func (c *Codec) DecodeHeap() error {
	for len(c.refs.Head()) > 0 {
		for _, ref := range c.refs.Truncate() {
			c.DecodeWithBytes(c.refs.Heap()[ref.Offset:], ref.Value)
		}
	}
	return nil
}

// Read.
func (c *Codec) Read(p []byte) error {
	_, err := c.buf.Read(p)
	return c.Err(err)
}

// ReadData function reads the generic data from the buffer.
func (c *Codec) ReadData(data any) error {

	if c.Err() != nil {
		return c.Err()
	}

	if d, ok := data.(interface{ Decode(*Codec) error }); ok {
		// it's a decoder interface.
		return c.Err(d.Decode(c))
	}

	switch data := data.(type) {
	case func(*Codec) error:
		// it's a decoder function.
		return c.Err(data(c))
	case []byte:
		// binary data.
		return c.Read(data)
	case *[]*Qualifier:
		// qualifiers are special case.
		return c.DecodeWithLength32(func(c *Codec) error {
			c.Begin("qualifiers")
			for c.Len() > 0 {
				q := Qualifier{}
				c.ReadData(&q)
				*data = append(*data, &q)
			}
			return c.Done()
		})
	case *string:
		// strings.
		var flags uint8
		c.ReadData(&flags)
		switch flags {
		case 0x01: // unicode-string.
			ret, buf := []uint16{}, uint16(0)
			for c.Err(binary.Read(c.buf, binary.LittleEndian, &buf)) == nil && buf != 0 {
				ret = append(ret, buf)
			}
			*data = string(utf16.Decode(ret))
		case 0x00: // ascii-string.
			ret, buf := []byte{}, [1]byte{}
			for c.Read(buf[:]) == nil && buf[0] != 0 {
				ret = append(ret, buf[:]...)
			}
			*data = string(ret)
		default:
			return c.Errf("invalid string flag 0x%02x", flags)
		}
	default:
		return c.Err(binary.Read(c.buf, binary.LittleEndian, data))
	}
	return nil
}

// WriteDataOnHeap function is used to write reference and refered value
// to the heap. First, we allocate space for the reference (or data that
// includes a reference), and then we write the actual referred data right
// after the reference.
func (c *Codec) WriteDataOnHeap(sz int, data any) error {

	// get current heap.
	h := c.refs.Heap()

	// allocate new heap and store the data + allocate size.
	c.refs.Push()
	c.refs.SetHeap(append(h, make([]byte, sz)...))

	// capture current length.
	before := c.Len()
	c.WriteData(data)

	if actualSz := c.Len() - before; actualSz < sz {
		// write missing bytes (expected size was more than actual).
		c.WriteData(make([]byte, sz-actualSz))
	} else if actualSz > sz {
		// throw an error, expected size was insufficient.
		return c.Errf("write_data_on_heap: actual size is greater than expected: %d > %d", actualSz, sz)
	}

	// get the heap data.
	h = c.refs.Heap()[len(h)+sz:]
	// remove the virtual-heap.
	c.refs.Pop()
	// append the virtual-heap data to the real heap.
	c.refs.HeapBuffer().Write(h)

	return nil
}

// Write.
func (c *Codec) Write(p []byte) error {
	_, err := c.buf.Write(p)
	return c.Err(err)
}

// WriteData function writes the generic data into the buffer.
func (c *Codec) WriteData(data any) error {

	if d, ok := data.(interface{ Encode(*Codec) error }); ok {
		// it's encoder interface.
		return c.Err(d.Encode(c))
	}

	switch data := data.(type) {
	case func(*Codec) error:
		// it's encoder function.
		return c.Err(data(c))
	case []byte:
		// binary data.
		return c.Write(data)
	case []*Qualifier:
		return c.EncodeWithLength32(func(c *Codec) error {
			c.Begin("qualifiers")
			for i := range data {
				c.WriteData(data[i])
			}
			return c.Done()
		})
	case string:

		var flags uint8

		if !IsASCII(data) {
			flags = 0x01
		}

		c.WriteData(flags)

		switch flags {
		case 0x01: // unicode-string.
			return c.Err(binary.Write(c.buf, binary.LittleEndian, append(utf16.Encode(([]rune)(data)), uint16(0))))
		case 0x00: // ascii-string.
			return c.Write(append([]byte(data), 0))
		}
	default:
		return c.Err(binary.Write(c.buf, binary.LittleEndian, data))
	}

	return nil
}

// IsASCII.
func IsASCII(data string) bool {
	for i := 0; i < len(data); i++ {
		if data[i] > unicode.MaxASCII {
			return false
		}
	}

	return true
}
