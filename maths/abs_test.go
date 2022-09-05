package maths

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestAbsInt8(t *testing.T) {
	type args struct {
		v int8
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(math.MinInt8)
			assert.Equalf(t, tt.want, AbsInt8(tt.args.v), "AbsInt8(%v)", tt.args.v)
		})
	}
}
