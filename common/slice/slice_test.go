package slice

import "testing"

func TestIntContains(t *testing.T) {
	type args struct {
		s []int
		e int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "IntContains false",
			args: args{
				s: []int{1, 2, 3},
				e: 4,
			},
			want: false,
		},
		{
			name: "IntContains true",
			args: args{
				s: []int{1, 2, 3},
				e: 2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntContains(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("IntContains() = %v, want %v", got, tt.want)
			}
		})
	}
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
