package dtyp

import (
	"encoding/binary"
	"encoding/json"
	"time"
)

func (ft *Filetime) MarshalJSON() ([]byte, error) {
	if ft.IsNever() {
		return json.Marshal("never")
	}
	return json.Marshal(ft.AsTime())
}

func (ft *Filetime) IsNever() bool {
	return ft.LowDateTime == 0xFFFFFFFF && ft.HighDateTime == 0x7FFFFFFF
}

// AsTime function returns the time.Time.
func (ft *Filetime) AsTime() time.Time {
	if ft.LowDateTime == 0 && ft.HighDateTime == 0 {
		return time.Time{}
	}
	// 100-nanosecond intervals since January 1, 1601
	nsec := (int64(ft.HighDateTime) << 32) + int64(ft.LowDateTime)
	// change starting time to the Epoch (00:00:00 UTC, January 1, 1970)
	nsec -= 116444736000000000
	// convert into nanoseconds
	// nsec *= 100
	return time.Unix(nsec/10000000, nsec%10000000).UTC()
}

func (ft *Filetime) DecodeBinary(b []byte) error {
	u := binary.LittleEndian.Uint64(b)
	ft.HighDateTime = uint32(u >> 32)
	ft.LowDateTime = uint32(u)
	return nil
}
