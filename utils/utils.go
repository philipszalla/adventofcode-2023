package utils

import (
	"fmt"
	"os"
	"strings"
	"testing"
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

func TestPart(t *testing.T, fn partFn, content string, expectedResult int) {
	lines := strings.Split(content, "\n")

	result := fn(lines)

	if result != expectedResult {
		t.Fatalf("Expected result is %d. But got %d", expectedResult, result)
	}
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
