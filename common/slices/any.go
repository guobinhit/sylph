package slices

// AnyInt applies a fn to each element of s (if s is empty, returns false),
// if there exist at least one element make that fn return true then AnyInt will return true, else return false.
func AnyInt(s []int, fn func(v int) bool) bool {
	if len(s) == 0 {
		return false
	}
	for _, elem := range s {
		if fn(elem) {
			return true
		}
	}
	return false
}

// AnyInt64 applies a fn to each element of s (if s is empty, returns false),
// if there exist at least one element make that fn return true then AnyInt64 will return true, else return false.
func AnyInt64(s []int64, fn func(v int64) bool) bool {
	if len(s) == 0 {
		return false
	}
	for _, elem := range s {
		if fn(elem) {
			return true
		}
	}
	return false
}

// AnyString applies a fn to each element of s (if s is empty, returns false),
// if there exist at least one element make that fn return true then AnyString will return true, else return false.
func AnyString(s []string, fn func(v string) bool) bool {
	if len(s) == 0 {
		return false
	}
	for _, elem := range s {
		if fn(elem) {
			return true
		}
	}
	return false
}
