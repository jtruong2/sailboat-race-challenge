package calculate

import (
	"errors"
	"fmt"
	"time"
)

const (
	dateLayout string = "2006-01-02T15:04:05"
	timeFormat string = "15:04"
	timeLayout string = time.Kitchen
)

type AverageMinutesCalculator interface{
	AverageMinutes(startDate string, finishTimes []string) (int, error)
}

type Service struct {}

func (s *Service) AverageMinutes(startDate string, finishTimes []string) (int, error) {
	if len(finishTimes) == 0 { return 0, errors.New("finishTimes cannot be empty") }
	if len(finishTimes) > 50 { return 0, errors.New("Cannot have more than 50 finishTimes") }

	var durations []int64
	startDatetime, err := ConvertToDatetime(dateLayout, startDate)
	if err != nil { return 0, errors.New(fmt.Sprintf("startDate is invalid: %v", startDate)) }

	for _, t := range finishTimes {
		if isValid := ValidateFormat(t); isValid != true {
			return 0, errors.New(fmt.Sprintf("finishTime is invalid: %v", t))
		}

		durationInMinutes := GetDurationInMinutes(startDatetime, t)
		durations = append(durations, durationInMinutes)
	}

	return int(GetAverageFromDurations(durations)), nil
}