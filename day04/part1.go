package day04

import (
	"slices"
	"sync"
)

func processLine(line string) int {
	game := parseCard(line)

	points := 0

	for _, your := range game.yours {
		if !slices.Contains(game.winning, your) {
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
