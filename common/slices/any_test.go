package slices

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnyInts(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AnyInts([]int{11, 20}, func(v int) bool {
		return v%2 == 0
	}))
	is.Equal(false, AnyInts([]int{11, 20}, func(v int) bool {
		return v%3 == 0
	}))
}

func TestAnyInt8s(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AnyInt8s([]int8{11, 20}, func(v int8) bool {
		return v%2 == 0
	}))
	is.Equal(false, AnyInt8s([]int8{11, 20}, func(v int8) bool {
		return v%3 == 0
	}))
}

func TestAnyInt16s(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AnyInt16s([]int16{11, 20}, func(v int16) bool {
		return v%2 == 0
	}))
	is.Equal(false, AnyInt16s([]int16{11, 20}, func(v int16) bool {
		return v%3 == 0
	}))
}

func TestAnyInt32s(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AnyInt32s([]int32{11, 20}, func(v int32) bool {
		return v%2 == 0
	}))
	is.Equal(false, AnyInt32s([]int32{11, 20}, func(v int32) bool {
		return v%3 == 0
	}))
}

func TestAnyInt64s(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AnyInt64s([]int64{11, 20}, func(v int64) bool {
		return v%2 == 0
	}))
	is.Equal(false, AnyInt64s([]int64{11, 20}, func(v int64) bool {
		return v%3 == 0
	}))
}

func TestAnyFloat32s(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AnyFloat32s([]float32{11, 20}, func(v float32) bool {
		return v > 0
	}))
	is.Equal(true, AnyFloat32s([]float32{11, 20}, func(v float32) bool {
		return v > 15
	}))
}

func TestAnyFloat64s(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AnyFloat64s([]float64{11, 20}, func(v float64) bool {
		return v > 0
	}))
	is.Equal(true, AnyFloat64s([]float64{11, 20}, func(v float64) bool {
		return v > 15
	}))
}

func TestAnyStrings(t *testing.T) {
	is := assert.New(t)
	is.Equal(false, AnyStrings([]string{"abc", "zora"}, func(v string) bool {
		return len(v) > 10
	}))
	is.Equal(true, AnyStrings([]string{"abc", "zora"}, func(v string) bool {
		return len(v) > 3
	}))
}
