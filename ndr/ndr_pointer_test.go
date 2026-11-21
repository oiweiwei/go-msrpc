package ndr

import (
	"context"
	"encoding/binary"
	"testing"
)

type u32Val struct{ v uint32 }

func (c *u32Val) UnmarshalNDR(ctx context.Context, r Reader) error { return r.ReadData(&c.v) }

// ptrPair reads two pointer fields followed by their deferred data.
func ptrPair(t *testing.T, r Reader, full bool) (a, b *u32Val, aliased int, target **u32Val) {
	return ptrPairMode(t, r, map[bool]string{false: "plain", true: "full"}[full])
}

// ptrPairMode reads with mode: "plain" (ReadPointer), "full" (ReadFullPointer),
// "hookfull" (ReadPointerWithHook with Full), "fullhook" (ReadFullPointerWithHook), "hook" (ReadPointerWithHook, Full unset).
func ptrPairMode(t *testing.T, r Reader, mode string) (a, b *u32Val, aliased int, target **u32Val) {
	t.Helper()
	read := func(dst **u32Val) error {
		v := &u32Val{}
		set := func(p any) {
			aliased++
			target, _ = p.(**u32Val)
		}
		def := UnmarshalNDRFunc(func(ctx context.Context, r Reader) error {
			*dst = v
			return v.UnmarshalNDR(ctx, r)
		})
		switch mode {
		case "plain":
			return r.ReadPointer(dst, set, def)
		case "full":
			return r.ReadFullPointer(dst, set, def)
		case "hookfull":
			return r.ReadPointerWithHook(dst, PointerHook{OnSet: set, Full: true}, def)
		case "fullhook":
			return r.ReadFullPointerWithHook(dst, PointerHook{OnSet: set}, def)
		}
		return r.ReadPointerWithHook(dst, PointerHook{OnSet: set}, def)
	}
	if err := read(&a); err != nil {
		t.Fatal(err)
	}
	if err := read(&b); err != nil {
		t.Fatal(err)
	}
	if err := r.ReadDeferred(); err != nil {
		t.Fatal(err)
	}
	return
}

func u32s(v ...uint32) []byte {
	b := make([]byte, 4*len(v))
	for i := range v {
		binary.LittleEndian.PutUint32(b[4*i:], v[i])
	}
	return b
}

func TestUniquePointersSameReferentIDDoNotAlias(t *testing.T) {
	r := NDR20(u32s(1, 1, 0xAAAA, 0xBBBB))
	a, b, al, _ := ptrPair(t, r, false)
	if al != 0 || a == nil || b == nil || a == b || a.v != 0xAAAA || b.v != 0xBBBB {
		t.Fatalf("unique: aliased=%d a=%v b=%v", al, a, b)
	}
}

func TestFullPointersSameReferentIDAlias(t *testing.T) {
	r := NDR20(u32s(1, 1, 0xAAAA, 0xBBBB))
	a, b, al, target := ptrPair(t, r, true)
	if al != 1 || a == nil || a.v != 0xAAAA || b != nil || target == nil || *target != a {
		t.Fatalf("full: aliased=%d a=%v b=%v target=%v", al, a, b, target)
	}
	// the second pointer carries no data: the 0xBBBB word is left in the buffer.
	var rest uint32
	if err := r.ReadData(&rest); err != nil || rest != 0xBBBB {
		t.Fatalf("remaining word = %#x err=%v", rest, err)
	}
}

func TestFullPointersDistinctIDsDoNotAlias(t *testing.T) {
	r := NDR20(u32s(1, 2, 0xAAAA, 0xBBBB))
	a, b, al, _ := ptrPair(t, r, true)
	if al != 0 || a == b || a.v != 0xAAAA || b.v != 0xBBBB {
		t.Fatalf("full distinct: aliased=%d a=%v b=%v", al, a, b)
	}
}

func TestUserMarkerUnchanged(t *testing.T) {
	for _, full := range []bool{false, true} {
		// 'User' marker twice: both are deferred with data, never aliased.
		r := NDR20(u32s(0x72657355, 0x72657355, 7, 9))
		a, b, al, _ := ptrPair(t, r, full)
		if al != 0 || a == b || a.v != 7 || b.v != 9 {
			t.Fatalf("User full=%v: aliased=%d a=%v b=%v", full, al, a, b)
		}
	}
}

func TestNDR64PointersNeverAlias(t *testing.T) {
	for _, full := range []bool{false, true} {
		b := make([]byte, 0, 32)
		for _, v := range []uint64{1, 1} {
			b = binary.LittleEndian.AppendUint64(b, v)
		}
		b = binary.LittleEndian.AppendUint32(b, 0xAAAA)
		b = binary.LittleEndian.AppendUint32(b, 0xBBBB)
		r := NDR64(b)
		a, c, al, _ := ptrPair(t, r, full)
		if al != 0 || a == c || a.v != 0xAAAA || c.v != 0xBBBB {
			t.Fatalf("ndr64 full=%v: aliased=%d a=%v c=%v", full, al, a, c)
		}
	}
}

func TestFullPointerHookFormsAlias(t *testing.T) {
	for _, mode := range []string{"hookfull", "fullhook"} {
		r := NDR20(u32s(1, 1, 0xAAAA, 0xBBBB))
		a, b, al, target := ptrPairMode(t, r, mode)
		if al != 1 || a == nil || a.v != 0xAAAA || b != nil || target == nil || *target != a {
			t.Fatalf("%s: aliased=%d a=%v b=%v", mode, al, a, b)
		}
	}
	r := NDR20(u32s(1, 1, 0xAAAA, 0xBBBB))
	a, b, al, _ := ptrPairMode(t, r, "hook")
	if al != 0 || a == b || a.v != 0xAAAA || b.v != 0xBBBB {
		t.Fatalf("hook without Full aliased: %d %v %v", al, a, b)
	}
}
