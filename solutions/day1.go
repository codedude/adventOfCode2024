/*
Day 1: https://adventofcode.com/2024/day/1
*/
package solutions

import (
	"aoc/internal"
	"slices"
	"strconv"
	"strings"
)

func Day1_1() int {
	data, _ := internal.ReadFileLines("resources/day_1_1")
	columns := [2][]int{}
	for _, v := range data {
		numbers := strings.Split(v, "   ")
		left, _ := strconv.Atoi(numbers[0])
		right, _ := strconv.Atoi(numbers[1])
		columns[0] = append(columns[0], left)
		columns[1] = append(columns[1], right)
	}

	slices.Sort(columns[0])
	slices.Sort(columns[1])
	total := 0
	for i := range len(data) {
		total += internal.AbsInt(columns[0][i] - columns[1][i])
	}

	return total
}

func Day1_2() int {
	data, _ := internal.ReadFileLines("resources/day_1")
	occurrences := make(map[int]int, len(data))
	leftColumn := []int{}
	for _, v := range data {
		numbers := strings.Split(v, "   ")
		left, _ := strconv.Atoi(numbers[0])
		right, _ := strconv.Atoi(numbers[1])

		leftColumn = append(leftColumn, left)
		v, ok := occurrences[right]
		if !ok {
			occurrences[right] = 1
		} else {
			occurrences[right] = v + 1
		}
	}
	total := 0
	for _, v := range leftColumn {
		occurs, ok := occurrences[v]
		if !ok {
			continue
		}
		total += v * occurs
	}

	return total
}
