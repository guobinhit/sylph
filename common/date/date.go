package date

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
