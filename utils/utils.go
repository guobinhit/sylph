package utils

import (
	"encoding/json"
)

// If functions similar to ternary operators，if b is ture return t, else return f.
func If(b bool, t, f interface{}) interface{} {
	if b {
		return t
	}
	return f
}

// Json returns a json string.
func Json(i interface{}) string {
	if i == nil {
		return ""
	}
	v, _ := json.Marshal(i)
	return string(v)
}

// DeepCopy source and target must have the same structure, and target must be a pointer.
func DeepCopy(source, target interface{}) {
	bytes, _ := json.Marshal(source)
	_ = json.Unmarshal(bytes, target)
}
