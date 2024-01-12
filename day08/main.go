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

func euclid(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func leastCommonMultiple(a, b int) int {
	return (a * b) / euclid(a, b)
}

func part2(lines []string) int {
	schema, mapDict := parse(lines)

	starts := []string{}
	for start := range mapDict {
		if strings.HasSuffix(start, "A") {
			starts = append(starts, start)
		}
	}
	c := make(chan int, len(starts))

	for index, position := range starts {
		go func(index int, position string) {
			count := 0
			for !strings.HasSuffix(position, "Z") {
				directionIndex := count % len(schema)
				direction := schema[directionIndex]

				position = getNextPosition(mapDict, position, direction)
				starts[index] = position

				count++
			}

			c <- count
		}(index, position)
	}

	counts := make([]int, len(starts))
	for index := range counts {
		counts[index] = <-c
	}

	sum := counts[0]

	if len(counts) > 1 {
		for _, count := range counts[1:] {
			sum = leastCommonMultiple(sum, count)
		}
	}

	return sum
}
