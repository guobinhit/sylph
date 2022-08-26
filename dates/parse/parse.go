package parse

import (
	"github.com/guobinhit/sylph/constant/date_const"
	"time"
)

// GetMmDd returns time.Time with '01-02' format value，others location is default value.
func GetMmDd(dStr string) (time.Time, error) {
	return time.ParseInLocation(date_const.MmDd, dStr, time.Local)
}

// GetYyyyMm returns time.Time with '2006-01' format value，others location is default value.
func GetYyyyMm(dStr string) (time.Time, error) {
	return time.ParseInLocation(date_const.YyyyMm, dStr, time.Local)
}

// GetYyyyMmDd returns time.Time with '2006-01-02' format value，others location is default value.
func GetYyyyMmDd(dStr string) (time.Time, error) {
	return time.ParseInLocation(date_const.YyyyMmDd, dStr, time.Local)
}

// GetMmDdHhMm returns time.Time with '01-02 15:04' format value，others location is default value.
func GetMmDdHhMm(dStr string) (time.Time, error) {
	return time.ParseInLocation(date_const.MmDdHhMm, dStr, time.Local)
}

// GetMmDdHhMmSs returns time.Time with '01-02 15:04:05' format value，others location is default value.
func GetMmDdHhMmSs(dStr string) (time.Time, error) {
	return time.ParseInLocation(date_const.MmDdHhMmSs, dStr, time.Local)
}

// GetMmDdHhMmSsSss returns time.Time with '01-02 15:04:05.000' format value，others location is default value.
func GetMmDdHhMmSsSss(dStr string) (time.Time, error) {
	return time.ParseInLocation(date_const.MmDdHhMmSsSss, dStr, time.Local)
}

// GetYyyyMmDdHhMm returns time.Time with '2006-01-02 15:04' format value，others location is default value.
func GetYyyyMmDdHhMm(dStr string) (time.Time, error) {
	return time.ParseInLocation(date_const.YyyyMmDdHhMm, dStr, time.Local)
}

// GetYyyyMmDdHhMmSs returns time.Time with '2006-01-02 15:04:05' format value，others location is default value.
func GetYyyyMmDdHhMmSs(dStr string) (time.Time, error) {
	return time.ParseInLocation(date_const.YyyyMmDdHhMmSs, dStr, time.Local)
}

// GetYyyyMmDdHhMmSsSss returns time.Time with '2006-01-02 15:04:05.000' format value，others location is default value.
func GetYyyyMmDdHhMmSsSss(dStr string) (time.Time, error) {
	return time.ParseInLocation(date_const.YyyyMmDdHhMmSsSss, dStr, time.Local)
}
