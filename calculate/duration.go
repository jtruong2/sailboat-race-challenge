package calculate

import (
	"math"
	"time"
)

func GetDurationInMinutes(startDatetime time.Time, t string) int64 {
	day := ExtractDay(t)
	hour, minute := ExtractTime(t)

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

func GetAverageFromDurations(durations []int64) float64 {
	var sum int64 = 0
	for i := 0; i < len(durations); i++ {
		sum += durations[i]
	}
	avg := float64(sum) / (float64(len(durations)))

	return math.Round(avg)
}
