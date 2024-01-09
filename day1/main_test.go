package day1

import (
	"testing"

	"github.com/philipszalla/adventofcode-2023/utils"
)

var example = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

func TestPart1(t *testing.T) {
	utils.TestPart(t, part1, example, 142)
}

var example2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func TestPart2(t *testing.T) {
	utils.TestPart(t, part2, example2, 281)
}
