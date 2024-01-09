package day3

import (
	"testing"

	"github.com/philipszalla/adventofcode-2023/utils"
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

func TestPart1(t *testing.T) {
	utils.TestPart(t, part1, example, 4361)
}

func TestPart2(t *testing.T) {
	utils.TestPart(t, part2, example, 467835)
}
