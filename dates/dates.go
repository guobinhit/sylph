package dates

import "time"

// GetTodayStart returns start time of today.
func GetTodayStart() time.Time {
	return GetDayStart(time.Now())
}

// GetTodayEnd returns end time of today.
func GetTodayEnd() time.Time {
	return GetDayEnd(time.Now())
}

// GetSecondStart returns specified start time of second.
func GetSecondStart(d time.Time) time.Time {
	return d.Truncate(time.Second)
}

// GetSecondEnd returns specified end time of second.
func GetSecondEnd(d time.Time) time.Time {
	return GetSecondStart(d).Add(time.Second - time.Nanosecond)
}

// GetMinuteStart returns specified start time of minute.
func GetMinuteStart(d time.Time) time.Time {
	return d.Truncate(time.Minute)
}

// GetMinuteEnd returns specified end time of minute.
func GetMinuteEnd(d time.Time) time.Time {
	return GetMinuteStart(d).Add(time.Minute - time.Nanosecond)
}

// GetHourStart returns specified start time of hour.
func GetHourStart(d time.Time) time.Time {
	return d.Truncate(time.Hour)
}

// GetHourEnd returns specified end time of hour.
func GetHourEnd(d time.Time) time.Time {
	return GetHourStart(d).Add(time.Hour - time.Nanosecond)
}

// GetDayStart returns specified start time day.
func GetDayStart(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// GetDayEnd returns specified end time of day.
func GetDayEnd(d time.Time) time.Time {
	return GetDayStart(d).AddDate(0, 0, 1).Add(-time.Nanosecond)
}

// GetWeekStart returns specified start time of week.
// Param of weekStartDay is used to set the start time of week.
func GetWeekStart(d time.Time, weekStartDay time.Weekday) time.Time {
	t := GetDayStart(d)
	weekday := int(t.Weekday())
	if weekStartDay != time.Sunday {
		weekStartDayInt := int(weekStartDay)
		if weekday < weekStartDayInt {
			weekday = weekday + 7 - weekStartDayInt
		} else {
			weekday = weekday - weekStartDayInt
		}
	}
	return t.AddDate(0, 0, -weekday)
}

// GetWeekEnd returns specified end time of week.
// Param of weekStartDay is used to set the start time of week.
func GetWeekEnd(d time.Time, weekStartDay time.Weekday) time.Time {
	return GetWeekStart(d, weekStartDay).AddDate(0, 0, 7).Add(-time.Nanosecond)
}

// GetMonthStart returns specified start time of month.
func GetMonthStart(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), 1, 0, 0, 0, 0, d.Location())
}

// GetMonthEnd returns specified end time of month.
func GetMonthEnd(d time.Time) time.Time {
	return GetMonthStart(d).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// GetYearStart returns specified start time of year.
func GetYearStart(d time.Time) time.Time {
	return time.Date(d.Year(), time.January, 1, 0, 0, 0, 0, d.Location())
}

// GetYearEnd returns specified end time of year.
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
// 1. t1 after t2 (like t1 is 2022-05-13 10:00:00, t2 is 2022-04-13 10:00:00), the result is negative;
// 2. t1 equal t2 (like t1 is 2022-04-13 10:00:00, t2 is 2022-04-13 10:00:00), the result is 0;
// 3. t1 before t2 (like t1 is 2022-04-13 10:00:00, t2 is 2022-05-13 10:00:00), the result is positive.
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
