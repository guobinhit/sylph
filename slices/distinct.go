package slices

// DistinctInts remove duplicate element in slices.
func DistinctInts(s []int) []int {
	if s == nil || len(s) == 0 {
		return s
	}
	ans := make([]int, 0)
	aMap := make(map[int]struct{})
	for _, v := range s {
		if _, ok := aMap[v]; !ok {
			ans = append(ans, v)
			aMap[v] = struct{}{}
		}
	}
	return ans
}

// DistinctInt8s remove duplicate element in slices.
func DistinctInt8s(s []int8) []int8 {
	if s == nil || len(s) == 0 {
		return s
	}
	ans := make([]int8, 0)
	aMap := make(map[int8]struct{})
	for _, v := range s {
		if _, ok := aMap[v]; !ok {
			ans = append(ans, v)
			aMap[v] = struct{}{}
		}
	}
	return ans
}

// DistinctInt16s remove duplicate element in slices.
func DistinctInt16s(s []int16) []int16 {
	if s == nil || len(s) == 0 {
		return s
	}
	ans := make([]int16, 0)
	aMap := make(map[int16]struct{})
	for _, v := range s {
		if _, ok := aMap[v]; !ok {
			ans = append(ans, v)
			aMap[v] = struct{}{}
		}
	}
	return ans
}

// DistinctInt32s remove duplicate element in slices.
func DistinctInt32s(s []int32) []int32 {
	if s == nil || len(s) == 0 {
		return s
	}
	ans := make([]int32, 0)
	aMap := make(map[int32]struct{})
	for _, v := range s {
		if _, ok := aMap[v]; !ok {
			ans = append(ans, v)
			aMap[v] = struct{}{}
		}
	}
	return ans
}

// DistinctInt64s remove duplicate element in slices.
func DistinctInt64s(s []int64) []int64 {
	if s == nil || len(s) == 0 {
		return s
	}
	ans := make([]int64, 0)
	aMap := make(map[int64]struct{})
	for _, v := range s {
		if _, ok := aMap[v]; !ok {
			ans = append(ans, v)
			aMap[v] = struct{}{}
		}
	}
	return ans
}

// DistinctFloat32s remove duplicate element in slices.
func DistinctFloat32s(s []float32) []float32 {
	if s == nil || len(s) == 0 {
		return s
	}
	ans := make([]float32, 0)
	aMap := make(map[float32]struct{})
	for _, v := range s {
		if _, ok := aMap[v]; !ok {
			ans = append(ans, v)
			aMap[v] = struct{}{}
		}
	}
	return ans
}

// DistinctFloat64s remove duplicate element in slices.
func DistinctFloat64s(s []float64) []float64 {
	if s == nil || len(s) == 0 {
		return s
	}
	ans := make([]float64, 0)
	aMap := make(map[float64]struct{})
	for _, v := range s {
		if _, ok := aMap[v]; !ok {
			ans = append(ans, v)
			aMap[v] = struct{}{}
		}
	}
	return ans
}

// DistinctStrings remove duplicate element in slices.
func DistinctStrings(s []string) []string {
	if s == nil || len(s) == 0 {
		return s
	}
	ans := make([]string, 0)
	aMap := make(map[string]struct{})
	for _, v := range s {
		if _, ok := aMap[v]; !ok {
			ans = append(ans, v)
			aMap[v] = struct{}{}
		}
	}
	return ans
}
