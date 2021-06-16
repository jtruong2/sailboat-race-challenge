package calculate

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	dateLayout string = "2006-01-02T15:04:05"
	timeFormat string = "15:04"
	timeLayout string = time.Kitchen
)

type Calculator interface{
	AverageMinutes(startDate string, finishTimes []string) (int, error)
}

type CalculateService struct {}

func (cs *CalculateService) AverageMinutes(startDate string, finishTimes []string) (int, error) {
	if len(finishTimes) == 0 { return 0, errors.New("finishTimes cannot be empty") }
	if len(finishTimes) > 50 { return 0, errors.New("Cannot have more than 50 finishTimes") }

	var durations []int64
	startDatetime, err := convertToDatetime(dateLayout, startDate)
	if err != nil { return 0, errors.New(fmt.Sprintf("startDate is invalid: %v", startDate)) }

	for _, t := range finishTimes {
		if isValid := validateFormat(t); isValid != true {
			return 0, errors.New(fmt.Sprintf("finishTime is invalid: %v", t))
		}

		durationInMinutes := getDurationInMinutes(startDatetime, t)
		durations = append(durations, durationInMinutes)
	}

	return int(getAverageFromDurations(durations)), nil
}

func getDurationInMinutes(startDatetime time.Time, t string) int64 {
		day := extractDay(t)
		hour, minute := extractTime(t)

		finishDate := startDatetime.AddDate(0, 0, day-1)
		finishTimeInSeconds := time.Date(
			finishDate.Year(),
			finishDate.Month(),
			finishDate.Day(),
			hour,
			minute,
			0,
			0,
			finishDate.Location(),
		).Unix()

		return (finishTimeInSeconds - startDatetime.Unix()) / 60
}

func getAverageFromDurations(durations []int64) float64 {
	var sum int64 = 0
	for i := 0; i < len(durations); i++ {
		sum += (durations[i])
	}
	avg := (float64(sum) / (float64(len(durations))))

	return math.Round(avg)
}

func extractDay(t string) int {
	return convertStringToInt(strings.TrimSpace(strings.Split(t, "DAY")[1]))
}

func extractTime(t string) (int, int) {
	unformattedTime := strings.ReplaceAll(strings.Split(t, ",")[0], " ", "")
	parsedTime, _ := convertToDatetime(timeLayout, unformattedTime)
	formattedTime := strings.Split(parsedTime.Format(timeFormat), ":")
	hour := convertStringToInt(formattedTime[0])
	minute := convertStringToInt(formattedTime[1])

	return hour, minute
}

func convertStringToInt(text string) int {
	retVal, _ := strconv.Atoi(text)

	return retVal
}

func convertToDatetime(layout string, date string) (time.Time, error) {
	return time.Parse(layout, date)
}

/*
	Validates that the time is formatted in hh:mm xM, DAY n
*/
func validateFormat(t string) bool {
	isValid, _ := regexp.MatchString(
		"^\\b(0?[1-9]|[1][0-2]):(0?[0-9]|[1-5][0-9])\\s[A|P]M,\\sDAY\\s(0?[1-9]|[1-9][0-9])$\\b",
		t,
	)
	return isValid
}
