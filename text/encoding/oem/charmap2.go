package oem

import "unicode"

type Charmap2 struct {
	toUTF8   [256]rune
	fromUTF8 map[rune]byte
}

func (c2 Charmap2) Encode(s string) ([]byte, error) {
	out := make([]byte, len([]rune(s)))
	for i, r := range []rune(s) {
		c, ok := c2.fromUTF8[r]
		if !ok {
			c = 0x1A
		}
		out[i] = c
	}
	return out, nil
}

func (c2 Charmap2) Decode(b []byte) (string, error) {
	out := make([]rune, len(b))
	for i := range b {
		out[i] = c2.toUTF8[b[i]]
	}
	return string(out), nil
}

func NewCharmap2(toUTF8 [256]rune) *Charmap2 {
	chrmap := Charmap2{toUTF8: toUTF8, fromUTF8: make(map[rune]byte)}
	for i := range toUTF8 {
		if toUTF8[i] == unicode.ReplacementChar {
			continue
		}
		chrmap.fromUTF8[toUTF8[i]] = byte(i)
	}
	return &chrmap
}
