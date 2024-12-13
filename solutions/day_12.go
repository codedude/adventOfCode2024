/*
https://adventofcode.com/2024/day/12
*/
package solutions

import (
	"aoc/internal"
	"fmt"
)

type Fence struct {
	Dir  WalkDirection // reuse
	I, J int           // position of the region
}

type Region struct {
	Area       int
	Perimeter  int
	EntryPoint [2]int
	Sides      map[int]bool
}

func getInput12() [][]int {
	data, _ := internal.ReadFileLines("resources/day_12")
	lines := make([][]int, len(data))
	for i := range data {
		lines[i] = make([]int, len(data[i]))
		for j := range lines[i] {
			lines[i][j] = int(data[i][j])
		}
	}
	return lines
}

func getFencesPart2(data [][]int, i, j int) int {
	region := data[i][j]
	fences := 0
	maxI := len(data) - 1
	maxJ := len(data[0]) - 1
	// 1 : test if the top one is not same region, then check top-left one if it is, then left one that is not
	// test top
	if (i == 0 || data[i-1][j] != region) && (j == 0 || (i != 0 && data[i-1][j-1] == region || data[i][j-1] != region)) {
		fences += 1
	}
	// test bottom
	if (i == maxI || data[i+1][j] != region) && (j == maxJ || (i != maxI && data[i+1][j+1] == region || data[i][j+1] != region)) {
		fences += 1
	}
	// test left
	if (j == 0 || data[i][j-1] != region) && (i == maxI || (j != 0 && data[i+1][j-1] == region || data[i+1][j] != region)) {
		fences += 1
	}
	// test right
	if (j == maxJ || data[i][j+1] != region) && (i == 0 || (j != maxJ && data[i-1][j+1] == region || data[i-1][j] != region)) {
		fences += 1
	}
	return fences
}

// regions can be duplicated, so set them unique
func markRegions(data [][]int, part2 bool) ([][]int, map[int]*Region) {
	stats := make(map[int]*Region, 26)
	markData := make([][]int, len(data))
	for i := range markData {
		markData[i] = make([]int, len(data[i]))
		for j := range markData[i] {
			markData[i][j] = -1
		}
	}
	id := 0
	for i := range data {
		for j := range data[i] {
			if markData[i][j] != -1 { // already marked
				continue
			}
			stack := Stack[[2]int]{}
			stack.Push([2]int{i, j})
			stats[id] = &Region{Area: 1, Perimeter: 0, Sides: make(map[int]bool), EntryPoint: [2]int{i, j}}
			markData[i][j] = id
			for {
				e, err := stack.Pop()
				if err != nil {
					break
				}
				neighboors := getNeighboorsNotMarked(data, e[0], e[1], markData)
				nFences := 4 - len(neighboors)
				if part2 {
					stats[id].Perimeter += getFencesPart2(data, e[0], e[1])
				} else {
					stats[id].Perimeter += nFences
				}
				for _, n := range neighboors {
					if n[0] == -1 {
						continue
					}
					stats[id].Area += 1
					markData[n[0]][n[1]] = id
					stack.Push(n)
				}
			}
			id += 1 // new id for the next reigon
		}
	}
	return markData, stats
}

// Return neighboors of same region not yet visited
func getNeighboorsNotMarked(data [][]int, i, j int, markData [][]int) [][2]int {
	region := data[i][j]
	neighboors := make([][2]int, 0, 4)
	if i > 0 {
		if data[i-1][j] == region {
			if markData[i-1][j] == -1 {
				neighboors = append(neighboors, [2]int{i - 1, j})
			} else {
				neighboors = append(neighboors, [2]int{-1, -1})
			}
		}
	}
	if i < len(data)-1 {
		if data[i+1][j] == region {
			if markData[i+1][j] == -1 {
				neighboors = append(neighboors, [2]int{i + 1, j})
			} else {
				neighboors = append(neighboors, [2]int{-1, -1})
			}
		}
	}
	if j > 0 {
		if data[i][j-1] == region {
			if markData[i][j-1] == -1 {
				neighboors = append(neighboors, [2]int{i, j - 1})
			} else {
				neighboors = append(neighboors, [2]int{-1, -1})
			}
		}
	}
	if j < len(data[i])-1 {
		if data[i][j+1] == region {
			if markData[i][j+1] == -1 {
				neighboors = append(neighboors, [2]int{i, j + 1})
			} else {
				neighboors = append(neighboors, [2]int{-1, -1})
			}
		}
	}
	return neighboors
}

func solve12(stats map[int]*Region) int {
	price := 0
	for k, stat := range stats {
		fmt.Println(k, stat.Perimeter)
		price += stat.Area * stat.Perimeter
	}
	return price
}

func Day12_1() int {
	data := getInput12()
	data, stats := markRegions(data, false) // set unique id for duplicates regions not adjacent
	soluce := solve12(stats)
	return soluce
}

func Day12_2() int {
	data := getInput12()
	data, stats := markRegions(data, true) // set unique id for duplicates regions not adjacent
	soluce := solve12(stats)
	return soluce
}
