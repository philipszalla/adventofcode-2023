package main

import (
	"os"
	"strconv"

	"github.com/philipszalla/adventofcode-2023/day01"
	"github.com/philipszalla/adventofcode-2023/day02"
	"github.com/philipszalla/adventofcode-2023/day03"
	"github.com/philipszalla/adventofcode-2023/day04"
	"github.com/philipszalla/adventofcode-2023/day05"
	"github.com/philipszalla/adventofcode-2023/day06"
	"github.com/philipszalla/adventofcode-2023/day07"
	"github.com/philipszalla/adventofcode-2023/day08"
	"github.com/philipszalla/adventofcode-2023/day09"
)

func main() {
	selectedDay := 0
	if len(os.Args) > 1 {
		selectedDay, _ = strconv.Atoi(os.Args[1])
	}

	switch selectedDay {
	case 1:
		day01.Run()
	case 2:
		day02.Run()
	case 3:
		day03.Run()
	case 4:
		day04.Run()
	case 5:
		day05.Run()
	case 6:
		day06.Run()
	case 7:
		day07.Run()
	case 8:
		day08.Run()
	case 9:
		day09.Run()
	default:
	}

}
