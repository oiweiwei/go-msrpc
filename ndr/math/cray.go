package math

// Cray machine architecture defines only a double-precision floating-point format.
// However, because Cray machines may be required to handle single-precision
// floating-point values (for instance, if single-precision values are specified
// in an interface definition), NDR defines a single-precision floating-point format
// for Cray machines; this format is identical to IEEE big-endian single-precision
// format.

/*

	Cray NDR Single-precision Floating-point Format (IEEE Big-Endian Single-Precision)

	+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
	| 31 | 30 | 29 | 28 | 27 | 26 | 25 | 24 | 23 | 22 | 21 | 20 | 19 | 18 | 17 | 16 |
	+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
	|                     f3                |                 f2                    |
	+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
	| 15 | 14 | 13 | 12 | 11 | 10 |  9 |  8 |  7 |  6 |  5 |  4 |  3 |  2 |  1 |  0 |
	+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
	| e2 |                f1                |  s |                 e1               |
	+-------------------------------------------------------------------------------+

*/

// CrayFloat32frombits ...
func CrayFloat32frombits(b uint32) float32 {

	var (
		r uint32
	)

	r |= (b & 0x000000ff) << 24
	r |= (b & 0x0000ff00) << 8
	r |= (b & 0x00ff0000) >> 8
	r |= (b & 0xff000000) >> 24

	return Float32frombits(r)
}

// CrayFloat32bits ...
func CrayFloat32bits(v float32) uint32 {

	var (
		b uint32 = Float32bits(v)

		r uint32
	)

	r |= (b & 0x000000ff) << 24
	r |= (b & 0x0000ff00) << 8
	r |= (b & 0x00ff0000) >> 8
	r |= (b & 0xff000000) >> 24

	return r
}

/*

	Cray NDR Double-precision Floating-point Format

	+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
	| 63 | 62 | 61 | 60 | 59 | 58 | 57 | 56 | 55 | 54 | 53 | 52 | 51 | 50 | 49 | 48 |
	+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
	|                f6                     |                f5                     |
	+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
	| 47 | 46 | 45 | 44 | 43 | 42 | 41 | 40 | 39 | 38 | 37 | 36 | 35 | 34 | 33 | 32 |
	+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
	|                f4                     |                f3                     |
	+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
	| 31 | 30 | 29 | 28 | 27 | 26 | 25 | 24 | 23 | 22 | 21 | 20 | 19 | 18 | 17 | 16 |
	+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
	|                f2                     |                f1                     |
	+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
	| 15 | 14 | 13 | 12 | 11 | 10 |  9 |  8 |  7 |  6 |  5 |  4 |  3 |  2 |  1 |  0 |
	+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
	|                e2                     |  s |           e1                     |
	+---------------------------------------+----+----------------------------------+

*/

func CrayFloat64bits(v float64) uint64 {

	var (
		b uint64 = Float64bits(v)

		exp     int
		s, e, m uint64
	)

	s |= (b & IEEE64Sign) >> IEEE64SignPos
	s <<= 7

	exp = IEEE64Exponent(b)
	exp -= 16384

	e |= (uint64(exp) & 0x00000000000000ff) << 8
	e |= (uint64(exp) & 0x0000000000007f00) >> 8

	m |= (b & 0x000ff00000000000) >> 28
	m |= (b & 0x00000ff000000000) >> 12
	m |= (b & 0x0000000ff0000000) << 4
	m |= (b & 0x000000000ff00000) << 20
	m |= (b & 0x00000000000ff000) << 36

	return (s | e | m)
}

func CrayFloat64frombits(b uint64) float64 {

	var (
		exp     int
		s, e, m uint64
	)

	s = (b & 0x0000000000000080) >> 7
	s <<= IEEE64SignPos

	e |= (b & 0x000000000000007f) << 8
	e |= (b & 0x000000000000ff00) >> 8

	exp = int(e)
	exp += 16384

	e = ToIEEE64Exponent(exp)

	m |= (b & 0x0000000000ff0000) << 28
	m |= (b & 0x00000000ff000000) << 12
	m |= (b & 0x000000ff00000000) >> 4
	m |= (b & 0x0000ff0000000000) >> 20
	m |= (b & 0x00ff000000000000) >> 36

	m &= IEEE64Mantissa

	return Float64frombits(s | e | m)

}

var Cray = floatFormat{
	float32frombits: CrayFloat32frombits,
	float32bits:     CrayFloat32bits,
	float64frombits: CrayFloat64frombits,
	float64bits:     CrayFloat64bits,
}
