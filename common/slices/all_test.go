package slices

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAllInt(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AllInt([]int{11, 20}, func(v int) bool {
		return v > 0
	}))
	is.Equal(false, AllInt([]int{11, 20}, func(v int) bool {
		return v%3 == 0
	}))
}

func TestAllInt8(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AllInt8([]int8{11, 20}, func(v int8) bool {
		return v > 0
	}))
	is.Equal(false, AllInt8([]int8{11, 20}, func(v int8) bool {
		return v%3 == 0
	}))
}

func TestAllInt16(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AllInt16([]int16{11, 20}, func(v int16) bool {
		return v > 0
	}))
	is.Equal(false, AllInt16([]int16{11, 20}, func(v int16) bool {
		return v%3 == 0
	}))
}

func TestAllInt32(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AllInt32([]int32{11, 20}, func(v int32) bool {
		return v > 0
	}))
	is.Equal(false, AllInt32([]int32{11, 20}, func(v int32) bool {
		return v%3 == 0
	}))
}

func TestAllInt64(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AllInt64([]int64{11, 20}, func(v int64) bool {
		return v > 0
	}))
	is.Equal(false, AllInt64([]int64{11, 20}, func(v int64) bool {
		return v%3 == 0
	}))
}

func TestAllFloat32(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AllFloat32([]float32{11, 20}, func(v float32) bool {
		return v > 0
	}))
	is.Equal(false, AllFloat32([]float32{11, 20}, func(v float32) bool {
		return v > 100
	}))
}

func TestAllFloat64(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AllFloat64([]float64{11, 20}, func(v float64) bool {
		return v > 0
	}))
	is.Equal(false, AllFloat64([]float64{11, 20}, func(v float64) bool {
		return v > 100
	}))
}

func TestAllString(t *testing.T) {
	is := assert.New(t)
	is.Equal(false, AllString([]string{"abc", "zora"}, func(v string) bool {
		return len(v) > 10
	}))
	is.Equal(true, AllString([]string{"abc", "zora"}, func(v string) bool {
		return len(v) > 0
	}))
}
