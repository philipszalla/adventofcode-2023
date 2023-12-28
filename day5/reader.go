package day5

import (
	"os"
	"strconv"
	"strings"
)

func ReadFile(filename string) ([]int, []Mapper) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	content := string(bytes)
	content = strings.ReplaceAll(content, "\r", "")

	seeds, maps := parse(content)

	return seeds, maps
}

func parse(content string) ([]int, []Mapper) {
	lines := strings.Split(content, "\n")

	seeds := []int{}
	maps := []Mapper{}

	currentMapIndex := -1

	for _, line := range lines {
		if len(line) == 0 {
			currentMapIndex = -1
			continue
		}

		if strings.HasPrefix(line, "seeds: ") {
			seedsString := strings.Split(strings.TrimPrefix(line, "seeds: "), " ")

			for _, seedString := range seedsString {
				seed, _ := strconv.Atoi(seedString)

				seeds = append(seeds, seed)
			}

			continue
		}

		if currentMapIndex != -1 {
			ruleParts := strings.Split(line, " ")
			if len(ruleParts) != 3 {
				panic("Map rule must contain 3 numbers!")
			}

			currentMap := maps[currentMapIndex]

			destinationStart, _ := strconv.Atoi(ruleParts[0])
			sourceStart, _ := strconv.Atoi(ruleParts[1])
			length, _ := strconv.Atoi(ruleParts[2])
			rule := MapperRule{
				destinationStart: destinationStart,
				sourceStart:      sourceStart,
				length:           length,
			}

			currentMap.rules = append(currentMap.rules, rule)

			maps[currentMapIndex] = currentMap

			continue
		}

		if strings.HasSuffix(line, " map:") {
			mapName := strings.TrimSuffix(line, " map:")

			currentMap := Mapper{
				name:  mapName,
				rules: []MapperRule{},
			}

			maps = append(maps, currentMap)
			currentMapIndex = len(maps) - 1

			continue
		}
	}

	return seeds, maps
}
