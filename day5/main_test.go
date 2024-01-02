package day5

import "testing"

var example = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

var expectedResult = 35

func TestPart1(t *testing.T) {
	seeds, mappers := parse(example)

	result := part1(seeds, mappers)

	if result != expectedResult {
		t.Fatalf("Expected result is %d. But got %d", expectedResult, result)
	}
}

var expectedResult2 = 46

func TestPart2(t *testing.T) {
	seeds, mappers := parse(example)

	result := part2(seeds, mappers)

	if result != expectedResult2 {
		t.Fatalf("Expected result is %d. But got %d", expectedResult2, result)
	}
}