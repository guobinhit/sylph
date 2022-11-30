package strings

import "strings"

// Equals returns true if s1 equal s2.
func Equals(s1, s2 string) bool {
	return s1 == s2
}

// EqualsIgnoreCase returns true if s1 equal s2, ignore case of element.
func EqualsIgnoreCase(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if strings.ToLower(s1) != strings.ToLower(s2) {
		return false
	}
	return true
}

// StarsWith tests whether the string s begins with prefix, wrapper strings.HasPrefix method.
func StarsWith(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// EndsWith tests whether the string s ends with suffix, wrapper strings.HasSuffix method.
func EndsWith(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}
