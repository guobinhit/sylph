package unpointer

import (
	"github.com/guobinhit/sylph/common/pointer"
	"testing"
)

func TestInt(t *testing.T) {
	type args struct {
		i *int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Int 1",
			args: args{
				i: pointer.Int(123),
			},
			want: int(123),
		},
		{
			name: "Int 2",
			args: args{
				i: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.args.i); got != tt.want {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8(t *testing.T) {
	type args struct {
		i *int8
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{
			name: "Int8 1",
			args: args{
				i: pointer.Int8(8),
			},
			want: int8(8),
		},
		{
			name: "Int8 2",
			args: args{
				i: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int8(tt.args.i); got != tt.want {
				t.Errorf("Int8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16(t *testing.T) {
	type args struct {
		i *int16
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{
			name: "Int16 1",
			args: args{
				i: pointer.Int16(123),
			},
			want: int16(123),
		},
		{
			name: "Int16 2",
			args: args{
				i: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int16(tt.args.i); got != tt.want {
				t.Errorf("Int16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32(t *testing.T) {
	type args struct {
		i *int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "Int32 1",
			args: args{
				i: pointer.Int32(123),
			},
			want: int32(123),
		},
		{
			name: "Int32 2",
			args: args{
				i: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32(tt.args.i); got != tt.want {
				t.Errorf("Int32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	type args struct {
		i *int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Int64 1",
			args: args{
				i: pointer.Int64(123),
			},
			want: int64(123),
		},
		{
			name: "Int64 2",
			args: args{
				i: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64(tt.args.i); got != tt.want {
				t.Errorf("Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32(t *testing.T) {
	type args struct {
		f *float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "Float32 1",
			args: args{
				f: pointer.Float32(123),
			},
			want: float32(123),
		},
		{
			name: "Float32 2",
			args: args{
				f: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32(tt.args.f); got != tt.want {
				t.Errorf("Float32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	type args struct {
		f *float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Float64 1",
			args: args{
				f: pointer.Float64(123),
			},
			want: float64(123),
		},
		{
			name: "Float64 2",
			args: args{
				f: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64(tt.args.f); got != tt.want {
				t.Errorf("Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBool(t *testing.T) {
	type args struct {
		b *bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Bool 1",
			args: args{
				b: pointer.Bool(true),
			},
			want: true,
		},
		{
			name: "Bool 2",
			args: args{
				b: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bool(tt.args.b); got != tt.want {
				t.Errorf("Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	type args struct {
		s *string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "String 1",
			args: args{
				s: pointer.String("zora"),
			},
			want: "zora",
		},
		{
			name: "String 2",
			args: args{
				s: nil,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.s); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntOrDefault(t *testing.T) {
	type args struct {
		i  *int
		dv int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "IntOrDefault 1",
			args: args{
				i: pointer.Int(123),
			},
			want: int(123),
		},
		{
			name: "IntOrDefault 2",
			args: args{
				i:  nil,
				dv: 123,
			},
			want: 123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntOrDefault(tt.args.i, tt.args.dv); got != tt.want {
				t.Errorf("IntOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8OrDefault(t *testing.T) {
	type args struct {
		i  *int8
		dv int8
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{
			name: "Int8OrDefault 1",
			args: args{
				i: pointer.Int8(8),
			},
			want: int8(8),
		},
		{
			name: "Int8OrDefault 2",
			args: args{
				i:  nil,
				dv: int8(8),
			},
			want: int8(8),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int8OrDefault(tt.args.i, tt.args.dv); got != tt.want {
				t.Errorf("Int8OrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16OrDefault(t *testing.T) {
	type args struct {
		i  *int16
		dv int16
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{
			name: "Int16OrDefault 1",
			args: args{
				i: pointer.Int16(123),
			},
			want: int16(123),
		},
		{
			name: "Int16OrDefault 2",
			args: args{
				i:  nil,
				dv: int16(123),
			},
			want: int16(123),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int16OrDefault(tt.args.i, tt.args.dv); got != tt.want {
				t.Errorf("Int16OrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32OrDefault(t *testing.T) {
	type args struct {
		i  *int32
		dv int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "Int32OrDefault 1",
			args: args{
				i: pointer.Int32(123),
			},
			want: int32(123),
		},
		{
			name: "Int32OrDefault 2",
			args: args{
				i:  nil,
				dv: int32(123),
			},
			want: int32(123),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32OrDefault(tt.args.i, tt.args.dv); got != tt.want {
				t.Errorf("Int32OrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64OrDefault(t *testing.T) {
	type args struct {
		i  *int64
		dv int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Int64OrDefault 1",
			args: args{
				i: pointer.Int64(123),
			},
			want: int64(123),
		},
		{
			name: "Int64OrDefault 2",
			args: args{
				i:  nil,
				dv: int64(123),
			},
			want: int64(123),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64OrDefault(tt.args.i, tt.args.dv); got != tt.want {
				t.Errorf("Int64OrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32OrDefault(t *testing.T) {
	type args struct {
		f  *float32
		dv float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "Float32OrDefault 1",
			args: args{
				f: pointer.Float32(123),
			},
			want: float32(123),
		},
		{
			name: "Float32OrDefault 2",
			args: args{
				f:  nil,
				dv: float32(123),
			},
			want: float32(123),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32OrDefault(tt.args.f, tt.args.dv); got != tt.want {
				t.Errorf("Float32OrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64OrDefault(t *testing.T) {
	type args struct {
		f  *float64
		dv float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Float64OrDefault 1",
			args: args{
				f: pointer.Float64(123),
			},
			want: float64(123),
		},
		{
			name: "Float64OrDefault 2",
			args: args{
				f:  nil,
				dv: float64(123),
			},
			want: float64(123),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64OrDefault(tt.args.f, tt.args.dv); got != tt.want {
				t.Errorf("Float64OrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolOrDefault(t *testing.T) {
	type args struct {
		b  *bool
		dv bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "BoolOrDefault 1",
			args: args{
				b: pointer.Bool(true),
			},
			want: true,
		},
		{
			name: "BoolOrDefault 2",
			args: args{
				b:  nil,
				dv: true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoolOrDefault(tt.args.b, tt.args.dv); got != tt.want {
				t.Errorf("BoolOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringOrDefault(t *testing.T) {
	type args struct {
		s  *string
		dv string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "StringOrDefault 1",
			args: args{
				s: pointer.String("zora"),
			},
			want: "zora",
		},
		{
			name: "StringOrDefault 2",
			args: args{
				s:  nil,
				dv: "zora",
			},
			want: "zora",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringOrDefault(tt.args.s, tt.args.dv); got != tt.want {
				t.Errorf("StringOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
