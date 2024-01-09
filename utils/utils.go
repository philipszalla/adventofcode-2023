package utils

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type partFn func([]string) int

func RunPart(fn partFn, part int, lines []string) {
	fmt.Printf("Starting part %d...\n", part)

	start := time.Now()

	result := fn(lines)

	end := time.Now()
	elapsed := end.Sub(start)

	fmt.Printf("Finished processing! Result: %d, Elapsed time: %s\n", result, elapsed)
}

func ReadFile(filename string) []string {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	content := string(bytes)
	content = strings.ReplaceAll(content, "\r", "")
	lines := strings.Split(content, "\n")

	return lines
}
