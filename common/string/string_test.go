package string

import "testing"

func TestEquals(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Equals false 1",
			args: args{
				s1: "abc",
				s2: "",
			},
			want: false,
		},
		{
			name: "Equals false 2",
			args: args{
				s1: "abc",
				s2: "bcd",
			},
			want: false,
		},
		{
			name: "Equals false 3",
			args: args{
				s1: "abc",
				s2: "aBc",
			},
			want: false,
		},
		{
			name: "Equals true 1",
			args: args{
				s1: "abc",
				s2: "abc",
			},
			want: true,
		},
		{
			name: "Equals true 2",
			args: args{
				s1: "a&b!?c",
				s2: "a&b!?c",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equals(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqualsIgnoreCase(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "EqualsIgnoreCase false 1",
			args: args{
				s1: "abc",
				s2: "",
			},
			want: false,
		},
		{
			name: "EqualsIgnoreCase false 2",
			args: args{
				s1: "abc",
				s2: "bcd",
			},
			want: false,
		},
		{
			name: "EqualsIgnoreCase false 3",
			args: args{
				s1: "welcome",
				s2: "come",
			},
			want: false,
		},
		{
			name: "EqualsIgnoreCase true 1",
			args: args{
				s1: "abc",
				s2: "abc",
			},
			want: true,
		},
		{
			name: "EqualsIgnoreCase true 2",
			args: args{
				s1: "abc",
				s2: "AbC",
			},
			want: true,
		},
		{
			name: "EqualsIgnoreCase true 3",
			args: args{
				s1: "a&b!?c",
				s2: "a&b!?c",
			},
			want: true,
		},
		{
			name: "EqualsIgnoreCase true 4",
			args: args{
				s1: "a&b!?c",
				s2: "A&B!?c",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EqualsIgnoreCase(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("EqualsIgnoreCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
