package day5

import (
	"fmt"
)

func Run() {
	seeds, mappers := ReadFile("day5/input.txt")

	fmt.Println(seeds, mappers)

	lowestValue := -1
	for index, seed := range seeds {
		value := processSeed(seed, mappers)
		if index == 0 || value < lowestValue {
			lowestValue = value
		}
	}

	fmt.Println("Lowest value:", lowestValue)
}

func processSeed(seed int, mappers []Mapper) int {
	fmt.Printf("--- processSeed %d ---\n", seed)

	currentValue := seed

	for _, mapper := range mappers {
		currentValue = processMapper(currentValue, mapper)
	}

	return currentValue
}

func processMapper(input int, mapper Mapper) int {
	fmt.Printf("--- processMapper %d ---\n", input)

	for _, rule := range mapper.rules {
		diff := input - rule.sourceStart

		if diff >= 0 && diff < rule.length {
			return rule.destinationStart + diff
		}
	}

	return input
}
