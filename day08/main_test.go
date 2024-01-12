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

var example3 = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

func TestPart2(t *testing.T) {
	utils.TestPart(t, part2, example, 2)

	utils.TestPart(t, part2, example2, 6)

	utils.TestPart(t, part2, example3, 6)
}
