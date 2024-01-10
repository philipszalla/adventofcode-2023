package day4

import (
	"slices"
	"strings"
	"sync"
)

func processLine(line string) int {
	if !strings.HasPrefix(line, "Card ") {
		return 0
	}

	_, winning, yours := parseLine(line)

	points := 0

	for _, your := range yours {
		if !slices.Contains(winning, your) {
			continue
		}

		if points != 0 {
			points *= 2
		} else {
			points = 1
		}
	}

	return points
}

func part1(lines []string) int {
	c := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(lines))

	for _, line := range lines {
		go func(line string) {
			c <- processLine(line)
			wg.Done()
		}(line)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	sum := 0
	for points := range c {
		sum += points
	}

	return sum
}
