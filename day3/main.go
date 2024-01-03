package day3

import (
	"fmt"
	"os"
	"slices"
	"sync"
	"time"
)

func Run() {
	filepath := ""
	if len(os.Args) > 2 {
		filepath = os.Args[2]
	}

	fmt.Println("Loading file", filepath)
	lines := ReadFile(filepath)

	// Part 1
	fmt.Println("Starting part 1...")

	start := time.Now()

	result := part1(lines)

	end := time.Now()
	elapsed := end.Sub(start)

	fmt.Printf("Finished processing! Result: %d, Elapsed time: %s\n", result, elapsed)

	// Part 2
	fmt.Println("Starting part 2...")

	start = time.Now()

	result = part2(lines)

	end = time.Now()
	elapsed = end.Sub(start)

	fmt.Printf("Finished processing! Result: %d, Elapsed time: %s\n", result, elapsed)
}

var symbols = []rune{'.', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func part1(lines []string) int {
	sum := 0

	allNumbers := findAllNumbers(lines)

	count := 0
	for y, line := range lines {
		for x, char := range line {
			if !slices.Contains(symbols, char) {
				sum += sumAdjacent(lines, allNumbers, y, x)
				count++
			}
		}
	}

	fmt.Printf("count: %v\n", count)

	return sum
}

func part2(lines []string) int {
	sum := 0

	allNumbers := findAllNumbers(lines)

	count := 0
	for y, line := range lines {
		for x, char := range line {
			if char == '*' {
				sum += calcGear(lines, allNumbers, y, x)
				count++
			}
		}
	}

	fmt.Printf("count: %v\n", count)

	return sum
}

func findAllNumbers(lines []string) [][]Number {
	allNumbers := [][]Number{}

	var wg sync.WaitGroup
	wg.Add(len(lines))

	for index, line := range lines {
		allNumbers = append(allNumbers, nil)

		go func(index int, line string) {
			allNumbers[index] = findLineNumbers(line)
			wg.Done()
		}(index, line)
	}

	wg.Wait()

	return allNumbers
}

func findLineNumbers(line string) []Number {
	lineNumbers := []Number{}

	startIndex := -1
	lastValue := 0
	for x, char := range line {
		value := int(char) - 48

		if value < 0 || value > 9 {
			startIndex = -1
			lastValue = 0

			number := Number{0, x, 1}
			lineNumbers = append(lineNumbers, number)

			continue
		}

		if startIndex == -1 {
			startIndex = x
		} else {
			lastValue *= 10
		}

		lastValue += value

		number := Number{lastValue, startIndex, x - startIndex + 1}
		lineNumbers = append(lineNumbers, number)

		for i := startIndex; i < x; i++ {
			lineNumbers[i] = number
		}
	}

	return lineNumbers
}

func getBoundaries(lines []string, row int, col int) (int, int, int, int) {
	startRow := 0
	if row > 0 {
		startRow = row - 1
	}
	endRow := len(lines[row])
	if row < endRow {
		endRow = row + 1
	}
	startCol := 0
	if col > 0 {
		startCol = col - 1
	}
	endCol := len(lines[row])
	if col < endCol {
		endCol = col + 1
	}

	return startRow, endRow, startCol, endCol
}

func sumAdjacent(lines []string, allNumbers [][]Number, row int, col int) int {
	sum := 0

	startRow, endRow, startCol, endCol := getBoundaries(lines, row, col)

	for y := startRow; y <= endRow; y++ {
		for x := startCol; x <= endCol; x++ {
			number := allNumbers[y][x]
			// fmt.Printf("y: %v x: %v number: %v\n", y, x, number)

			sum += number.value
			x += number.index - x + number.length - 1
		}
	}

	fmt.Printf("row: %v col: %v sum: %v startRow: %v endRow: %v startCol: %v endCol: %v\n", row, col, sum, startRow, endRow, startCol, endCol)

	return sum
}

func calcGear(lines []string, allNumbers [][]Number, row int, col int) int {
	product := 1

	startRow, endRow, startCol, endCol := getBoundaries(lines, row, col)

	adjacents := 0
	for y := startRow; y <= endRow; y++ {
		for x := startCol; x <= endCol; x++ {
			number := allNumbers[y][x]
			// fmt.Printf("y: %v x: %v number: %v\n", y, x, number)

			if number.value == 0 {
				continue
			}

			product *= number.value
			x += number.index - x + number.length - 1
			adjacents++
		}
	}

	fmt.Printf("row: %v col: %v product: %v adjacents: %v startRow: %v endRow: %v startCol: %v endCol: %v\n", row, col, product, adjacents, startRow, endRow, startCol, endCol)

	if adjacents != 2 {
		return 0
	}

	return product
}
