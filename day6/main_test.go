package day6

import (
	"testing"

	"github.com/philipszalla/adventofcode-2023/utils"
)

var example = `Time:      7  15   30
Distance:  9  40  200`

func TestPart1(t *testing.T) {
	utils.TestPart(t, part1, example, 288)
}

func TestPart2(t *testing.T) {
	utils.TestPart(t, part2, example, 71503)
}
