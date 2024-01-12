package day09

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/philipszalla/adventofcode-2023/utils"
)

func Run() {
	filepath := ""
	if len(os.Args) > 2 {
		filepath = os.Args[2]
	}

	fmt.Println("Loading file", filepath)
	lines := utils.ReadFile(filepath)

	utils.RunPart(part1, 1, lines)
}

func getNextLayer(layer []int) ([]int, bool) {
	end := true
	nextLayer := []int{}

	for index, b := range layer[1:] {
		a := layer[index]
		diff := b - a

		nextLayer = append(nextLayer, diff)
		end = end && diff == 0
	}

	return nextLayer, end
}

func parseLine(line string) [][]int {
	pyramid := [][]int{}
	pyramid = append(pyramid, []int{})

	// Get first layer
	parts := strings.Split(line, " ")
	for _, part := range parts {
		number, err := strconv.Atoi(part)
		if err != nil {
			continue
		}

		pyramid[0] = append(pyramid[0], number)
	}

	return pyramid
}

func buildPyramid(pyramid [][]int) [][]int {
	for {
		layer := pyramid[len(pyramid)-1]
		nextLayer, end := getNextLayer(layer)

		pyramid = append(pyramid, nextLayer)

		if end {
			break
		}
	}

	return pyramid
}

func extrapolatePyramid(pyramid [][]int) [][]int {
	pyramid[len(pyramid)-1] = append(pyramid[len(pyramid)-1], 0)
	for index := len(pyramid) - 2; index >= 0; index-- {
		prevLayer := pyramid[index+1]
		lastDiff := prevLayer[len(prevLayer)-1]

		layer := pyramid[index]
		lastValue := layer[len(layer)-1]

		layer = append(layer, lastValue+lastDiff)

		pyramid[index] = layer
	}

	return pyramid
}

func processLine(line string) int {
	pyramid := parseLine(line)

	pyramid = buildPyramid(pyramid)

	pyramid = extrapolatePyramid(pyramid)

	return pyramid[0][len(pyramid[0])-1]
}

func processAllLines(lines []string, fn func(string) int) int {
	sum := 0

	c := make(chan int, len(lines))
	for _, line := range lines {
		go func(line string) {
			c <- processLine(line)
		}(line)
	}

	for index := 0; index < len(lines); index++ {
		sum += <-c
	}

	return sum
}

func part1(lines []string) int {
	sum := processAllLines(lines, processLine)

	return sum
}
