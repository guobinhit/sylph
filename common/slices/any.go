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

// AnyInt8 applies a fn to each element of s (if s is empty, returns false),
// if there exist at least one element make that fn return true then AnyInt8 will return true, else return false.
func AnyInt8(s []int8, fn func(v int8) bool) bool {
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

// AnyInt16 applies a fn to each element of s (if s is empty, returns false),
// if there exist at least one element make that fn return true then AnyInt16 will return true, else return false.
func AnyInt16(s []int16, fn func(v int16) bool) bool {
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

// AnyInt32 applies a fn to each element of s (if s is empty, returns false),
// if there exist at least one element make that fn return true then AnyInt32 will return true, else return false.
func AnyInt32(s []int32, fn func(v int32) bool) bool {
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

// AnyFloat32 applies a fn to each element of s (if s is empty, returns false),
// if there exist at least one element make that fn return true then AnyFloat32 will return true, else return false.
func AnyFloat32(s []float32, fn func(v float32) bool) bool {
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

// AnyFloat64 applies a fn to each element of s (if s is empty, returns false),
// if there exist at least one element make that fn return true then AnyFloat64 will return true, else return false.
func AnyFloat64(s []float64, fn func(v float64) bool) bool {
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
