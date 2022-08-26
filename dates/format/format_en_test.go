package format

import (
	"github.com/guobinhit/sylph/constant/date_const"
	"testing"
	"time"
)

func TestGetEnOfMmDd(t *testing.T) {
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-13 10:00:00", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "GetEnOfMmDd",
			args: args{d: d},
			want: "04/13",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnOfMmDd(tt.args.d); got != tt.want {
				t.Errorf("GetEnOfMmDd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnOfYyyyMm(t *testing.T) {
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-13 10:00:00", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "GetEnOfYyyyMm",
			args: args{d: d},
			want: "2022/04",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnOfYyyyMm(tt.args.d); got != tt.want {
				t.Errorf("GetEnOfYyyyMm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnOfYyyyMmDd(t *testing.T) {
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-13 10:00:00", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "GetEnOfYyyyMmDd",
			args: args{d: d},
			want: "2022/04/13",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnOfYyyyMmDd(tt.args.d); got != tt.want {
				t.Errorf("GetEnOfYyyyMmDd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnOfMmDdHhMm(t *testing.T) {
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-13 10:00:00", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "GetEnOfMmDdHhMm",
			args: args{d: d},
			want: "04/13 10:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnOfMmDdHhMm(tt.args.d); got != tt.want {
				t.Errorf("GetEnOfMmDdHhMm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnOfMmDdHhMmSs(t *testing.T) {
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-13 10:00:00", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "GetEnOfMmDdHhMmSs",
			args: args{d: d},
			want: "04/13 10:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnOfMmDdHhMmSs(tt.args.d); got != tt.want {
				t.Errorf("GetEnOfMmDdHhMmSs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnOfYyyyMmDdHhMm(t *testing.T) {
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSsSss, "2022-04-13 10:00:00.999", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "GetEnOfYyyyMmDdHhMm",
			args: args{d: d},
			want: "2022/04/13 10:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnOfYyyyMmDdHhMm(tt.args.d); got != tt.want {
				t.Errorf("GetEnOfYyyyMmDdHhMm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnOfYyyyMmDdHhMmSs(t *testing.T) {
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSsSss, "2022-04-13 10:00:00.999", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "GetEnOfYyyyMmDdHhMmSs",
			args: args{d: d},
			want: "2022/04/13 10:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnOfYyyyMmDdHhMmSs(tt.args.d); got != tt.want {
				t.Errorf("GetEnOfYyyyMmDdHhMmSs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnOfYyyyMmDdHhMmSsSss(t *testing.T) {
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSsSss, "2022-04-13 10:20:59.999", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "GetEnOfYyyyMmDdHhMmSsSss",
			args: args{d: d},
			want: "2022/04/13 10:20:59.999",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnOfYyyyMmDdHhMmSsSss(tt.args.d); got != tt.want {
				t.Errorf("GetEnOfYyyyMmDdHhMmSsSss() = %v, want %v", got, tt.want)
			}
		})
	}
}
