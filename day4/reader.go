package day4

import (
	"strconv"
	"strings"
)

func parseLine(line string) (int, []int, []int) {
	winning := []int{}
	yours := []int{}

	if !strings.HasPrefix(line, "Card ") {
		return 0, winning, yours
	}

	line = strings.TrimPrefix(line, "Card ")

	cardParts := strings.Split(line, ": ")
	card, _ := strconv.Atoi(cardParts[0])

	values := strings.Split(cardParts[1], " | ")

	winningParts := strings.Split(values[0], " ")
	for _, part := range winningParts {
		value, err := strconv.Atoi(part)
		if err != nil {
			continue
		}

		winning = append(winning, value)
	}

	yoursParts := strings.Split(values[1], " ")
	for _, part := range yoursParts {
		value, err := strconv.Atoi(part)
		if err != nil {
			continue
		}

		yours = append(yours, value)
	}

	return card, winning, yours
}
