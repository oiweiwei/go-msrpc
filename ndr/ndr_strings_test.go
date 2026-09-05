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
