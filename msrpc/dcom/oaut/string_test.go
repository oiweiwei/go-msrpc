package oaut

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/oiweiwei/go-msrpc/ndr"
)

func TestStringWireFixtures(t *testing.T) {
	// NDR32: conformant count, cBytes, clSize, then clSize UTF-16 code units.
	// These fixtures are independent of the generated encoder.
	for _, tt := range []struct {
		name  string
		value String
		wire  string
	}{
		{"empty", String{}, "000000000000000000000000"},
		{"null", String{BytesCount: 0xffffffff}, "00000000ffffffff00000000"},
		{"embedded_nul", String{Data: "a\x00b"}, "030000000600000003000000610000006200"},
		{"trailing_nuls", String{Data: "a\x00\x00"}, "030000000600000003000000610000000000"},
		{"only_nul", String{Data: "\x00"}, "0100000002000000010000000000"},
		{"supplementary", String{Data: "\U0001f600\x00"}, "0300000006000000030000003dd800de0000"},
		// MS-OAUT allows an odd byte count; clSize rounds up to code units.
		{"odd_byte_count", String{BytesCount: 1, Data: "a"}, "0100000001000000010000006100"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			wire, err := hex.DecodeString(tt.wire)
			if err != nil {
				t.Fatal(err)
			}
			for _, blob := range []bool{false, true} {
				value := tt.value
				var encoded []byte
				if blob {
					encoded, err = ndr.Marshal((*FlaggedWordBlob)(&value))
				} else {
					encoded, err = ndr.Marshal(&value)
				}
				if err != nil {
					t.Fatal(err)
				}
				if !bytes.Equal(encoded, wire) {
					t.Errorf("blob=%t: wire = %x, want %x", blob, encoded, wire)
				}
				var decoded String
				if blob {
					err = ndr.Unmarshal(wire, (*FlaggedWordBlob)(&decoded))
				} else {
					err = ndr.Unmarshal(wire, &decoded)
				}
				if err != nil {
					t.Fatal(err)
				}
				if decoded != value {
					t.Errorf("blob=%t: decoded = %#v, want %#v", blob, decoded, value)
				}
				for end := range len(wire) {
					if err := ndr.Unmarshal(wire[:end], &String{}); err == nil {
						t.Fatalf("accepted truncation at byte %d", end)
					}
				}
			}
		})
	}
}
