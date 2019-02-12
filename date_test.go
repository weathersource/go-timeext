package timeext

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDate(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"1950-01-01", "1950-01-01"},
		{"1969-12-31", "1969-12-31"},
		{"1970-01-01", "1970-01-01"},
		{"2019-01-01", "2019-01-01"},
	}
	for _, test := range tests {
		out, err := Date(test.in)
		assert.Nil(t, err)
		assert.Equal(t, test.out, out.Format(dateFormat))
	}
}

func TestDateError(t *testing.T) {
	tests := []struct {
		d string
	}{
		{"20190-01"},
		{"2019-1-1"},
		{"2019-01-0"},
		{"2019-01-32"},
		{"2019-01-32A"},
		{""},
	}
	for _, test := range tests {
		_, err := Date(test.d)
		assert.NotNil(t, err)
	}
}

func TestDayCount(t *testing.T) {
	tests := []struct {
		d1, d2   string
		expected int
	}{
		{"2019-01-01", "2019-01-01", 1},
		{"2019-01-01", "2019-01-03", 3},
	}
	for _, test := range tests {
		count, err := DayCount(test.d1, test.d2)
		assert.Nil(t, err)
		assert.Equal(t, test.expected, count)
	}
}

func TestDayCountError(t *testing.T) {
	tests := []struct {
		d1, d2 string
	}{
		{"2019-01-01A", "2019-01-03"},
		{"2019-01-01", "2019-01-03A"},
		{"2019-01-03", "2019-01-01"},
	}
	for _, test := range tests {
		_, err := DayCount(test.d1, test.d2)
		assert.NotNil(t, err)
	}
}

func TestToDate(t *testing.T) {
	tests := []struct {
		in  time.Time
		out string
	}{
		{
			time.Unix(1546560000, 0),
			"2019-01-04",
		},
		{
			time.Date(2010, 12, 31, 0, 0, 0, 0, time.Local),
			"2010-12-31",
		},
	}
	for _, test := range tests {
		out := ToDate(test.in)
		assert.Equal(t, test.out, out)
	}
}

func TestDateYear(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			"2019-01-04",
			2019,
		},
		{
			"2010-12-31",
			2010,
		},
	}
	for _, test := range tests {
		out, err := DateYear(test.in)
		assert.Nil(t, err)
		assert.Equal(t, test.out, out)
	}

	testerrors := []struct {
		in string
	}{
		{
			"201",
		},
		{
			"201-01-04",
		},
	}
	for _, test := range testerrors {
		_, err := DateYear(test.in)
		assert.NotNil(t, err)
	}
}
