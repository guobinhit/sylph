package util

import (
	"reflect"
	"testing"
)

func TestIf(t *testing.T) {
	type args struct {
		b bool
		f interface{}
		s interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "If false",
			args: args{
				b: false,
				f: 1,
				s: 2,
			},
			want: 2,
		},
		{
			name: "If true",
			args: args{
				b: false,
				f: "hello",
				s: "world",
			},
			want: "world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := If(tt.args.b, tt.args.f, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("If() = %v, want %v", got, tt.want)
			}
		})
	}
}
