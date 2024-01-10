package day7

import (
	"slices"
	"strconv"
	"strings"
)

var cardValues = []rune{'_', '_', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}

func parse(lines []string) []Hand {
	hands := []Hand{}

	for _, line := range lines {
		parts := strings.Split(line, " ")

		cards := []int{}
		for _, cardType := range parts[0] {
			cards = append(cards, slices.Index(cardValues, cardType))
		}

		bid, _ := strconv.Atoi(parts[1])

		weight := 0

		cardTypes := []int{}
		cardCounts := []int{}
		for _, cardType := range cards {
			cardTypeIndex := slices.Index(cardTypes, cardType)
			if cardTypeIndex == -1 {
				cardTypeIndex = len(cardTypes)

				cardTypes = append(cardTypes, cardType)
				cardCounts = append(cardCounts, 0)
			}

			cardCounts[cardTypeIndex]++
		}

		switch len(cardTypes) {
		case 1:
			// Five of a kind
			weight = 70
		case 2:
			max := slices.Max(cardCounts)
			if max == 4 {
				// Four of a kind
				weight = 60
			} else {
				// Full house
				weight = 50
			}
		case 3:
			max := slices.Max(cardCounts)
			if max == 3 {
				// Three of a kind
				weight = 40
			} else {
				// Two pair
				weight = 30
			}
		case 4:
			// One pair
			weight = 20
		case 5:
			// High card
			weight = 10
		}

		hand := Hand{parts[0], cards, bid, weight}

		hands = append(hands, hand)
	}

	return hands
}
