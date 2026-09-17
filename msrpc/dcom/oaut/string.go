package oaut

import "unicode/utf16"

// NullString returns the NULL-data representation of BSTR blob.
func NullString() *String {
	return &String{BytesCount: 0xFFFFFFFF}
}

// EmptyString returns the empty-data representation of BSTR blob.
func EmptyString() *String {
	return &String{IsEmpty: true}
}

// Canonical returns the BSTR canonical representation. All trimmed NULL-s are
// added to the data buffer, BytesCount is populated to respective value.
func (o *String) Canonical() *String {

	if o == nil || o.BytesCount == 0xFFFFFFFF {
		return &String{BytesCount: 0xFFFFFFFF}
	}

	// Size counts UTF-16 code units, not Go string bytes.
	data := make([]uint16, o.Size)
	copy(data, utf16.Encode([]rune(o.Data)))

	return &String{
		BytesCount: o.Size * 2,
		Size:       o.Size,
		Data:       string(utf16.Decode(data)),
		IsEmpty:    len(data) == 0,
	}
}
