package oaut

import (
	"bytes"
	"testing"

	"github.com/oiweiwei/go-msrpc/ndr"
)

func TestStringCanonical(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input *String
		want  String
	}{
		{"nil", nil, String{BytesCount: 0xFFFFFFFF}},
		{"null", NullString(), String{BytesCount: 0xFFFFFFFF}},
		{"empty", EmptyString(), String{IsEmpty: true}},
		{"zero_size", &String{Data: "ignored"}, String{IsEmpty: true}},
		{"ascii", &String{Data: "abc", Size: 3}, String{Data: "abc", Size: 3, BytesCount: 6}},
		{"ascii_padding", &String{Data: "a", Size: 3}, String{Data: "a\x00\x00", Size: 3, BytesCount: 6}},
		{"ascii_truncation", &String{Data: "abc", Size: 2}, String{Data: "ab", Size: 2, BytesCount: 4}},
		{"nul_only", &String{Size: 2}, String{Data: "\x00\x00", Size: 2, BytesCount: 4}},
		{"bmp", &String{Data: "\u4e2d", Size: 1}, String{Data: "\u4e2d", Size: 1, BytesCount: 2}},
		{"bmp_padding", &String{Data: "\u4e2d", Size: 2}, String{Data: "\u4e2d\x00", Size: 2, BytesCount: 4}},
		{"accented_padding", &String{Data: "\u00e9", Size: 2}, String{Data: "\u00e9\x00", Size: 2, BytesCount: 4}},
		{"supplementary", &String{Data: "\U0001f680", Size: 2}, String{Data: "\U0001f680", Size: 2, BytesCount: 4}},
		{"supplementary_padding", &String{Data: "\U0001f680", Size: 3}, String{Data: "\U0001f680\x00", Size: 3, BytesCount: 6}},
		{"supplementary_truncation", &String{Data: "\U0001f680x", Size: 2}, String{Data: "\U0001f680", Size: 2, BytesCount: 4}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var before String
			if tc.input != nil {
				before = *tc.input
			}
			got := tc.input.Canonical()
			if *got != tc.want {
				t.Fatalf("Canonical()=%+v, want %+v", *got, tc.want)
			}
			if tc.input != nil && *tc.input != before {
				t.Fatal("Canonical mutated its receiver")
			}
			if again := got.Canonical(); *again != *got {
				t.Fatalf("Canonical is not idempotent: %+v then %+v", *got, *again)
			}
		})
	}
}

func TestStringCanonicalRestoresDecodedPayload(t *testing.T) {
	for _, tc := range []struct {
		name string
		text string
	}{
		{"ascii", "a\x00"},
		{"nul_only", "\x00\x00"},
		{"bmp", "\u4e2d\x00\x00"},
		{"supplementary", "\U0001f680\x00"},
		{"mixed", "a\x00\u4e2d\U0001f680\x00"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			wire, err := ndr.Marshal(&String{Data: tc.text})
			if err != nil {
				t.Fatal(err)
			}
			var decoded String
			if err := ndr.Unmarshal(wire, &decoded); err != nil {
				t.Fatal(err)
			}
			canonical := decoded.Canonical()
			if canonical.Data != tc.text {
				t.Fatalf("restored=%q, want %q", canonical.Data, tc.text)
			}
			reencoded, err := ndr.Marshal(canonical)
			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(reencoded, wire) {
				t.Fatalf("wire=%x, want %x", reencoded, wire)
			}
		})
	}
}
