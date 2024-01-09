package day2

import (
	"strconv"
	"strings"
)

func Parse(lines []string) []Game {

	games := []Game{}

	for _, line := range lines {
		splitLine := strings.Split(line, ": ")
		if len(splitLine) != 2 {
			continue
		}

		id, _ := strconv.Atoi(strings.TrimPrefix(splitLine[0], "Game "))

		game := Game{id, []CubeSet{}}

		lineSets := strings.Split(splitLine[1], "; ")
		for _, lineSet := range lineSets {
			set := CubeSet{0, 0, 0}

			colors := strings.Split(lineSet, ", ")
			for _, color := range colors {
				if strings.HasSuffix(color, "red") {
					count, _ := strconv.Atoi(strings.TrimSuffix(color, " red"))
					set.red = count
					continue
				}

				if strings.HasSuffix(color, "green") {
					count, _ := strconv.Atoi(strings.TrimSuffix(color, " green"))
					set.green = count
					continue
				}

				if strings.HasSuffix(color, "blue") {
					count, _ := strconv.Atoi(strings.TrimSuffix(color, " blue"))
					set.blue = count
					continue
				}
			}

			game.sets = append(game.sets, set)
		}

		games = append(games, game)
	}

	return games
}
