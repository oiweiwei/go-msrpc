package oaut

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/oiweiwei/go-msrpc/ndr"
)

func TestVariantByRefCharDeferred(t *testing.T) {
	// A sibling uint32 follows the pointer marker before the deferred CHAR.
	wire, err := hex.DecodeString("04000000000000001040000000000000104000000100000078563412f9")
	if err != nil {
		t.Fatal(err)
	}
	var value Variant
	var sibling uint32
	err = ndr.Unmarshal(wire, ndr.UnmarshalNDRFunc(func(ctx context.Context, r ndr.Reader) error {
		if err := value.UnmarshalNDR(ctx, r); err != nil {
			return err
		}
		return r.ReadData(&sibling)
	}))
	if err != nil {
		t.Fatal(err)
	}
	if value.VarUnion.GetValue() != int8(-7) || sibling != 0x12345678 {
		t.Fatalf("value=%#v sibling=%#x", value.VarUnion.GetValue(), sibling)
	}
	encoded, err := ndr.Marshal(ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := value.MarshalNDR(ctx, w); err != nil {
			return err
		}
		return w.WriteData(sibling)
	}))
	if err != nil {
		t.Fatal(err)
	}
	if len(encoded) != len(wire) || binary.LittleEndian.Uint32(encoded[20:24]) == 0 {
		t.Fatalf("invalid deferred CHAR: %x", encoded)
	}
	binary.LittleEndian.PutUint32(encoded[20:24], 1)
	if !bytes.Equal(encoded, wire) {
		t.Fatalf("wire=%x want=%x", encoded, wire)
	}
}

func TestVariantByRefCharWireFixtures(t *testing.T) {
	for _, value := range []int8{0, 127, -128, -7} {
		t.Run(fmt.Sprint(value), func(t *testing.T) {
			// NDR32 VARIANT header, VT_I1|VT_BYREF switch, unique-pointer
			// marker and deferred signed byte. The marker is not a string count.
			wire, err := hex.DecodeString("04000000000000001040000000000000104000000100000000")
			if err != nil {
				t.Fatal(err)
			}
			wire[len(wire)-1] = byte(value)
			var decoded Variant
			if err := ndr.Unmarshal(wire, &decoded); err != nil {
				t.Fatal(err)
			}
			if got := decoded.VarUnion.GetValue(); got != value {
				t.Fatalf("value = %#v (%T), want %d (int8)", got, got, value)
			}
			encoded, err := ndr.Marshal(&decoded)
			if err != nil {
				t.Fatal(err)
			}
			if len(encoded) != len(wire) {
				t.Fatalf("wire length = %d, want %d: %x", len(encoded), len(wire), encoded)
			}
			if got := binary.LittleEndian.Uint32(encoded[20:24]); got == 0 {
				t.Fatal("non-null CHAR encoded as null")
			}
			// Referent identifiers need only be nonzero, not equal to the fixture.
			binary.LittleEndian.PutUint32(encoded[20:24], 1)
			if hex.EncodeToString(encoded) != hex.EncodeToString(wire) {
				t.Fatalf("wire = %x, want %x", encoded, wire)
			}
			for end := range len(wire) {
				if err := ndr.Unmarshal(wire[:end], &Variant{}); err == nil {
					t.Fatalf("accepted truncation at byte %d", end)
				}
			}
		})
	}
}

func TestVariantByRefCharNull(t *testing.T) {
	wire, err := hex.DecodeString("030000000000000010400000000000001040000000000000")
	if err != nil {
		t.Fatal(err)
	}
	var decoded Variant
	previous, err := hex.DecodeString("040000000000000010400000000000001040000001000000f9")
	if err != nil {
		t.Fatal(err)
	}
	if err := ndr.Unmarshal(previous, &decoded); err != nil {
		t.Fatal(err)
	}
	if got := decoded.VarUnion.GetValue(); got != int8(-7) {
		t.Fatalf("initial CHAR = %#v, want -7", got)
	}
	if err := ndr.Unmarshal(wire, &decoded); err != nil {
		t.Fatal(err)
	}
	if got := decoded.VarUnion.GetValue(); got != nil {
		t.Fatalf("null CHAR value = %#v, want nil", got)
	}
	encoded, err := ndr.Marshal(&decoded)
	if err != nil {
		t.Fatal(err)
	}
	if hex.EncodeToString(encoded) != hex.EncodeToString(wire) {
		t.Fatalf("wire = %x, want %x", encoded, wire)
	}
}
