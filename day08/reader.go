package day08

import "strings"

func parse(lines []string) ([]rune, MapDict) {
	schema := []rune(lines[0])
	mapDict := make(MapDict, len(lines)-2)

	for _, line := range lines[2:] {
		parts := strings.Split(line, " = ")

		start := parts[0]
		parts = strings.Split(parts[1], ", ")
		destLeft := strings.TrimPrefix(parts[0], "(")
		destRight := strings.TrimSuffix(parts[1], ")")

		mapDict[start] = Dir{destLeft, destRight}
	}

	return schema, mapDict
}
