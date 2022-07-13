package slices

// DistinctSliceInt remove duplicate element in slices.
func DistinctSliceInt(s []int) []int {
	if s == nil || len(s) == 0 {
		return s
	}
	ans := make([]int, 0)
	aMap := make(map[int]struct{})
	for _, i := range s {
		if _, ok := aMap[i]; !ok {
			ans = append(ans, i)
			aMap[i] = struct{}{}
		}
	}
	return ans
}

// DistinctSliceInt8 remove duplicate element in slices.
func DistinctSliceInt8(s []int8) []int8 {
	if s == nil || len(s) == 0 {
		return s
	}
	ans := make([]int8, 0)
	aMap := make(map[int8]struct{})
	for _, i := range s {
		if _, ok := aMap[i]; !ok {
			ans = append(ans, i)
			aMap[i] = struct{}{}
		}
	}
	return ans
}

// DistinctSliceInt16 remove duplicate element in slices.
func DistinctSliceInt16(s []int16) []int16 {
	if s == nil || len(s) == 0 {
		return s
	}
	ans := make([]int16, 0)
	aMap := make(map[int16]struct{})
	for _, i := range s {
		if _, ok := aMap[i]; !ok {
			ans = append(ans, i)
			aMap[i] = struct{}{}
		}
	}
	return ans
}

// DistinctSliceInt32 remove duplicate element in slices.
func DistinctSliceInt32(s []int32) []int32 {
	if s == nil || len(s) == 0 {
		return s
	}
	ans := make([]int32, 0)
	aMap := make(map[int32]struct{})
	for _, i := range s {
		if _, ok := aMap[i]; !ok {
			ans = append(ans, i)
			aMap[i] = struct{}{}
		}
	}
	return ans
}

// DistinctSliceInt64 remove duplicate element in slices.
func DistinctSliceInt64(s []int64) []int64 {
	if s == nil || len(s) == 0 {
		return s
	}
	ans := make([]int64, 0)
	aMap := make(map[int64]struct{})
	for _, i := range s {
		if _, ok := aMap[i]; !ok {
			ans = append(ans, i)
			aMap[i] = struct{}{}
		}
	}
	return ans
}

// DistinctSliceFloat32 remove duplicate element in slices.
func DistinctSliceFloat32(s []float32) []float32 {
	if s == nil || len(s) == 0 {
		return s
	}
	ans := make([]float32, 0)
	aMap := make(map[float32]struct{})
	for _, i := range s {
		if _, ok := aMap[i]; !ok {
			ans = append(ans, i)
			aMap[i] = struct{}{}
		}
	}
	return ans
}

// DistinctSliceFloat64 remove duplicate element in slices.
func DistinctSliceFloat64(s []float64) []float64 {
	if s == nil || len(s) == 0 {
		return s
	}
	ans := make([]float64, 0)
	aMap := make(map[float64]struct{})
	for _, i := range s {
		if _, ok := aMap[i]; !ok {
			ans = append(ans, i)
			aMap[i] = struct{}{}
		}
	}
	return ans
}

// DistinctSliceString remove duplicate element in slices.
func DistinctSliceString(s []string) []string {
	if s == nil || len(s) == 0 {
		return s
	}
	ans := make([]string, 0)
	aMap := make(map[string]struct{})
	for _, i := range s {
		if _, ok := aMap[i]; !ok {
			ans = append(ans, i)
			aMap[i] = struct{}{}
		}
	}
	return ans
}
