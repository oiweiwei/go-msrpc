package binxml

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
	"strings"

	"github.com/oiweiwei/go-msrpc/text/encoding/utf16le"
)

var (
	Debug = false
)

type Reader struct {
	binary.ByteOrder
	r   *bytes.Reader
	sz  int
	err error
	ctx []string
}

func NewReader(b []byte) *Reader {
	return &Reader{binary.LittleEndian, bytes.NewReader(b), len(b), nil, nil}
}

func (r *Reader) cloneWithBytes(b []byte) *Reader {
	return &Reader{r.ByteOrder, bytes.NewReader(b), len(b), r.err, r.ctx}
}

func (r *Reader) Begin(ctx string) { r.ctx = append(r.ctx, ctx) }

func (r *Reader) Done() error {
	r.ctx = r.ctx[:len(r.ctx)-1]
	return r.err
}

func (r *Reader) SetSize(sz *uint32) {
	*sz = uint32(r.sz)
}

func (r *Reader) ReadUTF16NStringWithSize(s any) error {
	var sz uint16
	r.Read(&sz)
	return r.readUTF16String(s, int(sz), true)
}

func (r *Reader) ReadUTF16StringWithSize(s any) error {
	var sz uint16
	r.Read(&sz)
	return r.readUTF16String(s, int(sz), false)
}

func (r *Reader) ReadUTF16String(s any, sz int) error {
	return r.readUTF16String(s, sz, false)
}

func (r *Reader) readUTF16String(s any, sz int, null bool) error {
	if sz == 0 {
		return nil
	}
	b := make([]byte, sz*2)
	if r.Read(b); null {
		r.Read(make([]byte, 2))
	}
	ds, err := utf16le.Decode(b)
	if err != nil {
		return r.WithErr(err)
	}

	if s, ok := s.(*string); ok {
		*s = ds
	}

	return nil
}

type excludeSize struct{}

var (
	ExcludeSize = excludeSize{}
)

func (r *Reader) ReadWithLen(f interface{}, excludeSize ...excludeSize) error {
	l := uint32(0)
	r.Read(&l)
	if len(excludeSize) > 0 && l >= 4 {
		l -= 4
	}
	b := make([]byte, l)
	r.Read(b)
	return r.ReadWithBytes(b, f)
}

func (r *Reader) ReadWithBytes(b []byte, f interface{}) error {

	rr := r.cloneWithBytes(b)

	switch f := f.(type) {
	case func(*Reader) error:
		if err := f(rr); err != nil {
			return r.WithErr(err)
		}
	default:
		rr.Read(f)
	}

	if rr.err != nil {
		return r.WithErr(rr.err)
	}

	if rr.r.Len() != 0 {
		return r.WithErrf("unread portion of bytes of size %d", rr.r.Len())
	}

	return nil
}

func (r *Reader) WithErr(err ...error) error {
	if r.err != nil {
		return r.err
	}
	if len(err) > 0 && err[0] != nil {
		r.err = fmt.Errorf("%s: %v", strings.Join(r.ctx, "."), err[0])
	}
	return r.err
}

func (r *Reader) WithErrf(frmt string, args ...interface{}) error {
	return r.WithErr(fmt.Errorf(frmt, args...))
}

func reflectIsDecoder(v interface{}) (interface{ Decode(r *Reader) error }, bool) {
	ptr := reflect.ValueOf(v)
	if ptr.Type().Kind() != reflect.Pointer {
		return nil, false
	}

	ptr = ptr.Elem()

	if ptr.Type().Kind() != reflect.Pointer {
		return nil, false
	}

	if ptr.IsNil() {
		ptr.Set(reflect.New(ptr.Type().Elem()))
	}

	dr, ok := ptr.Interface().(interface{ Decode(r *Reader) error })
	return dr, ok

}

func (r *Reader) Read(data any) (int, error) {

	if r.err != nil {
		return 0, r.err
	}

	if Debug {
		fmt.Println(strings.Join(r.ctx, "."), fmt.Sprintf("read %T", data))
	}

	b, ok := data.([]byte)
	if ok {
		n, err := r.r.Read(b)
		if err != nil {
			return n, r.WithErr(err)
		}
		if n != len(b) {
			return n, r.WithErrf("read %d out of %d bytes", n, len(b))
		}
		return n, nil
	}

	l := r.r.Len()

	if dr, ok := reflectIsDecoder(data); ok {
		err := dr.Decode(r)
		return l - r.r.Len(), r.WithErr(err)
	}

	if err := binary.Read(r.r, binary.LittleEndian, data); err != nil {
		return l - r.r.Len(), r.WithErr(err)
	}

	if Debug {
		fmt.Println("data out:", data)
	}

	return l - r.r.Len(), nil
}

func (r *Reader) ReadTagAndMore(expect ...uint8) (uint8, bool) {
	return r.readTag(expect...)
}

func (r *Reader) ReadTag(expect ...uint8) uint8 {
	tag, _ := r.readTag(expect...)
	return tag
}

func (r *Reader) readTag(expect ...uint8) (uint8, bool) {

	t := [1]byte{}
	if _, err := r.Read(t[:]); err != nil {
		r.WithErr(err)
		return 0xFF, false
	}

	tag, more := uint8(t[0] & ^uint8(0x40)), t[0]&0x40 != 0

	if len(expect) > 0 {
		found := false
		for _, exp := range expect {
			if tag == exp {
				found = true
				break
			}
		}
		if !found {
			r.WithErrf("expected one of %v, got %d", expect, t[0])
			return 0xFF, false
		}
	}

	return tag, more
}
