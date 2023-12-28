package day1

import (
	"fmt"
	"os"
	"slices"
	"time"
)

func Run() {
	filepath := ""
	if len(os.Args) > 2 {
		filepath = os.Args[2]
	}

	fmt.Println("Loading file", filepath)
	lines := ReadFile(filepath)

	// Part 1
	fmt.Println("Starting part 1...")

	start := time.Now()

	result := part1(lines)

	end := time.Now()
	elapsed := end.Sub(start)

	fmt.Printf("Finished processing! Result: %d, Elapsed time: %s\n", result, elapsed)
}

var asciiCodeZero = 48

func part1(lines []string) int {
	sum := 0

	for _, line := range lines {
		value := 0

		// First digit
		for _, char := range line {
			code := int(char) - asciiCodeZero
			if code >= 0 && code < 10 {
				value += code * 10
				break
			}
		}

		chars := []rune(line)
		slices.Reverse(chars)

		// Second digit
		for _, char := range chars {
			code := int(char) - asciiCodeZero
			if code >= 0 && code < 10 {
				value += code
				break
			}
		}

		sum += value
	}

	return sum
}
