package slice

import (
	"strings"
)

// IntContains Return true if s contain e
func IntContains(s []int, e int) bool {
	if s == nil {
		return false
	}
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Int8Contains Return true if s contain e
func Int8Contains(s []int8, e int8) bool {
	if s == nil {
		return false
	}
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Int16Contains Return true if s contain e
func Int16Contains(s []int16, e int16) bool {
	if s == nil {
		return false
	}
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Int32Contains Return true if s contain e
func Int32Contains(s []int32, e int32) bool {
	if s == nil {
		return false
	}
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Int64Contains Return true if s contain e
func Int64Contains(s []int64, e int64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Float32Contains Return true if s contain e
func Float32Contains(s []float32, e float32) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Float64Contains Return true if s contain e
func Float64Contains(s []float64, e float64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// BoolContains Return true if s contain e
func BoolContains(s []bool, e bool) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// StringContains Return true if s contain e
func StringContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// StringContainsIgnoreCase Return true if s contain e, ignore case of element
func StringContainsIgnoreCase(s []string, e string) bool {
	for _, a := range s {
		if strings.ToLower(a) == strings.ToLower(e) {
			return true
		}
	}
	return false
}

// DistinctSliceInt Remove duplicate element in slice
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

// DistinctSliceInt8 Remove duplicate element in slice
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

// DistinctSliceInt16 Remove duplicate element in slice
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

// DistinctSliceInt32 Remove duplicate element in slice
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

// DistinctSliceInt64 Remove duplicate element in slice
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

// DistinctSliceFloat32 Remove duplicate element in slice
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

// DistinctSliceFloat64 Remove duplicate element in slice
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

// DistinctSliceString Remove duplicate element in slice
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
