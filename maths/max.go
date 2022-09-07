package maths

import "math"

// MaxInts returns the largest of elements, return int(0) when values are empty.
func MaxInts(values ...int) int {
	if len(values) == 0 {
		return int(0)
	}
	hitVal := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > hitVal {
			hitVal = values[i]
		}
	}
	return hitVal
}

// MaxInt8s returns the largest of elements, return int8(0) when values are empty.
func MaxInt8s(values ...int8) int8 {
	if len(values) == 0 {
		return int8(0)
	}
	hitVal := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > hitVal {
			hitVal = values[i]
		}
	}
	return hitVal
}

// MaxInt16s returns the largest of elements, return int16(0) when values are empty.
func MaxInt16s(values ...int16) int16 {
	if len(values) == 0 {
		return int16(0)
	}
	hitVal := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > hitVal {
			hitVal = values[i]
		}
	}
	return hitVal
}

// MaxInt32s returns the largest of elements, return int32(0) when values are empty.
func MaxInt32s(values ...int32) int32 {
	if len(values) == 0 {
		return int32(0)
	}
	hitVal := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > hitVal {
			hitVal = values[i]
		}
	}
	return hitVal
}

// MaxInt64s returns the largest of elements, return int64(0) when values are empty.
func MaxInt64s(values ...int64) int64 {
	if len(values) == 0 {
		return int64(0)
	}
	hitVal := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > hitVal {
			hitVal = values[i]
		}
	}
	return hitVal
}

// MaxFloat64s returns the largest of elements, return float64(0) when values are empty.
func MaxFloat64s(values ...float64) float64 {
	if len(values) == 0 {
		return float64(0)
	}
	hitVal := values[0]
	for i := 1; i < len(values); i++ {
		hitVal = math.Max(hitVal, values[i])
	}
	return hitVal
}
