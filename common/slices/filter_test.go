package slices

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilterInts(t *testing.T) {
	is := assert.New(t)
	is.Equal([]int{11, 20}, FilterInts([]int{11, 20}, func(v int) bool {
		return v > 0
	}))
	is.Equal([]int{20}, FilterInts([]int{11, 20}, func(v int) bool {
		return v%2 == 0
	}))
}

func TestFilterInt8s(t *testing.T) {
	is := assert.New(t)
	is.Equal([]int8{11, 20}, FilterInt8s([]int8{11, 20}, func(v int8) bool {
		return v > 0
	}))
	is.Equal([]int8{20}, FilterInt8s([]int8{11, 20}, func(v int8) bool {
		return v%2 == 0
	}))
}

func TestFilterInt16s(t *testing.T) {
	is := assert.New(t)
	is.Equal([]int16{11, 20}, FilterInt16s([]int16{11, 20}, func(v int16) bool {
		return v > 0
	}))
	is.Equal([]int16{20}, FilterInt16s([]int16{11, 20}, func(v int16) bool {
		return v%2 == 0
	}))
}

func TestFilterInt32s(t *testing.T) {
	is := assert.New(t)
	is.Equal([]int32{11, 20}, FilterInt32s([]int32{11, 20}, func(v int32) bool {
		return v > 0
	}))
	is.Equal([]int32{20}, FilterInt32s([]int32{11, 20}, func(v int32) bool {
		return v%2 == 0
	}))
}

func TestFilterInt64s(t *testing.T) {
	is := assert.New(t)
	is.Equal([]int64{11, 20}, FilterInt64s([]int64{11, 20}, func(v int64) bool {
		return v > 0
	}))
	is.Equal([]int64{20}, FilterInt64s([]int64{11, 20}, func(v int64) bool {
		return v%2 == 0
	}))
}

func TestFilterFloat32s(t *testing.T) {
	is := assert.New(t)
	is.Equal([]float32{11, 20}, FilterFloat32s([]float32{11, 20}, func(v float32) bool {
		return v > 0
	}))
	is.Equal([]float32{20}, FilterFloat32s([]float32{11, 20}, func(v float32) bool {
		return v > 15
	}))
}

func TestFilterFloat64s(t *testing.T) {
	is := assert.New(t)
	is.Equal([]float64{11, 20}, FilterFloat64s([]float64{11, 20}, func(v float64) bool {
		return v > 0
	}))
	is.Equal([]float64{20}, FilterFloat64s([]float64{11, 20}, func(v float64) bool {
		return v > 15
	}))
}

func TestFilterStrings(t *testing.T) {
	is := assert.New(t)
	is.Equal([]string{}, FilterStrings([]string{"abc", "zora"}, func(v string) bool {
		return len(v) > 4
	}))
	is.Equal([]string{"zora"}, FilterStrings([]string{"abc", "zora"}, func(v string) bool {
		return len(v) > 3
	}))
}
