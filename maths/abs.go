package maths

import (
	"fmt"
	"math"
)

// AbsInt returns the absolute value of v, panic when overflowed.
func AbsInt(v int) int {
	if v <= minInt {
		panic(fmt.Errorf("int overflow: abs(%d)", v))
	}
	if v < 0 {
		v = -v
	}
	return v
}

// AbsInt8 returns the absolute value of v, panic when overflowed.
func AbsInt8(v int8) int8 {
	if v <= math.MinInt8 {
		panic(fmt.Errorf("int8 overflow: abs(%d)", v))
	}
	if v < 0 {
		v = -v
	}
	return v
}

// AbsInt16 returns the absolute value of v, panic when overflowed.
func AbsInt16(v int16) int16 {
	if v <= math.MinInt16 {
		panic(fmt.Errorf("int16 overflow: abs(%d)", v))
	}
	if v < 0 {
		v = -v
	}
	return v
}

// AbsInt32 returns the absolute value of v, panic when overflowed.
func AbsInt32(v int32) int32 {
	if v <= math.MinInt32 {
		panic(fmt.Errorf("int32 overflow: abs(%d)", v))
	}
	if v < 0 {
		v = -v
	}
	return v
}

// AbsInt64 returns the absolute value of v, panic when overflowed.
func AbsInt64(v int64) int64 {
	if v <= math.MinInt64 {
		panic(fmt.Errorf("int64 overflow: abs(%d)", v))
	}
	if v < 0 {
		v = -v
	}
	return v
}

// AbsFloat32 returns the absolute value of v.
func AbsFloat32(x float32) float32 {
	return math.Float32frombits(math.Float32bits(x) &^ (1 << 31))
}

// AbsFloat64 returns the absolute value of v.
func AbsFloat64(v float64) float64 {
	return math.Abs(v)
}
