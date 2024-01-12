package day08

import (
	"testing"

	"github.com/philipszalla/adventofcode-2023/utils"
)

var example = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

var example2 = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

func TestPart1(t *testing.T) {
	utils.TestPart(t, part1, example, 2)

	utils.TestPart(t, part1, example2, 6)
}
