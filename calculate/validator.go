package calculate

import "regexp"

/*
	Validates that the finishTime is formatted in hh:mm xM, DAY n
*/
func ValidateFormat(t string) bool {
	isValid, _ := regexp.MatchString(
		"^\\b(0?[1-9]|[1][0-2]):(0?[0-9]|[1-5][0-9])\\s[A|P]M,\\sDAY\\s(0?[1-9]|[1-9][0-9])$\\b",
		t,
	)
	return isValid
}
