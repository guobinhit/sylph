package slices

// FilterInt applies a fn to each element of s, return a slices of make fn true.
func FilterInt(s []int, fn func(v int) bool) []int {
	ans := make([]int, 0)
	for _, v := range s {
		isOk := fn(v)
		if isOk {
			ans = append(ans, v)
		}
	}
	return ans
}

// FilterInt8 applies a fn to each element of s, return a slices of make fn true.
func FilterInt8(s []int8, fn func(v int8) bool) []int8 {
	ans := make([]int8, 0)
	for _, v := range s {
		isOk := fn(v)
		if isOk {
			ans = append(ans, v)
		}
	}
	return ans
}

// FilterInt16 applies a fn to each element of s, return a slices of make fn true.
func FilterInt16(s []int16, fn func(v int16) bool) []int16 {
	ans := make([]int16, 0)
	for _, v := range s {
		isOk := fn(v)
		if isOk {
			ans = append(ans, v)
		}
	}
	return ans
}

// FilterInt32 applies a fn to each element of s, return a slices of make fn true.
func FilterInt32(s []int32, fn func(v int32) bool) []int32 {
	ans := make([]int32, 0)
	for _, v := range s {
		isOk := fn(v)
		if isOk {
			ans = append(ans, v)
		}
	}
	return ans
}

// FilterInt64 applies a fn to each element of s, return a slices of make fn true.
func FilterInt64(s []int64, fn func(v int64) bool) []int64 {
	ans := make([]int64, 0)
	for _, v := range s {
		isOk := fn(v)
		if isOk {
			ans = append(ans, v)
		}
	}
	return ans
}

// FilterFloat32 applies a fn to each element of s, return a slices of make fn true.
func FilterFloat32(s []float32, fn func(v float32) bool) []float32 {
	ans := make([]float32, 0)
	for _, v := range s {
		isOk := fn(v)
		if isOk {
			ans = append(ans, v)
		}
	}
	return ans
}

// FilterFloat64 applies a fn to each element of s, return a slices of make fn true.
func FilterFloat64(s []float64, fn func(v float64) bool) []float64 {
	ans := make([]float64, 0)
	for _, v := range s {
		isOk := fn(v)
		if isOk {
			ans = append(ans, v)
		}
	}
	return ans
}

// FilterString applies a fn to each element of s, return a slices of make fn true.
func FilterString(s []string, fn func(v string) bool) []string {
	ans := make([]string, 0)
	for _, v := range s {
		isOk := fn(v)
		if isOk {
			ans = append(ans, v)
		}
	}
	return ans
}
