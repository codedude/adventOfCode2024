/*
https://adventofcode.com/2024/day/11
*/
package solutions

import (
	"aoc/internal"
	"strconv"
	"strings"
)

func getInput11() []int {
	data, _ := internal.ReadFileLines("resources/day_11")
	input := strings.Split(data[0], " ")
	inputList := make([]int, len(input))
	for i, v := range input {
		inputList[i], _ = strconv.Atoi(v)
	}
	return inputList
}

func hasEvenDigits(n int) bool {
	nStr := strconv.Itoa(n)
	return len(nStr)%2 == 0
}

func splitDigits(n int) ([]int, bool) {
	nStr := strconv.Itoa(n)
	if len(nStr)%2 != 0 {
		return nil, false
	}
	leftPart, _ := strconv.Atoi(nStr[:len(nStr)/2])
	rightPart, _ := strconv.Atoi(nStr[len(nStr)/2:])
	return []int{leftPart, rightPart}, true
}

// naive solve, use too much memory
func solve11(data []int, part2 bool) int {
	currentStones := make([]int, len(data), 1000000)
	newStones := make([]int, 0, 1000000)
	copy(currentStones, data)
	var blinkTimes int
	if part2 {
		blinkTimes = 75 // way too much memory
	} else {
		blinkTimes = 25
	}
	for round := 0; round < blinkTimes; round++ {
		for i := 0; i < len(currentStones); i++ {
			if currentStones[i] == 0 {
				newStones = append(newStones, 1)
			} else if hasEvenDigits(currentStones[i]) {

			} else {
				newStones = append(newStones, currentStones[i]*2024)
			}
		}
		currentStones = currentStones[:0]
		currentStones = append(currentStones, newStones...)
		newStones = newStones[:0]
	}
	return len(currentStones)
}

func recCalcStone(input int, tree map[int][]int, step int, cache map[int64]int) int {
	if input == -1 {
		return 0
	}
	if step == 0 {
		return 1
	}
	if stonesCount, ok := cache[(int64(input)<<8)|int64(step)]; ok {
		return stonesCount
	} else {
		r := recCalcStone(tree[input][0], tree, step-1, cache) + recCalcStone(tree[input][1], tree, step-1, cache)
		cache[(int64(input)<<8)|int64(step)] = r
		return r
	}
}

func solve11_2(data []int, part2 bool) int {
	var blinkTimes, stonesSize int
	if part2 {
		blinkTimes = 75
		stonesSize = 64000
	} else {
		blinkTimes = 25
		stonesSize = 32000
	}
	tree := make(map[int][]int, 1024)
	stones := make([]int, len(data), stonesSize)
	copy(stones, data)
	for round := 0; round < blinkTimes; round++ {
		endOfCurrentStone := len(stones)
		for i := 0; i < endOfCurrentStone; i++ {
			currentStone := stones[i]
			if _, ok := tree[currentStone]; !ok {
				tree[currentStone] = []int{-1, -1}
				if currentStone == 0 {
					tree[currentStone] = []int{1, -1}
					stones[i] = 1
				} else if parts, ok := splitDigits(currentStone); ok {
					tree[currentStone] = parts
					stones = append(stones, parts[0])
					stones = append(stones, parts[1])
				} else {
					tree[currentStone] = []int{currentStone * 2024, -1}
					stones[i] = currentStone * 2024
				}
			}
		}
	}
	totalStone := 0

	cache := make(map[int64]int, 1024)
	for _, v := range data {
		totalStone += recCalcStone(v, tree, blinkTimes, cache)
	}

	return totalStone
}

func Day11_1() int {
	data := getInput11()
	soluce := solve11_2(data, false)
	return soluce
}

func Day11_2() int {
	data := getInput11()
	soluce := solve11_2(data, true)
	return soluce
}
