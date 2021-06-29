package ask

import (
	"bufio"
	"fmt"
	"io"
)

type Asker interface {
	AskForFinishTimes(stdin io.Reader, finishTime *string)
	AskToAddMoreTimes(stdin io.Reader, addMore *string)
}

type Service struct {}

func (as *Service) AskForFinishTimes(stdin io.Reader, finishTime *string) {
	scanner := bufio.NewScanner(stdin)
	fmt.Println("Enter a finish time in the correct format (ex: 02:00 PM, DAY 19 )")
	scanner.Scan()
 	*finishTime = scanner.Text()
}

func (as *Service) AskToAddMoreTimes(stdin io.Reader, addMore *string) {
	scanner := bufio.NewScanner(stdin)
	fmt.Println("Would you like to add more finish times? (y/n)")
	scanner.Scan()
	*addMore = scanner.Text()
}
