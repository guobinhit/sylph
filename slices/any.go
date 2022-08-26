package slices

// AnyInts applies a fn to each element of s (if s is empty, returns false),
// if there exist at least one element make that fn return true then AnyInts will return true, else return false.
func AnyInts(s []int, fn func(v int) bool) bool {
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

// AnyInt8s applies a fn to each element of s (if s is empty, returns false),
// if there exist at least one element make that fn return true then AnyInt8s will return true, else return false.
func AnyInt8s(s []int8, fn func(v int8) bool) bool {
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

// AnyInt16s applies a fn to each element of s (if s is empty, returns false),
// if there exist at least one element make that fn return true then AnyInt16s will return true, else return false.
func AnyInt16s(s []int16, fn func(v int16) bool) bool {
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

// AnyInt32s applies a fn to each element of s (if s is empty, returns false),
// if there exist at least one element make that fn return true then AnyInt32s will return true, else return false.
func AnyInt32s(s []int32, fn func(v int32) bool) bool {
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

// AnyInt64s applies a fn to each element of s (if s is empty, returns false),
// if there exist at least one element make that fn return true then AnyInt64s will return true, else return false.
func AnyInt64s(s []int64, fn func(v int64) bool) bool {
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

// AnyFloat32s applies a fn to each element of s (if s is empty, returns false),
// if there exist at least one element make that fn return true then AnyFloat32s will return true, else return false.
func AnyFloat32s(s []float32, fn func(v float32) bool) bool {
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

// AnyFloat64s applies a fn to each element of s (if s is empty, returns false),
// if there exist at least one element make that fn return true then AnyFloat64s will return true, else return false.
func AnyFloat64s(s []float64, fn func(v float64) bool) bool {
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

// AnyStrings applies a fn to each element of s (if s is empty, returns false),
// if there exist at least one element make that fn return true then AnyStrings will return true, else return false.
func AnyStrings(s []string, fn func(v string) bool) bool {
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
