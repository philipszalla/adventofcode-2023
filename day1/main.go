package day1

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func Run() {
	filepath := ""
	filepath2 := ""
	if len(os.Args) > 3 {
		filepath = os.Args[2]
		filepath2 = os.Args[3]
	}

	fmt.Println("Loading file", filepath)
	lines := ReadFile(filepath)

	fmt.Println("Loading file", filepath2)
	lines2 := ReadFile(filepath2)

	// Part 1
	fmt.Println("Starting part 1...")

	start := time.Now()

	result := part1(lines)

	end := time.Now()
	elapsed := end.Sub(start)

	fmt.Printf("Finished processing! Result: %d, Elapsed time: %s\n", result, elapsed)

	// Part 2
	fmt.Println("Starting part 2...")

	start = time.Now()

	result = part2(lines2)

	end = time.Now()
	elapsed = end.Sub(start)

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

var dictionary = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func part2(lines []string) int {
	sum := 0

	for _, line := range lines {
		value := 0

		// First digit
		firstDigit := 0
		firstIndex := -1
		for index, char := range line {
			code := int(char) - asciiCodeZero
			if code >= 0 && code < 10 {
				firstDigit = code
				firstIndex = index
				break
			}
		}

		// Check for first word
		for wordValue, word := range dictionary {
			wordIndex := strings.Index(line, word)
			if wordIndex != -1 && (wordIndex < firstIndex || firstIndex == -1) {
				firstDigit = wordValue
				firstIndex = wordIndex
			}
		}

		value += firstDigit * 10

		chars := []rune(line)
		slices.Reverse(chars)

		// Second digit
		secondDigit := 0
		secondIndex := -1
		for index, char := range chars {
			code := int(char) - asciiCodeZero
			if code >= 0 && code < 10 {
				secondDigit = code
				secondIndex = index
				break
			}
		}
		if secondIndex != -1 {
			secondIndex = len(line) - secondIndex - 1
		}

		// Check for second word
		for wordValue, word := range dictionary {
			wordIndex := strings.LastIndex(line, word)
			if wordIndex != -1 && (wordIndex > secondIndex || secondIndex == -1) {
				secondDigit = wordValue
				secondIndex = wordIndex
			}
		}

		value += secondDigit

		fmt.Printf("%s -> %d\n", line, value)

		sum += value
	}

	return sum
}
