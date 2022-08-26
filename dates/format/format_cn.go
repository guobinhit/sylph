package format

import (
	"github.com/guobinhit/sylph/constant/date_const"
	"time"
)

// GetCnOfMmDd returns '01月02日' format date string.
func GetCnOfMmDd(d time.Time) string {
	return d.Format(date_const.CnOfMmDd)
}

// GetCnOfYyyyMm returns '2006年01月' format date string.
func GetCnOfYyyyMm(d time.Time) string {
	return d.Format(date_const.CnOfYyyyMm)
}

// GetCnOfYyyyMmDd returns '2006-01-02' format date string.
func GetCnOfYyyyMmDd(d time.Time) string {
	return d.Format(date_const.CnOfYyyyMmDd)
}

// GetCnOfMmDdHhMm returns '01月02日 15:04' format date string.
func GetCnOfMmDdHhMm(d time.Time) string {
	return d.Format(date_const.CnOfMmDdHhMm)
}

// GetCnOfMmDdHhMmSs returns '01月02日 15:04:05' format date string.
func GetCnOfMmDdHhMmSs(d time.Time) string {
	return d.Format(date_const.CnOfMmDdHhMmSs)
}

// GetCnOfYyyyMmDdHhMm returns '2006年01月02日 15:04' format date string.
func GetCnOfYyyyMmDdHhMm(d time.Time) string {
	return d.Format(date_const.CnOfYyyyMmDdHhMm)
}

// GetCnOfYyyyMmDdHhMmSs returns '2006年01月02日 15:04:05' format date string.
func GetCnOfYyyyMmDdHhMmSs(d time.Time) string {
	return d.Format(date_const.CnOfYyyyMmDdHhMmSs)
}

// GetCnOfYyyyMmDdHhMmSsSss returns '2006年01月02日 15:04:05.000' format date string.
func GetCnOfYyyyMmDdHhMmSsSss(d time.Time) string {
	return d.Format(date_const.CnOfYyyyMmDdHhMmSsSss)
}
