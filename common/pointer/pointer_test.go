package pointer

import (
	"reflect"
	"testing"
)

func TestInt(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want *int
	}{
		{
			name: "Int",
			args: args{
				i: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int(tt.args.i)
			if got == nil || reflect.ValueOf(got).Kind() != reflect.Ptr {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8(t *testing.T) {
	type args struct {
		i int8
	}
	tests := []struct {
		name string
		args args
		want *int8
	}{
		{
			name: "Int8",
			args: args{
				i: int8(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int8(tt.args.i)
			if got == nil || reflect.ValueOf(got).Kind() != reflect.Ptr {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16(t *testing.T) {
	type args struct {
		i int16
	}
	tests := []struct {
		name string
		args args
		want *int16
	}{
		{
			name: "Int16",
			args: args{
				i: int16(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int16(tt.args.i)
			if got == nil || reflect.ValueOf(got).Kind() != reflect.Ptr {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32(t *testing.T) {
	type args struct {
		i int32
	}
	tests := []struct {
		name string
		args args
		want *int32
	}{
		{
			name: "Int32",
			args: args{
				i: int32(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int32(tt.args.i)
			if got == nil || reflect.ValueOf(got).Kind() != reflect.Ptr {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	type args struct {
		i int64
	}
	tests := []struct {
		name string
		args args
		want *int64
	}{
		{
			name: "Int64",
			args: args{
				i: int64(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int64(tt.args.i)
			if got == nil || reflect.ValueOf(got).Kind() != reflect.Ptr {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32(t *testing.T) {
	type args struct {
		f float32
	}
	tests := []struct {
		name string
		args args
		want *float32
	}{
		{
			name: "Float32",
			args: args{
				f: float32(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Float32(tt.args.f)
			if got == nil || reflect.ValueOf(got).Kind() != reflect.Ptr {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		args args
		want *float64
	}{
		{
			name: "Float64",
			args: args{
				f: float64(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Float64(tt.args.f)
			if got == nil || reflect.ValueOf(got).Kind() != reflect.Ptr {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBool(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		args args
		want *bool
	}{
		{
			name: "Bool",
			args: args{
				b: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Bool(tt.args.b)
			if got == nil || reflect.ValueOf(got).Kind() != reflect.Ptr {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{
			name: "String",
			args: args{
				s: "abc",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := String(tt.args.s)
			if got == nil || reflect.ValueOf(got).Kind() != reflect.Ptr {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}
