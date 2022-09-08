package maths

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestAbs(t *testing.T) {
	is := assert.New(t)

	is.PanicsWithValue("int8 overflow: abs(-128)", func() { AbsInt8(-128) })
	is.Equal(int8(0), AbsInt8(0))
	is.Equal(int8(127), AbsInt8(math.MinInt8+1))
	is.Equal(int8(127), AbsInt8(math.MaxInt8))

	is.PanicsWithValue("int16 overflow: abs(-32768)", func() { AbsInt16(-32768) })
	is.Equal(int16(0), AbsInt16(0))
	is.Equal(int16(32767), AbsInt16(math.MinInt16+1))
	is.Equal(int16(32767), AbsInt16(math.MaxInt16))

	is.PanicsWithValue("int32 overflow: abs(-2147483648)", func() { AbsInt32(-2147483648) })
	is.Equal(int32(0), AbsInt32(0))
	is.Equal(int32(2147483647), AbsInt32(math.MinInt32+1))
	is.Equal(int32(2147483647), AbsInt32(math.MaxInt32))

	is.PanicsWithValue("int64 overflow: abs(-9223372036854775808)", func() { AbsInt64(-9223372036854775808) })
	is.Equal(int64(0), AbsInt64(0))
	is.Equal(int64(9223372036854775807), AbsInt64(math.MinInt64+1))
	is.Equal(int64(9223372036854775807), AbsInt64(math.MaxInt64))

	is.Equal(float32(0), AbsFloat32(0))
	is.Equal(float32(math.MaxFloat32), AbsFloat32(math.MaxFloat32+1))
	is.Equal(float32(math.MaxFloat32), AbsFloat32(math.MaxFloat32))

	is.Equal(float64(0), AbsFloat64(0))
	is.Equal(math.MaxFloat64, AbsFloat64(math.MaxFloat64+1))
	is.Equal(math.MaxFloat64, AbsFloat64(math.MaxFloat64))
}
