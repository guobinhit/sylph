package util

import (
	"reflect"
	"time"
)

// If functions similar to ternary operators，if b is ture return f, else return s
func If(b bool, f, s interface{}) interface{} {
	if b {
		return f
	}
	return s
}

// Equals general comparison method, if equal return true, else return false.
func Equals(a, b interface{}) bool {
	switch a1 := a.(type) {
	case string:
		b1, ok := b.(string)
		return ok && a1 == b1
	case int:
		b1, ok := b.(int)
		return ok && a1 == b1
	case int8:
		b1, ok := b.(int8)
		return ok && a1 == b1
	case int16:
		b1, ok := b.(int16)
		return ok && a1 == b1
	case int32:
		b1, ok := b.(int32)
		return ok && a1 == b1
	case int64:
		b1, ok := b.(int64)
		return ok && a1 == b1
	case uint:
		b1, ok := b.(uint)
		return ok && a1 == b1
	case uint8:
		b1, ok := b.(uint8)
		return ok && a1 == b1
	case uint16:
		b1, ok := b.(uint16)
		return ok && a1 == b1
	case uint32:
		b1, ok := b.(uint32)
		return ok && a1 == b1
	case uint64:
		b1, ok := b.(uint64)
		return ok && a1 == b1
	case float32:
		b1, ok := b.(float32)
		return ok && a1 == b1
	case float64:
		b1, ok := b.(float64)
		return ok && a1 == b1
	case time.Time:
		b1, ok := b.(time.Time)
		return ok && a1.Equal(b1)
	}
	return reflect.DeepEqual(a, b)
}
