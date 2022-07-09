package slice

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

func TestAnyInt64(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, AnyInt64([]int64{11, 20}, func(v int64) bool {
		return v%2 == 0
	}))
	is.Equal(false, AnyInt64([]int64{11, 20}, func(v int64) bool {
		return v%3 == 0
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
