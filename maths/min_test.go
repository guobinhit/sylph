package maths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMin(t *testing.T) {
	is := assert.New(t)

	is.Equal(int(0), MinInts())
	is.Equal(int(0), MinInts([]int{}...))
	is.Equal(int(1), MinInts(1))
	is.Equal(int(1), MinInts(1, 2, 3))
	is.Equal(int(1), MinInts(3, 2, 1))
	is.Equal(int(-3), MinInts(-3, 1, 2))

	is.Equal(int8(0), MinInt8s())
	is.Equal(int8(0), MinInt8s([]int8{}...))
	is.Equal(int8(1), MinInt8s(1))
	is.Equal(int8(1), MinInt8s(1, 2, 3))
	is.Equal(int8(1), MinInt8s(3, 2, 1))
	is.Equal(int8(-3), MinInt8s(-3, 1, 2))

	is.Equal(int16(0), MinInt16s())
	is.Equal(int16(0), MinInt16s([]int16{}...))
	is.Equal(int16(1), MinInt16s(1))
	is.Equal(int16(1), MinInt16s(1, 2, 3))
	is.Equal(int16(1), MinInt16s(3, 2, 1))
	is.Equal(int16(-3), MinInt16s(-3, 1, 2))

	is.Equal(int32(0), MinInt32s())
	is.Equal(int32(0), MinInt32s([]int32{}...))
	is.Equal(int32(1), MinInt32s(1))
	is.Equal(int32(1), MinInt32s(1, 2, 3))
	is.Equal(int32(1), MinInt32s(3, 2, 1))
	is.Equal(int32(-3), MinInt32s(-3, 1, 2))

	is.Equal(int64(0), MinInt64s())
	is.Equal(int64(0), MinInt64s([]int64{}...))
	is.Equal(int64(1), MinInt64s(1))
	is.Equal(int64(1), MinInt64s(1, 2, 3))
	is.Equal(int64(1), MinInt64s(3, 2, 1))
	is.Equal(int64(-3), MinInt64s(-3, 1, 2))

	is.Equal(float64(0), MinFloat64s())
	is.Equal(float64(0), MinFloat64s([]float64{}...))
	is.Equal(1.0, MinFloat64s(1.0))
	is.Equal(1.0, MinFloat64s(1.0, 2.0, 3.0))
	is.Equal(1.0, MinFloat64s(3.0, 2.0, 1.0))
	is.Equal(-3.0, MinFloat64s(-3.0, 1.0, 2.0))
}
