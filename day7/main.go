package day7

import (
	"fmt"
	"os"
	"slices"
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
	utils.RunPart(part2, 2, lines)
}

var cardValues = []rune{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
var cardValues2 = []rune{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}

func sortHands(hands []Hand, i, j int, values []rune) bool {
	a := hands[i]
	b := hands[j]

	// Sort by weight
	if a.weight != b.weight {
		return a.weight < b.weight
	}

	// Sort by first higher card
	aCards := []rune(a.cards)
	bCards := []rune(b.cards)
	for index, aCard := range aCards {
		bCard := bCards[index]

		aValue := slices.Index(values, aCard)
		bValue := slices.Index(values, bCard)

		if aValue == bValue {
			continue
		}

		return aValue < bValue
	}

	return false
}

func sum(hands []Hand) int {
	sum := 0
	for index, hand := range hands {
		sum += hand.bid * (index + 1)
	}

	return sum
}

func part1(lines []string) int {
	hands := parse(lines, false)

	sort.SliceStable(hands, func(i, j int) bool {
		return sortHands(hands, i, j, cardValues)
	})

	return sum(hands)
}

func part2(lines []string) int {
	hands := parse(lines, true)

	sort.SliceStable(hands, func(i, j int) bool {
		return sortHands(hands, i, j, cardValues2)
	})

	return sum(hands)
}
