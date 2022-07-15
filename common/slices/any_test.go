package slices

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnyInt(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AnyInt([]int{11, 20}, func(v int) bool {
		return v%2 == 0
	}))
	is.Equal(false, AnyInt([]int{11, 20}, func(v int) bool {
		return v%3 == 0
	}))
}

func TestAnyInt8(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AnyInt8([]int8{11, 20}, func(v int8) bool {
		return v%2 == 0
	}))
	is.Equal(false, AnyInt8([]int8{11, 20}, func(v int8) bool {
		return v%3 == 0
	}))
}

func TestAnyInt16(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AnyInt16([]int16{11, 20}, func(v int16) bool {
		return v%2 == 0
	}))
	is.Equal(false, AnyInt16([]int16{11, 20}, func(v int16) bool {
		return v%3 == 0
	}))
}

func TestAnyInt32(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AnyInt32([]int32{11, 20}, func(v int32) bool {
		return v%2 == 0
	}))
	is.Equal(false, AnyInt32([]int32{11, 20}, func(v int32) bool {
		return v%3 == 0
	}))
}

func TestAnyInt64(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AnyInt64([]int64{11, 20}, func(v int64) bool {
		return v%2 == 0
	}))
	is.Equal(false, AnyInt64([]int64{11, 20}, func(v int64) bool {
		return v%3 == 0
	}))
}

func TestAnFloat32(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AllFloat32([]float32{11, 20}, func(v float32) bool {
		return v > 0
	}))
	is.Equal(false, AllFloat32([]float32{11, 20}, func(v float32) bool {
		return v > 15
	}))
}

func TestAnFloat64(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AllFloat64([]float64{11, 20}, func(v float64) bool {
		return v > 0
	}))
	is.Equal(false, AllFloat64([]float64{11, 20}, func(v float64) bool {
		return v > 15
	}))
}

func TestAnyString(t *testing.T) {
	is := assert.New(t)
	is.Equal(false, AnyString([]string{"abc", "zora"}, func(v string) bool {
		return len(v) > 10
	}))
	is.Equal(true, AnyString([]string{"abc", "zora"}, func(v string) bool {
		return len(v) > 3
	}))
}
