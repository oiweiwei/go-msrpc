package binxml

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"

	"github.com/oiweiwei/go-msrpc/midl/uuid"
	"github.com/oiweiwei/go-msrpc/msrpc/dtyp"
)

type Value struct {
	ID      int           `json:"id"`
	Length  uint16        `json:"length,omitempty"`
	Type    uint8         `json:"type"`
	IsArray bool          `json:"is_array,omitempty"`
	Data    []interface{} `json:"data,omitempty"`
	Raw     []byte        `json:"-"`
}

func (v *Value) push(d interface{}) {
	v.Data = append(v.Data, d)
}

func (v *Value) Decode(r *Reader) error {

	r.Begin("value")

	r.Read(&v.Length)

	if r.Read(&v.Type); v.Type > 0x80 {
		v.Type, v.IsArray = v.Type-0x80, true
	}

	// pad.
	r.ReadTag(0x00)

	return r.Done()
}

const (
	Null       = 0x00
	String     = 0x01
	ANSIString = 0x02
	Int8       = 0x03
	Uint8      = 0x04
	Int16      = 0x05
	Uint16     = 0x06
	Int32      = 0x07
	Uint32     = 0x08
	Int64      = 0x09
	Uint64     = 0x0A
	Float32    = 0x0B
	Float64    = 0x0C
	Bool       = 0x0D
	Binary     = 0x0E
	GUID       = 0x0F
	SizeT      = 0x10
	FileTime   = 0x11
	SysTime    = 0x12
	SID        = 0x13
	HexInt32   = 0x14
	HexInt64   = 0x15
	BinXML     = 0x21
)

func (o *Value) DecodeValue(b []byte) error {

	r := NewReader(nil)

	r.Begin(fmt.Sprintf("value_%d", o.Type))

	o.Raw = b

	if len(b) == 0 {
		return nil
	}

	switch o.Type {
	case Null:
		o.Data = nil
	case String:

		if !o.IsArray {
			var s string
			r.ReadWithBytes(b, func(r *Reader) error {
				return r.ReadUTF16String(&s, len(b)/2)
			})
			o.push(s)
			break
		}

		start := 0
		for i := 0; i < len(b); i += 2 {
			if b[i] == 0x00 && b[i+1] == 0x00 {
				var s string
				r.ReadWithBytes(b[start:i], func(r *Reader) error {
					return r.ReadUTF16String(&s, len(b[start:i])/2)
				})
				o.push(s)
				start = i + 2
			}
		}
	case ANSIString:
		for _, b := range bytes.Split(b, []byte{0x00}) {
			o.push(string(b))
		}
	case Int8:
		for i := 0; i < len(b); i++ {
			o.push(int8(b[i]))
		}
	case Uint8:
		for i := 0; i < len(b); i++ {
			o.push(uint8(b[i]))
		}
	case Int16:
		for i := 0; i < len(b); i += 2 {
			o.push(int16(r.Uint16(b[i : i+2])))
		}
	case Uint16:
		for i := 0; i < len(b); i += 2 {
			o.push(r.Uint16(b[i : i+2]))
		}
	case Int32:
		for i := 0; i < len(b); i += 4 {
			o.push(int32(r.Uint32(b[i : i+4])))
		}
	case Uint32:
		for i := 0; i < len(b); i += 4 {
			o.push(r.Uint32(b[i : i+4]))
		}
	case Int64:
		for i := 0; i < len(b); i += 8 {
			o.push(int64(r.Uint64(b[i : i+8])))
		}
	case Uint64:
		for i := 0; i < len(b); i += 8 {
			o.push(r.Uint64(b[i : i+8]))
		}
	case Float32:
		for i := 0; i < len(b); i += 4 {
			o.push(math.Float32frombits(r.Uint32(b[i : i+4])))
		}
	case Float64:
		for i := 0; i < len(b); i += 8 {
			o.push(math.Float64frombits(r.Uint64(b[i : i+8])))
		}
	case Bool:
		for i := 0; i < len(b); i++ {
			o.push(b[i] != 0)
		}
	case Binary:
		o.push(hex.EncodeToString(b))
	case GUID:
		for i := 0; i < len(b); i += 16 {
			uuid := &uuid.UUID{}
			if err := uuid.DecodeBinary(b[i : i+16]); err != nil {
				return err
			}
			o.push(uuid)
		}
	case SizeT:
		panic("size_t is not supported")
		/*
			switch SizeTSize {
			case 8:
				for i := 0; i < len(b); i += 8 {
					o.push(binary.LittleEndian.Uint64(b[i : i+8]))
				}
			case 4:
				for i := 0; i < len(b); i += 4 {
					o.push(binary.LittleEndian.Uint32(b[i : i+4]))
				}
			}
		*/
	case FileTime:
		for i := 0; i < len(b); i += 8 {
			t := &dtyp.Filetime{}
			r.ReadWithBytes(b[i:i+8], t)
			o.push(t.AsTime())
		}
	case SysTime:
		for i := 0; i < len(b); i += 16 {
			t := &dtyp.SystemTime{}
			r.ReadWithBytes(b[i:i+16], t)
			o.push(t.AsTime())
		}
	case SID:
		for len(b) > 0 {
			sid := &dtyp.SID{}
			if err := sid.DecodeBinary(b); err != nil {
				return err
			}
			b = b[8+sid.SubAuthorityCount*4:]
			o.push(sid)
		}
	case HexInt32:
		for i := 0; i < len(b); i += 4 {
			o.push(fmt.Sprintf("0x%04x", binary.LittleEndian.Uint32(b[i:i+4])))
		}
	case HexInt64:
		for i := 0; i < len(b); i += 8 {
			o.push(fmt.Sprintf("0x%08x", binary.LittleEndian.Uint64(b[i:i+8])))
		}
	case BinXML:
		v := &Document{}
		r.ReadWithBytes(b, &v)
		o.push(v)
	}

	return r.Done()
}
