package calculate

import (
	"strconv"
	"time"
)

func ConvertStringToInt(text string) int {
	retVal, _ := strconv.Atoi(text)

	return retVal
}

func ConvertToDatetime(layout string, date string) (time.Time, error) {
	return time.Parse(layout, date)
}
