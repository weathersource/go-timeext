package timeext

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimestamp(t *testing.T) {
	tests := []struct {
		ts  string
		out string
	}{
		{"1950-01-01T00:00:00Z", "1950-01-01T00:00:00Z"},
		{"1969-12-31T00:00:00-01:00", "1969-12-31T01:00:00Z"},
		{"1970-01-01T00:00:00+01:00", "1969-12-31T23:00:00Z"},
		{"2019-01-01T00:00:00.1Z", "2019-01-01T00:00:00Z"},
	}
	for _, test := range tests {
		out, err := Timestamp(test.ts)
		assert.Nil(t, err)
		assert.Equal(t, test.out, out.Format(timestampFormat))
	}
}

func TestTimestampError(t *testing.T) {
	tests := []struct {
		ts string
	}{
		{"2019-01-01T00:00:00"},
		{"2019-01-01T00:00Z"},
		{"2019-01-01T00:00:00.Z"},
		{"2019-01-01T00:00:00+01"},
		{"2019-01-01T00:00:00+0100"},
		{""},
	}
	for _, test := range tests {
		_, err := Timestamp(test.ts)
		assert.NotNil(t, err)
	}
}

func TestTimestampRounded(t *testing.T) {
	tests := []struct {
		ts string
	}{
		{"2019-01-01T00:30:00Z"},
	}
	for _, test := range tests {
		_, err := TimestampRounded(test.ts)
		assert.Nil(t, err)
	}
}

func TestTimestampRoundedError(t *testing.T) {
	tests := []struct {
		ts string
	}{
		{"2019-01-01T00:30:00A"},
	}
	for _, test := range tests {
		_, err := TimestampRounded(test.ts)
		assert.NotNil(t, err)
	}
}

func TestTimestampRoundedString(t *testing.T) {
	tests := []struct {
		ts string
	}{
		{"2019-01-01T00:30:00Z"},
	}
	for _, test := range tests {
		_, err := TimestampRoundedString(test.ts)
		assert.Nil(t, err)
	}
}

func TestTimestampRoundedStringError(t *testing.T) {
	tests := []struct {
		ts string
	}{
		{"2019-01-01T00:30:00A"},
	}
	for _, test := range tests {
		_, err := TimestampRoundedString(test.ts)
		assert.NotNil(t, err)
	}
}

func TestTimestampRoundedQuarterHour(t *testing.T) {
	tests := []struct {
		ts string
		out string
	}{
		{"2019-01-01T00:33:01Z", "2019-01-01T00:30:00Z"},
		{"2019-01-01T00:58:13Z", "2019-01-01T00:45:00Z"},
		{"2019-01-01T00:11:59Z", "2019-01-01T00:00:00Z"},
		{"2019-01-01T00:29:50Z", "2019-01-01T00:15:00Z"},
		{"2019-01-01T23:52:30Z", "2019-01-01T23:45:00Z"},
		{"2019-01-01T23:52:29Z", "2019-01-01T23:45:00Z"},
	}
	for _, test := range tests {
		out, err := TimestampRoundedQuarterHour(test.ts)
		assert.Nil(t, err)
		assert.Equal(t, test.out, out.Format(timestampFormat))
	}
}

func TestTimestampRoundedQuarterHourError(t *testing.T) {
	tests := []struct {
		ts string
	}{
		{"2019-01-01T00:30:00A"},
	}
	for _, test := range tests {
		_, err := TimestampRoundedQuarterHour(test.ts)
		assert.NotNil(t, err)
	}
}

func TestTimestampRoundedQuarterHourString(t *testing.T) {
	tests := []struct {
		ts string
		out string
	}{
		{"2019-01-01T00:33:00Z", "2019-01-01T00:30:00Z"},
	}
	for _, test := range tests {
		out, err := TimestampRoundedQuarterHourString(test.ts)
		assert.Nil(t, err)
		assert.Equal(t, test.out, out)
	}
}

func TestTimestampRoundedQuarterHourStringError(t *testing.T) {
	tests := []struct {
		ts string
	}{
		{"2019-01-01T00:30:00A"},
	}
	for _, test := range tests {
		_, err := TimestampRoundedQuarterHourString(test.ts)
		assert.NotNil(t, err)
	}
}

func TestHourCount(t *testing.T) {
    tests := []struct {
        ts1, ts2 string
        expected int
    }{
        {"2019-01-01T00:00:00Z", "2019-01-01T00:00:00.1Z", 1},
        {"2019-01-01T00:00:00Z", "2019-01-01T02:00:00.1Z", 3},
        {"2019-01-01T00:00:00Z", "2019-01-01T02:30:00.1Z", 3},
    }
    for _, test := range tests {
        count, err := HourCount(test.ts1, test.ts2)
        assert.Nil(t, err)
        assert.Equal(t, test.expected, count)
    }
}

func TestHourCountError(t *testing.T) {
    tests := []struct {
        ts1, ts2 string
    }{
        {"2019-01-01T00:00:00A", "2019-01-01T00:00:00.1Z"},
        {"2019-01-01T00:00:00Z", "2019-01-01T00:00:00.1A"},
        {"2019-01-01T02:00:00Z", "2019-01-01T00:00:00.1Z"},
    }
    for _, test := range tests {
        _, err := HourCount(test.ts1, test.ts2)
        assert.NotNil(t, err)
    }
}

func TestQuarterHourCount(t *testing.T) {
    tests := []struct {
        ts1, ts2 string
        expected int
    }{
        {"2019-01-01T00:00:00Z", "2019-01-01T00:00:00.1Z", 1},
        {"2019-01-01T00:00:00Z", "2019-01-01T02:00:00.1Z", 9},
        {"2019-01-01T00:00:00Z", "2019-01-01T02:29:00.1Z", 10},
    }
    for _, test := range tests {
        count, err := QuarterHourCount(test.ts1, test.ts2)
        assert.Nil(t, err)
        assert.Equal(t, test.expected, count)
    }
}

func TestQuarterHourCountError(t *testing.T) {
    tests := []struct {
        ts1, ts2 string
    }{
        {"2019-01-01T00:00:00A", "2019-01-01T00:00:00.1Z"},
        {"2019-01-01T00:00:00Z", "2019-01-01T00:00:00.1A"},
        {"2019-01-01T02:00:00Z", "2019-01-01T00:00:00.1Z"},
    }
    for _, test := range tests {
        _, err := QuarterHourCount(test.ts1, test.ts2)
        assert.NotNil(t, err)
    }
}

func TestToTimestamp(t *testing.T) {
	loc, _ := time.LoadLocation("America/Los_Angeles")

	tests := []struct {
		in  time.Time
		out string
	}{
		{
			time.Unix(1546560000, 0),
			"2019-01-04T00:00:00Z",
		},
		{
			time.Date(2010, 12, 31, 0, 0, 0, 0, loc),
			"2010-12-31T08:00:00Z",
		},
	}
	for _, test := range tests {
		out := ToTimestamp(test.in)
		assert.Equal(t, test.out, out)
	}
}
