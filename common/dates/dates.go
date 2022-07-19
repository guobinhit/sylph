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

// GetDayStart returns specified day start time.
func GetDayStart(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// GetDayEnd returns specified day end time.
func GetDayEnd(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 999999999, d.Location())
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
