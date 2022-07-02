package string

import "strings"

// Equals Return true if s1 equal s2
func Equals(s1, s2 string) bool {
	return s1 == s2
}

// EqualsIgnoreCase Return true if s1 equal s2, ignore case of element
func EqualsIgnoreCase(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if strings.ToLower(s1) != strings.ToLower(s2) {
		return false
	}
	return true
}
