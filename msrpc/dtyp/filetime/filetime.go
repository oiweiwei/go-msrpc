package filetime

import (
	"time"

	"github.com/oiweiwei/go-msrpc/msrpc/dtyp"
)

// Never returns a Filetime that represents the maximum time.
func Never() *dtyp.Filetime {
	return &dtyp.Filetime{
		LowDateTime:  0xFFFFFFFF,
		HighDateTime: 0x7FFFFFFF,
	}
}

// Now returns a Filetime that represents the current UTC time.
func Now() *dtyp.Filetime {
	return FromTime(time.Now())
}

// FromTime returns a Filetime from a time.Time.
func FromTime(t time.Time) *dtyp.Filetime {
	ft := (t.UnixNano() / 100) + 116444736000000000
	return &dtyp.Filetime{
		LowDateTime:  uint32(ft),
		HighDateTime: uint32(ft >> 32),
	}
}
