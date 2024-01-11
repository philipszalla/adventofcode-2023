package day04

import (
	"slices"
	"sync"
)

func processCard(cards []Card, card Card, c chan int, wg *sync.WaitGroup) {
	c <- 1

	count := 0

	for _, your := range card.yours {
		if !slices.Contains(card.winning, your) {
			continue
		}

		count++
	}

	for i := card.id; i < card.id+count && i < len(cards); i++ {
		wg.Add(1)
		go processCard(cards, cards[i], c, wg)
	}

	wg.Done()
}

func part2(lines []string) int {
	cards := []Card{}
	for _, line := range lines {
		cards = append(cards, parseCard(line))
	}

	c := make(chan int)

	var wg sync.WaitGroup

	for _, card := range cards {
		wg.Add(1)
		go processCard(cards, card, c, &wg)
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
