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
