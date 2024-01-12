package day08

import (
	"fmt"
	"os"
	"slices"

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
}

func part1(lines []string) int {
	schema, starts, destsLeft, destsRight := parse(lines)
	position := "AAA"

	count := 0
	for position != "ZZZ" {
		positionIndex := slices.Index(starts, position)

		directionIndex := count % len(schema)
		direction := schema[directionIndex]

		if direction == 'L' {
			position = destsLeft[positionIndex]
		} else {
			position = destsRight[positionIndex]
		}

		count++
	}

	return count
}
