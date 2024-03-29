package maths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax(t *testing.T) {
	is := assert.New(t)

	is.Equal(int(0), MaxInts())
	is.Equal(int(0), MaxInts([]int{}...))
	is.Equal(int(1), MaxInts(1))
	is.Equal(int(3), MaxInts(1, 2, 3))
	is.Equal(int(3), MaxInts(3, 2, 1))
	is.Equal(int(2), MaxInts(-3, 1, 2))

	is.Equal(int8(0), MaxInt8s())
	is.Equal(int8(0), MaxInt8s([]int8{}...))
	is.Equal(int8(1), MaxInt8s(1))
	is.Equal(int8(3), MaxInt8s(1, 2, 3))
	is.Equal(int8(3), MaxInt8s(3, 2, 1))
	is.Equal(int8(2), MaxInt8s(-3, 1, 2))

	is.Equal(int16(0), MaxInt16s())
	is.Equal(int16(0), MaxInt16s([]int16{}...))
	is.Equal(int16(1), MaxInt16s(1))
	is.Equal(int16(3), MaxInt16s(1, 2, 3))
	is.Equal(int16(3), MaxInt16s(3, 2, 1))
	is.Equal(int16(2), MaxInt16s(-3, 1, 2))

	is.Equal(int32(0), MaxInt32s())
	is.Equal(int32(0), MaxInt32s([]int32{}...))
	is.Equal(int32(1), MaxInt32s(1))
	is.Equal(int32(3), MaxInt32s(1, 2, 3))
	is.Equal(int32(3), MaxInt32s(3, 2, 1))
	is.Equal(int32(2), MaxInt32s(-3, 1, 2))

	is.Equal(int64(0), MaxInt64s())
	is.Equal(int64(0), MaxInt64s([]int64{}...))
	is.Equal(int64(1), MaxInt64s(1))
	is.Equal(int64(3), MaxInt64s(1, 2, 3))
	is.Equal(int64(3), MaxInt64s(3, 2, 1))
	is.Equal(int64(2), MaxInt64s(-3, 1, 2))

	is.Equal(float64(0), MaxFloat64s())
	is.Equal(float64(0), MaxFloat64s([]float64{}...))
	is.Equal(1.0, MaxFloat64s(1.0))
	is.Equal(3.0, MaxFloat64s(1.0, 2.0, 3.0))
	is.Equal(3.0, MaxFloat64s(3.0, 2.0, 1.0))
	is.Equal(2.0, MaxFloat64s(-3.0, 1.0, 2.0))
}
