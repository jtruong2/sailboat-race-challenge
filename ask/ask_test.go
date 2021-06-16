package ask

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAskForFinishTimes(t *testing.T) {
	var stdin bytes.Buffer
	var finishTime string
	stdin.Write([]byte("02:00 PM, DAY 19\n"))

	as := AskService{}
	as.AskForFinishTimes(&stdin, &finishTime)

	assert.Equal(t, finishTime, "02:00 PM, DAY 19")
}

func TestAskToAddMoreTimes(t *testing.T) {
	var stdin bytes.Buffer
	var addMore string
	stdin.Write([]byte("n\n"))

	as := AskService{}
	as.AskToAddMoreTimes(&stdin, &addMore)

	assert.Equal(t, addMore, "n")
}
