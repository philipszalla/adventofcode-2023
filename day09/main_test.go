package day09

import (
	"testing"

	"github.com/philipszalla/adventofcode-2023/utils"
)

func TestProcessLine(t *testing.T) {
	result := processLine("0 3 6 9 12 15")

	if result != 18 {
		t.Fatalf("Expected %v but got %v", 18, result)
	}
}

var example = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func TestPart1(t *testing.T) {
	utils.TestPart(t, part1, example, 114)
}
