package slices

// AllInts applies a fn to each element of s (if s is empty, returns false),
// if all elements make that fn return true then AllInts will return true, else return false.
func AllInts(s []int, fn func(v int) bool) bool {
	if len(s) == 0 {
		return true
	}
	for _, v := range s {
		if !fn(v) {
			return false
		}
	}
	return true
}

// AllInt8s applies a fn to each element of s (if s is empty, returns false),
// if all elements make that fn return true then AllInt8s will return true, else return false.
func AllInt8s(s []int8, fn func(v int8) bool) bool {
	if len(s) == 0 {
		return true
	}
	for _, v := range s {
		if !fn(v) {
			return false
		}
	}
	return true
}

// AllInt16s applies a fn to each element of s (if s is empty, returns false),
// if all elements make that fn return true then AllInt16s will return true, else return false.
func AllInt16s(s []int16, fn func(v int16) bool) bool {
	if len(s) == 0 {
		return true
	}
	for _, v := range s {
		if !fn(v) {
			return false
		}
	}
	return true
}

// AllInt32s applies a fn to each element of s (if s is empty, returns false),
// if all elements make that fn return true then AllInt32s will return true, else return false.
func AllInt32s(s []int32, fn func(v int32) bool) bool {
	if len(s) == 0 {
		return true
	}
	for _, v := range s {
		if !fn(v) {
			return false
		}
	}
	return true
}

// AllInt64s applies a fn to each element of s (if s is empty, returns false),
// if all elements make that fn return true then AllInt64s will return true, else return false.
func AllInt64s(s []int64, fn func(v int64) bool) bool {
	if len(s) == 0 {
		return true
	}
	for _, v := range s {
		if !fn(v) {
			return false
		}
	}
	return true
}

// AllFloat32s applies a fn to each element of s (if s is empty, returns false),
// if all elements make that fn return true then AllFloat32s will return true, else return false.
func AllFloat32s(s []float32, fn func(v float32) bool) bool {
	if len(s) == 0 {
		return true
	}
	for _, v := range s {
		if !fn(v) {
			return false
		}
	}
	return true
}

// AllFloat64s applies a fn to each element of s (if s is empty, returns false),
// if all elements make that fn return true then AllFloat64s will return true, else return false.
func AllFloat64s(s []float64, fn func(v float64) bool) bool {
	if len(s) == 0 {
		return true
	}
	for _, v := range s {
		if !fn(v) {
			return false
		}
	}
	return true
}

// AllStrings applies a fn to each element of s (if s is empty, returns false),
// if all elements make that fn return true then AllStrings will return true, else return false.
func AllStrings(s []string, fn func(v string) bool) bool {
	if len(s) == 0 {
		return true
	}
	for _, v := range s {
		if !fn(v) {
			return false
		}
	}
	return true
}
