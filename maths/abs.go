package maths

import (
	"fmt"
	"math"
)

// AbsInt8 returns the absolute value of v, panic when overflowed.
func AbsInt8(v int8) int8 {
	if v <= math.MinInt8 {
		panic(fmt.Sprintf("int8 overflow: abs(%d)", v))
	}
	if v < 0 {
		v = -v
	}
	return v
}

// AbsInt16 returns the absolute value of v, panic when overflowed.
func AbsInt16(v int16) int16 {
	if v <= math.MinInt16 {
		panic(fmt.Sprintf("int16 overflow: abs(%d)", v))
	}
	if v < 0 {
		v = -v
	}
	return v
}

// AbsInt32 returns the absolute value of v, panic when overflowed.
func AbsInt32(v int32) int32 {
	if v <= math.MinInt32 {
		panic(fmt.Sprintf("int32 overflow: abs(%d)", v))
	}
	if v < 0 {
		v = -v
	}
	return v
}

// AbsInt64 returns the absolute value of v, panic when overflowed.
func AbsInt64(v int64) int64 {
	if v <= math.MinInt64 {
		panic(fmt.Sprintf("int64 overflow: abs(%d)", v))
	}
	if v < 0 {
		v = -v
	}
	return v
}

// AbsFloat32 returns the absolute value of v, panic when overflowed.
func AbsFloat32(v float32) float32 {
	return math.Float32frombits(math.Float32bits(v) &^ (1 << 31))
}

// AbsFloat64 returns the absolute value of v, panic when overflowed.
func AbsFloat64(v float64) float64 {
	return math.Abs(v)
}
