package format

import (
	"github.com/guobinhit/sylph/constant/date_const"
	"testing"
	"time"
)

func TestGetCnOfMmDd(t *testing.T) {
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
			name: "GetCnOfMmDd",
			args: args{d: d},
			want: "04月13日",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCnOfMmDd(tt.args.d); got != tt.want {
				t.Errorf("GetCnOfMmDd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCnOfYyyyMm(t *testing.T) {
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
			name: "GetCnOfYyyyMm",
			args: args{d: d},
			want: "2022年04月",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCnOfYyyyMm(tt.args.d); got != tt.want {
				t.Errorf("GetCnOfYyyyMm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCnOfYyyyMmDd(t *testing.T) {
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
			name: "GetCnOfYyyyMmDd",
			args: args{d: d},
			want: "2022年04月13日",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCnOfYyyyMmDd(tt.args.d); got != tt.want {
				t.Errorf("GetCnOfYyyyMmDd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCnOfMmDdHhMm(t *testing.T) {
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
			name: "GetCnOfMmDdHhMm",
			args: args{d: d},
			want: "04月13日 10:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCnOfMmDdHhMm(tt.args.d); got != tt.want {
				t.Errorf("GetCnOfMmDdHhMm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCnOfMmDdHhMmSs(t *testing.T) {
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
			name: "GetCnOfMmDdHhMmSs",
			args: args{d: d},
			want: "04月13日 10:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCnOfMmDdHhMmSs(tt.args.d); got != tt.want {
				t.Errorf("GetCnOfMmDdHhMmSs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCnOfYyyyMmDdHhMm(t *testing.T) {
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
			name: "GetCnOfYyyyMmDdHhMm",
			args: args{d: d},
			want: "2022年04月13日 10:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCnOfYyyyMmDdHhMm(tt.args.d); got != tt.want {
				t.Errorf("GetCnOfYyyyMmDdHhMm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCnOfYyyyMmDdHhMmSs(t *testing.T) {
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
			name: "GetCnOfYyyyMmDdHhMmSs",
			args: args{d: d},
			want: "2022年04月13日 10:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCnOfYyyyMmDdHhMmSs(tt.args.d); got != tt.want {
				t.Errorf("GetCnOfYyyyMmDdHhMmSs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCnOfYyyyMmDdHhMmSsSss(t *testing.T) {
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
			name: "GetCnOfYyyyMmDdHhMmSsSss",
			args: args{d: d},
			want: "2022年04月13日 10:00:00.999",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCnOfYyyyMmDdHhMmSsSss(tt.args.d); got != tt.want {
				t.Errorf("GetCnOfYyyyMmDdHhMmSsSss() = %v, want %v", got, tt.want)
			}
		})
	}
}
