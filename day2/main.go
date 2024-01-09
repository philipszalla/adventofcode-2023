package day2

import (
	"fmt"
	"os"

	"github.com/philipszalla/adventofcode-2023/utils"
)

func Run() {
	filepath := ""
	if len(os.Args) > 2 {
		filepath = os.Args[2]
	}

	fmt.Println("Loading file", filepath)
	lines := utils.ReadFile(filepath)

	utils.RunPart(part1, 1, lines)
	utils.RunPart(part2, 2, lines)
}

func part1(lines []string) int {
	games := Parse(lines)

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

func part2(lines []string) int {
	games := Parse(lines)

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
