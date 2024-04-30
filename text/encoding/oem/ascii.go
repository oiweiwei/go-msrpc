package oem

import "unicode"

type ascii struct{}

const (
	asciiReplacementChar = 0x1A
)

func (ascii) Encode(s string) ([]byte, error) {
	runes := []rune(s)
	b := make([]byte, len(runes))
	for i := 0; i < len(b); i++ {
		if runes[i] > unicode.MaxASCII {
			b[i] = asciiReplacementChar
		} else {
			b[i] = byte(runes[i])
		}
	}
	return b, nil
}

func (ascii) Decode(b []byte) (string, error) {
	runes := make([]rune, len(b))
	for i := 0; i < len(b); i++ {
		if b[i] == asciiReplacementChar {
			runes[i] = unicode.ReplacementChar
		} else {
			runes[i] = rune(b[i])
		}
	}
	return string(runes), nil
}
