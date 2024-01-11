package day06

import (
	"regexp"
	"strconv"
)

func parse(lines []string) ([]int, []int) {
	if len(lines) != 2 {
		return []int{}, []int{}
	}
	numberRegexp := regexp.MustCompile(`\d+`)

	timeStrings := numberRegexp.FindAllString(lines[0], -1)
	times := []int{}
	for _, timeString := range timeStrings {
		time, err := strconv.Atoi(timeString)
		if err != nil {
			continue
		}

		times = append(times, time)
	}

	distanceStrings := numberRegexp.FindAllString(lines[1], -1)
	distances := []int{}
	for _, distanceString := range distanceStrings {
		distance, err := strconv.Atoi(distanceString)
		if err != nil {
			continue
		}

		distances = append(distances, distance)
	}

	return times, distances
}
