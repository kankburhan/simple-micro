package util

import (
	"time"
)

var (
	LayoutDefault  = "2006-01-02 15:04:05"
	LayoutDateOnly = "2006-01-02"
	Loc, _         = time.LoadLocation("Asia/Jakarta")
)

func DateRangeValidation(dateStart, dateEnd string) (validDate, validRange bool) {
	validDate = true
	validRange = false
	loc, _ := time.LoadLocation("Asia/Jakarta")

	dateStartT, err := time.ParseInLocation(LayoutDateOnly, dateStart, loc)
	if err != nil {
		validDate = false
	}

	dateEndT, err := time.ParseInLocation(LayoutDateOnly, dateEnd, loc)
	if err != nil {
		validDate = false
	}

	if dateEndT.After(dateStartT) {
		validRange = true
	}

	return validDate, validRange
}

func DateValidationYYYY_MM_DD(date string) (validDate bool) {
	validDate = true
	loc, _ := time.LoadLocation("Asia/Jakarta")

	_, err := time.ParseInLocation(LayoutDateOnly, date, loc)
	if err != nil {
		validDate = false
	}
	return validDate
}

func ConvertToDateTime(dateTime string) (time.Time, error) {
	DateTime, err := time.ParseInLocation(LayoutDefault, dateTime, Loc)
	if err != nil {
		return time.Time{}, err
	}

	return DateTime, err
}
