package parse

import (
	"github.com/guobinhit/sylph/constant/date_const"
	"reflect"
	"testing"
	"time"
)

func TestGetMmDd(t *testing.T) {
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "0000-04-13 00:00:00", time.Local)
	type args struct {
		dStr string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "GetMmDd error",
			args:    args{dStr: "2022-04-13"},
			wantErr: true,
		},
		{
			name: "GetMmDd success",
			args: args{dStr: "04-13"},
			want: d,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMmDd(tt.args.dStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMmDd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMmDd() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetYyyyMm(t *testing.T) {
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-01 00:00:00", time.Local)
	type args struct {
		dStr string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "GetYyyyMm error",
			args:    args{dStr: "2022-04-13"},
			wantErr: true,
		},
		{
			name: "GetYyyyMm success",
			args: args{dStr: "2022-04"},
			want: d,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetYyyyMm(tt.args.dStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetYyyyMm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetYyyyMm() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetYyyyMmDd(t *testing.T) {
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-13 00:00:00", time.Local)
	type args struct {
		dStr string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "GetYyyyMmDd error",
			args:    args{dStr: "2022-04-13 00:00:00"},
			wantErr: true,
		},
		{
			name: "GetYyyyMmDd success",
			args: args{dStr: "2022-04-13"},
			want: d,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetYyyyMmDd(tt.args.dStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetYyyyMmDd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetYyyyMmDd() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMmDdHhMm(t *testing.T) {
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "0000-04-13 10:20:00", time.Local)
	type args struct {
		dStr string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "GetMmDdHhMm error",
			args:    args{dStr: "2022-04-13 00:00:00"},
			wantErr: true,
		},
		{
			name: "GetMmDdHhMm success",
			args: args{dStr: "04-13 10:20"},
			want: d,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMmDdHhMm(tt.args.dStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMmDdHhMm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMmDdHhMm() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMmDdHhMmSs(t *testing.T) {
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "0000-04-13 10:20:30", time.Local)
	type args struct {
		dStr string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "GetMmDdHhMmSs error",
			args:    args{dStr: "2022-04-13 00:00:00.123"},
			wantErr: true,
		},
		{
			name: "GetMmDdHhMmSs success",
			args: args{dStr: "04-13 10:20:30"},
			want: d,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMmDdHhMmSs(tt.args.dStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMmDdHhMmSs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMmDdHhMmSs() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMmDdHhMmSsSss(t *testing.T) {
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSsSss, "0000-04-13 10:20:30.123", time.Local)
	type args struct {
		dStr string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "GetMmDdHhMmSsSss error",
			args:    args{dStr: "2022-04-13 00:00:00.123"},
			wantErr: true,
		},
		{
			name: "GetMmDdHhMmSsSss success",
			args: args{dStr: "04-13 10:20:30.123"},
			want: d,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMmDdHhMmSsSss(tt.args.dStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMmDdHhMmSsSss() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMmDdHhMmSsSss() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetYyyyMmDdHhMm(t *testing.T) {
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSsSss, "2022-04-13 10:20:00.000", time.Local)
	type args struct {
		dStr string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "GetYyyyMmDdHhMm error",
			args:    args{dStr: "2022-04-13 00:00:00.123"},
			wantErr: true,
		},
		{
			name: "GetYyyyMmDdHhMm success",
			args: args{dStr: "2022-04-13 10:20"},
			want: d,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetYyyyMmDdHhMm(tt.args.dStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetYyyyMmDdHhMm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetYyyyMmDdHhMm() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetYyyyMmDdHhMmSs(t *testing.T) {
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSsSss, "2022-04-13 10:20:30.000", time.Local)
	type args struct {
		dStr string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "GetYyyyMmDdHhMmSs error",
			args:    args{dStr: "04-13 00:00:00.123"},
			wantErr: true,
		},
		{
			name: "GetYyyyMmDdHhMmSs success",
			args: args{dStr: "2022-04-13 10:20:30"},
			want: d,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetYyyyMmDdHhMmSs(tt.args.dStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetYyyyMmDdHhMmSs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetYyyyMmDdHhMmSs() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetYyyyMmDdHhMmSsSss(t *testing.T) {
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSsSss, "2022-04-13 10:20:30.999", time.Local)
	type args struct {
		dStr string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "GetYyyyMmDdHhMmSsSss error",
			args:    args{dStr: "04-13 00:00:00.123"},
			wantErr: true,
		},
		{
			name: "GetYyyyMmDdHhMmSsSss success",
			args: args{dStr: "2022-04-13 10:20:30.999"},
			want: d,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetYyyyMmDdHhMmSsSss(tt.args.dStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetYyyyMmDdHhMmSsSss() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetYyyyMmDdHhMmSsSss() got = %v, want %v", got, tt.want)
			}
		})
	}
}
