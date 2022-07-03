package slice

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestIntContains(t *testing.T) {
	is := assert.New(t)
	is.Equal(false, IntContains(nil, 0))
	is.Equal(false, IntContains([]int{1, 2, 3}, 4))
	is.Equal(true, IntContains([]int{1, 2, 3}, 2))
}

func TestInt8Contains(t *testing.T) {
	type args struct {
		s []int8
		e int8
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Int8Contains false",
			args: args{
				s: []int8{1, 2, 3},
				e: 4,
			},
			want: false,
		},
		{
			name: "Int8Contains true",
			args: args{
				s: []int8{1, 2, 3},
				e: 2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int8Contains(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("Int8Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16Contains(t *testing.T) {
	type args struct {
		s []int16
		e int16
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Int16Contains false",
			args: args{
				s: []int16{1, 2, 3},
				e: 4,
			},
			want: false,
		},
		{
			name: "Int16Contains true",
			args: args{
				s: []int16{1, 2, 3},
				e: 2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int16Contains(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("Int16Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32Contains(t *testing.T) {
	type args struct {
		s []int32
		e int32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Int32Contains false",
			args: args{
				s: []int32{1, 2, 3},
				e: 4,
			},
			want: false,
		},
		{
			name: "Int32Contains true",
			args: args{
				s: []int32{1, 2, 3},
				e: 2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32Contains(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("Int32Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64Contains(t *testing.T) {
	type args struct {
		s []int64
		e int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Int64Contains false",
			args: args{
				s: []int64{1, 2, 3},
				e: 4,
			},
			want: false,
		},
		{
			name: "Int64Contains true",
			args: args{
				s: []int64{1, 2, 3},
				e: 2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64Contains(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("Int64Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32Contains(t *testing.T) {
	type args struct {
		s []float32
		e float32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Float32Contains false",
			args: args{
				s: []float32{1, 2, 3},
				e: 4,
			},
			want: false,
		},
		{
			name: "Float32Contains true",
			args: args{
				s: []float32{1, 2, 3},
				e: 2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32Contains(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("Float32Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64Contains(t *testing.T) {
	type args struct {
		s []float64
		e float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Float64Contains false",
			args: args{
				s: []float64{1, 2, 3},
				e: 4,
			},
			want: false,
		},
		{
			name: "Float64Contains true",
			args: args{
				s: []float64{1, 2, 3},
				e: 2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64Contains(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("Float64Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolContains(t *testing.T) {
	type args struct {
		s []bool
		e bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "BoolContains false",
			args: args{
				s: []bool{false, false, false},
				e: true,
			},
			want: false,
		},
		{
			name: "BoolContains true",
			args: args{
				s: []bool{false, true, false},
				e: true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoolContains(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("BoolContains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringContains(t *testing.T) {
	type args struct {
		s []string
		e string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "StringContains false 1",
			args: args{
				s: []string{"1", "a", "!"},
				e: "",
			},
			want: false,
		},
		{
			name: "StringContains false 2",
			args: args{
				s: []string{"1", "a", "!"},
				e: "nice",
			},
			want: false,
		},
		{
			name: "StringContains true",
			args: args{
				s: []string{"1", "a", "!"},
				e: "a",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringContains(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("StringContains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringContainsIgnoreCase(t *testing.T) {
	type args struct {
		s []string
		e string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "StringContainsIgnoreCase false 1",
			args: args{
				s: []string{"abC", "eFg", "IGk"},
				e: "",
			},
			want: false,
		},
		{
			name: "StringContainsIgnoreCase false 2",
			args: args{
				s: []string{"abC", "eFg", "IGk"},
				e: "ab",
			},
			want: false,
		},
		{
			name: "StringContainsIgnoreCase true 1",
			args: args{
				s: []string{"abC", "eFg", "IGk"},
				e: "abc",
			},
			want: true,
		},
		{
			name: "StringContainsIgnoreCase true 2",
			args: args{
				s: []string{"abC", "eFg", "IGk"},
				e: "IGK",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringContainsIgnoreCase(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("StringContainsIgnoreCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctSliceInt(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "DistinctSliceInt 1",
			args: args{
				s: nil,
			},
			want: nil,
		},
		{
			name: "DistinctSliceInt 2",
			args: args{
				s: []int{},
			},
			want: []int{},
		},
		{
			name: "DistinctSliceInt 3",
			args: args{
				s: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "DistinctSliceInt 4",
			args: args{
				s: []int{1, 2, 1},
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctSliceInt(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctSliceInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctSliceInt8(t *testing.T) {
	type args struct {
		s []int8
	}
	tests := []struct {
		name string
		args args
		want []int8
	}{
		{
			name: "DistinctSliceInt8 1",
			args: args{
				s: nil,
			},
			want: nil,
		},
		{
			name: "DistinctSliceInt8 2",
			args: args{
				s: []int8{},
			},
			want: []int8{},
		},
		{
			name: "DistinctSliceInt8 3",
			args: args{
				s: []int8{1, 2, 3},
			},
			want: []int8{1, 2, 3},
		},
		{
			name: "DistinctSliceInt8 4",
			args: args{
				s: []int8{1, 2, 1},
			},
			want: []int8{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctSliceInt8(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctSliceInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctSliceInt16(t *testing.T) {
	type args struct {
		s []int16
	}
	tests := []struct {
		name string
		args args
		want []int16
	}{
		{
			name: "DistinctSliceInt16 1",
			args: args{
				s: nil,
			},
			want: nil,
		},
		{
			name: "DistinctSliceInt16 2",
			args: args{
				s: []int16{},
			},
			want: []int16{},
		},
		{
			name: "DistinctSliceInt16 3",
			args: args{
				s: []int16{1, 2, 3},
			},
			want: []int16{1, 2, 3},
		},
		{
			name: "DistinctSliceInt16 4",
			args: args{
				s: []int16{1, 2, 1},
			},
			want: []int16{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctSliceInt16(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctSliceInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctSliceInt32(t *testing.T) {
	type args struct {
		s []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "DistinctSliceInt32 1",
			args: args{
				s: nil,
			},
			want: nil,
		},
		{
			name: "DistinctSliceInt32 2",
			args: args{
				s: []int32{},
			},
			want: []int32{},
		},
		{
			name: "DistinctSliceInt32 3",
			args: args{
				s: []int32{1, 2, 3},
			},
			want: []int32{1, 2, 3},
		},
		{
			name: "DistinctSliceInt32 4",
			args: args{
				s: []int32{1, 2, 1},
			},
			want: []int32{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctSliceInt32(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctSliceInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctSliceInt64(t *testing.T) {
	type args struct {
		s []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "DistinctSliceInt64 1",
			args: args{
				s: nil,
			},
			want: nil,
		},
		{
			name: "DistinctSliceInt64 2",
			args: args{
				s: []int64{},
			},
			want: []int64{},
		},
		{
			name: "DistinctSliceInt64 3",
			args: args{
				s: []int64{1, 2, 3},
			},
			want: []int64{1, 2, 3},
		},
		{
			name: "DistinctSliceInt64 4",
			args: args{
				s: []int64{1, 2, 1},
			},
			want: []int64{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctSliceInt64(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctSliceInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctSliceFloat32(t *testing.T) {
	type args struct {
		s []float32
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{
			name: "DistinctSliceFloat32 1",
			args: args{
				s: nil,
			},
			want: nil,
		},
		{
			name: "DistinctSliceFloat32 2",
			args: args{
				s: []float32{},
			},
			want: []float32{},
		},
		{
			name: "DistinctSliceFloat32 3",
			args: args{
				s: []float32{1, 2, 3},
			},
			want: []float32{1, 2, 3},
		},
		{
			name: "DistinctSliceFloat32 4",
			args: args{
				s: []float32{1, 2, 1},
			},
			want: []float32{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctSliceFloat32(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctSliceFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctSliceFloat64(t *testing.T) {
	type args struct {
		s []float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "DistinctSliceFloat64 1",
			args: args{
				s: nil,
			},
			want: nil,
		},
		{
			name: "DistinctSliceFloat64 2",
			args: args{
				s: []float64{},
			},
			want: []float64{},
		},
		{
			name: "DistinctSliceFloat64 3",
			args: args{
				s: []float64{1, 2, 3},
			},
			want: []float64{1, 2, 3},
		},
		{
			name: "DistinctSliceFloat64 4",
			args: args{
				s: []float64{1, 2, 1},
			},
			want: []float64{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctSliceFloat64(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctSliceFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctSliceString(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "DistinctSliceString 1",
			args: args{
				s: nil,
			},
			want: nil,
		},
		{
			name: "DistinctSliceString 2",
			args: args{
				s: []string{},
			},
			want: []string{},
		},
		{
			name: "DistinctSliceString 3",
			args: args{
				s: []string{"1", "2", "3"},
			},
			want: []string{"1", "2", "3"},
		},
		{
			name: "DistinctSliceString 4",
			args: args{
				s: []string{"1", "2", "1"},
			},
			want: []string{"1", "2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctSliceString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctSliceString() = %v, want %v", got, tt.want)
			}
		})
	}
}
