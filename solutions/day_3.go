/*
Day 3: https://adventofcode.com/2024/day/3
*/
package solutions

import (
	"aoc/internal"
	"fmt"
	"regexp"
	"strconv"
)

func calcMul(data []string) (int, error) {
	r, err := regexp.Compile(`mul\((?P<Left>[0-9]{1,3}),(?P<Right>[0-9]{1,3})\)`)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	total := 0
	for _, line := range data {
		result := r.FindAllStringSubmatch(line, -1)
		if result == nil {
			continue
		}
		for _, mul := range result {
			left, _ := strconv.Atoi(mul[1])
			right, _ := strconv.Atoi(mul[2])
			total += left * right
		}
	}

	return total, nil
}

func Day3_1() int {
	data, _ := internal.ReadFileLines("resources/day_3")
	total, _ := calcMul(data)
	return total
}

func Day3_2() int {
	data, _ := internal.ReadFileLines("resources/day_3")

	rCleanInput, err := regexp.Compile(`(do\(\)|don't\(\)|mul\([0-9]{1,3},[0-9]{1,3}\))`)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	cleanData := make([]string, 512)
	do := true
	for _, line := range data {
		result := rCleanInput.FindAllString(line, -1)
		if result == nil {
			continue
		}
		for _, s := range result {
			if s == "do()" {
				do = true
			} else if s == "don't()" {
				do = false
			} else {
				if !do {
					continue
				} else {
					cleanData = append(cleanData, s)
				}
			}
		}
	}
	total, _ := calcMul(cleanData)
	return total
}
