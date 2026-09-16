package oaut

import (
	"strings"

	"github.com/oiweiwei/go-msrpc/ndr"
)

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

	data := o.Data

	if l := ndr.UTF16Len(data); l < uint64(o.Size) {
		data += strings.Repeat("\x00", int(o.Size)-int(l))
	}

	data = data[:int(o.Size)]

	return &String{
		BytesCount: o.Size * 2,
		Size:       o.Size,
		Data:       data,
		IsEmpty:    len(data) == 0,
	}
}
