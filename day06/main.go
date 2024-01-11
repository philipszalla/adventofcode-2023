package day06

import (
	"fmt"
	"os"
	"strings"

	"github.com/philipszalla/adventofcode-2023/utils"
)

func Run() {
	filepath := ""
	if len(os.Args) > 2 {
		filepath = os.Args[2]
	}

	fmt.Println("Loading file", filepath)
	lines := utils.ReadFile(filepath)

	utils.RunPart(part1, 1, lines)
	utils.RunPart(part2, 2, lines)
}

func calcSolutions(time int, distance int) int {
	solutions := 0

	for i := 1; i < time; i++ {
		newDistance := (time - i) * i

		if newDistance > distance {
			solutions++
		}
	}

	return solutions
}

func part1(lines []string) int {
	times, distances := parse(lines)

	result := 1
	for index, time := range times {
		distance := distances[index]

		result *= calcSolutions(time, distance)
	}

	return result
}

func part2(lines []string) int {
	lines[0] = strings.Replace(lines[0], " ", "", -1)
	lines[1] = strings.Replace(lines[1], " ", "", -1)

	times, distances := parse(lines)

	result := 1
	for index, time := range times {
		distance := distances[index]

		result *= calcSolutions(time, distance)
	}

	return result
}
