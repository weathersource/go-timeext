// This file provides special handling for strings formatted YYYY-MM-DD

package timeext

import (
	"math"
	"time"

	errors "github.com/weathersource/go-errors"
)

const dateFormat = "2006-01-02"

// Date validates t is formatted YYYY-MM-DD and returns time object. This should
// only be used in contexts where the timezone does not matter (i.e. local time
// is not needed)
func Date(d string) (date time.Time, err error) {
	date, err = time.ParseInLocation(dateFormat, d, time.UTC)
	if err != nil {
		err = errors.NewInvalidArgumentError("Date must be formatted \"YYYY-MM-DD\": "+d, err)
		return
	}
	return
}

// DayCount returns the count of days inclusively bounded by dateStart and dateEnd.
// dateStart and dateEnd must be formatted "YYYY-MM-DD"
func DayCount(dateStart string, dateEnd string) (int, error) {

	tStart, err := Date(dateStart)
	if err != nil {
		return 0, err
	}

	tEnd, err := Date(dateEnd)
	if err != nil {
		return 0, err
	}

	if tEnd.Before(tStart) {
		return 0, errors.NewInvalidArgumentError("Start Date (" + dateStart + ") must not be after End Date (" + dateEnd + ")")
	}

	return int(math.Floor(tEnd.Sub(tStart).Hours()/24 + 1)), nil
}

// ToDate converts a time object to a date string
func ToDate(date time.Time) string {
	return date.UTC().Format(dateFormat)
}

// Returns a timestamp representing a day period in a given timezone, accounting
// for when DST occurs at midnight (see https://github.com/golang/go/issues/35508).
// This should be used in contexts where local time matters.
func FirstTimeInDay(d string, loc *time.Location) (date time.Time, err error) {
	date, err = time.ParseInLocation(dateFormat, d, loc)
	if err != nil {
		err = errors.NewInvalidArgumentError("Date must be formatted \"YYYY-MM-DD\": "+d, err)
		return
	}

	// check for DST at midnight
	if date.Format(dateFormat) != d {
		date = date.Add(time.Hour)
	}

	return
}
