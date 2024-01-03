package day3

import (
	"os"
	"strings"
)

func ReadFile(filename string) []string {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	content := string(bytes)
	content = strings.ReplaceAll(content, "\r", "")

	games := parse(content)

	return games
}

func parse(content string) []string {
	lines := strings.Split(content, "\n")

	return lines
}
