package format

import (
	"github.com/guobinhit/sylph/constant/date_const"
	"time"
)

// GetCptOfYyyyMm returns '200601' format date string.
func GetCptOfYyyyMm(d time.Time) string {
	return d.Format(date_const.CptOfYyyyMm)
}

// GetCptOfYyyyMmDd returns '20060102' format date string.
func GetCptOfYyyyMmDd(d time.Time) string {
	return d.Format(date_const.CptOfYyyyMmDd)
}

// GetCptOfYyyyMmDdHh returns '2006010215' format date string.
func GetCptOfYyyyMmDdHh(d time.Time) string {
	return d.Format(date_const.CptOfYyyyMmDdHh)
}

// GetCptOfYyyyMmDdHhMm returns '200601021504' format date string.
func GetCptOfYyyyMmDdHhMm(d time.Time) string {
	return d.Format(date_const.CptOfYyyyMmDdHhMm)
}

// GetCptOfYyyyMmDdHhMmSs returns '20060102150405' format date string.
func GetCptOfYyyyMmDdHhMmSs(d time.Time) string {
	return d.Format(date_const.CptOfYyyyMmDdHhMmSs)
}
