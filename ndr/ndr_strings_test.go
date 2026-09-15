package ndr

import (
	"context"
	"encoding/binary"
	"runtime"
	"testing"
)

type vChar struct{ s string }

func (c *vChar) UnmarshalNDR(ctx context.Context, r Reader) error {
	return ReadCharNString(ctx, r, &c.s)
}

type vU16 struct{ s string }

func (c *vU16) UnmarshalNDR(ctx context.Context, r Reader) error {
	return ReadUTF16String(ctx, r, &c.s)
}

type utf16NWithFollowing struct {
	s         string
	following uint32
}

func (v *utf16NWithFollowing) MarshalNDR(ctx context.Context, w Writer) error {
	if err := WriteUTF16NString(ctx, w, v.s); err != nil {
		return err
	}
	return w.WriteData(v.following)
}

func (v *utf16NWithFollowing) UnmarshalNDR(ctx context.Context, r Reader) error {
	if err := ReadUTF16NString(ctx, r, &v.s); err != nil {
		return err
	}
	return r.ReadData(&v.following)
}

type utf16WithFollowing struct {
	s         string
	following uint32
}

var utf16LengthSink uint64

func (v *utf16WithFollowing) MarshalNDR(ctx context.Context, w Writer) error {
	if err := WriteUTF16String(ctx, w, v.s); err != nil {
		return err
	}
	return w.WriteData(v.following)
}

func (v *utf16WithFollowing) UnmarshalNDR(ctx context.Context, r Reader) error {
	if err := ReadUTF16String(ctx, r, &v.s); err != nil {
		return err
	}
	return r.ReadData(&v.following)
}

func TestUTF16LengthsUseCodeUnits(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		want  uint64
		wantN uint64
	}{
		{name: "ASCII", s: "abc", want: 3, wantN: 4},
		{name: "BMP", s: "\u03a9", want: 1, wantN: 2},
		{name: "supplementary", s: "\U0001f642", want: 2, wantN: 3},
		{name: "empty", s: "", want: 0, wantN: 1},
		{name: "leading NUL", s: "\x00abc", want: 4, wantN: 1},
		{name: "multiple NULs", s: "a\x00\x00b", want: 4, wantN: 2},
		{name: "interior NUL", s: "a\x00b", want: 3, wantN: 2},
		{name: "invalid UTF-8", s: string([]byte{0xff}), want: 1, wantN: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UTF16Len(tt.s); got != tt.want {
				t.Fatalf("UTF16Len(%q) = %d, want %d", tt.s, got, tt.want)
			}
			if got := UTF16NLen(tt.s); got != tt.wantN {
				t.Fatalf("UTF16NLen(%q) = %d, want %d", tt.s, got, tt.wantN)
			}
		})
	}
}

func TestUTF16LengthHelpersDoNotAllocate(t *testing.T) {
	tests := []struct {
		name string
		fn   func() uint64
	}{
		{name: "UTF16Len", fn: func() uint64 { return UTF16Len("\U0001f642") }},
		{name: "UTF16NLen stops at leading NUL", fn: func() uint64 { return UTF16NLen("\x00abc") }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testing.AllocsPerRun(100, func() { utf16LengthSink = tt.fn() }); got != 0 {
				t.Fatalf("allocations per run = %f, want 0", got)
			}
		})
	}
}

func TestWriteUTF16NStringMatchesHeaderAndFollowingField(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		count uint32
		units []uint16
	}{
		{name: "ASCII", s: "abc", count: 4, units: []uint16{0x0061, 0x0062, 0x0063, 0x0000}},
		{name: "BMP", s: "\u03a9", count: 2, units: []uint16{0x03a9, 0x0000}},
		{name: "supplementary", s: "\U0001f642", count: 3, units: []uint16{0xd83d, 0xde42, 0x0000}},
		{name: "empty", s: "", count: 1, units: []uint16{0x0000}},
		{name: "leading NUL", s: "\x00abc", count: 1, units: []uint16{0x0000}},
		{name: "multiple NULs", s: "a\x00\x00b", count: 2, units: []uint16{0x0061, 0x0000}},
		{name: "interior NUL", s: "a\x00b", count: 2, units: []uint16{0x0061, 0x0000}},
		{name: "trailing NUL", s: "a\x00", count: 2, units: []uint16{0x0061, 0x0000}},
		{name: "invalid UTF-8", s: string([]byte{0xff}), count: 2, units: []uint16{0xfffd, 0x0000}},
	}

	const marker = uint32(0xa1b2c3d4)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wire, err := Marshal(&utf16NWithFollowing{s: tt.s, following: marker})
			if err != nil {
				t.Fatalf("Marshal() error = %v", err)
			}

			for _, offset := range []int{0, 8} {
				if got := binary.LittleEndian.Uint32(wire[offset:]); got != tt.count {
					t.Fatalf("count at wire offset %d = %d, want %d", offset, got, tt.count)
				}
			}
			if got := binary.LittleEndian.Uint32(wire[4:]); got != 0 {
				t.Fatalf("offset = %d, want 0", got)
			}
			for i, want := range tt.units {
				if got := binary.LittleEndian.Uint16(wire[12+i*2:]); got != want {
					t.Fatalf("unit %d = 0x%04x, want 0x%04x", i, got, want)
				}
			}

			followingOffset := (12 + len(tt.units)*2 + 3) &^ 3
			if got := binary.LittleEndian.Uint32(wire[followingOffset:]); got != marker {
				t.Fatalf("following field at wire offset %d = 0x%08x, want 0x%08x", followingOffset, got, marker)
			}
			if wantLen := followingOffset + 4; len(wire) != wantLen {
				t.Fatalf("wire length = %d, want %d", len(wire), wantLen)
			}

			var decoded utf16NWithFollowing
			if err := Unmarshal(wire, &decoded); err != nil {
				t.Fatalf("Unmarshal() error = %v", err)
			}
			if decoded.following != marker {
				t.Fatalf("decoded following field = 0x%08x, want 0x%08x", decoded.following, marker)
			}
		})
	}
}

func TestWriteUTF16StringPreservesInteriorNUL(t *testing.T) {
	const marker = uint32(0xa1b2c3d4)
	wire, err := Marshal(&utf16WithFollowing{s: "a\x00b", following: marker})
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}
	if got := binary.LittleEndian.Uint32(wire[0:]); got != 3 {
		t.Fatalf("max_count = %d, want 3", got)
	}
	for i, want := range []uint16{0x0061, 0x0000, 0x0062} {
		if got := binary.LittleEndian.Uint16(wire[12+i*2:]); got != want {
			t.Fatalf("unit %d = 0x%04x, want 0x%04x", i, got, want)
		}
	}
	if got := binary.LittleEndian.Uint32(wire[20:]); got != marker {
		t.Fatalf("following field = 0x%08x, want 0x%08x", got, marker)
	}

	var decoded utf16WithFollowing
	if err := Unmarshal(wire, &decoded); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}
	if decoded.s != "a\x00b" {
		t.Fatalf("decoded string = %q, want %q", decoded.s, "a\x00b")
	}
	if decoded.following != marker {
		t.Fatalf("decoded following field = 0x%08x, want 0x%08x", decoded.following, marker)
	}
}

func TestReadStringRejectsOversizedCount(t *testing.T) {
	for _, name := range []string{"char", "utf16"} {
		var m0, m1 runtime.MemStats
		runtime.ReadMemStats(&m0)
		b := make([]byte, 12)
		binary.LittleEndian.PutUint32(b[8:], 0xFFFFFFFF)
		var err error
		if name == "char" {
			err = Unmarshal(b, &vChar{})
		} else {
			err = Unmarshal(b, &vU16{})
		}
		runtime.ReadMemStats(&m1)
		mb := (m1.TotalAlloc - m0.TotalAlloc) / (1024 * 1024)
		t.Logf("%s: err=%v alloc=%dMB", name, err, mb)
		if err == nil {
			t.Errorf("%s: expected error for oversized count", name)
		}
		if mb > 64 {
			t.Errorf("%s: allocated %dMB from a 12-byte message (DoS not fixed)", name, mb)
		}
	}
}

// valid short string still decodes
func TestReadCharStringRoundTrip(t *testing.T) {
	// max_count=5, offset=0, actual_count=5, then "hello"
	b := make([]byte, 12+5)
	binary.LittleEndian.PutUint32(b[0:], 5)
	binary.LittleEndian.PutUint32(b[8:], 5)
	copy(b[12:], "hello")
	c := &vChar{}
	if err := Unmarshal(b, c); err != nil {
		t.Fatalf("valid string failed: %v", err)
	}
	if c.s != "hello" {
		t.Fatalf("got %q want hello", c.s)
	}
}
