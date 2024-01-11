package day04

import (
	"strconv"
	"strings"
)

func parseCard(line string) Card {
	winning := []int{}
	yours := []int{}

	if !strings.HasPrefix(line, "Card") {
		return Card{0, winning, yours}
	}

	line = strings.TrimPrefix(line, "Card")

	cardParts := strings.Split(line, ":")
	id, _ := strconv.Atoi(strings.TrimSpace(cardParts[0]))

	values := strings.Split(cardParts[1], "|")

	winningParts := strings.Split(strings.TrimSpace(values[0]), " ")
	for _, part := range winningParts {
		value, err := strconv.Atoi(part)
		if err != nil {
			continue
		}

		winning = append(winning, value)
	}

	yoursParts := strings.Split(strings.TrimSpace(values[1]), " ")
	for _, part := range yoursParts {
		value, err := strconv.Atoi(part)
		if err != nil {
			continue
		}

		yours = append(yours, value)
	}

	return Card{id, winning, yours}
}
