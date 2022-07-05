package date

import (
	"github.com/guobinhit/sylph/constant/date_const"
	"reflect"
	"testing"
	"time"
)

func TestGetTodayStart(t *testing.T) {
	tests := []struct {
		name string
		want time.Time
	}{
		{
			name: "GetTodayStart",
			want: GetDayStart(time.Now()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTodayStart(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTodayStart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTodayEnd(t *testing.T) {
	tests := []struct {
		name string
		want time.Time
	}{
		{
			name: "GetTodayEnd",
			want: GetDayEnd(time.Now()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTodayEnd(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTodayEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDayStart(t *testing.T) {
	dateStart, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-13 00:00:00", time.Local)
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-13 09:00:00", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "GetDayStart",
			args: args{
				d: d,
			},
			want: dateStart,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDayStart(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDayStart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDayEnd(t *testing.T) {
	dateEnd, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-13 23:59:59.999999999", time.Local)
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-13 09:00:00", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "GetDayEnd",
			args: args{
				d: d,
			},
			want: dateEnd,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDayEnd(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDayEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
