package timeext

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDate(t *testing.T) {
	tests := []struct {
		d string
	}{
		{"1950-01-01"},
		{"1969-12-31"},
		{"1970-01-01"},
		{"2019-01-01"},
	}
	for _, test := range tests {
		_, err := Date(test.d)
		assert.Nil(t, err)
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
		count, err := DayCount(test.d1, test.d2)
		if !assert.NotNil(t, err) {
			fmt.Println(count)
		}
	}
}
