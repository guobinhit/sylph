package maths

import "math"

// MinInts returns the smallest of elements, return int(0) when values are empty.
func MinInts(values ...int) int {
	if len(values) == 0 {
		return int(0)
	}
	hitVal := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < hitVal {
			hitVal = values[i]
		}
	}
	return hitVal
}

// MinInt8s returns the smallest of elements, return int8(0) when values are empty.
func MinInt8s(values ...int8) int8 {
	if len(values) == 0 {
		return int8(0)
	}
	hitVal := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < hitVal {
			hitVal = values[i]
		}
	}
	return hitVal
}

// MinInt16s returns the smallest of elements, return int16(0) when values are empty.
func MinInt16s(values ...int16) int16 {
	if len(values) == 0 {
		return int16(0)
	}
	hitVal := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < hitVal {
			hitVal = values[i]
		}
	}
	return hitVal
}

// MinInt32s returns the smallest of elements, return int32(0) when values are empty.
func MinInt32s(values ...int32) int32 {
	if len(values) == 0 {
		return int32(0)
	}
	hitVal := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < hitVal {
			hitVal = values[i]
		}
	}
	return hitVal
}

// MinInt64s returns the smallest of elements, return int64(0) when values are empty.
func MinInt64s(values ...int64) int64 {
	if len(values) == 0 {
		return int64(0)
	}
	hitVal := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < hitVal {
			hitVal = values[i]
		}
	}
	return hitVal
}

// MinFloat64s returns the smallest of elements, return float64(0) when values are empty.
func MinFloat64s(values ...float64) float64 {
	if len(values) == 0 {
		return float64(0)
	}
	hitVal := values[0]
	for i := 1; i < len(values); i++ {
		hitVal = math.Min(hitVal, values[i])
	}
	return hitVal
}
