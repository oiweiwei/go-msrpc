package ndr

import (
	"context"
	"strings"
	"unicode/utf16"
)

func MultiSzLen(s []string) uint64 {
	l := uint64(len(s)) + 1
	for i := range s {
		l += UTF16Len(s[i])
	}
	return l
}

// UTF16NLen function ...
func UTF16NLen(s string) uint64 {
	l := len(s)
	for _, r := range s {
		if r == 0x0000 {
			l--
			// null-terminated string.
			break
		}
		if r >= 0x10000 && r <= '\U0010FFFF' {
			l++
		}
	}
	return uint64(l) + 1
}

func UTF16Len(s string) uint64 {
	l := len(s)
	for _, r := range s {
		if r >= 0x10000 && r <= '\U0010FFFF' {
			l++
		}
	}
	return uint64(l)
}

func CharLen(s string) uint64 {
	return uint64(len([]byte(s)))
}

func CharNLen(s string) uint64 {

	l := 0
	for _, b := range []byte(s) {
		if b == 0x0 {
			break
		}
		l++
	}
	return uint64(l) + 1
}

func WriteCharNString(ctx context.Context, w Writer, s string) error {

	l := CharNLen(s)

	if err := w.WriteSize(l); err != nil {
		return err
	}
	if err := w.WriteSize(uint64(0)); err != nil {
		return err
	}
	if err := w.WriteSize(l); err != nil {
		return err
	}

	for _, chr := range []byte(s) {
		if err := w.WriteData(chr); err != nil {
			return err
		}
	}

	return w.WriteData(uint8(0))
}

// ReadUTF16NString ...
func ReadCharNString(ctx context.Context, r Reader, s *string) error {

	sz := uint64(0)

	// max_count.
	if err := r.ReadSize(&sz); err != nil {
		return err
	}
	// offset.
	if err := r.ReadSize(&sz); err != nil {
		return err
	}
	// actual_count.
	if err := r.ReadSize(&sz); err != nil {
		return err
	}

	var buf = make([]byte, sz)

	for i := range buf {
		if err := r.ReadData(&buf[i]); err != nil {
			return err
		}
	}

	*s = strings.TrimRight(string(buf), ZeroString)

	return nil
}

func WriteCharString(ctx context.Context, w Writer, s string) error {

	l := CharLen(s)

	if err := w.WriteSize(l); err != nil {
		return err
	}
	if err := w.WriteSize(uint64(0)); err != nil {
		return err
	}
	if err := w.WriteSize(l); err != nil {
		return err
	}

	for _, chr := range []byte(s) {
		if err := w.WriteData(chr); err != nil {
			return err
		}
	}

	return nil
}

// ReadCharString ...
func ReadCharString(ctx context.Context, r Reader, s *string) error {

	sz := uint64(0)

	// max_count.
	if err := r.ReadSize(&sz); err != nil {
		return err
	}
	// offset.
	if err := r.ReadSize(&sz); err != nil {
		return err
	}
	// actual_count.
	if err := r.ReadSize(&sz); err != nil {
		return err
	}

	var buf = make([]byte, sz)

	for i := range buf {
		if err := r.ReadData(&buf[i]); err != nil {
			return err
		}
	}

	*s = string(buf)

	return nil
}

func WriteUTF16String(ctx context.Context, w Writer, s string) error {

	l := UTF16Len(s)

	if err := w.WriteSize(l); err != nil {
		return err
	}
	if err := w.WriteSize(uint64(0)); err != nil {
		return err
	}
	if err := w.WriteSize(l); err != nil {
		return err
	}

	for _, chr := range utf16.Encode([]rune(s)) {
		if err := w.WriteData(chr); err != nil {
			return err
		}
	}

	return nil
}

// ReadUTF16NString ...
func ReadUTF16String(ctx context.Context, r Reader, s *string) error {

	sz := uint64(0)

	// max_count.
	if err := r.ReadSize(&sz); err != nil {
		return err
	}
	// offset.
	if err := r.ReadSize(&sz); err != nil {
		return err
	}
	// actual_count.
	if err := r.ReadSize(&sz); err != nil {
		return err
	}

	var buf = make([]uint16, sz)

	for i := range buf {
		if err := r.ReadData(&buf[i]); err != nil {
			return err
		}
	}

	*s = string(utf16.Decode(buf))

	return nil
}

// WriteUTF16NString function ...
func WriteUTF16NString(ctx context.Context, w Writer, s string) error {

	l := UTF16NLen(s)

	if err := w.WriteSize(l); err != nil {
		return err
	}
	if err := w.WriteSize(uint64(0)); err != nil {
		return err
	}
	if err := w.WriteSize(l); err != nil {
		return err
	}

	s = strings.TrimRight(s, ZeroString)

	for _, chr := range utf16.Encode([]rune(s)) {
		if err := w.WriteData(chr); err != nil {
			return err
		}
	}

	return w.WriteData(uint16(0))
}

// ReadUTF16NString ...
func ReadUTF16NString(ctx context.Context, r Reader, s *string) error {

	sz := uint64(0)

	// max_count.
	if err := r.ReadSize(&sz); err != nil {
		return err
	}
	// offset.
	if err := r.ReadSize(&sz); err != nil {
		return err
	}
	// actual_count.
	if err := r.ReadSize(&sz); err != nil {
		return err
	}

	var buf = make([]uint16, sz)

	for i := range buf {
		if err := r.ReadData(&buf[i]); err != nil {
			return err
		}
	}

	*s = strings.TrimRight(string(utf16.Decode(buf)), ZeroString)

	return nil
}

func ReadUTF16NStringNoSize(ctx context.Context, r Reader, s *string) error {

	var buf = make([]uint16, 0)
	for {
		var chr uint16
		if err := r.ReadData(&chr); err != nil {
			return err
		}
		if chr == 0x0000 {
			break
		}
		buf = append(buf, chr)
	}

	*s = strings.TrimRight(string(utf16.Decode(buf)), ZeroString)

	return nil
}
