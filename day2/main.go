package day2

import (
	"fmt"
	"os"
	"time"
)

func Run() {
	filepath := ""
	// filepath2 := ""
	if len(os.Args) > 2 {
		filepath = os.Args[2]
	}

	fmt.Println("Loading file", filepath)
	games := ReadFile(filepath)

	// fmt.Println("Loading file", filepath2)
	// lines2 := ReadFile(filepath2)

	// Part 1
	fmt.Println("Starting part 1...")

	start := time.Now()

	result := part1(games)

	end := time.Now()
	elapsed := end.Sub(start)

	fmt.Printf("Finished processing! Result: %d, Elapsed time: %s\n", result, elapsed)

	// Part 2
	// fmt.Println("Starting part 2...")

	// start = time.Now()

	// result = part2(lines2)

	// end = time.Now()
	// elapsed = end.Sub(start)

	// fmt.Printf("Finished processing! Result: %d, Elapsed time: %s\n", result, elapsed)
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
