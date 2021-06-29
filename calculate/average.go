package calculate

import (
	"errors"
	"fmt"
)

type AverageMinutesCalculator interface{
	AverageMinutes(finishTimes []string) (int, error)
}

type AverageMinuteService struct {}

func (s *AverageMinuteService) AverageMinutes(finishTimes []string) (int, error) {
	if len(finishTimes) == 0 { return 0, errors.New("finishTimes cannot be empty") }
	if len(finishTimes) > 50 { return 0, errors.New("Cannot have more than 50 finishTimes") }

	var durations []int64
	startDatetime, err := ConvertToDatetime(DateLayout, StartDate)
	if err != nil { return 0, errors.New(fmt.Sprintf("StartDate is invalid: %v", StartDate)) }

	for _, t := range finishTimes {
		if isValid := ValidateFormat(t); isValid != true {
			return 0, errors.New(fmt.Sprintf("finishTime is invalid: %v", t))
		}

		durationInMinutes := GetDurationInMinutes(startDatetime, t)
		durations = append(durations, durationInMinutes)
	}

	return int(GetAverageFromDurations(durations)), nil
}