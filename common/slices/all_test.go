package slices

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAllInts(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AllInts([]int{11, 20}, func(v int) bool {
		return v > 0
	}))
	is.Equal(false, AllInts([]int{11, 20}, func(v int) bool {
		return v%3 == 0
	}))
}

func TestAllInt8s(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AllInt8s([]int8{11, 20}, func(v int8) bool {
		return v > 0
	}))
	is.Equal(false, AllInt8s([]int8{11, 20}, func(v int8) bool {
		return v%3 == 0
	}))
}

func TestAllInt16s(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AllInt16s([]int16{11, 20}, func(v int16) bool {
		return v > 0
	}))
	is.Equal(false, AllInt16s([]int16{11, 20}, func(v int16) bool {
		return v%3 == 0
	}))
}

func TestAllInt32s(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AllInt32s([]int32{11, 20}, func(v int32) bool {
		return v > 0
	}))
	is.Equal(false, AllInt32s([]int32{11, 20}, func(v int32) bool {
		return v%3 == 0
	}))
}

func TestAllInt64s(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AllInt64s([]int64{11, 20}, func(v int64) bool {
		return v > 0
	}))
	is.Equal(false, AllInt64s([]int64{11, 20}, func(v int64) bool {
		return v%3 == 0
	}))
}

func TestAllFloat32s(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AllFloat32s([]float32{11, 20}, func(v float32) bool {
		return v > 0
	}))
	is.Equal(false, AllFloat32s([]float32{11, 20}, func(v float32) bool {
		return v > 100
	}))
}

func TestAllFloat64s(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AllFloat64s([]float64{11, 20}, func(v float64) bool {
		return v > 0
	}))
	is.Equal(false, AllFloat64s([]float64{11, 20}, func(v float64) bool {
		return v > 100
	}))
}

func TestAllStrings(t *testing.T) {
	is := assert.New(t)
	is.Equal(false, AllStrings([]string{"abc", "zora"}, func(v string) bool {
		return len(v) > 10
	}))
	is.Equal(true, AllStrings([]string{"abc", "zora"}, func(v string) bool {
		return len(v) > 0
	}))
}
