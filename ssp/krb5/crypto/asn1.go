package crypto

import (
	"encoding/asn1"
	"fmt"
)

var (
	KRB5OID = asn1.ObjectIdentifier{1, 2, 840, 113554, 1, 2, 2}
)

// IsASN1Value function returns true if the given byte slice is an ASN.1 value
// with application tag.
func IsASN1Value(b []byte) bool {
	return len(b) > 0 && b[0] == 0x60
}

// EncodeASN1Value function encodes the given byte slice into an ASN.1 value
// with the given object identifier. If dceStyle is true, the function will
// encode the value in DCE style, else it will encode the extra value size into
// the header.
func EncodeASN1Value(b []byte, oid asn1.ObjectIdentifier, dceStyle bool, extra [][]byte) ([]byte, error) {
	if dceStyle {
		return EncodeASN1ValueWithExtraSize(b, oid)
	}
	return EncodeASN1ValueWithExtraSize(b, oid, extra...)
}

// EncodeASN1ValueWithExtraSize function encodes the given byte slice into an
// ASN.1 value with the given object identifier. The function will encode the
// extra value size into the header.
func EncodeASN1ValueWithExtraSize(b []byte, oid asn1.ObjectIdentifier, extra ...[]byte) ([]byte, error) {

	ob, err := asn1.Marshal(oid)
	if err != nil {
		return nil, fmt.Errorf("marshal oid: %w", err)
	}

	// allocate buffer for header.
	h := make([]byte, 6 /* app-tag and size */ +len(ob)+len(b))

	// compute size that includes the header data, oid data and extra data
	// size.
	sz := uint32(len(b)) + uint32(len(ob))
	for _, e := range extra {
		sz += uint32(len(e))
	}

	// write asn1 tag and size.
	n, err := EncodeASN1HeaderTo(h, sz)
	if err != nil {
		return nil, fmt.Errorf("encode asn1 size header: %w", err)
	}

	// copy oid and data.
	copy(h[n:], ob)
	copy(h[n+len(ob):], b)

	if len(extra) != 0 {
		// leave padding bytes in the end.
		return h, nil
	}

	// use exact size.
	return h[:n+len(ob)+len(b)], nil
}

// ParseASN1Value function parses the given byte slice as an ASN.1 value. The
// function returns the remaining data, the object identifier and an error if
// any.
func ParseASN1Value(b []byte, extra [][]byte) ([]byte, asn1.ObjectIdentifier, error) {

	var oid asn1.ObjectIdentifier

	if !IsASN1Value(b) {
		return b, oid, nil
	}

	b, sz, ok := DecodeASN1Header(b)
	if !ok {
		return nil, oid, fmt.Errorf("invalid asn1 size header")
	}

	extraL := uint32(0)
	for _, e := range extra {
		extraL += uint32(len(e))
	}

	if sz > uint32(len(b)) && extraL < sz {
		sz -= extraL
	}

	b = b[:sz]

	// trim oid.
	b, err := asn1.Unmarshal(b, &oid)
	if err != nil {
		return nil, oid, fmt.Errorf("unmarshal oid: %w", err)
	}

	return b, oid, nil
}

// EncodeASN1Header function encodes the given size into an ASN.1 header and
// returns the encoded byte slice. The function returns an error if any.
func EncodeASN1Header(sz uint32) ([]byte, error) {
	b := make([]byte, 5)
	n, err := EncodeASN1HeaderTo(b, sz)
	return b[:n], err
}

// EncodeASN1HeaderTo function encodes the given size into an ASN.1 header and
// writes it to the given byte slice. The function returns the number of bytes
// written and an error if any.
func EncodeASN1HeaderTo(b []byte, sz uint32) (int, error) {

	if len(b) < 5 {
		return 0, fmt.Errorf("buffer too small")
	}

	b[0] = 0x60

	cnt := 2

	if sz <= 0x7F {
		b[1] = uint8(sz)
		return cnt, nil
	}

	switch {
	case sz > 0xFFFFFF:
		b[cnt], cnt = uint8(sz>>24), cnt+1
		fallthrough
	case sz > 0xFFFF:
		b[cnt], cnt = uint8(sz>>16), cnt+1
		fallthrough
	case sz > 0xFF:
		b[cnt], cnt = uint8(sz>>8), cnt+1
		fallthrough
	default:
		b[cnt], cnt = uint8(sz), cnt+1
	}

	b[1] = 0x80 | (uint8(cnt) - 2)
	return cnt, nil
}

// DecodeASN1Header function decodes the given byte slice as an ASN.1 header and
// returns the remaining data, the size and a boolean value indicating if the
// header is valid.
func DecodeASN1Header(b []byte) ([]byte, uint32, bool) {

	if len(b) == 0 || b[0] != 0x60 {
		// not an asn1 value.
		return b, uint32(len(b)), true
	}

	if b = b[1:]; len(b) < 1 {
		// trim asn1 app tag.
		return nil, 0, false
	}

	var sz uint32

	if b, sz = b[1:], uint32(b[0]); sz <= 0x7F {
		// trim primitive size.
		return b, sz, true
	}

	if sz&0x80 != 0x80 || (sz&0x0F) > uint32(len(b)) {
		// check number of bytes.
		return nil, 0, false
	}

	sz &= 0x0F

	var cnt int

	switch cnt = int(sz); cnt {
	case 1:
		b, sz = b[cnt:], uint32(b[0])
	case 2:
		b, sz = b[cnt:], uint32(b[0])<<8|uint32(b[1])
	case 3:
		b, sz = b[cnt:], uint32(b[0])<<16|uint32(b[1])<<8|uint32(b[2])
	case 4:
		b, sz = b[cnt:], uint32(b[0])<<24|uint32(b[1])<<16|uint32(b[2])<<8|uint32(b[3])
	default:
		return nil, 0, false
	}

	return b, sz, true
}
