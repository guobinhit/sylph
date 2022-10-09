package dates

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

func TestGetMinuteStart(t *testing.T) {
	dateWant, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-13 12:23:00", time.Local)
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-13 12:23:34", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "GetMinuteStart",
			args: args{
				d: d,
			},
			want: dateWant,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMinuteStart(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMinuteStart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMinuteEnd(t *testing.T) {
	dateWant, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-13 12:23:59.999999999", time.Local)
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-13 12:23:34", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "GetMinuteEnd",
			args: args{
				d: d,
			},
			want: dateWant,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMinuteEnd(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMinuteEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetHourStart(t *testing.T) {
	dateWant, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-13 12:00:00", time.Local)
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-13 12:23:34", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "GetHourStart",
			args: args{
				d: d,
			},
			want: dateWant,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHourStart(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHourStart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetHourEnd(t *testing.T) {
	dateWant, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-13 12:59:59.999999999", time.Local)
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-13 12:23:34", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "GetHourEnd",
			args: args{
				d: d,
			},
			want: dateWant,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHourEnd(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHourEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDayStart(t *testing.T) {
	dateWant, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-13 00:00:00", time.Local)
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
			want: dateWant,
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

func TestGetTimeAddYears(t *testing.T) {
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-13 10:00:00", time.Local)
	type args struct {
		d     time.Time
		years int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "GetTimeAddYears + 10",
			args: args{
				d:     d,
				years: 10,
			},
			want: "2032-04-13 10:00:00",
		},
		{
			name: "GetTimeAddYears - 10",
			args: args{
				d:     d,
				years: -10,
			},
			want: "2012-04-13 10:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetTimeAddYears(tt.args.d, tt.args.years)
			gotStr := got.Format(date_const.YyyyMmDdHhMmSs)
			if !reflect.DeepEqual(gotStr, tt.want) {
				t.Errorf("GetTimeAddYears() = %v, want %v", gotStr, tt.want)
			}
		})
	}
}

func TestGetTimeAddMonths(t *testing.T) {
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-13 10:00:00", time.Local)
	type args struct {
		d      time.Time
		months int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "GetTimeAddMonths + 10",
			args: args{
				d:      d,
				months: 10,
			},
			want: "2023-02-13 10:00:00",
		},
		{
			name: "GetTimeAddMonths - 10",
			args: args{
				d:      d,
				months: -10,
			},
			want: "2021-06-13 10:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetTimeAddMonths(tt.args.d, tt.args.months)
			gotStr := got.Format(date_const.YyyyMmDdHhMmSs)
			if !reflect.DeepEqual(gotStr, tt.want) {
				t.Errorf("GetTimeAddMonths() = %v, want %v", gotStr, tt.want)
			}
		})
	}
}

func TestGetTimeAddDays(t *testing.T) {
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-13 10:00:00", time.Local)
	type args struct {
		d    time.Time
		days int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "GetTimeAddDays + 10",
			args: args{
				d:    d,
				days: 10,
			},
			want: "2022-04-23 10:00:00",
		},
		{
			name: "GetTimeAddDays - 10",
			args: args{
				d:    d,
				days: -10,
			},
			want: "2022-04-03 10:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetTimeAddDays(tt.args.d, tt.args.days)
			gotStr := got.Format(date_const.YyyyMmDdHhMmSs)
			if !reflect.DeepEqual(gotStr, tt.want) {
				t.Errorf("GetTimeAddDays() = %v, want %v", gotStr, tt.want)
			}
		})
	}
}

func TestGetTimeAddHours(t *testing.T) {
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-13 10:00:00", time.Local)
	type args struct {
		d     time.Time
		hours int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "GetTimeAddHours + 10",
			args: args{
				d:     d,
				hours: 10,
			},
			want: "2022-04-13 20:00:00",
		},
		{
			name: "GetTimeAddHours - 10",
			args: args{
				d:     d,
				hours: -10,
			},
			want: "2022-04-13 00:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetTimeAddHours(tt.args.d, tt.args.hours)
			gotStr := got.Format(date_const.YyyyMmDdHhMmSs)
			if !reflect.DeepEqual(gotStr, tt.want) {
				t.Errorf("GetTimeAddHours() = %v, want %v", gotStr, tt.want)
			}
		})
	}
}

func TestGetTimeAddMinutes(t *testing.T) {
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-13 10:00:00", time.Local)
	type args struct {
		d       time.Time
		minutes int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "GetTimeAddMinutes + 10",
			args: args{
				d:       d,
				minutes: 10,
			},
			want: "2022-04-13 10:10:00",
		},
		{
			name: "GetTimeAddMinutes - 10",
			args: args{
				d:       d,
				minutes: -10,
			},
			want: "2022-04-13 09:50:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetTimeAddMinutes(tt.args.d, tt.args.minutes)
			gotStr := got.Format(date_const.YyyyMmDdHhMmSs)
			if !reflect.DeepEqual(gotStr, tt.want) {
				t.Errorf("GetTimeAddMinutes() = %v, want %v", gotStr, tt.want)
			}
		})
	}
}

func TestGetTimeAddSeconds(t *testing.T) {
	d, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-13 10:00:00", time.Local)
	type args struct {
		d       time.Time
		seconds int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "GetTimeAddSeconds + 10",
			args: args{
				d:       d,
				seconds: 10,
			},
			want: "2022-04-13 10:00:10",
		},
		{
			name: "GetTimeAddSeconds - 10",
			args: args{
				d:       d,
				seconds: -10,
			},
			want: "2022-04-13 09:59:50",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetTimeAddSeconds(tt.args.d, tt.args.seconds)
			gotStr := got.Format(date_const.YyyyMmDdHhMmSs)
			if !reflect.DeepEqual(gotStr, tt.want) {
				t.Errorf("GetTimeAddSeconds() = %v, want %v", gotStr, tt.want)
			}
		})
	}
}

func TestGetTimeSubDays(t *testing.T) {
	t1, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-04-13 10:00:00", time.Local)
	t2, _ := time.ParseInLocation(date_const.YyyyMmDdHhMmSs, "2022-05-13 10:00:00", time.Local)
	type args struct {
		t1 time.Time
		t2 time.Time
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1 > t2",
			args: args{
				t1: t1,
				t2: t2,
			},
			want: 30,
		},
		{
			name: "t1 < t2",
			args: args{
				t1: t2,
				t2: t1,
			},
			want: -30,
		},
		{
			name: "t1 == t2",
			args: args{
				t1: t1,
				t2: t1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTimeSubDays(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("GetTimeSubDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
