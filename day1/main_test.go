package day1

import (
	"strings"
	"testing"
)

var example = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
var expectedResult = 142

func TestPart1(t *testing.T) {
	lines := strings.Split(example, "\n")

	result := part1(lines)

	if result != expectedResult {
		t.Fatalf("Result must be %d", expectedResult)
	}
}

var example2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
var expectedResult2 = 281

func TestPart2(t *testing.T) {
	lines := strings.Split(example2, "\n")

	result := part2(lines)

	if result != expectedResult2 {
		t.Fatalf("Result must be %d", expectedResult2)
	}
}
