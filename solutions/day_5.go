/*
https://adventofcode.com/2024/day/5
*/
package solutions

import (
	"aoc/internal"
	"slices"
	"strconv"
	"strings"
)

// soluce : 61, 53, 29 = 143

func getInput() ([][2]int, [][]int) {
	data, _ := internal.ReadFileLinesWithBlanks("resources/day_5")
	rules := make([][2]int, 0, 1024)
	updates := make([][]int, 0, 512)
	for i := range updates {
		updates[i] = make([]int, 0, 32)
	}
	rulePart := true
	for _, line := range data {
		if len(line) == 0 {
			rulePart = false
			continue
		}
		if rulePart {
			numbers := strings.Split(line, "|")
			ln, _ := strconv.Atoi(numbers[0])
			rn, _ := strconv.Atoi(numbers[1])
			rules = append(rules, [2]int{ln, rn})
		} else {
			numbers := strings.Split(line, ",")
			nList := make([]int, 0, 32)
			for _, n := range numbers {
				tmpNumber, _ := strconv.Atoi(n)
				nList = append(nList, tmpNumber)
			}
			updates = append(updates, nList)
		}
	}
	return rules, updates
}

func getRulesTree(rules [][2]int) map[int][]int {
	rulesTree := make(map[int][]int, 32)
	for _, rule := range rules {
		_, ok := rulesTree[rule[0]]
		if !ok {
			rulesTree[rule[0]] = []int{rule[1]}
		} else {
			rulesTree[rule[0]] = append(rulesTree[rule[0]], rule[1])
		}
	}
	return rulesTree
}

func getCorrectUpdates(rulesTree map[int][]int, updates [][]int) []int {
	correctUpdate := []int{}
	for ui, update := range updates {
		i := 0
		for i = 1; i < len(update); i++ {
			idx := slices.Index(rulesTree[update[i-1]], update[i])
			if idx == -1 {
				break
			}
		}
		if i == len(update) {
			correctUpdate = append(correctUpdate, ui)
		}
	}
	return correctUpdate
}

func getSumMidNumber(updates [][]int) int {
	total := 0
	for _, update := range updates {
		total += update[len(update)/2]
	}
	return total
}

func Day5_1() int {
	rules, updates := getInput()
	rulesTree := getRulesTree(rules)
	correctUpdatesIndex := getCorrectUpdates(rulesTree, updates)
	correctUpdates := make([][]int, 0, len(updates))
	for _, k := range correctUpdatesIndex {
		correctUpdates = append(correctUpdates, updates[k])
	}
	sumMidNumber := getSumMidNumber(correctUpdates)
	return sumMidNumber
}

func getIncorrectUpdates(rulesTree map[int][]int, updates [][]int) []int {
	correctUpdate := []int{}
	for ui, update := range updates {
		i := 0
		for i = 1; i < len(update); i++ {
			idx := slices.Index(rulesTree[update[i-1]], update[i])
			if idx == -1 {
				correctUpdate = append(correctUpdate, ui)
				break
			}
		}
	}
	return correctUpdate
}

func sortIncorrectUpdates(rulesTree map[int][]int, updates [][]int, incorrectUpdatesIndex []int) [][]int {
	correctUpdates := make([][]int, len(incorrectUpdatesIndex), len(incorrectUpdatesIndex))
	for i, ui := range incorrectUpdatesIndex {
		correctUpdates[i] = make([]int, 0, len(updates[i]))
		correctUpdates[i] = append(correctUpdates[i], updates[ui][0])
		for j := 1; j < len(updates[ui]); j++ {
			var iToInsert int
			var lastIdInsert int
			for iToInsert = 0; iToInsert < len(correctUpdates[i]); iToInsert++ {
				n := correctUpdates[i][iToInsert]
				_, ok := rulesTree[n]
				if !ok { // No rules, so place at the end
					lastIdInsert = len(correctUpdates[i]) - 1
					break
				} else {
					idx := slices.Index(rulesTree[n], updates[ui][j])
					if idx == -1 { // no path to the next, go further
						continue
					} else { // We found a path, continue if we found an other one more recent
						lastIdInsert = iToInsert + 1
						continue
					}
				}
			}
			correctUpdates[i] = slices.Insert(correctUpdates[i], lastIdInsert, updates[ui][j])
		}
	}
	return correctUpdates
}

func Day5_2() int {
	rules, updates := getInput()
	rulesTree := getRulesTree(rules)
	incorrectUpdatesIndex := getIncorrectUpdates(rulesTree, updates)
	correctedUpdates := sortIncorrectUpdates(rulesTree, updates, incorrectUpdatesIndex)
	sumMidNumber := getSumMidNumber(correctedUpdates)
	return sumMidNumber
}
