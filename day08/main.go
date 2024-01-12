package day08

import (
	"fmt"
	"os"
	"strings"

	"github.com/philipszalla/adventofcode-2023/utils"
)

type Dir struct {
	left  string
	right string
}

type MapDict map[string]Dir

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

func getNextPosition(mapDict MapDict, position string, direction rune) string {
	dict := mapDict[position]

	if direction == 'L' {
		position = dict.left
	} else {
		position = dict.right
	}

	return position
}

func part1(lines []string) int {
	schema, mapDict := parse(lines)
	position := "AAA"

	count := 0
	for position != "ZZZ" {
		directionIndex := count % len(schema)
		direction := schema[directionIndex]

		position = getNextPosition(mapDict, position, direction)

		count++
	}

	return count
}

func part2(lines []string) int {
	schema, mapDict := parse(lines)

	positions := []string{}
	for start := range mapDict {
		if strings.HasSuffix(start, "A") {
			positions = append(positions, start)
		}
	}

	count := 0
	for {
		directionIndex := count % len(schema)
		direction := schema[directionIndex]

		end := true
		for index, position := range positions {
			position = getNextPosition(mapDict, position, direction)
			positions[index] = position

			end = end && strings.HasSuffix(position, "Z")
		}

		count++

		if end {
			break
		}
	}

	return count
}
