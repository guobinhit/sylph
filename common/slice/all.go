package slice

// AllInt applies a fn to each element of s (if s is empty, returns false),
// if all elements make that fn return true then AllInt will return true, else return false.
func AllInt(s []int, fn func(v int) bool) bool {
	if len(s) == 0 {
		return true
	}
	for _, elem := range s {
		if !fn(elem) {
			return false
		}
	}
	return true
}

// AllInt64 applies a fn to each element of s (if s is empty, returns false),
// if all elements make that fn return true then AllInt64 will return true, else return false.
func AllInt64(s []int64, fn func(v int64) bool) bool {
	if len(s) == 0 {
		return true
	}
	for _, elem := range s {
		if !fn(elem) {
			return false
		}
	}
	return true
}

// AllString applies a fn to each element of s (if s is empty, returns false),
// if all elements make that fn return true then AllString will return true, else return false.
func AllString(s []string, fn func(v string) bool) bool {
	if len(s) == 0 {
		return true
	}
	for _, elem := range s {
		if !fn(elem) {
			return false
		}
	}
	return true
}
