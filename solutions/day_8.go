/*
https://adventofcode.com/2024/day/8
*/
package solutions

import (
	"aoc/internal"
)

func mapKey(i, j int) int {
	return (i << 16) | j
}

func getInput8() (map[byte][][2]int, int, int) {
	data, _ := internal.ReadFileLines("resources/day_8")
	dataMap := make(map[byte][][2]int, 32) // map[freq] = antenna pos list [[i,j], ...]
	for i := range data {
		for j := range data[i] {
			v := data[i][j]
			if v == '.' {
				continue
			}
			newValue := [2]int{i, j}
			_, ok := dataMap[v]
			if !ok {
				dataMap[v] = [][2]int{newValue}
			} else {
				dataMap[v] = append(dataMap[v], newValue)
			}
		}
	}
	return dataMap, len(data), len(data[0])
}

// Find the slop between two point, topmost point to bottom most, correct values if needed
func getSlope(ant1, ant2 [2]int) [2]int {
	slope := [2]int{-(ant1[0] - ant2[0]), -(ant1[1] - ant2[1])}
	return slope
}

func isAntinodesInMap(ant [2]int, width, height int) bool {
	return !(ant[0] < 0 || ant[1] < 0 || ant[0] >= height || ant[1] >= width)
}

func getAntinodes(ant1, ant2 [2]int, slope [2]int, width, height int, part2 bool) [][2]int {
	var antinodes [][2]int
	bothNodeOut := false
	newNode1, newNode2 := [2]int{ant1[0], ant1[1]}, [2]int{ant2[0], ant2[1]}
	if part2 {
		antinodes = append(antinodes, [2]int{ant1[0], ant1[1]})
		antinodes = append(antinodes, [2]int{ant2[0], ant2[1]})
	}
	for {
		newNode1 = [2]int{newNode1[0] - slope[0], newNode1[1] - slope[1]}
		if isAntinodesInMap(newNode1, width, height) {
			antinodes = append(antinodes, newNode1)
		} else {
			bothNodeOut = true
		}
		newNode2 = [2]int{newNode2[0] + slope[0], newNode2[1] + slope[1]}
		if isAntinodesInMap(newNode2, width, height) {
			antinodes = append(antinodes, newNode2)
			bothNodeOut = false
		}
		if bothNodeOut || !part2 {
			break // only 1 turn for part1
		}
	}
	return antinodes
}

func solve81(data map[byte][][2]int, height, width int, part2 bool) int {
	antinodes := make(map[int]bool, 32)
	for _, ant := range data { // Iterate over each frequencies
		for i := 0; i < len(ant); i++ { //Test each combination of antennas
			for j := i + 1; j < len(ant); j++ { // i and j are 2 antennas
				slope := getSlope(ant[i], ant[j])
				if slope[0] == 0 || slope[1] == 0 { // Antennas are aligned
					continue
				}
				tmpAntinodes := getAntinodes(ant[i], ant[j], slope, width, height, part2)
				for _, ant := range tmpAntinodes {
					antinodes[mapKey(ant[0], ant[1])] = true
				}
			}
		}
	}
	return len(antinodes)
}

func Day8_1() int {
	data, height, width := getInput8()
	soluce := solve81(data, height, width, false)
	return soluce
}

func Day8_2() int {
	data, height, width := getInput8()
	soluce := solve81(data, height, width, true)
	return soluce
}
