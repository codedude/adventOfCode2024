/*
https://adventofcode.com/2024/day/7
*/
package solutions

import (
	"aoc/internal"
	"strconv"
	"strings"
)

// soluce : 11387

type Operator int

const (
	ADD Operator = iota
	MUL
	CONCAT
)

func getInput7() map[int][]int {
	data, _ := internal.ReadFileLines("resources/day_7")
	dataList := make(map[int][]int, len(data))
	for _, equation := range data {
		eqSplit := strings.Split(equation, ": ")
		eqNumbers := strings.Split(eqSplit[1], " ")
		mapKey, _ := strconv.Atoi(eqSplit[0])
		dataList[mapKey] = make([]int, len(eqNumbers))
		for j, n := range eqNumbers {
			dataList[mapKey][j], _ = strconv.Atoi(n)
		}
	}
	return dataList
}

func recurseEq1(equation []int, result int, pos int, operator Operator, acc int) bool {
	if pos == len(equation) {
		return acc == result
	}
	if operator == ADD {
		acc = acc + equation[pos]
	} else if operator == MUL {
		acc = acc * equation[pos]
	}
	return recurseEq1(equation, result, pos+1, ADD, acc) || recurseEq1(equation, result, pos+1, MUL, acc)
}

func solve71(equations map[int][]int) []int {
	goodEquations := []int{}
	for result, eq := range equations {
		isGoodEq := recurseEq1(eq, result, 1, ADD, eq[0]) || recurseEq1(eq, result, 1, MUL, eq[0])
		if isGoodEq {
			goodEquations = append(goodEquations, result)
		}
	}
	return goodEquations
}

func recurseEq2(equation []int, result int, pos int, operator Operator, acc int) bool {
	if pos == len(equation) {
		return acc == result
	}
	if operator == ADD {
		acc = acc + equation[pos]
	} else if operator == MUL {
		acc = acc * equation[pos]
	} else if operator == CONCAT {
		lns := strconv.Itoa(acc)
		rns := strconv.Itoa(equation[pos])
		nstr := lns + rns
		nint, _ := strconv.Atoi(nstr)
		acc = nint
	}
	return recurseEq2(equation, result, pos+1, ADD, acc) || recurseEq2(equation, result, pos+1, MUL, acc) || recurseEq2(equation, result, pos+1, CONCAT, acc)
}

func solve72(equations map[int][]int) []int {
	goodEquations := []int{}
	for result, eq := range equations {
		isGoodEq := recurseEq2(eq, result, 1, ADD, eq[0]) || recurseEq2(eq, result, 1, MUL, eq[0]) || recurseEq2(eq, result, 1, CONCAT, eq[0])
		if isGoodEq {
			goodEquations = append(goodEquations, result)
		}
	}
	return goodEquations
}

func Day7_1() int {
	data := getInput7()
	goodEquations := solve71(data)
	soluce := 0
	for _, r := range goodEquations {
		soluce += r
	}
	return soluce
}

func Day7_2() int {
	data := getInput7()
	goodEquations := solve72(data)
	soluce := 0
	for _, r := range goodEquations {
		soluce += r
	}
	return soluce
}
