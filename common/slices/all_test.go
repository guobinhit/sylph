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
	is.Equal(false, AnyInt([]int{11, 20}, func(v int) bool {
		return v%3 == 0
	}))
}

func TestAllInt64(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AllInt64([]int64{11, 20}, func(v int64) bool {
		return v > 0
	}))
	is.Equal(false, AnyInt64([]int64{11, 20}, func(v int64) bool {
		return v%3 == 0
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
