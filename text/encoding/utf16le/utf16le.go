package utf16le

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

var UTF16LE = unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)

type UTF16LECodec struct {
	encoding.Encoding
}

func NewEncoding() *UTF16LECodec {
	return &UTF16LECodec{UTF16LE}
}

func (c *UTF16LECodec) Encode(s string) ([]byte, error) {
	return c.Encoding.NewEncoder().Bytes([]byte(s))
}

func (c *UTF16LECodec) Decode(b []byte) (string, error) {
	b, err := c.Encoding.NewDecoder().Bytes(b)
	return string(b), err
}

func Encode(s string) ([]byte, error) {
	return NewEncoding().Encode(s)
}

func Decode(b []byte) (string, error) {
	return NewEncoding().Decode(b)
}
