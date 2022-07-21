package format

import (
	"github.com/guobinhit/sylph/constant/date_const"
	"testing"
	"time"
)

func TestGetMmDd(t *testing.T) {
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
			name: "GetMmDd",
			args: args{d: d},
			want: "04-13",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMmDd(tt.args.d); got != tt.want {
				t.Errorf("GetMmDd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetYyyyMm(t *testing.T) {
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
			name: "GetYyyyMm",
			args: args{d: d},
			want: "2022-04",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetYyyyMm(tt.args.d); got != tt.want {
				t.Errorf("GetYyyyMm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetYyyyMmDd(t *testing.T) {
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
			name: "GetYyyyMmDd",
			args: args{d: d},
			want: "2022-04-13",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetYyyyMmDd(tt.args.d); got != tt.want {
				t.Errorf("GetYyyyMmDd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMmDdHhMm(t *testing.T) {
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
			name: "GetMmDdHhMm",
			args: args{d: d},
			want: "04-13 10:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMmDdHhMm(tt.args.d); got != tt.want {
				t.Errorf("GetMmDdHhMm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMmDdHhMmSs(t *testing.T) {
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
			name: "GetMmDdHhMmSs",
			args: args{d: d},
			want: "04-13 10:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMmDdHhMmSs(tt.args.d); got != tt.want {
				t.Errorf("GetMmDdHhMmSs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMmDdHhMmSsSss(t *testing.T) {
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
			name: "GetMmDdHhMmSsSss",
			args: args{d: d},
			want: "04-13 10:00:00.999",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMmDdHhMmSsSss(tt.args.d); got != tt.want {
				t.Errorf("GetMmDdHhMmSsSss() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetYyyyMmDdHhMm(t *testing.T) {
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
			name: "GetYyyyMmDdHhMm",
			args: args{d: d},
			want: "2022-04-13 10:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetYyyyMmDdHhMm(tt.args.d); got != tt.want {
				t.Errorf("GetYyyyMmDdHhMm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetYyyyMmDdHhMmSs(t *testing.T) {
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
			name: "GetYyyyMmDdHhMmSs",
			args: args{d: d},
			want: "2022-04-13 10:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetYyyyMmDdHhMmSs(tt.args.d); got != tt.want {
				t.Errorf("GetYyyyMmDdHhMmSs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetYyyyMmDdHhMmSsSss(t *testing.T) {
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
			name: "GetYyyyMmDdHhMmSsSss",
			args: args{d: d},
			want: "2022-04-13 10:00:00.999",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetYyyyMmDdHhMmSsSss(tt.args.d); got != tt.want {
				t.Errorf("GetYyyyMmDdHhMmSsSss() = %v, want %v", got, tt.want)
			}
		})
	}
}
