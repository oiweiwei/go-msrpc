package math

type FloatFormat interface {
	Float32bits(float32) uint32
	Float32frombits(uint32) float32
	Float64bits(float64) uint64
	Float64frombits(uint64) float64
}

type floatFormat struct {
	float32bits     func(float32) uint32
	float32frombits func(uint32) float32
	float64bits     func(float64) uint64
	float64frombits func(uint64) float64
}

func (ff floatFormat) Float32bits(f float32) uint32 {
	return ff.float32bits(f)
}

func (ff floatFormat) Float32frombits(b uint32) float32 {
	return ff.float32frombits(b)
}

func (ff floatFormat) Float64bits(f float64) uint64 {
	return ff.float64bits(f)
}

func (ff floatFormat) Float64frombits(b uint64) float64 {
	return ff.float64frombits(b)
}
