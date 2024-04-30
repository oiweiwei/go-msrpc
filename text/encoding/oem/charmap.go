package oem

import "golang.org/x/text/encoding"

type Charmap struct {
	charmap encoding.Encoding
}

func (e Charmap) Encode(s string) ([]byte, error) {
	return e.charmap.NewEncoder().Bytes([]byte(s))
}

func (e Charmap) Decode(b []byte) (string, error) {
	b, err := e.charmap.NewDecoder().Bytes(b)
	return string(b), err
}
