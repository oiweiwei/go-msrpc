package wmio

import (
	"encoding/json"
	"fmt"
)

type CIMType uint32

const CIMArray CIMType = 0x2000

const Inherited uint32 = 0x4000

var (
	Int8           CIMType = 16
	Uint8          CIMType = 17
	Int16          CIMType = 2
	Uint16         CIMType = 18
	Int32          CIMType = 3
	Uint32         CIMType = 19
	Int64          CIMType = 20
	Uint64         CIMType = 21
	Float32        CIMType = 4
	Float64        CIMType = 5
	Bool           CIMType = 11
	String         CIMType = 8
	DateTime       CIMType = 101
	Ref            CIMType = 102
	Rune           CIMType = 103
	CIMObject      CIMType = 13
	Int8Array      CIMType = Int8 | CIMArray
	Uint8Array     CIMType = Uint8 | CIMArray
	Int16Array     CIMType = Int16 | CIMArray
	Uint16Array    CIMType = Uint16 | CIMArray
	Int32Array     CIMType = Int32 | CIMArray
	Uint32Array    CIMType = Uint32 | CIMArray
	Int64Array     CIMType = Int64 | CIMArray
	Uint64Array    CIMType = Uint64 | CIMArray
	Float32Array   CIMType = Float32 | CIMArray
	Float64Array   CIMType = Float64 | CIMArray
	BoolArray      CIMType = Bool | CIMArray
	StringArray    CIMType = String | CIMArray
	DateTimeArray  CIMType = DateTime | CIMArray
	RefArray       CIMType = Ref | CIMArray
	RuneArray      CIMType = Rune | CIMArray
	CIMObjectArray CIMType = CIMObject | CIMArray
)

func (o CIMType) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.String())
}

func (o CIMType) IsArray() bool {
	return o&CIMArray != 0
}

func (o CIMType) String() string {

	switch o {
	case Int8:
		return "int8"
	case Uint8:
		return "uint8"
	case Int16:
		return "int16"
	case Uint16:
		return "uint16"
	case Int32:
		return "int32"
	case Uint32:
		return "uint32"
	case Int64:
		return "int64"
	case Uint64:
		return "uint64"
	case Float32:
		return "float32"
	case Float64:
		return "float64"
	case Bool:
		return "bool"
	case String:
		return "string"
	case DateTime:
		return "datetime"
	case Ref:
		return "ref"
	case Rune:
		return "rune"
	case CIMObject:
		return "object"
	case Int8Array:
		return "[]int8"
	case Uint8Array:
		return "[]uint8"
	case Int16Array:
		return "[]int16"
	case Uint16Array:
		return "[]uint16"
	case Int32Array:
		return "[]int32"
	case Uint32Array:
		return "[]uint32"
	case Int64Array:
		return "[]int64"
	case Uint64Array:
		return "[]uint64"
	case Float32Array:
		return "[]float32"
	case Float64Array:
		return "[]float64"
	case BoolArray:
		return "[]bool"
	case StringArray:
		return "[]string"
	case DateTimeArray:
		return "[]datetime"
	case RefArray:
		return "[]ref"
	case RuneArray:
		return "[]rune"
	case CIMObjectArray:
		return "[]object"
	}

	return fmt.Sprintf("novalue(%d)", o)
}
