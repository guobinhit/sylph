package unpointer

// IntOrDefault Return default value if i is nil
func IntOrDefault(i *int) int {
	if i == nil {
		return int(0)
	}
	return *i
}

// Int16OrDefault Return default value if i is nil
func Int16OrDefault(i *int16) int16 {
	if i == nil {
		return int16(0)
	}
	return *i
}

// Int32OrDefault Return default value if i is nil
func Int32OrDefault(i *int32) int32 {
	if i == nil {
		return int32(0)
	}
	return *i
}

// Int64OrDefault Return default value if i is nil
func Int64OrDefault(i *int64) int64 {
	if i == nil {
		return int64(0)
	}
	return *i
}

// Float32OrDefault Return default value if f is nil
func Float32OrDefault(f *float32) float32 {
	if f == nil {
		return float32(0)
	}
	return *f
}

// Float64OrDefault Return default value if b is nil
func Float64OrDefault(f *float64) float64 {
	if f == nil {
		return float64(0)
	}
	return *f
}

// BoolOrDefault Return default value if b is nil
func BoolOrDefault(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

// StringOrDefault Return default value if s is nil
func StringOrDefault(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
