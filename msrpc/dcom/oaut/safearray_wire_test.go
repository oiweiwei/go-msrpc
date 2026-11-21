package oaut

import (
	"encoding/hex"
	"testing"

	"github.com/oiweiwei/go-msrpc/ndr"
)

// RESP stubs of EXP-890 (oleaut32 on Windows Server 2019, IHarnessSource echo of SAFEARRAY(BSTR) and
// SAFEARRAY(VARIANT)): ORPCTHAT, the 'User' marker and the word after it are stripped, the rest is
// the SAFEARRAY wire form followed by the HRESULT. oleaut32 writes the referent id 1 for every
// [unique] BSTR pointer of the array: they must not alias.
const (
	saBstrResp    = "00000000000000005573657201000000010000000100800104000000000008000800000003000000b0756205030000000000000003000000010000000100000001000000060000000c0000000600000061006c00700068006100210001000000020000000100000021000000050000000a000000050000006400e9006a00e0002100000000000000"
	saVariantResp = "0000000000000000557365720100000001000000010080081000000000000c000c00000005000000804a6105050000000000000005000000557365725573657255736572557365725573657200000000060000000000000008000000000000000800000019596105060000000c00000006000000680065006c006c006f00210003000000000000000300000000000000030000002a00000004000000000000000500000000000000050000000000000000000000000008400f00000000000000082000000000000000200000f1bf600501000000010000000100800104000000000008000800000002000000502d5f0502000000000000000200000001000000010000000300000006000000030000006e003100210000000300000006000000030000006e0032002100000000000000030000000000000000000000000000000000000000000000"
)

// saWire strips ORPCTHAT (8 bytes), the 'User' marker and the word after it (4+4 bytes; value 1).
func saWire(resp string) string { return resp[32:] }

func decodeSA(t *testing.T, h string) *SafeArray {
	t.Helper()
	b, err := hex.DecodeString(h)
	if err != nil {
		t.Fatal(err)
	}
	var sa SafeArray
	if err := ndr.Unmarshal(b, &sa); err != nil {
		t.Fatalf("decode: %v", err)
	}
	return &sa
}

func TestSafeArrayBstrRepeatedReferentIDs(t *testing.T) {
	sa := decodeSA(t, saWire(saBstrResp))
	s := sa.ArrayStructs.Value.(*SafeArrayUnion_String).String
	want := []string{"alpha!", "!", "d\u00e9j\u00e0!"}
	if len(s.String) != len(want) {
		t.Fatalf("elements = %d, want %d", len(s.String), len(want))
	}
	for i, w := range want {
		if s.String[i] == nil || s.String[i].Data != w {
			t.Errorf("element %d = %+v, want %q", i, s.String[i], w)
		}
	}
}

func TestSafeArrayVariantRepeatedReferentIDs(t *testing.T) {
	sa := decodeSA(t, saWire(saVariantResp))
	v := sa.ArrayStructs.Value.(*SafeArrayUnion_Variant).Variant
	if len(v.Variant) != 5 {
		t.Fatalf("elements = %d", len(v.Variant))
	}
	bstr := v.Variant[0].VarUnion.Value.(*Variant_VarUnion_BSTR).BSTR
	if v.Variant[0].VT != 8 || bstr == nil || bstr.Data != "hello!" {
		t.Errorf("element 0 = %+v %+v", v.Variant[0], bstr)
	}
	if l := v.Variant[1].VarUnion.Value.(*Variant_VarUnion_Long).Long; v.Variant[1].VT != 3 || l != 42 {
		t.Errorf("element 1 = %+v", v.Variant[1])
	}
	nested := v.Variant[3].VarUnion.Value.(*Variant_VarUnion_SafeArray).SafeArray
	if v.Variant[3].VT != 0x2008 || nested == nil {
		t.Fatalf("element 3 = %+v", v.Variant[3])
	}
	ns := nested.ArrayStructs.Value.(*SafeArrayUnion_String).String.String
	if len(ns) != 2 || ns[0] == nil || ns[0].Data != "n1!" || ns[1] == nil || ns[1].Data != "n2!" {
		t.Errorf("nested = %+v", ns)
	}
}
