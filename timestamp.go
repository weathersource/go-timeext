// This file provides special handling for timestamp strings formatted RFC 3339

package timeext

import (
	"math"
	"time"

	errors "github.com/weathersource/go-errors"
)

const timestampFormat = time.RFC3339

// Timestamp validates t is formatted RFC 3339 and returns time object.
func Timestamp(t string) (timestamp time.Time, err error) {
	timestamp, err = time.Parse(timestampFormat, t)
	if err != nil {
		err = errors.NewInvalidArgumentError("Timestamp must be formatted RFC 3339: "+t, err)
		return
	}
	timestamp = timestamp.UTC()
	return
}

// TimestampRounded validates t is formatted RFC 3339 and returns a time object
// rounded to the top of the UTC hour
func TimestampRounded(t string) (time.Time, error) {

	ts, err := Timestamp(t)
	if err != nil {
		return time.Time{}, err
	}

	rounded := time.Date(ts.Year(), ts.Month(), ts.Day(), ts.Hour(), 0, 0, 0, time.UTC)
	return rounded, nil
}

// TimestampRoundedString validates t is formatted RFC 3339 and returns a string formatted RFC 3339
// rounded to the top of the UTC hour
func TimestampRoundedString(t string) (string, error) {

	timestamp, err := TimestampRounded(t)
	if err != nil {
		return "", err
	}

	return timestamp.Format(timestampFormat), nil
}

// TimestampRoundedQuarterHour validates t is formatted RFC 3339 and returns a time object
// rounded down to the nearest UTC quarter hour
func TimestampRoundedQuarterHour(t string) (time.Time, error) {

	ts, err := Timestamp(t)
	if err != nil {
		return time.Time{}, err
	}

	rounded := time.Date(ts.Year(), ts.Month(), ts.Day(), ts.Hour(), (ts.Minute()) - (ts.Minute() % 15), 0, 0, time.UTC)
	return rounded, nil
}

// TimestampRoundedQuarterHourString validates t is formatted RFC 3339 and returns a string formatted RFC 3339
// rounded down to the nearest UTC quarter hour
func TimestampRoundedQuarterHourString(t string) (string, error) {

	timestamp, err := TimestampRoundedQuarterHour(t)
	if err != nil {
		return "", err
	}

	return timestamp.Format(timestampFormat), nil
}

// HourCount returns the count of days inclusively bounded by dateStart and dateEnd.
// dateStart and dateEnd must be formatted "YYYY-MM-DD"
func HourCount(timestampStart string, timestampEnd string) (int, error) {

    tStart, err := Timestamp(timestampStart)
    if err != nil {
        return 0, err
    }

    tEnd, err := Timestamp(timestampEnd)
    if err != nil {
        return 0, err
    }

    if tEnd.Before(tStart) {
        return 0, errors.NewInvalidArgumentError("Start Timestamp (" +
            timestampStart + ") must not be after End Timestamp (" + timestampEnd + ")")
    }

    return int(math.Floor(tEnd.Sub(tStart).Hours() + 1)), nil
}

// QuarterHourCount returns the count of days inclusively bounded by dateStart and dateEnd.
// dateStart and dateEnd must be formatted "YYYY-MM-DD"
func QuarterHourCount(timestampStart string, timestampEnd string) (int, error) {

    tStart, err := TimestampRoundedQuarterHour(timestampStart)
    if err != nil {
        return 0, err
    }

    tEnd, err := TimestampRoundedQuarterHour(timestampEnd)
    if err != nil {
        return 0, err
    }

    if tEnd.Before(tStart) {
        return 0, errors.NewInvalidArgumentError("Start Timestamp (" +
            timestampStart + ") must not be after End Timestamp (" + timestampEnd + ")")
    }

    return int((tEnd.Unix()-tStart.Unix())/(15*60) + 1), nil
}

// ToTimestamp converts a time object to a RFC 9993 timestamp string
func ToTimestamp(timestamp time.Time) string {
	return timestamp.UTC().Format(timestampFormat)
}
