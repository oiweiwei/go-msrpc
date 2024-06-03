package wmio

import "strconv"

func JSONValueToType(value any, typ CIMType) (any, bool) {
	switch value := value.(type) {
	case float64:
		switch typ {
		case Uint8:
			return uint8(value), true
		case Int8:
			return int8(value), true
		case Uint16:
			return uint16(value), true
		case Int16:
			return int16(value), true
		case Uint32:
			return uint32(value), true
		case Int32:
			return int32(value), true
		case Uint64:
			return uint64(value), true
		case Int64:
			return int64(value), true
		case Float32:
			return float32(value), true
		case Float64:
			return value, true
		}
	case []float64:
		switch typ {
		case Uint8Array:
			ret := make([]uint8, len(value))
			for i := range value {
				ret[i] = uint8(value[i])
			}
			return ret, true
		case Int8Array:
			ret := make([]int8, len(value))
			for i := range value {
				ret[i] = int8(value[i])
			}
			return ret, true
		case Uint16Array:
			ret := make([]uint16, len(value))
			for i := range value {
				ret[i] = uint16(value[i])
			}
			return ret, true
		case Int16Array:
			ret := make([]int16, len(value))
			for i := range value {
				ret[i] = int16(value[i])
			}
			return ret, true
		case Uint32Array:
			ret := make([]uint32, len(value))
			for i := range value {
				ret[i] = uint32(value[i])
			}
			return ret, true
		case Int32Array:
			ret := make([]int32, len(value))
			for i := range value {
				ret[i] = int32(value[i])
			}
			return ret, true
		case Uint64Array:
			ret := make([]uint64, len(value))
			for i := range value {
				ret[i] = uint64(value[i])
			}
			return ret, true
		case Int64Array:
			ret := make([]int64, len(value))
			for i := range value {
				ret[i] = int64(value[i])
			}
			return ret, true
		case Float32Array:
			ret := make([]float32, len(value))
			for i := range value {
				ret[i] = float32(value[i])
			}
			return value, true
		case Float64Array:
			return value, true
		}
	case string:
		switch typ {
		case Uint8:
			ret, _ := strconv.ParseUint(value, 10, 8)
			return uint8(ret), true
		case Int8:
			ret, _ := strconv.ParseInt(value, 10, 8)
			return uint8(ret), true
		case Uint16:
			ret, _ := strconv.ParseUint(value, 10, 16)
			return uint16(ret), true
		case Int16:
			ret, _ := strconv.ParseInt(value, 10, 16)
			return uint16(ret), true
		case Uint32:
			ret, _ := strconv.ParseUint(value, 10, 32)
			return uint32(ret), true
		case Int32:
			ret, _ := strconv.ParseInt(value, 10, 32)
			return uint32(ret), true
		case Uint64:
			ret, _ := strconv.ParseUint(value, 10, 64)
			return uint64(ret), true
		case Int64:
			ret, _ := strconv.ParseInt(value, 10, 64)
			return uint64(ret), true
		case Float32:
			ret, _ := strconv.ParseFloat(value, 32)
			return float32(ret), true
		case Float64:
			ret, _ := strconv.ParseFloat(value, 64)
			return float64(ret), true
		}
	case []string:
		switch typ {
		case Uint8Array:
			ret := make([]uint8, len(value))
			for i := range value {
				rret, _ := strconv.ParseUint(value[i], 10, 8)
				ret[i] = uint8(rret)
			}
			return ret, true
		case Int8Array:
			ret := make([]int8, len(value))
			for i := range value {
				rret, _ := strconv.ParseInt(value[i], 10, 8)
				ret[i] = int8(rret)
			}
			return ret, true
		case Uint16Array:
			ret := make([]uint16, len(value))
			for i := range value {
				rret, _ := strconv.ParseUint(value[i], 10, 16)
				ret[i] = uint16(rret)
			}
			return ret, true
		case Int16Array:
			ret := make([]int16, len(value))
			for i := range value {
				rret, _ := strconv.ParseInt(value[i], 10, 16)
				ret[i] = int16(rret)
			}
			return ret, true
		case Uint32Array:
			ret := make([]uint32, len(value))
			for i := range value {
				rret, _ := strconv.ParseUint(value[i], 10, 32)
				ret[i] = uint32(rret)
			}
			return ret, true
		case Int32Array:
			ret := make([]int32, len(value))
			for i := range value {
				rret, _ := strconv.ParseInt(value[i], 10, 32)
				ret[i] = int32(rret)
			}
			return ret, true
		case Uint64Array:
			ret := make([]uint64, len(value))
			for i := range value {
				rret, _ := strconv.ParseUint(value[i], 10, 64)
				ret[i] = uint64(rret)
			}
			return ret, true
		case Int64Array:
			ret := make([]int64, len(value))
			for i := range value {
				rret, _ := strconv.ParseInt(value[i], 10, 64)
				ret[i] = int64(rret)
			}
			return ret, true
		}
	}

	return value, false
}
