package uuid

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
)

type UUID struct {
	TimeLow               uint32
	TimeMid               uint16
	TimeHiAndVersion      uint16
	ClockSeqHiAndReserved uint8
	ClockSeqLow           uint8
	Node                  [6]byte
}

func (u *UUID) Equals(ou *UUID) bool {
	if u == nil || ou == nil {
		return ou == u
	}
	return u.TimeLow == ou.TimeLow &&
		u.TimeMid == ou.TimeMid &&
		u.TimeHiAndVersion == ou.TimeHiAndVersion &&
		u.ClockSeqHiAndReserved == ou.ClockSeqHiAndReserved &&
		u.ClockSeqLow == ou.ClockSeqLow &&
		u.Node == ou.Node
}

func New(low uint32, mid, hiAndVers uint16, clockSeqHi, clockSeqLow uint8, node [6]byte) *UUID {
	return &UUID{TimeLow: low, TimeMid: mid, TimeHiAndVersion: hiAndVers,
		ClockSeqHiAndReserved: clockSeqHi, ClockSeqLow: clockSeqLow, Node: node}
}

func (u *UUID) Read(r io.Reader) error {
	b := make([]byte, 16)
	n, err := r.Read(b)
	if err != nil || n != 16 {
		return err
	}
	return u.DecodeBinary(b)
}

func (u *UUID) Write(w io.Writer) error {
	b := u.EncodeBinary()
	_, err := w.Write(b)
	return err
}

func (u *UUID) DecodeBinary(b []byte) error {
	if len(b) < 16 {
		return nil
	}
	u.TimeLow = binary.LittleEndian.Uint32(b[:4])
	u.TimeMid = binary.LittleEndian.Uint16(b[4:6])
	u.TimeHiAndVersion = binary.LittleEndian.Uint16(b[6:8])
	u.ClockSeqHiAndReserved = b[8]
	u.ClockSeqLow = b[9]
	copy(u.Node[:], b[10:])
	return nil
}

func (u *UUID) EncodeBinary() []byte {

	b := make([]byte, 16)
	if u == nil {
		return b
	}
	binary.LittleEndian.PutUint32(b[:4], u.TimeLow)
	binary.LittleEndian.PutUint16(b[4:6], u.TimeMid)
	binary.LittleEndian.PutUint16(b[6:8], u.TimeHiAndVersion)
	b[8] = u.ClockSeqHiAndReserved
	b[9] = u.ClockSeqLow
	copy(b[10:16], u.Node[:])
	return b
}

func (u *UUID) String() string {

	var (
		buf = bytes.NewBuffer(nil)
	)

	binary.Write(buf, binary.BigEndian, u)
	ustr := hex.EncodeToString(buf.Bytes())
	return strings.Join([]string{
		ustr[:8],
		ustr[8:12],
		ustr[12:16],
		ustr[16:20],
		ustr[20:32],
	}, "-")
}

func (u *UUID) MarshalJSON() ([]byte, error) {

	if u == nil {
		return []byte("null"), nil
	}

	return append([]byte{'"'}, append([]byte(u.String()), '"')...), nil
}

func IsUUID(c rune) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
}

func Parse(in string) (*UUID, error) {
	return Unmarshal(in)
}

func MustParse(in string) *UUID {
	u, err := Parse(in)
	if err != nil {
		panic(err)
	}
	return u
}

func Unmarshal(in string) (*UUID, error) {

	var (
		s   []string
		b   []byte
		err error
	)

	if s = strings.Split(in, "-"); len(s) != 5 {
		return nil, fmt.Errorf("uuid: invalid uuid literal: %s", in)
	} else if len(s[0]) != 8 {
		return nil, fmt.Errorf("uuid: invalid uuid literal: invalid time_low: %s", s[0])
	} else if len(s[1]) != 4 {
		return nil, fmt.Errorf("uuid: invalid uuid literal: invalid time_mid: %s", s[1])
	} else if len(s[2]) != 4 {
		return nil, fmt.Errorf("uuid: invalid uuid literal: invalid time_hi_and_version: %s", s[2])
	} else if len(s[3]) != 4 {
		return nil, fmt.Errorf("uuid: invalid uuid literal: invalid clock_seq_hi_and_resverd/clock_seq_row: %s", s[3])
	} else if len(s[4]) != 12 {
		return nil, fmt.Errorf("uuid: invalid uuid literal: invalid node: %s", s[4])
	}

	if b, err = hex.DecodeString(strings.Join(s, "")); err != nil {
		return nil, fmt.Errorf("uuid: invalid uuid literal: invalid hex string: %s", strings.Join(s, "-"))
	}

	var ret = new(UUID)

	if err = binary.Read(bytes.NewBuffer(b), binary.BigEndian, ret); err != nil {
		return nil, fmt.Errorf("uuid: invalid uuid literal: invalid binary: %s", strings.Join(s, "-"))
	}

	return ret, nil
}
