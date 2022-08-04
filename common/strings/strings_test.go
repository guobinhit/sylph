package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEquals(t *testing.T) {
	is := assert.New(t)
	is.Equal(false, Equals("abc", ""))
	is.Equal(false, Equals("abc", "bcd"))
	is.Equal(false, Equals("abc", "aBc"))
	is.Equal(true, Equals("abc", "abc"))
	is.Equal(true, Equals("a&b!?c", "a&b!?c"))
}

func TestEqualsIgnoreCase(t *testing.T) {
	is := assert.New(t)
	is.Equal(false, EqualsIgnoreCase("abc", ""))
	is.Equal(false, EqualsIgnoreCase("abc", "bcd"))
	is.Equal(true, EqualsIgnoreCase("abc", "aBc"))
	is.Equal(true, EqualsIgnoreCase("abc", "abc"))
	is.Equal(true, EqualsIgnoreCase("a&b!?c", "a&b!?c"))
	is.Equal(true, EqualsIgnoreCase("a&b!?c", "A&b!?C"))
}

func TestStarsWith(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, StarsWith("", ""))
	is.Equal(true, StarsWith("abc", ""))
	is.Equal(true, StarsWith("abc", "a"))
	is.Equal(true, StarsWith("abc", "ab"))
	is.Equal(true, StarsWith("abc", "abc"))
	is.Equal(false, StarsWith("abc", "abcd"))
	is.Equal(false, StarsWith("abc", "bc"))
}

func TestEndsWith(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, EndsWith("", ""))
	is.Equal(true, EndsWith("abc", ""))
	is.Equal(false, EndsWith("abc", "a"))
	is.Equal(false, EndsWith("abc", "ab"))
	is.Equal(true, EndsWith("abc", "abc"))
	is.Equal(false, EndsWith("abc", "b"))
	is.Equal(true, EndsWith("abc", "bc"))
}
