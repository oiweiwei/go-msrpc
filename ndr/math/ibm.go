package math

import (
	"math/bits"
)

/*
	IBM Hexadecimal Single-precision Floating-point Format

	+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
	| 31 | 30 | 29 | 28 | 27 | 26 | 25 | 24 | 23 | 22 | 21 | 20 | 19 | 18 | 17 | 16 |
	+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
	| s  |                           e      |                 f1                    |
	+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
	| 15 | 14 | 13 | 12 | 11 | 10 |  9 |  8 |  7 |  6 |  5 |  4 |  3 |  2 |  1 |  0 |
	+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
	|                                         f2                                    |
	+-------------------------------------------------------------------------------+

*/

// IBMHexfloat32frombits ...
func IBMHexfloat32frombits(b uint32) float32 {

	var (
		exp     int
		s, e, m uint32
	)

	var (
		lz int = bits.LeadingZeros32(b & 0x00ffffff)
	)

	s = b & IEEE32Sign

	exp = int((b & 0x7f000000) >> 24)
	exp -= 64
	exp *= 4
	exp -= 1
	exp -= lz - 8

	e = ToIEEE32Exponent(exp)

	m = b & 0x00ffffff
	m ^= 1 << (31 - lz)
	m <<= (lz - 8)
	m &= IEEE32Mantissa

	return Float32frombits(s | e | m)

}

// IBMHexfloat32bits ...
func IBMHexfloat32bits(v float32) uint32 {

	var (
		b uint32 = Float32bits(v)

		exp     int
		s, e, m uint32
	)

	s = b & 0x80000000

	exp = IEEE32Exponent(b)
	exp += 1

	m = b & 0x007fffff
	m += 0x00800000
	m >>= uint((4 - exp%4) % 4)

	exp += (4 - exp%4) % 4
	exp /= 4
	exp += 64

	e = uint32(exp) << 24

	return (s | e | m)
}

/*
	IBM Hexadecimal Double-precision Floating-point Format

	+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
	| 63 | 62 | 61 | 60 | 59 | 58 | 57 | 56 | 55 | 54 | 53 | 52 | 51 | 50 | 49 | 48 |
	+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
	| s  |                 e                |                 f1                    |
	+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
	| 47 | 46 | 45 | 44 | 43 | 42 | 41 | 40 | 39 | 38 | 37 | 36 | 35 | 34 | 33 | 32 |
	+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
	|                                         f2                                    |
	+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
	| 31 | 30 | 29 | 28 | 27 | 26 | 25 | 24 | 23 | 22 | 21 | 20 | 19 | 18 | 17 | 16 |
	+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
	|                                         f3                                    |
	+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
	| 15 | 14 | 13 | 12 | 11 | 10 |  9 |  8 |  7 |  6 |  5 |  4 |  3 |  2 |  1 |  0 |
	+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
	|                                         f4                                    |
	+-------------------------------------------------------------------------------+

*/

// IBMHexfloat64frombits ...
func IBMHexfloat64frombits(b uint64) float64 {

	var (
		exp     int
		s, e, m uint64
	)

	var (
		lz int = bits.LeadingZeros64(b & 0x00ffffffffffffff)
	)

	s = b & IEEE64Sign

	exp = int((b & 0x7f00000000000000) >> 56)
	exp -= 64
	exp *= 4
	exp -= 1
	exp -= lz - 8

	e = ToIEEE64Exponent(exp)

	m = b & 0x00ffffffffffffff
	m ^= 1 << (63 - lz)
	m <<= uint(lz) - 8
	m = (m >> 3) & IEEE64Mantissa

	return Float64frombits(s | e | m)
}

// IBMHexfloat64bits ...
func IBMHexfloat64bits(v float64) uint64 {

	var (
		b uint64 = Float64bits(v)

		exp     int
		s, e, m uint64
	)

	s = b & 0x8000000000000000

	exp = IEEE64Exponent(b)
	exp += 1

	m = (b & IEEE64Mantissa) << 3
	m += 0x0080000000000000
	m >>= uint((4 - (exp)%4) % 4)
	m &= 0x00ffffffffffffff

	exp += (4 - (exp)%4) % 4
	exp /= 4
	exp += 64

	e = uint64(exp) << 56

	return (s | e | m)
}

var IBMHex = floatFormat{
	float32frombits: IBMHexfloat32frombits,
	float32bits:     IBMHexfloat32bits,
	float64frombits: IBMHexfloat64frombits,
	float64bits:     IBMHexfloat64bits,
}
