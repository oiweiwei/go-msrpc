package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVaxFFloat32(t *testing.T) {

	var (
		ret float32
	)

	// source for the test data : https://pubs.usgs.gov/of/2005/1424/

	for _, testCase := range []struct {
		V   float32
		Ret uint32
	}{
		{1.000000, 0x00004080},
		{-1.000000, 0x0000C080},
		{3.500000, 0x00004160},
		{-3.500000, 0x0000C160},
		{3.141590, 0x0FD04149},
		{-3.141590, 0x0FD0C149},
		{9.9999999e+36, 0xBDC27DF0},
		{-9.9999999e+36, 0xBDC2FDF0},
		{9.9999999e-38, 0x1CEA0308},
		{-9.9999999e-38, 0x1CEA8308},
		{1.234568, 0x0653409E},
		{-1.234568, 0x0653C09E},
	} {

		t.Logf("checking value for %v...", testCase.V)
		assert.Equal(t, testCase.Ret, VaxFfloat32bits(testCase.V))
	}

	for _, testCase := range []float32{
		-118.625,
		0.1526,
		25.34,
		25.33,
		234566.23,
		-0.1752,
		0.001525,
		0.00023,
		0.000215,
		0.000015,
		123.45,
		1e3,
		1e-45,
	} {

		ret = VaxFfloat32frombits(VaxFfloat32bits(testCase))
		t.Logf("checking %v... ~ %v [abs: %v, rel: %v]", testCase, ret, ret-testCase, (ret-testCase)/testCase)
		t.Logf("vax: 0x%08x, float: 0x%08x", VaxFfloat32bits(testCase), Float32bits(testCase))

		abs := ret-testCase < cDelta32 && ret-testCase > -cDelta32

		rel := (ret-testCase)/testCase < cDelta32 && (ret-testCase)/testCase > -cDelta32

		assert.True(t, abs || rel)

	}
}

func TestVaxGfloat64(t *testing.T) {

	var (
		ret float64
	)

	// source for the test data : https://pubs.usgs.gov/of/2005/1424/

	for _, testCase := range []struct {
		V   float64
		Ret uint64
	}{

		{1.000000000000000, 0x0000000000004010},
		{-1.000000000000000, 0x000000000000C010},
		{3.500000000000000, 0x000000000000402C},
		{-3.500000000000000, 0x000000000000C02C},
		{3.141592653589793, 0x2D18544421FB4029},
		{-3.141592653589793, 0x2D18544421FBC029},
		{1.0000000000000000e+37, 0x691B435717B847BE},
		{-1.0000000000000000e+37, 0x691B435717B8C7BE},
		{9.9999999999999999e-38, 0x8B8F428A039D3861},
		{-9.9999999999999999e-38, 0x8B8F428A039DB861},
		{1.234567890123450, 0x59DD428CC0CA4013},
		{-1.234567890123450, 0x59DD428CC0CAC013},
	} {
		t.Logf("checking value for %v..., float: %v", testCase.V, VaxGfloat64frombits(VaxGfloat64bits(testCase.V)))
		assert.Equal(t, testCase.Ret, VaxGfloat64bits(testCase.V))
	}

	for _, testCase := range []float64{
		-118.625,
		0.1526,
		25.34,
		25.33,
		234566.23,
		-0.1752,
		0.001525,
		0.00023,
		0.000215,
		0.000015,
		123.45,
		1e3,
		1e-45,
	} {

		ret = VaxGfloat64frombits(VaxGfloat64bits(testCase))
		t.Logf("checking %v... ~ %v [abs: %v, rel: %v]", testCase, ret, ret-testCase, (ret-testCase)/testCase)
		t.Logf("vax: 0x%08x, float: 0x%08x", VaxGfloat64bits(testCase), Float64bits(testCase))

		abs := ret-testCase < cDelta64 && ret-testCase > -cDelta64

		rel := (ret-testCase)/testCase < cDelta64 && (ret-testCase)/testCase > -cDelta64

		assert.True(t, abs || rel)

	}
}
