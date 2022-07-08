package util

// If functions similar to ternary operators，if b is ture return f, else return s
func If(b bool, f, s interface{}) interface{} {
	if b {
		return f
	}
	return s
}
