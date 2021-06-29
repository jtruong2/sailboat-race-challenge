package main

import (

	a "github.com/jtruong2/sailboat-race/ask"
	c "github.com/jtruong2/sailboat-race/calculate"
	"github.com/jtruong2/sailboat-race/handler"
)

func main() {
	var asker a.Service
	var calculator c.AverageMinuteService

	handler.HandleInputs(&asker, &calculator)
}
