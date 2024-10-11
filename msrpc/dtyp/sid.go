package dtyp

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	ndr "github.com/oiweiwei/go-msrpc/ndr"
)

func (o *SID) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.String())
}

func (o *SID) Bytes() ([]byte, error) {
	return ndr.Marshal(o, ndr.Opaque)
}

func ParseSID(s string) (*SID, error) {
	o := new(SID)
	if err := o.Parse(s); err != nil {
		return nil, err
	}
	return o, nil
}

func (o *SID) Parse(s string) error {

	if s == "S-1-0-0" || s == "" || o == nil {
		return nil
	}

	if !strings.HasPrefix(s, "S-") {
		return fmt.Errorf("sid: parse: invalid prefix")
	}

	*o = SID{}

	for i, p := range strings.Split(s[2:], "-") {
		switch i {
		case 0:
			revision, err := strconv.ParseUint(p, 10, 8)
			if err != nil {
				return fmt.Errorf("sid: parse: revision: %v", err)
			}
			o.Revision = uint8(revision)
		case 1:
			if strings.HasPrefix(p, "0x") {
				iauth, err := hex.DecodeString(p[2:])
				if err != nil {
					return fmt.Errorf("sid: parse: idauthority: %v", err)
				}
				o.IDAuthority = &SIDIDAuthority{Value: iauth}
			} else {
				iauth, err := strconv.ParseUint(p, 10, 32)
				if err != nil {
					return fmt.Errorf("sid: parse: idauthority: %v", err)
				}
				o.IDAuthority = &SIDIDAuthority{Value: make([]byte, 6)}
				binary.BigEndian.PutUint32(o.IDAuthority.Value[2:], uint32(iauth))
			}
		default:
			subauth, err := strconv.ParseUint(p, 10, 32)
			if err != nil {
				return fmt.Errorf("sid: parse: subauthority: %v", err)
			}
			o.SubAuthority = append(o.SubAuthority, uint32(subauth))
			o.SubAuthorityCount++
		}
	}

	return nil
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
