package day2

import (
	"fmt"
	"os"
	"time"
)

func Run() {
	filepath := ""
	if len(os.Args) > 2 {
		filepath = os.Args[2]
	}

	fmt.Println("Loading file", filepath)
	games := ReadFile(filepath)

	// Part 1
	fmt.Println("Starting part 1...")

	start := time.Now()

	result := part1(games)

	end := time.Now()
	elapsed := end.Sub(start)

	fmt.Printf("Finished processing! Result: %d, Elapsed time: %s\n", result, elapsed)

	// Part 2
	fmt.Println("Starting part 2...")

	start = time.Now()

	result = part2(games)

	end = time.Now()
	elapsed = end.Sub(start)

	fmt.Printf("Finished processing! Result: %d, Elapsed time: %s\n", result, elapsed)
}

func part1(games []Game) int {
	existingCubes := CubeSet{12, 13, 14}

	sum := 0

	for _, game := range games {
		fmt.Printf("game.id: %v\n", game.id)

		possible := true
		for _, set := range game.sets {
			currentCubes := existingCubes

			currentCubes.red -= set.red
			currentCubes.green -= set.green
			currentCubes.blue -= set.blue

			if currentCubes.red < 0 || currentCubes.green < 0 || currentCubes.blue < 0 {
				possible = false
				fmt.Println("not possible")
				break
			}
		}

		if !possible {
			continue
		}

		sum += game.id
	}

	return sum
}

func part2(games []Game) int {
	sum := 0

	for _, game := range games {
		fmt.Printf("game.id: %v\n", game.id)

		minimum := CubeSet{0, 0, 0}

		for _, set := range game.sets {
			fmt.Printf("game.sets: %v\n", set)

			if minimum.red < set.red {
				minimum.red = set.red
			}

			if minimum.green < set.green {
				minimum.green = set.green
			}

			if minimum.blue < set.blue {
				minimum.blue = set.blue
			}
		}

		fmt.Printf("minimum: %v\n", minimum)

		product := minimum.red * minimum.green * minimum.blue
		fmt.Printf("product: %v\n", product)

		sum += product
	}

	return sum
}
