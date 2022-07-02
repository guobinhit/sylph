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
