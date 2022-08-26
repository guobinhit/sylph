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
			name: "DistinctInts 1",
			args: args{
				s: nil,
			},
			want: nil,
		},
		{
			name: "DistinctInts 2",
			args: args{
				s: []int{},
			},
			want: []int{},
		},
		{
			name: "DistinctInts 3",
			args: args{
				s: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "DistinctInts 4",
			args: args{
				s: []int{1, 2, 1},
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctInts(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctInt8s(t *testing.T) {
	type args struct {
		s []int8
	}
	tests := []struct {
		name string
		args args
		want []int8
	}{
		{
			name: "DistinctInt8s 1",
			args: args{
				s: nil,
			},
			want: nil,
		},
		{
			name: "DistinctInt8s 2",
			args: args{
				s: []int8{},
			},
			want: []int8{},
		},
		{
			name: "DistinctInt8s 3",
			args: args{
				s: []int8{1, 2, 3},
			},
			want: []int8{1, 2, 3},
		},
		{
			name: "DistinctInt8s 4",
			args: args{
				s: []int8{1, 2, 1},
			},
			want: []int8{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctInt8s(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctInt8s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctInt16s(t *testing.T) {
	type args struct {
		s []int16
	}
	tests := []struct {
		name string
		args args
		want []int16
	}{
		{
			name: "DistinctInt16s 1",
			args: args{
				s: nil,
			},
			want: nil,
		},
		{
			name: "DistinctInt16s 2",
			args: args{
				s: []int16{},
			},
			want: []int16{},
		},
		{
			name: "DistinctInt16s 3",
			args: args{
				s: []int16{1, 2, 3},
			},
			want: []int16{1, 2, 3},
		},
		{
			name: "DistinctInt16s 4",
			args: args{
				s: []int16{1, 2, 1},
			},
			want: []int16{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctInt16s(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctInt16s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctInt32s(t *testing.T) {
	type args struct {
		s []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "DistinctInt32s 1",
			args: args{
				s: nil,
			},
			want: nil,
		},
		{
			name: "DistinctInt32s 2",
			args: args{
				s: []int32{},
			},
			want: []int32{},
		},
		{
			name: "DistinctInt32s 3",
			args: args{
				s: []int32{1, 2, 3},
			},
			want: []int32{1, 2, 3},
		},
		{
			name: "DistinctInt32s 4",
			args: args{
				s: []int32{1, 2, 1},
			},
			want: []int32{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctInt32s(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctInt32s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctInt64s(t *testing.T) {
	type args struct {
		s []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "DistinctInt64s 1",
			args: args{
				s: nil,
			},
			want: nil,
		},
		{
			name: "DistinctInt64s 2",
			args: args{
				s: []int64{},
			},
			want: []int64{},
		},
		{
			name: "DistinctInt64s 3",
			args: args{
				s: []int64{1, 2, 3},
			},
			want: []int64{1, 2, 3},
		},
		{
			name: "DistinctInt64s 4",
			args: args{
				s: []int64{1, 2, 1},
			},
			want: []int64{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctInt64s(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctInt64s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctFloat32s(t *testing.T) {
	type args struct {
		s []float32
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{
			name: "DistinctFloat32s 1",
			args: args{
				s: nil,
			},
			want: nil,
		},
		{
			name: "DistinctFloat32s 2",
			args: args{
				s: []float32{},
			},
			want: []float32{},
		},
		{
			name: "DistinctFloat32s 3",
			args: args{
				s: []float32{1, 2, 3},
			},
			want: []float32{1, 2, 3},
		},
		{
			name: "DistinctFloat32s 4",
			args: args{
				s: []float32{1, 2, 1},
			},
			want: []float32{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctFloat32s(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctFloat32s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctFloat64s(t *testing.T) {
	type args struct {
		s []float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "DistinctFloat64s 1",
			args: args{
				s: nil,
			},
			want: nil,
		},
		{
			name: "DistinctFloat64s 2",
			args: args{
				s: []float64{},
			},
			want: []float64{},
		},
		{
			name: "DistinctFloat64s 3",
			args: args{
				s: []float64{1, 2, 3},
			},
			want: []float64{1, 2, 3},
		},
		{
			name: "DistinctFloat64s 4",
			args: args{
				s: []float64{1, 2, 1},
			},
			want: []float64{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctFloat64s(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctFloat64s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctStrings(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "DistinctStrings 1",
			args: args{
				s: nil,
			},
			want: nil,
		},
		{
			name: "DistinctStrings 2",
			args: args{
				s: []string{},
			},
			want: []string{},
		},
		{
			name: "DistinctStrings 3",
			args: args{
				s: []string{"1", "2", "3"},
			},
			want: []string{"1", "2", "3"},
		},
		{
			name: "DistinctStrings 4",
			args: args{
				s: []string{"1", "2", "1"},
			},
			want: []string{"1", "2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctStrings(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
