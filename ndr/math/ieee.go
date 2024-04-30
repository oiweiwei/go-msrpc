package math

import (
	"math"
)

var (
	Float32bits     = math.Float32bits
	Float32frombits = math.Float32frombits
	Float64bits     = math.Float64bits
	Float64frombits = math.Float64frombits
)

const (
	IEEE32Excess = 127
	IEEE64Excess = 1023

	IEEE32Mantissa = 0x007fffff
	IEEE64Mantissa = 0x000fffffffffffff

	IEEE32Sign = 0x80000000
	IEEE64Sign = 0x8000000000000000

	IEEE32SignPos = 31
	IEEE64SignPos = 63
)

func ToIEEE32Exponent(e int) uint32 {
	return (uint32(e+IEEE32Excess) << 23) & 0x7f800000
}

func IEEE32Exponent(u uint32) int {
	return int((u&0x7f800000)>>23) - IEEE32Excess
}

func ToIEEE64Exponent(e int) uint64 {
	return (uint64(e+IEEE64Excess) << 52) & 0x7ff0000000000000
}

func IEEE64Exponent(u uint64) int {
	return int((u&0x7ff0000000000000)>>52) - IEEE64Excess
}

var IEEE = floatFormat{
	float32frombits: Float32frombits,
	float32bits:     Float32bits,
	float64frombits: Float64frombits,
	float64bits:     Float64bits,
}
