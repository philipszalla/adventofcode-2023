package day3

import (
	"strings"
	"testing"
)

var example = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

var expectedResult = 4361

func TestPart1(t *testing.T) {
	data := strings.Split(example, "\n")

	result := part1(data)

	if result != expectedResult {
		t.Fatalf("Expected result is %d. But got %d", expectedResult, result)
	}
}

var expectedResult2 = 467835

func TestPart2(t *testing.T) {
	data := strings.Split(example, "\n")

	result := part2(data)

	if result != expectedResult2 {
		t.Fatalf("Expected result is %d. But got %d", expectedResult2, result)
	}
}
