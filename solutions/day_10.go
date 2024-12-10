/*
https://adventofcode.com/2024/day/10
*/
package solutions

import (
	"aoc/internal"
	"errors"
	"fmt"
)

type Stack[T any] struct {
	Items []T
}

func (s *Stack[T]) Push(e T) {
	s.Items = append(s.Items, e)
}

func (s *Stack[T]) Pop() (T, error) {
	var e T
	if len(s.Items) == 0 {
		return e, errors.New("Empty stack")
	}
	e = s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return e, nil
}

func getInput10() ([][]int, [][2]int) {
	data, _ := internal.ReadFileLines("resources/day_10")
	input := make([][]int, len(data))
	startPoints := make([][2]int, 0)
	for i := range data {
		input[i] = make([]int, len(data[i]))
		for j := range data[0] {
			value := int(data[i][j]) - 48 // byte 0-9 only
			input[i][j] = value
			if value == 0 {
				startPoints = append(startPoints, [2]int{i, j})
			}
		}
	}
	return input, startPoints
}

func solve10(data [][]int, startPoints [][2]int, part2 bool) int {
	topHeights := make(map[int]int, len(startPoints))
	for _, start := range startPoints {
		visited := make(map[int]bool, len(startPoints))
		stack := Stack[[3]int]{}
		stack.Push([3]int{start[0], start[1], 0})
		for {
			e, err := stack.Pop()
			if err != nil {
				break
			}
			if data[e[0]][e[1]] == 9 {
				if !part2 {
					_, ok := visited[mapKey(e[0], e[1])]
					if !ok {
						visited[mapKey(e[0], e[1])] = true
					} else {
						continue
					}
				}
				v, ok := topHeights[mapKey(start[0], start[1])]
				if !ok {
					topHeights[mapKey(start[0], start[1])] = 1
				} else {
					topHeights[mapKey(start[0], start[1])] = v + 1
				}
				continue
			}
			step := e[2] + 1
			if e[0] > 0 && data[e[0]-1][e[1]] == step {
				stack.Push([3]int{e[0] - 1, e[1], step})
			}
			if e[0] < len(data)-1 && data[e[0]+1][e[1]] == step {
				stack.Push([3]int{e[0] + 1, e[1], step})
			}
			if e[1] > 0 && data[e[0]][e[1]-1] == step {
				stack.Push([3]int{e[0], e[1] - 1, step})
			}
			if e[1] < len(data[0])-1 && data[e[0]][e[1]+1] == step {
				stack.Push([3]int{e[0], e[1] + 1, step})
			}
		}
	}
	total := 0
	for k, tops := range topHeights {
		fmt.Println(k>>16, k&0xff, tops)
		total += tops
	}
	return total
}

func Day10_1() int {
	data, startPoints := getInput10()
	soluce := solve10(data, startPoints, false)
	return soluce
}

func Day10_2() int {
	data, startPoints := getInput10()
	soluce := solve10(data, startPoints, true)
	return soluce
}
