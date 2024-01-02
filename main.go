package main

import (
	"os"
	"strconv"

	"github.com/philipszalla/adventofcode-2023/day1"
	"github.com/philipszalla/adventofcode-2023/day2"
	"github.com/philipszalla/adventofcode-2023/day5"
)

func main() {
	selectedDay := 0
	if len(os.Args) > 1 {
		selectedDay, _ = strconv.Atoi(os.Args[1])
	}

	switch selectedDay {
	case 1:
		day1.Run()
	case 2:
		day2.Run()
	case 5:
		day5.Run()
	default:
		println("Unknown option")
	}

}
