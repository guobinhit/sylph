package dates

import "time"

// GetTodayStart returns today start time.
func GetTodayStart() time.Time {
	return GetDayStart(time.Now())
}

// GetTodayEnd returns today end time.
func GetTodayEnd() time.Time {
	return GetDayEnd(time.Now())
}

// GetSecondStart returns specified second start time.
func GetSecondStart(d time.Time) time.Time {
	return d.Truncate(time.Second)
}

// GetSecondEnd returns specified second end time.
func GetSecondEnd(d time.Time) time.Time {
	return GetSecondStart(d).Add(time.Second - time.Nanosecond)
}

// GetMinuteStart returns specified minute start time.
func GetMinuteStart(d time.Time) time.Time {
	return d.Truncate(time.Minute)
}

// GetMinuteEnd returns specified minute end time.
func GetMinuteEnd(d time.Time) time.Time {
	return GetMinuteStart(d).Add(time.Minute - time.Nanosecond)
}

// GetHourStart returns specified hour start time.
func GetHourStart(d time.Time) time.Time {
	return d.Truncate(time.Hour)
}

// GetHourEnd returns specified hour end time.
func GetHourEnd(d time.Time) time.Time {
	return GetHourStart(d).Add(time.Hour - time.Nanosecond)
}

// GetDayStart returns specified day start time.
func GetDayStart(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// GetDayEnd returns specified day end time.
func GetDayEnd(d time.Time) time.Time {
	return GetDayStart(d).AddDate(0, 0, 1).Add(-time.Nanosecond)
}

// GetMonthStart returns specified month start time.
func GetMonthStart(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), 1, 0, 0, 0, 0, d.Location())
}

// GetMonthEnd returns specified month end time.
func GetMonthEnd(d time.Time) time.Time {
	return GetMonthStart(d).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// GetYearStart returns specified year start time.
func GetYearStart(d time.Time) time.Time {
	return time.Date(d.Year(), time.January, 1, 0, 0, 0, 0, d.Location())
}

// GetYearEnd returns specified year end time.
func GetYearEnd(d time.Time) time.Time {
	return GetYearStart(d).AddDate(1, 0, 0).Add(-time.Nanosecond)
}

// GetTimeAddYears returns specified time by add years.
func GetTimeAddYears(d time.Time, years int) time.Time {
	return d.AddDate(years, 0, 0)
}

// GetTimeAddMonths returns specified time by add months.
func GetTimeAddMonths(d time.Time, months int) time.Time {
	return d.AddDate(0, months, 0)
}

// GetTimeAddDays returns specified time by add days.
func GetTimeAddDays(d time.Time, days int) time.Time {
	return d.AddDate(0, 0, days)
}

// GetTimeAddHours returns specified time by add hours.
func GetTimeAddHours(d time.Time, hours int) time.Time {
	return d.Add(time.Hour * time.Duration(hours))
}

// GetTimeAddMinutes returns specified time by add minutes.
func GetTimeAddMinutes(d time.Time, minutes int) time.Time {
	return d.Add(time.Minute * time.Duration(minutes))
}

// GetTimeAddSeconds returns specified time by add seconds.
func GetTimeAddSeconds(d time.Time, seconds int) time.Time {
	return d.Add(time.Second * time.Duration(seconds))
}

// GetTimeSubDays returns number of days before two times
// 1. If t1 after t2 (like t1 is 2022-05-13 10:00:00, t2 is 2022-04-13 10:00:00), the result is negative;
// 2. If t1 equal t2 (like t1 is 2022-04-13 10:00:00, t2 is 2022-04-13 10:00:00), the result is 0;
// 3. If t1 before t2 (like t1 is 2022-04-13 10:00:00, t2 is 2022-05-13 10:00:00), the result is positive.
func GetTimeSubDays(t1, t2 time.Time) int {
	var days int
	isSwap := false
	if t1.Unix() > t2.Unix() {
		t1, t2 = t2, t1
		isSwap = true
	}
	ans := t1.Add(time.Duration(t2.Sub(t1).Milliseconds()%86400000) * time.Millisecond)
	days = int(t2.Sub(t1).Hours() / 24)
	if ans.Day() != t1.Day() {
		days += 1
	}
	if isSwap {
		days = -days
	}
	return days
}
