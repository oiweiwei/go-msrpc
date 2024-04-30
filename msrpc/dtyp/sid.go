package dtyp

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
)

func (o *SID) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.String())
}

// String ...
func (o *SID) String() string {

	if o == nil || o.IDAuthority == nil || len(o.IDAuthority.Value) < 6 {
		return "S-1-0-0"
	}

	ret := fmt.Sprintf("S-%d", o.Revision)

	iauth := o.IDAuthority.Value

	if iauth[0] != 0 || iauth[1] != 0 {
		ret += fmt.Sprintf("-0x%x", hex.EncodeToString(iauth))
	} else {
		ret += fmt.Sprintf("-%d", binary.BigEndian.Uint32(iauth[2:]))
	}

	for _, sa := range o.SubAuthority {
		ret += "-" + strconv.Itoa(int(sa))
	}

	return ret
}

func (o *SID) DecodeBinary(b []byte) error {
	if err := o.decodeBinary(b, binary.LittleEndian); err != nil {
		return fmt.Errorf("sid: decode_binary: %v", err)
	}
	return nil
}

func (o *SID) decodeBinary(b []byte, order binary.ByteOrder) error {

	r := bytes.NewReader(b)

	if err := binary.Read(r, order, &o.Revision); err != nil {
		return err
	}

	if err := binary.Read(r, order, &o.SubAuthorityCount); err != nil {
		return err
	}

	o.IDAuthority = &SIDIDAuthority{Value: make([]byte, 6)}

	if _, err := r.Read(o.IDAuthority.Value); err != nil {
		return err
	}

	o.SubAuthority = make([]uint32, o.SubAuthorityCount)
	for i := range o.SubAuthority {
		if err := binary.Read(r, order, &o.SubAuthority[i]); err != nil {
			return err
		}
	}

	return nil
}
