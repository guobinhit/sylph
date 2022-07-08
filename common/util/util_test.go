package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIf(t *testing.T) {
	is := assert.New(t)
	is.Equal(1, If(true, 1, 2))
	is.Equal(2, If(false, 1, 2))
}
