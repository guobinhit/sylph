package slices

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainInts(t *testing.T) {
	is := assert.New(t)
	is.Equal(false, ContainInts(nil, 0))
	is.Equal(false, ContainInts([]int{1, 2, 3}, 4))
	is.Equal(true, ContainInts([]int{1, 2, 3}, 2))
}

func TestContainInt8s(t *testing.T) {
	is := assert.New(t)
	is.Equal(false, ContainInt8s(nil, 0))
	is.Equal(false, ContainInt8s([]int8{1, 2, 3}, 4))
	is.Equal(true, ContainInt8s([]int8{1, 2, 3}, 2))
}

func TestContainInt16s(t *testing.T) {
	is := assert.New(t)
	is.Equal(false, ContainInt16s(nil, 0))
	is.Equal(false, ContainInt16s([]int16{1, 2, 3}, 4))
	is.Equal(true, ContainInt16s([]int16{1, 2, 3}, 2))
}

func TestContainInt32s(t *testing.T) {
	is := assert.New(t)
	is.Equal(false, ContainInt32s(nil, 0))
	is.Equal(false, ContainInt32s([]int32{1, 2, 3}, 4))
	is.Equal(true, ContainInt32s([]int32{1, 2, 3}, 2))
}

func TestContainInt64s(t *testing.T) {
	is := assert.New(t)
	is.Equal(false, ContainInt64s(nil, 0))
	is.Equal(false, ContainInt64s([]int64{1, 2, 3}, 4))
	is.Equal(true, ContainInt64s([]int64{1, 2, 3}, 2))
}

func TestContainFloat32s(t *testing.T) {
	is := assert.New(t)
	is.Equal(false, ContainFloat32s(nil, 0))
	is.Equal(false, ContainFloat32s([]float32{1, 2, 3}, 4))
	is.Equal(true, ContainFloat32s([]float32{1, 2, 3}, 2))
}

func TestContainFloat64s(t *testing.T) {
	is := assert.New(t)
	is.Equal(false, ContainFloat64s(nil, 0))
	is.Equal(false, ContainFloat64s([]float64{1, 2, 3}, 4))
	is.Equal(true, ContainFloat64s([]float64{1, 2, 3}, 2))
}

func TestContainStrings(t *testing.T) {
	is := assert.New(t)
	is.Equal(false, ContainStrings(nil, ""))
	is.Equal(false, ContainStrings([]string{"1", "a", "!"}, ""))
	is.Equal(false, ContainStrings([]string{"1", "a", "!"}, "nice"))
	is.Equal(true, ContainStrings([]string{"1", "a", "!"}, "a"))
}

func TestContainStringsIgnoreCase(t *testing.T) {
	is := assert.New(t)
	is.Equal(false, ContainStringsIgnoreCase(nil, ""))
	is.Equal(false, ContainStringsIgnoreCase([]string{"abC", "eFg", "IGk"}, ""))
	is.Equal(false, ContainStringsIgnoreCase([]string{"abC", "eFg", "IGk"}, "ab"))
	is.Equal(true, ContainStringsIgnoreCase([]string{"abC", "eFg", "IGk"}, "abc"))
	is.Equal(true, ContainStringsIgnoreCase([]string{"abC", "eFg", "IGk"}, "IGK"))
}
