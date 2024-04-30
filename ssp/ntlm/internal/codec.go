package internal

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/oiweiwei/go-msrpc/text/encoding/oem"
)

type StringCodec = oem.Encoding

// CodecValue implements the marshaling/unmarshaling methods.
type CodecValue interface {
	Marshal(context.Context) ([]byte, error)
	Unmarshal(context.Context, []byte) error
}

type Payload struct {
	StartAt           int
	Length, MaxLength *uint16
	Offset            *uint32
	Value             any
}

func PayloadField(length, maxLength *uint16, offset *uint32, value any) *Payload {
	return &Payload{0, length, maxLength, offset, value}
}

// The NTLM coder/decoder interface.
type Codec interface {
	// WithStringCodec sets the preferred string encoding/decoding
	// method (OEM, UTF16-LE).
	WithStringCodec(context.Context, StringCodec) Codec
	// WritePayload function writes the data to the payload buffer.
	WritePayload(context.Context, *Payload) error
	// WriteData function writes the data to the header buffer.
	WriteData(context.Context, any) error
	// WriteBytes function writes the bytes to the buffer.
	WriteBytes(context.Context, []byte, int) error
	// Bytes function returns the written bytes.
	Bytes(context.Context) ([]byte, error)
	// ReadPayload function reads the payload data from the buffer.
	ReadPayload(context.Context, *Payload) error
	// ReadData function reads the data from the buffer.
	ReadData(context.Context, any) error
	// ReadBytes reads the bytes into the allocated buffer.
	ReadBytes(context.Context, *[]byte, int) error
	// Complete the read by reading deferred payload fields.
	ReadAll(context.Context) error
	// Len returns the current offset (or bytes read).
	Len(context.Context) int
}

// NewCodec function creates the new coder/decoder.
func NewCodec(ctx context.Context, buf []byte) Codec {
	return &codec{
		stringCodec: oem.FromContext(ctx),
		h:           bytes.NewBuffer(buf),
		p:           bytes.NewBuffer(nil),
		buf:         buf,
	}
}

// The NTLM codec implementation.
type codec struct {
	// payload fields.
	pfs []*Payload
	// payload and header buffers.
	p, h *bytes.Buffer
	// full decode buffer.
	buf []byte
	// error.
	err error
	// string encoder.
	stringCodec StringCodec
}

// Err returns the error.
func (e *codec) Err(ctx context.Context) error {
	return e.err
}

// withErr sets the error.
func (e *codec) withErr(err error) error {
	if e.err == nil {
		e.err = err
	}
	return e.err
}

// WithStringCodec function sets the strings encoder/decoder.
func (e *codec) WithStringCodec(ctx context.Context, codec StringCodec) Codec {
	e.stringCodec = codec
	return e
}

// Len function returns the offset of the headers buffer.
func (e *codec) Len(ctx context.Context) int {
	return e.h.Len()
}

// Bytes function flushes the buffer into array of bytes.
func (e *codec) Bytes(ctx context.Context) ([]byte, error) {

	if e.err != nil {
		return nil, e.err
	}

	// initial offset for the payload.
	offset := uint32(e.h.Len())

	// write payload buffer.
	if _, err := e.p.WriteTo(e.h); err != nil {
		return nil, e.withErr(err)
	}

	b := e.h.Bytes()

	// iterate over payloads and write offset/length values.
	for _, p := range e.pfs {
		if *p.Length != 0 {
			// advance offset.
			*p.Offset, offset = offset, offset+uint32(*p.Length)
		}
		// encode len/max_len/offset on given position.
		binary.LittleEndian.PutUint16(b[p.StartAt:], *p.Length)
		binary.LittleEndian.PutUint16(b[p.StartAt+2:], *p.MaxLength)
		binary.LittleEndian.PutUint32(b[p.StartAt+4:], *p.Offset)
	}

	return b, nil

}

// WriteData function writes the data into the headers buffer.
func (e *codec) WriteData(ctx context.Context, data any) error {

	if e.err != nil {
		return e.err
	}

	if data == nil {
		return nil
	}

	// try encoder value first.
	if codecValue, ok := data.(CodecValue); ok {
		b, err := codecValue.Marshal(ctx)
		if err != nil {
			return e.withErr(err)
		}
		return e.WriteBytes(ctx, b, len(b))
	}

	// encode with default encoding/binary package.
	return e.withErr(binary.Write(e.h, binary.LittleEndian, data))
}

// ReadData function reads the data from the buffer.
func (e *codec) ReadData(ctx context.Context, data any) error {

	if e.err != nil {
		return e.err
	}

	if data == nil {
		return nil
	}

	// try decoder value first.
	if codecValue, ok := data.(CodecValue); ok {
		return e.withErr(codecValue.Unmarshal(ctx, e.h.Bytes()))
	}

	// decode with default encoding/binary package.
	return e.withErr(binary.Read(e.h, binary.LittleEndian, data))
}

// WriteBytes function writes the fixed bytes data to the headers buffer.
func (e *codec) WriteBytes(ctx context.Context, b []byte, sz int) error {

	if e.err != nil {
		return e.err
	}

	// trim input.
	if sz >= 0 && len(b) > sz {
		b = b[:sz]
	}

	// write data.
	n, err := e.h.Write(b)
	if err != nil {
		return e.withErr(err)
	}

	// pad to necessary size.
	for n := n; n < sz; n++ {
		e.withErr(e.h.WriteByte(0))
	}

	return e.err
}

// WriteBytes function writes the fixed bytes data to the headers buffer.
func (e *codec) ReadBytes(ctx context.Context, b *[]byte, sz int) error {

	if e.err != nil {
		return e.err
	}

	if sz <= 0 {
		return nil
	}

	if b == nil {
		for i := 0; i < sz; i++ {
			_, err := e.h.ReadByte()
			if err != nil {
				return e.withErr(err)
			}
		}
		return nil
	}

	*b = make([]byte, sz)

	n, err := e.h.Read(*b)
	if err != nil {
		return e.withErr(err)
	}
	// check length.
	if n != len(*b) {
		return e.withErr(io.ErrUnexpectedEOF)
	}
	return nil
}

// WritePayload function writes the payload to the payload buffer and reserves
// space for writing an length/max_length/offset information in header buffer.
func (e *codec) WritePayload(ctx context.Context, p *Payload) error {

	if e.err != nil {
		return e.err
	}

	var (
		b   []byte
		err error
	)

	switch v := p.Value.(type) {
	case string:
		// encode oem/utf16le string.
		if b, err = e.stringCodec.Encode(v); err != nil {
			return e.withErr(err)
		}
	case []byte:
		// encode bytes.
		b = v
	case CodecValue:
		// encodervalue.
		if b, err = v.Marshal(ctx); err != nil {
			return e.withErr(err)
		}
	default:
		return e.withErr(fmt.Errorf("unsupported type %T for WritePayload", v))
	}

	// write bytes to the payload buffer.
	if len(b) > 0 {
		if _, err := e.p.Write(b); err != nil {
			return e.withErr(err)
		}
	}

	// headers buffer start position.
	p.StartAt = e.h.Len()

	// set proper length.
	*p.Length = uint16(len(b))
	*p.MaxLength = uint16(len(b))

	e.pfs = append(e.pfs, p)

	// write 8-byte pad for field payload.
	return e.WriteBytes(ctx, nil, 8)
}

func (e *codec) ReadPayload(ctx context.Context, p *Payload) error {

	if e.err != nil {
		return e.err
	}

	e.ReadData(ctx, p.Length)
	e.ReadData(ctx, p.MaxLength)
	e.ReadData(ctx, p.Offset)

	if *p.Length > 0 {
		e.pfs = append(e.pfs, p)
	}

	return nil
}

func (e *codec) ReadAll(ctx context.Context) error {

	for _, p := range e.pfs {

		length, offset := int(*p.Length), int(*p.Offset)

		// check that offset doesn't exceeds the buffer length
		// and doesn't overlap with header length.
		if offset+length > len(e.buf) || (offset != 0 && length != 0 && offset < len(e.buf)-e.h.Len()) {
			return e.withErr(fmt.Errorf("unexpected offset %d+%d for buffer length %d (header length %d)",
				offset, length, len(e.buf), e.h.Len()))
		}

		if length == 0 {
			continue
		}

		// get chunk.
		b := e.buf[offset : offset+length]

		switch value := p.Value.(type) {
		case *string:
			// decode oem/utf16le string.
			dec, err := e.stringCodec.Decode(b)
			if err != nil {
				return e.withErr(err)
			}
			*value = dec
		case *[]byte:
			// decode bytes.
			*value = make([]byte, len(b))
			copy(*value, b)
		case CodecValue:
			return e.withErr(value.Unmarshal(ctx, b))
		default:
			return e.withErr(fmt.Errorf("unsupported type %T for Complete", value))
		}

	}

	return nil
}
