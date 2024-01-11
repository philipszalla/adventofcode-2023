package day07

import (
	"testing"

	"github.com/philipszalla/adventofcode-2023/utils"
)

var example = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func TestPart1(t *testing.T) {
	utils.TestPart(t, part1, example, 6440)
}

func TestPart2(t *testing.T) {
	utils.TestPart(t, part2, example, 5905)
}
