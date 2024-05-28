package dtyp

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
)

func (g *GUID) String() string {
	if g != nil && len(g.Data4) == 8 {
		return fmt.Sprintf("%08x-%04x-%04x-%02x%02x-%012s",
			g.Data1, g.Data2, g.Data3, g.Data4[0], g.Data4[1], hex.EncodeToString(g.Data4[2:]))
	}
	return ""
}

func (g *GUID) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.String())
}

func GUIDFromBytes(b []byte) (*GUID, error) {

	u := &uuid.UUID{}
	if err := u.DecodeBinary(b); err != nil {
		return nil, err
	}

	return GUIDFromUUID(u), nil
}

func GUIDFromUUID(u *uuid.UUID) *GUID {
	return &GUID{
		Data1: u.TimeLow,
		Data2: u.TimeMid,
		Data3: u.TimeHiAndVersion,
		Data4: []byte{u.ClockSeqHiAndReserved, u.ClockSeqLow, u.Node[0], u.Node[1], u.Node[2], u.Node[3], u.Node[4], u.Node[5]},
	}
}

func (o *GUID) UUID() *uuid.UUID {

	data4 := o.Data4
	if len(data4) < 8 {
		data4 = make([]byte, 8)
	}

	return &uuid.UUID{
		TimeLow:               o.Data1,
		TimeMid:               o.Data2,
		TimeHiAndVersion:      o.Data3,
		ClockSeqHiAndReserved: data4[0],
		ClockSeqLow:           data4[1],
		Node:                  [6]byte{data4[2], data4[3], data4[4], data4[5], data4[6], data4[7]},
	}
}
