package slices

import "strings"

// ContainInts returns true if s contain e.
func ContainInts(s []int, e int) bool {
	if s == nil {
		return false
	}
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

// ContainInt8s returns true if s contain e.
func ContainInt8s(s []int8, e int8) bool {
	if s == nil {
		return false
	}
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

// ContainInt16s returns true if s contain e.
func ContainInt16s(s []int16, e int16) bool {
	if s == nil {
		return false
	}
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

// ContainInt32s returns true if s contain e.
func ContainInt32s(s []int32, e int32) bool {
	if s == nil {
		return false
	}
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

// ContainInt64s returns true if s contain e.
func ContainInt64s(s []int64, e int64) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

// ContainFloat32s returns true if s contain e.
func ContainFloat32s(s []float32, e float32) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

// ContainFloat64s returns true if s contain e.
func ContainFloat64s(s []float64, e float64) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

// ContainStrings returns true if s contain e.
func ContainStrings(s []string, e string) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

// ContainStringsIgnoreCase returns true if s contain e, ignore case of element.
func ContainStringsIgnoreCase(s []string, e string) bool {
	for _, v := range s {
		if strings.ToLower(v) == strings.ToLower(e) {
			return true
		}
	}
	return false
}
