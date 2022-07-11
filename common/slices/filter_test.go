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

func TestFilterInt64(t *testing.T) {
	is := assert.New(t)
	is.Equal([]int64{11, 20}, FilterInt64([]int64{11, 20}, func(v int64) bool {
		return v > 0
	}))
	is.Equal([]int64{20}, FilterInt64([]int64{11, 20}, func(v int64) bool {
		return v%2 == 0
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
