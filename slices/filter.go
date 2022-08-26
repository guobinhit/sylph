package slices

// FilterInts applies a fn to each element of s, return a slices of make fn true.
func FilterInts(s []int, fn func(v int) bool) []int {
	ans := make([]int, 0)
	for _, v := range s {
		isOk := fn(v)
		if isOk {
			ans = append(ans, v)
		}
	}
	return ans
}

// FilterInt8s applies a fn to each element of s, return a slices of make fn true.
func FilterInt8s(s []int8, fn func(v int8) bool) []int8 {
	ans := make([]int8, 0)
	for _, v := range s {
		isOk := fn(v)
		if isOk {
			ans = append(ans, v)
		}
	}
	return ans
}

// FilterInt16s applies a fn to each element of s, return a slices of make fn true.
func FilterInt16s(s []int16, fn func(v int16) bool) []int16 {
	ans := make([]int16, 0)
	for _, v := range s {
		isOk := fn(v)
		if isOk {
			ans = append(ans, v)
		}
	}
	return ans
}

// FilterInt32s applies a fn to each element of s, return a slices of make fn true.
func FilterInt32s(s []int32, fn func(v int32) bool) []int32 {
	ans := make([]int32, 0)
	for _, v := range s {
		isOk := fn(v)
		if isOk {
			ans = append(ans, v)
		}
	}
	return ans
}

// FilterInt64s applies a fn to each element of s, return a slices of make fn true.
func FilterInt64s(s []int64, fn func(v int64) bool) []int64 {
	ans := make([]int64, 0)
	for _, v := range s {
		isOk := fn(v)
		if isOk {
			ans = append(ans, v)
		}
	}
	return ans
}

// FilterFloat32s applies a fn to each element of s, return a slices of make fn true.
func FilterFloat32s(s []float32, fn func(v float32) bool) []float32 {
	ans := make([]float32, 0)
	for _, v := range s {
		isOk := fn(v)
		if isOk {
			ans = append(ans, v)
		}
	}
	return ans
}

// FilterFloat64s applies a fn to each element of s, return a slices of make fn true.
func FilterFloat64s(s []float64, fn func(v float64) bool) []float64 {
	ans := make([]float64, 0)
	for _, v := range s {
		isOk := fn(v)
		if isOk {
			ans = append(ans, v)
		}
	}
	return ans
}

// FilterStrings applies a fn to each element of s, return a slices of make fn true.
func FilterStrings(s []string, fn func(v string) bool) []string {
	ans := make([]string, 0)
	for _, v := range s {
		isOk := fn(v)
		if isOk {
			ans = append(ans, v)
		}
	}
	return ans
}
