package slices

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntContains(t *testing.T) {
	is := assert.New(t)
	is.Equal(false, IntContains(nil, 0))
	is.Equal(false, IntContains([]int{1, 2, 3}, 4))
	is.Equal(true, IntContains([]int{1, 2, 3}, 2))
}

func TestInt8Contains(t *testing.T) {
	is := assert.New(t)
	is.Equal(false, Int8Contains(nil, 0))
	is.Equal(false, Int8Contains([]int8{1, 2, 3}, 4))
	is.Equal(true, Int8Contains([]int8{1, 2, 3}, 2))
}

func TestInt16Contains(t *testing.T) {
	is := assert.New(t)
	is.Equal(false, Int16Contains(nil, 0))
	is.Equal(false, Int16Contains([]int16{1, 2, 3}, 4))
	is.Equal(true, Int16Contains([]int16{1, 2, 3}, 2))
}

func TestInt32Contains(t *testing.T) {
	is := assert.New(t)
	is.Equal(false, Int32Contains(nil, 0))
	is.Equal(false, Int32Contains([]int32{1, 2, 3}, 4))
	is.Equal(true, Int32Contains([]int32{1, 2, 3}, 2))
}

func TestInt64Contains(t *testing.T) {
	is := assert.New(t)
	is.Equal(false, Int64Contains(nil, 0))
	is.Equal(false, Int64Contains([]int64{1, 2, 3}, 4))
	is.Equal(true, Int64Contains([]int64{1, 2, 3}, 2))
}

func TestFloat32Contains(t *testing.T) {
	is := assert.New(t)
	is.Equal(false, Float32Contains(nil, 0))
	is.Equal(false, Float32Contains([]float32{1, 2, 3}, 4))
	is.Equal(true, Float32Contains([]float32{1, 2, 3}, 2))
}

func TestFloat64Contains(t *testing.T) {
	is := assert.New(t)
	is.Equal(false, Float64Contains(nil, 0))
	is.Equal(false, Float64Contains([]float64{1, 2, 3}, 4))
	is.Equal(true, Float64Contains([]float64{1, 2, 3}, 2))
}

func TestBoolContains(t *testing.T) {
	is := assert.New(t)
	is.Equal(false, BoolContains(nil, false))
	is.Equal(false, BoolContains([]bool{false, false, false}, true))
	is.Equal(true, BoolContains([]bool{false, false, false}, false))
}

func TestStringContains(t *testing.T) {
	is := assert.New(t)
	is.Equal(false, StringContains(nil, ""))
	is.Equal(false, StringContains([]string{"1", "a", "!"}, ""))
	is.Equal(false, StringContains([]string{"1", "a", "!"}, "nice"))
	is.Equal(true, StringContains([]string{"1", "a", "!"}, "a"))
}

func TestStringContainsIgnoreCase(t *testing.T) {
	is := assert.New(t)
	is.Equal(false, StringContainsIgnoreCase(nil, ""))
	is.Equal(false, StringContainsIgnoreCase([]string{"abC", "eFg", "IGk"}, ""))
	is.Equal(false, StringContainsIgnoreCase([]string{"abC", "eFg", "IGk"}, "ab"))
	is.Equal(true, StringContainsIgnoreCase([]string{"abC", "eFg", "IGk"}, "abc"))
	is.Equal(true, StringContainsIgnoreCase([]string{"abC", "eFg", "IGk"}, "IGK"))
}
