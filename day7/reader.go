package day7

import (
	"slices"
	"strconv"
	"strings"
)

func countCards(cards string) ([]rune, []int) {
	cardTypes := []rune{}
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

	return cardTypes, cardCounts
}

func useJoker(cardTypes []rune, cardCounts []int) ([]rune, []int) {
	jokerIndex := slices.Index(cardTypes, 'J')
	if jokerIndex == -1 || len(cardTypes) == 1 {
		return cardTypes, cardCounts
	}

	// Remove joker from cards before searching for card with most appearances
	jokerCount := cardCounts[jokerIndex]
	cardTypes = append(cardTypes[:jokerIndex], cardTypes[jokerIndex+1:]...)
	cardCounts = append(cardCounts[:jokerIndex], cardCounts[jokerIndex+1:]...)

	// Determine card with most apperances and add jokers
	maxSameCardCount := slices.Max(cardCounts)
	maxSameCardIndex := slices.Index(cardCounts, maxSameCardCount)
	cardCounts[maxSameCardIndex] += jokerCount

	return cardTypes, cardCounts
}

func getWeight(cardTypes []rune, cardCounts []int) int {
	weight := 0

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

	return weight
}

func parse(lines []string, part2 bool) []Hand {
	hands := []Hand{}

	for _, line := range lines {
		parts := strings.Split(line, " ")

		cards := parts[0]
		bid, _ := strconv.Atoi(parts[1])

		cardTypes, cardCounts := countCards(cards)

		if part2 {
			cardTypes, cardCounts = useJoker(cardTypes, cardCounts)
		}

		weight := getWeight(cardTypes, cardCounts)

		hand := Hand{cards, bid, weight}

		hands = append(hands, hand)
	}

	return hands
}
