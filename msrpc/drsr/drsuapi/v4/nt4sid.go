package drsuapi

import (
	"encoding/json"

	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
)

// IsZero returns true if the NT4SID is nil or all zeros.
func (o *NT4SID) IsZero() bool {

	if o == nil {
		return true
	}

	for _, v := range o.Data {
		if v != 0 {
			return false
		}
	}
	return true
}

// MarshalJSON implements the json.Marshaler interface.
func (o *NT4SID) MarshalJSON() ([]byte, error) {

	sid, err := o.SID()
	if err != nil {
		return nil, err
	}

	if sid == nil {
		return []byte(`null`), nil
	}

	return json.Marshal(sid.String())
}

// SID returns the SID from the NT4SID.
func (o *NT4SID) SID() (*dtyp.SID, error) {

	if o.IsZero() {
		return nil, nil
	}

	sid := &dtyp.SID{}

	if err := sid.DecodeBinary(o.Data); err != nil {
		return nil, err
	}

	return sid, nil
}
