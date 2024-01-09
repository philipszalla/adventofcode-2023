package day2

import (
	"strings"
	"testing"
)

var example = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

var expectedResult = 8

func TestPart1(t *testing.T) {
	lines := strings.Split(example, "\n")
	result := part1(lines)

	if result != expectedResult {
		t.Fatalf("Expected result is %d. But got %d", expectedResult, result)
	}
}

var expectedResult2 = 2286

func TestPart2(t *testing.T) {
	lines := strings.Split(example, "\n")
	result := part2(lines)

	if result != expectedResult2 {
		t.Fatalf("Expected result is %d. But got %d", expectedResult2, result)
	}
}
