package unpointer

// Int returns default value if i is nil.
func Int(i *int) int {
	if i == nil {
		return int(0)
	}
	return *i
}

// Int8 returns default value if i is nil.
func Int8(i *int8) int8 {
	if i == nil {
		return int8(0)
	}
	return *i
}

// Int16 returns default value if i is nil.
func Int16(i *int16) int16 {
	if i == nil {
		return int16(0)
	}
	return *i
}

// Int32 returns default value if i is nil.
func Int32(i *int32) int32 {
	if i == nil {
		return int32(0)
	}
	return *i
}

// Int64 returns default value if i is nil.
func Int64(i *int64) int64 {
	if i == nil {
		return int64(0)
	}
	return *i
}

// Float32 returns default value if f is nil.
func Float32(f *float32) float32 {
	if f == nil {
		return float32(0)
	}
	return *f
}

// Float64 returns default value if f is nil.
func Float64(f *float64) float64 {
	if f == nil {
		return float64(0)
	}
	return *f
}

// Bool returns default value if b is nil.
func Bool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

// String returns default value if s is nil.
func String(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// IntOrDefault returns dv if i is nil.
func IntOrDefault(i *int, dv int) int {
	if i == nil {
		return dv
	}
	return *i
}

// Int8OrDefault returns dv if i is nil.
func Int8OrDefault(i *int8, dv int8) int8 {
	if i == nil {
		return dv
	}
	return *i
}

// Int16OrDefault returns dv if i is nil.
func Int16OrDefault(i *int16, dv int16) int16 {
	if i == nil {
		return dv
	}
	return *i
}

// Int32OrDefault returns dv if i is nil.
func Int32OrDefault(i *int32, dv int32) int32 {
	if i == nil {
		return dv
	}
	return *i
}

// Int64OrDefault returns default value if i is nil.
func Int64OrDefault(i *int64, dv int64) int64 {
	if i == nil {
		return dv
	}
	return *i
}

// Float32OrDefault returns dv if f is nil.
func Float32OrDefault(f *float32, dv float32) float32 {
	if f == nil {
		return dv
	}
	return *f
}

// Float64OrDefault returns dv if f is nil.
func Float64OrDefault(f *float64, dv float64) float64 {
	if f == nil {
		return dv
	}
	return *f
}

// BoolOrDefault returns dv if b is nil.
func BoolOrDefault(b *bool, dv bool) bool {
	if b == nil {
		return dv
	}
	return *b
}

// StringOrDefault returns dv if s is nil.
func StringOrDefault(s *string, dv string) string {
	if s == nil {
		return dv
	}
	return *s
}
