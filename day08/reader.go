package day08

import "strings"

func parse(lines []string) ([]rune, []string, []string, []string) {
	schema := []rune{}
	starts, destsLeft, destsRight := []string{}, []string{}, []string{}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, " = ")
		if len(parts) == 1 {
			schema = []rune(parts[0])
			continue
		}

		start := parts[0]
		parts = strings.Split(parts[1], ", ")
		destLeft := strings.TrimPrefix(parts[0], "(")
		destRight := strings.TrimSuffix(parts[1], ")")

		starts = append(starts, start)
		destsLeft = append(destsLeft, destLeft)
		destsRight = append(destsRight, destRight)
	}

	return schema, starts, destsLeft, destsRight
}
