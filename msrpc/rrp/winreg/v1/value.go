package winreg

import (
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/oiweiwei/go-msrpc/ndr"
	"github.com/oiweiwei/go-msrpc/text/encoding/utf16le"
)

type None struct{}

func isUTF16(value []byte) bool {
	if len(value)%2 != 0 {
		return false // not even length
	}
	if len(value) < 2 {
		return false // too short for UTF-16
	}
	if value[len(value)-2] != 0 || value[len(value)-1] != 0 {
		return false // not null-terminated
	}
	return true
}

func decodeString(value []byte) (string, error) {
	if isUTF16(value) {
		return utf16le.Decode(value)
	}
	return string(value), nil
}

func EncodeValue(value any, valueType uint32) ([]byte, error) {

	var (
		ret []byte
		err error
	)

	switch valueType {
	case RegDword:
		ret = make([]byte, 4)
		switch v := value.(type) {
		case int32:
			binary.LittleEndian.PutUint32(ret, uint32(v))
		case uint32:
			binary.LittleEndian.PutUint32(ret, v)
		case int:
			binary.LittleEndian.PutUint32(ret, uint32(v))
		default:
			return nil, fmt.Errorf("invalid value type for dword: %T", value)
		}
	case RegDwordBigEndian:
		ret = make([]byte, 4)
		switch v := value.(type) {
		case int32:
			binary.BigEndian.PutUint32(ret, uint32(v))
		case uint32:
			binary.BigEndian.PutUint32(ret, v)
		case int:
			binary.BigEndian.PutUint32(ret, uint32(v))
		default:
			return nil, fmt.Errorf("invalid value type for dword big endian: %T", value)
		}
	case RegQword:
		ret = make([]byte, 8)
		switch v := value.(type) {
		case int64:
			binary.LittleEndian.PutUint64(ret, uint64(v))
		case uint64:
			binary.LittleEndian.PutUint64(ret, v)
		case int:
			binary.LittleEndian.PutUint64(ret, uint64(v))
		default:
			return nil, fmt.Errorf("invalid value type for qword: %T", value)
		}
	case RegString, RegExpandString:
		s, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("invalid value type for string: %T", value)
		}

		ret, err = utf16le.Encode(s + ndr.ZeroString)
		if err != nil {
			return nil, fmt.Errorf("string value data encode error: %w", err)
		}
	case RegMultistring:
		s, ok := value.([]string)
		if !ok {
			return nil, fmt.Errorf("invalid value type for multistring: %T", value)
		}
		ret, err = utf16le.Encode(strings.Join(s, ndr.ZeroString) + ndr.ZeroString) // ensure the last zero terminator
		if err != nil {
			return nil, fmt.Errorf("multistring value data encode error: %w", err)
		}
	case RegNone:
	case RegBinary:
		switch v := value.(type) {
		case []byte:
			ret = v
		case string:
			ret = []byte(v)
		default:
			return nil, fmt.Errorf("invalid value type for binary: %T", value)
		}
	}

	return ret, nil
}

func DecodeValue(valueType uint32, valueData []byte) (any, error) {
	switch valueType {
	case RegDword:
		if len(valueData) < 4 {
			return nil, fmt.Errorf("dword value data too short: %d", len(valueData))
		}
		return binary.LittleEndian.Uint32(valueData), nil
	case RegDwordBigEndian:
		if len(valueData) < 4 {
			return nil, fmt.Errorf("dword big endian value data too short: %d", len(valueData))
		}
		return binary.BigEndian.Uint32(valueData), nil
	case RegQword:
		if len(valueData) < 8 {
			return nil, fmt.Errorf("qword value data too short: %d", len(valueData))
		}
		return binary.LittleEndian.Uint64(valueData), nil
	case RegString, RegExpandString:
		s, err := decodeString(valueData)
		if err != nil {
			return nil, fmt.Errorf("string value data decode error: %w", err)
		}
		return strings.TrimRight(s, ndr.ZeroString), nil
	case RegMultistring:
		s, err := decodeString(valueData)
		if err != nil {
			return nil, fmt.Errorf("multistring value data decode error: %w", err)
		}
		if s = strings.TrimRight(s, ndr.ZeroString); len(s) == 0 {
			return ([]string)(nil), nil // empty multistring
		}
		return strings.Split(s, ndr.ZeroString), nil
	case RegNone:
		return None{}, nil // no data
	default:
		return valueData, nil // raw data
	}
}
