package oaut

import (
	"encoding/binary"
	"encoding/hex"
	"testing"

	"github.com/oiweiwei/go-msrpc/ndr"
)

func TestVariantTypedNullSafeArray(t *testing.T) {
	for _, vt := range []uint16{0x2003, 0x2005, 0x2008, 0x200c, 0x6003, 0x6005, 0x6008, 0x600c} {
		value := &Variant{Size: 3, VT: vt, VarUnion: &Variant_VarUnion{}}
		if vt&0x4000 != 0 {
			value.VarUnion.Value = &Variant_VarUnion_SafeArrayPtr{}
		} else {
			value.VarUnion.Value = &Variant_VarUnion_SafeArray{}
		}
		wire, err := ndr.Marshal(value)
		if err != nil {
			t.Errorf("VT %#x: marshal: %v", vt, err)
			continue
		}
		if len(wire) != 24 {
			t.Errorf("VT %#x: wire length = %d, want 24", vt, len(wire))
			continue
		}
		if got := binary.LittleEndian.Uint16(wire[8:]); got != vt {
			t.Errorf("VT %#x: header VT = %#x", vt, got)
		}
		wantSwitch := uint32(0x2000)
		if vt&0x4000 != 0 {
			wantSwitch = 0x6000
		}
		if got := binary.LittleEndian.Uint32(wire[16:]); got != wantSwitch {
			t.Errorf("VT %#x: union switch = %#x, want %#x", vt, got, wantSwitch)
		}
		var decoded Variant
		if err := ndr.Unmarshal(wire, &decoded); err != nil {
			t.Errorf("VT %#x: unmarshal: %v", vt, err)
		} else if decoded.VT != vt {
			t.Errorf("decoded VT = %#x, want %#x", decoded.VT, vt)
		}
	}
}

func TestVariantTypedSafeArrayWireFixture(t *testing.T) {
	// NDR32 VT_ARRAY|VT_I4, two elements (1, 2), lower bound -2.
	// The header VT is 0x2003, while the separate union switch is 0x2000.
	wire, err := hex.DecodeString("0a0000000000000003200000000000000020000001000000020000000100000001008000040000000000030003000000020000000300000002000000feffffff020000000100000002000000")
	if err != nil {
		t.Fatal(err)
	}
	var value Variant
	if err := ndr.Unmarshal(wire, &value); err != nil {
		t.Fatal(err)
	}
	if value.VT != 0x2003 {
		t.Fatalf("VT = %#x, want 0x2003", value.VT)
	}
	array := value.VarUnion.Value.(*Variant_VarUnion_SafeArray).SafeArray
	if len(array.Bound) != 1 || array.Bound[0].ElementsCount != 2 || array.Bound[0].LowerBound != -2 {
		t.Fatalf("array bounds = %+v", array.Bound)
	}
	encoded, err := ndr.Marshal(&value)
	if err != nil {
		t.Fatal(err)
	}
	if got := binary.LittleEndian.Uint32(encoded[16:]); got != 0x2000 {
		t.Fatalf("union switch = %#x, want 0x2000", got)
	}
	var decoded Variant
	if err := ndr.Unmarshal(encoded, &decoded); err != nil {
		t.Fatal(err)
	}
	values := decoded.VarUnion.Value.(*Variant_VarUnion_SafeArray).SafeArray.ArrayStructs.Value.(*SafeArrayUnion_Long).Long
	if values.Size != 2 || len(values.Data) != 2 || values.Data[0] != 1 || values.Data[1] != 2 {
		t.Fatalf("decoded values = %+v", values)
	}
	for end := range len(wire) {
		if err := ndr.Unmarshal(wire[:end], &Variant{}); err == nil {
			t.Fatalf("accepted truncation at byte %d", end)
		}
	}
}
