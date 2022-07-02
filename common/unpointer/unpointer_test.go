package unpointer

import (
	"github.com/guobinhit/sylph/common/pointer"
	"testing"
)

func TestIntOrDefault(t *testing.T) {
	type args struct {
		i *int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "IntOrDefault 1",
			args: args{
				i: pointer.Int(1120),
			},
			want: int(1120),
		},
		{
			name: "IntOrDefault 2",
			args: args{
				i: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntOrDefault(tt.args.i); got != tt.want {
				t.Errorf("IntOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8OrDefault(t *testing.T) {
	type args struct {
		i *int8
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
				i: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int8OrDefault(tt.args.i); got != tt.want {
				t.Errorf("Int8OrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16OrDefault(t *testing.T) {
	type args struct {
		i *int16
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{
			name: "Int16OrDefault 1",
			args: args{
				i: pointer.Int16(1120),
			},
			want: int16(1120),
		},
		{
			name: "Int16OrDefault 2",
			args: args{
				i: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int16OrDefault(tt.args.i); got != tt.want {
				t.Errorf("Int16OrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32OrDefault(t *testing.T) {
	type args struct {
		i *int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "Int32OrDefault 1",
			args: args{
				i: pointer.Int32(1120),
			},
			want: int32(1120),
		},
		{
			name: "Int32OrDefault 2",
			args: args{
				i: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32OrDefault(tt.args.i); got != tt.want {
				t.Errorf("Int32OrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64OrDefault(t *testing.T) {
	type args struct {
		i *int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Int64OrDefault 1",
			args: args{
				i: pointer.Int64(1120),
			},
			want: int64(1120),
		},
		{
			name: "Int64OrDefault 2",
			args: args{
				i: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64OrDefault(tt.args.i); got != tt.want {
				t.Errorf("Int64OrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32OrDefault(t *testing.T) {
	type args struct {
		f *float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "Float32OrDefault 1",
			args: args{
				f: pointer.Float32(1120),
			},
			want: float32(1120),
		},
		{
			name: "Float32OrDefault 2",
			args: args{
				f: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32OrDefault(tt.args.f); got != tt.want {
				t.Errorf("Float32OrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64OrDefault(t *testing.T) {
	type args struct {
		f *float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Float64OrDefault 1",
			args: args{
				f: pointer.Float64(1120),
			},
			want: float64(1120),
		},
		{
			name: "Float64OrDefault 2",
			args: args{
				f: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64OrDefault(tt.args.f); got != tt.want {
				t.Errorf("Float64OrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolOrDefault(t *testing.T) {
	type args struct {
		b *bool
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
				b: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoolOrDefault(tt.args.b); got != tt.want {
				t.Errorf("BoolOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringOrDefault(t *testing.T) {
	type args struct {
		s *string
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
				s: nil,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringOrDefault(tt.args.s); got != tt.want {
				t.Errorf("StringOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
