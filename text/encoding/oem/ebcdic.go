package oem

import (
	"github.com/indece-official/go-ebcdic"
)

type EBCDIC struct {
	CodePage int
}

func (e EBCDIC) Encode(s string) ([]byte, error) {
	return ebcdic.Encode(s, e.CodePage)
}

func (e EBCDIC) Decode(b []byte) (string, error) {
	return ebcdic.Decode(b, e.CodePage)
}
