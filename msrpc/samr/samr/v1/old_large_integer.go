package samr

import "encoding/json"

func (o *OldLargeInteger) Uint64() uint64 {
	return uint64(o.LowPart) + uint64(o.HighPart)<<32
}

func (o *OldLargeInteger) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.Uint64())
}
