package wmio

import (
	"fmt"
	"strings"
)

type Value struct {
	Type  CIMType `json:"type"`
	Value any     `json:"value"`
}

func ValueType(value any) (CIMType, error) {

	var typ CIMType

	switch value.(type) {
	case uint8:
		typ = Uint8
	case []uint8:
		typ = Uint8Array
	case int8:
		typ = Int8
	case []int8:
		typ = Int8Array
	case uint16:
		typ = Uint16
	case []uint16:
		typ = Uint16Array
	case int16:
		typ = Int16
	case []int16:
		typ = Int16Array
	case bool:
		typ = Bool
	case []bool:
		typ = BoolArray
	case uint32:
		typ = Uint32
	case []uint32:
		typ = Uint32Array
	case int32:
		typ = Int32
	case []int32:
		typ = Int32Array
	case uint64:
		typ = Uint64
	case []uint64:
		typ = Uint64Array
	case int64:
		typ = Int64
	case []int64:
		typ = Int64Array
	case float32:
		typ = Float32
	case []float32:
		typ = Float32Array
	case float64:
		typ = Float64
	case []float64:
		typ = Float64Array
	case *Object:
		typ = CIMObject
	case []*Object:
		typ = CIMObjectArray
	case string:
		typ = String
	case []string:
		typ = StringArray
	case Value:
		// no-op
	default:
		return 0, fmt.Errorf("unknown type %T", value)
	}

	return typ, nil
}

func (o Value) EncodeValueSize() int {

	if o.Type.IsArray() /* refable */ {
		return 4
	}

	switch o.Type {
	case Int8, Uint8:
		return 1
	case Int16, Uint16, Rune, Bool:
		return 2
	case Int32, Uint32, Float32:
		return 4
	case Int64, Uint64, Float64:
		return 8
	case /* refable */ String, DateTime, Ref, CIMObject:
		return 4
	}

	return 0
}

func (o Value) String() string {
	if o.Value == nil {
		return "NULL"
	}
	switch o.Type {
	case Int8, Uint8, Int16, Uint16, Int32, Uint32, Int64, Uint64, Float32, Float64, Bool, Rune:
		return fmt.Sprintf("%v", o.Value)
	case String, DateTime, Ref:
		return fmt.Sprintf("%q", o.Value)
	case CIMObject:
		return "Not Implemented"
	default:
		ret := []string{}
		switch o.Type {
		case Int8Array:
			v, _ := o.Value.([]int8)
			for i := range v {
				ret = append(ret, fmt.Sprintf("%v", v[i]))
			}
		case Uint8Array:
			v, _ := o.Value.([]uint8)
			for i := range v {
				ret = append(ret, fmt.Sprintf("%v", v[i]))
			}
		case Int16Array:
			v, _ := o.Value.([]int16)
			for i := range v {
				ret = append(ret, fmt.Sprintf("%v", v[i]))
			}
		case Uint16Array:
			v, _ := o.Value.([]uint16)
			for i := range v {
				ret = append(ret, fmt.Sprintf("%v", v[i]))
			}
		case Int32Array:
			v, _ := o.Value.([]int32)
			for i := range v {
				ret = append(ret, fmt.Sprintf("%v", v[i]))
			}
		case Uint32Array:
			v, _ := o.Value.([]uint32)
			for i := range v {
				ret = append(ret, fmt.Sprintf("%v", v[i]))
			}
		case Int64Array:
			v, _ := o.Value.([]int64)
			for i := range v {
				ret = append(ret, fmt.Sprintf("%v", v[i]))
			}
		case Uint64Array:
			v, _ := o.Value.([]uint64)
			for i := range v {
				ret = append(ret, fmt.Sprintf("%v", v[i]))
			}
		case Float32Array:
			v, _ := o.Value.([]float32)
			for i := range v {
				ret = append(ret, fmt.Sprintf("%v", v[i]))
			}
		case Float64Array:
			v, _ := o.Value.([]float64)
			for i := range v {
				ret = append(ret, fmt.Sprintf("%v", v[i]))
			}
		case BoolArray:
			v, _ := o.Value.([]bool)
			for i := range v {
				ret = append(ret, fmt.Sprintf("%v", v[i]))
			}
		case RuneArray:
			v, _ := o.Value.([]uint16)
			for i := range v {
				ret = append(ret, fmt.Sprintf("%v", v[i]))
			}
		case StringArray, DateTimeArray, RefArray:
			v, _ := o.Value.([]string)
			for i := range v {
				ret = append(ret, fmt.Sprintf("%q", v[i]))
			}
		case CIMObjectArray:
			return "Not Implemented"
		}

		return strings.Join(ret, ", ")
	}
}

func (o *Value) Encode(r *Codec) error {
	r.Begin("value." + o.Type.String())
	switch o.Type {
	case Int8, Uint8, Int16, Uint16, Int32, Uint32, Int64, Uint64, Float32, Float64, Rune:
		r.WriteData(o.Value)
	case Bool:
		if v, _ := o.Value.(bool); v {
			r.WriteData((uint16)(0xFFFF))
		} else {
			r.WriteData((uint16)(0x0000))
		}
	case 0:
		// no-op.
	default:
		if o.Value == nil {
			return r.WriteRef(nil)
		}
		r.WriteRef(func(r *Codec) error {
			r.Begin("value.*")

			var sz int

			switch o.Type {
			case String, DateTime, Ref, CIMObject:
				// no-op. (write_data).
			case Int8Array:
				v, _ := o.Value.([]int8)
				sz = len(v)
			case Uint8Array:
				v, _ := o.Value.([]uint8)
				sz = len(v)
			case Int16Array, RuneArray:
				v, _ := o.Value.([]int16)
				sz = len(v)
			case Uint16Array:
				v, _ := o.Value.([]uint16)
				sz = len(v)
			case BoolArray:
				v, _ := o.Value.([]bool)
				sz = len(v)
			case Int32Array:
				v, _ := o.Value.([]int32)
				sz = len(v)
			case Uint32Array:
				v, _ := o.Value.([]uint32)
				sz = len(v)
			case Int64Array:
				v, _ := o.Value.([]int64)
				sz = len(v)
			case Uint64Array:
				v, _ := o.Value.([]uint64)
				sz = len(v)
			case Float32Array:
				v, _ := o.Value.([]float32)
				sz = len(v)
			case Float64Array:
				v, _ := o.Value.([]float64)
				sz = len(v)
			case StringArray, DateTimeArray, RefArray:
				v, _ := o.Value.([]string)
				sz = len(v)
			case CIMObjectArray:
				v, _ := o.Value.([]*Object)
				sz = len(v)
			}

			if o.Type.IsArray() {
				r.WriteData((uint32)(sz))
			}

			switch o.Type {
			case BoolArray:
				v, _ := o.Value.([]bool)
				for i := range v {
					if v[i] {
						r.WriteData((uint16)(0xFFFF))
					} else {
						r.WriteData((uint16)(0x0000))
					}
				}
			case StringArray, DateTimeArray, RefArray:
				v, _ := o.Value.([]string)
				r.WriteDataOnHeap(len(v)*4, func(r *Codec) error {
					r.Begin("string_array")
					for i := range v {
						r.WriteRef(v[i])
					}
					return r.Done()
				})
			case CIMObjectArray:
				v, _ := o.Value.([]*Object)
				r.WriteDataOnHeap(len(v)*4, func(r *Codec) error {
					r.Begin("object_array")
					for i := range v {
						r.WriteRef(v[i])
					}
					return r.Done()
				})
			default:
				r.WriteData(o.Value)
			}

			return r.Done()
		}, o.Type.String())
	}
	return r.Done()
}

func (o *Value) Decode(r *Codec) error {
	r.Begin("value." + o.Type.String())
	switch o.Type {
	case Int8:
		v := int8(0)
		r.ReadData(&v)
		o.Value = v
	case Uint8:
		v := uint8(0)
		r.ReadData(&v)
		o.Value = v
	case Int16:
		v := int16(0)
		r.ReadData(&v)
		o.Value = v
	case Uint16, Rune:
		v := uint16(0)
		r.ReadData(&v)
		o.Value = v
	case Int32:
		v := int32(0)
		r.ReadData(&v)
		o.Value = v
	case Uint32:
		v := uint32(0)
		r.ReadData(&v)
		o.Value = v
	case Int64:
		v := int64(0)
		r.ReadData(&v)
		o.Value = v
	case Uint64:
		v := uint64(0)
		r.ReadData(&v)
		o.Value = v
	case Float32:
		v := float32(0)
		r.ReadData(&v)
		o.Value = v
	case Float64:
		v := float64(0)
		r.ReadData(&v)
		o.Value = v
	case Bool:
		v := uint16(0)
		r.ReadData(&v)
		o.Value = v == 0xFFFF
	case 0:
		// no-op.
	default:
		r.ReadRef(func(r *Codec) error {
			r.Begin("*")
			sz := uint32(0)
			if o.Type.IsArray() {
				r.ReadData(&sz)
			}
			switch o.Type {
			case CIMObject:
				v := Object{}
				r.ReadData(&v)
				o.Value = &v
			case String, DateTime, Ref:
				v := ""
				r.ReadData(&v)
				o.Value = v
			case Int8Array:
				aV := make([]int8, sz)
				r.ReadData(&aV)
				o.Value = aV
			case Uint8Array:
				aV := make([]uint8, sz)
				r.ReadData(&aV)
				o.Value = aV
			case Int16Array:
				aV := make([]int16, sz)
				r.ReadData(&aV)
				o.Value = aV
			case Uint16Array, RuneArray:
				aV := make([]uint16, sz)
				r.ReadData(&aV)
				o.Value = aV
			case Int32Array:
				aV := make([]int32, sz)
				r.ReadData(&aV)
				o.Value = aV
			case Uint32Array:
				aV := make([]uint32, sz)
				r.ReadData(&aV)
				o.Value = aV
			case Int64Array:
				aV := make([]int64, sz)
				r.ReadData(&aV)
				o.Value = aV
			case Uint64Array:
				aV := make([]uint64, sz)
				r.ReadData(&aV)
				o.Value = aV
			case Float32Array:
				aV := make([]float32, sz)
				r.ReadData(&aV)
				o.Value = aV
			case Float64Array:
				aV := make([]float64, sz)
				r.ReadData(&aV)
				o.Value = aV
			case BoolArray:
				aV := make([]bool, sz)
				aaV := make([]uint16, sz)
				r.ReadData(&aaV)
				for i := range aaV {
					aV[i] = aaV[i] == 0xFFFF
				}
				o.Value = aV
			case StringArray, DateTimeArray, RefArray:
				aV := make([]string, sz)
				for i := range aV {
					r.ReadRef(&aV[i])
				}
				o.Value = aV
			case CIMObjectArray:
				aV := make([]*Object, sz)
				for i := range aV {
					v := Object{}
					r.ReadRef(&v)
					aV[i] = &v
				}
				o.Value = aV
			default:
				return fmt.Errorf("unknown type %s", o.Type)
			}
			return r.Done()
		}, "ref")
	}
	return r.Done()
}
