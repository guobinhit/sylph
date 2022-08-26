package format

import (
	"github.com/guobinhit/sylph/constant/date_const"
	"time"
)

// GetMmDd returns '01-02' format date string.
func GetMmDd(d time.Time) string {
	return d.Format(date_const.MmDd)
}

// GetYyyyMm returns '2006-01' format date string.
func GetYyyyMm(d time.Time) string {
	return d.Format(date_const.YyyyMm)
}

// GetYyyyMmDd returns '2006-01-02' format date string.
func GetYyyyMmDd(d time.Time) string {
	return d.Format(date_const.YyyyMmDd)
}

// GetMmDdHhMm returns '01-02 15:04' format date string.
func GetMmDdHhMm(d time.Time) string {
	return d.Format(date_const.MmDdHhMm)
}

// GetMmDdHhMmSs returns '01-02 15:04:05' format date string.
func GetMmDdHhMmSs(d time.Time) string {
	return d.Format(date_const.MmDdHhMmSs)
}

// GetMmDdHhMmSsSss returns '01-02 15:04:05.000' format date string.
func GetMmDdHhMmSsSss(d time.Time) string {
	return d.Format(date_const.MmDdHhMmSsSss)
}

// GetYyyyMmDdHhMm returns '2006-01-02 15:04' format date string.
func GetYyyyMmDdHhMm(d time.Time) string {
	return d.Format(date_const.YyyyMmDdHhMm)
}

// GetYyyyMmDdHhMmSs returns '2006-01-02 15:04:05' format date string.
func GetYyyyMmDdHhMmSs(d time.Time) string {
	return d.Format(date_const.YyyyMmDdHhMmSs)
}

// GetYyyyMmDdHhMmSsSss returns '2006-01-02 15:04:05.000' format date string.
func GetYyyyMmDdHhMmSsSss(d time.Time) string {
	return d.Format(date_const.YyyyMmDdHhMmSsSss)
}
