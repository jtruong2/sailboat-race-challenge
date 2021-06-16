package handler

import (
	"fmt"
	"html"
	"os"
	"strconv"
	"strings"

	"github.com/jtruong2/sailboat-race/ask"
	"github.com/jtruong2/sailboat-race/calculate"
)

const (
	startDate    string = "2021-01-01T08:00:00"
)

func HandleInputs(asker ask.Asker, calculator calculate.Calculator) {
	var finishTimes []string
	handleTimesInput(asker, calculator, finishTimes)
}

func handleTimesInput(asker ask.Asker, calculator calculate.Calculator, finishTimes []string) {
	var finishTime string
	asker.AskForFinishTimes(os.Stdin, &finishTime)

	finishTimes = append(finishTimes, strings.TrimSpace(finishTime))
	handleContinueInput(asker, calculator, finishTimes)
}

func handleContinueInput(asker ask.Asker, calculator calculate.Calculator, finishTimes []string) {
	var addMore string
	asker.AskToAddMoreTimes(os.Stdin, &addMore)

	switch addMore {
	case "y", "yes":
		handleTimesInput(asker, calculator, finishTimes)
	case "n", "no":
		averageDuration, err := calculator.AverageMinutes(startDate, finishTimes)
		if err != nil {
			logWarning(fmt.Sprintf("There was an error processing your request: %v\n", err))
			break
		}
		logSuccess(averageDuration)
	default:
		logWarning("Invalid input. Please try again.\n")
		handleContinueInput(asker, calculator, finishTimes)
	}
}

func logSuccess(averageDuration int) {
	const successPrint string = "\033[1;32m%s\033[0m"

	fmt.Printf(successPrint, fmt.Sprintf("The average duration is %v minutes %v\n", averageDuration, boatMoji()))
}

func logWarning(message string) {
	const warningPrint string = "\033[1;33m%s\033[0m"

	fmt.Printf(warningPrint, message)
}

func boatMoji() string {
	return html.UnescapeString(
		"&#" + strconv.Itoa(128675) + ";&#" + strconv.Itoa(128674) + ";&#" + strconv.Itoa(128676) + ";",
	)
}
