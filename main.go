package main

import (

	a "github.com/jtruong2/sailboat-race/ask"
	c "github.com/jtruong2/sailboat-race/calculate"
	"github.com/jtruong2/sailboat-race/handler"
)

func main() {
	var asker a.AskService
	var calculator c.CalculateService

	handler.HandleInputs(&asker, &calculator)
}
