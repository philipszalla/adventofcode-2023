package day7

import (
	"fmt"
	"os"
	"sort"

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
}

func part1(lines []string) int {
	hands := parse(lines)

	sort.SliceStable(hands, func(i, j int) bool {
		a := hands[i]
		b := hands[j]

		// Sort by weight
		if a.weight != b.weight {
			return a.weight < b.weight
		}

		// Sort by first higher card
		for index, aCard := range a.cards {
			bCard := b.cards[index]

			if aCard == bCard {
				continue
			}

			return aCard < bCard
		}

		return false
	})

	sum := 0
	for index, hand := range hands {
		sum += hand.bid*index + 1
	}

	return sum
}
