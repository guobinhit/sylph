package slices

import (
	"reflect"
	"testing"
)

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
