package format

import (
	"github.com/guobinhit/sylph/constant/date_const"
	"testing"
	"time"
)

func TestGetCptOfYyyyMm(t *testing.T) {
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-30 10:00:00", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "GetCptOfYyyyMm",
			args: args{d: d},
			want: "202204",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCptOfYyyyMm(tt.args.d); got != tt.want {
				t.Errorf("GetCptOfYyyyMm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCptOfYyyyMmDd(t *testing.T) {
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-01 10:00:00", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "GetCptOfYyyyMmDd",
			args: args{d: d},
			want: "20220401",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCptOfYyyyMmDd(tt.args.d); got != tt.want {
				t.Errorf("GetCptOfYyyyMmDd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCptOfYyyyMmDdHh(t *testing.T) {
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-13 23:59:00", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "GetCptOfYyyyMmDdHh",
			args: args{d: d},
			want: "2022041323",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCptOfYyyyMmDdHh(tt.args.d); got != tt.want {
				t.Errorf("GetCptOfYyyyMmDdHh() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCptOfYyyyMmDdHhMm(t *testing.T) {
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSsSss, "2022-04-13 10:59:00.999", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "GetCptOfYyyyMmDdHhMm",
			args: args{d: d},
			want: "202204131059",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCptOfYyyyMmDdHhMm(tt.args.d); got != tt.want {
				t.Errorf("GetCptOfYyyyMmDdHhMm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCptOfYyyyMmDdHhMmSs(t *testing.T) {
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSsSss, "2022-04-13 13:20:01.123", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "GetCptOfYyyyMmDdHhMmSs",
			args: args{d: d},
			want: "20220413132001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCptOfYyyyMmDdHhMmSs(tt.args.d); got != tt.want {
				t.Errorf("GetCptOfYyyyMmDdHhMmSs() = %v, want %v", got, tt.want)
			}
		})
	}
}
