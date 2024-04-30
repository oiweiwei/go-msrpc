package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCrayFloat32(t *testing.T) {

	var (
		ret float32
	)

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
		ret = CrayFloat32frombits(CrayFloat32bits(testCase))
		t.Logf("checking %v... ~ %v [abs: %v, rel: %v]", testCase, ret, ret-testCase, (ret-testCase)/testCase)
		t.Logf("0x%08x", CrayFloat32bits(testCase))

		abs := ret-testCase < cDelta32 && ret-testCase > -cDelta32

		rel := (ret-testCase)/testCase < cDelta32 && (ret-testCase)/testCase > -cDelta32

		assert.True(t, abs || rel)

	}
}

func TestCrayFloat64(t *testing.T) {

	var (
		ret float64
	)

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
		ret = CrayFloat64frombits(CrayFloat64bits(testCase))
		t.Logf("checking %v... ~ %v [abs: %v, rel: %v]", testCase, ret, ret-testCase, (ret-testCase)/testCase)
		t.Logf("0x%08x", CrayFloat64bits(testCase))

		abs := ret-testCase < cDelta64 && ret-testCase > -cDelta64

		rel := (ret-testCase)/testCase < cDelta64 && (ret-testCase)/testCase > -cDelta64

		assert.True(t, abs || rel)

	}
}
