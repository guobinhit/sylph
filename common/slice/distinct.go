package slice

// DistinctSliceInt remove duplicate element in slice.
func DistinctSliceInt(s []int) []int {
	if s == nil || len(s) == 0 {
		return s
	}
	ans := make([]int, 0)
	aMap := make(map[int]bool)
	for _, i := range s {
		if ok := aMap[i]; !ok {
			ans = append(ans, i)
			aMap[i] = true
		}
	}
	return ans
}

// DistinctSliceInt8 remove duplicate element in slice.
func DistinctSliceInt8(s []int8) []int8 {
	if s == nil || len(s) == 0 {
		return s
	}
	ans := make([]int8, 0)
	aMap := make(map[int8]bool)
	for _, i := range s {
		if ok := aMap[i]; !ok {
			ans = append(ans, i)
			aMap[i] = true
		}
	}
	return ans
}

// DistinctSliceInt16 remove duplicate element in slice.
func DistinctSliceInt16(s []int16) []int16 {
	if s == nil || len(s) == 0 {
		return s
	}
	ans := make([]int16, 0)
	aMap := make(map[int16]bool)
	for _, i := range s {
		if ok := aMap[i]; !ok {
			ans = append(ans, i)
			aMap[i] = true
		}
	}
	return ans
}

// DistinctSliceInt32 remove duplicate element in slice.
func DistinctSliceInt32(s []int32) []int32 {
	if s == nil || len(s) == 0 {
		return s
	}
	ans := make([]int32, 0)
	aMap := make(map[int32]bool)
	for _, i := range s {
		if ok := aMap[i]; !ok {
			ans = append(ans, i)
			aMap[i] = true
		}
	}
	return ans
}

// DistinctSliceInt64 remove duplicate element in slice.
func DistinctSliceInt64(s []int64) []int64 {
	if s == nil || len(s) == 0 {
		return s
	}
	ans := make([]int64, 0)
	aMap := make(map[int64]bool)
	for _, i := range s {
		if ok := aMap[i]; !ok {
			ans = append(ans, i)
			aMap[i] = true
		}
	}
	return ans
}

// DistinctSliceFloat32 remove duplicate element in slice.
func DistinctSliceFloat32(s []float32) []float32 {
	if s == nil || len(s) == 0 {
		return s
	}
	ans := make([]float32, 0)
	aMap := make(map[float32]bool)
	for _, i := range s {
		if ok := aMap[i]; !ok {
			ans = append(ans, i)
			aMap[i] = true
		}
	}
	return ans
}

// DistinctSliceFloat64 remove duplicate element in slice.
func DistinctSliceFloat64(s []float64) []float64 {
	if s == nil || len(s) == 0 {
		return s
	}
	ans := make([]float64, 0)
	aMap := make(map[float64]bool)
	for _, i := range s {
		if ok := aMap[i]; !ok {
			ans = append(ans, i)
			aMap[i] = true
		}
	}
	return ans
}

// DistinctSliceString remove duplicate element in slice.
func DistinctSliceString(s []string) []string {
	if s == nil || len(s) == 0 {
		return s
	}
	ans := make([]string, 0)
	aMap := make(map[string]bool)
	for _, i := range s {
		if ok := aMap[i]; !ok {
			ans = append(ans, i)
			aMap[i] = true
		}
	}
	return ans
}
