package format

import (
	"github.com/guobinhit/sylph/constant/date_const"
	"time"
)

// GetEnOfMmDd returns '01/02' format date string.
func GetEnOfMmDd(d time.Time) string {
	return d.Format(date_const.EnOfMmDd)
}

// GetEnOfYyyyMm returns '2006/01' format date string.
func GetEnOfYyyyMm(d time.Time) string {
	return d.Format(date_const.EnOfYyyyMm)
}

// GetEnOfYyyyMmDd returns '2006/01/02' format date string.
func GetEnOfYyyyMmDd(d time.Time) string {
	return d.Format(date_const.EnOfYyyyMmDd)
}

// GetEnOfMmDdHhMm returns '01/02 15:04' format date string.
func GetEnOfMmDdHhMm(d time.Time) string {
	return d.Format(date_const.EnOfMmDdHhMm)
}

// GetEnOfMmDdHhMmSs returns '01/02 15:04:05' format date string.
func GetEnOfMmDdHhMmSs(d time.Time) string {
	return d.Format(date_const.EnOfMmDdHhMmSs)
}

// GetEnOfYyyyMmDdHhMm returns '2006/01/02 15:04' format date string.
func GetEnOfYyyyMmDdHhMm(d time.Time) string {
	return d.Format(date_const.EnOfYyyyMmDdHhMm)
}

// GetEnOfYyyyMmDdHhMmSs returns '2006/01/02 15:04:05' format date string.
func GetEnOfYyyyMmDdHhMmSs(d time.Time) string {
	return d.Format(date_const.EnOfYyyyMmDdHhMmSs)
}

// GetEnOfYyyyMmDdHhMmSsSss returns '2006/01/02 15:04:05.000' format date string.
func GetEnOfYyyyMmDdHhMmSsSss(d time.Time) string {
	return d.Format(date_const.EnOfYyyyMmDdHhMmSsSss)
}
