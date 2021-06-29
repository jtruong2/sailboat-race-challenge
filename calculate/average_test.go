package calculate

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

const startDate string = "2021-01-01T08:00:00"

func TestAverageMinutes(t *testing.T) {
	cases := []struct {
		name 		  string
		startDate     string
		finishTimes   []string
		expected      int
	} {
		{
			name: "when retVal ends in .5",
			finishTimes: []string{"12:00 PM, DAY 1", "12:01 PM, DAY 1"},
			startDate: startDate,
			expected: 241,
		},
		{
			name: "when one finish time",
			startDate: startDate,
			finishTimes: []string{"12:00 AM, DAY 2"},
			expected: 960,
		},
		{
			name: "when multiple finish times",
			startDate: startDate,
			finishTimes: []string{"02:00 PM, DAY 19", "02:00 PM, DAY 20", "01:58 PM, DAY 20"},
			expected: 27239,
		},
		{
			name: "when finished on day 99",
			startDate: startDate,
			finishTimes: []string{"01:58 PM, DAY 99"},
			expected: 141478,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cs := Service{}
			output, err := cs.AverageMinutes(tc.startDate, tc.finishTimes)
			if err != nil { t.Logf(`There was an error with spec TestAverageMinutes: %v`, err) }

			assert.Equal(t, tc.expected, output)
		})
	}
}

func TestAverageMinutesSadPath(t *testing.T) {
	cases := []struct {
		name  	 			string
		startDate     string
		finishTimes   []string
		err           error
	} {
		{
			name: "when startDate is invalid",
			startDate: "021-1-01T08:00",
			finishTimes: []string{"12:00 PM, DAY 1"},
			err: errors.New("startDate is invalid: 021-1-01T08:00"),
		},
		{
			name: "when finishTime is invalid",
			startDate: startDate,
			finishTimes: []string{"12:00 PM, DAY 1 fdafdsk"},
			err: errors.New("finishTime is invalid: 12:00 PM, DAY 1 fdafdsk"),
		},
		{
			name: "when finishTimes contains more than 50 elements",
			startDate: startDate,
			finishTimes: []string{
				"12:00 PM, DAY 1", "12:00 PM, DAY 2", "12:00 PM, DAY 3", "12:00 PM, DAY 4", "12:00 PM, DAY 5", "12:00 PM, DAY 6",
				"12:00 PM, DAY 7", "12:00 PM, DAY 8", "12:00 PM, DAY 9", "12:00 PM, DAY 10", "12:00 PM, DAY 11", "12:00 PM, DAY 12", "12:00 PM, DAY 13",
				"12:00 PM, DAY 14", "12:00 PM, DAY 15", "12:00 PM, DAY 16", "12:00 PM, DAY 17", "12:00 PM, DAY 18", "12:00 PM, DAY 19", "12:00 PM, DAY 20",
				"12:00 PM, DAY 21", "12:00 PM, DAY 22", "12:00 PM, DAY 23", "12:00 PM, DAY 24", "12:00 PM, DAY 25", "12:00 PM, DAY 26", "12:00 PM, DAY 27",
				"12:00 PM, DAY 28", "12:00 PM, DAY 29", "12:00 PM, DAY 30", "12:00 PM, DAY 31", "12:00 PM, DAY 32", "12:00 PM, DAY 33", "12:00 PM, DAY 34",
				"12:00 PM, DAY 35", "12:00 PM, DAY 36", "12:00 PM, DAY 37", "12:00 PM, DAY 38", "12:00 PM, DAY 39", "12:00 PM, DAY 40", "12:00 PM, DAY 41",
				"12:00 PM, DAY 42", "12:00 PM, DAY 43", "12:00 PM, DAY 44", "12:00 PM, DAY 45", "12:00 PM, DAY 46", "12:00 PM, DAY 47", "12:00 PM, DAY 48",
				"12:00 PM, DAY 49", "12:00 PM, DAY 50", "12:00 PM, DAY 51",
			},
			err: errors.New("Cannot have more than 50 finishTimes"),
		},
		{
			name: "when finishTimes is empty",
			startDate: startDate,
			finishTimes: []string{},
			err: errors.New("finishTimes cannot be empty"),
		},
		{
			name: "when hour is invalid",
			startDate: startDate,
			finishTimes: []string{"13:00 PM, DAY 1"},
			err: errors.New("finishTime is invalid: 13:00 PM, DAY 1"),
		},
		{
			name: "when minute is invalid",
			startDate: startDate,
			finishTimes: []string{"12:60 PM, DAY 1"},
			err: errors.New("finishTime is invalid: 12:60 PM, DAY 1"),
		},
		{
			name: "when day is more than 99",
			startDate: startDate,
			finishTimes: []string{"12:00 PM, DAY 100"},
			err: errors.New("finishTime is invalid: 12:00 PM, DAY 100"),
		},
		{
			name: "when day is less than 1",
			startDate: startDate,
			finishTimes: []string{"12:00 PM, DAY 0"},
			err: errors.New("finishTime is invalid: 12:00 PM, DAY 0"),
		},
		{
			name: "when finishTime is before the startDate",
			startDate: startDate,
			finishTimes: []string{"4:00 AM, DAY 0"},
			err: errors.New("finishTime is invalid: 4:00 AM, DAY 0"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cs := Service{}
			_, err := cs.AverageMinutes(tc.startDate, tc.finishTimes)

			assert.NotNil(t, err)
			assert.Equal(t, tc.err, err)
		})
	}
}
