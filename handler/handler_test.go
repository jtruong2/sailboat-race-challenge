package handler

import (
	"io"
	"testing"
	"errors"
)

type mockAskService struct {
	finishTimeInput string
	addMoreInput string
}

type mockCalculatorService struct {
	err error
}

func TestHandleInputs(t *testing.T) {
	cases := []struct{
		name string
		mas *mockAskService
		mcs *mockCalculatorService
	}{
		{
			name: "when addMore equals no",
			mas: &mockAskService{
				finishTimeInput: "12:00 PM, DAY 1",
				addMoreInput: "n",
			},
			mcs: &mockCalculatorService{
				err: nil,
			},
		},
		{
			name: "when error in AverageMinutes",
			mas: &mockAskService{
				finishTimeInput: "12:00 PM, DAY 1",
				addMoreInput: "n",
			},
			mcs: &mockCalculatorService{
				err: errors.New("failed"),
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			HandleInputs(tc.mas, tc.mcs)
		})
	}
}

func (mas *mockAskService) AskForFinishTimes(stdin io.Reader, finishTime *string) {
	*finishTime = mas.finishTimeInput
}

func (mas *mockAskService) AskToAddMoreTimes(stdin io.Reader, addMore *string) {
	*addMore = mas.addMoreInput
}

func (mcs *mockCalculatorService) AverageMinutes(finishTimes []string) (int, error) {
	return 0, mcs.err
}
