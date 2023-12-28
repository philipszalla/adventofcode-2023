package day1

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

	return strings.Split(content, "\n")
}
