package day5

import (
	"fmt"
	"math"
	"os"
	"slices"
	"sync"
	"time"
)

func Run() {
	filename := "day5/example.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	fmt.Println("Loading file", filename)
	seeds, mappers := ReadFile(filename)

	// Part 1
	fmt.Println("Starting part 1...")

	start := time.Now()

	lowestValue := part1(seeds, mappers)

	end := time.Now()
	elapsed := end.Sub(start)

	fmt.Printf("Finished processing! Result: %d, Elapsed time: %s\n", lowestValue, elapsed)

	// Part 2
	fmt.Println("Starting part 2...")

	start = time.Now()

	lowestValue = part2(seeds, mappers)

	end = time.Now()
	elapsed = end.Sub(start)

	fmt.Printf("Finished processing! Result: %d, Elapsed time: %s\n", lowestValue, elapsed)
}

func part2(seeds []int, mappers []Mapper) int {
	if len(seeds)%2 != 0 {
		panic("Seed count must be even")
	}

	fmt.Printf("|------------|------------|\n")
	fmt.Printf("| Seed start | Result     |\n")
	fmt.Printf("|------------|------------|\n")

	results := []int{}

	var wg sync.WaitGroup
	wg.Add(len(seeds) / 2)

	i := 0
	for i < len(seeds) {
		results = append(results, -1)

		go func(i int, seeds []int, mappers []Mapper) {
			results[i/2] = processSeedRange(seeds[i], seeds[i+1], mappers)
			wg.Done()
		}(i, seeds, mappers)

		i += 2
	}

	wg.Wait()

	fmt.Printf("|------------|------------|\n")

	lowestValue := slices.Min(results)
	return lowestValue
}

func processSeedRange(seedStart int, length int, mappers []Mapper) int {
	lowestValue := math.Inf(0)

	i := 0
	for i < length {
		value := processSeed(seedStart+i, mappers)
		lowestValue = math.Min(float64(lowestValue), float64(value))

		i++
	}

	fmt.Printf("| %10d | %10.0f |\n", seedStart, lowestValue)

	return int(lowestValue)
}

func part1(seeds []int, mappers []Mapper) int {
	lowestValue := math.Inf(0)

	for _, seed := range seeds {
		value := processSeed(seed, mappers)
		lowestValue = math.Min(float64(lowestValue), float64(value))
	}

	return int(lowestValue)
}

func processSeed(seed int, mappers []Mapper) int {
	// fmt.Printf("--- processSeed %d ---\n", seed)

	currentValue := seed

	for _, mapper := range mappers {
		currentValue = processMapper(currentValue, mapper)
	}

	return currentValue
}

func processMapper(input int, mapper Mapper) int {
	// fmt.Printf("--- processMapper %d ---\n", input)

	for _, rule := range mapper.rules {
		diff := input - rule.sourceStart

		if diff >= 0 && diff < rule.length {
			return rule.destinationStart + diff
		}
	}

	return input
}
