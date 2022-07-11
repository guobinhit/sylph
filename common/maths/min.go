package maths

import "math"

// MinInts returns the smallest of elements, panics when values are empty.
func MinInts(values ...int) int {
	if len(values) == 0 {
		panic(valuesIsEmptyMsg)
	}
	hitVal := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < hitVal {
			hitVal = values[i]
		}
	}
	return hitVal
}

// MinInt8s returns the smallest of elements, panics when values are empty.
func MinInt8s(values ...int8) int8 {
	if len(values) == 0 {
		panic(valuesIsEmptyMsg)
	}
	hitVal := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < hitVal {
			hitVal = values[i]
		}
	}
	return hitVal
}

// MinInt16s returns the smallest of elements, panics when values are empty.
func MinInt16s(values ...int16) int16 {
	if len(values) == 0 {
		panic(valuesIsEmptyMsg)
	}
	hitVal := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < hitVal {
			hitVal = values[i]
		}
	}
	return hitVal
}

// MinInt32s returns the smallest of elements, panics when values are empty.
func MinInt32s(values ...int32) int32 {
	if len(values) == 0 {
		panic(valuesIsEmptyMsg)
	}
	hitVal := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < hitVal {
			hitVal = values[i]
		}
	}
	return hitVal
}

// MinInt64s returns the smallest of elements, panics when values are empty.
func MinInt64s(values ...int64) int64 {
	if len(values) == 0 {
		panic(valuesIsEmptyMsg)
	}
	hitVal := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < hitVal {
			hitVal = values[i]
		}
	}
	return hitVal
}

// MinFloat64s returns the smallest of elements, panics when values are empty.
func MinFloat64s(values ...float64) float64 {
	if len(values) == 0 {
		panic(valuesIsEmptyMsg)
	}
	hitVal := values[0]
	for i := 1; i < len(values); i++ {
		hitVal = math.Min(hitVal, values[i])
	}
	return hitVal
}
