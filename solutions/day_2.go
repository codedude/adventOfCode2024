/*
Day 2: https://adventofcode.com/2024/day/2
*/
package solutions

import (
	"aoc/internal"
	"strconv"
	"strings"
)

func isReportSafe(levels []int) bool {
	orderAsc := levels[0]-levels[1] < 0 // >0 = DESC, <0 = ASC
	for i := 1; i < len(levels); i++ {
		diff := levels[i-1] - levels[i]
		absDiff := internal.AbsInt(diff)
		if absDiff < 1 || absDiff > 3 {
			return false
		}
		if (orderAsc && diff > 0) || (!orderAsc && diff < 0) {
			return false
		}
	}
	return true
}

func Day2_1() int {
	data, _ := internal.ReadFileLines("resources/day_2")

	safeReports := 0
	for _, report := range data {
		levels := strings.Split(report, " ")
		numbers := []int{}
		for _, v := range levels {
			i, _ := strconv.Atoi(v)
			numbers = append(numbers, i)
		}
		if isReportSafe(numbers) {
			safeReports += 1
		}
	}

	return safeReports
}

func rmISlice(slice []int, i int) []int {
	sliceCopy := make([]int, len(slice))
	copy(sliceCopy, slice)
	return append(sliceCopy[:i], sliceCopy[i+1:]...)
}

func isReportSafeDampener(levels []int, errors int) bool {
	lastOrderAsc := levels[0]-levels[1] < 0 // >0 = DESC, <0 = ASC
	for i := 1; i < len(levels); i++ {
		diff := levels[i-1] - levels[i]
		absDiff := internal.AbsInt(diff)
		if absDiff < 1 || absDiff > 3 {
			if errors == 1 {
				return false
			}
			ret := isReportSafeDampener(rmISlice(levels, i), errors+1) || isReportSafeDampener(rmISlice(levels, i-1), errors+1)
			return ret
		} else if (lastOrderAsc && diff > 0) || (!lastOrderAsc && diff < 0) {
			if errors == 1 {
				return false
			}
			ret := isReportSafeDampener(rmISlice(levels, i), errors+1) || isReportSafeDampener(rmISlice(levels, i-1), errors+1)
			if i > 1 {
				ret = ret || isReportSafeDampener(rmISlice(levels, i-2), errors+1)
			}
			return ret
		}
		lastOrderAsc = diff < 0
	}
	return true
}

func Day2_2() int {
	data, _ := internal.ReadFileLines("resources/day_2")

	safeReports := 0
	for _, report := range data {
		levels := strings.Split(report, " ")
		numbers := []int{}
		for _, v := range levels {
			i, _ := strconv.Atoi(v)
			numbers = append(numbers, i)
		}
		if isReportSafeDampener(numbers, 0) {
			safeReports += 1
		}
	}

	return safeReports
}
