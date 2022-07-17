package slices

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilterInt(t *testing.T) {
	is := assert.New(t)
	is.Equal([]int{11, 20}, FilterInt([]int{11, 20}, func(v int) bool {
		return v > 0
	}))
	is.Equal([]int{20}, FilterInt([]int{11, 20}, func(v int) bool {
		return v%2 == 0
	}))
}

func TestFilterInt8(t *testing.T) {
	is := assert.New(t)
	is.Equal([]int8{11, 20}, FilterInt8([]int8{11, 20}, func(v int8) bool {
		return v > 0
	}))
	is.Equal([]int8{20}, FilterInt8([]int8{11, 20}, func(v int8) bool {
		return v%2 == 0
	}))
}

func TestFilterInt16(t *testing.T) {
	is := assert.New(t)
	is.Equal([]int16{11, 20}, FilterInt16([]int16{11, 20}, func(v int16) bool {
		return v > 0
	}))
	is.Equal([]int16{20}, FilterInt16([]int16{11, 20}, func(v int16) bool {
		return v%2 == 0
	}))
}

func TestFilterInt32(t *testing.T) {
	is := assert.New(t)
	is.Equal([]int32{11, 20}, FilterInt32([]int32{11, 20}, func(v int32) bool {
		return v > 0
	}))
	is.Equal([]int32{20}, FilterInt32([]int32{11, 20}, func(v int32) bool {
		return v%2 == 0
	}))
}

func TestFilterInt64(t *testing.T) {
	is := assert.New(t)
	is.Equal([]int64{11, 20}, FilterInt64([]int64{11, 20}, func(v int64) bool {
		return v > 0
	}))
	is.Equal([]int64{20}, FilterInt64([]int64{11, 20}, func(v int64) bool {
		return v%2 == 0
	}))
}

func TestFilterFloat32(t *testing.T) {
	is := assert.New(t)
	is.Equal([]float32{11, 20}, FilterFloat32([]float32{11, 20}, func(v float32) bool {
		return v > 0
	}))
	is.Equal([]float32{20}, FilterFloat32([]float32{11, 20}, func(v float32) bool {
		return v > 15
	}))
}

func TestFilterFloat64(t *testing.T) {
	is := assert.New(t)
	is.Equal([]float64{11, 20}, FilterFloat64([]float64{11, 20}, func(v float64) bool {
		return v > 0
	}))
	is.Equal([]float64{20}, FilterFloat64([]float64{11, 20}, func(v float64) bool {
		return v > 15
	}))
}

func TestFilterString(t *testing.T) {
	is := assert.New(t)
	is.Equal([]string{}, FilterString([]string{"abc", "zora"}, func(v string) bool {
		return len(v) > 4
	}))
	is.Equal([]string{"zora"}, FilterString([]string{"abc", "zora"}, func(v string) bool {
		return len(v) > 3
	}))
}
