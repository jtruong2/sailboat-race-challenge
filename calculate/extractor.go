package calculate

import "strings"

func ExtractDay(t string) int {
	return ConvertStringToInt(strings.TrimSpace(strings.Split(strings.ToUpper(t), "DAY")[1]))
}

func ExtractTime(t string) (int, int) {
	unformattedTime := strings.ReplaceAll(strings.Split(strings.ToUpper(t), ",")[0], " ", "")
	parsedTime, _ := ConvertToDatetime(TimeLayout, unformattedTime)
	formattedTime := strings.Split(parsedTime.Format(TimeFormat), ":")
	hour := ConvertStringToInt(formattedTime[0])
	minute := ConvertStringToInt(formattedTime[1])

	return hour, minute
}
