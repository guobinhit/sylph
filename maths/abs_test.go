package maths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbs(t *testing.T) {
	is := assert.New(t)

	is.PanicsWithValue("int8 overflow: abs(-128)", func() { AbsInt8(-128) })
	is.Equal(int8(0), AbsInt8(0))
	is.Equal(int8(127), AbsInt8(int8(-127)))
	is.Equal(int8(127), AbsInt8(int8(127)))

	is.PanicsWithValue("int16 overflow: abs(-32768)", func() { AbsInt16(-32768) })
	is.Equal(int16(0), AbsInt16(0))
	is.Equal(int16(32767), AbsInt16(int16(-32767)))
	is.Equal(int16(32767), AbsInt16(int16(32767)))

	is.PanicsWithValue("int32 overflow: abs(-2147483648)", func() { AbsInt32(-2147483648) })
	is.Equal(int32(0), AbsInt32(0))
	is.Equal(int32(2147483647), AbsInt32(int32(-2147483647)))
	is.Equal(int32(2147483647), AbsInt32(int32(2147483647)))

	is.PanicsWithValue("int64 overflow: abs(-9223372036854775808)", func() { AbsInt64(-9223372036854775808) })
	is.Equal(int64(0), AbsInt64(0))
	is.Equal(int64(9223372036854775807), AbsInt64(int64(-9223372036854775807)))
	is.Equal(int64(9223372036854775807), AbsInt64(int64(9223372036854775807)))

	is.Equal(float32(0), AbsFloat32(0))
	is.Equal(float32(9223372036854775807), AbsFloat32(float32(-9223372036854775807)))
	is.Equal(float32(9223372036854775807), AbsFloat32(float32(9223372036854775807)))

	is.Equal(float64(0), AbsFloat64(0))
	is.Equal(float64(9223372036854775807), AbsFloat64(float64(-9223372036854775807)))
	is.Equal(float64(9223372036854775807), AbsFloat64(float64(9223372036854775807)))
}
