package slices

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

// AllInt8 applies a fn to each element of s (if s is empty, returns false),
// if all elements make that fn return true then AllInt8 will return true, else return false.
func AllInt8(s []int8, fn func(v int8) bool) bool {
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

// AllInt16 applies a fn to each element of s (if s is empty, returns false),
// if all elements make that fn return true then AllInt16 will return true, else return false.
func AllInt16(s []int16, fn func(v int16) bool) bool {
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

// AllInt32 applies a fn to each element of s (if s is empty, returns false),
// if all elements make that fn return true then AllInt32 will return true, else return false.
func AllInt32(s []int32, fn func(v int32) bool) bool {
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

// AllFloat32 applies a fn to each element of s (if s is empty, returns false),
// if all elements make that fn return true then AllFloat32 will return true, else return false.
func AllFloat32(s []float32, fn func(v float32) bool) bool {
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

// AllFloat64 applies a fn to each element of s (if s is empty, returns false),
// if all elements make that fn return true then AllFloat64 will return true, else return false.
func AllFloat64(s []float64, fn func(v float64) bool) bool {
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
